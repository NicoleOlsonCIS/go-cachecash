// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: faucet.proto

package ccmsg

import (
	context "context"
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
	grpc "google.golang.org/grpc"
	io "io"
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
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type GetCoinsRequest struct {
	PublicKey            *PublicKey `protobuf:"bytes,1,opt,name=public_key,json=publicKey,proto3" json:"public_key,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *GetCoinsRequest) Reset()         { *m = GetCoinsRequest{} }
func (m *GetCoinsRequest) String() string { return proto.CompactTextString(m) }
func (*GetCoinsRequest) ProtoMessage()    {}
func (*GetCoinsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_1ed390957cc073ed, []int{0}
}
func (m *GetCoinsRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GetCoinsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GetCoinsRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GetCoinsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetCoinsRequest.Merge(m, src)
}
func (m *GetCoinsRequest) XXX_Size() int {
	return m.Size()
}
func (m *GetCoinsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetCoinsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetCoinsRequest proto.InternalMessageInfo

func (m *GetCoinsRequest) GetPublicKey() *PublicKey {
	if m != nil {
		return m.PublicKey
	}
	return nil
}

func init() {
	proto.RegisterType((*GetCoinsRequest)(nil), "ccmsg.GetCoinsRequest")
}

func init() { proto.RegisterFile("faucet.proto", fileDescriptor_1ed390957cc073ed) }

var fileDescriptor_1ed390957cc073ed = []byte{
	// 215 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x49, 0x4b, 0x2c, 0x4d,
	0x4e, 0x2d, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x4d, 0x4e, 0xce, 0x2d, 0x4e, 0x97,
	0xd2, 0x4d, 0xcf, 0x2c, 0xc9, 0x28, 0x4d, 0xd2, 0x4b, 0xce, 0xcf, 0xd5, 0x4f, 0xcf, 0x4f, 0xcf,
	0xd7, 0x07, 0xcb, 0x26, 0x95, 0xa6, 0x81, 0x79, 0x60, 0x0e, 0x98, 0x05, 0xd1, 0x25, 0x25, 0x9d,
	0x9e, 0x9f, 0x9f, 0x9e, 0x93, 0x8a, 0x50, 0x95, 0x9a, 0x5b, 0x50, 0x52, 0x09, 0x95, 0xe4, 0x49,
	0xce, 0xcf, 0xcd, 0xcd, 0xcf, 0x83, 0xf0, 0x94, 0x9c, 0xb8, 0xf8, 0xdd, 0x53, 0x4b, 0x9c, 0xf3,
	0x33, 0xf3, 0x8a, 0x83, 0x52, 0x0b, 0x4b, 0x53, 0x8b, 0x4b, 0x84, 0xf4, 0xb9, 0xb8, 0x0a, 0x4a,
	0x93, 0x72, 0x32, 0x93, 0xe3, 0xb3, 0x53, 0x2b, 0x25, 0x18, 0x15, 0x18, 0x35, 0xb8, 0x8d, 0x04,
	0xf4, 0xc0, 0x0e, 0xd1, 0x0b, 0x00, 0x4b, 0x78, 0xa7, 0x56, 0x06, 0x71, 0x16, 0xc0, 0x98, 0x46,
	0x6e, 0x5c, 0x6c, 0x6e, 0x60, 0x47, 0x0b, 0xd9, 0x70, 0x71, 0xc0, 0x4c, 0x13, 0x12, 0x83, 0x6a,
	0x41, 0x33, 0x5e, 0x4a, 0x4c, 0x0f, 0xe2, 0x3a, 0x3d, 0x98, 0xeb, 0xf4, 0x5c, 0x41, 0xae, 0x53,
	0x62, 0x70, 0x12, 0x38, 0xf1, 0x48, 0x8e, 0xf1, 0xc2, 0x23, 0x39, 0xc6, 0x07, 0x8f, 0xe4, 0x18,
	0x67, 0x3c, 0x96, 0x63, 0x48, 0x62, 0x03, 0xab, 0x31, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0x79,
	0x51, 0x22, 0x22, 0x15, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// FaucetClient is the client API for Faucet service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type FaucetClient interface {
	GetCoins(ctx context.Context, in *GetCoinsRequest, opts ...grpc.CallOption) (*empty.Empty, error)
}

type faucetClient struct {
	cc *grpc.ClientConn
}

func NewFaucetClient(cc *grpc.ClientConn) FaucetClient {
	return &faucetClient{cc}
}

func (c *faucetClient) GetCoins(ctx context.Context, in *GetCoinsRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/ccmsg.Faucet/GetCoins", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FaucetServer is the server API for Faucet service.
type FaucetServer interface {
	GetCoins(context.Context, *GetCoinsRequest) (*empty.Empty, error)
}

func RegisterFaucetServer(s *grpc.Server, srv FaucetServer) {
	s.RegisterService(&_Faucet_serviceDesc, srv)
}

func _Faucet_GetCoins_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCoinsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FaucetServer).GetCoins(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ccmsg.Faucet/GetCoins",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FaucetServer).GetCoins(ctx, req.(*GetCoinsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Faucet_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ccmsg.Faucet",
	HandlerType: (*FaucetServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetCoins",
			Handler:    _Faucet_GetCoins_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "faucet.proto",
}

func (m *GetCoinsRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GetCoinsRequest) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.PublicKey != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintFaucet(dAtA, i, uint64(m.PublicKey.Size()))
		n1, err1 := m.PublicKey.MarshalTo(dAtA[i:])
		if err1 != nil {
			return 0, err1
		}
		i += n1
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func encodeVarintFaucet(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *GetCoinsRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.PublicKey != nil {
		l = m.PublicKey.Size()
		n += 1 + l + sovFaucet(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovFaucet(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozFaucet(x uint64) (n int) {
	return sovFaucet(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *GetCoinsRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowFaucet
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
			return fmt.Errorf("proto: GetCoinsRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GetCoinsRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PublicKey", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFaucet
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
				return ErrInvalidLengthFaucet
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthFaucet
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.PublicKey == nil {
				m.PublicKey = &PublicKey{}
			}
			if err := m.PublicKey.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipFaucet(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthFaucet
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthFaucet
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipFaucet(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowFaucet
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
					return 0, ErrIntOverflowFaucet
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowFaucet
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
				return 0, ErrInvalidLengthFaucet
			}
			iNdEx += length
			if iNdEx < 0 {
				return 0, ErrInvalidLengthFaucet
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowFaucet
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipFaucet(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
				if iNdEx < 0 {
					return 0, ErrInvalidLengthFaucet
				}
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthFaucet = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowFaucet   = fmt.Errorf("proto: integer overflow")
)