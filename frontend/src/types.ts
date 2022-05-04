export enum MsgType {
  Connected = 1,
  Message,
  Disconnected,
}

export interface Msg {
  Type: MsgType;
  User: string;
  Message?: string;
  Timestamp: Date;
}
