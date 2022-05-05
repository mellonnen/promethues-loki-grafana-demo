import { Msg, MsgType } from "./types";

type ChatMessageProps = {
  myUsername: string;
  msg: Msg;
};

const ChatMessage = (props: ChatMessageProps) => {
  const { msg, myUsername } = props;
  const time = new Date(msg.timestamp)
    .toLocaleTimeString([], { hour: "2-digit", minute: "2-digit" })
    .replace(" AM", "")
    .replace(" PM", "");

  const dataTime = msg.user === myUsername ? time : `${msg.user} ${time}`;

  const msgClass =
    msg.type === MsgType.Disconnect || msg.type === MsgType.Connect
      ? "msg info"
      : msg.user === myUsername
      ? "msg sent"
      : "msg rcvd";

  const msgContent =
    msg.type === MsgType.Disconnect
      ? `${msg.user} has left the chat.`
      : msg.type === MsgType.Connect
      ? `${msg.user} has joined the chat.`
      : msg.message;

  return (
    <div data-time={dataTime} className={msgClass}>
      {msgContent}
    </div>
  );
};

export default ChatMessage;
