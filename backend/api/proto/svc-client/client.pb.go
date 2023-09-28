// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.12.4
// source: api/proto/svc-client/client.proto

package svc_client

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

type GetDataClientRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId int64 `protobuf:"varint,1,opt,name=userId,proto3" json:"userId,omitempty"`
}

func (x *GetDataClientRequest) Reset() {
	*x = GetDataClientRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_svc_client_client_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetDataClientRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetDataClientRequest) ProtoMessage() {}

func (x *GetDataClientRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_svc_client_client_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetDataClientRequest.ProtoReflect.Descriptor instead.
func (*GetDataClientRequest) Descriptor() ([]byte, []int) {
	return file_api_proto_svc_client_client_proto_rawDescGZIP(), []int{0}
}

func (x *GetDataClientRequest) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

type GetDataClientResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DataClient     *DataClient     `protobuf:"bytes,1,opt,name=dataClient,proto3" json:"dataClient,omitempty"`
	ResponseStatus *ResponseStatus `protobuf:"bytes,2,opt,name=responseStatus,proto3" json:"responseStatus,omitempty"`
}

func (x *GetDataClientResponse) Reset() {
	*x = GetDataClientResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_svc_client_client_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetDataClientResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetDataClientResponse) ProtoMessage() {}

func (x *GetDataClientResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_svc_client_client_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetDataClientResponse.ProtoReflect.Descriptor instead.
func (*GetDataClientResponse) Descriptor() ([]byte, []int) {
	return file_api_proto_svc_client_client_proto_rawDescGZIP(), []int{1}
}

func (x *GetDataClientResponse) GetDataClient() *DataClient {
	if x != nil {
		return x.DataClient
	}
	return nil
}

func (x *GetDataClientResponse) GetResponseStatus() *ResponseStatus {
	if x != nil {
		return x.ResponseStatus
	}
	return nil
}

type SaveDataClientRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DataClient *DataClient `protobuf:"bytes,1,opt,name=dataClient,proto3" json:"dataClient,omitempty"`
}

func (x *SaveDataClientRequest) Reset() {
	*x = SaveDataClientRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_svc_client_client_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SaveDataClientRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SaveDataClientRequest) ProtoMessage() {}

func (x *SaveDataClientRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_svc_client_client_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SaveDataClientRequest.ProtoReflect.Descriptor instead.
func (*SaveDataClientRequest) Descriptor() ([]byte, []int) {
	return file_api_proto_svc_client_client_proto_rawDescGZIP(), []int{2}
}

func (x *SaveDataClientRequest) GetDataClient() *DataClient {
	if x != nil {
		return x.DataClient
	}
	return nil
}

type SaveDataClientResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ResponseStatus *ResponseStatus `protobuf:"bytes,1,opt,name=responseStatus,proto3" json:"responseStatus,omitempty"`
}

func (x *SaveDataClientResponse) Reset() {
	*x = SaveDataClientResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_svc_client_client_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SaveDataClientResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SaveDataClientResponse) ProtoMessage() {}

func (x *SaveDataClientResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_svc_client_client_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SaveDataClientResponse.ProtoReflect.Descriptor instead.
func (*SaveDataClientResponse) Descriptor() ([]byte, []int) {
	return file_api_proto_svc_client_client_proto_rawDescGZIP(), []int{3}
}

func (x *SaveDataClientResponse) GetResponseStatus() *ResponseStatus {
	if x != nil {
		return x.ResponseStatus
	}
	return nil
}

type DataClient struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId            int64              `protobuf:"varint,1,opt,name=userId,proto3" json:"userId,omitempty"`
	FirstName         string             `protobuf:"bytes,2,opt,name=firstName,proto3" json:"firstName,omitempty"`
	LastName          string             `protobuf:"bytes,3,opt,name=lastName,proto3" json:"lastName,omitempty"`
	RegDate           string             `protobuf:"bytes,4,opt,name=regDate,proto3" json:"regDate,omitempty"`
	LanguageItems     []*LanguageItems   `protobuf:"bytes,5,rep,name=languageItems,proto3" json:"languageItems,omitempty"`
	AboutMe           string             `protobuf:"bytes,6,opt,name=aboutMe,proto3" json:"aboutMe,omitempty"`
	Avatar            string             `protobuf:"bytes,7,opt,name=avatar,proto3" json:"avatar,omitempty"`
	DeleteAvatar      string             `protobuf:"bytes,8,opt,name=deleteAvatar,proto3" json:"deleteAvatar,omitempty"`
	InterfaceLang     string             `protobuf:"bytes,9,opt,name=interfaceLang,proto3" json:"interfaceLang,omitempty"`
	IsTeacher         bool               `protobuf:"varint,10,opt,name=isTeacher,proto3" json:"isTeacher,omitempty"`
	RegistrationToken *RegistrationToken `protobuf:"bytes,11,opt,name=registrationToken,proto3" json:"registrationToken,omitempty"`
	SubStudents       []*SubStudents     `protobuf:"bytes,12,rep,name=subStudents,proto3" json:"subStudents,omitempty"`
}

func (x *DataClient) Reset() {
	*x = DataClient{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_svc_client_client_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DataClient) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DataClient) ProtoMessage() {}

func (x *DataClient) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_svc_client_client_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DataClient.ProtoReflect.Descriptor instead.
func (*DataClient) Descriptor() ([]byte, []int) {
	return file_api_proto_svc_client_client_proto_rawDescGZIP(), []int{4}
}

func (x *DataClient) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *DataClient) GetFirstName() string {
	if x != nil {
		return x.FirstName
	}
	return ""
}

func (x *DataClient) GetLastName() string {
	if x != nil {
		return x.LastName
	}
	return ""
}

func (x *DataClient) GetRegDate() string {
	if x != nil {
		return x.RegDate
	}
	return ""
}

func (x *DataClient) GetLanguageItems() []*LanguageItems {
	if x != nil {
		return x.LanguageItems
	}
	return nil
}

func (x *DataClient) GetAboutMe() string {
	if x != nil {
		return x.AboutMe
	}
	return ""
}

func (x *DataClient) GetAvatar() string {
	if x != nil {
		return x.Avatar
	}
	return ""
}

func (x *DataClient) GetDeleteAvatar() string {
	if x != nil {
		return x.DeleteAvatar
	}
	return ""
}

func (x *DataClient) GetInterfaceLang() string {
	if x != nil {
		return x.InterfaceLang
	}
	return ""
}

func (x *DataClient) GetIsTeacher() bool {
	if x != nil {
		return x.IsTeacher
	}
	return false
}

func (x *DataClient) GetRegistrationToken() *RegistrationToken {
	if x != nil {
		return x.RegistrationToken
	}
	return nil
}

func (x *DataClient) GetSubStudents() []*SubStudents {
	if x != nil {
		return x.SubStudents
	}
	return nil
}

type SubStudents struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Avatar    string `protobuf:"bytes,2,opt,name=avatar,proto3" json:"avatar,omitempty"`
	FirstName string `protobuf:"bytes,3,opt,name=firstName,proto3" json:"firstName,omitempty"`
	LastName  string `protobuf:"bytes,4,opt,name=lastName,proto3" json:"lastName,omitempty"`
}

func (x *SubStudents) Reset() {
	*x = SubStudents{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_svc_client_client_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SubStudents) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubStudents) ProtoMessage() {}

func (x *SubStudents) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_svc_client_client_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubStudents.ProtoReflect.Descriptor instead.
func (*SubStudents) Descriptor() ([]byte, []int) {
	return file_api_proto_svc_client_client_proto_rawDescGZIP(), []int{5}
}

func (x *SubStudents) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *SubStudents) GetAvatar() string {
	if x != nil {
		return x.Avatar
	}
	return ""
}

func (x *SubStudents) GetFirstName() string {
	if x != nil {
		return x.FirstName
	}
	return ""
}

func (x *SubStudents) GetLastName() string {
	if x != nil {
		return x.LastName
	}
	return ""
}

type RegistrationToken struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RegToken     string `protobuf:"bytes,1,opt,name=regToken,proto3" json:"regToken,omitempty"`
	CreatedDate  string `protobuf:"bytes,2,opt,name=createdDate,proto3" json:"createdDate,omitempty"`
	DeadlineDate string `protobuf:"bytes,3,opt,name=deadlineDate,proto3" json:"deadlineDate,omitempty"`
}

func (x *RegistrationToken) Reset() {
	*x = RegistrationToken{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_svc_client_client_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegistrationToken) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegistrationToken) ProtoMessage() {}

func (x *RegistrationToken) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_svc_client_client_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegistrationToken.ProtoReflect.Descriptor instead.
func (*RegistrationToken) Descriptor() ([]byte, []int) {
	return file_api_proto_svc_client_client_proto_rawDescGZIP(), []int{6}
}

func (x *RegistrationToken) GetRegToken() string {
	if x != nil {
		return x.RegToken
	}
	return ""
}

func (x *RegistrationToken) GetCreatedDate() string {
	if x != nil {
		return x.CreatedDate
	}
	return ""
}

func (x *RegistrationToken) GetDeadlineDate() string {
	if x != nil {
		return x.DeadlineDate
	}
	return ""
}

type LanguageItems struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Lang string `protobuf:"bytes,1,opt,name=lang,proto3" json:"lang,omitempty"`
	Lvl  string `protobuf:"bytes,2,opt,name=lvl,proto3" json:"lvl,omitempty"`
}

func (x *LanguageItems) Reset() {
	*x = LanguageItems{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_svc_client_client_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LanguageItems) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LanguageItems) ProtoMessage() {}

func (x *LanguageItems) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_svc_client_client_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LanguageItems.ProtoReflect.Descriptor instead.
func (*LanguageItems) Descriptor() ([]byte, []int) {
	return file_api_proto_svc_client_client_proto_rawDescGZIP(), []int{7}
}

func (x *LanguageItems) GetLang() string {
	if x != nil {
		return x.Lang
	}
	return ""
}

func (x *LanguageItems) GetLvl() string {
	if x != nil {
		return x.Lvl
	}
	return ""
}

type ResponseStatus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status int64  `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	Error  string `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *ResponseStatus) Reset() {
	*x = ResponseStatus{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_svc_client_client_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResponseStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResponseStatus) ProtoMessage() {}

func (x *ResponseStatus) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_svc_client_client_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResponseStatus.ProtoReflect.Descriptor instead.
func (*ResponseStatus) Descriptor() ([]byte, []int) {
	return file_api_proto_svc_client_client_proto_rawDescGZIP(), []int{8}
}

func (x *ResponseStatus) GetStatus() int64 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *ResponseStatus) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

var File_api_proto_svc_client_client_proto protoreflect.FileDescriptor

var file_api_proto_svc_client_client_proto_rawDesc = []byte{
	0x0a, 0x21, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73, 0x76, 0x63, 0x2d,
	0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x06, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x22, 0x2e, 0x0a, 0x14, 0x47,
	0x65, 0x74, 0x44, 0x61, 0x74, 0x61, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x8b, 0x01, 0x0a, 0x15,
	0x47, 0x65, 0x74, 0x44, 0x61, 0x74, 0x61, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x32, 0x0a, 0x0a, 0x64, 0x61, 0x74, 0x61, 0x43, 0x6c, 0x69,
	0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x63, 0x6c, 0x69, 0x65,
	0x6e, 0x74, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x52, 0x0a, 0x64,
	0x61, 0x74, 0x61, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x12, 0x3e, 0x0a, 0x0e, 0x72, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x16, 0x2e, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x0e, 0x72, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x4b, 0x0a, 0x15, 0x53, 0x61, 0x76,
	0x65, 0x44, 0x61, 0x74, 0x61, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x32, 0x0a, 0x0a, 0x64, 0x61, 0x74, 0x61, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e,
	0x44, 0x61, 0x74, 0x61, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x52, 0x0a, 0x64, 0x61, 0x74, 0x61,
	0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x22, 0x58, 0x0a, 0x16, 0x53, 0x61, 0x76, 0x65, 0x44, 0x61,
	0x74, 0x61, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x3e, 0x0a, 0x0e, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x63, 0x6c, 0x69, 0x65, 0x6e,
	0x74, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x52, 0x0e, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x22, 0xcf, 0x03, 0x0a, 0x0a, 0x44, 0x61, 0x74, 0x61, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x12,
	0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x66, 0x69, 0x72, 0x73, 0x74,
	0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x66, 0x69, 0x72, 0x73,
	0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x61, 0x73, 0x74, 0x4e, 0x61, 0x6d,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x61, 0x73, 0x74, 0x4e, 0x61, 0x6d,
	0x65, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x67, 0x44, 0x61, 0x74, 0x65, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x72, 0x65, 0x67, 0x44, 0x61, 0x74, 0x65, 0x12, 0x3b, 0x0a, 0x0d, 0x6c,
	0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x05, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x15, 0x2e, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x4c, 0x61, 0x6e, 0x67,
	0x75, 0x61, 0x67, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x52, 0x0d, 0x6c, 0x61, 0x6e, 0x67, 0x75,
	0x61, 0x67, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x62, 0x6f, 0x75,
	0x74, 0x4d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x62, 0x6f, 0x75, 0x74,
	0x4d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x12, 0x22, 0x0a, 0x0c, 0x64, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0c, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x12, 0x24,
	0x0a, 0x0d, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x4c, 0x61, 0x6e, 0x67, 0x18,
	0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65,
	0x4c, 0x61, 0x6e, 0x67, 0x12, 0x1c, 0x0a, 0x09, 0x69, 0x73, 0x54, 0x65, 0x61, 0x63, 0x68, 0x65,
	0x72, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x69, 0x73, 0x54, 0x65, 0x61, 0x63, 0x68,
	0x65, 0x72, 0x12, 0x47, 0x0a, 0x11, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e,
	0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x11, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x35, 0x0a, 0x0b, 0x73,
	0x75, 0x62, 0x53, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x0c, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x13, 0x2e, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x53, 0x75, 0x62, 0x53, 0x74, 0x75,
	0x64, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x0b, 0x73, 0x75, 0x62, 0x53, 0x74, 0x75, 0x64, 0x65, 0x6e,
	0x74, 0x73, 0x22, 0x6f, 0x0a, 0x0b, 0x53, 0x75, 0x62, 0x53, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74,
	0x73, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x12, 0x1c, 0x0a, 0x09, 0x66, 0x69, 0x72,
	0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x66, 0x69,
	0x72, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x61, 0x73, 0x74, 0x4e,
	0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x61, 0x73, 0x74, 0x4e,
	0x61, 0x6d, 0x65, 0x22, 0x75, 0x0a, 0x11, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x67, 0x54,
	0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x65, 0x67, 0x54,
	0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x20, 0x0a, 0x0b, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x44,
	0x61, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x64, 0x44, 0x61, 0x74, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x64, 0x65, 0x61, 0x64, 0x6c, 0x69,
	0x6e, 0x65, 0x44, 0x61, 0x74, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x64, 0x65,
	0x61, 0x64, 0x6c, 0x69, 0x6e, 0x65, 0x44, 0x61, 0x74, 0x65, 0x22, 0x35, 0x0a, 0x0d, 0x4c, 0x61,
	0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x6c,
	0x61, 0x6e, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6c, 0x61, 0x6e, 0x67, 0x12,
	0x10, 0x0a, 0x03, 0x6c, 0x76, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6c, 0x76,
	0x6c, 0x22, 0x3e, 0x0a, 0x0e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x65,
	0x72, 0x72, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f,
	0x72, 0x32, 0xb2, 0x01, 0x0a, 0x0d, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x4e, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x44, 0x61, 0x74, 0x61, 0x43, 0x6c,
	0x69, 0x65, 0x6e, 0x74, 0x12, 0x1c, 0x2e, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x47, 0x65,
	0x74, 0x44, 0x61, 0x74, 0x61, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x47, 0x65, 0x74, 0x44,
	0x61, 0x74, 0x61, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x12, 0x51, 0x0a, 0x0e, 0x53, 0x61, 0x76, 0x65, 0x44, 0x61, 0x74, 0x61, 0x43,
	0x6c, 0x69, 0x65, 0x6e, 0x74, 0x12, 0x1d, 0x2e, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x53,
	0x61, 0x76, 0x65, 0x44, 0x61, 0x74, 0x61, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x53, 0x61,
	0x76, 0x65, 0x44, 0x61, 0x74, 0x61, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x18, 0x5a, 0x16, 0x2e, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73, 0x76, 0x63, 0x2d, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_proto_svc_client_client_proto_rawDescOnce sync.Once
	file_api_proto_svc_client_client_proto_rawDescData = file_api_proto_svc_client_client_proto_rawDesc
)

func file_api_proto_svc_client_client_proto_rawDescGZIP() []byte {
	file_api_proto_svc_client_client_proto_rawDescOnce.Do(func() {
		file_api_proto_svc_client_client_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_proto_svc_client_client_proto_rawDescData)
	})
	return file_api_proto_svc_client_client_proto_rawDescData
}

var file_api_proto_svc_client_client_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_api_proto_svc_client_client_proto_goTypes = []interface{}{
	(*GetDataClientRequest)(nil),   // 0: client.GetDataClientRequest
	(*GetDataClientResponse)(nil),  // 1: client.GetDataClientResponse
	(*SaveDataClientRequest)(nil),  // 2: client.SaveDataClientRequest
	(*SaveDataClientResponse)(nil), // 3: client.SaveDataClientResponse
	(*DataClient)(nil),             // 4: client.DataClient
	(*SubStudents)(nil),            // 5: client.SubStudents
	(*RegistrationToken)(nil),      // 6: client.RegistrationToken
	(*LanguageItems)(nil),          // 7: client.LanguageItems
	(*ResponseStatus)(nil),         // 8: client.ResponseStatus
}
var file_api_proto_svc_client_client_proto_depIdxs = []int32{
	4, // 0: client.GetDataClientResponse.dataClient:type_name -> client.DataClient
	8, // 1: client.GetDataClientResponse.responseStatus:type_name -> client.ResponseStatus
	4, // 2: client.SaveDataClientRequest.dataClient:type_name -> client.DataClient
	8, // 3: client.SaveDataClientResponse.responseStatus:type_name -> client.ResponseStatus
	7, // 4: client.DataClient.languageItems:type_name -> client.LanguageItems
	6, // 5: client.DataClient.registrationToken:type_name -> client.RegistrationToken
	5, // 6: client.DataClient.subStudents:type_name -> client.SubStudents
	0, // 7: client.ClientService.GetDataClient:input_type -> client.GetDataClientRequest
	2, // 8: client.ClientService.SaveDataClient:input_type -> client.SaveDataClientRequest
	1, // 9: client.ClientService.GetDataClient:output_type -> client.GetDataClientResponse
	3, // 10: client.ClientService.SaveDataClient:output_type -> client.SaveDataClientResponse
	9, // [9:11] is the sub-list for method output_type
	7, // [7:9] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_api_proto_svc_client_client_proto_init() }
func file_api_proto_svc_client_client_proto_init() {
	if File_api_proto_svc_client_client_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_proto_svc_client_client_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetDataClientRequest); i {
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
		file_api_proto_svc_client_client_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetDataClientResponse); i {
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
		file_api_proto_svc_client_client_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SaveDataClientRequest); i {
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
		file_api_proto_svc_client_client_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SaveDataClientResponse); i {
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
		file_api_proto_svc_client_client_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DataClient); i {
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
		file_api_proto_svc_client_client_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SubStudents); i {
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
		file_api_proto_svc_client_client_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegistrationToken); i {
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
		file_api_proto_svc_client_client_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LanguageItems); i {
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
		file_api_proto_svc_client_client_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResponseStatus); i {
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
			RawDescriptor: file_api_proto_svc_client_client_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_proto_svc_client_client_proto_goTypes,
		DependencyIndexes: file_api_proto_svc_client_client_proto_depIdxs,
		MessageInfos:      file_api_proto_svc_client_client_proto_msgTypes,
	}.Build()
	File_api_proto_svc_client_client_proto = out.File
	file_api_proto_svc_client_client_proto_rawDesc = nil
	file_api_proto_svc_client_client_proto_goTypes = nil
	file_api_proto_svc_client_client_proto_depIdxs = nil
}
