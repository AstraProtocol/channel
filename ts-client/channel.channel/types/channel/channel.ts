/* eslint-disable */
import { Writer, Reader } from "protobufjs/minimal";

export const protobufPackage = "channel.channel";

export interface Channel {
  index: string;
  multisigAddr: string;
  partA: string;
  partB: string;
  denom: string;
}

const baseChannel: object = {
  index: "",
  multisigAddr: "",
  partA: "",
  partB: "",
  denom: "",
};

export const Channel = {
  encode(message: Channel, writer: Writer = Writer.create()): Writer {
    if (message.Index !== "") {
      writer.uint32(10).string(message.Index);
    }
    if (message.MultisigAddr !== "") {
      writer.uint32(18).string(message.MultisigAddr);
    }
    if (message.PartA !== "") {
      writer.uint32(26).string(message.PartA);
    }
    if (message.PartB !== "") {
      writer.uint32(34).string(message.PartB);
    }
    if (message.Denom !== "") {
      writer.uint32(42).string(message.Denom);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): Channel {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseChannel } as Channel;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.Index = reader.string();
          break;
        case 2:
          message.MultisigAddr = reader.string();
          break;
        case 3:
          message.PartA = reader.string();
          break;
        case 4:
          message.PartB = reader.string();
          break;
        case 5:
          message.Denom = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): Channel {
    const message = { ...baseChannel } as Channel;
    if (object.Index !== undefined && object.Index !== null) {
      message.Index = String(object.Index);
    } else {
      message.Index = "";
    }
    if (object.MultisigAddr !== undefined && object.MultisigAddr !== null) {
      message.MultisigAddr = String(object.MultisigAddr);
    } else {
      message.MultisigAddr = "";
    }
    if (object.PartA !== undefined && object.PartA !== null) {
      message.PartA = String(object.PartA);
    } else {
      message.PartA = "";
    }
    if (object.PartB !== undefined && object.PartB !== null) {
      message.PartB = String(object.PartB);
    } else {
      message.PartB = "";
    }
    if (object.Denom !== undefined && object.Denom !== null) {
      message.Denom = String(object.Denom);
    } else {
      message.Denom = "";
    }

    return message;
  },

  toJSON(message: Channel): unknown {
    const obj: any = {};

    message.index !== undefined && (obj.index = message.index);
    message.multisigAddr !== undefined &&
      (obj.multisigAddr = message.multisigAddr);
    message.partA !== undefined && (obj.partA = message.partA);
    message.partB !== undefined && (obj.partB = message.partB);
    message.denom !== undefined && (obj.denom = message.denom);

    return obj;
  },

  fromPartial(object: DeepPartial<Channel>): Channel {
    const message = { ...baseChannel } as Channel;
    if (object.Index !== undefined && object.Index !== null) {
      message.Index = object.Index;
    } else {
      message.Index = "";
    }
    if (object.MultisigAddr !== undefined && object.MultisigAddr !== null) {
      message.MultisigAddr = object.MultisigAddr;
    } else {
      message.MultisigAddr = "";
    }
    if (object.PartA !== undefined && object.PartA !== null) {
      message.PartA = object.PartA;
    } else {
      message.PartA = "";
    }
    if (object.PartB !== undefined && object.PartB !== null) {
      message.PartB = object.PartB;
    } else {
      message.PartB = "";
    }
    if (object.Denom !== undefined && object.Denom !== null) {
      message.Denom = object.Denom;
    } else {
      message.Denom = "";
    }

    return message;
  },
};

type Builtin = Date | Function | Uint8Array | string | number | undefined;
export type DeepPartial<T> = T extends Builtin
  ? T
  : T extends Array<infer U>
  ? Array<DeepPartial<U>>
  : T extends ReadonlyArray<infer U>
  ? ReadonlyArray<DeepPartial<U>>
  : T extends {}
  ? { [K in keyof T]?: DeepPartial<T[K]> }
  : Partial<T>;
