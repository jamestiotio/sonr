// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.15.6
// source: core.proto

package models

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

// Current Node Lifecycle State
type LifecycleState int32

const (
	LifecycleState_Active  LifecycleState = 0 // Resume or Start
	LifecycleState_Paused  LifecycleState = 1 // User Paused
	LifecycleState_Stopped LifecycleState = 2 // User Quit
)

// Enum value maps for LifecycleState.
var (
	LifecycleState_name = map[int32]string{
		0: "Active",
		1: "Paused",
		2: "Stopped",
	}
	LifecycleState_value = map[string]int32{
		"Active":  0,
		"Paused":  1,
		"Stopped": 2,
	}
)

func (x LifecycleState) Enum() *LifecycleState {
	p := new(LifecycleState)
	*p = x
	return p
}

func (x LifecycleState) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (LifecycleState) Descriptor() protoreflect.EnumDescriptor {
	return file_core_proto_enumTypes[0].Descriptor()
}

func (LifecycleState) Type() protoreflect.EnumType {
	return &file_core_proto_enumTypes[0]
}

func (x LifecycleState) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use LifecycleState.Descriptor instead.
func (LifecycleState) EnumDescriptor() ([]byte, []int) {
	return file_core_proto_rawDescGZIP(), []int{0}
}

// Sonr Default Protocols
type SonrProtocol int32

const (
	SonrProtocol_Unspecified    SonrProtocol = 0 // Default Value
	SonrProtocol_Authorize      SonrProtocol = 1 // Auth Protocol for Invite/Response
	SonrProtocol_Devices        SonrProtocol = 2 // Device Protocol for User Linked Devices
	SonrProtocol_Linker         SonrProtocol = 3 // Protocol for Linking Device to User
	SonrProtocol_LocalTransfer  SonrProtocol = 4 // Protocol for Local Transfer to Peer
	SonrProtocol_RemoteTransfer SonrProtocol = 5 // Protocol for Remote Transfer to Peer
)

// Enum value maps for SonrProtocol.
var (
	SonrProtocol_name = map[int32]string{
		0: "Unspecified",
		1: "Authorize",
		2: "Devices",
		3: "Linker",
		4: "LocalTransfer",
		5: "RemoteTransfer",
	}
	SonrProtocol_value = map[string]int32{
		"Unspecified":    0,
		"Authorize":      1,
		"Devices":        2,
		"Linker":         3,
		"LocalTransfer":  4,
		"RemoteTransfer": 5,
	}
)

func (x SonrProtocol) Enum() *SonrProtocol {
	p := new(SonrProtocol)
	*p = x
	return p
}

func (x SonrProtocol) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (SonrProtocol) Descriptor() protoreflect.EnumDescriptor {
	return file_core_proto_enumTypes[1].Descriptor()
}

func (SonrProtocol) Type() protoreflect.EnumType {
	return &file_core_proto_enumTypes[1]
}

func (x SonrProtocol) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use SonrProtocol.Descriptor instead.
func (SonrProtocol) EnumDescriptor() ([]byte, []int) {
	return file_core_proto_rawDescGZIP(), []int{1}
}

// Sent on Data Transfer to Add piece of File - Binary
type Chunk struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Buffer     []byte `protobuf:"bytes,1,opt,name=buffer,proto3" json:"buffer,omitempty"`          // Binary Value of Chunk
	Size       int32  `protobuf:"varint,2,opt,name=size,proto3" json:"size,omitempty"`             // Size of this Chunk
	IsComplete bool   `protobuf:"varint,3,opt,name=isComplete,proto3" json:"isComplete,omitempty"` // Whether this is final chunk
}

func (x *Chunk) Reset() {
	*x = Chunk{}
	if protoimpl.UnsafeEnabled {
		mi := &file_core_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Chunk) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Chunk) ProtoMessage() {}

func (x *Chunk) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use Chunk.ProtoReflect.Descriptor instead.
func (*Chunk) Descriptor() ([]byte, []int) {
	return file_core_proto_rawDescGZIP(), []int{0}
}

func (x *Chunk) GetBuffer() []byte {
	if x != nil {
		return x.Buffer
	}
	return nil
}

func (x *Chunk) GetSize() int32 {
	if x != nil {
		return x.Size
	}
	return 0
}

func (x *Chunk) GetIsComplete() bool {
	if x != nil {
		return x.IsComplete
	}
	return false
}

type PushMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Peer *Peer             `protobuf:"bytes,1,opt,name=peer,proto3" json:"peer,omitempty"`                                                                                         // Peer to send to
	Data map[string]string `protobuf:"bytes,2,rep,name=data,proto3" json:"data,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"` // Data to send
}

func (x *PushMessage) Reset() {
	*x = PushMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_core_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PushMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PushMessage) ProtoMessage() {}

func (x *PushMessage) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use PushMessage.ProtoReflect.Descriptor instead.
func (*PushMessage) Descriptor() ([]byte, []int) {
	return file_core_proto_rawDescGZIP(), []int{1}
}

func (x *PushMessage) GetPeer() *Peer {
	if x != nil {
		return x.Peer
	}
	return nil
}

func (x *PushMessage) GetData() map[string]string {
	if x != nil {
		return x.Data
	}
	return nil
}

var File_core_proto protoreflect.FileDescriptor

var file_core_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x6d, 0x6f,
	0x64, 0x65, 0x6c, 0x73, 0x1a, 0x0a, 0x70, 0x65, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x53, 0x0a, 0x05, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x12, 0x16, 0x0a, 0x06, 0x62, 0x75, 0x66,
	0x66, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x06, 0x62, 0x75, 0x66, 0x66, 0x65,
	0x72, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x04, 0x73, 0x69, 0x7a, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x69, 0x73, 0x43, 0x6f, 0x6d, 0x70, 0x6c,
	0x65, 0x74, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x69, 0x73, 0x43, 0x6f, 0x6d,
	0x70, 0x6c, 0x65, 0x74, 0x65, 0x22, 0x9b, 0x01, 0x0a, 0x0b, 0x50, 0x75, 0x73, 0x68, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x20, 0x0a, 0x04, 0x70, 0x65, 0x65, 0x72, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x50, 0x65, 0x65,
	0x72, 0x52, 0x04, 0x70, 0x65, 0x65, 0x72, 0x12, 0x31, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18,
	0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x50,
	0x75, 0x73, 0x68, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x1a, 0x37, 0x0a, 0x09, 0x44, 0x61,
	0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a,
	0x02, 0x38, 0x01, 0x2a, 0x35, 0x0a, 0x0e, 0x4c, 0x69, 0x66, 0x65, 0x63, 0x79, 0x63, 0x6c, 0x65,
	0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x0a, 0x0a, 0x06, 0x41, 0x63, 0x74, 0x69, 0x76, 0x65, 0x10,
	0x00, 0x12, 0x0a, 0x0a, 0x06, 0x50, 0x61, 0x75, 0x73, 0x65, 0x64, 0x10, 0x01, 0x12, 0x0b, 0x0a,
	0x07, 0x53, 0x74, 0x6f, 0x70, 0x70, 0x65, 0x64, 0x10, 0x02, 0x2a, 0x6e, 0x0a, 0x0c, 0x53, 0x6f,
	0x6e, 0x72, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x12, 0x0f, 0x0a, 0x0b, 0x55, 0x6e,
	0x73, 0x70, 0x65, 0x63, 0x69, 0x66, 0x69, 0x65, 0x64, 0x10, 0x00, 0x12, 0x0d, 0x0a, 0x09, 0x41,
	0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x65, 0x10, 0x01, 0x12, 0x0b, 0x0a, 0x07, 0x44, 0x65,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x10, 0x02, 0x12, 0x0a, 0x0a, 0x06, 0x4c, 0x69, 0x6e, 0x6b, 0x65,
	0x72, 0x10, 0x03, 0x12, 0x11, 0x0a, 0x0d, 0x4c, 0x6f, 0x63, 0x61, 0x6c, 0x54, 0x72, 0x61, 0x6e,
	0x73, 0x66, 0x65, 0x72, 0x10, 0x04, 0x12, 0x12, 0x0a, 0x0e, 0x52, 0x65, 0x6d, 0x6f, 0x74, 0x65,
	0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x10, 0x05, 0x42, 0x09, 0x5a, 0x07, 0x2f, 0x6d,
	0x6f, 0x64, 0x65, 0x6c, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
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

var file_core_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_core_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_core_proto_goTypes = []interface{}{
	(LifecycleState)(0), // 0: models.LifecycleState
	(SonrProtocol)(0),   // 1: models.SonrProtocol
	(*Chunk)(nil),       // 2: models.Chunk
	(*PushMessage)(nil), // 3: models.PushMessage
	nil,                 // 4: models.PushMessage.DataEntry
	(*Peer)(nil),        // 5: models.Peer
}
var file_core_proto_depIdxs = []int32{
	5, // 0: models.PushMessage.peer:type_name -> models.Peer
	4, // 1: models.PushMessage.data:type_name -> models.PushMessage.DataEntry
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
	file_peer_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_core_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Chunk); i {
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
			switch v := v.(*PushMessage); i {
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
			NumEnums:      2,
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
