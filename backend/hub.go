package main

import (
	"github.com/sirupsen/logrus"
)

type Msg struct {
	User     string `json:"user"`
	Message  string `json:"message,omitempty"`
	LeftChat bool   `json:"left_chat"`
}

type Hub struct {
	clients map[*Client]bool

	broadcast chan Msg

	register   chan *Client
	unregister chan *Client
	logger     *logrus.Entry
}

func newHub(logger *logrus.Entry) *Hub {
	l := logger.WithField("component", "Hub")
	return &Hub{
		broadcast:  make(chan Msg, 100),
		register:   make(chan *Client),
		unregister: make(chan *Client),
		clients:    make(map[*Client]bool),
		logger:     l,
	}
}

func (h *Hub) Run() {
	h.logger.Infoln("Started hub")
	for {
		select {
		case client := <-h.register:
			h.clients[client] = true
			h.logger.Infof("registered client with username: %s", client.name)

		case client := <-h.unregister:
			if _, ok := h.clients[client]; ok {
				h.logger.Infof("unregistered client with username: %s", client.name)
				delete(h.clients, client)
				close(client.send)
				h.broadcast <- Msg{User: client.name, LeftChat: true}
			}
		case msg := <-h.broadcast:
			for client := range h.clients {
				client.send <- msg
			}
		}
	}
}
