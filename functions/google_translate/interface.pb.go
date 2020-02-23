// Code generated by protoc-gen-go. DO NOT EDIT.
// source: interface.proto

package main

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type Input struct {
	InputLanguage        string   `protobuf:"bytes,1,opt,name=inputLanguage,proto3" json:"inputLanguage,omitempty"`
	OutputLanguage       string   `protobuf:"bytes,2,opt,name=outputLanguage,proto3" json:"outputLanguage,omitempty"`
	Text                 string   `protobuf:"bytes,3,opt,name=text,proto3" json:"text,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Input) Reset()         { *m = Input{} }
func (m *Input) String() string { return proto.CompactTextString(m) }
func (*Input) ProtoMessage()    {}
func (*Input) Descriptor() ([]byte, []int) {
	return fileDescriptor_3ef53c9e620778f1, []int{0}
}

func (m *Input) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Input.Unmarshal(m, b)
}
func (m *Input) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Input.Marshal(b, m, deterministic)
}
func (m *Input) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Input.Merge(m, src)
}
func (m *Input) XXX_Size() int {
	return xxx_messageInfo_Input.Size(m)
}
func (m *Input) XXX_DiscardUnknown() {
	xxx_messageInfo_Input.DiscardUnknown(m)
}

var xxx_messageInfo_Input proto.InternalMessageInfo

func (m *Input) GetInputLanguage() string {
	if m != nil {
		return m.InputLanguage
	}
	return ""
}

func (m *Input) GetOutputLanguage() string {
	if m != nil {
		return m.OutputLanguage
	}
	return ""
}

func (m *Input) GetText() string {
	if m != nil {
		return m.Text
	}
	return ""
}

type Output struct {
	Text                 string   `protobuf:"bytes,1,opt,name=text,proto3" json:"text,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Output) Reset()         { *m = Output{} }
func (m *Output) String() string { return proto.CompactTextString(m) }
func (*Output) ProtoMessage()    {}
func (*Output) Descriptor() ([]byte, []int) {
	return fileDescriptor_3ef53c9e620778f1, []int{1}
}

func (m *Output) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Output.Unmarshal(m, b)
}
func (m *Output) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Output.Marshal(b, m, deterministic)
}
func (m *Output) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Output.Merge(m, src)
}
func (m *Output) XXX_Size() int {
	return xxx_messageInfo_Output.Size(m)
}
func (m *Output) XXX_DiscardUnknown() {
	xxx_messageInfo_Output.DiscardUnknown(m)
}

var xxx_messageInfo_Output proto.InternalMessageInfo

func (m *Output) GetText() string {
	if m != nil {
		return m.Text
	}
	return ""
}

func init() {
	proto.RegisterType((*Input)(nil), "main.Input")
	proto.RegisterType((*Output)(nil), "main.Output")
}

func init() { proto.RegisterFile("interface.proto", fileDescriptor_3ef53c9e620778f1) }

var fileDescriptor_3ef53c9e620778f1 = []byte{
	// 125 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xcf, 0xcc, 0x2b, 0x49,
	0x2d, 0x4a, 0x4b, 0x4c, 0x4e, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0xc9, 0x4d, 0xcc,
	0xcc, 0x53, 0xca, 0xe4, 0x62, 0xf5, 0xcc, 0x2b, 0x28, 0x2d, 0x11, 0x52, 0xe1, 0xe2, 0xcd, 0x04,
	0x31, 0x7c, 0x12, 0xf3, 0xd2, 0x4b, 0x13, 0xd3, 0x53, 0x25, 0x18, 0x15, 0x18, 0x35, 0x38, 0x83,
	0x50, 0x05, 0x85, 0xd4, 0xb8, 0xf8, 0xf2, 0x4b, 0x4b, 0x90, 0x95, 0x31, 0x81, 0x95, 0xa1, 0x89,
	0x0a, 0x09, 0x71, 0xb1, 0x94, 0xa4, 0x56, 0x94, 0x48, 0x30, 0x83, 0x65, 0xc1, 0x6c, 0x25, 0x19,
	0x2e, 0x36, 0x7f, 0xb0, 0x2a, 0xb8, 0x2c, 0x23, 0x42, 0x36, 0x89, 0x0d, 0xec, 0x2a, 0x63, 0x40,
	0x00, 0x00, 0x00, 0xff, 0xff, 0x49, 0xc4, 0x87, 0xb8, 0xa8, 0x00, 0x00, 0x00,
}