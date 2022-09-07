// Code generated by protoc-gen-go. DO NOT EDIT.
// source: processor_message.proto

package pb

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

type CPU struct {
	Brand                string   `protobuf:"bytes,1,opt,name=brand,proto3" json:"brand,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	NumberCores          uint32   `protobuf:"varint,3,opt,name=number_cores,json=numberCores,proto3" json:"number_cores,omitempty"`
	NumberThreads        uint32   `protobuf:"varint,4,opt,name=number_threads,json=numberThreads,proto3" json:"number_threads,omitempty"`
	MinGhz               float64  `protobuf:"fixed64,5,opt,name=min_ghz,json=minGhz,proto3" json:"min_ghz,omitempty"`
	MaxGhz               float64  `protobuf:"fixed64,6,opt,name=max_ghz,json=maxGhz,proto3" json:"max_ghz,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CPU) Reset()         { *m = CPU{} }
func (m *CPU) String() string { return proto.CompactTextString(m) }
func (*CPU) ProtoMessage()    {}
func (*CPU) Descriptor() ([]byte, []int) {
	return fileDescriptor_466578cecc6db379, []int{0}
}

func (m *CPU) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CPU.Unmarshal(m, b)
}
func (m *CPU) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CPU.Marshal(b, m, deterministic)
}
func (m *CPU) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CPU.Merge(m, src)
}
func (m *CPU) XXX_Size() int {
	return xxx_messageInfo_CPU.Size(m)
}
func (m *CPU) XXX_DiscardUnknown() {
	xxx_messageInfo_CPU.DiscardUnknown(m)
}

var xxx_messageInfo_CPU proto.InternalMessageInfo

func (m *CPU) GetBrand() string {
	if m != nil {
		return m.Brand
	}
	return ""
}

func (m *CPU) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *CPU) GetNumberCores() uint32 {
	if m != nil {
		return m.NumberCores
	}
	return 0
}

func (m *CPU) GetNumberThreads() uint32 {
	if m != nil {
		return m.NumberThreads
	}
	return 0
}

func (m *CPU) GetMinGhz() float64 {
	if m != nil {
		return m.MinGhz
	}
	return 0
}

func (m *CPU) GetMaxGhz() float64 {
	if m != nil {
		return m.MaxGhz
	}
	return 0
}

type GPU struct {
	Brand                string   `protobuf:"bytes,1,opt,name=brand,proto3" json:"brand,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	MinGhz               float64  `protobuf:"fixed64,3,opt,name=min_ghz,json=minGhz,proto3" json:"min_ghz,omitempty"`
	MaxGhz               float64  `protobuf:"fixed64,4,opt,name=max_ghz,json=maxGhz,proto3" json:"max_ghz,omitempty"`
	Memory               *Memory  `protobuf:"bytes,5,opt,name=memory,proto3" json:"memory,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GPU) Reset()         { *m = GPU{} }
func (m *GPU) String() string { return proto.CompactTextString(m) }
func (*GPU) ProtoMessage()    {}
func (*GPU) Descriptor() ([]byte, []int) {
	return fileDescriptor_466578cecc6db379, []int{1}
}

func (m *GPU) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GPU.Unmarshal(m, b)
}
func (m *GPU) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GPU.Marshal(b, m, deterministic)
}
func (m *GPU) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GPU.Merge(m, src)
}
func (m *GPU) XXX_Size() int {
	return xxx_messageInfo_GPU.Size(m)
}
func (m *GPU) XXX_DiscardUnknown() {
	xxx_messageInfo_GPU.DiscardUnknown(m)
}

var xxx_messageInfo_GPU proto.InternalMessageInfo

func (m *GPU) GetBrand() string {
	if m != nil {
		return m.Brand
	}
	return ""
}

func (m *GPU) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *GPU) GetMinGhz() float64 {
	if m != nil {
		return m.MinGhz
	}
	return 0
}

func (m *GPU) GetMaxGhz() float64 {
	if m != nil {
		return m.MaxGhz
	}
	return 0
}

func (m *GPU) GetMemory() *Memory {
	if m != nil {
		return m.Memory
	}
	return nil
}

func init() {
	proto.RegisterType((*CPU)(nil), "sample.pcbook.CPU")
	proto.RegisterType((*GPU)(nil), "sample.pcbook.GPU")
}

func init() {
	proto.RegisterFile("processor_message.proto", fileDescriptor_466578cecc6db379)
}

var fileDescriptor_466578cecc6db379 = []byte{
	// 253 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x90, 0xcd, 0x4a, 0xc3, 0x40,
	0x14, 0x85, 0x99, 0x26, 0x8d, 0x78, 0x6b, 0x5c, 0x0c, 0x95, 0x0e, 0xae, 0x62, 0x41, 0xc8, 0xc6,
	0x2c, 0xf4, 0x0d, 0xec, 0x22, 0x2b, 0x41, 0x82, 0x6e, 0xdc, 0x84, 0x99, 0xf4, 0xd2, 0x14, 0x9d,
	0x1f, 0x66, 0x2a, 0xd4, 0x3c, 0x85, 0xef, 0xe1, 0x4b, 0x8a, 0x77, 0xb2, 0x68, 0x17, 0x2e, 0xba,
	0x9b, 0xf9, 0xce, 0x07, 0xe7, 0x72, 0x60, 0xe1, 0xbc, 0xed, 0x30, 0x04, 0xeb, 0x5b, 0x8d, 0x21,
	0xc8, 0x0d, 0x56, 0xce, 0xdb, 0x9d, 0xe5, 0x79, 0x90, 0xda, 0x7d, 0x60, 0xe5, 0x3a, 0x65, 0xed,
	0xfb, 0xf5, 0x5c, 0xa3, 0xb6, 0xfe, 0xeb, 0x58, 0x5a, 0xfe, 0x30, 0x48, 0x56, 0xcf, 0xaf, 0x7c,
	0x0e, 0x53, 0xe5, 0xa5, 0x59, 0x0b, 0x56, 0xb0, 0xf2, 0xbc, 0x89, 0x1f, 0xce, 0x21, 0x35, 0x52,
	0xa3, 0x98, 0x10, 0xa4, 0x37, 0xbf, 0x81, 0x0b, 0xf3, 0xa9, 0x15, 0xfa, 0xb6, 0xb3, 0x1e, 0x83,
	0x48, 0x0a, 0x56, 0xe6, 0xcd, 0x2c, 0xb2, 0xd5, 0x1f, 0xe2, 0xb7, 0x70, 0x39, 0x2a, 0xbb, 0xde,
	0xa3, 0x5c, 0x07, 0x91, 0x92, 0x94, 0x47, 0xfa, 0x12, 0x21, 0x5f, 0xc0, 0x99, 0xde, 0x9a, 0x76,
	0xd3, 0x0f, 0x62, 0x5a, 0xb0, 0x92, 0x35, 0x99, 0xde, 0x9a, 0xba, 0x1f, 0x28, 0x90, 0x7b, 0x0a,
	0xb2, 0x31, 0x90, 0xfb, 0xba, 0x1f, 0x96, 0xdf, 0x0c, 0x92, 0xfa, 0xa4, 0x6b, 0x0f, 0x3a, 0x92,
	0xff, 0x3a, 0xd2, 0xc3, 0x0e, 0x7e, 0x07, 0x59, 0x5c, 0x8a, 0x8e, 0x9a, 0xdd, 0x5f, 0x55, 0x47,
	0x3b, 0x56, 0x4f, 0x14, 0x36, 0xa3, 0xf4, 0x98, 0xbe, 0x4d, 0x9c, 0x52, 0x19, 0xad, 0xf9, 0xf0,
	0x1b, 0x00, 0x00, 0xff, 0xff, 0xf9, 0x20, 0x9c, 0xbf, 0x8d, 0x01, 0x00, 0x00,
}
