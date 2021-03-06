// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.15.8
// source: proto/messages.proto

package proto

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Event_Type int32

const (
	Event_CLIENT_JOINED  Event_Type = 0
	Event_CLIENT_LEFT    Event_Type = 1
	Event_CLIENT_UPDATED Event_Type = 2
	Event_SERVER_WELCOME Event_Type = 3
)

// Enum value maps for Event_Type.
var (
	Event_Type_name = map[int32]string{
		0: "CLIENT_JOINED",
		1: "CLIENT_LEFT",
		2: "CLIENT_UPDATED",
		3: "SERVER_WELCOME",
	}
	Event_Type_value = map[string]int32{
		"CLIENT_JOINED":  0,
		"CLIENT_LEFT":    1,
		"CLIENT_UPDATED": 2,
		"SERVER_WELCOME": 3,
	}
)

func (x Event_Type) Enum() *Event_Type {
	p := new(Event_Type)
	*p = x
	return p
}

func (x Event_Type) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Event_Type) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_messages_proto_enumTypes[0].Descriptor()
}

func (Event_Type) Type() protoreflect.EnumType {
	return &file_proto_messages_proto_enumTypes[0]
}

func (x Event_Type) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Event_Type.Descriptor instead.
func (Event_Type) EnumDescriptor() ([]byte, []int) {
	return file_proto_messages_proto_rawDescGZIP(), []int{0, 0}
}

type Event_Client_ConnectionStatus int32

const (
	Event_Client_OFFLINE Event_Client_ConnectionStatus = 0
	Event_Client_ONLINE  Event_Client_ConnectionStatus = 1
)

// Enum value maps for Event_Client_ConnectionStatus.
var (
	Event_Client_ConnectionStatus_name = map[int32]string{
		0: "OFFLINE",
		1: "ONLINE",
	}
	Event_Client_ConnectionStatus_value = map[string]int32{
		"OFFLINE": 0,
		"ONLINE":  1,
	}
)

func (x Event_Client_ConnectionStatus) Enum() *Event_Client_ConnectionStatus {
	p := new(Event_Client_ConnectionStatus)
	*p = x
	return p
}

func (x Event_Client_ConnectionStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Event_Client_ConnectionStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_messages_proto_enumTypes[1].Descriptor()
}

func (Event_Client_ConnectionStatus) Type() protoreflect.EnumType {
	return &file_proto_messages_proto_enumTypes[1]
}

func (x Event_Client_ConnectionStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Event_Client_ConnectionStatus.Descriptor instead.
func (Event_Client_ConnectionStatus) EnumDescriptor() ([]byte, []int) {
	return file_proto_messages_proto_rawDescGZIP(), []int{0, 0, 0}
}

type Event_Client_State int32

const (
	Event_Client_JOINING Event_Client_State = 0
	Event_Client_JOINED  Event_Client_State = 1
)

// Enum value maps for Event_Client_State.
var (
	Event_Client_State_name = map[int32]string{
		0: "JOINING",
		1: "JOINED",
	}
	Event_Client_State_value = map[string]int32{
		"JOINING": 0,
		"JOINED":  1,
	}
)

func (x Event_Client_State) Enum() *Event_Client_State {
	p := new(Event_Client_State)
	*p = x
	return p
}

func (x Event_Client_State) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Event_Client_State) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_messages_proto_enumTypes[2].Descriptor()
}

func (Event_Client_State) Type() protoreflect.EnumType {
	return &file_proto_messages_proto_enumTypes[2]
}

func (x Event_Client_State) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Event_Client_State.Descriptor instead.
func (Event_Client_State) EnumDescriptor() ([]byte, []int) {
	return file_proto_messages_proto_rawDescGZIP(), []int{0, 0, 1}
}

type ClientEvent_Type int32

const (
	ClientEvent_NAME_CHANGE ClientEvent_Type = 0
)

// Enum value maps for ClientEvent_Type.
var (
	ClientEvent_Type_name = map[int32]string{
		0: "NAME_CHANGE",
	}
	ClientEvent_Type_value = map[string]int32{
		"NAME_CHANGE": 0,
	}
)

func (x ClientEvent_Type) Enum() *ClientEvent_Type {
	p := new(ClientEvent_Type)
	*p = x
	return p
}

func (x ClientEvent_Type) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ClientEvent_Type) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_messages_proto_enumTypes[3].Descriptor()
}

func (ClientEvent_Type) Type() protoreflect.EnumType {
	return &file_proto_messages_proto_enumTypes[3]
}

func (x ClientEvent_Type) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ClientEvent_Type.Descriptor instead.
func (ClientEvent_Type) EnumDescriptor() ([]byte, []int) {
	return file_proto_messages_proto_rawDescGZIP(), []int{1, 0}
}

type Event struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type Event_Type `protobuf:"varint,1,opt,name=type,proto3,enum=Event_Type" json:"type,omitempty"`
	// Types that are assignable to Data:
	//	*Event_Client_
	//	*Event_Room_
	//	*Event_Welcome
	Data isEvent_Data `protobuf_oneof:"data"`
}

func (x *Event) Reset() {
	*x = Event{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_messages_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Event) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Event) ProtoMessage() {}

func (x *Event) ProtoReflect() protoreflect.Message {
	mi := &file_proto_messages_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Event.ProtoReflect.Descriptor instead.
func (*Event) Descriptor() ([]byte, []int) {
	return file_proto_messages_proto_rawDescGZIP(), []int{0}
}

func (x *Event) GetType() Event_Type {
	if x != nil {
		return x.Type
	}
	return Event_CLIENT_JOINED
}

func (m *Event) GetData() isEvent_Data {
	if m != nil {
		return m.Data
	}
	return nil
}

func (x *Event) GetClient() *Event_Client {
	if x, ok := x.GetData().(*Event_Client_); ok {
		return x.Client
	}
	return nil
}

func (x *Event) GetRoom() *Event_Room {
	if x, ok := x.GetData().(*Event_Room_); ok {
		return x.Room
	}
	return nil
}

func (x *Event) GetWelcome() *Event_ServerWelcome {
	if x, ok := x.GetData().(*Event_Welcome); ok {
		return x.Welcome
	}
	return nil
}

type isEvent_Data interface {
	isEvent_Data()
}

type Event_Client_ struct {
	Client *Event_Client `protobuf:"bytes,2,opt,name=client,proto3,oneof"`
}

type Event_Room_ struct {
	Room *Event_Room `protobuf:"bytes,3,opt,name=room,proto3,oneof"`
}

type Event_Welcome struct {
	Welcome *Event_ServerWelcome `protobuf:"bytes,4,opt,name=welcome,proto3,oneof"`
}

func (*Event_Client_) isEvent_Data() {}

func (*Event_Room_) isEvent_Data() {}

func (*Event_Welcome) isEvent_Data() {}

type ClientEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type ClientEvent_Type `protobuf:"varint,1,opt,name=type,proto3,enum=ClientEvent_Type" json:"type,omitempty"`
	// Types that are assignable to Data:
	//	*ClientEvent_NameChange_
	Data isClientEvent_Data `protobuf_oneof:"data"`
}

func (x *ClientEvent) Reset() {
	*x = ClientEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_messages_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClientEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClientEvent) ProtoMessage() {}

func (x *ClientEvent) ProtoReflect() protoreflect.Message {
	mi := &file_proto_messages_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClientEvent.ProtoReflect.Descriptor instead.
func (*ClientEvent) Descriptor() ([]byte, []int) {
	return file_proto_messages_proto_rawDescGZIP(), []int{1}
}

func (x *ClientEvent) GetType() ClientEvent_Type {
	if x != nil {
		return x.Type
	}
	return ClientEvent_NAME_CHANGE
}

func (m *ClientEvent) GetData() isClientEvent_Data {
	if m != nil {
		return m.Data
	}
	return nil
}

func (x *ClientEvent) GetNameChange() *ClientEvent_NameChange {
	if x, ok := x.GetData().(*ClientEvent_NameChange_); ok {
		return x.NameChange
	}
	return nil
}

type isClientEvent_Data interface {
	isClientEvent_Data()
}

type ClientEvent_NameChange_ struct {
	NameChange *ClientEvent_NameChange `protobuf:"bytes,2,opt,name=nameChange,proto3,oneof"`
}

func (*ClientEvent_NameChange_) isClientEvent_Data() {}

type Event_Client struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id               string                        `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name             string                        `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	ConnectionStatus Event_Client_ConnectionStatus `protobuf:"varint,3,opt,name=connectionStatus,proto3,enum=Event_Client_ConnectionStatus" json:"connectionStatus,omitempty"`
	State            Event_Client_State            `protobuf:"varint,4,opt,name=state,proto3,enum=Event_Client_State" json:"state,omitempty"`
}

func (x *Event_Client) Reset() {
	*x = Event_Client{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_messages_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Event_Client) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Event_Client) ProtoMessage() {}

func (x *Event_Client) ProtoReflect() protoreflect.Message {
	mi := &file_proto_messages_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Event_Client.ProtoReflect.Descriptor instead.
func (*Event_Client) Descriptor() ([]byte, []int) {
	return file_proto_messages_proto_rawDescGZIP(), []int{0, 0}
}

func (x *Event_Client) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Event_Client) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Event_Client) GetConnectionStatus() Event_Client_ConnectionStatus {
	if x != nil {
		return x.ConnectionStatus
	}
	return Event_Client_OFFLINE
}

func (x *Event_Client) GetState() Event_Client_State {
	if x != nil {
		return x.State
	}
	return Event_Client_JOINING
}

type Event_Room struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      string                   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Clients map[string]*Event_Client `protobuf:"bytes,2,rep,name=clients,proto3" json:"clients,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *Event_Room) Reset() {
	*x = Event_Room{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_messages_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Event_Room) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Event_Room) ProtoMessage() {}

func (x *Event_Room) ProtoReflect() protoreflect.Message {
	mi := &file_proto_messages_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Event_Room.ProtoReflect.Descriptor instead.
func (*Event_Room) Descriptor() ([]byte, []int) {
	return file_proto_messages_proto_rawDescGZIP(), []int{0, 1}
}

func (x *Event_Room) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Event_Room) GetClients() map[string]*Event_Client {
	if x != nil {
		return x.Clients
	}
	return nil
}

type Event_ServerWelcome struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ClientId string      `protobuf:"bytes,1,opt,name=clientId,proto3" json:"clientId,omitempty"`
	Room     *Event_Room `protobuf:"bytes,2,opt,name=room,proto3" json:"room,omitempty"`
}

func (x *Event_ServerWelcome) Reset() {
	*x = Event_ServerWelcome{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_messages_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Event_ServerWelcome) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Event_ServerWelcome) ProtoMessage() {}

func (x *Event_ServerWelcome) ProtoReflect() protoreflect.Message {
	mi := &file_proto_messages_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Event_ServerWelcome.ProtoReflect.Descriptor instead.
func (*Event_ServerWelcome) Descriptor() ([]byte, []int) {
	return file_proto_messages_proto_rawDescGZIP(), []int{0, 2}
}

func (x *Event_ServerWelcome) GetClientId() string {
	if x != nil {
		return x.ClientId
	}
	return ""
}

func (x *Event_ServerWelcome) GetRoom() *Event_Room {
	if x != nil {
		return x.Room
	}
	return nil
}

type ClientEvent_NameChange struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=Name,proto3" json:"Name,omitempty"`
}

func (x *ClientEvent_NameChange) Reset() {
	*x = ClientEvent_NameChange{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_messages_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClientEvent_NameChange) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClientEvent_NameChange) ProtoMessage() {}

func (x *ClientEvent_NameChange) ProtoReflect() protoreflect.Message {
	mi := &file_proto_messages_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClientEvent_NameChange.ProtoReflect.Descriptor instead.
func (*ClientEvent_NameChange) Descriptor() ([]byte, []int) {
	return file_proto_messages_proto_rawDescGZIP(), []int{1, 0}
}

func (x *ClientEvent_NameChange) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

var File_proto_messages_proto protoreflect.FileDescriptor

var file_proto_messages_proto_rawDesc = []byte{
	0x0a, 0x14, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xdd, 0x05, 0x0a, 0x05, 0x45, 0x76, 0x65, 0x6e, 0x74,
	0x12, 0x1f, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0b,
	0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70,
	0x65, 0x12, 0x27, 0x0a, 0x06, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x0d, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74,
	0x48, 0x00, 0x52, 0x06, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x12, 0x21, 0x0a, 0x04, 0x72, 0x6f,
	0x6f, 0x6d, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74,
	0x2e, 0x52, 0x6f, 0x6f, 0x6d, 0x48, 0x00, 0x52, 0x04, 0x72, 0x6f, 0x6f, 0x6d, 0x12, 0x30, 0x0a,
	0x07, 0x77, 0x65, 0x6c, 0x63, 0x6f, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14,
	0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x57, 0x65, 0x6c,
	0x63, 0x6f, 0x6d, 0x65, 0x48, 0x00, 0x52, 0x07, 0x77, 0x65, 0x6c, 0x63, 0x6f, 0x6d, 0x65, 0x1a,
	0xf2, 0x01, 0x0a, 0x06, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x4a,
	0x0a, 0x10, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1e, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74,
	0x2e, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x10, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x29, 0x0a, 0x05, 0x73, 0x74,
	0x61, 0x74, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x13, 0x2e, 0x45, 0x76, 0x65, 0x6e,
	0x74, 0x2e, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x05,
	0x73, 0x74, 0x61, 0x74, 0x65, 0x22, 0x2b, 0x0a, 0x10, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x0b, 0x0a, 0x07, 0x4f, 0x46, 0x46,
	0x4c, 0x49, 0x4e, 0x45, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06, 0x4f, 0x4e, 0x4c, 0x49, 0x4e, 0x45,
	0x10, 0x01, 0x22, 0x20, 0x0a, 0x05, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x0b, 0x0a, 0x07, 0x4a,
	0x4f, 0x49, 0x4e, 0x49, 0x4e, 0x47, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06, 0x4a, 0x4f, 0x49, 0x4e,
	0x45, 0x44, 0x10, 0x01, 0x1a, 0x95, 0x01, 0x0a, 0x04, 0x52, 0x6f, 0x6f, 0x6d, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x32, 0x0a,
	0x07, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x18,
	0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x52, 0x6f, 0x6f, 0x6d, 0x2e, 0x43, 0x6c, 0x69, 0x65,
	0x6e, 0x74, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x07, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74,
	0x73, 0x1a, 0x49, 0x0a, 0x0c, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x73, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x6b, 0x65, 0x79, 0x12, 0x23, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x43, 0x6c, 0x69, 0x65, 0x6e,
	0x74, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x4c, 0x0a, 0x0d,
	0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x57, 0x65, 0x6c, 0x63, 0x6f, 0x6d, 0x65, 0x12, 0x1a, 0x0a,
	0x08, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x04, 0x72, 0x6f, 0x6f,
	0x6d, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x2e,
	0x52, 0x6f, 0x6f, 0x6d, 0x52, 0x04, 0x72, 0x6f, 0x6f, 0x6d, 0x22, 0x52, 0x0a, 0x04, 0x54, 0x79,
	0x70, 0x65, 0x12, 0x11, 0x0a, 0x0d, 0x43, 0x4c, 0x49, 0x45, 0x4e, 0x54, 0x5f, 0x4a, 0x4f, 0x49,
	0x4e, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0f, 0x0a, 0x0b, 0x43, 0x4c, 0x49, 0x45, 0x4e, 0x54, 0x5f,
	0x4c, 0x45, 0x46, 0x54, 0x10, 0x01, 0x12, 0x12, 0x0a, 0x0e, 0x43, 0x4c, 0x49, 0x45, 0x4e, 0x54,
	0x5f, 0x55, 0x50, 0x44, 0x41, 0x54, 0x45, 0x44, 0x10, 0x02, 0x12, 0x12, 0x0a, 0x0e, 0x53, 0x45,
	0x52, 0x56, 0x45, 0x52, 0x5f, 0x57, 0x45, 0x4c, 0x43, 0x4f, 0x4d, 0x45, 0x10, 0x03, 0x42, 0x06,
	0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0xb2, 0x01, 0x0a, 0x0b, 0x43, 0x6c, 0x69, 0x65, 0x6e,
	0x74, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x25, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x11, 0x2e, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x45, 0x76, 0x65,
	0x6e, 0x74, 0x2e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x39, 0x0a,
	0x0a, 0x6e, 0x61, 0x6d, 0x65, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x17, 0x2e, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x2e,
	0x4e, 0x61, 0x6d, 0x65, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x48, 0x00, 0x52, 0x0a, 0x6e, 0x61,
	0x6d, 0x65, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x1a, 0x20, 0x0a, 0x0a, 0x4e, 0x61, 0x6d, 0x65,
	0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x17, 0x0a, 0x04, 0x54, 0x79,
	0x70, 0x65, 0x12, 0x0f, 0x0a, 0x0b, 0x4e, 0x41, 0x4d, 0x45, 0x5f, 0x43, 0x48, 0x41, 0x4e, 0x47,
	0x45, 0x10, 0x00, 0x42, 0x06, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x42, 0x27, 0x5a, 0x25, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x54, 0x61, 0x69, 0x74, 0x65, 0x72,
	0x73, 0x2f, 0x6d, 0x6f, 0x6e, 0x73, 0x74, 0x65, 0x72, 0x6d, 0x61, 0x6b, 0x65, 0x72, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_messages_proto_rawDescOnce sync.Once
	file_proto_messages_proto_rawDescData = file_proto_messages_proto_rawDesc
)

func file_proto_messages_proto_rawDescGZIP() []byte {
	file_proto_messages_proto_rawDescOnce.Do(func() {
		file_proto_messages_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_messages_proto_rawDescData)
	})
	return file_proto_messages_proto_rawDescData
}

var file_proto_messages_proto_enumTypes = make([]protoimpl.EnumInfo, 4)
var file_proto_messages_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_proto_messages_proto_goTypes = []interface{}{
	(Event_Type)(0),                    // 0: Event.Type
	(Event_Client_ConnectionStatus)(0), // 1: Event.Client.ConnectionStatus
	(Event_Client_State)(0),            // 2: Event.Client.State
	(ClientEvent_Type)(0),              // 3: ClientEvent.Type
	(*Event)(nil),                      // 4: Event
	(*ClientEvent)(nil),                // 5: ClientEvent
	(*Event_Client)(nil),               // 6: Event.Client
	(*Event_Room)(nil),                 // 7: Event.Room
	(*Event_ServerWelcome)(nil),        // 8: Event.ServerWelcome
	nil,                                // 9: Event.Room.ClientsEntry
	(*ClientEvent_NameChange)(nil),     // 10: ClientEvent.NameChange
}
var file_proto_messages_proto_depIdxs = []int32{
	0,  // 0: Event.type:type_name -> Event.Type
	6,  // 1: Event.client:type_name -> Event.Client
	7,  // 2: Event.room:type_name -> Event.Room
	8,  // 3: Event.welcome:type_name -> Event.ServerWelcome
	3,  // 4: ClientEvent.type:type_name -> ClientEvent.Type
	10, // 5: ClientEvent.nameChange:type_name -> ClientEvent.NameChange
	1,  // 6: Event.Client.connectionStatus:type_name -> Event.Client.ConnectionStatus
	2,  // 7: Event.Client.state:type_name -> Event.Client.State
	9,  // 8: Event.Room.clients:type_name -> Event.Room.ClientsEntry
	7,  // 9: Event.ServerWelcome.room:type_name -> Event.Room
	6,  // 10: Event.Room.ClientsEntry.value:type_name -> Event.Client
	11, // [11:11] is the sub-list for method output_type
	11, // [11:11] is the sub-list for method input_type
	11, // [11:11] is the sub-list for extension type_name
	11, // [11:11] is the sub-list for extension extendee
	0,  // [0:11] is the sub-list for field type_name
}

func init() { file_proto_messages_proto_init() }
func file_proto_messages_proto_init() {
	if File_proto_messages_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_messages_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Event); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_messages_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClientEvent); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_messages_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Event_Client); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_messages_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Event_Room); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_messages_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Event_ServerWelcome); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_messages_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClientEvent_NameChange); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	file_proto_messages_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*Event_Client_)(nil),
		(*Event_Room_)(nil),
		(*Event_Welcome)(nil),
	}
	file_proto_messages_proto_msgTypes[1].OneofWrappers = []interface{}{
		(*ClientEvent_NameChange_)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_messages_proto_rawDesc,
			NumEnums:      4,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_messages_proto_goTypes,
		DependencyIndexes: file_proto_messages_proto_depIdxs,
		EnumInfos:         file_proto_messages_proto_enumTypes,
		MessageInfos:      file_proto_messages_proto_msgTypes,
	}.Build()
	File_proto_messages_proto = out.File
	file_proto_messages_proto_rawDesc = nil
	file_proto_messages_proto_goTypes = nil
	file_proto_messages_proto_depIdxs = nil
}
