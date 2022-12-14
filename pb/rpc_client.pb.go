// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.10
// source: rpc_client.proto

package pb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ClientParamId struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *ClientParamId) Reset() {
	*x = ClientParamId{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_client_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClientParamId) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClientParamId) ProtoMessage() {}

func (x *ClientParamId) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_client_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClientParamId.ProtoReflect.Descriptor instead.
func (*ClientParamId) Descriptor() ([]byte, []int) {
	return file_rpc_client_proto_rawDescGZIP(), []int{0}
}

func (x *ClientParamId) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type Client struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id             string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Role           int32                  `protobuf:"varint,2,opt,name=role,proto3" json:"role,omitempty"`
	Name           string                 `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Email          string                 `protobuf:"bytes,4,opt,name=email,proto3" json:"email,omitempty"`
	PhoneNumber    string                 `protobuf:"bytes,5,opt,name=phone_number,json=phoneNumber,proto3" json:"phone_number,omitempty"`
	DateOfBith     string                 `protobuf:"bytes,6,opt,name=date_of_bith,json=dateOfBith,proto3" json:"date_of_bith,omitempty"`
	IdentityCard   string                 `protobuf:"bytes,7,opt,name=identity_card,json=identityCard,proto3" json:"identity_card,omitempty"`
	Address        string                 `protobuf:"bytes,8,opt,name=address,proto3" json:"address,omitempty"`
	MembershipCard string                 `protobuf:"bytes,9,opt,name=membership_card,json=membershipCard,proto3" json:"membership_card,omitempty"`
	Password       string                 `protobuf:"bytes,10,opt,name=password,proto3" json:"password,omitempty"`
	Status         int32                  `protobuf:"varint,11,opt,name=status,proto3" json:"status,omitempty"`
	CreatedAt      *timestamppb.Timestamp `protobuf:"bytes,12,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt      *timestamppb.Timestamp `protobuf:"bytes,13,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
}

func (x *Client) Reset() {
	*x = Client{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_client_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Client) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Client) ProtoMessage() {}

func (x *Client) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_client_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Client.ProtoReflect.Descriptor instead.
func (*Client) Descriptor() ([]byte, []int) {
	return file_rpc_client_proto_rawDescGZIP(), []int{1}
}

func (x *Client) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Client) GetRole() int32 {
	if x != nil {
		return x.Role
	}
	return 0
}

func (x *Client) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Client) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *Client) GetPhoneNumber() string {
	if x != nil {
		return x.PhoneNumber
	}
	return ""
}

func (x *Client) GetDateOfBith() string {
	if x != nil {
		return x.DateOfBith
	}
	return ""
}

func (x *Client) GetIdentityCard() string {
	if x != nil {
		return x.IdentityCard
	}
	return ""
}

func (x *Client) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *Client) GetMembershipCard() string {
	if x != nil {
		return x.MembershipCard
	}
	return ""
}

func (x *Client) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *Client) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *Client) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *Client) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

type SearchClientRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name         string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Email        string `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`
	PhoneNumber  string `protobuf:"bytes,3,opt,name=phone_number,json=phoneNumber,proto3" json:"phone_number,omitempty"`
	IdentityCard string `protobuf:"bytes,4,opt,name=identity_card,json=identityCard,proto3" json:"identity_card,omitempty"`
}

func (x *SearchClientRequest) Reset() {
	*x = SearchClientRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_client_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchClientRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchClientRequest) ProtoMessage() {}

func (x *SearchClientRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_client_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchClientRequest.ProtoReflect.Descriptor instead.
func (*SearchClientRequest) Descriptor() ([]byte, []int) {
	return file_rpc_client_proto_rawDescGZIP(), []int{2}
}

func (x *SearchClientRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *SearchClientRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *SearchClientRequest) GetPhoneNumber() string {
	if x != nil {
		return x.PhoneNumber
	}
	return ""
}

func (x *SearchClientRequest) GetIdentityCard() string {
	if x != nil {
		return x.IdentityCard
	}
	return ""
}

type SearchClientResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Client []*Client `protobuf:"bytes,1,rep,name=client,proto3" json:"client,omitempty"`
}

func (x *SearchClientResponse) Reset() {
	*x = SearchClientResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_client_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchClientResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchClientResponse) ProtoMessage() {}

func (x *SearchClientResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_client_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchClientResponse.ProtoReflect.Descriptor instead.
func (*SearchClientResponse) Descriptor() ([]byte, []int) {
	return file_rpc_client_proto_rawDescGZIP(), []int{3}
}

func (x *SearchClientResponse) GetClient() []*Client {
	if x != nil {
		return x.Client
	}
	return nil
}

type ChangePasswordRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ClientId        string `protobuf:"bytes,1,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty"`
	OldPassword     string `protobuf:"bytes,2,opt,name=old_password,json=oldPassword,proto3" json:"old_password,omitempty"`
	NewPassword     string `protobuf:"bytes,3,opt,name=new_password,json=newPassword,proto3" json:"new_password,omitempty"`
	ConfirmPassword string `protobuf:"bytes,4,opt,name=confirm_password,json=confirmPassword,proto3" json:"confirm_password,omitempty"`
}

func (x *ChangePasswordRequest) Reset() {
	*x = ChangePasswordRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_client_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChangePasswordRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChangePasswordRequest) ProtoMessage() {}

func (x *ChangePasswordRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_client_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChangePasswordRequest.ProtoReflect.Descriptor instead.
func (*ChangePasswordRequest) Descriptor() ([]byte, []int) {
	return file_rpc_client_proto_rawDescGZIP(), []int{4}
}

func (x *ChangePasswordRequest) GetClientId() string {
	if x != nil {
		return x.ClientId
	}
	return ""
}

func (x *ChangePasswordRequest) GetOldPassword() string {
	if x != nil {
		return x.OldPassword
	}
	return ""
}

func (x *ChangePasswordRequest) GetNewPassword() string {
	if x != nil {
		return x.NewPassword
	}
	return ""
}

func (x *ChangePasswordRequest) GetConfirmPassword() string {
	if x != nil {
		return x.ConfirmPassword
	}
	return ""
}

type ChangePasswordResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    int32  `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *ChangePasswordResponse) Reset() {
	*x = ChangePasswordResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_client_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChangePasswordResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChangePasswordResponse) ProtoMessage() {}

func (x *ChangePasswordResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_client_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChangePasswordResponse.ProtoReflect.Descriptor instead.
func (*ChangePasswordResponse) Descriptor() ([]byte, []int) {
	return file_rpc_client_proto_rawDescGZIP(), []int{5}
}

func (x *ChangePasswordResponse) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *ChangePasswordResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_rpc_client_proto protoreflect.FileDescriptor

var file_rpc_client_proto_rawDesc = []byte{
	0x0a, 0x10, 0x72, 0x70, 0x63, 0x5f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x0e, 0x66, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x5f, 0x62, 0x6f, 0x6f, 0x6b, 0x69,
	0x6e, 0x67, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x1f, 0x0a, 0x0d, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x50, 0x61, 0x72,
	0x61, 0x6d, 0x49, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x69, 0x64, 0x22, 0xad, 0x03, 0x0a, 0x06, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x12, 0x0a, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x72,
	0x6f, 0x6c, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x21, 0x0a,
	0x0c, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72,
	0x12, 0x20, 0x0a, 0x0c, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x6f, 0x66, 0x5f, 0x62, 0x69, 0x74, 0x68,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x64, 0x61, 0x74, 0x65, 0x4f, 0x66, 0x42, 0x69,
	0x74, 0x68, 0x12, 0x23, 0x0a, 0x0d, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x5f, 0x63,
	0x61, 0x72, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x69, 0x64, 0x65, 0x6e, 0x74,
	0x69, 0x74, 0x79, 0x43, 0x61, 0x72, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65,
	0x73, 0x73, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x12, 0x27, 0x0a, 0x0f, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x68, 0x69, 0x70, 0x5f,
	0x63, 0x61, 0x72, 0x64, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x6d, 0x65, 0x6d, 0x62,
	0x65, 0x72, 0x73, 0x68, 0x69, 0x70, 0x43, 0x61, 0x72, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61,
	0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61,
	0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x18, 0x0b, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x39,
	0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x0c, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x39, 0x0a, 0x0a, 0x75, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x64, 0x41, 0x74, 0x22, 0x87, 0x01, 0x0a, 0x13, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x43,
	0x6c, 0x69, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x21, 0x0a, 0x0c, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x5f,
	0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x68,
	0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x23, 0x0a, 0x0d, 0x69, 0x64, 0x65,
	0x6e, 0x74, 0x69, 0x74, 0x79, 0x5f, 0x63, 0x61, 0x72, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0c, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x43, 0x61, 0x72, 0x64, 0x22, 0x46,
	0x0a, 0x14, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2e, 0x0a, 0x06, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x66, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x5f,
	0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x52, 0x06,
	0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x22, 0xa5, 0x01, 0x0a, 0x15, 0x43, 0x68, 0x61, 0x6e, 0x67,
	0x65, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x1b, 0x0a, 0x09, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x21, 0x0a,
	0x0c, 0x6f, 0x6c, 0x64, 0x5f, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x6f, 0x6c, 0x64, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64,
	0x12, 0x21, 0x0a, 0x0c, 0x6e, 0x65, 0x77, 0x5f, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x6e, 0x65, 0x77, 0x50, 0x61, 0x73, 0x73, 0x77,
	0x6f, 0x72, 0x64, 0x12, 0x29, 0x0a, 0x10, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x5f, 0x70,
	0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x63,
	0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x22, 0x46,
	0x0a, 0x16, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x32, 0x8a, 0x03, 0x0a, 0x09, 0x52, 0x50, 0x43, 0x43, 0x6c,
	0x69, 0x65, 0x6e, 0x74, 0x12, 0x41, 0x0a, 0x08, 0x46, 0x69, 0x6e, 0x64, 0x42, 0x79, 0x49, 0x64,
	0x12, 0x1d, 0x2e, 0x66, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x5f, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e,
	0x67, 0x2e, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x49, 0x64, 0x1a,
	0x16, 0x2e, 0x66, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x5f, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67,
	0x2e, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x12, 0x3e, 0x0a, 0x0c, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x12, 0x16, 0x2e, 0x66, 0x6c, 0x69, 0x67, 0x68, 0x74,
	0x5f, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x1a,
	0x16, 0x2e, 0x66, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x5f, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67,
	0x2e, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x12, 0x3e, 0x0a, 0x0c, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x12, 0x16, 0x2e, 0x66, 0x6c, 0x69, 0x67, 0x68, 0x74,
	0x5f, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x1a,
	0x16, 0x2e, 0x66, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x5f, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67,
	0x2e, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x12, 0x5f, 0x0a, 0x0e, 0x43, 0x68, 0x61, 0x6e, 0x67,
	0x65, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x25, 0x2e, 0x66, 0x6c, 0x69, 0x67,
	0x68, 0x74, 0x5f, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x43, 0x68, 0x61, 0x6e, 0x67,
	0x65, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x26, 0x2e, 0x66, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x5f, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e,
	0x67, 0x2e, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x59, 0x0a, 0x0c, 0x53, 0x65, 0x61, 0x72,
	0x63, 0x68, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x12, 0x23, 0x2e, 0x66, 0x6c, 0x69, 0x67, 0x68,
	0x74, 0x5f, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68,
	0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e,
	0x66, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x5f, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x53,
	0x65, 0x61, 0x72, 0x63, 0x68, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x42, 0x07, 0x5a, 0x05, 0x2e, 0x2f, 0x3b, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_rpc_client_proto_rawDescOnce sync.Once
	file_rpc_client_proto_rawDescData = file_rpc_client_proto_rawDesc
)

func file_rpc_client_proto_rawDescGZIP() []byte {
	file_rpc_client_proto_rawDescOnce.Do(func() {
		file_rpc_client_proto_rawDescData = protoimpl.X.CompressGZIP(file_rpc_client_proto_rawDescData)
	})
	return file_rpc_client_proto_rawDescData
}

var file_rpc_client_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_rpc_client_proto_goTypes = []interface{}{
	(*ClientParamId)(nil),          // 0: flight_booking.ClientParamId
	(*Client)(nil),                 // 1: flight_booking.Client
	(*SearchClientRequest)(nil),    // 2: flight_booking.SearchClientRequest
	(*SearchClientResponse)(nil),   // 3: flight_booking.SearchClientResponse
	(*ChangePasswordRequest)(nil),  // 4: flight_booking.ChangePasswordRequest
	(*ChangePasswordResponse)(nil), // 5: flight_booking.ChangePasswordResponse
	(*timestamppb.Timestamp)(nil),  // 6: google.protobuf.Timestamp
}
var file_rpc_client_proto_depIdxs = []int32{
	6, // 0: flight_booking.Client.created_at:type_name -> google.protobuf.Timestamp
	6, // 1: flight_booking.Client.updated_at:type_name -> google.protobuf.Timestamp
	1, // 2: flight_booking.SearchClientResponse.client:type_name -> flight_booking.Client
	0, // 3: flight_booking.RPCClient.FindById:input_type -> flight_booking.ClientParamId
	1, // 4: flight_booking.RPCClient.CreateClient:input_type -> flight_booking.Client
	1, // 5: flight_booking.RPCClient.UpdateClient:input_type -> flight_booking.Client
	4, // 6: flight_booking.RPCClient.ChangePassword:input_type -> flight_booking.ChangePasswordRequest
	2, // 7: flight_booking.RPCClient.SearchClient:input_type -> flight_booking.SearchClientRequest
	1, // 8: flight_booking.RPCClient.FindById:output_type -> flight_booking.Client
	1, // 9: flight_booking.RPCClient.CreateClient:output_type -> flight_booking.Client
	1, // 10: flight_booking.RPCClient.UpdateClient:output_type -> flight_booking.Client
	5, // 11: flight_booking.RPCClient.ChangePassword:output_type -> flight_booking.ChangePasswordResponse
	3, // 12: flight_booking.RPCClient.SearchClient:output_type -> flight_booking.SearchClientResponse
	8, // [8:13] is the sub-list for method output_type
	3, // [3:8] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_rpc_client_proto_init() }
func file_rpc_client_proto_init() {
	if File_rpc_client_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_rpc_client_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClientParamId); i {
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
		file_rpc_client_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Client); i {
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
		file_rpc_client_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchClientRequest); i {
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
		file_rpc_client_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchClientResponse); i {
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
		file_rpc_client_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChangePasswordRequest); i {
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
		file_rpc_client_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChangePasswordResponse); i {
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
			RawDescriptor: file_rpc_client_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_rpc_client_proto_goTypes,
		DependencyIndexes: file_rpc_client_proto_depIdxs,
		MessageInfos:      file_rpc_client_proto_msgTypes,
	}.Build()
	File_rpc_client_proto = out.File
	file_rpc_client_proto_rawDesc = nil
	file_rpc_client_proto_goTypes = nil
	file_rpc_client_proto_depIdxs = nil
}
