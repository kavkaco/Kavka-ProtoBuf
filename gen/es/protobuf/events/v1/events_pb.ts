// @generated by protoc-gen-es v1.10.0 with parameter "target=ts,import_extension=.ts"
// @generated from file protobuf/events/v1/events.proto (package events.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage } from "@bufbuild/protobuf";
import { Message, proto3 } from "@bufbuild/protobuf";
import { Chat } from "../../model/chat/v1/chat_pb.ts";
import { Message as Message$1 } from "../../model/message/v1/message_pb.ts";

/**
 * @generated from message events.v1.StreamEvent
 */
export class StreamEvent extends Message<StreamEvent> {
  /**
   * @generated from field: string sender_user_id = 1;
   */
  senderUserId = "";

  /**
   * @generated from field: repeated string receivers_user_id = 2;
   */
  receiversUserId: string[] = [];

  /**
   * @generated from field: bytes payload = 3;
   */
  payload = new Uint8Array(0);

  constructor(data?: PartialMessage<StreamEvent>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "events.v1.StreamEvent";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "sender_user_id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "receivers_user_id", kind: "scalar", T: 9 /* ScalarType.STRING */, repeated: true },
    { no: 3, name: "payload", kind: "scalar", T: 12 /* ScalarType.BYTES */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): StreamEvent {
    return new StreamEvent().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): StreamEvent {
    return new StreamEvent().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): StreamEvent {
    return new StreamEvent().fromJsonString(jsonString, options);
  }

  static equals(a: StreamEvent | PlainMessage<StreamEvent> | undefined, b: StreamEvent | PlainMessage<StreamEvent> | undefined): boolean {
    return proto3.util.equals(StreamEvent, a, b);
  }
}

/**
 * @generated from message events.v1.SubscribeEventsStreamRequest
 */
export class SubscribeEventsStreamRequest extends Message<SubscribeEventsStreamRequest> {
  constructor(data?: PartialMessage<SubscribeEventsStreamRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "events.v1.SubscribeEventsStreamRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): SubscribeEventsStreamRequest {
    return new SubscribeEventsStreamRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): SubscribeEventsStreamRequest {
    return new SubscribeEventsStreamRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): SubscribeEventsStreamRequest {
    return new SubscribeEventsStreamRequest().fromJsonString(jsonString, options);
  }

  static equals(a: SubscribeEventsStreamRequest | PlainMessage<SubscribeEventsStreamRequest> | undefined, b: SubscribeEventsStreamRequest | PlainMessage<SubscribeEventsStreamRequest> | undefined): boolean {
    return proto3.util.equals(SubscribeEventsStreamRequest, a, b);
  }
}

/**
 * @generated from message events.v1.SubscribeEventsStreamResponse
 */
export class SubscribeEventsStreamResponse extends Message<SubscribeEventsStreamResponse> {
  /**
   * @generated from field: string name = 1;
   */
  name = "";

  /**
   * @generated from field: events.v1.SubscribeEventsStreamResponse.Type type = 2;
   */
  type = SubscribeEventsStreamResponse_Type.UNSPECIFIED;

  /**
   * @generated from oneof events.v1.SubscribeEventsStreamResponse.payload
   */
  payload: {
    /**
     * @generated from field: events.v1.AddChat add_chat = 3;
     */
    value: AddChat;
    case: "addChat";
  } | {
    /**
     * @generated from field: events.v1.AddMessage add_message = 4;
     */
    value: AddMessage;
    case: "addMessage";
  } | { case: undefined; value?: undefined } = { case: undefined };

  constructor(data?: PartialMessage<SubscribeEventsStreamResponse>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "events.v1.SubscribeEventsStreamResponse";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "name", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "type", kind: "enum", T: proto3.getEnumType(SubscribeEventsStreamResponse_Type) },
    { no: 3, name: "add_chat", kind: "message", T: AddChat, oneof: "payload" },
    { no: 4, name: "add_message", kind: "message", T: AddMessage, oneof: "payload" },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): SubscribeEventsStreamResponse {
    return new SubscribeEventsStreamResponse().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): SubscribeEventsStreamResponse {
    return new SubscribeEventsStreamResponse().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): SubscribeEventsStreamResponse {
    return new SubscribeEventsStreamResponse().fromJsonString(jsonString, options);
  }

  static equals(a: SubscribeEventsStreamResponse | PlainMessage<SubscribeEventsStreamResponse> | undefined, b: SubscribeEventsStreamResponse | PlainMessage<SubscribeEventsStreamResponse> | undefined): boolean {
    return proto3.util.equals(SubscribeEventsStreamResponse, a, b);
  }
}

/**
 * @generated from enum events.v1.SubscribeEventsStreamResponse.Type
 */
export enum SubscribeEventsStreamResponse_Type {
  /**
   * @generated from enum value: TYPE_UNSPECIFIED = 0;
   */
  UNSPECIFIED = 0,

  /**
   * @generated from enum value: TYPE_ADD_CHAT = 1;
   */
  ADD_CHAT = 1,

  /**
   * @generated from enum value: TYPE_REMOVE_CHAT = 2;
   */
  REMOVE_CHAT = 2,

  /**
   * @generated from enum value: TYPE_UPDATE_CHAT = 3;
   */
  UPDATE_CHAT = 3,

  /**
   * @generated from enum value: TYPE_ADD_MESSAGE = 4;
   */
  ADD_MESSAGE = 4,

  /**
   * @generated from enum value: TYPE_REMOVE_MESSAGE = 5;
   */
  REMOVE_MESSAGE = 5,

  /**
   * @generated from enum value: TYPE_UPDATE_MESSAGE = 6;
   */
  UPDATE_MESSAGE = 6,

  /**
   * @generated from enum value: TYPE_CLEAR_CHAT = 7;
   */
  CLEAR_CHAT = 7,

  /**
   * @generated from enum value: TYPE_MESSAGE_SEEN = 8;
   */
  MESSAGE_SEEN = 8,
}
// Retrieve enum metadata with: proto3.getEnumType(SubscribeEventsStreamResponse_Type)
proto3.util.setEnumType(SubscribeEventsStreamResponse_Type, "events.v1.SubscribeEventsStreamResponse.Type", [
  { no: 0, name: "TYPE_UNSPECIFIED" },
  { no: 1, name: "TYPE_ADD_CHAT" },
  { no: 2, name: "TYPE_REMOVE_CHAT" },
  { no: 3, name: "TYPE_UPDATE_CHAT" },
  { no: 4, name: "TYPE_ADD_MESSAGE" },
  { no: 5, name: "TYPE_REMOVE_MESSAGE" },
  { no: 6, name: "TYPE_UPDATE_MESSAGE" },
  { no: 7, name: "TYPE_CLEAR_CHAT" },
  { no: 8, name: "TYPE_MESSAGE_SEEN" },
]);

/**
 * @generated from message events.v1.AddChat
 */
export class AddChat extends Message<AddChat> {
  /**
   * @generated from field: model.chat.v1.Chat chat = 1;
   */
  chat?: Chat;

  constructor(data?: PartialMessage<AddChat>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "events.v1.AddChat";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "chat", kind: "message", T: Chat },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): AddChat {
    return new AddChat().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): AddChat {
    return new AddChat().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): AddChat {
    return new AddChat().fromJsonString(jsonString, options);
  }

  static equals(a: AddChat | PlainMessage<AddChat> | undefined, b: AddChat | PlainMessage<AddChat> | undefined): boolean {
    return proto3.util.equals(AddChat, a, b);
  }
}

/**
 * @generated from message events.v1.AddMessage
 */
export class AddMessage extends Message<AddMessage> {
  /**
   * @generated from field: string chat_id = 1;
   */
  chatId = "";

  /**
   * @generated from field: model.message.v1.Message message = 2;
   */
  message?: Message$1;

  constructor(data?: PartialMessage<AddMessage>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "events.v1.AddMessage";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "chat_id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "message", kind: "message", T: Message$1 },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): AddMessage {
    return new AddMessage().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): AddMessage {
    return new AddMessage().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): AddMessage {
    return new AddMessage().fromJsonString(jsonString, options);
  }

  static equals(a: AddMessage | PlainMessage<AddMessage> | undefined, b: AddMessage | PlainMessage<AddMessage> | undefined): boolean {
    return proto3.util.equals(AddMessage, a, b);
  }
}

