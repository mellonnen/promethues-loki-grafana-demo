import { useEffect, useRef, useState } from "react";
import names from "./names";
import { Msg, MsgType } from "./types";

type ChatRoomProps = {
  myUsername: string;
};

const ChatRoom = (props: ChatRoomProps) => {
  let { myUsername } = props;
  if (myUsername === "") {
    myUsername = names()!;
  }
  const [currentMsg, setCurrentMsg] = useState("");
  const [msgs, setMsgs] = useState<Msg[]>([]);
  const ws = useRef<WebSocket>();
  useEffect(() => {
    ws.current = new WebSocket("ws://localhost:8080/chat");
    ws.current.onopen = () => {
      const msg: Msg = {
        type: MsgType.Connect,
        user: myUsername,
        timestamp: new Date(),
      };
      ws.current?.send(JSON.stringify(msg));
    };
    ws.current.onmessage = (e) => {
      const msg: Msg = JSON.parse(e.data);
      setMsgs((prev) => [...prev, msg]);
    };
  }, []);
  return (
    <div>
      <ol>
        {msgs.map((m) => (
          <li>{m.user}</li>
        ))}
      </ol>
      <input
        type="text"
        value={currentMsg}
        onChange={(e) => setCurrentMsg(e.target.value)}
        onKeyDown={(e) => {
          if (e.key !== "Enter") return;
          const msg: Msg = {
            type: MsgType.Message,
            user: myUsername,
            message: currentMsg,
            timestamp: new Date(),
          };
          ws.current?.send(JSON.stringify(msg));
        }}
        placeholder="write message"
      />
    </div>
  );
};

export default ChatRoom;
