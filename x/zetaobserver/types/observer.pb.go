// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: zetaobserver/observer.proto

package types

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
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

type ObserverChain int32

const (
	ObserverChain_EmptyObserver        ObserverChain = 0
	ObserverChain_EthChainObserver     ObserverChain = 1
	ObserverChain_ZetaChainObserver    ObserverChain = 2
	ObserverChain_BtcChainObserver     ObserverChain = 3
	ObserverChain_PolygonChainObserver ObserverChain = 4
	ObserverChain_BscChainObserver     ObserverChain = 5
)

var ObserverChain_name = map[int32]string{
	0: "EmptyObserver",
	1: "EthChainObserver",
	2: "ZetaChainObserver",
	3: "BtcChainObserver",
	4: "PolygonChainObserver",
	5: "BscChainObserver",
}

var ObserverChain_value = map[string]int32{
	"EmptyObserver":        0,
	"EthChainObserver":     1,
	"ZetaChainObserver":    2,
	"BtcChainObserver":     3,
	"PolygonChainObserver": 4,
	"BscChainObserver":     5,
}

func (x ObserverChain) String() string {
	return proto.EnumName(ObserverChain_name, int32(x))
}

func (ObserverChain) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_27c587491f61fd50, []int{0}
}

type ObservationType int32

const (
	ObservationType_EmptyObserverType ObservationType = 0
	ObservationType_InBoundTx         ObservationType = 1
	ObservationType_OutBoundTx        ObservationType = 2
	ObservationType_GasPrice          ObservationType = 3
)

var ObservationType_name = map[int32]string{
	0: "EmptyObserverType",
	1: "InBoundTx",
	2: "OutBoundTx",
	3: "GasPrice",
}

var ObservationType_value = map[string]int32{
	"EmptyObserverType": 0,
	"InBoundTx":         1,
	"OutBoundTx":        2,
	"GasPrice":          3,
}

func (x ObservationType) String() string {
	return proto.EnumName(ObservationType_name, int32(x))
}

func (ObservationType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_27c587491f61fd50, []int{1}
}

type ObserverMapper struct {
	Index           string          `protobuf:"bytes,1,opt,name=index,proto3" json:"index,omitempty"`
	ObserverChain   ObserverChain   `protobuf:"varint,2,opt,name=ObserverChain,proto3,enum=zetachain.zetacore.zetaobserver.ObserverChain" json:"ObserverChain,omitempty"`
	ObservationType ObservationType `protobuf:"varint,3,opt,name=ObservationType,proto3,enum=zetachain.zetacore.zetaobserver.ObservationType" json:"ObservationType,omitempty"`
	ObserverList    []string        `protobuf:"bytes,4,rep,name=observerList,proto3" json:"observerList,omitempty"`
}

func (m *ObserverMapper) Reset()         { *m = ObserverMapper{} }
func (m *ObserverMapper) String() string { return proto.CompactTextString(m) }
func (*ObserverMapper) ProtoMessage()    {}
func (*ObserverMapper) Descriptor() ([]byte, []int) {
	return fileDescriptor_27c587491f61fd50, []int{0}
}
func (m *ObserverMapper) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ObserverMapper) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ObserverMapper.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ObserverMapper) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ObserverMapper.Merge(m, src)
}
func (m *ObserverMapper) XXX_Size() int {
	return m.Size()
}
func (m *ObserverMapper) XXX_DiscardUnknown() {
	xxx_messageInfo_ObserverMapper.DiscardUnknown(m)
}

var xxx_messageInfo_ObserverMapper proto.InternalMessageInfo

func (m *ObserverMapper) GetIndex() string {
	if m != nil {
		return m.Index
	}
	return ""
}

func (m *ObserverMapper) GetObserverChain() ObserverChain {
	if m != nil {
		return m.ObserverChain
	}
	return ObserverChain_EmptyObserver
}

func (m *ObserverMapper) GetObservationType() ObservationType {
	if m != nil {
		return m.ObservationType
	}
	return ObservationType_EmptyObserverType
}

func (m *ObserverMapper) GetObserverList() []string {
	if m != nil {
		return m.ObserverList
	}
	return nil
}

func init() {
	proto.RegisterEnum("zetachain.zetacore.zetaobserver.ObserverChain", ObserverChain_name, ObserverChain_value)
	proto.RegisterEnum("zetachain.zetacore.zetaobserver.ObservationType", ObservationType_name, ObservationType_value)
	proto.RegisterType((*ObserverMapper)(nil), "zetachain.zetacore.zetaobserver.ObserverMapper")
}

func init() { proto.RegisterFile("zetaobserver/observer.proto", fileDescriptor_27c587491f61fd50) }

var fileDescriptor_27c587491f61fd50 = []byte{
	// 374 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0x41, 0x6b, 0xea, 0x40,
	0x10, 0xc7, 0xb3, 0x46, 0x1f, 0xcf, 0x41, 0x7d, 0x71, 0xc9, 0x83, 0xe0, 0x83, 0x3c, 0xf1, 0x24,
	0x42, 0x37, 0xa5, 0x7e, 0x03, 0x8b, 0x94, 0x42, 0x45, 0x09, 0x9e, 0xec, 0x29, 0xc6, 0x25, 0x06,
	0x6a, 0x36, 0x24, 0x6b, 0x31, 0xfd, 0x14, 0xbd, 0xf7, 0xda, 0x43, 0x3f, 0x4a, 0x8f, 0x1e, 0x7b,
	0x2c, 0xfa, 0x35, 0x7a, 0x28, 0xbb, 0x69, 0xa4, 0x6b, 0x0f, 0xed, 0x6d, 0xf6, 0x3f, 0xbf, 0xf9,
	0x4f, 0x32, 0x33, 0xf0, 0xef, 0x8e, 0x72, 0x8f, 0xcd, 0x53, 0x9a, 0xdc, 0xd2, 0xc4, 0x29, 0x02,
	0x12, 0x27, 0x8c, 0x33, 0xfc, 0x5f, 0x24, 0xfd, 0xa5, 0x17, 0x46, 0x44, 0x46, 0x2c, 0xa1, 0xe4,
	0x33, 0xdf, 0x32, 0x03, 0x16, 0x30, 0xc9, 0x3a, 0x22, 0xca, 0xcb, 0x3a, 0x6f, 0x08, 0x1a, 0xe3,
	0x0f, 0x64, 0xe4, 0xc5, 0x31, 0x4d, 0xb0, 0x09, 0x95, 0x30, 0x5a, 0xd0, 0x8d, 0x85, 0xda, 0xa8,
	0x5b, 0x75, 0xf3, 0x07, 0x9e, 0x42, 0xbd, 0xe0, 0xce, 0x45, 0x17, 0xab, 0xd4, 0x46, 0xdd, 0xc6,
	0x19, 0x21, 0xdf, 0xf4, 0x25, 0x4a, 0x95, 0xab, 0x9a, 0xe0, 0x19, 0xfc, 0xc9, 0x05, 0x8f, 0x87,
	0x2c, 0x9a, 0x66, 0x31, 0xb5, 0x74, 0xe9, 0x7b, 0xfa, 0x43, 0xdf, 0x43, 0x9d, 0x7b, 0x6c, 0x84,
	0x3b, 0x50, 0x2b, 0xe0, 0xab, 0x30, 0xe5, 0x56, 0xb9, 0xad, 0x77, 0xab, 0xae, 0xa2, 0xf5, 0x1e,
	0xd0, 0xd1, 0x6f, 0xe1, 0x26, 0xd4, 0x87, 0xab, 0x98, 0x67, 0x85, 0x6a, 0x68, 0xd8, 0x04, 0x63,
	0xc8, 0x97, 0x32, 0x7d, 0x50, 0x11, 0xfe, 0x0b, 0xcd, 0x19, 0xe5, 0x9e, 0x2a, 0x97, 0x04, 0x3c,
	0xe0, 0xbe, 0xaa, 0xea, 0xd8, 0x02, 0x73, 0xc2, 0x6e, 0xb2, 0x80, 0x45, 0x6a, 0xa6, 0x2c, 0xf9,
	0xf4, 0x88, 0xaf, 0xb4, 0xca, 0x4f, 0x8f, 0x36, 0xea, 0x5d, 0x7f, 0x99, 0x8e, 0xe8, 0xaa, 0x7c,
	0x9e, 0x10, 0x0d, 0x0d, 0xd7, 0xa1, 0x7a, 0x19, 0x0d, 0xd8, 0x3a, 0x5a, 0x4c, 0x37, 0x06, 0xc2,
	0x0d, 0x80, 0xf1, 0x9a, 0x17, 0xef, 0x12, 0xae, 0xc1, 0xef, 0x0b, 0x2f, 0x9d, 0x24, 0xa1, 0x4f,
	0x0d, 0x3d, 0x37, 0x1f, 0x8c, 0x9e, 0x77, 0x36, 0xda, 0xee, 0x6c, 0xf4, 0xba, 0xb3, 0xd1, 0xfd,
	0xde, 0xd6, 0xb6, 0x7b, 0x5b, 0x7b, 0xd9, 0xdb, 0xda, 0xac, 0x1f, 0x84, 0x7c, 0xb9, 0x9e, 0x13,
	0x9f, 0xad, 0x1c, 0x31, 0xf2, 0x13, 0xb9, 0x06, 0xa7, 0x58, 0x83, 0xb3, 0x71, 0x94, 0x43, 0xe4,
	0x59, 0x4c, 0xd3, 0xf9, 0x2f, 0x79, 0x4f, 0xfd, 0xf7, 0x00, 0x00, 0x00, 0xff, 0xff, 0x01, 0x7f,
	0xb1, 0x06, 0xa5, 0x02, 0x00, 0x00,
}

func (m *ObserverMapper) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ObserverMapper) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ObserverMapper) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.ObserverList) > 0 {
		for iNdEx := len(m.ObserverList) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.ObserverList[iNdEx])
			copy(dAtA[i:], m.ObserverList[iNdEx])
			i = encodeVarintObserver(dAtA, i, uint64(len(m.ObserverList[iNdEx])))
			i--
			dAtA[i] = 0x22
		}
	}
	if m.ObservationType != 0 {
		i = encodeVarintObserver(dAtA, i, uint64(m.ObservationType))
		i--
		dAtA[i] = 0x18
	}
	if m.ObserverChain != 0 {
		i = encodeVarintObserver(dAtA, i, uint64(m.ObserverChain))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Index) > 0 {
		i -= len(m.Index)
		copy(dAtA[i:], m.Index)
		i = encodeVarintObserver(dAtA, i, uint64(len(m.Index)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintObserver(dAtA []byte, offset int, v uint64) int {
	offset -= sovObserver(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *ObserverMapper) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Index)
	if l > 0 {
		n += 1 + l + sovObserver(uint64(l))
	}
	if m.ObserverChain != 0 {
		n += 1 + sovObserver(uint64(m.ObserverChain))
	}
	if m.ObservationType != 0 {
		n += 1 + sovObserver(uint64(m.ObservationType))
	}
	if len(m.ObserverList) > 0 {
		for _, s := range m.ObserverList {
			l = len(s)
			n += 1 + l + sovObserver(uint64(l))
		}
	}
	return n
}

func sovObserver(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozObserver(x uint64) (n int) {
	return sovObserver(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ObserverMapper) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowObserver
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
			return fmt.Errorf("proto: ObserverMapper: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ObserverMapper: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Index", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowObserver
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
				return ErrInvalidLengthObserver
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthObserver
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Index = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ObserverChain", wireType)
			}
			m.ObserverChain = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowObserver
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ObserverChain |= ObserverChain(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ObservationType", wireType)
			}
			m.ObservationType = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowObserver
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ObservationType |= ObservationType(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ObserverList", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowObserver
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
				return ErrInvalidLengthObserver
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthObserver
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ObserverList = append(m.ObserverList, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipObserver(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthObserver
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
func skipObserver(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowObserver
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
					return 0, ErrIntOverflowObserver
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
					return 0, ErrIntOverflowObserver
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
				return 0, ErrInvalidLengthObserver
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupObserver
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthObserver
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthObserver        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowObserver          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupObserver = fmt.Errorf("proto: unexpected end of group")
)