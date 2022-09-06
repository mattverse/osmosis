// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: osmosis/mint/v1beta1/mint.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-sdk/codec/types"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/gogo/protobuf/types"
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

// Minter represents the minting state.
type Minter struct {
	// epoch_provisions represent rewards for the current epoch.
	EpochProvisions github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,1,opt,name=epoch_provisions,json=epochProvisions,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"epoch_provisions" yaml:"epoch_provisions"`
}

func (m *Minter) Reset()         { *m = Minter{} }
func (m *Minter) String() string { return proto.CompactTextString(m) }
func (*Minter) ProtoMessage()    {}
func (*Minter) Descriptor() ([]byte, []int) {
	return fileDescriptor_ccb38f8335e0f45b, []int{0}
}
func (m *Minter) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Minter) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Minter.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Minter) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Minter.Merge(m, src)
}
func (m *Minter) XXX_Size() int {
	return m.Size()
}
func (m *Minter) XXX_DiscardUnknown() {
	xxx_messageInfo_Minter.DiscardUnknown(m)
}

var xxx_messageInfo_Minter proto.InternalMessageInfo

// WeightedAddress represents an address with a weight assigned to it.
// The weight is used to determine the proportion of the total minted
// tokens to be minted to the address.
type WeightedAddress struct {
	Address string                                 `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty" yaml:"address"`
	Weight  github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,2,opt,name=weight,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"weight" yaml:"weight"`
}

func (m *WeightedAddress) Reset()         { *m = WeightedAddress{} }
func (m *WeightedAddress) String() string { return proto.CompactTextString(m) }
func (*WeightedAddress) ProtoMessage()    {}
func (*WeightedAddress) Descriptor() ([]byte, []int) {
	return fileDescriptor_ccb38f8335e0f45b, []int{1}
}
func (m *WeightedAddress) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *WeightedAddress) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_WeightedAddress.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *WeightedAddress) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WeightedAddress.Merge(m, src)
}
func (m *WeightedAddress) XXX_Size() int {
	return m.Size()
}
func (m *WeightedAddress) XXX_DiscardUnknown() {
	xxx_messageInfo_WeightedAddress.DiscardUnknown(m)
}

var xxx_messageInfo_WeightedAddress proto.InternalMessageInfo

func (m *WeightedAddress) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

// DistributionProportions defines the distribution proportions of the minted
// denom. In other words, defines which stakeholders will receive the minted
// denoms and how much.
type DistributionProportions struct {
	// staking defines the proportion of the minted mint_denom that is to be
	// allocated as staking rewards.
	Staking github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,1,opt,name=staking,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"staking" yaml:"staking"`
	// pool_incentives defines the proportion of the minted mint_denom that is
	// to be allocated as pool incentives.
	PoolIncentives github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,2,opt,name=pool_incentives,json=poolIncentives,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"pool_incentives" yaml:"pool_incentives"`
	// developer_rewards defines the proportion of the minted mint_denom that is
	// to be allocated to developer rewards address.
	DeveloperRewards github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,3,opt,name=developer_rewards,json=developerRewards,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"developer_rewards" yaml:"developer_rewards"`
	// community_pool defines the proportion of the minted mint_denom that is
	// to be allocated to the community pool.
	CommunityPool github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,4,opt,name=community_pool,json=communityPool,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"community_pool" yaml:"community_pool"`
}

func (m *DistributionProportions) Reset()         { *m = DistributionProportions{} }
func (m *DistributionProportions) String() string { return proto.CompactTextString(m) }
func (*DistributionProportions) ProtoMessage()    {}
func (*DistributionProportions) Descriptor() ([]byte, []int) {
	return fileDescriptor_ccb38f8335e0f45b, []int{2}
}
func (m *DistributionProportions) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *DistributionProportions) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_DistributionProportions.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *DistributionProportions) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DistributionProportions.Merge(m, src)
}
func (m *DistributionProportions) XXX_Size() int {
	return m.Size()
}
func (m *DistributionProportions) XXX_DiscardUnknown() {
	xxx_messageInfo_DistributionProportions.DiscardUnknown(m)
}

var xxx_messageInfo_DistributionProportions proto.InternalMessageInfo

// Params holds parameters for the x/mint module.
type Params struct {
	// mint_denom is the denom of the coin to mint.
	MintDenom string `protobuf:"bytes,1,opt,name=mint_denom,json=mintDenom,proto3" json:"mint_denom,omitempty"`
	// genesis_epoch_provisions epoch provisions from the first epoch.
	GenesisEpochProvisions github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,2,opt,name=genesis_epoch_provisions,json=genesisEpochProvisions,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"genesis_epoch_provisions" yaml:"genesis_epoch_provisions"`
	// epoch_identifier mint epoch identifier e.g. (day, week).
	EpochIdentifier string `protobuf:"bytes,3,opt,name=epoch_identifier,json=epochIdentifier,proto3" json:"epoch_identifier,omitempty" yaml:"epoch_identifier"`
	// reduction_period_in_epochs the number of epochs it takes
	// to reduce the rewards.
	ReductionPeriodInEpochs int64 `protobuf:"varint,4,opt,name=reduction_period_in_epochs,json=reductionPeriodInEpochs,proto3" json:"reduction_period_in_epochs,omitempty" yaml:"reduction_period_in_epochs"`
	// reduction_factor is the reduction multiplier to execute
	// at the end of each period set by reduction_period_in_epochs.
	ReductionFactor github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,5,opt,name=reduction_factor,json=reductionFactor,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"reduction_factor" yaml:"reduction_factor"`
	// distribution_proportions defines the distribution proportions of the minted
	// denom. In other words, defines which stakeholders will receive the minted
	// denoms and how much.
	DistributionProportions DistributionProportions `protobuf:"bytes,6,opt,name=distribution_proportions,json=distributionProportions,proto3" json:"distribution_proportions"`
	// weighted_developer_rewards_receivers is the address to receive developer
	// rewards with weights assignedt to each address. The final amount that each
	// address receives is: epoch_provisions *
	// distribution_proportions.developer_rewards * Address's Weight.
	WeightedDeveloperRewardsReceivers []WeightedAddress `protobuf:"bytes,7,rep,name=weighted_developer_rewards_receivers,json=weightedDeveloperRewardsReceivers,proto3" json:"weighted_developer_rewards_receivers" yaml:"developer_rewards_receiver"`
	// minting_rewards_distribution_start_epoch start epoch to distribute minting
	// rewards
	MintingRewardsDistributionStartEpoch int64 `protobuf:"varint,8,opt,name=minting_rewards_distribution_start_epoch,json=mintingRewardsDistributionStartEpoch,proto3" json:"minting_rewards_distribution_start_epoch,omitempty" yaml:"minting_rewards_distribution_start_epoch"`
}

func (m *Params) Reset()      { *m = Params{} }
func (*Params) ProtoMessage() {}
func (*Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_ccb38f8335e0f45b, []int{3}
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

func (m *Params) GetMintDenom() string {
	if m != nil {
		return m.MintDenom
	}
	return ""
}

func (m *Params) GetEpochIdentifier() string {
	if m != nil {
		return m.EpochIdentifier
	}
	return ""
}

func (m *Params) GetReductionPeriodInEpochs() int64 {
	if m != nil {
		return m.ReductionPeriodInEpochs
	}
	return 0
}

func (m *Params) GetDistributionProportions() DistributionProportions {
	if m != nil {
		return m.DistributionProportions
	}
	return DistributionProportions{}
}

func (m *Params) GetWeightedDeveloperRewardsReceivers() []WeightedAddress {
	if m != nil {
		return m.WeightedDeveloperRewardsReceivers
	}
	return nil
}

func (m *Params) GetMintingRewardsDistributionStartEpoch() int64 {
	if m != nil {
		return m.MintingRewardsDistributionStartEpoch
	}
	return 0
}

func init() {
	proto.RegisterType((*Minter)(nil), "osmosis.mint.v1beta1.Minter")
	proto.RegisterType((*WeightedAddress)(nil), "osmosis.mint.v1beta1.WeightedAddress")
	proto.RegisterType((*DistributionProportions)(nil), "osmosis.mint.v1beta1.DistributionProportions")
	proto.RegisterType((*Params)(nil), "osmosis.mint.v1beta1.Params")
}

func init() { proto.RegisterFile("osmosis/mint/v1beta1/mint.proto", fileDescriptor_ccb38f8335e0f45b) }

var fileDescriptor_ccb38f8335e0f45b = []byte{
	// 767 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x55, 0xcd, 0x6e, 0xf3, 0x44,
	0x14, 0x8d, 0x49, 0x49, 0xf9, 0xe6, 0x53, 0x9b, 0x62, 0x95, 0xc6, 0x14, 0x11, 0xb7, 0x56, 0x8b,
	0x82, 0x44, 0x63, 0xd2, 0xee, 0xba, 0x01, 0xa2, 0x50, 0x48, 0x25, 0xa4, 0x30, 0x2c, 0x2a, 0x75,
	0x63, 0x39, 0xf6, 0xd4, 0x1d, 0x35, 0x9e, 0x31, 0x33, 0x93, 0x84, 0x6c, 0x78, 0x01, 0x36, 0x2c,
	0xbb, 0x84, 0xb7, 0xe9, 0xb2, 0xec, 0x10, 0x8b, 0x08, 0xb5, 0x6f, 0x90, 0x17, 0x00, 0xcd, 0x8f,
	0x93, 0xd6, 0x6d, 0x24, 0x22, 0x56, 0xf1, 0x9c, 0x7b, 0x72, 0xce, 0xf5, 0x9d, 0x7b, 0xaf, 0x81,
	0x4b, 0x79, 0x4a, 0x39, 0xe6, 0x7e, 0x8a, 0x89, 0xf0, 0x47, 0xad, 0x3e, 0x12, 0x61, 0x4b, 0x1d,
	0x9a, 0x19, 0xa3, 0x82, 0xda, 0xdb, 0x86, 0xd0, 0x54, 0x98, 0x21, 0xec, 0x6e, 0x27, 0x34, 0xa1,
	0x8a, 0xe0, 0xcb, 0x27, 0xcd, 0xdd, 0x75, 0x13, 0x4a, 0x93, 0x01, 0xf2, 0xd5, 0xa9, 0x3f, 0xbc,
	0xf2, 0x05, 0x4e, 0x11, 0x17, 0x61, 0x9a, 0x19, 0xc2, 0x87, 0x45, 0x42, 0x48, 0x26, 0x26, 0x54,
	0x2f, 0x86, 0xe2, 0x21, 0x0b, 0x05, 0xa6, 0x44, 0xc7, 0xbd, 0x9f, 0x41, 0xe5, 0x3b, 0x4c, 0x04,
	0x62, 0xb6, 0x00, 0x5b, 0x28, 0xa3, 0xd1, 0x75, 0x90, 0x31, 0x3a, 0xc2, 0x1c, 0x53, 0xc2, 0x1d,
	0x6b, 0xcf, 0x6a, 0xbc, 0x69, 0x77, 0xef, 0xa6, 0x6e, 0xe9, 0xaf, 0xa9, 0xfb, 0x49, 0x82, 0xc5,
	0xf5, 0xb0, 0xdf, 0x8c, 0x68, 0xea, 0x47, 0x2a, 0x7f, 0xf3, 0x73, 0xc4, 0xe3, 0x1b, 0x5f, 0x4c,
	0x32, 0xc4, 0x9b, 0x1d, 0x14, 0xcd, 0xa6, 0x6e, 0x6d, 0x12, 0xa6, 0x83, 0x53, 0xaf, 0xa8, 0xe7,
	0xc1, 0xaa, 0x82, 0x7a, 0x0b, 0xe4, 0xd6, 0x02, 0xd5, 0x0b, 0x84, 0x93, 0x6b, 0x81, 0xe2, 0xaf,
	0xe2, 0x98, 0x21, 0xce, 0xed, 0xcf, 0xc0, 0x7a, 0xa8, 0x1f, 0x4d, 0x02, 0xf6, 0x6c, 0xea, 0x6e,
	0x6a, 0x49, 0x13, 0xf0, 0x60, 0x4e, 0xb1, 0x2f, 0x40, 0x65, 0xac, 0x04, 0x9c, 0x77, 0x14, 0xf9,
	0x8b, 0x95, 0xb3, 0xdd, 0xd0, 0xd2, 0x5a, 0xc5, 0x83, 0x46, 0xce, 0xfb, 0xa3, 0x0c, 0x6a, 0x1d,
	0xcc, 0x05, 0xc3, 0xfd, 0xa1, 0xac, 0x58, 0x8f, 0xd1, 0x8c, 0x32, 0xf9, 0xc4, 0xed, 0x4b, 0xb0,
	0xce, 0x45, 0x78, 0x83, 0x49, 0x62, 0x52, 0xfc, 0x72, 0x65, 0x57, 0xf3, 0x42, 0x46, 0xc6, 0x83,
	0xb9, 0xa0, 0xfd, 0x23, 0xa8, 0x66, 0x94, 0x0e, 0x02, 0x4c, 0x22, 0x44, 0x04, 0x1e, 0x21, 0x6e,
	0xde, 0xec, 0xdb, 0x95, 0x3d, 0x76, 0xb4, 0x47, 0x41, 0xce, 0x83, 0x9b, 0x12, 0xe9, 0xce, 0x01,
	0x7b, 0x0c, 0xde, 0x8f, 0xd1, 0x08, 0x0d, 0x68, 0x86, 0x58, 0xc0, 0xd0, 0x38, 0x64, 0x31, 0x77,
	0xca, 0xca, 0xf4, 0x7c, 0x65, 0x53, 0x47, 0x9b, 0xbe, 0x10, 0xf4, 0xe0, 0xd6, 0x1c, 0x83, 0x1a,
	0xb2, 0x09, 0xd8, 0x8c, 0x68, 0x9a, 0x0e, 0x09, 0x16, 0x93, 0x40, 0x26, 0xe5, 0xac, 0x29, 0xd7,
	0x6f, 0x56, 0x76, 0xfd, 0x40, 0xbb, 0x3e, 0x57, 0xf3, 0xe0, 0xc6, 0x1c, 0xe8, 0xc9, 0xf3, 0x3f,
	0x15, 0x50, 0xe9, 0x85, 0x2c, 0x4c, 0xb9, 0xfd, 0x31, 0x00, 0x72, 0xf6, 0x82, 0x18, 0x11, 0x9a,
	0xea, 0x5b, 0x84, 0x6f, 0x24, 0xd2, 0x91, 0x80, 0xfd, 0x8b, 0x05, 0x9c, 0x04, 0x11, 0xc4, 0x31,
	0x0f, 0x5e, 0xcc, 0x85, 0xbe, 0x8f, 0xef, 0x57, 0x4e, 0xd2, 0xd5, 0x49, 0x2e, 0xd3, 0xf5, 0xe0,
	0x8e, 0x09, 0x7d, 0xfd, 0x7c, 0x4c, 0xec, 0xb3, 0x7c, 0x38, 0x71, 0x2c, 0xef, 0xec, 0x0a, 0x23,
	0x66, 0xee, 0xe7, 0xa3, 0xe2, 0xb8, 0x2d, 0x18, 0xf9, 0xb8, 0x75, 0xe7, 0x88, 0xdd, 0x07, 0xbb,
	0x0c, 0xc5, 0xc3, 0x48, 0x76, 0x71, 0x90, 0x21, 0x86, 0x69, 0x1c, 0x60, 0xa2, 0x13, 0xe1, 0xaa,
	0xf6, 0xe5, 0xf6, 0xe1, 0x6c, 0xea, 0xee, 0x6b, 0xc5, 0xe5, 0x5c, 0x0f, 0xd6, 0xe6, 0xc1, 0x9e,
	0x8a, 0x75, 0x89, 0x4a, 0x9a, 0xcb, 0x45, 0xb2, 0xf8, 0xdf, 0x55, 0x18, 0x09, 0xca, 0x9c, 0x77,
	0xff, 0xdf, 0x22, 0x29, 0xea, 0x79, 0xb0, 0x3a, 0x87, 0xce, 0x14, 0x62, 0x13, 0xe0, 0xc4, 0x4f,
	0x86, 0x55, 0x56, 0x35, 0x9f, 0x56, 0xa7, 0xb2, 0x67, 0x35, 0xde, 0x1e, 0x1f, 0x35, 0x5f, 0xdb,
	0xb9, 0xcd, 0x25, 0x23, 0xde, 0x5e, 0x93, 0xc9, 0xc2, 0x5a, 0xbc, 0x64, 0x03, 0xfc, 0x6e, 0x81,
	0x83, 0xb1, 0x59, 0x5c, 0xc1, 0x8b, 0x5e, 0x0f, 0x18, 0x8a, 0x10, 0x1e, 0x21, 0xc6, 0x9d, 0xf5,
	0xbd, 0x72, 0xe3, 0xed, 0xf1, 0xe1, 0xeb, 0xe6, 0x85, 0xd5, 0xd7, 0xfe, 0x54, 0x9a, 0x2e, 0xea,
	0xbf, 0x5c, 0xd7, 0x83, 0xfb, 0xb9, 0x7b, 0xa7, 0x30, 0x54, 0x30, 0xb7, 0x96, 0x3d, 0xdc, 0x90,
	0x76, 0x98, 0x24, 0x73, 0x81, 0x67, 0x45, 0xe2, 0x22, 0x64, 0x42, 0xdf, 0xa8, 0xf3, 0x9e, 0xba,
	0xfc, 0x93, 0xd9, 0xd4, 0xf5, 0xb5, 0xf9, 0x7f, 0xfd, 0xa7, 0x07, 0x0f, 0x0c, 0xd5, 0x24, 0xf0,
	0xb4, 0xa2, 0x3f, 0x48, 0x9e, 0x6a, 0x8c, 0xd3, 0xb5, 0xdb, 0xdf, 0xdc, 0x52, 0xfb, 0xfc, 0xee,
	0xa1, 0x6e, 0xdd, 0x3f, 0xd4, 0xad, 0xbf, 0x1f, 0xea, 0xd6, 0xaf, 0x8f, 0xf5, 0xd2, 0xfd, 0x63,
	0xbd, 0xf4, 0xe7, 0x63, 0xbd, 0x74, 0xf9, 0xf9, 0x93, 0xae, 0x30, 0xc5, 0x3a, 0x1a, 0x84, 0x7d,
	0x9e, 0x1f, 0xfc, 0x51, 0xab, 0xe5, 0xff, 0xa4, 0xbf, 0xa8, 0xaa, 0x47, 0xfa, 0x15, 0xf5, 0x0d,
	0x3b, 0xf9, 0x37, 0x00, 0x00, 0xff, 0xff, 0x6e, 0xf8, 0x2b, 0xb5, 0x6e, 0x07, 0x00, 0x00,
}

func (m *Minter) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Minter) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Minter) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.EpochProvisions.Size()
		i -= size
		if _, err := m.EpochProvisions.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintMint(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *WeightedAddress) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *WeightedAddress) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *WeightedAddress) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.Weight.Size()
		i -= size
		if _, err := m.Weight.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintMint(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintMint(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *DistributionProportions) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *DistributionProportions) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *DistributionProportions) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.CommunityPool.Size()
		i -= size
		if _, err := m.CommunityPool.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintMint(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x22
	{
		size := m.DeveloperRewards.Size()
		i -= size
		if _, err := m.DeveloperRewards.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintMint(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	{
		size := m.PoolIncentives.Size()
		i -= size
		if _, err := m.PoolIncentives.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintMint(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	{
		size := m.Staking.Size()
		i -= size
		if _, err := m.Staking.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintMint(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
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
	if m.MintingRewardsDistributionStartEpoch != 0 {
		i = encodeVarintMint(dAtA, i, uint64(m.MintingRewardsDistributionStartEpoch))
		i--
		dAtA[i] = 0x40
	}
	if len(m.WeightedDeveloperRewardsReceivers) > 0 {
		for iNdEx := len(m.WeightedDeveloperRewardsReceivers) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.WeightedDeveloperRewardsReceivers[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintMint(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x3a
		}
	}
	{
		size, err := m.DistributionProportions.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintMint(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x32
	{
		size := m.ReductionFactor.Size()
		i -= size
		if _, err := m.ReductionFactor.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintMint(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x2a
	if m.ReductionPeriodInEpochs != 0 {
		i = encodeVarintMint(dAtA, i, uint64(m.ReductionPeriodInEpochs))
		i--
		dAtA[i] = 0x20
	}
	if len(m.EpochIdentifier) > 0 {
		i -= len(m.EpochIdentifier)
		copy(dAtA[i:], m.EpochIdentifier)
		i = encodeVarintMint(dAtA, i, uint64(len(m.EpochIdentifier)))
		i--
		dAtA[i] = 0x1a
	}
	{
		size := m.GenesisEpochProvisions.Size()
		i -= size
		if _, err := m.GenesisEpochProvisions.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintMint(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if len(m.MintDenom) > 0 {
		i -= len(m.MintDenom)
		copy(dAtA[i:], m.MintDenom)
		i = encodeVarintMint(dAtA, i, uint64(len(m.MintDenom)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintMint(dAtA []byte, offset int, v uint64) int {
	offset -= sovMint(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Minter) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.EpochProvisions.Size()
	n += 1 + l + sovMint(uint64(l))
	return n
}

func (m *WeightedAddress) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovMint(uint64(l))
	}
	l = m.Weight.Size()
	n += 1 + l + sovMint(uint64(l))
	return n
}

func (m *DistributionProportions) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Staking.Size()
	n += 1 + l + sovMint(uint64(l))
	l = m.PoolIncentives.Size()
	n += 1 + l + sovMint(uint64(l))
	l = m.DeveloperRewards.Size()
	n += 1 + l + sovMint(uint64(l))
	l = m.CommunityPool.Size()
	n += 1 + l + sovMint(uint64(l))
	return n
}

func (m *Params) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.MintDenom)
	if l > 0 {
		n += 1 + l + sovMint(uint64(l))
	}
	l = m.GenesisEpochProvisions.Size()
	n += 1 + l + sovMint(uint64(l))
	l = len(m.EpochIdentifier)
	if l > 0 {
		n += 1 + l + sovMint(uint64(l))
	}
	if m.ReductionPeriodInEpochs != 0 {
		n += 1 + sovMint(uint64(m.ReductionPeriodInEpochs))
	}
	l = m.ReductionFactor.Size()
	n += 1 + l + sovMint(uint64(l))
	l = m.DistributionProportions.Size()
	n += 1 + l + sovMint(uint64(l))
	if len(m.WeightedDeveloperRewardsReceivers) > 0 {
		for _, e := range m.WeightedDeveloperRewardsReceivers {
			l = e.Size()
			n += 1 + l + sovMint(uint64(l))
		}
	}
	if m.MintingRewardsDistributionStartEpoch != 0 {
		n += 1 + sovMint(uint64(m.MintingRewardsDistributionStartEpoch))
	}
	return n
}

func sovMint(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozMint(x uint64) (n int) {
	return sovMint(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Minter) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMint
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
			return fmt.Errorf("proto: Minter: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Minter: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field EpochProvisions", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMint
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
				return ErrInvalidLengthMint
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMint
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.EpochProvisions.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMint(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthMint
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
func (m *WeightedAddress) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMint
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
			return fmt.Errorf("proto: WeightedAddress: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: WeightedAddress: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMint
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
				return ErrInvalidLengthMint
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMint
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Weight", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMint
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
				return ErrInvalidLengthMint
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMint
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Weight.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMint(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthMint
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
func (m *DistributionProportions) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMint
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
			return fmt.Errorf("proto: DistributionProportions: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: DistributionProportions: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Staking", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMint
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
				return ErrInvalidLengthMint
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMint
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Staking.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PoolIncentives", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMint
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
				return ErrInvalidLengthMint
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMint
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.PoolIncentives.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DeveloperRewards", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMint
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
				return ErrInvalidLengthMint
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMint
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.DeveloperRewards.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CommunityPool", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMint
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
				return ErrInvalidLengthMint
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMint
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.CommunityPool.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMint(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthMint
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
				return ErrIntOverflowMint
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
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MintDenom", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMint
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
				return ErrInvalidLengthMint
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMint
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.MintDenom = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field GenesisEpochProvisions", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMint
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
				return ErrInvalidLengthMint
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMint
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.GenesisEpochProvisions.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field EpochIdentifier", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMint
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
				return ErrInvalidLengthMint
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMint
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.EpochIdentifier = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ReductionPeriodInEpochs", wireType)
			}
			m.ReductionPeriodInEpochs = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMint
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ReductionPeriodInEpochs |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ReductionFactor", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMint
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
				return ErrInvalidLengthMint
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMint
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.ReductionFactor.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DistributionProportions", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMint
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
				return ErrInvalidLengthMint
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthMint
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.DistributionProportions.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field WeightedDeveloperRewardsReceivers", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMint
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
				return ErrInvalidLengthMint
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthMint
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.WeightedDeveloperRewardsReceivers = append(m.WeightedDeveloperRewardsReceivers, WeightedAddress{})
			if err := m.WeightedDeveloperRewardsReceivers[len(m.WeightedDeveloperRewardsReceivers)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 8:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MintingRewardsDistributionStartEpoch", wireType)
			}
			m.MintingRewardsDistributionStartEpoch = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMint
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MintingRewardsDistributionStartEpoch |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipMint(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthMint
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
func skipMint(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowMint
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
					return 0, ErrIntOverflowMint
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
					return 0, ErrIntOverflowMint
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
				return 0, ErrInvalidLengthMint
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupMint
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthMint
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthMint        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowMint          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupMint = fmt.Errorf("proto: unexpected end of group")
)
