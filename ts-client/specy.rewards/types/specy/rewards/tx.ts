/* eslint-disable */
import _m0 from "protobufjs/minimal";

export const protobufPackage = "specy.rewards";

export interface MsgClaim {
  creator: string;
  path: string[];
}

export interface MsgClaimResponse {
}

export interface MsgSetRewardList {
  creator: string;
  list: string[];
}

export interface MsgSetRewardListResponse {
}

function createBaseMsgClaim(): MsgClaim {
  return { creator: "", path: [] };
}

export const MsgClaim = {
  encode(message: MsgClaim, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    for (const v of message.path) {
      writer.uint32(18).string(v!);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgClaim {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgClaim();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.path.push(reader.string());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgClaim {
    return {
      creator: isSet(object.creator) ? String(object.creator) : "",
      path: Array.isArray(object?.path) ? object.path.map((e: any) => String(e)) : [],
    };
  },

  toJSON(message: MsgClaim): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    if (message.path) {
      obj.path = message.path.map((e) => e);
    } else {
      obj.path = [];
    }
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgClaim>, I>>(object: I): MsgClaim {
    const message = createBaseMsgClaim();
    message.creator = object.creator ?? "";
    message.path = object.path?.map((e) => e) || [];
    return message;
  },
};

function createBaseMsgClaimResponse(): MsgClaimResponse {
  return {};
}

export const MsgClaimResponse = {
  encode(_: MsgClaimResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgClaimResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgClaimResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(_: any): MsgClaimResponse {
    return {};
  },

  toJSON(_: MsgClaimResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgClaimResponse>, I>>(_: I): MsgClaimResponse {
    const message = createBaseMsgClaimResponse();
    return message;
  },
};

function createBaseMsgSetRewardList(): MsgSetRewardList {
  return { creator: "", list: [] };
}

export const MsgSetRewardList = {
  encode(message: MsgSetRewardList, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    for (const v of message.list) {
      writer.uint32(18).string(v!);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgSetRewardList {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgSetRewardList();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.list.push(reader.string());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgSetRewardList {
    return {
      creator: isSet(object.creator) ? String(object.creator) : "",
      list: Array.isArray(object?.list) ? object.list.map((e: any) => String(e)) : [],
    };
  },

  toJSON(message: MsgSetRewardList): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    if (message.list) {
      obj.list = message.list.map((e) => e);
    } else {
      obj.list = [];
    }
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgSetRewardList>, I>>(object: I): MsgSetRewardList {
    const message = createBaseMsgSetRewardList();
    message.creator = object.creator ?? "";
    message.list = object.list?.map((e) => e) || [];
    return message;
  },
};

function createBaseMsgSetRewardListResponse(): MsgSetRewardListResponse {
  return {};
}

export const MsgSetRewardListResponse = {
  encode(_: MsgSetRewardListResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgSetRewardListResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgSetRewardListResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(_: any): MsgSetRewardListResponse {
    return {};
  },

  toJSON(_: MsgSetRewardListResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgSetRewardListResponse>, I>>(_: I): MsgSetRewardListResponse {
    const message = createBaseMsgSetRewardListResponse();
    return message;
  },
};

/** Msg defines the Msg service. */
export interface Msg {
  Claim(request: MsgClaim): Promise<MsgClaimResponse>;
  /** this line is used by starport scaffolding # proto/tx/rpc */
  SetRewardList(request: MsgSetRewardList): Promise<MsgSetRewardListResponse>;
}

export class MsgClientImpl implements Msg {
  private readonly rpc: Rpc;
  constructor(rpc: Rpc) {
    this.rpc = rpc;
    this.Claim = this.Claim.bind(this);
    this.SetRewardList = this.SetRewardList.bind(this);
  }
  Claim(request: MsgClaim): Promise<MsgClaimResponse> {
    const data = MsgClaim.encode(request).finish();
    const promise = this.rpc.request("specy.rewards.Msg", "Claim", data);
    return promise.then((data) => MsgClaimResponse.decode(new _m0.Reader(data)));
  }

  SetRewardList(request: MsgSetRewardList): Promise<MsgSetRewardListResponse> {
    const data = MsgSetRewardList.encode(request).finish();
    const promise = this.rpc.request("specy.rewards.Msg", "SetRewardList", data);
    return promise.then((data) => MsgSetRewardListResponse.decode(new _m0.Reader(data)));
  }
}

interface Rpc {
  request(service: string, method: string, data: Uint8Array): Promise<Uint8Array>;
}

type Builtin = Date | Function | Uint8Array | string | number | boolean | undefined;

export type DeepPartial<T> = T extends Builtin ? T
  : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>>
  : T extends {} ? { [K in keyof T]?: DeepPartial<T[K]> }
  : Partial<T>;

type KeysOfUnion<T> = T extends T ? keyof T : never;
export type Exact<P, I extends P> = P extends Builtin ? P
  : P & { [K in keyof P]: Exact<P[K], I[K]> } & { [K in Exclude<keyof I, KeysOfUnion<P>>]: never };

function isSet(value: any): boolean {
  return value !== null && value !== undefined;
}