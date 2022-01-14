// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        (unknown)
// source: node/discover/v1/discover.proto

// Package Discover is used to find other Peers in the sonr network.

package discover

import (
	common "github.com/sonr-io/core/common"
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

type VisibilityRequest_Visibility int32

const (
	VisibilityRequest_VISIBILITY_UNSPECIFIED VisibilityRequest_Visibility = 0
	VisibilityRequest_VISIBILITY_AVAILABLE   VisibilityRequest_Visibility = 1 // Everyone can see this peer
	VisibilityRequest_VISIBILITY_HIDDEN      VisibilityRequest_Visibility = 2 // Only Linked Devices can see this peer
	VisibilityRequest_VISIBILITY_FRIENDS     VisibilityRequest_Visibility = 3 // Only Friends can see this peer
)

// Enum value maps for VisibilityRequest_Visibility.
var (
	VisibilityRequest_Visibility_name = map[int32]string{
		0: "VISIBILITY_UNSPECIFIED",
		1: "VISIBILITY_AVAILABLE",
		2: "VISIBILITY_HIDDEN",
		3: "VISIBILITY_FRIENDS",
	}
	VisibilityRequest_Visibility_value = map[string]int32{
		"VISIBILITY_UNSPECIFIED": 0,
		"VISIBILITY_AVAILABLE":   1,
		"VISIBILITY_HIDDEN":      2,
		"VISIBILITY_FRIENDS":     3,
	}
)

func (x VisibilityRequest_Visibility) Enum() *VisibilityRequest_Visibility {
	p := new(VisibilityRequest_Visibility)
	*p = x
	return p
}

func (x VisibilityRequest_Visibility) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (VisibilityRequest_Visibility) Descriptor() protoreflect.EnumDescriptor {
	return file_node_discover_v1_discover_proto_enumTypes[0].Descriptor()
}

func (VisibilityRequest_Visibility) Type() protoreflect.EnumType {
	return &file_node_discover_v1_discover_proto_enumTypes[0]
}

func (x VisibilityRequest_Visibility) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use VisibilityRequest_Visibility.Descriptor instead.
func (VisibilityRequest_Visibility) EnumDescriptor() ([]byte, []int) {
	return file_node_discover_v1_discover_proto_rawDescGZIP(), []int{1, 0}
}

type VisibilityResponse_Visibility int32

const (
	VisibilityResponse_VISIBILITY_UNSPECIFIED VisibilityResponse_Visibility = 0
	VisibilityResponse_VISIBILITY_AVAILABLE   VisibilityResponse_Visibility = 1 // Everyone can see this peer
	VisibilityResponse_VISIBILITY_HIDDEN      VisibilityResponse_Visibility = 2 // Only Linked Devices can see this peer
	VisibilityResponse_VISIBILITY_FRIENDS     VisibilityResponse_Visibility = 3 // Only Friends can see this peer
)

// Enum value maps for VisibilityResponse_Visibility.
var (
	VisibilityResponse_Visibility_name = map[int32]string{
		0: "VISIBILITY_UNSPECIFIED",
		1: "VISIBILITY_AVAILABLE",
		2: "VISIBILITY_HIDDEN",
		3: "VISIBILITY_FRIENDS",
	}
	VisibilityResponse_Visibility_value = map[string]int32{
		"VISIBILITY_UNSPECIFIED": 0,
		"VISIBILITY_AVAILABLE":   1,
		"VISIBILITY_HIDDEN":      2,
		"VISIBILITY_FRIENDS":     3,
	}
)

func (x VisibilityResponse_Visibility) Enum() *VisibilityResponse_Visibility {
	p := new(VisibilityResponse_Visibility)
	*p = x
	return p
}

func (x VisibilityResponse_Visibility) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (VisibilityResponse_Visibility) Descriptor() protoreflect.EnumDescriptor {
	return file_node_discover_v1_discover_proto_enumTypes[1].Descriptor()
}

func (VisibilityResponse_Visibility) Type() protoreflect.EnumType {
	return &file_node_discover_v1_discover_proto_enumTypes[1]
}

func (x VisibilityResponse_Visibility) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use VisibilityResponse_Visibility.Descriptor instead.
func (VisibilityResponse_Visibility) EnumDescriptor() ([]byte, []int) {
	return file_node_discover_v1_discover_proto_rawDescGZIP(), []int{2, 0}
}

// LobbyMessage is message passed from Peer in Lobby
type LobbyMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Peer     *common.Peer     `protobuf:"bytes,1,opt,name=peer,proto3" json:"peer,omitempty"`             // Users Peer Data
	Message  *string          `protobuf:"bytes,2,opt,name=message,proto3,oneof" json:"message,omitempty"` // Message to be published
	Metadata *common.Metadata `protobuf:"bytes,3,opt,name=metadata,proto3" json:"metadata,omitempty"`     // Metadata
}

func (x *LobbyMessage) Reset() {
	*x = LobbyMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_node_discover_v1_discover_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LobbyMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LobbyMessage) ProtoMessage() {}

func (x *LobbyMessage) ProtoReflect() protoreflect.Message {
	mi := &file_node_discover_v1_discover_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LobbyMessage.ProtoReflect.Descriptor instead.
func (*LobbyMessage) Descriptor() ([]byte, []int) {
	return file_node_discover_v1_discover_proto_rawDescGZIP(), []int{0}
}

func (x *LobbyMessage) GetPeer() *common.Peer {
	if x != nil {
		return x.Peer
	}
	return nil
}

func (x *LobbyMessage) GetMessage() string {
	if x != nil && x.Message != nil {
		return *x.Message
	}
	return ""
}

func (x *LobbyMessage) GetMetadata() *common.Metadata {
	if x != nil {
		return x.Metadata
	}
	return nil
}

// VisibilityRequest is Message for updating Peer Visibility in Exchange
type VisibilityRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SName      string                       `protobuf:"bytes,1,opt,name=s_name,json=sName,proto3" json:"s_name,omitempty"`                                                  // SName combined with Device ID and Hashed
	PublicKey  []byte                       `protobuf:"bytes,2,opt,name=public_key,json=publicKey,proto3" json:"public_key,omitempty"`                                      // Buffer of Public Key
	Visibility VisibilityRequest_Visibility `protobuf:"varint,3,opt,name=visibility,proto3,enum=node.discover.v1.VisibilityRequest_Visibility" json:"visibility,omitempty"` // Visibility
}

func (x *VisibilityRequest) Reset() {
	*x = VisibilityRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_node_discover_v1_discover_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VisibilityRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VisibilityRequest) ProtoMessage() {}

func (x *VisibilityRequest) ProtoReflect() protoreflect.Message {
	mi := &file_node_discover_v1_discover_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VisibilityRequest.ProtoReflect.Descriptor instead.
func (*VisibilityRequest) Descriptor() ([]byte, []int) {
	return file_node_discover_v1_discover_proto_rawDescGZIP(), []int{1}
}

func (x *VisibilityRequest) GetSName() string {
	if x != nil {
		return x.SName
	}
	return ""
}

func (x *VisibilityRequest) GetPublicKey() []byte {
	if x != nil {
		return x.PublicKey
	}
	return nil
}

func (x *VisibilityRequest) GetVisibility() VisibilityRequest_Visibility {
	if x != nil {
		return x.Visibility
	}
	return VisibilityRequest_VISIBILITY_UNSPECIFIED
}

// VisibilityResponse is response for VisibilityRequest
type VisibilityResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success    bool                          `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`                                                           // If Request was Successful
	Error      string                        `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`                                                                // Error Message if Request was not successful
	Visibility VisibilityResponse_Visibility `protobuf:"varint,3,opt,name=visibility,proto3,enum=node.discover.v1.VisibilityResponse_Visibility" json:"visibility,omitempty"` // Visibility
}

func (x *VisibilityResponse) Reset() {
	*x = VisibilityResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_node_discover_v1_discover_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VisibilityResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VisibilityResponse) ProtoMessage() {}

func (x *VisibilityResponse) ProtoReflect() protoreflect.Message {
	mi := &file_node_discover_v1_discover_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VisibilityResponse.ProtoReflect.Descriptor instead.
func (*VisibilityResponse) Descriptor() ([]byte, []int) {
	return file_node_discover_v1_discover_proto_rawDescGZIP(), []int{2}
}

func (x *VisibilityResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *VisibilityResponse) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

func (x *VisibilityResponse) GetVisibility() VisibilityResponse_Visibility {
	if x != nil {
		return x.Visibility
	}
	return VisibilityResponse_VISIBILITY_UNSPECIFIED
}

var File_node_discover_v1_discover_proto protoreflect.FileDescriptor

var file_node_discover_v1_discover_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x6e, 0x6f, 0x64, 0x65, 0x2f, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x2f,
	0x76, 0x31, 0x2f, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x10, 0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72,
	0x2e, 0x76, 0x31, 0x1a, 0x11, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x63, 0x6f, 0x72, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x89, 0x01, 0x0a, 0x0c, 0x4c, 0x6f, 0x62, 0x62, 0x79,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x20, 0x0a, 0x04, 0x70, 0x65, 0x65, 0x72, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x50,
	0x65, 0x65, 0x72, 0x52, 0x04, 0x70, 0x65, 0x65, 0x72, 0x12, 0x1d, 0x0a, 0x07, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x07, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x88, 0x01, 0x01, 0x12, 0x2c, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61,
	0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x52, 0x08, 0x6d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x42, 0x0a, 0x0a, 0x08, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x22, 0x8c, 0x02, 0x0a, 0x11, 0x56, 0x69, 0x73, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x74,
	0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x15, 0x0a, 0x06, 0x73, 0x5f, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x73, 0x4e, 0x61, 0x6d, 0x65, 0x12,
	0x1d, 0x0a, 0x0a, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0c, 0x52, 0x09, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x4b, 0x65, 0x79, 0x12, 0x4e,
	0x0a, 0x0a, 0x76, 0x69, 0x73, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x2e, 0x2e, 0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76,
	0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x56, 0x69, 0x73, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x56, 0x69, 0x73, 0x69, 0x62, 0x69, 0x6c, 0x69,
	0x74, 0x79, 0x52, 0x0a, 0x76, 0x69, 0x73, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x22, 0x71,
	0x0a, 0x0a, 0x56, 0x69, 0x73, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x12, 0x1a, 0x0a, 0x16,
	0x56, 0x49, 0x53, 0x49, 0x42, 0x49, 0x4c, 0x49, 0x54, 0x59, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45,
	0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x18, 0x0a, 0x14, 0x56, 0x49, 0x53, 0x49,
	0x42, 0x49, 0x4c, 0x49, 0x54, 0x59, 0x5f, 0x41, 0x56, 0x41, 0x49, 0x4c, 0x41, 0x42, 0x4c, 0x45,
	0x10, 0x01, 0x12, 0x15, 0x0a, 0x11, 0x56, 0x49, 0x53, 0x49, 0x42, 0x49, 0x4c, 0x49, 0x54, 0x59,
	0x5f, 0x48, 0x49, 0x44, 0x44, 0x45, 0x4e, 0x10, 0x02, 0x12, 0x16, 0x0a, 0x12, 0x56, 0x49, 0x53,
	0x49, 0x42, 0x49, 0x4c, 0x49, 0x54, 0x59, 0x5f, 0x46, 0x52, 0x49, 0x45, 0x4e, 0x44, 0x53, 0x10,
	0x03, 0x22, 0x88, 0x02, 0x0a, 0x12, 0x56, 0x69, 0x73, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x4f, 0x0a, 0x0a, 0x76, 0x69, 0x73, 0x69,
	0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x2f, 0x2e, 0x6e,
	0x6f, 0x64, 0x65, 0x2e, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e,
	0x56, 0x69, 0x73, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x2e, 0x56, 0x69, 0x73, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x52, 0x0a, 0x76,
	0x69, 0x73, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x22, 0x71, 0x0a, 0x0a, 0x56, 0x69, 0x73,
	0x69, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x12, 0x1a, 0x0a, 0x16, 0x56, 0x49, 0x53, 0x49, 0x42,
	0x49, 0x4c, 0x49, 0x54, 0x59, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45,
	0x44, 0x10, 0x00, 0x12, 0x18, 0x0a, 0x14, 0x56, 0x49, 0x53, 0x49, 0x42, 0x49, 0x4c, 0x49, 0x54,
	0x59, 0x5f, 0x41, 0x56, 0x41, 0x49, 0x4c, 0x41, 0x42, 0x4c, 0x45, 0x10, 0x01, 0x12, 0x15, 0x0a,
	0x11, 0x56, 0x49, 0x53, 0x49, 0x42, 0x49, 0x4c, 0x49, 0x54, 0x59, 0x5f, 0x48, 0x49, 0x44, 0x44,
	0x45, 0x4e, 0x10, 0x02, 0x12, 0x16, 0x0a, 0x12, 0x56, 0x49, 0x53, 0x49, 0x42, 0x49, 0x4c, 0x49,
	0x54, 0x59, 0x5f, 0x46, 0x52, 0x49, 0x45, 0x4e, 0x44, 0x53, 0x10, 0x03, 0x42, 0x27, 0x5a, 0x25,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6e, 0x72, 0x2d,
	0x69, 0x6f, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x6e, 0x6f, 0x64, 0x65, 0x2f, 0x64, 0x69, 0x73,
	0x63, 0x6f, 0x76, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_node_discover_v1_discover_proto_rawDescOnce sync.Once
	file_node_discover_v1_discover_proto_rawDescData = file_node_discover_v1_discover_proto_rawDesc
)

func file_node_discover_v1_discover_proto_rawDescGZIP() []byte {
	file_node_discover_v1_discover_proto_rawDescOnce.Do(func() {
		file_node_discover_v1_discover_proto_rawDescData = protoimpl.X.CompressGZIP(file_node_discover_v1_discover_proto_rawDescData)
	})
	return file_node_discover_v1_discover_proto_rawDescData
}

var file_node_discover_v1_discover_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_node_discover_v1_discover_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_node_discover_v1_discover_proto_goTypes = []interface{}{
	(VisibilityRequest_Visibility)(0),  // 0: node.discover.v1.VisibilityRequest.Visibility
	(VisibilityResponse_Visibility)(0), // 1: node.discover.v1.VisibilityResponse.Visibility
	(*LobbyMessage)(nil),               // 2: node.discover.v1.LobbyMessage
	(*VisibilityRequest)(nil),          // 3: node.discover.v1.VisibilityRequest
	(*VisibilityResponse)(nil),         // 4: node.discover.v1.VisibilityResponse
	(*common.Peer)(nil),                // 5: common.Peer
	(*common.Metadata)(nil),            // 6: common.Metadata
}
var file_node_discover_v1_discover_proto_depIdxs = []int32{
	5, // 0: node.discover.v1.LobbyMessage.peer:type_name -> common.Peer
	6, // 1: node.discover.v1.LobbyMessage.metadata:type_name -> common.Metadata
	0, // 2: node.discover.v1.VisibilityRequest.visibility:type_name -> node.discover.v1.VisibilityRequest.Visibility
	1, // 3: node.discover.v1.VisibilityResponse.visibility:type_name -> node.discover.v1.VisibilityResponse.Visibility
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_node_discover_v1_discover_proto_init() }
func file_node_discover_v1_discover_proto_init() {
	if File_node_discover_v1_discover_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_node_discover_v1_discover_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LobbyMessage); i {
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
		file_node_discover_v1_discover_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VisibilityRequest); i {
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
		file_node_discover_v1_discover_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VisibilityResponse); i {
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
	file_node_discover_v1_discover_proto_msgTypes[0].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_node_discover_v1_discover_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_node_discover_v1_discover_proto_goTypes,
		DependencyIndexes: file_node_discover_v1_discover_proto_depIdxs,
		EnumInfos:         file_node_discover_v1_discover_proto_enumTypes,
		MessageInfos:      file_node_discover_v1_discover_proto_msgTypes,
	}.Build()
	File_node_discover_v1_discover_proto = out.File
	file_node_discover_v1_discover_proto_rawDesc = nil
	file_node_discover_v1_discover_proto_goTypes = nil
	file_node_discover_v1_discover_proto_depIdxs = nil
}