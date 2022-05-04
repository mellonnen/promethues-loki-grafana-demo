type ChatRoomProps = {
  myUsername: string;
};

const ChatRoom = (props: ChatRoomProps) => {
  const { myUsername } = props;
  return <div>{myUsername}</div>;
};

export default ChatRoom;
