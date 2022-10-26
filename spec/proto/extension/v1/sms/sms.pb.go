// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.17.3
// source: spec/proto/extension/v1/sms/sms.proto

package sms

import (
	reflect "reflect"
	sync "sync"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// SendSmsRequest is the request of the `SendSms` method.
type SendSmsWithTemplateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. The saas service name
	//  If your system uses multiple SMS services at the same time,
	//  you can specify which service to use with this field.
	ComponentName string `protobuf:"bytes,1,opt,name=component_name,json=componentName,proto3" json:"component_name,omitempty"`
	// Required. The SMS receive phone numbers.
	PhoneNumbers []string `protobuf:"bytes,2,rep,name=phone_numbers,json=phoneNumbers,proto3" json:"phone_numbers,omitempty"`
	// Required.
	Template *Template `protobuf:"bytes,3,opt,name=template,proto3" json:"template,omitempty"`
	// The registered sign name
	SignName string `protobuf:"bytes,4,opt,name=sign_name,json=signName,proto3" json:"sign_name,omitempty"`
	// The SMS sender tag.
	SenderId string `protobuf:"bytes,5,opt,name=sender_id,json=senderId,proto3" json:"sender_id,omitempty"`
	// The metadata which will be sent to SMS components.
	Metadata map[string]string `protobuf:"bytes,6,rep,name=metadata,proto3" json:"metadata,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *SendSmsWithTemplateRequest) Reset() {
	*x = SendSmsWithTemplateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spec_proto_extension_v1_sms_sms_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendSmsWithTemplateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendSmsWithTemplateRequest) ProtoMessage() {}

func (x *SendSmsWithTemplateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_spec_proto_extension_v1_sms_sms_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendSmsWithTemplateRequest.ProtoReflect.Descriptor instead.
func (*SendSmsWithTemplateRequest) Descriptor() ([]byte, []int) {
	return file_spec_proto_extension_v1_sms_sms_proto_rawDescGZIP(), []int{0}
}

func (x *SendSmsWithTemplateRequest) GetComponentName() string {
	if x != nil {
		return x.ComponentName
	}
	return ""
}

func (x *SendSmsWithTemplateRequest) GetPhoneNumbers() []string {
	if x != nil {
		return x.PhoneNumbers
	}
	return nil
}

func (x *SendSmsWithTemplateRequest) GetTemplate() *Template {
	if x != nil {
		return x.Template
	}
	return nil
}

func (x *SendSmsWithTemplateRequest) GetSignName() string {
	if x != nil {
		return x.SignName
	}
	return ""
}

func (x *SendSmsWithTemplateRequest) GetSenderId() string {
	if x != nil {
		return x.SenderId
	}
	return ""
}

func (x *SendSmsWithTemplateRequest) GetMetadata() map[string]string {
	if x != nil {
		return x.Metadata
	}
	return nil
}

// Sms template
type Template struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required
	TemplateId string `protobuf:"bytes,1,opt,name=template_id,json=templateId,proto3" json:"template_id,omitempty"`
	// Required
	TemplateParams map[string]string `protobuf:"bytes,2,rep,name=template_params,json=templateParams,proto3" json:"template_params,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *Template) Reset() {
	*x = Template{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spec_proto_extension_v1_sms_sms_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Template) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Template) ProtoMessage() {}

func (x *Template) ProtoReflect() protoreflect.Message {
	mi := &file_spec_proto_extension_v1_sms_sms_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Template.ProtoReflect.Descriptor instead.
func (*Template) Descriptor() ([]byte, []int) {
	return file_spec_proto_extension_v1_sms_sms_proto_rawDescGZIP(), []int{1}
}

func (x *Template) GetTemplateId() string {
	if x != nil {
		return x.TemplateId
	}
	return ""
}

func (x *Template) GetTemplateParams() map[string]string {
	if x != nil {
		return x.TemplateParams
	}
	return nil
}

// SendSmsResponse is the reponse of the `SendSms` method.
type SendSmsWithTemplateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The unique requestId.
	RequestId string `protobuf:"bytes,1,opt,name=request_id,json=requestId,proto3" json:"request_id,omitempty"`
	// "OK" represents success.
	Code string `protobuf:"bytes,2,opt,name=code,proto3" json:"code,omitempty"`
	// The error message.
	Message string `protobuf:"bytes,3,opt,name=message,proto3" json:"message,omitempty"`
	// The metadata returned from SMS service.
	Metadata map[string]string `protobuf:"bytes,4,rep,name=metadata,proto3" json:"metadata,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *SendSmsWithTemplateResponse) Reset() {
	*x = SendSmsWithTemplateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spec_proto_extension_v1_sms_sms_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendSmsWithTemplateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendSmsWithTemplateResponse) ProtoMessage() {}

func (x *SendSmsWithTemplateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_spec_proto_extension_v1_sms_sms_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendSmsWithTemplateResponse.ProtoReflect.Descriptor instead.
func (*SendSmsWithTemplateResponse) Descriptor() ([]byte, []int) {
	return file_spec_proto_extension_v1_sms_sms_proto_rawDescGZIP(), []int{2}
}

func (x *SendSmsWithTemplateResponse) GetRequestId() string {
	if x != nil {
		return x.RequestId
	}
	return ""
}

func (x *SendSmsWithTemplateResponse) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *SendSmsWithTemplateResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *SendSmsWithTemplateResponse) GetMetadata() map[string]string {
	if x != nil {
		return x.Metadata
	}
	return nil
}

var File_spec_proto_extension_v1_sms_sms_proto protoreflect.FileDescriptor

var file_spec_proto_extension_v1_sms_sms_proto_rawDesc = []byte{
	0x0a, 0x25, 0x73, 0x70, 0x65, 0x63, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x65, 0x78, 0x74,
	0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x6d, 0x73, 0x2f, 0x73, 0x6d,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1b, 0x73, 0x70, 0x65, 0x63, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31,
	0x2e, 0x73, 0x6d, 0x73, 0x22, 0x85, 0x03, 0x0a, 0x1a, 0x53, 0x65, 0x6e, 0x64, 0x53, 0x6d, 0x73,
	0x57, 0x69, 0x74, 0x68, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x25, 0x0a, 0x0e, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74,
	0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x63, 0x6f, 0x6d,
	0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x70, 0x68,
	0x6f, 0x6e, 0x65, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28,
	0x09, 0x52, 0x0c, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x12,
	0x41, 0x0a, 0x08, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x25, 0x2e, 0x73, 0x70, 0x65, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x65,
	0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x73, 0x6d, 0x73, 0x2e,
	0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x52, 0x08, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61,
	0x74, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x69, 0x67, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x12,
	0x1b, 0x0a, 0x09, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x49, 0x64, 0x12, 0x61, 0x0a, 0x08,
	0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x45,
	0x2e, 0x73, 0x70, 0x65, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x65, 0x78, 0x74, 0x65,
	0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x73, 0x6d, 0x73, 0x2e, 0x53, 0x65, 0x6e,
	0x64, 0x53, 0x6d, 0x73, 0x57, 0x69, 0x74, 0x68, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x1a,
	0x3b, 0x0a, 0x0d, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b,
	0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0xd2, 0x01, 0x0a,
	0x08, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x74, 0x65, 0x6d,
	0x70, 0x6c, 0x61, 0x74, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a,
	0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x49, 0x64, 0x12, 0x62, 0x0a, 0x0f, 0x74, 0x65,
	0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x5f, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x18, 0x02, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x39, 0x2e, 0x73, 0x70, 0x65, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x73, 0x6d,
	0x73, 0x2e, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x2e, 0x54, 0x65, 0x6d, 0x70, 0x6c,
	0x61, 0x74, 0x65, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0e,
	0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x1a, 0x41,
	0x0a, 0x13, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38,
	0x01, 0x22, 0x8b, 0x02, 0x0a, 0x1b, 0x53, 0x65, 0x6e, 0x64, 0x53, 0x6d, 0x73, 0x57, 0x69, 0x74,
	0x68, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x64,
	0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x63, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x62,
	0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x46, 0x2e, 0x73, 0x70, 0x65, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x65, 0x78,
	0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x73, 0x6d, 0x73, 0x2e, 0x53,
	0x65, 0x6e, 0x64, 0x53, 0x6d, 0x73, 0x57, 0x69, 0x74, 0x68, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61,
	0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0x1a, 0x3b, 0x0a, 0x0d, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x32,
	0x99, 0x01, 0x0a, 0x0a, 0x53, 0x6d, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x8a,
	0x01, 0x0a, 0x13, 0x53, 0x65, 0x6e, 0x64, 0x53, 0x6d, 0x73, 0x57, 0x69, 0x74, 0x68, 0x54, 0x65,
	0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x12, 0x37, 0x2e, 0x73, 0x70, 0x65, 0x63, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31,
	0x2e, 0x73, 0x6d, 0x73, 0x2e, 0x53, 0x65, 0x6e, 0x64, 0x53, 0x6d, 0x73, 0x57, 0x69, 0x74, 0x68,
	0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x38, 0x2e, 0x73, 0x70, 0x65, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x65, 0x78, 0x74,
	0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x73, 0x6d, 0x73, 0x2e, 0x53, 0x65,
	0x6e, 0x64, 0x53, 0x6d, 0x73, 0x57, 0x69, 0x74, 0x68, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x31, 0x5a, 0x2f, 0x6d,
	0x6f, 0x73, 0x6e, 0x2e, 0x69, 0x6f, 0x2f, 0x6c, 0x61, 0x79, 0x6f, 0x74, 0x74, 0x6f, 0x2f, 0x73,
	0x70, 0x65, 0x63, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73,
	0x69, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x6d, 0x73, 0x3b, 0x73, 0x6d, 0x73, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_spec_proto_extension_v1_sms_sms_proto_rawDescOnce sync.Once
	file_spec_proto_extension_v1_sms_sms_proto_rawDescData = file_spec_proto_extension_v1_sms_sms_proto_rawDesc
)

func file_spec_proto_extension_v1_sms_sms_proto_rawDescGZIP() []byte {
	file_spec_proto_extension_v1_sms_sms_proto_rawDescOnce.Do(func() {
		file_spec_proto_extension_v1_sms_sms_proto_rawDescData = protoimpl.X.CompressGZIP(file_spec_proto_extension_v1_sms_sms_proto_rawDescData)
	})
	return file_spec_proto_extension_v1_sms_sms_proto_rawDescData
}

var file_spec_proto_extension_v1_sms_sms_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_spec_proto_extension_v1_sms_sms_proto_goTypes = []interface{}{
	(*SendSmsWithTemplateRequest)(nil),  // 0: spec.proto.extension.v1.sms.SendSmsWithTemplateRequest
	(*Template)(nil),                    // 1: spec.proto.extension.v1.sms.Template
	(*SendSmsWithTemplateResponse)(nil), // 2: spec.proto.extension.v1.sms.SendSmsWithTemplateResponse
	nil,                                 // 3: spec.proto.extension.v1.sms.SendSmsWithTemplateRequest.MetadataEntry
	nil,                                 // 4: spec.proto.extension.v1.sms.Template.TemplateParamsEntry
	nil,                                 // 5: spec.proto.extension.v1.sms.SendSmsWithTemplateResponse.MetadataEntry
}
var file_spec_proto_extension_v1_sms_sms_proto_depIdxs = []int32{
	1, // 0: spec.proto.extension.v1.sms.SendSmsWithTemplateRequest.template:type_name -> spec.proto.extension.v1.sms.Template
	3, // 1: spec.proto.extension.v1.sms.SendSmsWithTemplateRequest.metadata:type_name -> spec.proto.extension.v1.sms.SendSmsWithTemplateRequest.MetadataEntry
	4, // 2: spec.proto.extension.v1.sms.Template.template_params:type_name -> spec.proto.extension.v1.sms.Template.TemplateParamsEntry
	5, // 3: spec.proto.extension.v1.sms.SendSmsWithTemplateResponse.metadata:type_name -> spec.proto.extension.v1.sms.SendSmsWithTemplateResponse.MetadataEntry
	0, // 4: spec.proto.extension.v1.sms.SmsService.SendSmsWithTemplate:input_type -> spec.proto.extension.v1.sms.SendSmsWithTemplateRequest
	2, // 5: spec.proto.extension.v1.sms.SmsService.SendSmsWithTemplate:output_type -> spec.proto.extension.v1.sms.SendSmsWithTemplateResponse
	5, // [5:6] is the sub-list for method output_type
	4, // [4:5] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_spec_proto_extension_v1_sms_sms_proto_init() }
func file_spec_proto_extension_v1_sms_sms_proto_init() {
	if File_spec_proto_extension_v1_sms_sms_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_spec_proto_extension_v1_sms_sms_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendSmsWithTemplateRequest); i {
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
		file_spec_proto_extension_v1_sms_sms_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Template); i {
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
		file_spec_proto_extension_v1_sms_sms_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendSmsWithTemplateResponse); i {
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
			RawDescriptor: file_spec_proto_extension_v1_sms_sms_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_spec_proto_extension_v1_sms_sms_proto_goTypes,
		DependencyIndexes: file_spec_proto_extension_v1_sms_sms_proto_depIdxs,
		MessageInfos:      file_spec_proto_extension_v1_sms_sms_proto_msgTypes,
	}.Build()
	File_spec_proto_extension_v1_sms_sms_proto = out.File
	file_spec_proto_extension_v1_sms_sms_proto_rawDesc = nil
	file_spec_proto_extension_v1_sms_sms_proto_goTypes = nil
	file_spec_proto_extension_v1_sms_sms_proto_depIdxs = nil
}
