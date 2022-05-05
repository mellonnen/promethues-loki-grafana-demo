export enum MsgType {
  Connect = 0,
  Message,
  Disconnect,
}

export interface Msg {
  type: MsgType;
  user: string;
  message?: string;
  timestamp: Date;
}
