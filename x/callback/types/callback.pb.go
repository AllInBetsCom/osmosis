// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: osmosis/callback/v1beta1/callback.proto

package types

import (
	cosmossdk_io_math "cosmossdk.io/math"
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	_ "github.com/cosmos/cosmos-sdk/codec/types"
	types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
	io "io"
	math "math"
	math_bits "math/bits"
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

// Callback defines the callback structure.
type Callback struct {
	// contract_address is the address of the contract which is requesting the
	// callback (bech32 encoded).
	ContractAddress string `protobuf:"bytes,1,opt,name=contract_address,json=contractAddress,proto3" json:"contract_address,omitempty"`
	// job_id is an identifier the callback requestor can pass in to identify the
	// callback when it happens.
	JobId uint64 `protobuf:"varint,2,opt,name=job_id,json=jobId,proto3" json:"job_id,omitempty"`
	// callback_height is the height at which the callback is executed.
	CallbackHeight int64 `protobuf:"varint,3,opt,name=callback_height,json=callbackHeight,proto3" json:"callback_height,omitempty"`
	// fee_split is the breakdown of the fees paid by the contract to reserve the
	// callback
	FeeSplit *CallbackFeesFeeSplit `protobuf:"bytes,4,opt,name=fee_split,json=feeSplit,proto3" json:"fee_split,omitempty"`
	// reserved_by is the address which reserved the callback (bech32 encoded).
	ReservedBy string `protobuf:"bytes,5,opt,name=reserved_by,json=reservedBy,proto3" json:"reserved_by,omitempty"`
	// callback_gas_limit is the maximum gas that can be consumed by this
	// callback.
	MaxGasLimit uint64 `protobuf:"varint,6,opt,name=max_gas_limit,json=maxGasLimit,proto3" json:"max_gas_limit,omitempty"`
}

func (m *Callback) Reset()         { *m = Callback{} }
func (m *Callback) String() string { return proto.CompactTextString(m) }
func (*Callback) ProtoMessage()    {}
func (*Callback) Descriptor() ([]byte, []int) {
	return fileDescriptor_1737e09bd93c8faf, []int{0}
}
func (m *Callback) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Callback) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Callback.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Callback) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Callback.Merge(m, src)
}
func (m *Callback) XXX_Size() int {
	return m.Size()
}
func (m *Callback) XXX_DiscardUnknown() {
	xxx_messageInfo_Callback.DiscardUnknown(m)
}

var xxx_messageInfo_Callback proto.InternalMessageInfo

func (m *Callback) GetContractAddress() string {
	if m != nil {
		return m.ContractAddress
	}
	return ""
}

func (m *Callback) GetJobId() uint64 {
	if m != nil {
		return m.JobId
	}
	return 0
}

func (m *Callback) GetCallbackHeight() int64 {
	if m != nil {
		return m.CallbackHeight
	}
	return 0
}

func (m *Callback) GetFeeSplit() *CallbackFeesFeeSplit {
	if m != nil {
		return m.FeeSplit
	}
	return nil
}

func (m *Callback) GetReservedBy() string {
	if m != nil {
		return m.ReservedBy
	}
	return ""
}

func (m *Callback) GetMaxGasLimit() uint64 {
	if m != nil {
		return m.MaxGasLimit
	}
	return 0
}

// CallbackFeesFeeSplit is the breakdown of all the fees that need to be paid by
// the contract to reserve a callback
type CallbackFeesFeeSplit struct {
	// transaction_fees is the transaction fees for the callback based on its gas
	// consumption
	TransactionFees *types.Coin `protobuf:"bytes,1,opt,name=transaction_fees,json=transactionFees,proto3" json:"transaction_fees,omitempty"`
	// block_reservation_fees is the block reservation fees portion of the
	// callback reservation fees
	BlockReservationFees *types.Coin `protobuf:"bytes,2,opt,name=block_reservation_fees,json=blockReservationFees,proto3" json:"block_reservation_fees,omitempty"`
	// future_reservation_fees is the future reservation fees portion of the
	// callback reservation fees
	FutureReservationFees *types.Coin `protobuf:"bytes,3,opt,name=future_reservation_fees,json=futureReservationFees,proto3" json:"future_reservation_fees,omitempty"`
	// surplus_fees is any extra fees passed in for the registration of the
	// callback
	SurplusFees *types.Coin `protobuf:"bytes,4,opt,name=surplus_fees,json=surplusFees,proto3" json:"surplus_fees,omitempty"`
}

func (m *CallbackFeesFeeSplit) Reset()         { *m = CallbackFeesFeeSplit{} }
func (m *CallbackFeesFeeSplit) String() string { return proto.CompactTextString(m) }
func (*CallbackFeesFeeSplit) ProtoMessage()    {}
func (*CallbackFeesFeeSplit) Descriptor() ([]byte, []int) {
	return fileDescriptor_1737e09bd93c8faf, []int{1}
}
func (m *CallbackFeesFeeSplit) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *CallbackFeesFeeSplit) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_CallbackFeesFeeSplit.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *CallbackFeesFeeSplit) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CallbackFeesFeeSplit.Merge(m, src)
}
func (m *CallbackFeesFeeSplit) XXX_Size() int {
	return m.Size()
}
func (m *CallbackFeesFeeSplit) XXX_DiscardUnknown() {
	xxx_messageInfo_CallbackFeesFeeSplit.DiscardUnknown(m)
}

var xxx_messageInfo_CallbackFeesFeeSplit proto.InternalMessageInfo

func (m *CallbackFeesFeeSplit) GetTransactionFees() *types.Coin {
	if m != nil {
		return m.TransactionFees
	}
	return nil
}

func (m *CallbackFeesFeeSplit) GetBlockReservationFees() *types.Coin {
	if m != nil {
		return m.BlockReservationFees
	}
	return nil
}

func (m *CallbackFeesFeeSplit) GetFutureReservationFees() *types.Coin {
	if m != nil {
		return m.FutureReservationFees
	}
	return nil
}

func (m *CallbackFeesFeeSplit) GetSurplusFees() *types.Coin {
	if m != nil {
		return m.SurplusFees
	}
	return nil
}

// Params defines the module parameters.
type Params struct {
	// callback_gas_limit is the maximum gas that can be consumed by a callback.
	CallbackGasLimit uint64 `protobuf:"varint,1,opt,name=callback_gas_limit,json=callbackGasLimit,proto3" json:"callback_gas_limit,omitempty"`
	// max_block_reservation_limit is the maximum number of callbacks which can be
	// registered in a given block.
	MaxBlockReservationLimit uint64 `protobuf:"varint,2,opt,name=max_block_reservation_limit,json=maxBlockReservationLimit,proto3" json:"max_block_reservation_limit,omitempty"`
	// max_future_reservation_limit is the maximum number of blocks in the future
	// that a contract can request a callback in.
	MaxFutureReservationLimit uint64 `protobuf:"varint,3,opt,name=max_future_reservation_limit,json=maxFutureReservationLimit,proto3" json:"max_future_reservation_limit,omitempty"`
	// block_reservation_fee_multiplier is used to calculate a part of the
	// reservation fees which will need to be paid when requesting the callback.
	BlockReservationFeeMultiplier cosmossdk_io_math.LegacyDec `protobuf:"bytes,4,opt,name=block_reservation_fee_multiplier,json=blockReservationFeeMultiplier,proto3,customtype=cosmossdk.io/math.LegacyDec" json:"block_reservation_fee_multiplier"`
	// future_reservation_fee_multiplier is used to calculate a part of the
	// reservation fees which will need to be paid while requesting the callback.
	FutureReservationFeeMultiplier cosmossdk_io_math.LegacyDec `protobuf:"bytes,5,opt,name=future_reservation_fee_multiplier,json=futureReservationFeeMultiplier,proto3,customtype=cosmossdk.io/math.LegacyDec" json:"future_reservation_fee_multiplier"`
	// min_price_of_gas defines the minimum price for each single unit of gas in
	// the network.
	MinPriceOfGas types.Coin `protobuf:"bytes,6,opt,name=min_price_of_gas,json=minPriceOfGas,proto3" json:"min_price_of_gas"`
}

func (m *Params) Reset()         { *m = Params{} }
func (m *Params) String() string { return proto.CompactTextString(m) }
func (*Params) ProtoMessage()    {}
func (*Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_1737e09bd93c8faf, []int{2}
}
func (m *Params) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Params) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Params.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Params) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Params.Merge(m, src)
}
func (m *Params) XXX_Size() int {
	return m.Size()
}
func (m *Params) XXX_DiscardUnknown() {
	xxx_messageInfo_Params.DiscardUnknown(m)
}

var xxx_messageInfo_Params proto.InternalMessageInfo

func (m *Params) GetCallbackGasLimit() uint64 {
	if m != nil {
		return m.CallbackGasLimit
	}
	return 0
}

func (m *Params) GetMaxBlockReservationLimit() uint64 {
	if m != nil {
		return m.MaxBlockReservationLimit
	}
	return 0
}

func (m *Params) GetMaxFutureReservationLimit() uint64 {
	if m != nil {
		return m.MaxFutureReservationLimit
	}
	return 0
}

func (m *Params) GetMinPriceOfGas() types.Coin {
	if m != nil {
		return m.MinPriceOfGas
	}
	return types.Coin{}
}

func init() {
	proto.RegisterType((*Callback)(nil), "osmosis.callback.v1beta1.Callback")
	proto.RegisterType((*CallbackFeesFeeSplit)(nil), "osmosis.callback.v1beta1.CallbackFeesFeeSplit")
	proto.RegisterType((*Params)(nil), "osmosis.callback.v1beta1.Params")
}

func init() {
	proto.RegisterFile("osmosis/callback/v1beta1/callback.proto", fileDescriptor_1737e09bd93c8faf)
}

var fileDescriptor_1737e09bd93c8faf = []byte{
	// 654 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x54, 0xc1, 0x6f, 0xd3, 0x3e,
	0x14, 0x6e, 0xda, 0xae, 0xda, 0xdc, 0xdf, 0x7e, 0xab, 0xac, 0x0d, 0xd2, 0x0d, 0xb2, 0xd2, 0xcb,
	0x8a, 0xc4, 0x12, 0x6d, 0x20, 0x4e, 0x20, 0x44, 0x99, 0xb6, 0x21, 0x06, 0x1b, 0xe1, 0xc6, 0xc5,
	0x72, 0x52, 0x27, 0xf5, 0x16, 0xc7, 0x55, 0xec, 0x4c, 0x2d, 0xe2, 0xce, 0x95, 0x3f, 0x86, 0x33,
	0xe7, 0x1d, 0x27, 0x4e, 0x88, 0xc3, 0x84, 0xb6, 0xff, 0x82, 0x13, 0x8a, 0x9d, 0xb4, 0xd3, 0x56,
	0x98, 0xc4, 0xcd, 0xfe, 0xde, 0xfb, 0xbe, 0xf7, 0xfc, 0x3d, 0xdb, 0x60, 0x8d, 0x0b, 0xc6, 0x05,
	0x15, 0x8e, 0x8f, 0xa3, 0xc8, 0xc3, 0xfe, 0x91, 0x73, 0xbc, 0xe1, 0x11, 0x89, 0x37, 0xc6, 0x80,
	0x3d, 0x48, 0xb8, 0xe4, 0xd0, 0xcc, 0x13, 0xed, 0x31, 0x9e, 0x27, 0x2e, 0x37, 0x43, 0xce, 0xc3,
	0x88, 0x38, 0x2a, 0xcf, 0x4b, 0x03, 0x07, 0xc7, 0x23, 0x4d, 0x5a, 0x5e, 0x0c, 0x79, 0xc8, 0xd5,
	0xd2, 0xc9, 0x56, 0x39, 0x6a, 0xf9, 0x4a, 0xcb, 0xf1, 0xb0, 0x20, 0x93, 0x72, 0x9c, 0xc6, 0x79,
	0xbc, 0xa9, 0xe3, 0x48, 0x13, 0xf5, 0x46, 0x87, 0xda, 0x9f, 0xca, 0x60, 0xf6, 0x45, 0xde, 0x00,
	0xbc, 0x0f, 0x1a, 0x3e, 0x8f, 0x65, 0x82, 0x7d, 0x89, 0x70, 0xaf, 0x97, 0x10, 0x21, 0x4c, 0xa3,
	0x65, 0x74, 0xe6, 0xdc, 0x85, 0x02, 0x7f, 0xae, 0x61, 0xb8, 0x04, 0x6a, 0x87, 0xdc, 0x43, 0xb4,
	0x67, 0x96, 0x5b, 0x46, 0xa7, 0xea, 0xce, 0x1c, 0x72, 0xef, 0x65, 0x0f, 0xae, 0x81, 0x85, 0xe2,
	0x38, 0xa8, 0x4f, 0x68, 0xd8, 0x97, 0x66, 0xa5, 0x65, 0x74, 0x2a, 0xee, 0xff, 0x05, 0xbc, 0xab,
	0x50, 0xf8, 0x0a, 0xcc, 0x05, 0x84, 0x20, 0x31, 0x88, 0xa8, 0x34, 0xab, 0x2d, 0xa3, 0x53, 0xdf,
	0xb4, 0xed, 0x3f, 0x39, 0x62, 0x17, 0x1d, 0x6e, 0x13, 0x22, 0xb6, 0x09, 0x79, 0x97, 0xb1, 0xdc,
	0xd9, 0x20, 0x5f, 0xc1, 0x55, 0x50, 0x4f, 0x88, 0x20, 0xc9, 0x31, 0xe9, 0x21, 0x6f, 0x64, 0xce,
	0xa8, 0x96, 0x41, 0x01, 0x75, 0x47, 0xb0, 0x0d, 0xe6, 0x19, 0x1e, 0xa2, 0x10, 0x0b, 0x14, 0x51,
	0x46, 0xa5, 0x59, 0x53, 0x4d, 0xd7, 0x19, 0x1e, 0xee, 0x60, 0xb1, 0x97, 0x41, 0xed, 0xaf, 0x65,
	0xb0, 0x38, 0xad, 0x0e, 0xdc, 0x02, 0x0d, 0x99, 0xe0, 0x58, 0x60, 0x5f, 0x52, 0x1e, 0xa3, 0x80,
	0x10, 0xed, 0x4a, 0x7d, 0xb3, 0x69, 0xe7, 0x5e, 0x66, 0xc6, 0x4f, 0x9a, 0xe5, 0x34, 0x76, 0x17,
	0x2e, 0x51, 0x32, 0x35, 0xb8, 0x0f, 0x6e, 0x79, 0x11, 0xf7, 0x8f, 0x90, 0x6e, 0x0b, 0x4f, 0xb4,
	0xca, 0x37, 0x69, 0x2d, 0x2a, 0xa2, 0x3b, 0xe1, 0x29, 0xc1, 0xb7, 0xe0, 0x76, 0x90, 0xca, 0x34,
	0x21, 0xd7, 0x15, 0x2b, 0x37, 0x29, 0x2e, 0x69, 0xe6, 0x55, 0xc9, 0x27, 0xe0, 0x3f, 0x91, 0x26,
	0x83, 0x28, 0x15, 0x5a, 0xa7, 0x7a, 0x93, 0x4e, 0x3d, 0x4f, 0xcf, 0xd8, 0xed, 0x5f, 0x15, 0x50,
	0x3b, 0xc0, 0x09, 0x66, 0x02, 0x3e, 0x00, 0x70, 0x7c, 0x0d, 0x26, 0xa6, 0x1b, 0xca, 0xf4, 0x46,
	0x11, 0x29, 0x9c, 0x87, 0x4f, 0xc1, 0x4a, 0x36, 0x9d, 0xeb, 0xf6, 0x68, 0x9a, 0xbe, 0x60, 0x26,
	0xc3, 0xc3, 0xee, 0x15, 0x1f, 0x34, 0xfd, 0x19, 0xb8, 0x93, 0xd1, 0xa7, 0x98, 0xa1, 0xf9, 0x15,
	0xc5, 0x6f, 0x32, 0x3c, 0xdc, 0xbe, 0x7a, 0x6a, 0x2d, 0xf0, 0x01, 0xb4, 0xa6, 0x8e, 0x06, 0xb1,
	0x34, 0x92, 0x74, 0x10, 0x51, 0x92, 0x28, 0x2b, 0xe6, 0xba, 0x1b, 0x27, 0x67, 0xab, 0xa5, 0x1f,
	0x67, 0xab, 0x2b, 0xda, 0x11, 0xd1, 0x3b, 0xb2, 0x29, 0x77, 0x18, 0x96, 0x7d, 0x7b, 0x8f, 0x84,
	0xd8, 0x1f, 0x6d, 0x11, 0xff, 0xdb, 0x97, 0x75, 0x90, 0x1b, 0xb6, 0x45, 0x7c, 0xf7, 0xee, 0x94,
	0xe1, 0xbd, 0x1e, 0xeb, 0xc2, 0x8f, 0xe0, 0xde, 0xf4, 0x29, 0x5e, 0x2e, 0x3e, 0xf3, 0xaf, 0xc5,
	0xad, 0x69, 0x73, 0xbe, 0x54, 0x7d, 0x17, 0x34, 0x18, 0x8d, 0xd1, 0x20, 0xa1, 0x3e, 0x41, 0x3c,
	0xc8, 0x66, 0xa5, 0x9e, 0xc6, 0xdf, 0x86, 0xde, 0xad, 0x66, 0x7d, 0xb8, 0xf3, 0x8c, 0xc6, 0x07,
	0x19, 0x6f, 0x3f, 0xd8, 0xc1, 0xa2, 0xfb, 0xe6, 0xe4, 0xdc, 0x32, 0x4e, 0xcf, 0x2d, 0xe3, 0xe7,
	0xb9, 0x65, 0x7c, 0xbe, 0xb0, 0x4a, 0xa7, 0x17, 0x56, 0xe9, 0xfb, 0x85, 0x55, 0x7a, 0xff, 0x28,
	0xa4, 0xb2, 0x9f, 0x7a, 0xb6, 0xcf, 0x99, 0x93, 0x3f, 0xf0, 0xf5, 0x08, 0x7b, 0xa2, 0xd8, 0x38,
	0xc7, 0x9b, 0x8f, 0x9d, 0xe1, 0xe4, 0xbb, 0x94, 0xa3, 0x01, 0x11, 0x5e, 0x4d, 0x7d, 0x4f, 0x0f,
	0x7f, 0x07, 0x00, 0x00, 0xff, 0xff, 0x04, 0xa7, 0x5e, 0x82, 0x4f, 0x05, 0x00, 0x00,
}

func (m *Callback) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Callback) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Callback) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.MaxGasLimit != 0 {
		i = encodeVarintCallback(dAtA, i, uint64(m.MaxGasLimit))
		i--
		dAtA[i] = 0x30
	}
	if len(m.ReservedBy) > 0 {
		i -= len(m.ReservedBy)
		copy(dAtA[i:], m.ReservedBy)
		i = encodeVarintCallback(dAtA, i, uint64(len(m.ReservedBy)))
		i--
		dAtA[i] = 0x2a
	}
	if m.FeeSplit != nil {
		{
			size, err := m.FeeSplit.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintCallback(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x22
	}
	if m.CallbackHeight != 0 {
		i = encodeVarintCallback(dAtA, i, uint64(m.CallbackHeight))
		i--
		dAtA[i] = 0x18
	}
	if m.JobId != 0 {
		i = encodeVarintCallback(dAtA, i, uint64(m.JobId))
		i--
		dAtA[i] = 0x10
	}
	if len(m.ContractAddress) > 0 {
		i -= len(m.ContractAddress)
		copy(dAtA[i:], m.ContractAddress)
		i = encodeVarintCallback(dAtA, i, uint64(len(m.ContractAddress)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *CallbackFeesFeeSplit) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CallbackFeesFeeSplit) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *CallbackFeesFeeSplit) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.SurplusFees != nil {
		{
			size, err := m.SurplusFees.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintCallback(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x22
	}
	if m.FutureReservationFees != nil {
		{
			size, err := m.FutureReservationFees.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintCallback(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1a
	}
	if m.BlockReservationFees != nil {
		{
			size, err := m.BlockReservationFees.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintCallback(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if m.TransactionFees != nil {
		{
			size, err := m.TransactionFees.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintCallback(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *Params) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Params) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Params) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.MinPriceOfGas.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintCallback(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x32
	{
		size := m.FutureReservationFeeMultiplier.Size()
		i -= size
		if _, err := m.FutureReservationFeeMultiplier.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintCallback(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x2a
	{
		size := m.BlockReservationFeeMultiplier.Size()
		i -= size
		if _, err := m.BlockReservationFeeMultiplier.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintCallback(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x22
	if m.MaxFutureReservationLimit != 0 {
		i = encodeVarintCallback(dAtA, i, uint64(m.MaxFutureReservationLimit))
		i--
		dAtA[i] = 0x18
	}
	if m.MaxBlockReservationLimit != 0 {
		i = encodeVarintCallback(dAtA, i, uint64(m.MaxBlockReservationLimit))
		i--
		dAtA[i] = 0x10
	}
	if m.CallbackGasLimit != 0 {
		i = encodeVarintCallback(dAtA, i, uint64(m.CallbackGasLimit))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintCallback(dAtA []byte, offset int, v uint64) int {
	offset -= sovCallback(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Callback) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ContractAddress)
	if l > 0 {
		n += 1 + l + sovCallback(uint64(l))
	}
	if m.JobId != 0 {
		n += 1 + sovCallback(uint64(m.JobId))
	}
	if m.CallbackHeight != 0 {
		n += 1 + sovCallback(uint64(m.CallbackHeight))
	}
	if m.FeeSplit != nil {
		l = m.FeeSplit.Size()
		n += 1 + l + sovCallback(uint64(l))
	}
	l = len(m.ReservedBy)
	if l > 0 {
		n += 1 + l + sovCallback(uint64(l))
	}
	if m.MaxGasLimit != 0 {
		n += 1 + sovCallback(uint64(m.MaxGasLimit))
	}
	return n
}

func (m *CallbackFeesFeeSplit) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.TransactionFees != nil {
		l = m.TransactionFees.Size()
		n += 1 + l + sovCallback(uint64(l))
	}
	if m.BlockReservationFees != nil {
		l = m.BlockReservationFees.Size()
		n += 1 + l + sovCallback(uint64(l))
	}
	if m.FutureReservationFees != nil {
		l = m.FutureReservationFees.Size()
		n += 1 + l + sovCallback(uint64(l))
	}
	if m.SurplusFees != nil {
		l = m.SurplusFees.Size()
		n += 1 + l + sovCallback(uint64(l))
	}
	return n
}

func (m *Params) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.CallbackGasLimit != 0 {
		n += 1 + sovCallback(uint64(m.CallbackGasLimit))
	}
	if m.MaxBlockReservationLimit != 0 {
		n += 1 + sovCallback(uint64(m.MaxBlockReservationLimit))
	}
	if m.MaxFutureReservationLimit != 0 {
		n += 1 + sovCallback(uint64(m.MaxFutureReservationLimit))
	}
	l = m.BlockReservationFeeMultiplier.Size()
	n += 1 + l + sovCallback(uint64(l))
	l = m.FutureReservationFeeMultiplier.Size()
	n += 1 + l + sovCallback(uint64(l))
	l = m.MinPriceOfGas.Size()
	n += 1 + l + sovCallback(uint64(l))
	return n
}

func sovCallback(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozCallback(x uint64) (n int) {
	return sovCallback(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Callback) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCallback
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Callback: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Callback: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ContractAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCallback
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthCallback
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCallback
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ContractAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field JobId", wireType)
			}
			m.JobId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCallback
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.JobId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CallbackHeight", wireType)
			}
			m.CallbackHeight = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCallback
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.CallbackHeight |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FeeSplit", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCallback
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthCallback
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthCallback
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.FeeSplit == nil {
				m.FeeSplit = &CallbackFeesFeeSplit{}
			}
			if err := m.FeeSplit.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ReservedBy", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCallback
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthCallback
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCallback
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ReservedBy = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxGasLimit", wireType)
			}
			m.MaxGasLimit = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCallback
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MaxGasLimit |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipCallback(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthCallback
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *CallbackFeesFeeSplit) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCallback
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: CallbackFeesFeeSplit: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CallbackFeesFeeSplit: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TransactionFees", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCallback
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthCallback
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthCallback
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.TransactionFees == nil {
				m.TransactionFees = &types.Coin{}
			}
			if err := m.TransactionFees.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BlockReservationFees", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCallback
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthCallback
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthCallback
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.BlockReservationFees == nil {
				m.BlockReservationFees = &types.Coin{}
			}
			if err := m.BlockReservationFees.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FutureReservationFees", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCallback
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthCallback
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthCallback
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.FutureReservationFees == nil {
				m.FutureReservationFees = &types.Coin{}
			}
			if err := m.FutureReservationFees.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SurplusFees", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCallback
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthCallback
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthCallback
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.SurplusFees == nil {
				m.SurplusFees = &types.Coin{}
			}
			if err := m.SurplusFees.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipCallback(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthCallback
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Params) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCallback
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Params: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Params: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CallbackGasLimit", wireType)
			}
			m.CallbackGasLimit = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCallback
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.CallbackGasLimit |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxBlockReservationLimit", wireType)
			}
			m.MaxBlockReservationLimit = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCallback
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MaxBlockReservationLimit |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxFutureReservationLimit", wireType)
			}
			m.MaxFutureReservationLimit = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCallback
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MaxFutureReservationLimit |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BlockReservationFeeMultiplier", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCallback
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthCallback
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCallback
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.BlockReservationFeeMultiplier.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FutureReservationFeeMultiplier", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCallback
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthCallback
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCallback
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.FutureReservationFeeMultiplier.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MinPriceOfGas", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCallback
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthCallback
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthCallback
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.MinPriceOfGas.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipCallback(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthCallback
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipCallback(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowCallback
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowCallback
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowCallback
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthCallback
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupCallback
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthCallback
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthCallback        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowCallback          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupCallback = fmt.Errorf("proto: unexpected end of group")
)
