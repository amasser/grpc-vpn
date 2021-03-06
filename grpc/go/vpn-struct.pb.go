// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.12.1
// source: vpn-struct.proto

package vpn

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

type AuthType int32

const (
	AuthType_AT_NONE           AuthType = 0 // unknown
	AuthType_AT_TEST           AuthType = 1 // test
	AuthType_AT_GOOGLE_OPEN_ID AuthType = 2 // google open id
	AuthType_AT_AWS_IAM        AuthType = 3 // aws iam
)

// Enum value maps for AuthType.
var (
	AuthType_name = map[int32]string{
		0: "AT_NONE",
		1: "AT_TEST",
		2: "AT_GOOGLE_OPEN_ID",
		3: "AT_AWS_IAM",
	}
	AuthType_value = map[string]int32{
		"AT_NONE":           0,
		"AT_TEST":           1,
		"AT_GOOGLE_OPEN_ID": 2,
		"AT_AWS_IAM":        3,
	}
)

func (x AuthType) Enum() *AuthType {
	p := new(AuthType)
	*p = x
	return p
}

func (x AuthType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (AuthType) Descriptor() protoreflect.EnumDescriptor {
	return file_vpn_struct_proto_enumTypes[0].Descriptor()
}

func (AuthType) Type() protoreflect.EnumType {
	return &file_vpn_struct_proto_enumTypes[0]
}

func (x AuthType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use AuthType.Descriptor instead.
func (AuthType) EnumDescriptor() ([]byte, []int) {
	return file_vpn_struct_proto_rawDescGZIP(), []int{0}
}

type ErrorCode int32

const (
	ErrorCode_EC_UNKNOWN               ErrorCode = 0
	ErrorCode_EC_SUCCESS               ErrorCode = 1
	ErrorCode_EC_INVALID_AUTHORIZATION ErrorCode = 2
	ErrorCode_EC_EXPIRED_JWT           ErrorCode = 3
)

// Enum value maps for ErrorCode.
var (
	ErrorCode_name = map[int32]string{
		0: "EC_UNKNOWN",
		1: "EC_SUCCESS",
		2: "EC_INVALID_AUTHORIZATION",
		3: "EC_EXPIRED_JWT",
	}
	ErrorCode_value = map[string]int32{
		"EC_UNKNOWN":               0,
		"EC_SUCCESS":               1,
		"EC_INVALID_AUTHORIZATION": 2,
		"EC_EXPIRED_JWT":           3,
	}
)

func (x ErrorCode) Enum() *ErrorCode {
	p := new(ErrorCode)
	*p = x
	return p
}

func (x ErrorCode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ErrorCode) Descriptor() protoreflect.EnumDescriptor {
	return file_vpn_struct_proto_enumTypes[1].Descriptor()
}

func (ErrorCode) Type() protoreflect.EnumType {
	return &file_vpn_struct_proto_enumTypes[1]
}

func (x ErrorCode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ErrorCode.Descriptor instead.
func (ErrorCode) EnumDescriptor() ([]byte, []int) {
	return file_vpn_struct_proto_rawDescGZIP(), []int{1}
}

type IPPacketType int32

const (
	IPPacketType_IPPT_UNKNOWN    IPPacketType = 0
	IPPacketType_IPPT_RAW        IPPacketType = 1
	IPPacketType_IPPT_VPN_ASSIGN IPPacketType = 2
)

// Enum value maps for IPPacketType.
var (
	IPPacketType_name = map[int32]string{
		0: "IPPT_UNKNOWN",
		1: "IPPT_RAW",
		2: "IPPT_VPN_ASSIGN",
	}
	IPPacketType_value = map[string]int32{
		"IPPT_UNKNOWN":    0,
		"IPPT_RAW":        1,
		"IPPT_VPN_ASSIGN": 2,
	}
)

func (x IPPacketType) Enum() *IPPacketType {
	p := new(IPPacketType)
	*p = x
	return p
}

func (x IPPacketType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (IPPacketType) Descriptor() protoreflect.EnumDescriptor {
	return file_vpn_struct_proto_enumTypes[2].Descriptor()
}

func (IPPacketType) Type() protoreflect.EnumType {
	return &file_vpn_struct_proto_enumTypes[2]
}

func (x IPPacketType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use IPPacketType.Descriptor instead.
func (IPPacketType) EnumDescriptor() ([]byte, []int) {
	return file_vpn_struct_proto_rawDescGZIP(), []int{2}
}

type IPPacket struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ErrorCode  ErrorCode     `protobuf:"varint,1,opt,name=error_code,json=errorCode,proto3,enum=vpn.ErrorCode" json:"error_code,omitempty"`       // error code
	PacketType IPPacketType  `protobuf:"varint,2,opt,name=packet_type,json=packetType,proto3,enum=vpn.IPPacketType" json:"packet_type,omitempty"` // packet type
	Packet1    *IPPacket_Raw `protobuf:"bytes,10,opt,name=packet1,proto3" json:"packet1,omitempty"`                                               // raw packet
	Packet2    *IPPacket_Vpn `protobuf:"bytes,11,opt,name=packet2,proto3" json:"packet2,omitempty"`                                               // vpn packet
}

func (x *IPPacket) Reset() {
	*x = IPPacket{}
	if protoimpl.UnsafeEnabled {
		mi := &file_vpn_struct_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IPPacket) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IPPacket) ProtoMessage() {}

func (x *IPPacket) ProtoReflect() protoreflect.Message {
	mi := &file_vpn_struct_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IPPacket.ProtoReflect.Descriptor instead.
func (*IPPacket) Descriptor() ([]byte, []int) {
	return file_vpn_struct_proto_rawDescGZIP(), []int{0}
}

func (x *IPPacket) GetErrorCode() ErrorCode {
	if x != nil {
		return x.ErrorCode
	}
	return ErrorCode_EC_UNKNOWN
}

func (x *IPPacket) GetPacketType() IPPacketType {
	if x != nil {
		return x.PacketType
	}
	return IPPacketType_IPPT_UNKNOWN
}

func (x *IPPacket) GetPacket1() *IPPacket_Raw {
	if x != nil {
		return x.Packet1
	}
	return nil
}

func (x *IPPacket) GetPacket2() *IPPacket_Vpn {
	if x != nil {
		return x.Packet2
	}
	return nil
}

type AuthRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AuthType     AuthType                  `protobuf:"varint,1,opt,name=auth_type,json=authType,proto3,enum=vpn.AuthType" json:"auth_type,omitempty"` // auth type
	GoogleOpenId *AuthRequest_GoogleOpenID `protobuf:"bytes,2,opt,name=google_open_id,json=googleOpenId,proto3" json:"google_open_id,omitempty"`      // support google openid connect
	AwsIam       *AuthRequest_AwsIam       `protobuf:"bytes,3,opt,name=aws_iam,json=awsIam,proto3" json:"aws_iam,omitempty"`                          // support aws iam
}

func (x *AuthRequest) Reset() {
	*x = AuthRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_vpn_struct_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthRequest) ProtoMessage() {}

func (x *AuthRequest) ProtoReflect() protoreflect.Message {
	mi := &file_vpn_struct_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthRequest.ProtoReflect.Descriptor instead.
func (*AuthRequest) Descriptor() ([]byte, []int) {
	return file_vpn_struct_proto_rawDescGZIP(), []int{1}
}

func (x *AuthRequest) GetAuthType() AuthType {
	if x != nil {
		return x.AuthType
	}
	return AuthType_AT_NONE
}

func (x *AuthRequest) GetGoogleOpenId() *AuthRequest_GoogleOpenID {
	if x != nil {
		return x.GoogleOpenId
	}
	return nil
}

func (x *AuthRequest) GetAwsIam() *AuthRequest_AwsIam {
	if x != nil {
		return x.AwsIam
	}
	return nil
}

type AuthResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ErrorCode ErrorCode `protobuf:"varint,1,opt,name=error_code,json=errorCode,proto3,enum=vpn.ErrorCode" json:"error_code,omitempty"` // error code
	Jwt       string    `protobuf:"bytes,2,opt,name=jwt,proto3" json:"jwt,omitempty"`                                                  // jwt
}

func (x *AuthResponse) Reset() {
	*x = AuthResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_vpn_struct_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthResponse) ProtoMessage() {}

func (x *AuthResponse) ProtoReflect() protoreflect.Message {
	mi := &file_vpn_struct_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthResponse.ProtoReflect.Descriptor instead.
func (*AuthResponse) Descriptor() ([]byte, []int) {
	return file_vpn_struct_proto_rawDescGZIP(), []int{2}
}

func (x *AuthResponse) GetErrorCode() ErrorCode {
	if x != nil {
		return x.ErrorCode
	}
	return ErrorCode_EC_UNKNOWN
}

func (x *AuthResponse) GetJwt() string {
	if x != nil {
		return x.Jwt
	}
	return ""
}

type IPPacket_Raw struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Raw []byte `protobuf:"bytes,1,opt,name=raw,proto3" json:"raw,omitempty"` // raw packet
}

func (x *IPPacket_Raw) Reset() {
	*x = IPPacket_Raw{}
	if protoimpl.UnsafeEnabled {
		mi := &file_vpn_struct_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IPPacket_Raw) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IPPacket_Raw) ProtoMessage() {}

func (x *IPPacket_Raw) ProtoReflect() protoreflect.Message {
	mi := &file_vpn_struct_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IPPacket_Raw.ProtoReflect.Descriptor instead.
func (*IPPacket_Raw) Descriptor() ([]byte, []int) {
	return file_vpn_struct_proto_rawDescGZIP(), []int{0, 0}
}

func (x *IPPacket_Raw) GetRaw() []byte {
	if x != nil {
		return x.Raw
	}
	return nil
}

type IPPacket_Vpn struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	VpnAssignedIp []byte `protobuf:"bytes,1,opt,name=vpn_assigned_ip,json=vpnAssignedIp,proto3" json:"vpn_assigned_ip,omitempty"` // vpn  assigned ip
	VpnGateway    []byte `protobuf:"bytes,2,opt,name=vpn_gateway,json=vpnGateway,proto3" json:"vpn_gateway,omitempty"`            // vpn gateway
	VpnSubnetIp   []byte `protobuf:"bytes,3,opt,name=vpn_subnet_ip,json=vpnSubnetIp,proto3" json:"vpn_subnet_ip,omitempty"`       // vpn subnet ip
	VpnSubnetMask []byte `protobuf:"bytes,4,opt,name=vpn_subnet_mask,json=vpnSubnetMask,proto3" json:"vpn_subnet_mask,omitempty"` // vpn subnet mask
}

func (x *IPPacket_Vpn) Reset() {
	*x = IPPacket_Vpn{}
	if protoimpl.UnsafeEnabled {
		mi := &file_vpn_struct_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IPPacket_Vpn) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IPPacket_Vpn) ProtoMessage() {}

func (x *IPPacket_Vpn) ProtoReflect() protoreflect.Message {
	mi := &file_vpn_struct_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IPPacket_Vpn.ProtoReflect.Descriptor instead.
func (*IPPacket_Vpn) Descriptor() ([]byte, []int) {
	return file_vpn_struct_proto_rawDescGZIP(), []int{0, 1}
}

func (x *IPPacket_Vpn) GetVpnAssignedIp() []byte {
	if x != nil {
		return x.VpnAssignedIp
	}
	return nil
}

func (x *IPPacket_Vpn) GetVpnGateway() []byte {
	if x != nil {
		return x.VpnGateway
	}
	return nil
}

func (x *IPPacket_Vpn) GetVpnSubnetIp() []byte {
	if x != nil {
		return x.VpnSubnetIp
	}
	return nil
}

func (x *IPPacket_Vpn) GetVpnSubnetMask() []byte {
	if x != nil {
		return x.VpnSubnetMask
	}
	return nil
}

type AuthRequest_GoogleOpenID struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code string `protobuf:"bytes,1,opt,name=code,proto3" json:"code,omitempty"`
}

func (x *AuthRequest_GoogleOpenID) Reset() {
	*x = AuthRequest_GoogleOpenID{}
	if protoimpl.UnsafeEnabled {
		mi := &file_vpn_struct_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthRequest_GoogleOpenID) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthRequest_GoogleOpenID) ProtoMessage() {}

func (x *AuthRequest_GoogleOpenID) ProtoReflect() protoreflect.Message {
	mi := &file_vpn_struct_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthRequest_GoogleOpenID.ProtoReflect.Descriptor instead.
func (*AuthRequest_GoogleOpenID) Descriptor() ([]byte, []int) {
	return file_vpn_struct_proto_rawDescGZIP(), []int{1, 0}
}

func (x *AuthRequest_GoogleOpenID) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

type AuthRequest_AwsIam struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AccessKey       string `protobuf:"bytes,1,opt,name=access_key,json=accessKey,proto3" json:"access_key,omitempty"`
	SecretAccessKey string `protobuf:"bytes,2,opt,name=secret_access_key,json=secretAccessKey,proto3" json:"secret_access_key,omitempty"`
}

func (x *AuthRequest_AwsIam) Reset() {
	*x = AuthRequest_AwsIam{}
	if protoimpl.UnsafeEnabled {
		mi := &file_vpn_struct_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthRequest_AwsIam) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthRequest_AwsIam) ProtoMessage() {}

func (x *AuthRequest_AwsIam) ProtoReflect() protoreflect.Message {
	mi := &file_vpn_struct_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthRequest_AwsIam.ProtoReflect.Descriptor instead.
func (*AuthRequest_AwsIam) Descriptor() ([]byte, []int) {
	return file_vpn_struct_proto_rawDescGZIP(), []int{1, 1}
}

func (x *AuthRequest_AwsIam) GetAccessKey() string {
	if x != nil {
		return x.AccessKey
	}
	return ""
}

func (x *AuthRequest_AwsIam) GetSecretAccessKey() string {
	if x != nil {
		return x.SecretAccessKey
	}
	return ""
}

var File_vpn_struct_proto protoreflect.FileDescriptor

var file_vpn_struct_proto_rawDesc = []byte{
	0x0a, 0x10, 0x76, 0x70, 0x6e, 0x2d, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x03, 0x76, 0x70, 0x6e, 0x22, 0xfd, 0x02, 0x0a, 0x08, 0x49, 0x50, 0x50, 0x61,
	0x63, 0x6b, 0x65, 0x74, 0x12, 0x2d, 0x0a, 0x0a, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x5f, 0x63, 0x6f,
	0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0e, 0x2e, 0x76, 0x70, 0x6e, 0x2e, 0x45,
	0x72, 0x72, 0x6f, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x09, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x43,
	0x6f, 0x64, 0x65, 0x12, 0x32, 0x0a, 0x0b, 0x70, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x5f, 0x74, 0x79,
	0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x11, 0x2e, 0x76, 0x70, 0x6e, 0x2e, 0x49,
	0x50, 0x50, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0a, 0x70, 0x61, 0x63,
	0x6b, 0x65, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x2b, 0x0a, 0x07, 0x70, 0x61, 0x63, 0x6b, 0x65,
	0x74, 0x31, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x76, 0x70, 0x6e, 0x2e, 0x49,
	0x50, 0x50, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x2e, 0x52, 0x61, 0x77, 0x52, 0x07, 0x70, 0x61, 0x63,
	0x6b, 0x65, 0x74, 0x31, 0x12, 0x2b, 0x0a, 0x07, 0x70, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x32, 0x18,
	0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x76, 0x70, 0x6e, 0x2e, 0x49, 0x50, 0x50, 0x61,
	0x63, 0x6b, 0x65, 0x74, 0x2e, 0x56, 0x70, 0x6e, 0x52, 0x07, 0x70, 0x61, 0x63, 0x6b, 0x65, 0x74,
	0x32, 0x1a, 0x17, 0x0a, 0x03, 0x52, 0x61, 0x77, 0x12, 0x10, 0x0a, 0x03, 0x72, 0x61, 0x77, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x03, 0x72, 0x61, 0x77, 0x1a, 0x9a, 0x01, 0x0a, 0x03, 0x56,
	0x70, 0x6e, 0x12, 0x26, 0x0a, 0x0f, 0x76, 0x70, 0x6e, 0x5f, 0x61, 0x73, 0x73, 0x69, 0x67, 0x6e,
	0x65, 0x64, 0x5f, 0x69, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0d, 0x76, 0x70, 0x6e,
	0x41, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x65, 0x64, 0x49, 0x70, 0x12, 0x1f, 0x0a, 0x0b, 0x76, 0x70,
	0x6e, 0x5f, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52,
	0x0a, 0x76, 0x70, 0x6e, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x12, 0x22, 0x0a, 0x0d, 0x76,
	0x70, 0x6e, 0x5f, 0x73, 0x75, 0x62, 0x6e, 0x65, 0x74, 0x5f, 0x69, 0x70, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x0b, 0x76, 0x70, 0x6e, 0x53, 0x75, 0x62, 0x6e, 0x65, 0x74, 0x49, 0x70, 0x12,
	0x26, 0x0a, 0x0f, 0x76, 0x70, 0x6e, 0x5f, 0x73, 0x75, 0x62, 0x6e, 0x65, 0x74, 0x5f, 0x6d, 0x61,
	0x73, 0x6b, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0d, 0x76, 0x70, 0x6e, 0x53, 0x75, 0x62,
	0x6e, 0x65, 0x74, 0x4d, 0x61, 0x73, 0x6b, 0x22, 0xa9, 0x02, 0x0a, 0x0b, 0x41, 0x75, 0x74, 0x68,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2a, 0x0a, 0x09, 0x61, 0x75, 0x74, 0x68, 0x5f,
	0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0d, 0x2e, 0x76, 0x70, 0x6e,
	0x2e, 0x41, 0x75, 0x74, 0x68, 0x54, 0x79, 0x70, 0x65, 0x52, 0x08, 0x61, 0x75, 0x74, 0x68, 0x54,
	0x79, 0x70, 0x65, 0x12, 0x43, 0x0a, 0x0e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5f, 0x6f, 0x70,
	0x65, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x76, 0x70,
	0x6e, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x47, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x4f, 0x70, 0x65, 0x6e, 0x49, 0x44, 0x52, 0x0c, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x4f, 0x70, 0x65, 0x6e, 0x49, 0x64, 0x12, 0x30, 0x0a, 0x07, 0x61, 0x77, 0x73, 0x5f,
	0x69, 0x61, 0x6d, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x76, 0x70, 0x6e, 0x2e,
	0x41, 0x75, 0x74, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x41, 0x77, 0x73, 0x49,
	0x61, 0x6d, 0x52, 0x06, 0x61, 0x77, 0x73, 0x49, 0x61, 0x6d, 0x1a, 0x22, 0x0a, 0x0c, 0x47, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x4f, 0x70, 0x65, 0x6e, 0x49, 0x44, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f,
	0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x1a, 0x53,
	0x0a, 0x06, 0x41, 0x77, 0x73, 0x49, 0x61, 0x6d, 0x12, 0x1d, 0x0a, 0x0a, 0x61, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x61, 0x63,
	0x63, 0x65, 0x73, 0x73, 0x4b, 0x65, 0x79, 0x12, 0x2a, 0x0a, 0x11, 0x73, 0x65, 0x63, 0x72, 0x65,
	0x74, 0x5f, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0f, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73,
	0x4b, 0x65, 0x79, 0x22, 0x4f, 0x0a, 0x0c, 0x41, 0x75, 0x74, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x2d, 0x0a, 0x0a, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x5f, 0x63, 0x6f, 0x64,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0e, 0x2e, 0x76, 0x70, 0x6e, 0x2e, 0x45, 0x72,
	0x72, 0x6f, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x09, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x43, 0x6f,
	0x64, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6a, 0x77, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x6a, 0x77, 0x74, 0x2a, 0x4b, 0x0a, 0x08, 0x41, 0x75, 0x74, 0x68, 0x54, 0x79, 0x70, 0x65,
	0x12, 0x0b, 0x0a, 0x07, 0x41, 0x54, 0x5f, 0x4e, 0x4f, 0x4e, 0x45, 0x10, 0x00, 0x12, 0x0b, 0x0a,
	0x07, 0x41, 0x54, 0x5f, 0x54, 0x45, 0x53, 0x54, 0x10, 0x01, 0x12, 0x15, 0x0a, 0x11, 0x41, 0x54,
	0x5f, 0x47, 0x4f, 0x4f, 0x47, 0x4c, 0x45, 0x5f, 0x4f, 0x50, 0x45, 0x4e, 0x5f, 0x49, 0x44, 0x10,
	0x02, 0x12, 0x0e, 0x0a, 0x0a, 0x41, 0x54, 0x5f, 0x41, 0x57, 0x53, 0x5f, 0x49, 0x41, 0x4d, 0x10,
	0x03, 0x2a, 0x5d, 0x0a, 0x09, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x0e,
	0x0a, 0x0a, 0x45, 0x43, 0x5f, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x0e,
	0x0a, 0x0a, 0x45, 0x43, 0x5f, 0x53, 0x55, 0x43, 0x43, 0x45, 0x53, 0x53, 0x10, 0x01, 0x12, 0x1c,
	0x0a, 0x18, 0x45, 0x43, 0x5f, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x5f, 0x41, 0x55, 0x54,
	0x48, 0x4f, 0x52, 0x49, 0x5a, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x10, 0x02, 0x12, 0x12, 0x0a, 0x0e,
	0x45, 0x43, 0x5f, 0x45, 0x58, 0x50, 0x49, 0x52, 0x45, 0x44, 0x5f, 0x4a, 0x57, 0x54, 0x10, 0x03,
	0x2a, 0x43, 0x0a, 0x0c, 0x49, 0x50, 0x50, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x54, 0x79, 0x70, 0x65,
	0x12, 0x10, 0x0a, 0x0c, 0x49, 0x50, 0x50, 0x54, 0x5f, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e,
	0x10, 0x00, 0x12, 0x0c, 0x0a, 0x08, 0x49, 0x50, 0x50, 0x54, 0x5f, 0x52, 0x41, 0x57, 0x10, 0x01,
	0x12, 0x13, 0x0a, 0x0f, 0x49, 0x50, 0x50, 0x54, 0x5f, 0x56, 0x50, 0x4e, 0x5f, 0x41, 0x53, 0x53,
	0x49, 0x47, 0x4e, 0x10, 0x02, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_vpn_struct_proto_rawDescOnce sync.Once
	file_vpn_struct_proto_rawDescData = file_vpn_struct_proto_rawDesc
)

func file_vpn_struct_proto_rawDescGZIP() []byte {
	file_vpn_struct_proto_rawDescOnce.Do(func() {
		file_vpn_struct_proto_rawDescData = protoimpl.X.CompressGZIP(file_vpn_struct_proto_rawDescData)
	})
	return file_vpn_struct_proto_rawDescData
}

var file_vpn_struct_proto_enumTypes = make([]protoimpl.EnumInfo, 3)
var file_vpn_struct_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_vpn_struct_proto_goTypes = []interface{}{
	(AuthType)(0),                    // 0: vpn.AuthType
	(ErrorCode)(0),                   // 1: vpn.ErrorCode
	(IPPacketType)(0),                // 2: vpn.IPPacketType
	(*IPPacket)(nil),                 // 3: vpn.IPPacket
	(*AuthRequest)(nil),              // 4: vpn.AuthRequest
	(*AuthResponse)(nil),             // 5: vpn.AuthResponse
	(*IPPacket_Raw)(nil),             // 6: vpn.IPPacket.Raw
	(*IPPacket_Vpn)(nil),             // 7: vpn.IPPacket.Vpn
	(*AuthRequest_GoogleOpenID)(nil), // 8: vpn.AuthRequest.GoogleOpenID
	(*AuthRequest_AwsIam)(nil),       // 9: vpn.AuthRequest.AwsIam
}
var file_vpn_struct_proto_depIdxs = []int32{
	1, // 0: vpn.IPPacket.error_code:type_name -> vpn.ErrorCode
	2, // 1: vpn.IPPacket.packet_type:type_name -> vpn.IPPacketType
	6, // 2: vpn.IPPacket.packet1:type_name -> vpn.IPPacket.Raw
	7, // 3: vpn.IPPacket.packet2:type_name -> vpn.IPPacket.Vpn
	0, // 4: vpn.AuthRequest.auth_type:type_name -> vpn.AuthType
	8, // 5: vpn.AuthRequest.google_open_id:type_name -> vpn.AuthRequest.GoogleOpenID
	9, // 6: vpn.AuthRequest.aws_iam:type_name -> vpn.AuthRequest.AwsIam
	1, // 7: vpn.AuthResponse.error_code:type_name -> vpn.ErrorCode
	8, // [8:8] is the sub-list for method output_type
	8, // [8:8] is the sub-list for method input_type
	8, // [8:8] is the sub-list for extension type_name
	8, // [8:8] is the sub-list for extension extendee
	0, // [0:8] is the sub-list for field type_name
}

func init() { file_vpn_struct_proto_init() }
func file_vpn_struct_proto_init() {
	if File_vpn_struct_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_vpn_struct_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IPPacket); i {
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
		file_vpn_struct_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthRequest); i {
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
		file_vpn_struct_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthResponse); i {
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
		file_vpn_struct_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IPPacket_Raw); i {
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
		file_vpn_struct_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IPPacket_Vpn); i {
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
		file_vpn_struct_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthRequest_GoogleOpenID); i {
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
		file_vpn_struct_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthRequest_AwsIam); i {
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
			RawDescriptor: file_vpn_struct_proto_rawDesc,
			NumEnums:      3,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_vpn_struct_proto_goTypes,
		DependencyIndexes: file_vpn_struct_proto_depIdxs,
		EnumInfos:         file_vpn_struct_proto_enumTypes,
		MessageInfos:      file_vpn_struct_proto_msgTypes,
	}.Build()
	File_vpn_struct_proto = out.File
	file_vpn_struct_proto_rawDesc = nil
	file_vpn_struct_proto_goTypes = nil
	file_vpn_struct_proto_depIdxs = nil
}
