export enum MsgType {
  Connect = 0,
  Message,
  Disconnected,
}

export interface Msg {
  type: MsgType;
  user: string;
  message?: string;
  timestamp: Date;
}
