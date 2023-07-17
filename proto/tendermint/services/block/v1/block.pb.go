// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: tendermint/services/block/v1/block.proto

package v1

import (
	fmt "fmt"
	types "github.com/cometbft/cometbft/proto/tendermint/types"
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

type GetByHeightRequest struct {
	// the height of the block requested.
	Height int64 `protobuf:"varint,1,opt,name=height,proto3" json:"height,omitempty"`
}

func (m *GetByHeightRequest) Reset()         { *m = GetByHeightRequest{} }
func (m *GetByHeightRequest) String() string { return proto.CompactTextString(m) }
func (*GetByHeightRequest) ProtoMessage()    {}
func (*GetByHeightRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_d48acf20d1015667, []int{0}
}
func (m *GetByHeightRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GetByHeightRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GetByHeightRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GetByHeightRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetByHeightRequest.Merge(m, src)
}
func (m *GetByHeightRequest) XXX_Size() int {
	return m.Size()
}
func (m *GetByHeightRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetByHeightRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetByHeightRequest proto.InternalMessageInfo

func (m *GetByHeightRequest) GetHeight() int64 {
	if m != nil {
		return m.Height
	}
	return 0
}

type GetByHeightResponse struct {
	BlockId *types.BlockID `protobuf:"bytes,1,opt,name=block_id,json=blockId,proto3" json:"block_id,omitempty"`
	Block   *types.Block   `protobuf:"bytes,2,opt,name=block,proto3" json:"block,omitempty"`
}

func (m *GetByHeightResponse) Reset()         { *m = GetByHeightResponse{} }
func (m *GetByHeightResponse) String() string { return proto.CompactTextString(m) }
func (*GetByHeightResponse) ProtoMessage()    {}
func (*GetByHeightResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_d48acf20d1015667, []int{1}
}
func (m *GetByHeightResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GetByHeightResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GetByHeightResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GetByHeightResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetByHeightResponse.Merge(m, src)
}
func (m *GetByHeightResponse) XXX_Size() int {
	return m.Size()
}
func (m *GetByHeightResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetByHeightResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetByHeightResponse proto.InternalMessageInfo

func (m *GetByHeightResponse) GetBlockId() *types.BlockID {
	if m != nil {
		return m.BlockId
	}
	return nil
}

func (m *GetByHeightResponse) GetBlock() *types.Block {
	if m != nil {
		return m.Block
	}
	return nil
}

// GetLatestHeightRequest - empty message since no parameter is required
type GetLatestHeightRequest struct {
}

func (m *GetLatestHeightRequest) Reset()         { *m = GetLatestHeightRequest{} }
func (m *GetLatestHeightRequest) String() string { return proto.CompactTextString(m) }
func (*GetLatestHeightRequest) ProtoMessage()    {}
func (*GetLatestHeightRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_d48acf20d1015667, []int{2}
}
func (m *GetLatestHeightRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GetLatestHeightRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GetLatestHeightRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GetLatestHeightRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetLatestHeightRequest.Merge(m, src)
}
func (m *GetLatestHeightRequest) XXX_Size() int {
	return m.Size()
}
func (m *GetLatestHeightRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetLatestHeightRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetLatestHeightRequest proto.InternalMessageInfo

// GetLatestHeightResponse provides the height of the latest committed block.
type GetLatestHeightResponse struct {
	// The height of the latest committed block. Will be 0 if no data has been
	// committed yet.
	Height int64 `protobuf:"varint,1,opt,name=height,proto3" json:"height,omitempty"`
}

func (m *GetLatestHeightResponse) Reset()         { *m = GetLatestHeightResponse{} }
func (m *GetLatestHeightResponse) String() string { return proto.CompactTextString(m) }
func (*GetLatestHeightResponse) ProtoMessage()    {}
func (*GetLatestHeightResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_d48acf20d1015667, []int{3}
}
func (m *GetLatestHeightResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GetLatestHeightResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GetLatestHeightResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GetLatestHeightResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetLatestHeightResponse.Merge(m, src)
}
func (m *GetLatestHeightResponse) XXX_Size() int {
	return m.Size()
}
func (m *GetLatestHeightResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetLatestHeightResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetLatestHeightResponse proto.InternalMessageInfo

func (m *GetLatestHeightResponse) GetHeight() int64 {
	if m != nil {
		return m.Height
	}
	return 0
}

func init() {
	proto.RegisterType((*GetByHeightRequest)(nil), "tendermint.services.block.v1.GetByHeightRequest")
	proto.RegisterType((*GetByHeightResponse)(nil), "tendermint.services.block.v1.GetByHeightResponse")
	proto.RegisterType((*GetLatestHeightRequest)(nil), "tendermint.services.block.v1.GetLatestHeightRequest")
	proto.RegisterType((*GetLatestHeightResponse)(nil), "tendermint.services.block.v1.GetLatestHeightResponse")
}

func init() {
	proto.RegisterFile("tendermint/services/block/v1/block.proto", fileDescriptor_d48acf20d1015667)
}

var fileDescriptor_d48acf20d1015667 = []byte{
	// 274 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xd2, 0x28, 0x49, 0xcd, 0x4b,
	0x49, 0x2d, 0xca, 0xcd, 0xcc, 0x2b, 0xd1, 0x2f, 0x4e, 0x2d, 0x2a, 0xcb, 0x4c, 0x4e, 0x2d, 0xd6,
	0x4f, 0xca, 0xc9, 0x4f, 0xce, 0xd6, 0x2f, 0x33, 0x84, 0x30, 0xf4, 0x0a, 0x8a, 0xf2, 0x4b, 0xf2,
	0x85, 0x64, 0x10, 0x2a, 0xf5, 0x60, 0x2a, 0xf5, 0x20, 0x0a, 0xca, 0x0c, 0xa5, 0x90, 0x64, 0xf5,
	0x4b, 0x2a, 0x0b, 0x60, 0x86, 0x40, 0xf4, 0x62, 0x91, 0x05, 0x93, 0x10, 0x59, 0x25, 0x1d, 0x2e,
	0x21, 0xf7, 0xd4, 0x12, 0xa7, 0x4a, 0x8f, 0xd4, 0xcc, 0xf4, 0x8c, 0x92, 0xa0, 0xd4, 0xc2, 0xd2,
	0xd4, 0xe2, 0x12, 0x21, 0x31, 0x2e, 0xb6, 0x0c, 0xb0, 0x80, 0x04, 0xa3, 0x02, 0xa3, 0x06, 0x73,
	0x10, 0x94, 0xa7, 0x54, 0xc5, 0x25, 0x8c, 0xa2, 0xba, 0xb8, 0x20, 0x3f, 0xaf, 0x38, 0x55, 0xc8,
	0x84, 0x8b, 0x03, 0x6c, 0x63, 0x7c, 0x66, 0x0a, 0x58, 0x03, 0xb7, 0x91, 0xa4, 0x1e, 0x92, 0x8b,
	0x21, 0xf6, 0x39, 0x81, 0x54, 0x78, 0xba, 0x04, 0xb1, 0x83, 0x95, 0x7a, 0xa6, 0x08, 0xe9, 0x72,
	0xb1, 0x82, 0x99, 0x12, 0x4c, 0x60, 0x2d, 0xe2, 0x38, 0xb4, 0x04, 0x41, 0x54, 0x29, 0x49, 0x70,
	0x89, 0xb9, 0xa7, 0x96, 0xf8, 0x24, 0x96, 0xa4, 0x16, 0x97, 0xa0, 0xb8, 0x56, 0xc9, 0x90, 0x4b,
	0x1c, 0x43, 0x06, 0xea, 0x32, 0x1c, 0x1e, 0x71, 0x8a, 0x3c, 0xf1, 0x48, 0x8e, 0xf1, 0xc2, 0x23,
	0x39, 0xc6, 0x07, 0x8f, 0xe4, 0x18, 0x27, 0x3c, 0x96, 0x63, 0xb8, 0xf0, 0x58, 0x8e, 0xe1, 0xc6,
	0x63, 0x39, 0x86, 0x28, 0xfb, 0xf4, 0xcc, 0x92, 0x8c, 0xd2, 0x24, 0xbd, 0xe4, 0xfc, 0x5c, 0xfd,
	0xe4, 0xfc, 0xdc, 0xd4, 0x92, 0xa4, 0xb4, 0x12, 0x04, 0x03, 0x1c, 0x68, 0xfa, 0xf8, 0xe2, 0x2d,
	0x89, 0x0d, 0xac, 0xc6, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0xb2, 0xf1, 0x4b, 0x36, 0xde, 0x01,
	0x00, 0x00,
}

func (m *GetByHeightRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GetByHeightRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GetByHeightRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Height != 0 {
		i = encodeVarintBlock(dAtA, i, uint64(m.Height))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *GetByHeightResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GetByHeightResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GetByHeightResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Block != nil {
		{
			size, err := m.Block.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintBlock(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if m.BlockId != nil {
		{
			size, err := m.BlockId.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintBlock(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *GetLatestHeightRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GetLatestHeightRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GetLatestHeightRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *GetLatestHeightResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GetLatestHeightResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GetLatestHeightResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Height != 0 {
		i = encodeVarintBlock(dAtA, i, uint64(m.Height))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintBlock(dAtA []byte, offset int, v uint64) int {
	offset -= sovBlock(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *GetByHeightRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Height != 0 {
		n += 1 + sovBlock(uint64(m.Height))
	}
	return n
}

func (m *GetByHeightResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.BlockId != nil {
		l = m.BlockId.Size()
		n += 1 + l + sovBlock(uint64(l))
	}
	if m.Block != nil {
		l = m.Block.Size()
		n += 1 + l + sovBlock(uint64(l))
	}
	return n
}

func (m *GetLatestHeightRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *GetLatestHeightResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Height != 0 {
		n += 1 + sovBlock(uint64(m.Height))
	}
	return n
}

func sovBlock(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozBlock(x uint64) (n int) {
	return sovBlock(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *GetByHeightRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowBlock
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
			return fmt.Errorf("proto: GetByHeightRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GetByHeightRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Height", wireType)
			}
			m.Height = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBlock
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Height |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipBlock(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthBlock
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
func (m *GetByHeightResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowBlock
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
			return fmt.Errorf("proto: GetByHeightResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GetByHeightResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BlockId", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBlock
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
				return ErrInvalidLengthBlock
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthBlock
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.BlockId == nil {
				m.BlockId = &types.BlockID{}
			}
			if err := m.BlockId.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Block", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBlock
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
				return ErrInvalidLengthBlock
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthBlock
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Block == nil {
				m.Block = &types.Block{}
			}
			if err := m.Block.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipBlock(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthBlock
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
func (m *GetLatestHeightRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowBlock
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
			return fmt.Errorf("proto: GetLatestHeightRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GetLatestHeightRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipBlock(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthBlock
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
func (m *GetLatestHeightResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowBlock
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
			return fmt.Errorf("proto: GetLatestHeightResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GetLatestHeightResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Height", wireType)
			}
			m.Height = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBlock
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Height |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipBlock(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthBlock
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
func skipBlock(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowBlock
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
					return 0, ErrIntOverflowBlock
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
					return 0, ErrIntOverflowBlock
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
				return 0, ErrInvalidLengthBlock
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupBlock
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthBlock
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthBlock        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowBlock          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupBlock = fmt.Errorf("proto: unexpected end of group")
)
