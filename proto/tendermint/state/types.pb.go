// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: tendermint/state/types.proto

package state

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/gogo/protobuf/types"
	github_com_gogo_protobuf_types "github.com/gogo/protobuf/types"
	_ "github.com/PikeEcosystem/tendermint/abci/types"
	_ "github.com/PikeEcosystem/tendermint/proto/tendermint/types"
	types "github.com/tendermint/tendermint/abci/types"
	state "github.com/tendermint/tendermint/proto/tendermint/state"
	types1 "github.com/tendermint/tendermint/proto/tendermint/types"
	_ "github.com/tendermint/tendermint/proto/tendermint/version"
	io "io"
	math "math"
	math_bits "math/bits"
	time "time"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = time.Kitchen

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// ABCIResponses retains the responses
// of the various ABCI calls during block processing.
// It is persisted to disk for each height before calling Commit.
type ABCIResponses struct {
	DeliverTxs []*types.ResponseDeliverTx `protobuf:"bytes,1,rep,name=deliver_txs,json=deliverTxs,proto3" json:"deliver_txs,omitempty"`
	EndBlock   *types.ResponseEndBlock    `protobuf:"bytes,2,opt,name=end_block,json=endBlock,proto3" json:"end_block,omitempty"`
	BeginBlock *types.ResponseBeginBlock  `protobuf:"bytes,3,opt,name=begin_block,json=beginBlock,proto3" json:"begin_block,omitempty"`
}

func (m *ABCIResponses) Reset()         { *m = ABCIResponses{} }
func (m *ABCIResponses) String() string { return proto.CompactTextString(m) }
func (*ABCIResponses) ProtoMessage()    {}
func (*ABCIResponses) Descriptor() ([]byte, []int) {
	return fileDescriptor_898987a4421067cd, []int{0}
}
func (m *ABCIResponses) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ABCIResponses) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ABCIResponses.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ABCIResponses) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ABCIResponses.Merge(m, src)
}
func (m *ABCIResponses) XXX_Size() int {
	return m.Size()
}
func (m *ABCIResponses) XXX_DiscardUnknown() {
	xxx_messageInfo_ABCIResponses.DiscardUnknown(m)
}

var xxx_messageInfo_ABCIResponses proto.InternalMessageInfo

func (m *ABCIResponses) GetDeliverTxs() []*types.ResponseDeliverTx {
	if m != nil {
		return m.DeliverTxs
	}
	return nil
}

func (m *ABCIResponses) GetEndBlock() *types.ResponseEndBlock {
	if m != nil {
		return m.EndBlock
	}
	return nil
}

func (m *ABCIResponses) GetBeginBlock() *types.ResponseBeginBlock {
	if m != nil {
		return m.BeginBlock
	}
	return nil
}

type State struct {
	Version state.Version `protobuf:"bytes,1,opt,name=version,proto3" json:"version"`
	// immutable
	ChainID       string `protobuf:"bytes,2,opt,name=chain_id,json=chainId,proto3" json:"chain_id,omitempty"`
	InitialHeight int64  `protobuf:"varint,14,opt,name=initial_height,json=initialHeight,proto3" json:"initial_height,omitempty"`
	// LastBlockHeight=0 at genesis (ie. block(H=0) does not exist)
	LastBlockHeight int64          `protobuf:"varint,3,opt,name=last_block_height,json=lastBlockHeight,proto3" json:"last_block_height,omitempty"`
	LastBlockID     types1.BlockID `protobuf:"bytes,4,opt,name=last_block_id,json=lastBlockId,proto3" json:"last_block_id"`
	LastBlockTime   time.Time      `protobuf:"bytes,5,opt,name=last_block_time,json=lastBlockTime,proto3,stdtime" json:"last_block_time"`
	// LastValidators is used to validate block.LastCommit.
	// Validators are persisted to the database separately every time they change,
	// so we can query for historical validator sets.
	// Note that if s.LastBlockHeight causes a valset change,
	// we set s.LastHeightValidatorsChanged = s.LastBlockHeight + 1 + 1
	// Extra +1 due to nextValSet delay.
	NextValidators              *types1.ValidatorSet `protobuf:"bytes,6,opt,name=next_validators,json=nextValidators,proto3" json:"next_validators,omitempty"`
	Validators                  *types1.ValidatorSet `protobuf:"bytes,7,opt,name=validators,proto3" json:"validators,omitempty"`
	LastValidators              *types1.ValidatorSet `protobuf:"bytes,8,opt,name=last_validators,json=lastValidators,proto3" json:"last_validators,omitempty"`
	LastHeightValidatorsChanged int64                `protobuf:"varint,9,opt,name=last_height_validators_changed,json=lastHeightValidatorsChanged,proto3" json:"last_height_validators_changed,omitempty"`
	// Consensus parameters used for validating blocks.
	// Changes returned by EndBlock and updated after Commit.
	ConsensusParams                  types1.ConsensusParams `protobuf:"bytes,10,opt,name=consensus_params,json=consensusParams,proto3" json:"consensus_params"`
	LastHeightConsensusParamsChanged int64                  `protobuf:"varint,11,opt,name=last_height_consensus_params_changed,json=lastHeightConsensusParamsChanged,proto3" json:"last_height_consensus_params_changed,omitempty"`
	// Merkle root of the results from executing prev block
	LastResultsHash []byte `protobuf:"bytes,12,opt,name=last_results_hash,json=lastResultsHash,proto3" json:"last_results_hash,omitempty"`
	// the latest AppHash we've received from calling abci.Commit()
	AppHash []byte `protobuf:"bytes,13,opt,name=app_hash,json=appHash,proto3" json:"app_hash,omitempty"`
	// the VRF Proof value generated by the last Proposer
	LastProofHash []byte `protobuf:"bytes,1000,opt,name=last_proof_hash,json=lastProofHash,proto3" json:"last_proof_hash,omitempty"`
}

func (m *State) Reset()         { *m = State{} }
func (m *State) String() string { return proto.CompactTextString(m) }
func (*State) ProtoMessage()    {}
func (*State) Descriptor() ([]byte, []int) {
	return fileDescriptor_898987a4421067cd, []int{1}
}
func (m *State) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *State) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_State.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *State) XXX_Merge(src proto.Message) {
	xxx_messageInfo_State.Merge(m, src)
}
func (m *State) XXX_Size() int {
	return m.Size()
}
func (m *State) XXX_DiscardUnknown() {
	xxx_messageInfo_State.DiscardUnknown(m)
}

var xxx_messageInfo_State proto.InternalMessageInfo

func (m *State) GetVersion() state.Version {
	if m != nil {
		return m.Version
	}
	return state.Version{}
}

func (m *State) GetChainID() string {
	if m != nil {
		return m.ChainID
	}
	return ""
}

func (m *State) GetInitialHeight() int64 {
	if m != nil {
		return m.InitialHeight
	}
	return 0
}

func (m *State) GetLastBlockHeight() int64 {
	if m != nil {
		return m.LastBlockHeight
	}
	return 0
}

func (m *State) GetLastBlockID() types1.BlockID {
	if m != nil {
		return m.LastBlockID
	}
	return types1.BlockID{}
}

func (m *State) GetLastBlockTime() time.Time {
	if m != nil {
		return m.LastBlockTime
	}
	return time.Time{}
}

func (m *State) GetNextValidators() *types1.ValidatorSet {
	if m != nil {
		return m.NextValidators
	}
	return nil
}

func (m *State) GetValidators() *types1.ValidatorSet {
	if m != nil {
		return m.Validators
	}
	return nil
}

func (m *State) GetLastValidators() *types1.ValidatorSet {
	if m != nil {
		return m.LastValidators
	}
	return nil
}

func (m *State) GetLastHeightValidatorsChanged() int64 {
	if m != nil {
		return m.LastHeightValidatorsChanged
	}
	return 0
}

func (m *State) GetConsensusParams() types1.ConsensusParams {
	if m != nil {
		return m.ConsensusParams
	}
	return types1.ConsensusParams{}
}

func (m *State) GetLastHeightConsensusParamsChanged() int64 {
	if m != nil {
		return m.LastHeightConsensusParamsChanged
	}
	return 0
}

func (m *State) GetLastResultsHash() []byte {
	if m != nil {
		return m.LastResultsHash
	}
	return nil
}

func (m *State) GetAppHash() []byte {
	if m != nil {
		return m.AppHash
	}
	return nil
}

func (m *State) GetLastProofHash() []byte {
	if m != nil {
		return m.LastProofHash
	}
	return nil
}

func init() {
	proto.RegisterType((*ABCIResponses)(nil), "tendermint.state.ABCIResponses")
	proto.RegisterType((*State)(nil), "tendermint.state.State")
}

func init() { proto.RegisterFile("tendermint/state/types.proto", fileDescriptor_898987a4421067cd) }

var fileDescriptor_898987a4421067cd = []byte{
	// 703 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x94, 0x4f, 0x6f, 0xd3, 0x30,
	0x18, 0xc6, 0x1b, 0xba, 0xad, 0x9d, 0xb3, 0xb6, 0x10, 0x38, 0x64, 0x1d, 0xa4, 0x65, 0xfc, 0xab,
	0x90, 0x48, 0xa5, 0x71, 0xe2, 0x82, 0x44, 0x5a, 0xb4, 0x55, 0x9a, 0xd0, 0x94, 0x4d, 0x3b, 0x70,
	0x89, 0x9c, 0xc4, 0x4b, 0x2c, 0xd2, 0x38, 0x8a, 0xdd, 0x69, 0x7c, 0x05, 0x4e, 0xfb, 0x58, 0x3b,
	0xee, 0x88, 0x38, 0x0c, 0xd4, 0x5d, 0xf8, 0x18, 0xc8, 0x76, 0x9c, 0x66, 0x2b, 0x93, 0x76, 0x73,
	0xde, 0xe7, 0x79, 0x7f, 0x7a, 0x6c, 0xbf, 0x31, 0xe8, 0x12, 0xca, 0x72, 0x18, 0x90, 0x74, 0x48,
	0x19, 0x64, 0x68, 0xc8, 0xbe, 0x67, 0x88, 0xda, 0x59, 0x4e, 0x18, 0x31, 0xda, 0x4a, 0xb3, 0x85,
	0xd6, 0x7d, 0x12, 0x91, 0x88, 0x08, 0x69, 0xc8, 0x57, 0xd2, 0xd5, 0xdd, 0x2c, 0x09, 0xd0, 0x0f,
	0x70, 0x15, 0xd0, 0x5d, 0xc0, 0x45, 0xf5, 0x86, 0xd6, 0x67, 0x28, 0x0d, 0x51, 0x3e, 0xc5, 0x29,
	0x2b, 0xd4, 0x53, 0x98, 0xe0, 0x10, 0x32, 0x92, 0x17, 0x8e, 0x67, 0x4b, 0x8e, 0x0c, 0xe6, 0x70,
	0xaa, 0x00, 0x4f, 0x97, 0xe4, 0x2a, 0xde, 0xaa, 0xa8, 0xa7, 0x28, 0xa7, 0x58, 0x85, 0x28, 0xf4,
	0x5e, 0x44, 0x48, 0x94, 0xa0, 0xa1, 0xf8, 0xf2, 0x67, 0x27, 0x43, 0x86, 0xa7, 0x88, 0x32, 0x38,
	0xcd, 0xfe, 0x83, 0x5f, 0x3a, 0x9a, 0xee, 0x56, 0x45, 0xbd, 0xbd, 0xed, 0xed, 0x5f, 0x1a, 0x68,
	0x7d, 0x72, 0x46, 0x13, 0x17, 0xd1, 0x8c, 0xa4, 0x14, 0x51, 0x63, 0x04, 0xf4, 0x10, 0x25, 0xf8,
	0x14, 0xe5, 0x1e, 0x3b, 0xa3, 0xa6, 0xd6, 0xaf, 0x0f, 0xf4, 0x9d, 0x6d, 0x7b, 0x01, 0xb1, 0x39,
	0xc4, 0x56, 0x0d, 0x63, 0xe9, 0x3d, 0x3a, 0x73, 0x41, 0xa8, 0x96, 0xd4, 0xf8, 0x08, 0xd6, 0x51,
	0x1a, 0x7a, 0x7e, 0x42, 0x82, 0x6f, 0xe6, 0x83, 0xbe, 0x36, 0xd0, 0x77, 0x9e, 0xdf, 0x89, 0xf8,
	0x9c, 0x86, 0x0e, 0x37, 0xba, 0x4d, 0x54, 0xac, 0x8c, 0x31, 0xd0, 0x7d, 0x14, 0xe1, 0xb4, 0x20,
	0xd4, 0x05, 0xe1, 0xc5, 0x9d, 0x04, 0x87, 0x7b, 0x25, 0x03, 0xf8, 0xe5, 0x7a, 0xfb, 0x47, 0x03,
	0xac, 0x1e, 0xf2, 0xf3, 0x30, 0x3e, 0x80, 0x46, 0x71, 0xb2, 0xa6, 0x26, 0x58, 0x9b, 0x55, 0x96,
	0x38, 0x33, 0xfb, 0x58, 0x1a, 0x9c, 0x95, 0x8b, 0xab, 0x5e, 0xcd, 0x55, 0x7e, 0xe3, 0x35, 0x68,
	0x06, 0x31, 0xc4, 0xa9, 0x87, 0x43, 0xb1, 0x93, 0x75, 0x47, 0x9f, 0x5f, 0xf5, 0x1a, 0x23, 0x5e,
	0x9b, 0x8c, 0xdd, 0x86, 0x10, 0x27, 0xa1, 0xf1, 0x0a, 0xb4, 0x71, 0x8a, 0x19, 0x86, 0x89, 0x17,
	0x23, 0x1c, 0xc5, 0xcc, 0x6c, 0xf7, 0xb5, 0x41, 0xdd, 0x6d, 0x15, 0xd5, 0x3d, 0x51, 0x34, 0xde,
	0x82, 0x47, 0x09, 0xa4, 0x4c, 0x6e, 0x4c, 0x39, 0xeb, 0xc2, 0xd9, 0xe1, 0x82, 0x48, 0x5e, 0x78,
	0x5d, 0xd0, 0xaa, 0x78, 0x71, 0x68, 0xae, 0x2c, 0x67, 0x97, 0x97, 0x29, 0xba, 0x26, 0x63, 0xe7,
	0x31, 0xcf, 0x3e, 0xbf, 0xea, 0xe9, 0xfb, 0x0a, 0x35, 0x19, 0xbb, 0x7a, 0xc9, 0x9d, 0x84, 0xc6,
	0x3e, 0xe8, 0x54, 0x98, 0x7c, 0x92, 0xcc, 0x55, 0x41, 0xed, 0xda, 0x72, 0xcc, 0x6c, 0x35, 0x66,
	0xf6, 0x91, 0x1a, 0x33, 0xa7, 0xc9, 0xb1, 0xe7, 0xbf, 0x7b, 0x9a, 0xdb, 0x2a, 0x59, 0x5c, 0x35,
	0x76, 0x41, 0x27, 0x45, 0x67, 0xcc, 0x2b, 0xff, 0x07, 0x6a, 0xae, 0x09, 0x9a, 0xb5, 0x9c, 0xf1,
	0x58, 0x79, 0x0e, 0x11, 0x73, 0xdb, 0xbc, 0xad, 0xac, 0xf0, 0x81, 0x01, 0x15, 0x46, 0xe3, 0x5e,
	0x8c, 0x4a, 0x07, 0x0f, 0x22, 0xb6, 0x55, 0x81, 0x34, 0xef, 0x17, 0x84, 0xb7, 0x55, 0x82, 0x8c,
	0x80, 0x25, 0x40, 0xf2, 0x66, 0x2a, 0x3c, 0x2f, 0x88, 0x61, 0x1a, 0xa1, 0xd0, 0x5c, 0x17, 0x97,
	0xb5, 0xc5, 0x5d, 0xf2, 0x9e, 0x16, 0xdd, 0x23, 0x69, 0x31, 0x5c, 0xf0, 0x30, 0xe0, 0x73, 0x99,
	0xd2, 0x19, 0xf5, 0xe4, 0x4b, 0x60, 0x82, 0xe5, 0xbf, 0x40, 0xc6, 0x19, 0x29, 0xe7, 0x81, 0x30,
	0x16, 0xf3, 0xd7, 0x09, 0x6e, 0x96, 0x8d, 0x2f, 0xe0, 0x65, 0x35, 0xd8, 0x6d, 0x7e, 0x19, 0x4f,
	0x17, 0xf1, 0xfa, 0x8b, 0x78, 0xb7, 0xf8, 0x2a, 0xa3, 0x1a, 0xc4, 0x1c, 0xd1, 0x59, 0xc2, 0xa8,
	0x17, 0x43, 0x1a, 0x9b, 0x1b, 0x7d, 0x6d, 0xb0, 0x21, 0x07, 0xd1, 0x95, 0xf5, 0x3d, 0x48, 0x63,
	0x63, 0x13, 0x34, 0x61, 0x96, 0x49, 0x4b, 0x4b, 0x58, 0x1a, 0x30, 0xcb, 0x84, 0xf4, 0xa6, 0x38,
	0xf8, 0x2c, 0x27, 0xe4, 0x44, 0x3a, 0xfe, 0x36, 0x84, 0x45, 0x8c, 0xca, 0x01, 0x2f, 0x73, 0xa3,
	0xb3, 0x7b, 0x31, 0xb7, 0xb4, 0xcb, 0xb9, 0xa5, 0xfd, 0x99, 0x5b, 0xda, 0xf9, 0xb5, 0x55, 0xbb,
	0xbc, 0xb6, 0x6a, 0x3f, 0xaf, 0xad, 0xda, 0xd7, 0x77, 0x11, 0x66, 0xf1, 0xcc, 0xb7, 0x03, 0x32,
	0x1d, 0x26, 0x38, 0x45, 0xc3, 0xf2, 0x29, 0x96, 0x0f, 0xf8, 0xcd, 0x67, 0xdf, 0x5f, 0x13, 0xd5,
	0xf7, 0xff, 0x02, 0x00, 0x00, 0xff, 0xff, 0xe7, 0x82, 0xd1, 0xeb, 0x0f, 0x06, 0x00, 0x00,
}

func (m *ABCIResponses) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ABCIResponses) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ABCIResponses) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.BeginBlock != nil {
		{
			size, err := m.BeginBlock.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTypes(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1a
	}
	if m.EndBlock != nil {
		{
			size, err := m.EndBlock.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTypes(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if len(m.DeliverTxs) > 0 {
		for iNdEx := len(m.DeliverTxs) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.DeliverTxs[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintTypes(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *State) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *State) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *State) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.LastProofHash) > 0 {
		i -= len(m.LastProofHash)
		copy(dAtA[i:], m.LastProofHash)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.LastProofHash)))
		i--
		dAtA[i] = 0x3e
		i--
		dAtA[i] = 0xc2
	}
	if m.InitialHeight != 0 {
		i = encodeVarintTypes(dAtA, i, uint64(m.InitialHeight))
		i--
		dAtA[i] = 0x70
	}
	if len(m.AppHash) > 0 {
		i -= len(m.AppHash)
		copy(dAtA[i:], m.AppHash)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.AppHash)))
		i--
		dAtA[i] = 0x6a
	}
	if len(m.LastResultsHash) > 0 {
		i -= len(m.LastResultsHash)
		copy(dAtA[i:], m.LastResultsHash)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.LastResultsHash)))
		i--
		dAtA[i] = 0x62
	}
	if m.LastHeightConsensusParamsChanged != 0 {
		i = encodeVarintTypes(dAtA, i, uint64(m.LastHeightConsensusParamsChanged))
		i--
		dAtA[i] = 0x58
	}
	{
		size, err := m.ConsensusParams.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintTypes(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x52
	if m.LastHeightValidatorsChanged != 0 {
		i = encodeVarintTypes(dAtA, i, uint64(m.LastHeightValidatorsChanged))
		i--
		dAtA[i] = 0x48
	}
	if m.LastValidators != nil {
		{
			size, err := m.LastValidators.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTypes(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x42
	}
	if m.Validators != nil {
		{
			size, err := m.Validators.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTypes(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x3a
	}
	if m.NextValidators != nil {
		{
			size, err := m.NextValidators.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTypes(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x32
	}
	n7, err7 := github_com_gogo_protobuf_types.StdTimeMarshalTo(m.LastBlockTime, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdTime(m.LastBlockTime):])
	if err7 != nil {
		return 0, err7
	}
	i -= n7
	i = encodeVarintTypes(dAtA, i, uint64(n7))
	i--
	dAtA[i] = 0x2a
	{
		size, err := m.LastBlockID.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintTypes(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x22
	if m.LastBlockHeight != 0 {
		i = encodeVarintTypes(dAtA, i, uint64(m.LastBlockHeight))
		i--
		dAtA[i] = 0x18
	}
	if len(m.ChainID) > 0 {
		i -= len(m.ChainID)
		copy(dAtA[i:], m.ChainID)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.ChainID)))
		i--
		dAtA[i] = 0x12
	}
	{
		size, err := m.Version.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintTypes(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintTypes(dAtA []byte, offset int, v uint64) int {
	offset -= sovTypes(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *ABCIResponses) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.DeliverTxs) > 0 {
		for _, e := range m.DeliverTxs {
			l = e.Size()
			n += 1 + l + sovTypes(uint64(l))
		}
	}
	if m.EndBlock != nil {
		l = m.EndBlock.Size()
		n += 1 + l + sovTypes(uint64(l))
	}
	if m.BeginBlock != nil {
		l = m.BeginBlock.Size()
		n += 1 + l + sovTypes(uint64(l))
	}
	return n
}

func (m *State) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Version.Size()
	n += 1 + l + sovTypes(uint64(l))
	l = len(m.ChainID)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	if m.LastBlockHeight != 0 {
		n += 1 + sovTypes(uint64(m.LastBlockHeight))
	}
	l = m.LastBlockID.Size()
	n += 1 + l + sovTypes(uint64(l))
	l = github_com_gogo_protobuf_types.SizeOfStdTime(m.LastBlockTime)
	n += 1 + l + sovTypes(uint64(l))
	if m.NextValidators != nil {
		l = m.NextValidators.Size()
		n += 1 + l + sovTypes(uint64(l))
	}
	if m.Validators != nil {
		l = m.Validators.Size()
		n += 1 + l + sovTypes(uint64(l))
	}
	if m.LastValidators != nil {
		l = m.LastValidators.Size()
		n += 1 + l + sovTypes(uint64(l))
	}
	if m.LastHeightValidatorsChanged != 0 {
		n += 1 + sovTypes(uint64(m.LastHeightValidatorsChanged))
	}
	l = m.ConsensusParams.Size()
	n += 1 + l + sovTypes(uint64(l))
	if m.LastHeightConsensusParamsChanged != 0 {
		n += 1 + sovTypes(uint64(m.LastHeightConsensusParamsChanged))
	}
	l = len(m.LastResultsHash)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	l = len(m.AppHash)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	if m.InitialHeight != 0 {
		n += 1 + sovTypes(uint64(m.InitialHeight))
	}
	l = len(m.LastProofHash)
	if l > 0 {
		n += 2 + l + sovTypes(uint64(l))
	}
	return n
}

func sovTypes(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTypes(x uint64) (n int) {
	return sovTypes(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ABCIResponses) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypes
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
			return fmt.Errorf("proto: ABCIResponses: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ABCIResponses: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DeliverTxs", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DeliverTxs = append(m.DeliverTxs, &types.ResponseDeliverTx{})
			if err := m.DeliverTxs[len(m.DeliverTxs)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field EndBlock", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.EndBlock == nil {
				m.EndBlock = &types.ResponseEndBlock{}
			}
			if err := m.EndBlock.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BeginBlock", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.BeginBlock == nil {
				m.BeginBlock = &types.ResponseBeginBlock{}
			}
			if err := m.BeginBlock.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTypes
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
func (m *State) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypes
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
			return fmt.Errorf("proto: State: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: State: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Version", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Version.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ChainID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ChainID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field LastBlockHeight", wireType)
			}
			m.LastBlockHeight = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.LastBlockHeight |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LastBlockID", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.LastBlockID.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LastBlockTime", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(&m.LastBlockTime, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field NextValidators", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.NextValidators == nil {
				m.NextValidators = &types1.ValidatorSet{}
			}
			if err := m.NextValidators.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Validators", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Validators == nil {
				m.Validators = &types1.ValidatorSet{}
			}
			if err := m.Validators.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LastValidators", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.LastValidators == nil {
				m.LastValidators = &types1.ValidatorSet{}
			}
			if err := m.LastValidators.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 9:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field LastHeightValidatorsChanged", wireType)
			}
			m.LastHeightValidatorsChanged = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.LastHeightValidatorsChanged |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 10:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ConsensusParams", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.ConsensusParams.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 11:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field LastHeightConsensusParamsChanged", wireType)
			}
			m.LastHeightConsensusParamsChanged = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.LastHeightConsensusParamsChanged |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 12:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LastResultsHash", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.LastResultsHash = append(m.LastResultsHash[:0], dAtA[iNdEx:postIndex]...)
			if m.LastResultsHash == nil {
				m.LastResultsHash = []byte{}
			}
			iNdEx = postIndex
		case 13:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AppHash", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AppHash = append(m.AppHash[:0], dAtA[iNdEx:postIndex]...)
			if m.AppHash == nil {
				m.AppHash = []byte{}
			}
			iNdEx = postIndex
		case 14:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field InitialHeight", wireType)
			}
			m.InitialHeight = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.InitialHeight |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 1000:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LastProofHash", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.LastProofHash = append(m.LastProofHash[:0], dAtA[iNdEx:postIndex]...)
			if m.LastProofHash == nil {
				m.LastProofHash = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTypes
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
func skipTypes(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTypes
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
					return 0, ErrIntOverflowTypes
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
					return 0, ErrIntOverflowTypes
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
				return 0, ErrInvalidLengthTypes
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTypes
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTypes
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTypes        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTypes          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTypes = fmt.Errorf("proto: unexpected end of group")
)
