// Code generated by protoc-gen-go.
// source: api.proto
// DO NOT EDIT!

/*
Package api is a generated protocol buffer package.

It is generated from these files:
	api.proto

It has these top-level messages:
	ListingReqApi
	ListingRespApi
	Inventory
	OrderRespApi
	TransactionRecord
*/
package pb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
const _ = proto.ProtoPackageIsVersion1

type ListingReqApi struct {
	Listing     *Listing     `protobuf:"bytes,1,opt,name=listing" json:"listing,omitempty"`
	Inventory   []*Inventory `protobuf:"bytes,2,rep,name=inventory" json:"inventory,omitempty"`
	CurrentSlug string       `protobuf:"bytes,3,opt,name=currentSlug" json:"currentSlug,omitempty"`
}

func (m *ListingReqApi) Reset()                    { *m = ListingReqApi{} }
func (m *ListingReqApi) String() string            { return proto.CompactTextString(m) }
func (*ListingReqApi) ProtoMessage()               {}
func (*ListingReqApi) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{0} }

func (m *ListingReqApi) GetListing() *Listing {
	if m != nil {
		return m.Listing
	}
	return nil
}

func (m *ListingReqApi) GetInventory() []*Inventory {
	if m != nil {
		return m.Inventory
	}
	return nil
}

type ListingRespApi struct {
	Contract  *RicardianContract `protobuf:"bytes,1,opt,name=contract" json:"contract,omitempty"`
	Inventory []*Inventory       `protobuf:"bytes,2,rep,name=inventory" json:"inventory,omitempty"`
}

func (m *ListingRespApi) Reset()                    { *m = ListingRespApi{} }
func (m *ListingRespApi) String() string            { return proto.CompactTextString(m) }
func (*ListingRespApi) ProtoMessage()               {}
func (*ListingRespApi) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{1} }

func (m *ListingRespApi) GetContract() *RicardianContract {
	if m != nil {
		return m.Contract
	}
	return nil
}

func (m *ListingRespApi) GetInventory() []*Inventory {
	if m != nil {
		return m.Inventory
	}
	return nil
}

type Inventory struct {
	Item  string `protobuf:"bytes,1,opt,name=item" json:"item,omitempty"`
	Count uint64 `protobuf:"varint,2,opt,name=count" json:"count,omitempty"`
}

func (m *Inventory) Reset()                    { *m = Inventory{} }
func (m *Inventory) String() string            { return proto.CompactTextString(m) }
func (*Inventory) ProtoMessage()               {}
func (*Inventory) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{2} }

type OrderRespApi struct {
	Contract     *RicardianContract   `protobuf:"bytes,1,opt,name=contract" json:"contract,omitempty"`
	State        OrderState           `protobuf:"varint,2,opt,name=state,enum=OrderState" json:"state,omitempty"`
	Read         bool                 `protobuf:"varint,3,opt,name=read" json:"read,omitempty"`
	Funded       bool                 `protobuf:"varint,4,opt,name=funded" json:"funded,omitempty"`
	Transactions []*TransactionRecord `protobuf:"bytes,5,rep,name=transactions" json:"transactions,omitempty"`
}

func (m *OrderRespApi) Reset()                    { *m = OrderRespApi{} }
func (m *OrderRespApi) String() string            { return proto.CompactTextString(m) }
func (*OrderRespApi) ProtoMessage()               {}
func (*OrderRespApi) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{3} }

func (m *OrderRespApi) GetContract() *RicardianContract {
	if m != nil {
		return m.Contract
	}
	return nil
}

func (m *OrderRespApi) GetTransactions() []*TransactionRecord {
	if m != nil {
		return m.Transactions
	}
	return nil
}

type TransactionRecord struct {
	Txid  string `protobuf:"bytes,1,opt,name=txid" json:"txid,omitempty"`
	Value int64  `protobuf:"varint,2,opt,name=value" json:"value,omitempty"`
}

func (m *TransactionRecord) Reset()                    { *m = TransactionRecord{} }
func (m *TransactionRecord) String() string            { return proto.CompactTextString(m) }
func (*TransactionRecord) ProtoMessage()               {}
func (*TransactionRecord) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{4} }

func init() {
	proto.RegisterType((*ListingReqApi)(nil), "ListingReqApi")
	proto.RegisterType((*ListingRespApi)(nil), "ListingRespApi")
	proto.RegisterType((*Inventory)(nil), "Inventory")
	proto.RegisterType((*OrderRespApi)(nil), "OrderRespApi")
	proto.RegisterType((*TransactionRecord)(nil), "TransactionRecord")
}

var fileDescriptor5 = []byte{
	// 322 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x9c, 0x52, 0x3d, 0x4f, 0xc3, 0x30,
	0x14, 0x54, 0xe8, 0x07, 0xcd, 0x6b, 0x29, 0xc2, 0x42, 0x28, 0x62, 0x2a, 0x99, 0x3a, 0x79, 0x28,
	0x82, 0x8d, 0x01, 0x31, 0x21, 0x21, 0x21, 0xbd, 0xf2, 0x07, 0x8c, 0x6d, 0x2a, 0xa3, 0x62, 0x07,
	0xdb, 0xa9, 0x60, 0xe2, 0xd7, 0xf1, 0xbf, 0xb0, 0x9d, 0xa4, 0x05, 0xb1, 0x20, 0xb6, 0x77, 0xf7,
	0x4e, 0xf7, 0x2e, 0x17, 0x43, 0xce, 0x2a, 0x45, 0x2b, 0x6b, 0xbc, 0x39, 0x3d, 0xe4, 0x46, 0x7b,
	0xcb, 0xb8, 0x77, 0x2d, 0x31, 0x31, 0x56, 0x48, 0xdb, 0xa2, 0xf2, 0x03, 0x0e, 0xee, 0x94, 0xf3,
	0x4a, 0xaf, 0x50, 0xbe, 0x5e, 0x57, 0x8a, 0x94, 0xb0, 0xbf, 0x6e, 0x88, 0x22, 0x9b, 0x65, 0xf3,
	0xf1, 0x62, 0x44, 0x3b, 0x41, 0xb7, 0x20, 0x73, 0xc8, 0x95, 0xde, 0x48, 0xed, 0x8d, 0x7d, 0x2f,
	0xf6, 0x66, 0xbd, 0xa0, 0x02, 0x7a, 0xdb, 0x31, 0xb8, 0x5b, 0x92, 0x19, 0x8c, 0x79, 0x6d, 0x6d,
	0x40, 0xcb, 0x75, 0xbd, 0x2a, 0x7a, 0xc1, 0x31, 0xc7, 0xef, 0x54, 0xf9, 0x0c, 0xd3, 0x6d, 0x00,
	0x57, 0xc5, 0x04, 0x14, 0x46, 0x5d, 0xe6, 0x36, 0x02, 0xa1, 0xa8, 0x38, 0xb3, 0x42, 0x31, 0x7d,
	0xd3, 0x6e, 0x70, 0xab, 0xf9, 0x7b, 0x9a, 0xf2, 0x02, 0xf2, 0x2d, 0x4f, 0x08, 0xf4, 0x95, 0x97,
	0x2f, 0xe9, 0x44, 0x8e, 0x69, 0x26, 0xc7, 0x30, 0xe0, 0xa6, 0xd6, 0x3e, 0xd8, 0x64, 0xf3, 0x3e,
	0x36, 0xa0, 0xfc, 0xcc, 0x60, 0x72, 0x1f, 0x4b, 0xfb, 0x6f, 0xc2, 0x33, 0x18, 0x38, 0xcf, 0xbc,
	0x4c, 0xb6, 0xd3, 0xc5, 0x98, 0x26, 0xb7, 0x65, 0xa4, 0xb0, 0xd9, 0xc4, 0x34, 0x56, 0x32, 0x91,
	0x1a, 0x1a, 0x61, 0x9a, 0xc9, 0x09, 0x0c, 0x9f, 0x6a, 0x2d, 0xa4, 0x28, 0xfa, 0x89, 0x6d, 0x11,
	0xb9, 0x84, 0x49, 0xf0, 0xd5, 0x2e, 0x58, 0x2b, 0xa3, 0x5d, 0x31, 0x48, 0xdf, 0x4c, 0xe8, 0xc3,
	0x8e, 0x44, 0xc9, 0xc3, 0x6f, 0xc6, 0x1f, 0xba, 0xf2, 0x0a, 0x8e, 0x7e, 0x49, 0xe2, 0x61, 0xff,
	0xa6, 0x44, 0x57, 0x43, 0x9c, 0x63, 0x0d, 0x1b, 0xb6, 0xae, 0x9b, 0xbc, 0x3d, 0x6c, 0xc0, 0xe3,
	0x30, 0xbd, 0x98, 0xf3, 0xaf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x76, 0xf2, 0x5c, 0x62, 0x5d, 0x02,
	0x00, 0x00,
}
