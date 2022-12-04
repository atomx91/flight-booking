// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.10
// source: rpc_booking.proto

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

type BookingParamId struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *BookingParamId) Reset() {
	*x = BookingParamId{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_booking_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BookingParamId) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BookingParamId) ProtoMessage() {}

func (x *BookingParamId) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_booking_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BookingParamId.ProtoReflect.Descriptor instead.
func (*BookingParamId) Descriptor() ([]byte, []int) {
	return file_rpc_booking_proto_rawDescGZIP(), []int{0}
}

func (x *BookingParamId) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type ClientDTO struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id             string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Role           int32  `protobuf:"varint,2,opt,name=role,proto3" json:"role,omitempty"`
	Name           string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Email          string `protobuf:"bytes,4,opt,name=email,proto3" json:"email,omitempty"`
	PhoneNumber    string `protobuf:"bytes,5,opt,name=phone_number,json=phoneNumber,proto3" json:"phone_number,omitempty"`
	DateOfBith     string `protobuf:"bytes,6,opt,name=date_of_bith,json=dateOfBith,proto3" json:"date_of_bith,omitempty"`
	IdentityCard   string `protobuf:"bytes,7,opt,name=identity_card,json=identityCard,proto3" json:"identity_card,omitempty"`
	Address        string `protobuf:"bytes,8,opt,name=address,proto3" json:"address,omitempty"`
	MembershipCard string `protobuf:"bytes,9,opt,name=membership_card,json=membershipCard,proto3" json:"membership_card,omitempty"`
	Password       string `protobuf:"bytes,10,opt,name=password,proto3" json:"password,omitempty"`
	Status         int32  `protobuf:"varint,11,opt,name=status,proto3" json:"status,omitempty"`
	Audit          *Audit `protobuf:"bytes,12,opt,name=audit,proto3" json:"audit,omitempty"`
}

func (x *ClientDTO) Reset() {
	*x = ClientDTO{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_booking_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClientDTO) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClientDTO) ProtoMessage() {}

func (x *ClientDTO) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_booking_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClientDTO.ProtoReflect.Descriptor instead.
func (*ClientDTO) Descriptor() ([]byte, []int) {
	return file_rpc_booking_proto_rawDescGZIP(), []int{1}
}

func (x *ClientDTO) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ClientDTO) GetRole() int32 {
	if x != nil {
		return x.Role
	}
	return 0
}

func (x *ClientDTO) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ClientDTO) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *ClientDTO) GetPhoneNumber() string {
	if x != nil {
		return x.PhoneNumber
	}
	return ""
}

func (x *ClientDTO) GetDateOfBith() string {
	if x != nil {
		return x.DateOfBith
	}
	return ""
}

func (x *ClientDTO) GetIdentityCard() string {
	if x != nil {
		return x.IdentityCard
	}
	return ""
}

func (x *ClientDTO) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *ClientDTO) GetMembershipCard() string {
	if x != nil {
		return x.MembershipCard
	}
	return ""
}

func (x *ClientDTO) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *ClientDTO) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *ClientDTO) GetAudit() *Audit {
	if x != nil {
		return x.Audit
	}
	return nil
}

type FlightDTO struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name          string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	From          string                 `protobuf:"bytes,3,opt,name=from,proto3" json:"from,omitempty"`
	To            string                 `protobuf:"bytes,4,opt,name=to,proto3" json:"to,omitempty"`
	DepartDate    *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=depart_date,json=departDate,proto3" json:"depart_date,omitempty"`
	Status        string                 `protobuf:"bytes,6,opt,name=status,proto3" json:"status,omitempty"`
	AvailableSlot int32                  `protobuf:"varint,7,opt,name=available_slot,json=availableSlot,proto3" json:"available_slot,omitempty"`
	CreatedAt     *timestamppb.Timestamp `protobuf:"bytes,8,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt     *timestamppb.Timestamp `protobuf:"bytes,9,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
}

func (x *FlightDTO) Reset() {
	*x = FlightDTO{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_booking_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FlightDTO) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FlightDTO) ProtoMessage() {}

func (x *FlightDTO) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_booking_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FlightDTO.ProtoReflect.Descriptor instead.
func (*FlightDTO) Descriptor() ([]byte, []int) {
	return file_rpc_booking_proto_rawDescGZIP(), []int{2}
}

func (x *FlightDTO) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *FlightDTO) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *FlightDTO) GetFrom() string {
	if x != nil {
		return x.From
	}
	return ""
}

func (x *FlightDTO) GetTo() string {
	if x != nil {
		return x.To
	}
	return ""
}

func (x *FlightDTO) GetDepartDate() *timestamppb.Timestamp {
	if x != nil {
		return x.DepartDate
	}
	return nil
}

func (x *FlightDTO) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *FlightDTO) GetAvailableSlot() int32 {
	if x != nil {
		return x.AvailableSlot
	}
	return 0
}

func (x *FlightDTO) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *FlightDTO) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

type Booking struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	ClientId   string                 `protobuf:"bytes,2,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty"`
	FlightId   string                 `protobuf:"bytes,3,opt,name=flight_id,json=flightId,proto3" json:"flight_id,omitempty"`
	Code       string                 `protobuf:"bytes,4,opt,name=code,proto3" json:"code,omitempty"`
	BookedSlot int32                  `protobuf:"varint,5,opt,name=booked_slot,json=bookedSlot,proto3" json:"booked_slot,omitempty"`
	Status     string                 `protobuf:"bytes,6,opt,name=status,proto3" json:"status,omitempty"`
	BookedDate *timestamppb.Timestamp `protobuf:"bytes,7,opt,name=booked_date,json=bookedDate,proto3" json:"booked_date,omitempty"`
	CreatedAt  *timestamppb.Timestamp `protobuf:"bytes,8,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt  *timestamppb.Timestamp `protobuf:"bytes,9,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	Client     *ClientDTO             `protobuf:"bytes,10,opt,name=client,proto3" json:"client,omitempty"`
	Flight     *FlightDTO             `protobuf:"bytes,11,opt,name=flight,proto3" json:"flight,omitempty"`
}

func (x *Booking) Reset() {
	*x = Booking{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_booking_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Booking) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Booking) ProtoMessage() {}

func (x *Booking) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_booking_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Booking.ProtoReflect.Descriptor instead.
func (*Booking) Descriptor() ([]byte, []int) {
	return file_rpc_booking_proto_rawDescGZIP(), []int{3}
}

func (x *Booking) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Booking) GetClientId() string {
	if x != nil {
		return x.ClientId
	}
	return ""
}

func (x *Booking) GetFlightId() string {
	if x != nil {
		return x.FlightId
	}
	return ""
}

func (x *Booking) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *Booking) GetBookedSlot() int32 {
	if x != nil {
		return x.BookedSlot
	}
	return 0
}

func (x *Booking) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *Booking) GetBookedDate() *timestamppb.Timestamp {
	if x != nil {
		return x.BookedDate
	}
	return nil
}

func (x *Booking) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *Booking) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

func (x *Booking) GetClient() *ClientDTO {
	if x != nil {
		return x.Client
	}
	return nil
}

func (x *Booking) GetFlight() *FlightDTO {
	if x != nil {
		return x.Flight
	}
	return nil
}

type SearchBookingRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	ClientId string                 `protobuf:"bytes,2,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty"`
	FlightId string                 `protobuf:"bytes,3,opt,name=flight_id,json=flightId,proto3" json:"flight_id,omitempty"`
	Code     string                 `protobuf:"bytes,4,opt,name=code,proto3" json:"code,omitempty"`
	Status   string                 `protobuf:"bytes,5,opt,name=status,proto3" json:"status,omitempty"`
	FromDate *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=from_date,json=fromDate,proto3" json:"from_date,omitempty"`
	ToDate   *timestamppb.Timestamp `protobuf:"bytes,7,opt,name=to_date,json=toDate,proto3" json:"to_date,omitempty"`
}

func (x *SearchBookingRequest) Reset() {
	*x = SearchBookingRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_booking_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchBookingRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchBookingRequest) ProtoMessage() {}

func (x *SearchBookingRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_booking_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchBookingRequest.ProtoReflect.Descriptor instead.
func (*SearchBookingRequest) Descriptor() ([]byte, []int) {
	return file_rpc_booking_proto_rawDescGZIP(), []int{4}
}

func (x *SearchBookingRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *SearchBookingRequest) GetClientId() string {
	if x != nil {
		return x.ClientId
	}
	return ""
}

func (x *SearchBookingRequest) GetFlightId() string {
	if x != nil {
		return x.FlightId
	}
	return ""
}

func (x *SearchBookingRequest) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *SearchBookingRequest) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *SearchBookingRequest) GetFromDate() *timestamppb.Timestamp {
	if x != nil {
		return x.FromDate
	}
	return nil
}

func (x *SearchBookingRequest) GetToDate() *timestamppb.Timestamp {
	if x != nil {
		return x.ToDate
	}
	return nil
}

type SearchBookingResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Booking []*Booking `protobuf:"bytes,1,rep,name=booking,proto3" json:"booking,omitempty"`
}

func (x *SearchBookingResponse) Reset() {
	*x = SearchBookingResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_booking_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchBookingResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchBookingResponse) ProtoMessage() {}

func (x *SearchBookingResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_booking_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchBookingResponse.ProtoReflect.Descriptor instead.
func (*SearchBookingResponse) Descriptor() ([]byte, []int) {
	return file_rpc_booking_proto_rawDescGZIP(), []int{5}
}

func (x *SearchBookingResponse) GetBooking() []*Booking {
	if x != nil {
		return x.Booking
	}
	return nil
}

var File_rpc_booking_proto protoreflect.FileDescriptor

var file_rpc_booking_proto_rawDesc = []byte{
	0x0a, 0x11, 0x72, 0x70, 0x63, 0x5f, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x66, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x5f, 0x62, 0x6f, 0x6f, 0x6b,
	0x69, 0x6e, 0x67, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0b, 0x61, 0x75, 0x64, 0x69, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x20, 0x0a, 0x0e, 0x42, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x50, 0x61, 0x72, 0x61,
	0x6d, 0x49, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x64, 0x22, 0xe7, 0x02, 0x0a, 0x09, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x44, 0x54,
	0x4f, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x12, 0x0a, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x04, 0x72, 0x6f, 0x6c, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61,
	0x69, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12,
	0x21, 0x0a, 0x0c, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62,
	0x65, 0x72, 0x12, 0x20, 0x0a, 0x0c, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x6f, 0x66, 0x5f, 0x62, 0x69,
	0x74, 0x68, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x64, 0x61, 0x74, 0x65, 0x4f, 0x66,
	0x42, 0x69, 0x74, 0x68, 0x12, 0x23, 0x0a, 0x0d, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79,
	0x5f, 0x63, 0x61, 0x72, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x69, 0x64, 0x65,
	0x6e, 0x74, 0x69, 0x74, 0x79, 0x43, 0x61, 0x72, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x12, 0x27, 0x0a, 0x0f, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x68, 0x69,
	0x70, 0x5f, 0x63, 0x61, 0x72, 0x64, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x6d, 0x65,
	0x6d, 0x62, 0x65, 0x72, 0x73, 0x68, 0x69, 0x70, 0x43, 0x61, 0x72, 0x64, 0x12, 0x1a, 0x0a, 0x08,
	0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x12, 0x2b, 0x0a, 0x05, 0x61, 0x75, 0x64, 0x69, 0x74, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x15, 0x2e, 0x66, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x5f, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67,
	0x2e, 0x41, 0x75, 0x64, 0x69, 0x74, 0x52, 0x05, 0x61, 0x75, 0x64, 0x69, 0x74, 0x22, 0xc5, 0x02,
	0x0a, 0x09, 0x46, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x44, 0x54, 0x4f, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x12, 0x0a, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x66,
	0x72, 0x6f, 0x6d, 0x12, 0x0e, 0x0a, 0x02, 0x74, 0x6f, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x74, 0x6f, 0x12, 0x3b, 0x0a, 0x0b, 0x64, 0x65, 0x70, 0x61, 0x72, 0x74, 0x5f, 0x64, 0x61,
	0x74, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x52, 0x0a, 0x64, 0x65, 0x70, 0x61, 0x72, 0x74, 0x44, 0x61, 0x74, 0x65,
	0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x25, 0x0a, 0x0e, 0x61, 0x76, 0x61, 0x69,
	0x6c, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x73, 0x6c, 0x6f, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x0d, 0x61, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x53, 0x6c, 0x6f, 0x74, 0x12,
	0x39, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x08, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52,
	0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x39, 0x0a, 0x0a, 0x75, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0xb9, 0x03, 0x0a, 0x07, 0x42, 0x6f, 0x6f, 0x6b, 0x69, 0x6e,
	0x67, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x1b,
	0x0a, 0x09, 0x66, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x66, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x63,
	0x6f, 0x64, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12,
	0x1f, 0x0a, 0x0b, 0x62, 0x6f, 0x6f, 0x6b, 0x65, 0x64, 0x5f, 0x73, 0x6c, 0x6f, 0x74, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x62, 0x6f, 0x6f, 0x6b, 0x65, 0x64, 0x53, 0x6c, 0x6f, 0x74,
	0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x3b, 0x0a, 0x0b, 0x62, 0x6f, 0x6f, 0x6b,
	0x65, 0x64, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0a, 0x62, 0x6f, 0x6f, 0x6b, 0x65,
	0x64, 0x44, 0x61, 0x74, 0x65, 0x12, 0x39, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64,
	0x5f, 0x61, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74,
	0x12, 0x39, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x09,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x31, 0x0a, 0x06, 0x63,
	0x6c, 0x69, 0x65, 0x6e, 0x74, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x66, 0x6c,
	0x69, 0x67, 0x68, 0x74, 0x5f, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x43, 0x6c, 0x69,
	0x65, 0x6e, 0x74, 0x44, 0x54, 0x4f, 0x52, 0x06, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x12, 0x31,
	0x0a, 0x06, 0x66, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19,
	0x2e, 0x66, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x5f, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x2e,
	0x46, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x44, 0x54, 0x4f, 0x52, 0x06, 0x66, 0x6c, 0x69, 0x67, 0x68,
	0x74, 0x22, 0xfa, 0x01, 0x0a, 0x14, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x42, 0x6f, 0x6f, 0x6b,
	0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x6c,
	0x69, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63,
	0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x66, 0x6c, 0x69, 0x67, 0x68,
	0x74, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x6c, 0x69, 0x67,
	0x68, 0x74, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x12, 0x37, 0x0a, 0x09, 0x66, 0x72, 0x6f, 0x6d, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52,
	0x08, 0x66, 0x72, 0x6f, 0x6d, 0x44, 0x61, 0x74, 0x65, 0x12, 0x33, 0x0a, 0x07, 0x74, 0x6f, 0x5f,
	0x64, 0x61, 0x74, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x06, 0x74, 0x6f, 0x44, 0x61, 0x74, 0x65, 0x22, 0x4a,
	0x0a, 0x15, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x42, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x31, 0x0a, 0x07, 0x62, 0x6f, 0x6f, 0x6b, 0x69,
	0x6e, 0x67, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x66, 0x6c, 0x69, 0x67, 0x68,
	0x74, 0x5f, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x42, 0x6f, 0x6f, 0x6b, 0x69, 0x6e,
	0x67, 0x52, 0x07, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x32, 0xb5, 0x02, 0x0a, 0x0a, 0x52,
	0x50, 0x43, 0x42, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x12, 0x43, 0x0a, 0x08, 0x46, 0x69, 0x6e,
	0x64, 0x42, 0x79, 0x49, 0x64, 0x12, 0x1e, 0x2e, 0x66, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x5f, 0x62,
	0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x42, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x50, 0x61,
	0x72, 0x61, 0x6d, 0x49, 0x64, 0x1a, 0x17, 0x2e, 0x66, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x5f, 0x62,
	0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x42, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x12, 0x41,
	0x0a, 0x0d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x12,
	0x17, 0x2e, 0x66, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x5f, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67,
	0x2e, 0x42, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x1a, 0x17, 0x2e, 0x66, 0x6c, 0x69, 0x67, 0x68,
	0x74, 0x5f, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x42, 0x6f, 0x6f, 0x6b, 0x69, 0x6e,
	0x67, 0x12, 0x41, 0x0a, 0x0d, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x42, 0x6f, 0x6f, 0x6b, 0x69,
	0x6e, 0x67, 0x12, 0x17, 0x2e, 0x66, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x5f, 0x62, 0x6f, 0x6f, 0x6b,
	0x69, 0x6e, 0x67, 0x2e, 0x42, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x1a, 0x17, 0x2e, 0x66, 0x6c,
	0x69, 0x67, 0x68, 0x74, 0x5f, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x42, 0x6f, 0x6f,
	0x6b, 0x69, 0x6e, 0x67, 0x12, 0x5c, 0x0a, 0x0d, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x42, 0x6f,
	0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x12, 0x24, 0x2e, 0x66, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x5f, 0x62,
	0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x42, 0x6f, 0x6f,
	0x6b, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x25, 0x2e, 0x66, 0x6c,
	0x69, 0x67, 0x68, 0x74, 0x5f, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x53, 0x65, 0x61,
	0x72, 0x63, 0x68, 0x42, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x42, 0x07, 0x5a, 0x05, 0x2e, 0x2f, 0x3b, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_rpc_booking_proto_rawDescOnce sync.Once
	file_rpc_booking_proto_rawDescData = file_rpc_booking_proto_rawDesc
)

func file_rpc_booking_proto_rawDescGZIP() []byte {
	file_rpc_booking_proto_rawDescOnce.Do(func() {
		file_rpc_booking_proto_rawDescData = protoimpl.X.CompressGZIP(file_rpc_booking_proto_rawDescData)
	})
	return file_rpc_booking_proto_rawDescData
}

var file_rpc_booking_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_rpc_booking_proto_goTypes = []interface{}{
	(*BookingParamId)(nil),        // 0: flight_booking.BookingParamId
	(*ClientDTO)(nil),             // 1: flight_booking.ClientDTO
	(*FlightDTO)(nil),             // 2: flight_booking.FlightDTO
	(*Booking)(nil),               // 3: flight_booking.Booking
	(*SearchBookingRequest)(nil),  // 4: flight_booking.SearchBookingRequest
	(*SearchBookingResponse)(nil), // 5: flight_booking.SearchBookingResponse
	(*Audit)(nil),                 // 6: flight_booking.Audit
	(*timestamppb.Timestamp)(nil), // 7: google.protobuf.Timestamp
}
var file_rpc_booking_proto_depIdxs = []int32{
	6,  // 0: flight_booking.ClientDTO.audit:type_name -> flight_booking.Audit
	7,  // 1: flight_booking.FlightDTO.depart_date:type_name -> google.protobuf.Timestamp
	7,  // 2: flight_booking.FlightDTO.created_at:type_name -> google.protobuf.Timestamp
	7,  // 3: flight_booking.FlightDTO.updated_at:type_name -> google.protobuf.Timestamp
	7,  // 4: flight_booking.Booking.booked_date:type_name -> google.protobuf.Timestamp
	7,  // 5: flight_booking.Booking.created_at:type_name -> google.protobuf.Timestamp
	7,  // 6: flight_booking.Booking.updated_at:type_name -> google.protobuf.Timestamp
	1,  // 7: flight_booking.Booking.client:type_name -> flight_booking.ClientDTO
	2,  // 8: flight_booking.Booking.flight:type_name -> flight_booking.FlightDTO
	7,  // 9: flight_booking.SearchBookingRequest.from_date:type_name -> google.protobuf.Timestamp
	7,  // 10: flight_booking.SearchBookingRequest.to_date:type_name -> google.protobuf.Timestamp
	3,  // 11: flight_booking.SearchBookingResponse.booking:type_name -> flight_booking.Booking
	0,  // 12: flight_booking.RPCBooking.FindById:input_type -> flight_booking.BookingParamId
	3,  // 13: flight_booking.RPCBooking.CreateBooking:input_type -> flight_booking.Booking
	3,  // 14: flight_booking.RPCBooking.UpdateBooking:input_type -> flight_booking.Booking
	4,  // 15: flight_booking.RPCBooking.SearchBooking:input_type -> flight_booking.SearchBookingRequest
	3,  // 16: flight_booking.RPCBooking.FindById:output_type -> flight_booking.Booking
	3,  // 17: flight_booking.RPCBooking.CreateBooking:output_type -> flight_booking.Booking
	3,  // 18: flight_booking.RPCBooking.UpdateBooking:output_type -> flight_booking.Booking
	5,  // 19: flight_booking.RPCBooking.SearchBooking:output_type -> flight_booking.SearchBookingResponse
	16, // [16:20] is the sub-list for method output_type
	12, // [12:16] is the sub-list for method input_type
	12, // [12:12] is the sub-list for extension type_name
	12, // [12:12] is the sub-list for extension extendee
	0,  // [0:12] is the sub-list for field type_name
}

func init() { file_rpc_booking_proto_init() }
func file_rpc_booking_proto_init() {
	if File_rpc_booking_proto != nil {
		return
	}
	file_audit_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_rpc_booking_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BookingParamId); i {
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
		file_rpc_booking_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClientDTO); i {
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
		file_rpc_booking_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FlightDTO); i {
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
		file_rpc_booking_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Booking); i {
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
		file_rpc_booking_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchBookingRequest); i {
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
		file_rpc_booking_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchBookingResponse); i {
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
			RawDescriptor: file_rpc_booking_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_rpc_booking_proto_goTypes,
		DependencyIndexes: file_rpc_booking_proto_depIdxs,
		MessageInfos:      file_rpc_booking_proto_msgTypes,
	}.Build()
	File_rpc_booking_proto = out.File
	file_rpc_booking_proto_rawDesc = nil
	file_rpc_booking_proto_goTypes = nil
	file_rpc_booking_proto_depIdxs = nil
}
