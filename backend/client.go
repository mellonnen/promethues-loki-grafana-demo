package main

import (
	"net/http"
	"time"

	"github.com/gorilla/websocket"
	"github.com/sirupsen/logrus"
)

const (
	writeWait = 10 * time.Second

	pongWait = 60 * time.Second

	pingPeriod = (pongWait * 9) / 10

	maxMessageSize = 512
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true }, // allow all
}

type Client struct {
	name   string
	hub    *Hub
	send   chan Msg
	conn   *websocket.Conn
	logger *logrus.Entry
}

func (c *Client) read() {
	defer func() {
		c.hub.unregister <- c
		c.conn.Close()
	}()

	c.conn.SetReadLimit(maxMessageSize)
	c.conn.SetReadDeadline(time.Now().Add(pongWait))
	c.conn.SetPongHandler(func(string) error { c.conn.SetReadDeadline(time.Now().Add(pongWait)); return nil })

	for {
		var msg Msg
		err := c.conn.ReadJSON(&msg)
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				c.logger.WithError(err).Errorln("closing websocket connection")
			}
			return
		}
		c.hub.broadcast <- msg
	}
}

func (c *Client) write() {
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
		c.conn.Close()
	}()
	for {
		select {
		case msg, ok := <-c.send:
			c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if !ok {
				c.conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}
			c.conn.WriteJSON(msg)
			n := len(c.send)
			for i := 0; i < n; i++ {
				c.conn.WriteJSON(<-c.send)
			}
		case <-ticker.C:
			c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if err := c.conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				c.logger.WithError(err).Errorln("ping message")
				return
			}
		}
	}
}

func ClientHandler(hub *Hub) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logger := LoggerFromContext(r.Context())

		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			logger.WithError(err).Errorln("upgrading websocket connection")
			return
		}
		client := &Client{
			hub:    hub,
			conn:   conn,
			send:   make(chan Msg, 256),
			logger: logger,
		}
		var msg Msg

		client.conn.ReadJSON(&msg)

		if msg.Type != Connect {
			client.logger.Errorf("wrong message type, expected %d, got %d", Connect, msg.Type)
			conn.Close()
			return
		}

		client.name = msg.User
		client.hub.register <- client

		client.logger.Infoln("successfully initialized client, starting read/write loops")

		go client.read()
		go client.write()
		client.hub.broadcast <- msg
	})
}
