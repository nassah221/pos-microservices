// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.19.4
// source: contract/cashier.proto

package cashier

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

type Cashier struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name     string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Email    string `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
	Password string `protobuf:"bytes,4,opt,name=password,proto3" json:"password,omitempty"`
	Created  int64  `protobuf:"varint,5,opt,name=created,proto3" json:"created,omitempty"`
	Updated  int64  `protobuf:"varint,6,opt,name=updated,proto3" json:"updated,omitempty"`
}

func (x *Cashier) Reset() {
	*x = Cashier{}
	if protoimpl.UnsafeEnabled {
		mi := &file_contract_cashier_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Cashier) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Cashier) ProtoMessage() {}

func (x *Cashier) ProtoReflect() protoreflect.Message {
	mi := &file_contract_cashier_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Cashier.ProtoReflect.Descriptor instead.
func (*Cashier) Descriptor() ([]byte, []int) {
	return file_contract_cashier_proto_rawDescGZIP(), []int{0}
}

func (x *Cashier) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Cashier) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Cashier) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *Cashier) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *Cashier) GetCreated() int64 {
	if x != nil {
		return x.Created
	}
	return 0
}

func (x *Cashier) GetUpdated() int64 {
	if x != nil {
		return x.Updated
	}
	return 0
}

type GetCashierRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetCashierRequest) Reset() {
	*x = GetCashierRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_contract_cashier_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCashierRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCashierRequest) ProtoMessage() {}

func (x *GetCashierRequest) ProtoReflect() protoreflect.Message {
	mi := &file_contract_cashier_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCashierRequest.ProtoReflect.Descriptor instead.
func (*GetCashierRequest) Descriptor() ([]byte, []int) {
	return file_contract_cashier_proto_rawDescGZIP(), []int{1}
}

func (x *GetCashierRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type ListCashiersRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ListCashiersRequest) Reset() {
	*x = ListCashiersRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_contract_cashier_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListCashiersRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListCashiersRequest) ProtoMessage() {}

func (x *ListCashiersRequest) ProtoReflect() protoreflect.Message {
	mi := &file_contract_cashier_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListCashiersRequest.ProtoReflect.Descriptor instead.
func (*ListCashiersRequest) Descriptor() ([]byte, []int) {
	return file_contract_cashier_proto_rawDescGZIP(), []int{2}
}

type ListCashiersResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Cashiers []*Cashier `protobuf:"bytes,1,rep,name=cashiers,proto3" json:"cashiers,omitempty"`
}

func (x *ListCashiersResponse) Reset() {
	*x = ListCashiersResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_contract_cashier_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListCashiersResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListCashiersResponse) ProtoMessage() {}

func (x *ListCashiersResponse) ProtoReflect() protoreflect.Message {
	mi := &file_contract_cashier_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListCashiersResponse.ProtoReflect.Descriptor instead.
func (*ListCashiersResponse) Descriptor() ([]byte, []int) {
	return file_contract_cashier_proto_rawDescGZIP(), []int{3}
}

func (x *ListCashiersResponse) GetCashiers() []*Cashier {
	if x != nil {
		return x.Cashiers
	}
	return nil
}

type SigninRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Email    string `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
}

func (x *SigninRequest) Reset() {
	*x = SigninRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_contract_cashier_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SigninRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SigninRequest) ProtoMessage() {}

func (x *SigninRequest) ProtoReflect() protoreflect.Message {
	mi := &file_contract_cashier_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SigninRequest.ProtoReflect.Descriptor instead.
func (*SigninRequest) Descriptor() ([]byte, []int) {
	return file_contract_cashier_proto_rawDescGZIP(), []int{4}
}

func (x *SigninRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *SigninRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type SigninResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Cashier *Cashier `protobuf:"bytes,1,opt,name=cashier,proto3" json:"cashier,omitempty"`
	Token   string   `protobuf:"bytes,2,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *SigninResponse) Reset() {
	*x = SigninResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_contract_cashier_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SigninResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SigninResponse) ProtoMessage() {}

func (x *SigninResponse) ProtoReflect() protoreflect.Message {
	mi := &file_contract_cashier_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SigninResponse.ProtoReflect.Descriptor instead.
func (*SigninResponse) Descriptor() ([]byte, []int) {
	return file_contract_cashier_proto_rawDescGZIP(), []int{5}
}

func (x *SigninResponse) GetCashier() *Cashier {
	if x != nil {
		return x.Cashier
	}
	return nil
}

func (x *SigninResponse) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type DeleteCashierResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteCashierResponse) Reset() {
	*x = DeleteCashierResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_contract_cashier_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteCashierResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteCashierResponse) ProtoMessage() {}

func (x *DeleteCashierResponse) ProtoReflect() protoreflect.Message {
	mi := &file_contract_cashier_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteCashierResponse.ProtoReflect.Descriptor instead.
func (*DeleteCashierResponse) Descriptor() ([]byte, []int) {
	return file_contract_cashier_proto_rawDescGZIP(), []int{6}
}

func (x *DeleteCashierResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

var File_contract_cashier_proto protoreflect.FileDescriptor

var file_contract_cashier_proto_rawDesc = []byte{
	0x0a, 0x16, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x2f, 0x63, 0x61, 0x73, 0x68, 0x69,
	0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x63, 0x61, 0x73, 0x68, 0x69, 0x65,
	0x72, 0x22, 0x93, 0x01, 0x0a, 0x07, 0x43, 0x61, 0x73, 0x68, 0x69, 0x65, 0x72, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77,
	0x6f, 0x72, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77,
	0x6f, 0x72, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x12, 0x18, 0x0a,
	0x07, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07,
	0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x22, 0x23, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x43, 0x61,
	0x73, 0x68, 0x69, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x15, 0x0a, 0x13,
	0x4c, 0x69, 0x73, 0x74, 0x43, 0x61, 0x73, 0x68, 0x69, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x22, 0x44, 0x0a, 0x14, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x61, 0x73, 0x68, 0x69,
	0x65, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2c, 0x0a, 0x08, 0x63,
	0x61, 0x73, 0x68, 0x69, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e,
	0x63, 0x61, 0x73, 0x68, 0x69, 0x65, 0x72, 0x2e, 0x43, 0x61, 0x73, 0x68, 0x69, 0x65, 0x72, 0x52,
	0x08, 0x63, 0x61, 0x73, 0x68, 0x69, 0x65, 0x72, 0x73, 0x22, 0x41, 0x0a, 0x0d, 0x53, 0x69, 0x67,
	0x6e, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d,
	0x61, 0x69, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c,
	0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x22, 0x52, 0x0a, 0x0e,
	0x53, 0x69, 0x67, 0x6e, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2a,
	0x0a, 0x07, 0x63, 0x61, 0x73, 0x68, 0x69, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x10, 0x2e, 0x63, 0x61, 0x73, 0x68, 0x69, 0x65, 0x72, 0x2e, 0x43, 0x61, 0x73, 0x68, 0x69, 0x65,
	0x72, 0x52, 0x07, 0x63, 0x61, 0x73, 0x68, 0x69, 0x65, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f,
	0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e,
	0x22, 0x27, 0x0a, 0x15, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x61, 0x73, 0x68, 0x69, 0x65,
	0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x32, 0x8e, 0x03, 0x0a, 0x0e, 0x43, 0x61,
	0x73, 0x68, 0x69, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x36, 0x0a, 0x06,
	0x53, 0x69, 0x67, 0x6e, 0x75, 0x70, 0x12, 0x10, 0x2e, 0x63, 0x61, 0x73, 0x68, 0x69, 0x65, 0x72,
	0x2e, 0x43, 0x61, 0x73, 0x68, 0x69, 0x65, 0x72, 0x1a, 0x1a, 0x2e, 0x63, 0x61, 0x73, 0x68, 0x69,
	0x65, 0x72, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x61, 0x73, 0x68, 0x69, 0x65, 0x72, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x39, 0x0a, 0x06, 0x53, 0x69, 0x67, 0x6e, 0x69, 0x6e, 0x12, 0x16,
	0x2e, 0x63, 0x61, 0x73, 0x68, 0x69, 0x65, 0x72, 0x2e, 0x53, 0x69, 0x67, 0x6e, 0x69, 0x6e, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x63, 0x61, 0x73, 0x68, 0x69, 0x65, 0x72,
	0x2e, 0x53, 0x69, 0x67, 0x6e, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x3a, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x43, 0x61, 0x73, 0x68, 0x69, 0x65, 0x72, 0x12, 0x1a, 0x2e,
	0x63, 0x61, 0x73, 0x68, 0x69, 0x65, 0x72, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x61, 0x73, 0x68, 0x69,
	0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x10, 0x2e, 0x63, 0x61, 0x73, 0x68,
	0x69, 0x65, 0x72, 0x2e, 0x43, 0x61, 0x73, 0x68, 0x69, 0x65, 0x72, 0x12, 0x4b, 0x0a, 0x0c, 0x4c,
	0x69, 0x73, 0x74, 0x43, 0x61, 0x73, 0x68, 0x69, 0x65, 0x72, 0x73, 0x12, 0x1c, 0x2e, 0x63, 0x61,
	0x73, 0x68, 0x69, 0x65, 0x72, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x61, 0x73, 0x68, 0x69, 0x65,
	0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x63, 0x61, 0x73, 0x68,
	0x69, 0x65, 0x72, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x61, 0x73, 0x68, 0x69, 0x65, 0x72, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x33, 0x0a, 0x0d, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x43, 0x61, 0x73, 0x68, 0x69, 0x65, 0x72, 0x12, 0x10, 0x2e, 0x63, 0x61, 0x73, 0x68,
	0x69, 0x65, 0x72, 0x2e, 0x43, 0x61, 0x73, 0x68, 0x69, 0x65, 0x72, 0x1a, 0x10, 0x2e, 0x63, 0x61,
	0x73, 0x68, 0x69, 0x65, 0x72, 0x2e, 0x43, 0x61, 0x73, 0x68, 0x69, 0x65, 0x72, 0x12, 0x4b, 0x0a,
	0x0d, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x61, 0x73, 0x68, 0x69, 0x65, 0x72, 0x12, 0x1a,
	0x2e, 0x63, 0x61, 0x73, 0x68, 0x69, 0x65, 0x72, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x61, 0x73, 0x68,
	0x69, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x63, 0x61, 0x73,
	0x68, 0x69, 0x65, 0x72, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x61, 0x73, 0x68, 0x69,
	0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x0a, 0x5a, 0x08, 0x2f, 0x63,
	0x61, 0x73, 0x68, 0x69, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_contract_cashier_proto_rawDescOnce sync.Once
	file_contract_cashier_proto_rawDescData = file_contract_cashier_proto_rawDesc
)

func file_contract_cashier_proto_rawDescGZIP() []byte {
	file_contract_cashier_proto_rawDescOnce.Do(func() {
		file_contract_cashier_proto_rawDescData = protoimpl.X.CompressGZIP(file_contract_cashier_proto_rawDescData)
	})
	return file_contract_cashier_proto_rawDescData
}

var file_contract_cashier_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_contract_cashier_proto_goTypes = []interface{}{
	(*Cashier)(nil),               // 0: cashier.Cashier
	(*GetCashierRequest)(nil),     // 1: cashier.GetCashierRequest
	(*ListCashiersRequest)(nil),   // 2: cashier.ListCashiersRequest
	(*ListCashiersResponse)(nil),  // 3: cashier.ListCashiersResponse
	(*SigninRequest)(nil),         // 4: cashier.SigninRequest
	(*SigninResponse)(nil),        // 5: cashier.SigninResponse
	(*DeleteCashierResponse)(nil), // 6: cashier.DeleteCashierResponse
}
var file_contract_cashier_proto_depIdxs = []int32{
	0, // 0: cashier.ListCashiersResponse.cashiers:type_name -> cashier.Cashier
	0, // 1: cashier.SigninResponse.cashier:type_name -> cashier.Cashier
	0, // 2: cashier.CashierService.Signup:input_type -> cashier.Cashier
	4, // 3: cashier.CashierService.Signin:input_type -> cashier.SigninRequest
	1, // 4: cashier.CashierService.GetCashier:input_type -> cashier.GetCashierRequest
	2, // 5: cashier.CashierService.ListCashiers:input_type -> cashier.ListCashiersRequest
	0, // 6: cashier.CashierService.UpdateCashier:input_type -> cashier.Cashier
	1, // 7: cashier.CashierService.DeleteCashier:input_type -> cashier.GetCashierRequest
	1, // 8: cashier.CashierService.Signup:output_type -> cashier.GetCashierRequest
	5, // 9: cashier.CashierService.Signin:output_type -> cashier.SigninResponse
	0, // 10: cashier.CashierService.GetCashier:output_type -> cashier.Cashier
	3, // 11: cashier.CashierService.ListCashiers:output_type -> cashier.ListCashiersResponse
	0, // 12: cashier.CashierService.UpdateCashier:output_type -> cashier.Cashier
	6, // 13: cashier.CashierService.DeleteCashier:output_type -> cashier.DeleteCashierResponse
	8, // [8:14] is the sub-list for method output_type
	2, // [2:8] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_contract_cashier_proto_init() }
func file_contract_cashier_proto_init() {
	if File_contract_cashier_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_contract_cashier_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Cashier); i {
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
		file_contract_cashier_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCashierRequest); i {
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
		file_contract_cashier_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListCashiersRequest); i {
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
		file_contract_cashier_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListCashiersResponse); i {
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
		file_contract_cashier_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SigninRequest); i {
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
		file_contract_cashier_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SigninResponse); i {
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
		file_contract_cashier_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteCashierResponse); i {
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
			RawDescriptor: file_contract_cashier_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_contract_cashier_proto_goTypes,
		DependencyIndexes: file_contract_cashier_proto_depIdxs,
		MessageInfos:      file_contract_cashier_proto_msgTypes,
	}.Build()
	File_contract_cashier_proto = out.File
	file_contract_cashier_proto_rawDesc = nil
	file_contract_cashier_proto_goTypes = nil
	file_contract_cashier_proto_depIdxs = nil
}
