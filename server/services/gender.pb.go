// Code generated by protoc-gen-go. DO NOT EDIT.
// source: gender.proto

package services

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

type Gender int32

const (
	Gender_UNKNOWN Gender = 0
	Gender_MALE    Gender = 1
	Gender_FEMALE  Gender = 2
)

var Gender_name = map[int32]string{
	0: "UNKNOWN",
	1: "MALE",
	2: "FEMALE",
}

var Gender_value = map[string]int32{
	"UNKNOWN": 0,
	"MALE":    1,
	"FEMALE":  2,
}

func (x Gender) String() string {
	return proto.EnumName(Gender_name, int32(x))
}

func (Gender) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_9b0394242b5170f7, []int{0}
}

func init() {
	proto.RegisterEnum("services.Gender", Gender_name, Gender_value)
}

func init() {
	proto.RegisterFile("gender.proto", fileDescriptor_9b0394242b5170f7)
}

var fileDescriptor_9b0394242b5170f7 = []byte{
	// 100 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x49, 0x4f, 0xcd, 0x4b,
	0x49, 0x2d, 0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x28, 0x4e, 0x2d, 0x2a, 0xcb, 0x4c,
	0x4e, 0x2d, 0xd6, 0xd2, 0xe6, 0x62, 0x73, 0x07, 0xcb, 0x08, 0x71, 0x73, 0xb1, 0x87, 0xfa, 0x79,
	0xfb, 0xf9, 0x87, 0xfb, 0x09, 0x30, 0x08, 0x71, 0x70, 0xb1, 0xf8, 0x3a, 0xfa, 0xb8, 0x0a, 0x30,
	0x0a, 0x71, 0x71, 0xb1, 0xb9, 0xb9, 0x82, 0xd9, 0x4c, 0x4e, 0x3c, 0x51, 0x5c, 0x7a, 0xfa, 0x30,
	0xad, 0x49, 0x6c, 0x60, 0xb3, 0x8c, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0x8d, 0x4b, 0x45, 0xd8,
	0x5b, 0x00, 0x00, 0x00,
}