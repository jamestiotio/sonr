// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: proto/protocols/exchange.proto

package exchange

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

type VisibilityRequest_Visibility int32

const (
	VisibilityRequest_VISIBLE VisibilityRequest_Visibility = 0 // Everyone can see this peer
	VisibilityRequest_GHOST   VisibilityRequest_Visibility = 1 // Only Linked Devices can see this peer
	VisibilityRequest_FRIENDS VisibilityRequest_Visibility = 2 // Only Friends can see this peer
)

// Enum value maps for VisibilityRequest_Visibility.
var (
	VisibilityRequest_Visibility_name = map[int32]string{
		0: "VISIBLE",
		1: "GHOST",
		2: "FRIENDS",
	}
	VisibilityRequest_Visibility_value = map[string]int32{
		"VISIBLE": 0,
		"GHOST":   1,
		"FRIENDS": 2,
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
	return file_proto_protocols_exchange_proto_enumTypes[0].Descriptor()
}

func (VisibilityRequest_Visibility) Type() protoreflect.EnumType {
	return &file_proto_protocols_exchange_proto_enumTypes[0]
}

func (x VisibilityRequest_Visibility) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use VisibilityRequest_Visibility.Descriptor instead.
func (VisibilityRequest_Visibility) EnumDescriptor() ([]byte, []int) {
	return file_proto_protocols_exchange_proto_rawDescGZIP(), []int{4, 0}
}

type VisibilityResponse_Visibility int32

const (
	VisibilityResponse_VISIBLE VisibilityResponse_Visibility = 0 // Everyone can see this peer
	VisibilityResponse_GHOST   VisibilityResponse_Visibility = 1 // Only Linked Devices can see this peer
	VisibilityResponse_FRIENDS VisibilityResponse_Visibility = 2 // Only Friends can see this peer
)

// Enum value maps for VisibilityResponse_Visibility.
var (
	VisibilityResponse_Visibility_name = map[int32]string{
		0: "VISIBLE",
		1: "GHOST",
		2: "FRIENDS",
	}
	VisibilityResponse_Visibility_value = map[string]int32{
		"VISIBLE": 0,
		"GHOST":   1,
		"FRIENDS": 2,
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
	return file_proto_protocols_exchange_proto_enumTypes[1].Descriptor()
}

func (VisibilityResponse_Visibility) Type() protoreflect.EnumType {
	return &file_proto_protocols_exchange_proto_enumTypes[1]
}

func (x VisibilityResponse_Visibility) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use VisibilityResponse_Visibility.Descriptor instead.
func (VisibilityResponse_Visibility) EnumDescriptor() ([]byte, []int) {
	return file_proto_protocols_exchange_proto_rawDescGZIP(), []int{5, 0}
}

// AuthorizeRequest is Message for Verifying Request (Hmac Sha256)
type AuthorizeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SName    string `protobuf:"bytes,1,opt,name=sName,proto3" json:"sName,omitempty"`       // SName combined with Device ID and Hashed
	Mnemonic string `protobuf:"bytes,2,opt,name=mnemonic,proto3" json:"mnemonic,omitempty"` // Mnemonic Hashed with public key for fingerprint
	DeviceId string `protobuf:"bytes,3,opt,name=deviceId,proto3" json:"deviceId,omitempty"` // Device ID
}

func (x *AuthorizeRequest) Reset() {
	*x = AuthorizeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_protocols_exchange_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthorizeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthorizeRequest) ProtoMessage() {}

func (x *AuthorizeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_protocols_exchange_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthorizeRequest.ProtoReflect.Descriptor instead.
func (*AuthorizeRequest) Descriptor() ([]byte, []int) {
	return file_proto_protocols_exchange_proto_rawDescGZIP(), []int{0}
}

func (x *AuthorizeRequest) GetSName() string {
	if x != nil {
		return x.SName
	}
	return ""
}

func (x *AuthorizeRequest) GetMnemonic() string {
	if x != nil {
		return x.Mnemonic
	}
	return ""
}

func (x *AuthorizeRequest) GetDeviceId() string {
	if x != nil {
		return x.DeviceId
	}
	return ""
}

// AuthorizeResponse is Message for Verifying Response (Hmac Sha256)
type AuthorizeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool              `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`                                                                                        // If Values were Verified
	Error   string            `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`                                                                                             // Error Message
	Records map[string]string `protobuf:"bytes,3,rep,name=records,proto3" json:"records,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"` // Updated Domain TXT Records
}

func (x *AuthorizeResponse) Reset() {
	*x = AuthorizeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_protocols_exchange_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthorizeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthorizeResponse) ProtoMessage() {}

func (x *AuthorizeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_protocols_exchange_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthorizeResponse.ProtoReflect.Descriptor instead.
func (*AuthorizeResponse) Descriptor() ([]byte, []int) {
	return file_proto_protocols_exchange_proto_rawDescGZIP(), []int{1}
}

func (x *AuthorizeResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *AuthorizeResponse) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

func (x *AuthorizeResponse) GetRecords() map[string]string {
	if x != nil {
		return x.Records
	}
	return nil
}

// LookupRequest is Message for Verifying Request (Hmac Sha256)
type LookupRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SName string `protobuf:"bytes,1,opt,name=sName,proto3" json:"sName,omitempty"` // SName combined with Device ID and Hashed
}

func (x *LookupRequest) Reset() {
	*x = LookupRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_protocols_exchange_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LookupRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LookupRequest) ProtoMessage() {}

func (x *LookupRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_protocols_exchange_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LookupRequest.ProtoReflect.Descriptor instead.
func (*LookupRequest) Descriptor() ([]byte, []int) {
	return file_proto_protocols_exchange_proto_rawDescGZIP(), []int{2}
}

func (x *LookupRequest) GetSName() string {
	if x != nil {
		return x.SName
	}
	return ""
}

// LookupResponse is Message for Verifying Response (Hmac Sha256)
type LookupResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success   bool   `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`    // If Values were Verified
	PublicKey string `protobuf:"bytes,2,opt,name=publicKey,proto3" json:"publicKey,omitempty"` // Base64 Encoded Public Key
}

func (x *LookupResponse) Reset() {
	*x = LookupResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_protocols_exchange_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LookupResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LookupResponse) ProtoMessage() {}

func (x *LookupResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_protocols_exchange_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LookupResponse.ProtoReflect.Descriptor instead.
func (*LookupResponse) Descriptor() ([]byte, []int) {
	return file_proto_protocols_exchange_proto_rawDescGZIP(), []int{3}
}

func (x *LookupResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *LookupResponse) GetPublicKey() string {
	if x != nil {
		return x.PublicKey
	}
	return ""
}

// VisibilityRequest is Message for updating Peer Visibility in Exchange
type VisibilityRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SName      string                       `protobuf:"bytes,1,opt,name=sName,proto3" json:"sName,omitempty"`                                                                      // SName combined with Device ID and Hashed
	PublicKey  []byte                       `protobuf:"bytes,2,opt,name=publicKey,proto3" json:"publicKey,omitempty"`                                                              // Buffer of Public Key
	Visibility VisibilityRequest_Visibility `protobuf:"varint,3,opt,name=visibility,proto3,enum=sonr.protocols.exchange.VisibilityRequest_Visibility" json:"visibility,omitempty"` // Visibility
}

func (x *VisibilityRequest) Reset() {
	*x = VisibilityRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_protocols_exchange_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VisibilityRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VisibilityRequest) ProtoMessage() {}

func (x *VisibilityRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_protocols_exchange_proto_msgTypes[4]
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
	return file_proto_protocols_exchange_proto_rawDescGZIP(), []int{4}
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
	return VisibilityRequest_VISIBLE
}

// VisibilityResponse is response for VisibilityRequest
type VisibilityResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success    bool                          `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`                                                                  // If Request was Successful
	Error      string                        `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`                                                                       // Error Message if Request was not successful
	Visibility VisibilityResponse_Visibility `protobuf:"varint,3,opt,name=visibility,proto3,enum=sonr.protocols.exchange.VisibilityResponse_Visibility" json:"visibility,omitempty"` // Visibility
}

func (x *VisibilityResponse) Reset() {
	*x = VisibilityResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_protocols_exchange_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VisibilityResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VisibilityResponse) ProtoMessage() {}

func (x *VisibilityResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_protocols_exchange_proto_msgTypes[5]
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
	return file_proto_protocols_exchange_proto_rawDescGZIP(), []int{5}
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
	return VisibilityResponse_VISIBLE
}

var File_proto_protocols_exchange_proto protoreflect.FileDescriptor

var file_proto_protocols_exchange_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c,
	0x73, 0x2f, 0x65, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x17, 0x73, 0x6f, 0x6e, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x73,
	0x2e, 0x65, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x22, 0x60, 0x0a, 0x10, 0x41, 0x75, 0x74,
	0x68, 0x6f, 0x72, 0x69, 0x7a, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a,
	0x05, 0x73, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x73, 0x4e,
	0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6d, 0x6e, 0x65, 0x6d, 0x6f, 0x6e, 0x69, 0x63, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6d, 0x6e, 0x65, 0x6d, 0x6f, 0x6e, 0x69, 0x63, 0x12,
	0x1a, 0x0a, 0x08, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x49, 0x64, 0x22, 0xd2, 0x01, 0x0a, 0x11,
	0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x65,
	0x72, 0x72, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f,
	0x72, 0x12, 0x51, 0x0a, 0x07, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x73, 0x18, 0x03, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x37, 0x2e, 0x73, 0x6f, 0x6e, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63,
	0x6f, 0x6c, 0x73, 0x2e, 0x65, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x2e, 0x41, 0x75, 0x74,
	0x68, 0x6f, 0x72, 0x69, 0x7a, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x52,
	0x65, 0x63, 0x6f, 0x72, 0x64, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x07, 0x72, 0x65, 0x63,
	0x6f, 0x72, 0x64, 0x73, 0x1a, 0x3a, 0x0a, 0x0c, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x73, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01,
	0x22, 0x25, 0x0a, 0x0d, 0x4c, 0x6f, 0x6f, 0x6b, 0x75, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x73, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x48, 0x0a, 0x0e, 0x4c, 0x6f, 0x6f, 0x6b, 0x75,
	0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63,
	0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x12, 0x1c, 0x0a, 0x09, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x4b, 0x65, 0x79,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x4b, 0x65,
	0x79, 0x22, 0xd1, 0x01, 0x0a, 0x11, 0x56, 0x69, 0x73, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x4e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x73, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a,
	0x09, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x4b, 0x65, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c,
	0x52, 0x09, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x4b, 0x65, 0x79, 0x12, 0x55, 0x0a, 0x0a, 0x76,
	0x69, 0x73, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x35, 0x2e, 0x73, 0x6f, 0x6e, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x73,
	0x2e, 0x65, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x2e, 0x56, 0x69, 0x73, 0x69, 0x62, 0x69,
	0x6c, 0x69, 0x74, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x56, 0x69, 0x73, 0x69,
	0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x52, 0x0a, 0x76, 0x69, 0x73, 0x69, 0x62, 0x69, 0x6c, 0x69,
	0x74, 0x79, 0x22, 0x31, 0x0a, 0x0a, 0x56, 0x69, 0x73, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79,
	0x12, 0x0b, 0x0a, 0x07, 0x56, 0x49, 0x53, 0x49, 0x42, 0x4c, 0x45, 0x10, 0x00, 0x12, 0x09, 0x0a,
	0x05, 0x47, 0x48, 0x4f, 0x53, 0x54, 0x10, 0x01, 0x12, 0x0b, 0x0a, 0x07, 0x46, 0x52, 0x49, 0x45,
	0x4e, 0x44, 0x53, 0x10, 0x02, 0x22, 0xcf, 0x01, 0x0a, 0x12, 0x56, 0x69, 0x73, 0x69, 0x62, 0x69,
	0x6c, 0x69, 0x74, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07,
	0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73,
	0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x56, 0x0a, 0x0a,
	0x76, 0x69, 0x73, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x36, 0x2e, 0x73, 0x6f, 0x6e, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c,
	0x73, 0x2e, 0x65, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x2e, 0x56, 0x69, 0x73, 0x69, 0x62,
	0x69, 0x6c, 0x69, 0x74, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x56, 0x69,
	0x73, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x52, 0x0a, 0x76, 0x69, 0x73, 0x69, 0x62, 0x69,
	0x6c, 0x69, 0x74, 0x79, 0x22, 0x31, 0x0a, 0x0a, 0x56, 0x69, 0x73, 0x69, 0x62, 0x69, 0x6c, 0x69,
	0x74, 0x79, 0x12, 0x0b, 0x0a, 0x07, 0x56, 0x49, 0x53, 0x49, 0x42, 0x4c, 0x45, 0x10, 0x00, 0x12,
	0x09, 0x0a, 0x05, 0x47, 0x48, 0x4f, 0x53, 0x54, 0x10, 0x01, 0x12, 0x0b, 0x0a, 0x07, 0x46, 0x52,
	0x49, 0x45, 0x4e, 0x44, 0x53, 0x10, 0x02, 0x42, 0x26, 0x5a, 0x24, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6e, 0x72, 0x2d, 0x69, 0x6f, 0x2f, 0x63, 0x6f,
	0x72, 0x65, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x65, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_protocols_exchange_proto_rawDescOnce sync.Once
	file_proto_protocols_exchange_proto_rawDescData = file_proto_protocols_exchange_proto_rawDesc
)

func file_proto_protocols_exchange_proto_rawDescGZIP() []byte {
	file_proto_protocols_exchange_proto_rawDescOnce.Do(func() {
		file_proto_protocols_exchange_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_protocols_exchange_proto_rawDescData)
	})
	return file_proto_protocols_exchange_proto_rawDescData
}

var file_proto_protocols_exchange_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_proto_protocols_exchange_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_proto_protocols_exchange_proto_goTypes = []interface{}{
	(VisibilityRequest_Visibility)(0),  // 0: sonr.protocols.exchange.VisibilityRequest.Visibility
	(VisibilityResponse_Visibility)(0), // 1: sonr.protocols.exchange.VisibilityResponse.Visibility
	(*AuthorizeRequest)(nil),           // 2: sonr.protocols.exchange.AuthorizeRequest
	(*AuthorizeResponse)(nil),          // 3: sonr.protocols.exchange.AuthorizeResponse
	(*LookupRequest)(nil),              // 4: sonr.protocols.exchange.LookupRequest
	(*LookupResponse)(nil),             // 5: sonr.protocols.exchange.LookupResponse
	(*VisibilityRequest)(nil),          // 6: sonr.protocols.exchange.VisibilityRequest
	(*VisibilityResponse)(nil),         // 7: sonr.protocols.exchange.VisibilityResponse
	nil,                                // 8: sonr.protocols.exchange.AuthorizeResponse.RecordsEntry
}
var file_proto_protocols_exchange_proto_depIdxs = []int32{
	8, // 0: sonr.protocols.exchange.AuthorizeResponse.records:type_name -> sonr.protocols.exchange.AuthorizeResponse.RecordsEntry
	0, // 1: sonr.protocols.exchange.VisibilityRequest.visibility:type_name -> sonr.protocols.exchange.VisibilityRequest.Visibility
	1, // 2: sonr.protocols.exchange.VisibilityResponse.visibility:type_name -> sonr.protocols.exchange.VisibilityResponse.Visibility
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_proto_protocols_exchange_proto_init() }
func file_proto_protocols_exchange_proto_init() {
	if File_proto_protocols_exchange_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_protocols_exchange_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthorizeRequest); i {
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
		file_proto_protocols_exchange_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthorizeResponse); i {
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
		file_proto_protocols_exchange_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LookupRequest); i {
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
		file_proto_protocols_exchange_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LookupResponse); i {
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
		file_proto_protocols_exchange_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
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
		file_proto_protocols_exchange_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
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
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_protocols_exchange_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_protocols_exchange_proto_goTypes,
		DependencyIndexes: file_proto_protocols_exchange_proto_depIdxs,
		EnumInfos:         file_proto_protocols_exchange_proto_enumTypes,
		MessageInfos:      file_proto_protocols_exchange_proto_msgTypes,
	}.Build()
	File_proto_protocols_exchange_proto = out.File
	file_proto_protocols_exchange_proto_rawDesc = nil
	file_proto_protocols_exchange_proto_goTypes = nil
	file_proto_protocols_exchange_proto_depIdxs = nil
}