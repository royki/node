// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: zetachain/zetacore/crosschain/rate_limiter_flags.proto

package types

import (
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
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

type RateLimiterFlags struct {
	Enabled bool `protobuf:"varint,1,opt,name=enabled,proto3" json:"enabled,omitempty"`
	// window in blocks
	Window int64 `protobuf:"varint,2,opt,name=window,proto3" json:"window,omitempty"`
	// rate in azeta per block
	Rate github_com_cosmos_cosmos_sdk_types.Uint `protobuf:"bytes,3,opt,name=rate,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Uint" json:"rate"`
	// conversion in azeta per token
	Conversions []Conversion `protobuf:"bytes,4,rep,name=conversions,proto3" json:"conversions"`
}

func (m *RateLimiterFlags) Reset()         { *m = RateLimiterFlags{} }
func (m *RateLimiterFlags) String() string { return proto.CompactTextString(m) }
func (*RateLimiterFlags) ProtoMessage()    {}
func (*RateLimiterFlags) Descriptor() ([]byte, []int) {
	return fileDescriptor_9c435f4c2dabc0eb, []int{0}
}
func (m *RateLimiterFlags) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *RateLimiterFlags) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_RateLimiterFlags.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *RateLimiterFlags) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RateLimiterFlags.Merge(m, src)
}
func (m *RateLimiterFlags) XXX_Size() int {
	return m.Size()
}
func (m *RateLimiterFlags) XXX_DiscardUnknown() {
	xxx_messageInfo_RateLimiterFlags.DiscardUnknown(m)
}

var xxx_messageInfo_RateLimiterFlags proto.InternalMessageInfo

func (m *RateLimiterFlags) GetEnabled() bool {
	if m != nil {
		return m.Enabled
	}
	return false
}

func (m *RateLimiterFlags) GetWindow() int64 {
	if m != nil {
		return m.Window
	}
	return 0
}

func (m *RateLimiterFlags) GetConversions() []Conversion {
	if m != nil {
		return m.Conversions
	}
	return nil
}

type Conversion struct {
	Zrc20 string                                 `protobuf:"bytes,1,opt,name=zrc20,proto3" json:"zrc20,omitempty"`
	Rate  github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,2,opt,name=rate,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"rate"`
}

func (m *Conversion) Reset()         { *m = Conversion{} }
func (m *Conversion) String() string { return proto.CompactTextString(m) }
func (*Conversion) ProtoMessage()    {}
func (*Conversion) Descriptor() ([]byte, []int) {
	return fileDescriptor_9c435f4c2dabc0eb, []int{1}
}
func (m *Conversion) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Conversion) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Conversion.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Conversion) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Conversion.Merge(m, src)
}
func (m *Conversion) XXX_Size() int {
	return m.Size()
}
func (m *Conversion) XXX_DiscardUnknown() {
	xxx_messageInfo_Conversion.DiscardUnknown(m)
}

var xxx_messageInfo_Conversion proto.InternalMessageInfo

func (m *Conversion) GetZrc20() string {
	if m != nil {
		return m.Zrc20
	}
	return ""
}

func init() {
	proto.RegisterType((*RateLimiterFlags)(nil), "zetachain.zetacore.crosschain.RateLimiterFlags")
	proto.RegisterType((*Conversion)(nil), "zetachain.zetacore.crosschain.Conversion")
}

func init() {
	proto.RegisterFile("zetachain/zetacore/crosschain/rate_limiter_flags.proto", fileDescriptor_9c435f4c2dabc0eb)
}

var fileDescriptor_9c435f4c2dabc0eb = []byte{
	// 335 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x91, 0xcf, 0x4e, 0xfa, 0x40,
	0x10, 0xc7, 0xbb, 0xc0, 0x8f, 0x9f, 0x2c, 0x17, 0xb3, 0x21, 0xa6, 0x31, 0xb1, 0x34, 0x1c, 0xb4,
	0x1e, 0xd8, 0x2a, 0x26, 0x3e, 0x40, 0x31, 0x5e, 0xf4, 0x62, 0x13, 0x2f, 0x5e, 0x48, 0x59, 0x96,
	0xb2, 0x11, 0x3a, 0x64, 0x77, 0x15, 0xe5, 0x29, 0x7c, 0x2c, 0x8e, 0x1c, 0x8d, 0x31, 0xc4, 0xc0,
	0x8b, 0x98, 0x6e, 0xcb, 0x9f, 0x78, 0x30, 0x9e, 0x3a, 0xd3, 0xcc, 0xe7, 0xbb, 0x9f, 0xc9, 0xe0,
	0xcb, 0x29, 0xd7, 0x11, 0x1b, 0x44, 0x22, 0xf1, 0x4d, 0x05, 0x92, 0xfb, 0x4c, 0x82, 0x52, 0xd9,
	0x3f, 0x19, 0x69, 0xde, 0x19, 0x8a, 0x91, 0xd0, 0x5c, 0x76, 0xfa, 0xc3, 0x28, 0x56, 0x74, 0x2c,
	0x41, 0x03, 0x39, 0xda, 0x70, 0x74, 0xcd, 0xd1, 0x2d, 0x77, 0x58, 0x8b, 0x21, 0x06, 0x33, 0xe9,
	0xa7, 0x55, 0x06, 0x35, 0x3e, 0x11, 0xde, 0x0f, 0x23, 0xcd, 0x6f, 0xb3, 0xc0, 0xeb, 0x34, 0x8f,
	0xd8, 0xf8, 0x3f, 0x4f, 0xa2, 0xee, 0x90, 0xf7, 0x6c, 0xe4, 0x22, 0x6f, 0x2f, 0x5c, 0xb7, 0xe4,
	0x00, 0x97, 0x27, 0x22, 0xe9, 0xc1, 0xc4, 0x2e, 0xb8, 0xc8, 0x2b, 0x86, 0x79, 0x47, 0xda, 0xb8,
	0x94, 0x7a, 0xd9, 0x45, 0x17, 0x79, 0x95, 0xc0, 0x9f, 0x2d, 0xea, 0xd6, 0xc7, 0xa2, 0x7e, 0x12,
	0x0b, 0x3d, 0x78, 0xea, 0x52, 0x06, 0x23, 0x9f, 0x81, 0x1a, 0x81, 0xca, 0x3f, 0x4d, 0xd5, 0x7b,
	0xf4, 0xf5, 0xeb, 0x98, 0x2b, 0x7a, 0x2f, 0x12, 0x1d, 0x1a, 0x98, 0xdc, 0xe1, 0x2a, 0x83, 0xe4,
	0x99, 0x4b, 0x25, 0x20, 0x51, 0x76, 0xc9, 0x2d, 0x7a, 0xd5, 0xd6, 0x29, 0xfd, 0x75, 0x2d, 0xda,
	0xde, 0x10, 0x41, 0x29, 0x7d, 0x36, 0xdc, 0xcd, 0x68, 0xf4, 0x31, 0xde, 0x0e, 0x90, 0x1a, 0xfe,
	0x37, 0x95, 0xac, 0x75, 0x66, 0xb6, 0xaa, 0x84, 0x59, 0x43, 0x82, 0xdc, 0xbd, 0x60, 0xdc, 0x69,
	0xee, 0x7e, 0xfc, 0x07, 0xf7, 0x2b, 0xce, 0x32, 0xf5, 0xe0, 0x66, 0xb6, 0x74, 0xd0, 0x7c, 0xe9,
	0xa0, 0xaf, 0xa5, 0x83, 0xde, 0x56, 0x8e, 0x35, 0x5f, 0x39, 0xd6, 0xfb, 0xca, 0xb1, 0x1e, 0xce,
	0x77, 0x72, 0x52, 0xff, 0xe6, 0x8f, 0xcb, 0xbe, 0xec, 0xde, 0xd6, 0xc4, 0x76, 0xcb, 0xe6, 0x34,
	0x17, 0xdf, 0x01, 0x00, 0x00, 0xff, 0xff, 0x16, 0x2b, 0x9e, 0xbb, 0x09, 0x02, 0x00, 0x00,
}

func (m *RateLimiterFlags) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RateLimiterFlags) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *RateLimiterFlags) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Conversions) > 0 {
		for iNdEx := len(m.Conversions) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Conversions[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintRateLimiterFlags(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x22
		}
	}
	{
		size := m.Rate.Size()
		i -= size
		if _, err := m.Rate.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintRateLimiterFlags(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	if m.Window != 0 {
		i = encodeVarintRateLimiterFlags(dAtA, i, uint64(m.Window))
		i--
		dAtA[i] = 0x10
	}
	if m.Enabled {
		i--
		if m.Enabled {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *Conversion) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Conversion) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Conversion) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.Rate.Size()
		i -= size
		if _, err := m.Rate.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintRateLimiterFlags(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if len(m.Zrc20) > 0 {
		i -= len(m.Zrc20)
		copy(dAtA[i:], m.Zrc20)
		i = encodeVarintRateLimiterFlags(dAtA, i, uint64(len(m.Zrc20)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintRateLimiterFlags(dAtA []byte, offset int, v uint64) int {
	offset -= sovRateLimiterFlags(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *RateLimiterFlags) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Enabled {
		n += 2
	}
	if m.Window != 0 {
		n += 1 + sovRateLimiterFlags(uint64(m.Window))
	}
	l = m.Rate.Size()
	n += 1 + l + sovRateLimiterFlags(uint64(l))
	if len(m.Conversions) > 0 {
		for _, e := range m.Conversions {
			l = e.Size()
			n += 1 + l + sovRateLimiterFlags(uint64(l))
		}
	}
	return n
}

func (m *Conversion) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Zrc20)
	if l > 0 {
		n += 1 + l + sovRateLimiterFlags(uint64(l))
	}
	l = m.Rate.Size()
	n += 1 + l + sovRateLimiterFlags(uint64(l))
	return n
}

func sovRateLimiterFlags(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozRateLimiterFlags(x uint64) (n int) {
	return sovRateLimiterFlags(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *RateLimiterFlags) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRateLimiterFlags
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
			return fmt.Errorf("proto: RateLimiterFlags: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RateLimiterFlags: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Enabled", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRateLimiterFlags
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Enabled = bool(v != 0)
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Window", wireType)
			}
			m.Window = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRateLimiterFlags
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Window |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Rate", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRateLimiterFlags
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
				return ErrInvalidLengthRateLimiterFlags
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthRateLimiterFlags
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Rate.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Conversions", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRateLimiterFlags
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
				return ErrInvalidLengthRateLimiterFlags
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthRateLimiterFlags
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Conversions = append(m.Conversions, Conversion{})
			if err := m.Conversions[len(m.Conversions)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipRateLimiterFlags(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthRateLimiterFlags
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
func (m *Conversion) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRateLimiterFlags
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
			return fmt.Errorf("proto: Conversion: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Conversion: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Zrc20", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRateLimiterFlags
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
				return ErrInvalidLengthRateLimiterFlags
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthRateLimiterFlags
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Zrc20 = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Rate", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRateLimiterFlags
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
				return ErrInvalidLengthRateLimiterFlags
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthRateLimiterFlags
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Rate.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipRateLimiterFlags(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthRateLimiterFlags
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
func skipRateLimiterFlags(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowRateLimiterFlags
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
					return 0, ErrIntOverflowRateLimiterFlags
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
					return 0, ErrIntOverflowRateLimiterFlags
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
				return 0, ErrInvalidLengthRateLimiterFlags
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupRateLimiterFlags
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthRateLimiterFlags
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthRateLimiterFlags        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowRateLimiterFlags          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupRateLimiterFlags = fmt.Errorf("proto: unexpected end of group")
)