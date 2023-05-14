// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.23.0
// source: vstreamer_protos/commander/commander.proto

package _go

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

type Operation int32

const (
	Operation_TRANSCRIBE          Operation = 0
	Operation_TRANSLATE           Operation = 1
	Operation_SUBTITLE            Operation = 2
	Operation_SUBTITLE_TRANSLATED Operation = 3
	Operation_TTS                 Operation = 4
	Operation_VC                  Operation = 5
	Operation_PLAYBACK            Operation = 6
	Operation_PAUSE               Operation = 7
	Operation_RESUME              Operation = 8
	Operation_RELOAD              Operation = 9
	Operation_SET_FILTERS         Operation = 10
)

// Enum value maps for Operation.
var (
	Operation_name = map[int32]string{
		0:  "TRANSCRIBE",
		1:  "TRANSLATE",
		2:  "SUBTITLE",
		3:  "SUBTITLE_TRANSLATED",
		4:  "TTS",
		5:  "VC",
		6:  "PLAYBACK",
		7:  "PAUSE",
		8:  "RESUME",
		9:  "RELOAD",
		10: "SET_FILTERS",
	}
	Operation_value = map[string]int32{
		"TRANSCRIBE":          0,
		"TRANSLATE":           1,
		"SUBTITLE":            2,
		"SUBTITLE_TRANSLATED": 3,
		"TTS":                 4,
		"VC":                  5,
		"PLAYBACK":            6,
		"PAUSE":               7,
		"RESUME":              8,
		"RELOAD":              9,
		"SET_FILTERS":         10,
	}
)

func (x Operation) Enum() *Operation {
	p := new(Operation)
	*p = x
	return p
}

func (x Operation) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Operation) Descriptor() protoreflect.EnumDescriptor {
	return file_vstreamer_protos_commander_commander_proto_enumTypes[0].Descriptor()
}

func (Operation) Type() protoreflect.EnumType {
	return &file_vstreamer_protos_commander_commander_proto_enumTypes[0]
}

func (x Operation) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Operation.Descriptor instead.
func (Operation) EnumDescriptor() ([]byte, []int) {
	return file_vstreamer_protos_commander_commander_proto_rawDescGZIP(), []int{0}
}

type Command struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Chains   []*OperationChain `protobuf:"bytes,1,rep,name=chains,proto3" json:"chains,omitempty"`
	Sound    *Sound            `protobuf:"bytes,2,opt,name=sound,proto3" json:"sound,omitempty"`
	Text     string            `protobuf:"bytes,3,opt,name=text,proto3" json:"text,omitempty"`
	FilePath string            `protobuf:"bytes,4,opt,name=file_path,json=filePath,proto3" json:"file_path,omitempty"`
	Filters  []string          `protobuf:"bytes,5,rep,name=filters,proto3" json:"filters,omitempty"`
}

func (x *Command) Reset() {
	*x = Command{}
	if protoimpl.UnsafeEnabled {
		mi := &file_vstreamer_protos_commander_commander_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Command) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Command) ProtoMessage() {}

func (x *Command) ProtoReflect() protoreflect.Message {
	mi := &file_vstreamer_protos_commander_commander_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Command.ProtoReflect.Descriptor instead.
func (*Command) Descriptor() ([]byte, []int) {
	return file_vstreamer_protos_commander_commander_proto_rawDescGZIP(), []int{0}
}

func (x *Command) GetChains() []*OperationChain {
	if x != nil {
		return x.Chains
	}
	return nil
}

func (x *Command) GetSound() *Sound {
	if x != nil {
		return x.Sound
	}
	return nil
}

func (x *Command) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

func (x *Command) GetFilePath() string {
	if x != nil {
		return x.FilePath
	}
	return ""
}

func (x *Command) GetFilters() []string {
	if x != nil {
		return x.Filters
	}
	return nil
}

type Sound struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data     []byte `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	Rate     int32  `protobuf:"varint,2,opt,name=rate,proto3" json:"rate,omitempty"`
	Format   int32  `protobuf:"varint,3,opt,name=format,proto3" json:"format,omitempty"`
	Channels int32  `protobuf:"varint,4,opt,name=channels,proto3" json:"channels,omitempty"`
}

func (x *Sound) Reset() {
	*x = Sound{}
	if protoimpl.UnsafeEnabled {
		mi := &file_vstreamer_protos_commander_commander_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Sound) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Sound) ProtoMessage() {}

func (x *Sound) ProtoReflect() protoreflect.Message {
	mi := &file_vstreamer_protos_commander_commander_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Sound.ProtoReflect.Descriptor instead.
func (*Sound) Descriptor() ([]byte, []int) {
	return file_vstreamer_protos_commander_commander_proto_rawDescGZIP(), []int{1}
}

func (x *Sound) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *Sound) GetRate() int32 {
	if x != nil {
		return x.Rate
	}
	return 0
}

func (x *Sound) GetFormat() int32 {
	if x != nil {
		return x.Format
	}
	return 0
}

func (x *Sound) GetChannels() int32 {
	if x != nil {
		return x.Channels
	}
	return 0
}

type OperationChain struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Operations []*OperationRoute `protobuf:"bytes,1,rep,name=operations,proto3" json:"operations,omitempty"`
}

func (x *OperationChain) Reset() {
	*x = OperationChain{}
	if protoimpl.UnsafeEnabled {
		mi := &file_vstreamer_protos_commander_commander_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OperationChain) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OperationChain) ProtoMessage() {}

func (x *OperationChain) ProtoReflect() protoreflect.Message {
	mi := &file_vstreamer_protos_commander_commander_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OperationChain.ProtoReflect.Descriptor instead.
func (*OperationChain) Descriptor() ([]byte, []int) {
	return file_vstreamer_protos_commander_commander_proto_rawDescGZIP(), []int{2}
}

func (x *OperationChain) GetOperations() []*OperationRoute {
	if x != nil {
		return x.Operations
	}
	return nil
}

type OperationRoute struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Operation Operation `protobuf:"varint,1,opt,name=operation,proto3,enum=voicerecog.Operation" json:"operation,omitempty"`
	Remote    string    `protobuf:"bytes,2,opt,name=remote,proto3" json:"remote,omitempty"`
}

func (x *OperationRoute) Reset() {
	*x = OperationRoute{}
	if protoimpl.UnsafeEnabled {
		mi := &file_vstreamer_protos_commander_commander_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OperationRoute) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OperationRoute) ProtoMessage() {}

func (x *OperationRoute) ProtoReflect() protoreflect.Message {
	mi := &file_vstreamer_protos_commander_commander_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OperationRoute.ProtoReflect.Descriptor instead.
func (*OperationRoute) Descriptor() ([]byte, []int) {
	return file_vstreamer_protos_commander_commander_proto_rawDescGZIP(), []int{3}
}

func (x *OperationRoute) GetOperation() Operation {
	if x != nil {
		return x.Operation
	}
	return Operation_TRANSCRIBE
}

func (x *OperationRoute) GetRemote() string {
	if x != nil {
		return x.Remote
	}
	return ""
}

type Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result bool `protobuf:"varint,1,opt,name=result,proto3" json:"result,omitempty"`
}

func (x *Response) Reset() {
	*x = Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_vstreamer_protos_commander_commander_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Response) ProtoMessage() {}

func (x *Response) ProtoReflect() protoreflect.Message {
	mi := &file_vstreamer_protos_commander_commander_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Response.ProtoReflect.Descriptor instead.
func (*Response) Descriptor() ([]byte, []int) {
	return file_vstreamer_protos_commander_commander_proto_rawDescGZIP(), []int{4}
}

func (x *Response) GetResult() bool {
	if x != nil {
		return x.Result
	}
	return false
}

var File_vstreamer_protos_commander_commander_proto protoreflect.FileDescriptor

var file_vstreamer_protos_commander_commander_proto_rawDesc = []byte{
	0x0a, 0x2a, 0x76, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x65, 0x72, 0x5f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x73, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x65, 0x72, 0x2f, 0x63, 0x6f, 0x6d,
	0x6d, 0x61, 0x6e, 0x64, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x76, 0x6f,
	0x69, 0x63, 0x65, 0x72, 0x65, 0x63, 0x6f, 0x67, 0x22, 0xb1, 0x01, 0x0a, 0x07, 0x43, 0x6f, 0x6d,
	0x6d, 0x61, 0x6e, 0x64, 0x12, 0x32, 0x0a, 0x06, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x72, 0x65, 0x63, 0x6f,
	0x67, 0x2e, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x68, 0x61, 0x69, 0x6e,
	0x52, 0x06, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x73, 0x12, 0x27, 0x0a, 0x05, 0x73, 0x6f, 0x75, 0x6e,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x72,
	0x65, 0x63, 0x6f, 0x67, 0x2e, 0x53, 0x6f, 0x75, 0x6e, 0x64, 0x52, 0x05, 0x73, 0x6f, 0x75, 0x6e,
	0x64, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x78, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x74, 0x65, 0x78, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x70, 0x61,
	0x74, 0x68, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x50, 0x61,
	0x74, 0x68, 0x12, 0x18, 0x0a, 0x07, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x18, 0x05, 0x20,
	0x03, 0x28, 0x09, 0x52, 0x07, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x22, 0x63, 0x0a, 0x05,
	0x53, 0x6f, 0x75, 0x6e, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0c, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x12, 0x0a, 0x04, 0x72, 0x61, 0x74,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x72, 0x61, 0x74, 0x65, 0x12, 0x16, 0x0a,
	0x06, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x66,
	0x6f, 0x72, 0x6d, 0x61, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c,
	0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c,
	0x73, 0x22, 0x4c, 0x0a, 0x0e, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x68,
	0x61, 0x69, 0x6e, 0x12, 0x3a, 0x0a, 0x0a, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x72,
	0x65, 0x63, 0x6f, 0x67, 0x2e, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x6f,
	0x75, 0x74, 0x65, 0x52, 0x0a, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22,
	0x5d, 0x0a, 0x0e, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x6f, 0x75, 0x74,
	0x65, 0x12, 0x33, 0x0a, 0x09, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x15, 0x2e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x72, 0x65, 0x63, 0x6f,
	0x67, 0x2e, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x09, 0x6f, 0x70, 0x65,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x22, 0x22,
	0x0a, 0x08, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x2a, 0xa4, 0x01, 0x0a, 0x09, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x0e, 0x0a, 0x0a, 0x54, 0x52, 0x41, 0x4e, 0x53, 0x43, 0x52, 0x49, 0x42, 0x45, 0x10, 0x00,
	0x12, 0x0d, 0x0a, 0x09, 0x54, 0x52, 0x41, 0x4e, 0x53, 0x4c, 0x41, 0x54, 0x45, 0x10, 0x01, 0x12,
	0x0c, 0x0a, 0x08, 0x53, 0x55, 0x42, 0x54, 0x49, 0x54, 0x4c, 0x45, 0x10, 0x02, 0x12, 0x17, 0x0a,
	0x13, 0x53, 0x55, 0x42, 0x54, 0x49, 0x54, 0x4c, 0x45, 0x5f, 0x54, 0x52, 0x41, 0x4e, 0x53, 0x4c,
	0x41, 0x54, 0x45, 0x44, 0x10, 0x03, 0x12, 0x07, 0x0a, 0x03, 0x54, 0x54, 0x53, 0x10, 0x04, 0x12,
	0x06, 0x0a, 0x02, 0x56, 0x43, 0x10, 0x05, 0x12, 0x0c, 0x0a, 0x08, 0x50, 0x4c, 0x41, 0x59, 0x42,
	0x41, 0x43, 0x4b, 0x10, 0x06, 0x12, 0x09, 0x0a, 0x05, 0x50, 0x41, 0x55, 0x53, 0x45, 0x10, 0x07,
	0x12, 0x0a, 0x0a, 0x06, 0x52, 0x45, 0x53, 0x55, 0x4d, 0x45, 0x10, 0x08, 0x12, 0x0a, 0x0a, 0x06,
	0x52, 0x45, 0x4c, 0x4f, 0x41, 0x44, 0x10, 0x09, 0x12, 0x0f, 0x0a, 0x0b, 0x53, 0x45, 0x54, 0x5f,
	0x46, 0x49, 0x4c, 0x54, 0x45, 0x52, 0x53, 0x10, 0x0a, 0x32, 0x90, 0x01, 0x0a, 0x09, 0x43, 0x6f,
	0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x65, 0x72, 0x12, 0x3e, 0x0a, 0x0f, 0x70, 0x72, 0x6f, 0x63, 0x65,
	0x73, 0x73, 0x5f, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x12, 0x13, 0x2e, 0x76, 0x6f, 0x69,
	0x63, 0x65, 0x72, 0x65, 0x63, 0x6f, 0x67, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x1a,
	0x14, 0x2e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x72, 0x65, 0x63, 0x6f, 0x67, 0x2e, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x43, 0x0a, 0x14, 0x73, 0x79, 0x6e, 0x63, 0x5f,
	0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x12,
	0x13, 0x2e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x72, 0x65, 0x63, 0x6f, 0x67, 0x2e, 0x43, 0x6f, 0x6d,
	0x6d, 0x61, 0x6e, 0x64, 0x1a, 0x14, 0x2e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x72, 0x65, 0x63, 0x6f,
	0x67, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x2b, 0x5a, 0x29,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6e, 0x64, 0x65,
	0x6b, 0x6f, 0x31, 0x34, 0x33, 0x2f, 0x76, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x65, 0x72, 0x2d,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x67, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_vstreamer_protos_commander_commander_proto_rawDescOnce sync.Once
	file_vstreamer_protos_commander_commander_proto_rawDescData = file_vstreamer_protos_commander_commander_proto_rawDesc
)

func file_vstreamer_protos_commander_commander_proto_rawDescGZIP() []byte {
	file_vstreamer_protos_commander_commander_proto_rawDescOnce.Do(func() {
		file_vstreamer_protos_commander_commander_proto_rawDescData = protoimpl.X.CompressGZIP(file_vstreamer_protos_commander_commander_proto_rawDescData)
	})
	return file_vstreamer_protos_commander_commander_proto_rawDescData
}

var file_vstreamer_protos_commander_commander_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_vstreamer_protos_commander_commander_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_vstreamer_protos_commander_commander_proto_goTypes = []interface{}{
	(Operation)(0),         // 0: voicerecog.Operation
	(*Command)(nil),        // 1: voicerecog.Command
	(*Sound)(nil),          // 2: voicerecog.Sound
	(*OperationChain)(nil), // 3: voicerecog.OperationChain
	(*OperationRoute)(nil), // 4: voicerecog.OperationRoute
	(*Response)(nil),       // 5: voicerecog.Response
}
var file_vstreamer_protos_commander_commander_proto_depIdxs = []int32{
	3, // 0: voicerecog.Command.chains:type_name -> voicerecog.OperationChain
	2, // 1: voicerecog.Command.sound:type_name -> voicerecog.Sound
	4, // 2: voicerecog.OperationChain.operations:type_name -> voicerecog.OperationRoute
	0, // 3: voicerecog.OperationRoute.operation:type_name -> voicerecog.Operation
	1, // 4: voicerecog.Commander.process_command:input_type -> voicerecog.Command
	1, // 5: voicerecog.Commander.sync_process_command:input_type -> voicerecog.Command
	5, // 6: voicerecog.Commander.process_command:output_type -> voicerecog.Response
	5, // 7: voicerecog.Commander.sync_process_command:output_type -> voicerecog.Response
	6, // [6:8] is the sub-list for method output_type
	4, // [4:6] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_vstreamer_protos_commander_commander_proto_init() }
func file_vstreamer_protos_commander_commander_proto_init() {
	if File_vstreamer_protos_commander_commander_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_vstreamer_protos_commander_commander_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Command); i {
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
		file_vstreamer_protos_commander_commander_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Sound); i {
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
		file_vstreamer_protos_commander_commander_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OperationChain); i {
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
		file_vstreamer_protos_commander_commander_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OperationRoute); i {
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
		file_vstreamer_protos_commander_commander_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Response); i {
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
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_vstreamer_protos_commander_commander_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_vstreamer_protos_commander_commander_proto_goTypes,
		DependencyIndexes: file_vstreamer_protos_commander_commander_proto_depIdxs,
		EnumInfos:         file_vstreamer_protos_commander_commander_proto_enumTypes,
		MessageInfos:      file_vstreamer_protos_commander_commander_proto_msgTypes,
	}.Build()
	File_vstreamer_protos_commander_commander_proto = out.File
	file_vstreamer_protos_commander_commander_proto_rawDesc = nil
	file_vstreamer_protos_commander_commander_proto_goTypes = nil
	file_vstreamer_protos_commander_commander_proto_depIdxs = nil
}