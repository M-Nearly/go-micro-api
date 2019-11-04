// Code generated by protoc-gen-go. DO NOT EDIT.
// source: api.proto

package proto

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

type Pair struct {
	Key                  string   `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Values               []string `protobuf:"bytes,2,rep,name=values,proto3" json:"values,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Pair) Reset()         { *m = Pair{} }
func (m *Pair) String() string { return proto.CompactTextString(m) }
func (*Pair) ProtoMessage()    {}
func (*Pair) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{0}
}

func (m *Pair) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Pair.Unmarshal(m, b)
}
func (m *Pair) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Pair.Marshal(b, m, deterministic)
}
func (m *Pair) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Pair.Merge(m, src)
}
func (m *Pair) XXX_Size() int {
	return xxx_messageInfo_Pair.Size(m)
}
func (m *Pair) XXX_DiscardUnknown() {
	xxx_messageInfo_Pair.DiscardUnknown(m)
}

var xxx_messageInfo_Pair proto.InternalMessageInfo

func (m *Pair) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *Pair) GetValues() []string {
	if m != nil {
		return m.Values
	}
	return nil
}

// A HTTP request as RPC
// Forward by the api handler
type Request struct {
	Method               string           `protobuf:"bytes,1,opt,name=method,proto3" json:"method,omitempty"`
	Path                 string           `protobuf:"bytes,2,opt,name=path,proto3" json:"path,omitempty"`
	Header               map[string]*Pair `protobuf:"bytes,3,rep,name=header,proto3" json:"header,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Get                  map[string]*Pair `protobuf:"bytes,4,rep,name=get,proto3" json:"get,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Post                 map[string]*Pair `protobuf:"bytes,5,rep,name=post,proto3" json:"post,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Body                 string           `protobuf:"bytes,6,opt,name=body,proto3" json:"body,omitempty"`
	Url                  string           `protobuf:"bytes,7,opt,name=url,proto3" json:"url,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{1}
}

func (m *Request) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Request.Unmarshal(m, b)
}
func (m *Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Request.Marshal(b, m, deterministic)
}
func (m *Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Request.Merge(m, src)
}
func (m *Request) XXX_Size() int {
	return xxx_messageInfo_Request.Size(m)
}
func (m *Request) XXX_DiscardUnknown() {
	xxx_messageInfo_Request.DiscardUnknown(m)
}

var xxx_messageInfo_Request proto.InternalMessageInfo

func (m *Request) GetMethod() string {
	if m != nil {
		return m.Method
	}
	return ""
}

func (m *Request) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

func (m *Request) GetHeader() map[string]*Pair {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *Request) GetGet() map[string]*Pair {
	if m != nil {
		return m.Get
	}
	return nil
}

func (m *Request) GetPost() map[string]*Pair {
	if m != nil {
		return m.Post
	}
	return nil
}

func (m *Request) GetBody() string {
	if m != nil {
		return m.Body
	}
	return ""
}

func (m *Request) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

// A HTTP response as RPC
// Expected response for the api handler
type Response struct {
	StatusCode           int32            `protobuf:"varint,1,opt,name=statusCode,proto3" json:"statusCode,omitempty"`
	Header               map[string]*Pair `protobuf:"bytes,2,rep,name=header,proto3" json:"header,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Body                 string           `protobuf:"bytes,3,opt,name=body,proto3" json:"body,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{2}
}

func (m *Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response.Unmarshal(m, b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response.Marshal(b, m, deterministic)
}
func (m *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(m, src)
}
func (m *Response) XXX_Size() int {
	return xxx_messageInfo_Response.Size(m)
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

func (m *Response) GetStatusCode() int32 {
	if m != nil {
		return m.StatusCode
	}
	return 0
}

func (m *Response) GetHeader() map[string]*Pair {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *Response) GetBody() string {
	if m != nil {
		return m.Body
	}
	return ""
}

// A HTTP event as RPC
// Forwarded by the event handler
type Event struct {
	// e.g login
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// uuid
	Id string `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	// unix timestamp of event
	Timestamp int64 `protobuf:"varint,3,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	// event headers
	Header map[string]*Pair `protobuf:"bytes,4,rep,name=header,proto3" json:"header,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// the event data
	Data                 string   `protobuf:"bytes,5,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Event) Reset()         { *m = Event{} }
func (m *Event) String() string { return proto.CompactTextString(m) }
func (*Event) ProtoMessage()    {}
func (*Event) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{3}
}

func (m *Event) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Event.Unmarshal(m, b)
}
func (m *Event) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Event.Marshal(b, m, deterministic)
}
func (m *Event) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Event.Merge(m, src)
}
func (m *Event) XXX_Size() int {
	return xxx_messageInfo_Event.Size(m)
}
func (m *Event) XXX_DiscardUnknown() {
	xxx_messageInfo_Event.DiscardUnknown(m)
}

var xxx_messageInfo_Event proto.InternalMessageInfo

func (m *Event) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Event) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Event) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

func (m *Event) GetHeader() map[string]*Pair {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *Event) GetData() string {
	if m != nil {
		return m.Data
	}
	return ""
}

func init() {
	proto.RegisterType((*Pair)(nil), "proto.Pair")
	proto.RegisterType((*Request)(nil), "proto.Request")
	proto.RegisterMapType((map[string]*Pair)(nil), "proto.Request.GetEntry")
	proto.RegisterMapType((map[string]*Pair)(nil), "proto.Request.HeaderEntry")
	proto.RegisterMapType((map[string]*Pair)(nil), "proto.Request.PostEntry")
	proto.RegisterType((*Response)(nil), "proto.Response")
	proto.RegisterMapType((map[string]*Pair)(nil), "proto.Response.HeaderEntry")
	proto.RegisterType((*Event)(nil), "proto.Event")
	proto.RegisterMapType((map[string]*Pair)(nil), "proto.Event.HeaderEntry")
}

func init() { proto.RegisterFile("api.proto", fileDescriptor_00212fb1f9d3bf1c) }

var fileDescriptor_00212fb1f9d3bf1c = []byte{
	// 380 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x52, 0xc1, 0x6e, 0xe2, 0x30,
	0x10, 0x55, 0xe2, 0x24, 0x90, 0x89, 0xb4, 0x5a, 0xf9, 0xb0, 0x6b, 0xb1, 0xab, 0x15, 0x9b, 0x13,
	0x2b, 0xad, 0x10, 0x82, 0x4b, 0xd5, 0x2b, 0xa5, 0xed, 0x11, 0xf9, 0x0f, 0x8c, 0x62, 0x95, 0xa8,
	0x24, 0x4e, 0x63, 0x07, 0x89, 0x6f, 0xe3, 0x43, 0xfa, 0x3b, 0x95, 0x27, 0x26, 0x50, 0x44, 0x2f,
	0x88, 0x53, 0x26, 0xe3, 0x37, 0xcf, 0xef, 0xbd, 0x31, 0xc4, 0xa2, 0xca, 0xc7, 0x55, 0xad, 0x8c,
	0xa2, 0x21, 0x7e, 0xd2, 0x09, 0x04, 0x4b, 0x91, 0xd7, 0xf4, 0x3b, 0x90, 0x57, 0xb9, 0x63, 0xde,
	0xd0, 0x1b, 0xc5, 0xdc, 0x96, 0xf4, 0x07, 0x44, 0x5b, 0xb1, 0x69, 0xa4, 0x66, 0xfe, 0x90, 0x8c,
	0x62, 0xee, 0xfe, 0xd2, 0x3d, 0x81, 0x1e, 0x97, 0x6f, 0x8d, 0xd4, 0xc6, 0x62, 0x0a, 0x69, 0xd6,
	0x2a, 0x73, 0x83, 0xee, 0x8f, 0x52, 0x08, 0x2a, 0x61, 0xd6, 0xcc, 0xc7, 0x2e, 0xd6, 0x74, 0x0a,
	0xd1, 0x5a, 0x8a, 0x4c, 0xd6, 0x8c, 0x0c, 0xc9, 0x28, 0x99, 0x0e, 0x5a, 0x21, 0x63, 0xc7, 0x35,
	0x7e, 0xc6, 0xc3, 0x45, 0x69, 0xea, 0x1d, 0x77, 0x48, 0xfa, 0x0f, 0xc8, 0x8b, 0x34, 0x2c, 0xc0,
	0x81, 0x9f, 0x67, 0x03, 0x4f, 0xd2, 0xb4, 0x68, 0x8b, 0xa1, 0xff, 0x21, 0xa8, 0x94, 0x36, 0x2c,
	0x44, 0x2c, 0x3b, 0xc3, 0x2e, 0x95, 0x76, 0x60, 0x44, 0x59, 0x81, 0x2b, 0x95, 0xed, 0x58, 0xd4,
	0x0a, 0xb4, 0xb5, 0x8d, 0xa0, 0xa9, 0x37, 0xac, 0xd7, 0x46, 0xd0, 0xd4, 0x9b, 0xc1, 0x23, 0x24,
	0x27, 0xaa, 0x2e, 0x64, 0xf4, 0x17, 0x42, 0x4c, 0x05, 0x8d, 0x26, 0xd3, 0xc4, 0xdd, 0x6a, 0x13,
	0xe5, 0xed, 0xc9, 0xbd, 0x7f, 0xe7, 0x0d, 0xe6, 0xd0, 0x3f, 0x88, 0xbd, 0x9e, 0xe4, 0x01, 0xe2,
	0xce, 0xc5, 0xd5, 0x2c, 0xe9, 0xde, 0x83, 0x3e, 0x97, 0xba, 0x52, 0xa5, 0x96, 0xf4, 0x0f, 0x80,
	0x36, 0xc2, 0x34, 0x7a, 0xae, 0x32, 0x89, 0x64, 0x21, 0x3f, 0xe9, 0xd0, 0x59, 0xb7, 0x32, 0x1f,
	0x53, 0xfd, 0xd5, 0xa5, 0xda, 0x12, 0x5c, 0xdc, 0xd9, 0x21, 0x5a, 0x72, 0x8c, 0xf6, 0x56, 0x41,
	0xa6, 0xef, 0x1e, 0x84, 0x8b, 0xad, 0x2c, 0x71, 0x81, 0xa5, 0x28, 0xa4, 0xe3, 0xc0, 0x9a, 0x7e,
	0x03, 0x3f, 0xcf, 0xdc, 0x9b, 0xf3, 0xf3, 0x8c, 0xfe, 0x86, 0xd8, 0xe4, 0x85, 0xd4, 0x46, 0x14,
	0x15, 0xca, 0x21, 0xfc, 0xd8, 0xa0, 0x93, 0xce, 0x5c, 0xf0, 0xe9, 0xc9, 0x20, 0xff, 0x57, 0xce,
	0x32, 0x61, 0x04, 0x0b, 0xdb, 0x3b, 0x6d, 0x7d, 0x2b, 0x67, 0xab, 0x08, 0xdb, 0xb3, 0x8f, 0x00,
	0x00, 0x00, 0xff, 0xff, 0xfc, 0x06, 0xfe, 0xb4, 0xa2, 0x03, 0x00, 0x00,
}
