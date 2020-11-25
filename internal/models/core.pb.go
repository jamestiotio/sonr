// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.13.0
// source: core.proto

package models

import (
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type LobbyEvent_Event int32

const (
	LobbyEvent_NONE   LobbyEvent_Event = 0
	LobbyEvent_JOIN   LobbyEvent_Event = 1
	LobbyEvent_UPDATE LobbyEvent_Event = 2
	LobbyEvent_EXIT   LobbyEvent_Event = 3
)

// Enum value maps for LobbyEvent_Event.
var (
	LobbyEvent_Event_name = map[int32]string{
		0: "NONE",
		1: "JOIN",
		2: "UPDATE",
		3: "EXIT",
	}
	LobbyEvent_Event_value = map[string]int32{
		"NONE":   0,
		"JOIN":   1,
		"UPDATE": 2,
		"EXIT":   3,
	}
)

func (x LobbyEvent_Event) Enum() *LobbyEvent_Event {
	p := new(LobbyEvent_Event)
	*p = x
	return p
}

func (x LobbyEvent_Event) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (LobbyEvent_Event) Descriptor() protoreflect.EnumDescriptor {
	return file_core_proto_enumTypes[0].Descriptor()
}

func (LobbyEvent_Event) Type() protoreflect.EnumType {
	return &file_core_proto_enumTypes[0]
}

func (x LobbyEvent_Event) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use LobbyEvent_Event.Descriptor instead.
func (LobbyEvent_Event) EnumDescriptor() ([]byte, []int) {
	return file_core_proto_rawDescGZIP(), []int{0, 0}
}

// [CORE]
// Message Sent when peer messages Lobby Topic
type LobbyEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Event LobbyEvent_Event `protobuf:"varint,1,opt,name=event,proto3,enum=LobbyEvent_Event" json:"event,omitempty"`
	Peer  *Peer            `protobuf:"bytes,2,opt,name=peer,proto3" json:"peer,omitempty"`
	Id    string           `protobuf:"bytes,3,opt,name=id,proto3" json:"id,omitempty"` // Optional used for remove
}

func (x *LobbyEvent) Reset() {
	*x = LobbyEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_core_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LobbyEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LobbyEvent) ProtoMessage() {}

func (x *LobbyEvent) ProtoReflect() protoreflect.Message {
	mi := &file_core_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LobbyEvent.ProtoReflect.Descriptor instead.
func (*LobbyEvent) Descriptor() ([]byte, []int) {
	return file_core_proto_rawDescGZIP(), []int{0}
}

func (x *LobbyEvent) GetEvent() LobbyEvent_Event {
	if x != nil {
		return x.Event
	}
	return LobbyEvent_NONE
}

func (x *LobbyEvent) GetPeer() *Peer {
	if x != nil {
		return x.Peer
	}
	return nil
}

func (x *LobbyEvent) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

// Define Block Type: Sent on Data Transfer
type Block struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data    []byte `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	Size    int64  `protobuf:"varint,2,opt,name=size,proto3" json:"size,omitempty"`
	Current int64  `protobuf:"varint,3,opt,name=current,proto3" json:"current,omitempty"`
	Total   int64  `protobuf:"varint,4,opt,name=total,proto3" json:"total,omitempty"`
}

func (x *Block) Reset() {
	*x = Block{}
	if protoimpl.UnsafeEnabled {
		mi := &file_core_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Block) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Block) ProtoMessage() {}

func (x *Block) ProtoReflect() protoreflect.Message {
	mi := &file_core_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Block.ProtoReflect.Descriptor instead.
func (*Block) Descriptor() ([]byte, []int) {
	return file_core_proto_rawDescGZIP(), []int{1}
}

func (x *Block) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *Block) GetSize() int64 {
	if x != nil {
		return x.Size
	}
	return 0
}

func (x *Block) GetCurrent() int64 {
	if x != nil {
		return x.Current
	}
	return 0
}

func (x *Block) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

type File struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Base64 []string `protobuf:"bytes,1,rep,name=base64,proto3" json:"base64,omitempty"`
	FileId string   `protobuf:"bytes,2,opt,name=fileId,proto3" json:"fileId,omitempty"`
	Name   string   `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Path   string   `protobuf:"bytes,4,opt,name=path,proto3" json:"path,omitempty"`
	Size   int64    `protobuf:"varint,5,opt,name=size,proto3" json:"size,omitempty"`
	Blocks int64    `protobuf:"varint,6,opt,name=blocks,proto3" json:"blocks,omitempty"`
	Kind   string   `protobuf:"bytes,7,opt,name=kind,proto3" json:"kind,omitempty"`
}

func (x *File) Reset() {
	*x = File{}
	if protoimpl.UnsafeEnabled {
		mi := &file_core_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *File) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*File) ProtoMessage() {}

func (x *File) ProtoReflect() protoreflect.Message {
	mi := &file_core_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use File.ProtoReflect.Descriptor instead.
func (*File) Descriptor() ([]byte, []int) {
	return file_core_proto_rawDescGZIP(), []int{2}
}

func (x *File) GetBase64() []string {
	if x != nil {
		return x.Base64
	}
	return nil
}

func (x *File) GetFileId() string {
	if x != nil {
		return x.FileId
	}
	return ""
}

func (x *File) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *File) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *File) GetSize() int64 {
	if x != nil {
		return x.Size
	}
	return 0
}

func (x *File) GetBlocks() int64 {
	if x != nil {
		return x.Blocks
	}
	return 0
}

func (x *File) GetKind() string {
	if x != nil {
		return x.Kind
	}
	return ""
}

var File_core_proto protoreflect.FileDescriptor

var file_core_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0a, 0x64, 0x61,
	0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x93, 0x01, 0x0a, 0x0a, 0x4c, 0x6f, 0x62,
	0x62, 0x79, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x27, 0x0a, 0x05, 0x65, 0x76, 0x65, 0x6e, 0x74,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x11, 0x2e, 0x4c, 0x6f, 0x62, 0x62, 0x79, 0x45, 0x76,
	0x65, 0x6e, 0x74, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x05, 0x65, 0x76, 0x65, 0x6e, 0x74,
	0x12, 0x19, 0x0a, 0x04, 0x70, 0x65, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x05,
	0x2e, 0x50, 0x65, 0x65, 0x72, 0x52, 0x04, 0x70, 0x65, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x31, 0x0a, 0x05, 0x45,
	0x76, 0x65, 0x6e, 0x74, 0x12, 0x08, 0x0a, 0x04, 0x4e, 0x4f, 0x4e, 0x45, 0x10, 0x00, 0x12, 0x08,
	0x0a, 0x04, 0x4a, 0x4f, 0x49, 0x4e, 0x10, 0x01, 0x12, 0x0a, 0x0a, 0x06, 0x55, 0x50, 0x44, 0x41,
	0x54, 0x45, 0x10, 0x02, 0x12, 0x08, 0x0a, 0x04, 0x45, 0x58, 0x49, 0x54, 0x10, 0x03, 0x22, 0x5f,
	0x0a, 0x05, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x12, 0x0a, 0x04, 0x73,
	0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x12,
	0x18, 0x0a, 0x07, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x07, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74,
	0x61, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x22,
	0x9e, 0x01, 0x0a, 0x04, 0x46, 0x69, 0x6c, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x62, 0x61, 0x73, 0x65,
	0x36, 0x34, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x62, 0x61, 0x73, 0x65, 0x36, 0x34,
	0x12, 0x16, 0x0a, 0x06, 0x66, 0x69, 0x6c, 0x65, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x66, 0x69, 0x6c, 0x65, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04,
	0x70, 0x61, 0x74, 0x68, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x61, 0x74, 0x68,
	0x12, 0x12, 0x0a, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04,
	0x73, 0x69, 0x7a, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x73, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x73, 0x12, 0x12, 0x0a, 0x04,
	0x6b, 0x69, 0x6e, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6b, 0x69, 0x6e, 0x64,
	0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x3b, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_core_proto_rawDescOnce sync.Once
	file_core_proto_rawDescData = file_core_proto_rawDesc
)

func file_core_proto_rawDescGZIP() []byte {
	file_core_proto_rawDescOnce.Do(func() {
		file_core_proto_rawDescData = protoimpl.X.CompressGZIP(file_core_proto_rawDescData)
	})
	return file_core_proto_rawDescData
}

var file_core_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_core_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_core_proto_goTypes = []interface{}{
	(LobbyEvent_Event)(0), // 0: LobbyEvent.Event
	(*LobbyEvent)(nil),    // 1: LobbyEvent
	(*Block)(nil),         // 2: Block
	(*File)(nil),          // 3: File
	(*Peer)(nil),          // 4: Peer
}
var file_core_proto_depIdxs = []int32{
	0, // 0: LobbyEvent.event:type_name -> LobbyEvent.Event
	4, // 1: LobbyEvent.peer:type_name -> Peer
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_core_proto_init() }
func file_core_proto_init() {
	if File_core_proto != nil {
		return
	}
	file_data_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_core_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LobbyEvent); i {
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
		file_core_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Block); i {
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
		file_core_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*File); i {
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
			RawDescriptor: file_core_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_core_proto_goTypes,
		DependencyIndexes: file_core_proto_depIdxs,
		EnumInfos:         file_core_proto_enumTypes,
		MessageInfos:      file_core_proto_msgTypes,
	}.Build()
	File_core_proto = out.File
	file_core_proto_rawDesc = nil
	file_core_proto_goTypes = nil
	file_core_proto_depIdxs = nil
}
