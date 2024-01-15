// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        (unknown)
// source: customer/v1/customer.proto

package pb

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

// VerificationResult represents the possible values after attempting to verify a customer's phone number.
type VerificationResult int32

const (
	// Default value.
	VerificationResult_VERIFICATION_RESULT_UNSPECIFIED VerificationResult = 0
	// The customer has been sent the server generated code but has not yet submitted his code.
	VerificationResult_VERIFICATION_RESULT_IN_PROGRESS VerificationResult = 1
	// The code which the customer sent has expired.
	VerificationResult_VERIFICATION_RESULT_CODE_EXPIRED VerificationResult = 2
	// The code which the customer sent does not match the server generated code.
	VerificationResult_VERIFICATION_RESULT_INCORRECT_CODE VerificationResult = 3
	// The customer has reached the maximum number of attempts to verify the phone number and can not perform any further attempts.
	VerificationResult_VERIFICATION_RESULT_MAX_ATTEMPTS_REACHED VerificationResult = 4
	// The code which the customer sent matches the verification code and the customer's phone is now verified.
	VerificationResult_VERIFICATION_RESULT_SUCCESS VerificationResult = 5
)

// Enum value maps for VerificationResult.
var (
	VerificationResult_name = map[int32]string{
		0: "VERIFICATION_RESULT_UNSPECIFIED",
		1: "VERIFICATION_RESULT_IN_PROGRESS",
		2: "VERIFICATION_RESULT_CODE_EXPIRED",
		3: "VERIFICATION_RESULT_INCORRECT_CODE",
		4: "VERIFICATION_RESULT_MAX_ATTEMPTS_REACHED",
		5: "VERIFICATION_RESULT_SUCCESS",
	}
	VerificationResult_value = map[string]int32{
		"VERIFICATION_RESULT_UNSPECIFIED":          0,
		"VERIFICATION_RESULT_IN_PROGRESS":          1,
		"VERIFICATION_RESULT_CODE_EXPIRED":         2,
		"VERIFICATION_RESULT_INCORRECT_CODE":       3,
		"VERIFICATION_RESULT_MAX_ATTEMPTS_REACHED": 4,
		"VERIFICATION_RESULT_SUCCESS":              5,
	}
)

func (x VerificationResult) Enum() *VerificationResult {
	p := new(VerificationResult)
	*p = x
	return p
}

func (x VerificationResult) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (VerificationResult) Descriptor() protoreflect.EnumDescriptor {
	return file_customer_v1_customer_proto_enumTypes[0].Descriptor()
}

func (VerificationResult) Type() protoreflect.EnumType {
	return &file_customer_v1_customer_proto_enumTypes[0]
}

func (x VerificationResult) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use VerificationResult.Descriptor instead.
func (VerificationResult) EnumDescriptor() ([]byte, []int) {
	return file_customer_v1_customer_proto_rawDescGZIP(), []int{0}
}

// Customer represents a customer in the coffee shop system.
type Customer struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// A unique id for the customer.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// The customer's first name.
	FirstName string `protobuf:"bytes,2,opt,name=first_name,json=firstName,proto3" json:"first_name,omitempty"`
	// The customer's last name.
	LastName string `protobuf:"bytes,3,opt,name=last_name,json=lastName,proto3" json:"last_name,omitempty"`
	// The customer's email address.
	Email string `protobuf:"bytes,4,opt,name=email,proto3" json:"email,omitempty"`
	// The customer's phone number.
	PhoneNumber string `protobuf:"bytes,5,opt,name=phone_number,json=phoneNumber,proto3" json:"phone_number,omitempty"`
	// Whether the customers's phone number has been verified.
	PhoneVerified bool `protobuf:"varint,6,opt,name=phone_verified,json=phoneVerified,proto3" json:"phone_verified,omitempty"`
}

func (x *Customer) Reset() {
	*x = Customer{}
	if protoimpl.UnsafeEnabled {
		mi := &file_customer_v1_customer_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Customer) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Customer) ProtoMessage() {}

func (x *Customer) ProtoReflect() protoreflect.Message {
	mi := &file_customer_v1_customer_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Customer.ProtoReflect.Descriptor instead.
func (*Customer) Descriptor() ([]byte, []int) {
	return file_customer_v1_customer_proto_rawDescGZIP(), []int{0}
}

func (x *Customer) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Customer) GetFirstName() string {
	if x != nil {
		return x.FirstName
	}
	return ""
}

func (x *Customer) GetLastName() string {
	if x != nil {
		return x.LastName
	}
	return ""
}

func (x *Customer) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *Customer) GetPhoneNumber() string {
	if x != nil {
		return x.PhoneNumber
	}
	return ""
}

func (x *Customer) GetPhoneVerified() bool {
	if x != nil {
		return x.PhoneVerified
	}
	return false
}

// CreateCustomerRequest is the request issued to create a customer in the system.
// This method is idempotent.
type CreateCustomerRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// An idempotent request id to uniqely identify the create request.
	// This field is optional - if no request_id is provided, one will be automatically generated.
	RequestId string `protobuf:"bytes,1,opt,name=request_id,json=requestId,proto3" json:"request_id,omitempty"`
	// The customer to be created.
	Customer *Customer `protobuf:"bytes,2,opt,name=customer,proto3" json:"customer,omitempty"`
}

func (x *CreateCustomerRequest) Reset() {
	*x = CreateCustomerRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_customer_v1_customer_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateCustomerRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateCustomerRequest) ProtoMessage() {}

func (x *CreateCustomerRequest) ProtoReflect() protoreflect.Message {
	mi := &file_customer_v1_customer_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateCustomerRequest.ProtoReflect.Descriptor instead.
func (*CreateCustomerRequest) Descriptor() ([]byte, []int) {
	return file_customer_v1_customer_proto_rawDescGZIP(), []int{1}
}

func (x *CreateCustomerRequest) GetRequestId() string {
	if x != nil {
		return x.RequestId
	}
	return ""
}

func (x *CreateCustomerRequest) GetCustomer() *Customer {
	if x != nil {
		return x.Customer
	}
	return nil
}

// CreateCustomerResponse is the response returned after a customer has been created. It returns the newly created customer.
type CreateCustomerResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The created customer.
	Customer *Customer `protobuf:"bytes,1,opt,name=customer,proto3" json:"customer,omitempty"`
}

func (x *CreateCustomerResponse) Reset() {
	*x = CreateCustomerResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_customer_v1_customer_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateCustomerResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateCustomerResponse) ProtoMessage() {}

func (x *CreateCustomerResponse) ProtoReflect() protoreflect.Message {
	mi := &file_customer_v1_customer_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateCustomerResponse.ProtoReflect.Descriptor instead.
func (*CreateCustomerResponse) Descriptor() ([]byte, []int) {
	return file_customer_v1_customer_proto_rawDescGZIP(), []int{2}
}

func (x *CreateCustomerResponse) GetCustomer() *Customer {
	if x != nil {
		return x.Customer
	}
	return nil
}

// VerifyCustomerRequest performs verification on the customers phone number by comparing
type VerifyCustomerRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The id of the customer to verify.
	CustomerId string `protobuf:"bytes,2,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`
	// The user provided code for the verification process.
	VerificationCode string `protobuf:"bytes,3,opt,name=verification_code,json=verificationCode,proto3" json:"verification_code,omitempty"`
}

func (x *VerifyCustomerRequest) Reset() {
	*x = VerifyCustomerRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_customer_v1_customer_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VerifyCustomerRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VerifyCustomerRequest) ProtoMessage() {}

func (x *VerifyCustomerRequest) ProtoReflect() protoreflect.Message {
	mi := &file_customer_v1_customer_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VerifyCustomerRequest.ProtoReflect.Descriptor instead.
func (*VerifyCustomerRequest) Descriptor() ([]byte, []int) {
	return file_customer_v1_customer_proto_rawDescGZIP(), []int{3}
}

func (x *VerifyCustomerRequest) GetCustomerId() string {
	if x != nil {
		return x.CustomerId
	}
	return ""
}

func (x *VerifyCustomerRequest) GetVerificationCode() string {
	if x != nil {
		return x.VerificationCode
	}
	return ""
}

// VerifyCustomerResponse contains the result of the verification attempt made by a given customer.
type VerifyCustomerResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The id of the customer for which this verification result is attributed.
	CustomerId string `protobuf:"bytes,1,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`
	// The result of the verification.
	Result VerificationResult `protobuf:"varint,2,opt,name=result,proto3,enum=customer.v1.VerificationResult" json:"result,omitempty"`
}

func (x *VerifyCustomerResponse) Reset() {
	*x = VerifyCustomerResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_customer_v1_customer_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VerifyCustomerResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VerifyCustomerResponse) ProtoMessage() {}

func (x *VerifyCustomerResponse) ProtoReflect() protoreflect.Message {
	mi := &file_customer_v1_customer_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VerifyCustomerResponse.ProtoReflect.Descriptor instead.
func (*VerifyCustomerResponse) Descriptor() ([]byte, []int) {
	return file_customer_v1_customer_proto_rawDescGZIP(), []int{4}
}

func (x *VerifyCustomerResponse) GetCustomerId() string {
	if x != nil {
		return x.CustomerId
	}
	return ""
}

func (x *VerifyCustomerResponse) GetResult() VerificationResult {
	if x != nil {
		return x.Result
	}
	return VerificationResult_VERIFICATION_RESULT_UNSPECIFIED
}

// GetCustomerRequest is the request issued to get a single customer by their id.
type GetCustomerRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The id of the customer to retrieve.
	CustomerId string `protobuf:"bytes,1,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`
}

func (x *GetCustomerRequest) Reset() {
	*x = GetCustomerRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_customer_v1_customer_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCustomerRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCustomerRequest) ProtoMessage() {}

func (x *GetCustomerRequest) ProtoReflect() protoreflect.Message {
	mi := &file_customer_v1_customer_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCustomerRequest.ProtoReflect.Descriptor instead.
func (*GetCustomerRequest) Descriptor() ([]byte, []int) {
	return file_customer_v1_customer_proto_rawDescGZIP(), []int{5}
}

func (x *GetCustomerRequest) GetCustomerId() string {
	if x != nil {
		return x.CustomerId
	}
	return ""
}

// GetCustomerResponse is the response returned when GetCustomer is called.
type GetCustomerResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The retrieved customer.
	Customer *Customer `protobuf:"bytes,1,opt,name=customer,proto3" json:"customer,omitempty"`
}

func (x *GetCustomerResponse) Reset() {
	*x = GetCustomerResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_customer_v1_customer_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCustomerResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCustomerResponse) ProtoMessage() {}

func (x *GetCustomerResponse) ProtoReflect() protoreflect.Message {
	mi := &file_customer_v1_customer_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCustomerResponse.ProtoReflect.Descriptor instead.
func (*GetCustomerResponse) Descriptor() ([]byte, []int) {
	return file_customer_v1_customer_proto_rawDescGZIP(), []int{6}
}

func (x *GetCustomerResponse) GetCustomer() *Customer {
	if x != nil {
		return x.Customer
	}
	return nil
}

// BatchGetCustomersRequest is the request issued to get a batch of customers by their ids.
type BatchGetCustomersRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The id of the customer to retrieve.
	CustomerIds []string `protobuf:"bytes,1,rep,name=customer_ids,json=customerIds,proto3" json:"customer_ids,omitempty"`
	// The number of customers to return per page.
	PageSize uint32 `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	// The index from which to start fetching customers.
	PageIndex uint32 `protobuf:"varint,3,opt,name=page_index,json=pageIndex,proto3" json:"page_index,omitempty"`
}

func (x *BatchGetCustomersRequest) Reset() {
	*x = BatchGetCustomersRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_customer_v1_customer_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BatchGetCustomersRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BatchGetCustomersRequest) ProtoMessage() {}

func (x *BatchGetCustomersRequest) ProtoReflect() protoreflect.Message {
	mi := &file_customer_v1_customer_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BatchGetCustomersRequest.ProtoReflect.Descriptor instead.
func (*BatchGetCustomersRequest) Descriptor() ([]byte, []int) {
	return file_customer_v1_customer_proto_rawDescGZIP(), []int{7}
}

func (x *BatchGetCustomersRequest) GetCustomerIds() []string {
	if x != nil {
		return x.CustomerIds
	}
	return nil
}

func (x *BatchGetCustomersRequest) GetPageSize() uint32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *BatchGetCustomersRequest) GetPageIndex() uint32 {
	if x != nil {
		return x.PageIndex
	}
	return 0
}

// BatchGetCustomersResponse is the response returned when BatchGetCustomers is called.
// It is a paginated response.
type BatchGetCustomersResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The list of retrieved customers.
	Customers []*Customer `protobuf:"bytes,1,rep,name=customers,proto3" json:"customers,omitempty"`
	// This field is true if there are more pages of results and fasle otherwise.
	Next bool `protobuf:"varint,2,opt,name=next,proto3" json:"next,omitempty"`
}

func (x *BatchGetCustomersResponse) Reset() {
	*x = BatchGetCustomersResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_customer_v1_customer_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BatchGetCustomersResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BatchGetCustomersResponse) ProtoMessage() {}

func (x *BatchGetCustomersResponse) ProtoReflect() protoreflect.Message {
	mi := &file_customer_v1_customer_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BatchGetCustomersResponse.ProtoReflect.Descriptor instead.
func (*BatchGetCustomersResponse) Descriptor() ([]byte, []int) {
	return file_customer_v1_customer_proto_rawDescGZIP(), []int{8}
}

func (x *BatchGetCustomersResponse) GetCustomers() []*Customer {
	if x != nil {
		return x.Customers
	}
	return nil
}

func (x *BatchGetCustomersResponse) GetNext() bool {
	if x != nil {
		return x.Next
	}
	return false
}

var File_customer_v1_customer_proto protoreflect.FileDescriptor

var file_customer_v1_customer_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x75,
	0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x63, 0x75,
	0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb6, 0x01, 0x0a, 0x08, 0x43, 0x75, 0x73, 0x74,
	0x6f, 0x6d, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x66, 0x69, 0x72, 0x73, 0x74, 0x5f, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x66, 0x69, 0x72, 0x73, 0x74, 0x4e,
	0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x61, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65,
	0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x21, 0x0a, 0x0c, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x5f,
	0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x68,
	0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x25, 0x0a, 0x0e, 0x70, 0x68, 0x6f,
	0x6e, 0x65, 0x5f, 0x76, 0x65, 0x72, 0x69, 0x66, 0x69, 0x65, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x0d, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x56, 0x65, 0x72, 0x69, 0x66, 0x69, 0x65, 0x64,
	0x22, 0x69, 0x0a, 0x15, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d,
	0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x72, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x72,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x64, 0x12, 0x31, 0x0a, 0x08, 0x63, 0x75, 0x73, 0x74,
	0x6f, 0x6d, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x63, 0x75, 0x73,
	0x74, 0x6f, 0x6d, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65,
	0x72, 0x52, 0x08, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x22, 0x4b, 0x0a, 0x16, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x31, 0x0a, 0x08, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65,
	0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d,
	0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x52, 0x08,
	0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x22, 0x65, 0x0a, 0x15, 0x56, 0x65, 0x72, 0x69,
	0x66, 0x79, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5f, 0x69, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72,
	0x49, 0x64, 0x12, 0x2b, 0x0a, 0x11, 0x76, 0x65, 0x72, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x76,
	0x65, 0x72, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x64, 0x65, 0x22,
	0x72, 0x0a, 0x16, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65,
	0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x75, 0x73,
	0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a,
	0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x49, 0x64, 0x12, 0x37, 0x0a, 0x06, 0x72, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1f, 0x2e, 0x63, 0x75, 0x73,
	0x74, 0x6f, 0x6d, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x56, 0x65, 0x72, 0x69, 0x66, 0x69, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x52, 0x06, 0x72, 0x65, 0x73,
	0x75, 0x6c, 0x74, 0x22, 0x35, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d,
	0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x75, 0x73,
	0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a,
	0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x49, 0x64, 0x22, 0x48, 0x0a, 0x13, 0x47, 0x65,
	0x74, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x31, 0x0a, 0x08, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x2e, 0x76,
	0x31, 0x2e, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x52, 0x08, 0x63, 0x75, 0x73, 0x74,
	0x6f, 0x6d, 0x65, 0x72, 0x22, 0x79, 0x0a, 0x18, 0x42, 0x61, 0x74, 0x63, 0x68, 0x47, 0x65, 0x74,
	0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x21, 0x0a, 0x0c, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72,
	0x49, 0x64, 0x73, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65,
	0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x70, 0x61, 0x67, 0x65, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x22,
	0x64, 0x0a, 0x19, 0x42, 0x61, 0x74, 0x63, 0x68, 0x47, 0x65, 0x74, 0x43, 0x75, 0x73, 0x74, 0x6f,
	0x6d, 0x65, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x33, 0x0a, 0x09,
	0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x15, 0x2e, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x75,
	0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x52, 0x09, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72,
	0x73, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x65, 0x78, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x04, 0x6e, 0x65, 0x78, 0x74, 0x2a, 0xfb, 0x01, 0x0a, 0x12, 0x56, 0x65, 0x72, 0x69, 0x66, 0x69,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x23, 0x0a, 0x1f,
	0x56, 0x45, 0x52, 0x49, 0x46, 0x49, 0x43, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x52, 0x45, 0x53,
	0x55, 0x4c, 0x54, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10,
	0x00, 0x12, 0x23, 0x0a, 0x1f, 0x56, 0x45, 0x52, 0x49, 0x46, 0x49, 0x43, 0x41, 0x54, 0x49, 0x4f,
	0x4e, 0x5f, 0x52, 0x45, 0x53, 0x55, 0x4c, 0x54, 0x5f, 0x49, 0x4e, 0x5f, 0x50, 0x52, 0x4f, 0x47,
	0x52, 0x45, 0x53, 0x53, 0x10, 0x01, 0x12, 0x24, 0x0a, 0x20, 0x56, 0x45, 0x52, 0x49, 0x46, 0x49,
	0x43, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x52, 0x45, 0x53, 0x55, 0x4c, 0x54, 0x5f, 0x43, 0x4f,
	0x44, 0x45, 0x5f, 0x45, 0x58, 0x50, 0x49, 0x52, 0x45, 0x44, 0x10, 0x02, 0x12, 0x26, 0x0a, 0x22,
	0x56, 0x45, 0x52, 0x49, 0x46, 0x49, 0x43, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x52, 0x45, 0x53,
	0x55, 0x4c, 0x54, 0x5f, 0x49, 0x4e, 0x43, 0x4f, 0x52, 0x52, 0x45, 0x43, 0x54, 0x5f, 0x43, 0x4f,
	0x44, 0x45, 0x10, 0x03, 0x12, 0x2c, 0x0a, 0x28, 0x56, 0x45, 0x52, 0x49, 0x46, 0x49, 0x43, 0x41,
	0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x52, 0x45, 0x53, 0x55, 0x4c, 0x54, 0x5f, 0x4d, 0x41, 0x58, 0x5f,
	0x41, 0x54, 0x54, 0x45, 0x4d, 0x50, 0x54, 0x53, 0x5f, 0x52, 0x45, 0x41, 0x43, 0x48, 0x45, 0x44,
	0x10, 0x04, 0x12, 0x1f, 0x0a, 0x1b, 0x56, 0x45, 0x52, 0x49, 0x46, 0x49, 0x43, 0x41, 0x54, 0x49,
	0x4f, 0x4e, 0x5f, 0x52, 0x45, 0x53, 0x55, 0x4c, 0x54, 0x5f, 0x53, 0x55, 0x43, 0x43, 0x45, 0x53,
	0x53, 0x10, 0x05, 0x32, 0x8d, 0x04, 0x0a, 0x0f, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x73, 0x0a, 0x0e, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x12, 0x22, 0x2e, 0x63, 0x75, 0x73, 0x74,
	0x6f, 0x6d, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x75,
	0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x23, 0x2e,
	0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x18, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x12, 0x3a, 0x01, 0x2a, 0x22, 0x0d, 0x2f,
	0x76, 0x31, 0x2f, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x73, 0x12, 0x88, 0x01, 0x0a,
	0x0e, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x12,
	0x22, 0x2e, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x56, 0x65,
	0x72, 0x69, 0x66, 0x79, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x23, 0x2e, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x2e, 0x76,
	0x31, 0x2e, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x2d, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x27,
	0x3a, 0x01, 0x2a, 0x22, 0x22, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65,
	0x72, 0x73, 0x2f, 0x7b, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x7d,
	0x3a, 0x76, 0x65, 0x72, 0x69, 0x66, 0x79, 0x12, 0x75, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x43, 0x75,
	0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x12, 0x1f, 0x2e, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65,
	0x72, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d,
	0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65,
	0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x23, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x1d, 0x12, 0x1b, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x73,
	0x2f, 0x7b, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x7d, 0x12, 0x82,
	0x01, 0x0a, 0x11, 0x42, 0x61, 0x74, 0x63, 0x68, 0x47, 0x65, 0x74, 0x43, 0x75, 0x73, 0x74, 0x6f,
	0x6d, 0x65, 0x72, 0x73, 0x12, 0x25, 0x2e, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x2e,
	0x76, 0x31, 0x2e, 0x42, 0x61, 0x74, 0x63, 0x68, 0x47, 0x65, 0x74, 0x43, 0x75, 0x73, 0x74, 0x6f,
	0x6d, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x26, 0x2e, 0x63, 0x75,
	0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x42, 0x61, 0x74, 0x63, 0x68, 0x47,
	0x65, 0x74, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x1e, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x18, 0x12, 0x16, 0x2f, 0x76, 0x31,
	0x2f, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x73, 0x3a, 0x62, 0x61, 0x74, 0x63, 0x68,
	0x47, 0x65, 0x74, 0x42, 0x0e, 0x5a, 0x0c, 0x74, 0x6d, 0x70, 0x72, 0x6c, 0x64, 0x65, 0x6d, 0x6f,
	0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_customer_v1_customer_proto_rawDescOnce sync.Once
	file_customer_v1_customer_proto_rawDescData = file_customer_v1_customer_proto_rawDesc
)

func file_customer_v1_customer_proto_rawDescGZIP() []byte {
	file_customer_v1_customer_proto_rawDescOnce.Do(func() {
		file_customer_v1_customer_proto_rawDescData = protoimpl.X.CompressGZIP(file_customer_v1_customer_proto_rawDescData)
	})
	return file_customer_v1_customer_proto_rawDescData
}

var file_customer_v1_customer_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_customer_v1_customer_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_customer_v1_customer_proto_goTypes = []interface{}{
	(VerificationResult)(0),           // 0: customer.v1.VerificationResult
	(*Customer)(nil),                  // 1: customer.v1.Customer
	(*CreateCustomerRequest)(nil),     // 2: customer.v1.CreateCustomerRequest
	(*CreateCustomerResponse)(nil),    // 3: customer.v1.CreateCustomerResponse
	(*VerifyCustomerRequest)(nil),     // 4: customer.v1.VerifyCustomerRequest
	(*VerifyCustomerResponse)(nil),    // 5: customer.v1.VerifyCustomerResponse
	(*GetCustomerRequest)(nil),        // 6: customer.v1.GetCustomerRequest
	(*GetCustomerResponse)(nil),       // 7: customer.v1.GetCustomerResponse
	(*BatchGetCustomersRequest)(nil),  // 8: customer.v1.BatchGetCustomersRequest
	(*BatchGetCustomersResponse)(nil), // 9: customer.v1.BatchGetCustomersResponse
}
var file_customer_v1_customer_proto_depIdxs = []int32{
	1, // 0: customer.v1.CreateCustomerRequest.customer:type_name -> customer.v1.Customer
	1, // 1: customer.v1.CreateCustomerResponse.customer:type_name -> customer.v1.Customer
	0, // 2: customer.v1.VerifyCustomerResponse.result:type_name -> customer.v1.VerificationResult
	1, // 3: customer.v1.GetCustomerResponse.customer:type_name -> customer.v1.Customer
	1, // 4: customer.v1.BatchGetCustomersResponse.customers:type_name -> customer.v1.Customer
	2, // 5: customer.v1.CustomerService.CreateCustomer:input_type -> customer.v1.CreateCustomerRequest
	4, // 6: customer.v1.CustomerService.VerifyCustomer:input_type -> customer.v1.VerifyCustomerRequest
	6, // 7: customer.v1.CustomerService.GetCustomer:input_type -> customer.v1.GetCustomerRequest
	8, // 8: customer.v1.CustomerService.BatchGetCustomers:input_type -> customer.v1.BatchGetCustomersRequest
	3, // 9: customer.v1.CustomerService.CreateCustomer:output_type -> customer.v1.CreateCustomerResponse
	5, // 10: customer.v1.CustomerService.VerifyCustomer:output_type -> customer.v1.VerifyCustomerResponse
	7, // 11: customer.v1.CustomerService.GetCustomer:output_type -> customer.v1.GetCustomerResponse
	9, // 12: customer.v1.CustomerService.BatchGetCustomers:output_type -> customer.v1.BatchGetCustomersResponse
	9, // [9:13] is the sub-list for method output_type
	5, // [5:9] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_customer_v1_customer_proto_init() }
func file_customer_v1_customer_proto_init() {
	if File_customer_v1_customer_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_customer_v1_customer_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Customer); i {
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
		file_customer_v1_customer_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateCustomerRequest); i {
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
		file_customer_v1_customer_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateCustomerResponse); i {
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
		file_customer_v1_customer_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VerifyCustomerRequest); i {
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
		file_customer_v1_customer_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VerifyCustomerResponse); i {
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
		file_customer_v1_customer_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCustomerRequest); i {
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
		file_customer_v1_customer_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCustomerResponse); i {
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
		file_customer_v1_customer_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BatchGetCustomersRequest); i {
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
		file_customer_v1_customer_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BatchGetCustomersResponse); i {
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
			RawDescriptor: file_customer_v1_customer_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_customer_v1_customer_proto_goTypes,
		DependencyIndexes: file_customer_v1_customer_proto_depIdxs,
		EnumInfos:         file_customer_v1_customer_proto_enumTypes,
		MessageInfos:      file_customer_v1_customer_proto_msgTypes,
	}.Build()
	File_customer_v1_customer_proto = out.File
	file_customer_v1_customer_proto_rawDesc = nil
	file_customer_v1_customer_proto_goTypes = nil
	file_customer_v1_customer_proto_depIdxs = nil
}
