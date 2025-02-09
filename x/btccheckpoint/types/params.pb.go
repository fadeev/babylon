// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: babylon/btccheckpoint/v1/params.proto

package types

import (
	fmt "fmt"
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

// Params defines the parameters for the module.
type Params struct {
	// btc_confirmation_depth is the confirmation depth in BTC.
	// A block is considered irreversible only when it is at least k-deep in BTC
	// (k in research paper)
	BtcConfirmationDepth uint64 `protobuf:"varint,1,opt,name=btc_confirmation_depth,json=btcConfirmationDepth,proto3" json:"btc_confirmation_depth,omitempty" yaml:"btc_confirmation_depth"`
	// checkpoint_finalization_timeout is the maximum time window (measured in BTC
	// blocks) between a checkpoint
	// - being submitted to BTC, and
	// - being reported back to BBN
	// If a checkpoint has not been reported back within w BTC blocks, then BBN
	// has dishonest majority and is stalling checkpoints (w in research paper)
	CheckpointFinalizationTimeout uint64 `protobuf:"varint,2,opt,name=checkpoint_finalization_timeout,json=checkpointFinalizationTimeout,proto3" json:"checkpoint_finalization_timeout,omitempty" yaml:"checkpoint_finalization_timeout"`
}

func (m *Params) Reset()         { *m = Params{} }
func (m *Params) String() string { return proto.CompactTextString(m) }
func (*Params) ProtoMessage()    {}
func (*Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_5445a19005ae983c, []int{0}
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

func (m *Params) GetBtcConfirmationDepth() uint64 {
	if m != nil {
		return m.BtcConfirmationDepth
	}
	return 0
}

func (m *Params) GetCheckpointFinalizationTimeout() uint64 {
	if m != nil {
		return m.CheckpointFinalizationTimeout
	}
	return 0
}

func init() {
	proto.RegisterType((*Params)(nil), "babylon.btccheckpoint.v1.Params")
}

func init() {
	proto.RegisterFile("babylon/btccheckpoint/v1/params.proto", fileDescriptor_5445a19005ae983c)
}

var fileDescriptor_5445a19005ae983c = []byte{
	// 271 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x4d, 0x4a, 0x4c, 0xaa,
	0xcc, 0xc9, 0xcf, 0xd3, 0x4f, 0x2a, 0x49, 0x4e, 0xce, 0x48, 0x4d, 0xce, 0x2e, 0xc8, 0xcf, 0xcc,
	0x2b, 0xd1, 0x2f, 0x33, 0xd4, 0x2f, 0x48, 0x2c, 0x4a, 0xcc, 0x2d, 0xd6, 0x2b, 0x28, 0xca, 0x2f,
	0xc9, 0x17, 0x92, 0x80, 0x2a, 0xd3, 0x43, 0x51, 0xa6, 0x57, 0x66, 0x28, 0x25, 0x92, 0x9e, 0x9f,
	0x9e, 0x0f, 0x56, 0xa4, 0x0f, 0x62, 0x41, 0xd4, 0x2b, 0xdd, 0x66, 0xe4, 0x62, 0x0b, 0x00, 0x1b,
	0x20, 0x14, 0xce, 0x25, 0x96, 0x54, 0x92, 0x1c, 0x9f, 0x9c, 0x9f, 0x97, 0x96, 0x59, 0x94, 0x9b,
	0x58, 0x92, 0x99, 0x9f, 0x17, 0x9f, 0x92, 0x5a, 0x50, 0x92, 0x21, 0xc1, 0xa8, 0xc0, 0xa8, 0xc1,
	0xe2, 0xa4, 0xf8, 0xe9, 0x9e, 0xbc, 0x6c, 0x65, 0x62, 0x6e, 0x8e, 0x95, 0x12, 0x76, 0x75, 0x4a,
	0x41, 0x22, 0x49, 0x25, 0xc9, 0xce, 0x48, 0xe2, 0x2e, 0x20, 0x61, 0xa1, 0x22, 0x2e, 0x79, 0x84,
	0x53, 0xe2, 0xd3, 0x32, 0xf3, 0x12, 0x73, 0x32, 0xab, 0x20, 0xfa, 0x4a, 0x32, 0x73, 0x53, 0xf3,
	0x4b, 0x4b, 0x24, 0x98, 0xc0, 0x36, 0x68, 0x7d, 0xba, 0x27, 0xaf, 0x06, 0xb1, 0x81, 0x80, 0x06,
	0xa5, 0x20, 0x59, 0x84, 0x0a, 0x37, 0x24, 0x05, 0x21, 0x10, 0x79, 0x2b, 0x96, 0x17, 0x0b, 0xe4,
	0x19, 0x9d, 0xfc, 0x4f, 0x3c, 0x92, 0x63, 0xbc, 0xf0, 0x48, 0x8e, 0xf1, 0xc1, 0x23, 0x39, 0xc6,
	0x09, 0x8f, 0xe5, 0x18, 0x2e, 0x3c, 0x96, 0x63, 0xb8, 0xf1, 0x58, 0x8e, 0x21, 0xca, 0x34, 0x3d,
	0xb3, 0x24, 0xa3, 0x34, 0x49, 0x2f, 0x39, 0x3f, 0x57, 0x1f, 0x1a, 0x64, 0xc9, 0x19, 0x89, 0x99,
	0x79, 0x30, 0x8e, 0x7e, 0x05, 0x5a, 0x40, 0x97, 0x54, 0x16, 0xa4, 0x16, 0x27, 0xb1, 0x81, 0x43,
	0xcd, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0x67, 0xe2, 0x30, 0x84, 0x8e, 0x01, 0x00, 0x00,
}

func (this *Params) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Params)
	if !ok {
		that2, ok := that.(Params)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.BtcConfirmationDepth != that1.BtcConfirmationDepth {
		return false
	}
	if this.CheckpointFinalizationTimeout != that1.CheckpointFinalizationTimeout {
		return false
	}
	return true
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
	if m.CheckpointFinalizationTimeout != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.CheckpointFinalizationTimeout))
		i--
		dAtA[i] = 0x10
	}
	if m.BtcConfirmationDepth != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.BtcConfirmationDepth))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintParams(dAtA []byte, offset int, v uint64) int {
	offset -= sovParams(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Params) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.BtcConfirmationDepth != 0 {
		n += 1 + sovParams(uint64(m.BtcConfirmationDepth))
	}
	if m.CheckpointFinalizationTimeout != 0 {
		n += 1 + sovParams(uint64(m.CheckpointFinalizationTimeout))
	}
	return n
}

func sovParams(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozParams(x uint64) (n int) {
	return sovParams(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Params) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowParams
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
				return fmt.Errorf("proto: wrong wireType = %d for field BtcConfirmationDepth", wireType)
			}
			m.BtcConfirmationDepth = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.BtcConfirmationDepth |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CheckpointFinalizationTimeout", wireType)
			}
			m.CheckpointFinalizationTimeout = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.CheckpointFinalizationTimeout |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipParams(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthParams
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
func skipParams(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowParams
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
					return 0, ErrIntOverflowParams
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
					return 0, ErrIntOverflowParams
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
				return 0, ErrInvalidLengthParams
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupParams
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthParams
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthParams        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowParams          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupParams = fmt.Errorf("proto: unexpected end of group")
)
