// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: osmosis/callback/v1beta1/errors.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
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
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// ModuleErrors defines the module level error codes
type ModuleErrors int32

const (
	// ERR_UNKNOWN is the default error code
	ModuleErrors_ERR_UNKNOWN ModuleErrors = 0
	// ERR_OUT_OF_GAS is the error code when the contract callback exceeds the gas
	// limit allowed by the module
	ModuleErrors_ERR_OUT_OF_GAS ModuleErrors = 1
	// ERR_CONTRACT_EXECUTION_FAILED is the error code when the contract callback
	// execution fails
	ModuleErrors_ERR_CONTRACT_EXECUTION_FAILED ModuleErrors = 2
)

var ModuleErrors_name = map[int32]string{
	0: "ERR_UNKNOWN",
	1: "ERR_OUT_OF_GAS",
	2: "ERR_CONTRACT_EXECUTION_FAILED",
}

var ModuleErrors_value = map[string]int32{
	"ERR_UNKNOWN":                   0,
	"ERR_OUT_OF_GAS":                1,
	"ERR_CONTRACT_EXECUTION_FAILED": 2,
}

func (x ModuleErrors) String() string {
	return proto.EnumName(ModuleErrors_name, int32(x))
}

func (ModuleErrors) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_b09e851f2910e9a2, []int{0}
}

func init() {
	proto.RegisterEnum("osmosis.callback.v1beta1.ModuleErrors", ModuleErrors_name, ModuleErrors_value)
}

func init() {
	proto.RegisterFile("osmosis/callback/v1beta1/errors.proto", fileDescriptor_b09e851f2910e9a2)
}

var fileDescriptor_b09e851f2910e9a2 = []byte{
	// 235 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0xcd, 0x2f, 0xce, 0xcd,
	0x2f, 0xce, 0x2c, 0xd6, 0x4f, 0x4e, 0xcc, 0xc9, 0x49, 0x4a, 0x4c, 0xce, 0xd6, 0x2f, 0x33, 0x4c,
	0x4a, 0x2d, 0x49, 0x34, 0xd4, 0x4f, 0x2d, 0x2a, 0xca, 0x2f, 0x2a, 0xd6, 0x2b, 0x28, 0xca, 0x2f,
	0xc9, 0x17, 0x92, 0x80, 0x2a, 0xd3, 0x83, 0x29, 0xd3, 0x83, 0x2a, 0x93, 0x12, 0x49, 0xcf, 0x4f,
	0xcf, 0x07, 0x2b, 0xd2, 0x07, 0xb1, 0x20, 0xea, 0xb5, 0xc2, 0xb8, 0x78, 0x7c, 0xf3, 0x53, 0x4a,
	0x73, 0x52, 0x5d, 0xc1, 0xa6, 0x08, 0xf1, 0x73, 0x71, 0xbb, 0x06, 0x05, 0xc5, 0x87, 0xfa, 0x79,
	0xfb, 0xf9, 0x87, 0xfb, 0x09, 0x30, 0x08, 0x09, 0x71, 0xf1, 0x81, 0x04, 0xfc, 0x43, 0x43, 0xe2,
	0xfd, 0xdd, 0xe2, 0xdd, 0x1d, 0x83, 0x05, 0x18, 0x85, 0x14, 0xb9, 0x64, 0x41, 0x62, 0xce, 0xfe,
	0x7e, 0x21, 0x41, 0x8e, 0xce, 0x21, 0xf1, 0xae, 0x11, 0xae, 0xce, 0xa1, 0x21, 0x9e, 0xfe, 0x7e,
	0xf1, 0x6e, 0x8e, 0x9e, 0x3e, 0xae, 0x2e, 0x02, 0x4c, 0x4e, 0x7e, 0x27, 0x1e, 0xc9, 0x31, 0x5e,
	0x78, 0x24, 0xc7, 0xf8, 0xe0, 0x91, 0x1c, 0xe3, 0x84, 0xc7, 0x72, 0x0c, 0x17, 0x1e, 0xcb, 0x31,
	0xdc, 0x78, 0x2c, 0xc7, 0x10, 0x65, 0x92, 0x9e, 0x59, 0x92, 0x51, 0x9a, 0xa4, 0x97, 0x9c, 0x9f,
	0xab, 0x0f, 0x75, 0xac, 0x6e, 0x4e, 0x62, 0x52, 0x31, 0x8c, 0xa3, 0x5f, 0x66, 0x64, 0xa6, 0x5f,
	0x81, 0xf0, 0x66, 0x49, 0x65, 0x41, 0x6a, 0x71, 0x12, 0x1b, 0xd8, 0xb9, 0xc6, 0x80, 0x00, 0x00,
	0x00, 0xff, 0xff, 0xcd, 0x6c, 0x5b, 0xb4, 0x07, 0x01, 0x00, 0x00,
}
