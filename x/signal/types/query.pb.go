// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: celestia/signal/v1/query.proto

package types

import (
	context "context"
	fmt "fmt"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

// QueryVersionTallyRequest is the request type for the VersionTally query.
type QueryVersionTallyRequest struct {
	Version uint64 `protobuf:"varint,1,opt,name=version,proto3" json:"version,omitempty"`
}

func (m *QueryVersionTallyRequest) Reset()         { *m = QueryVersionTallyRequest{} }
func (m *QueryVersionTallyRequest) String() string { return proto.CompactTextString(m) }
func (*QueryVersionTallyRequest) ProtoMessage()    {}
func (*QueryVersionTallyRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_7af24246367e432c, []int{0}
}
func (m *QueryVersionTallyRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryVersionTallyRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryVersionTallyRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryVersionTallyRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryVersionTallyRequest.Merge(m, src)
}
func (m *QueryVersionTallyRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryVersionTallyRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryVersionTallyRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryVersionTallyRequest proto.InternalMessageInfo

func (m *QueryVersionTallyRequest) GetVersion() uint64 {
	if m != nil {
		return m.Version
	}
	return 0
}

// QueryVersionTallyResponse is the response type for the VersionTally query.
type QueryVersionTallyResponse struct {
	VotingPower      uint64 `protobuf:"varint,1,opt,name=voting_power,json=votingPower,proto3" json:"voting_power,omitempty"`
	ThresholdPower   uint64 `protobuf:"varint,2,opt,name=threshold_power,json=thresholdPower,proto3" json:"threshold_power,omitempty"`
	TotalVotingPower uint64 `protobuf:"varint,3,opt,name=total_voting_power,json=totalVotingPower,proto3" json:"total_voting_power,omitempty"`
}

func (m *QueryVersionTallyResponse) Reset()         { *m = QueryVersionTallyResponse{} }
func (m *QueryVersionTallyResponse) String() string { return proto.CompactTextString(m) }
func (*QueryVersionTallyResponse) ProtoMessage()    {}
func (*QueryVersionTallyResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_7af24246367e432c, []int{1}
}
func (m *QueryVersionTallyResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryVersionTallyResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryVersionTallyResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryVersionTallyResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryVersionTallyResponse.Merge(m, src)
}
func (m *QueryVersionTallyResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryVersionTallyResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryVersionTallyResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryVersionTallyResponse proto.InternalMessageInfo

func (m *QueryVersionTallyResponse) GetVotingPower() uint64 {
	if m != nil {
		return m.VotingPower
	}
	return 0
}

func (m *QueryVersionTallyResponse) GetThresholdPower() uint64 {
	if m != nil {
		return m.ThresholdPower
	}
	return 0
}

func (m *QueryVersionTallyResponse) GetTotalVotingPower() uint64 {
	if m != nil {
		return m.TotalVotingPower
	}
	return 0
}

func init() {
	proto.RegisterType((*QueryVersionTallyRequest)(nil), "celestia.signal.v1.QueryVersionTallyRequest")
	proto.RegisterType((*QueryVersionTallyResponse)(nil), "celestia.signal.v1.QueryVersionTallyResponse")
}

func init() { proto.RegisterFile("celestia/signal/v1/query.proto", fileDescriptor_7af24246367e432c) }

var fileDescriptor_7af24246367e432c = []byte{
	// 337 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x91, 0xcd, 0x4e, 0x3a, 0x31,
	0x14, 0xc5, 0x29, 0xff, 0x0f, 0x93, 0x4a, 0xd4, 0x74, 0x85, 0xa8, 0x13, 0xc5, 0x85, 0x2e, 0x60,
	0x1a, 0xd0, 0x27, 0x70, 0xed, 0x42, 0x89, 0x61, 0xe1, 0x86, 0x14, 0x68, 0x4a, 0x93, 0xda, 0x5b,
	0xda, 0xce, 0x28, 0x31, 0x6e, 0x7c, 0x02, 0x12, 0xe3, 0xc6, 0x27, 0x72, 0x49, 0xe2, 0xc6, 0xa5,
	0x01, 0x1f, 0xc4, 0xcc, 0x0c, 0x83, 0x1a, 0x35, 0x71, 0x37, 0x73, 0xee, 0xef, 0x9e, 0xde, 0x7b,
	0x0f, 0x0e, 0x7a, 0x5c, 0x71, 0xe7, 0x25, 0xa3, 0x4e, 0x0a, 0xcd, 0x14, 0x8d, 0x1b, 0x74, 0x18,
	0x71, 0x3b, 0x0a, 0x8d, 0x05, 0x0f, 0x84, 0xe4, 0xf5, 0x30, 0xab, 0x87, 0x71, 0xa3, 0xb2, 0x29,
	0x00, 0x84, 0xe2, 0x94, 0x19, 0x49, 0x99, 0xd6, 0xe0, 0x99, 0x97, 0xa0, 0x5d, 0xd6, 0x51, 0x3d,
	0xc4, 0xe5, 0xd3, 0xc4, 0xa0, 0xcd, 0xad, 0x93, 0xa0, 0xcf, 0x98, 0x52, 0xa3, 0x16, 0x1f, 0x46,
	0xdc, 0x79, 0x52, 0xc6, 0x4b, 0x71, 0x26, 0x97, 0xd1, 0x36, 0xda, 0xff, 0xdb, 0xca, 0x7f, 0xab,
	0xf7, 0x08, 0xaf, 0x7f, 0xd3, 0xe6, 0x0c, 0x68, 0xc7, 0xc9, 0x0e, 0x2e, 0xc5, 0xe0, 0xa5, 0x16,
	0x1d, 0x03, 0x97, 0xdc, 0xce, 0x9b, 0x97, 0x33, 0xed, 0x24, 0x91, 0xc8, 0x1e, 0x5e, 0xf5, 0x03,
	0xcb, 0xdd, 0x00, 0x54, 0x7f, 0x4e, 0x15, 0x53, 0x6a, 0x65, 0x21, 0x67, 0x60, 0x0d, 0x13, 0x0f,
	0x9e, 0xa9, 0xce, 0x27, 0xc7, 0x3f, 0x29, 0xbb, 0x96, 0x56, 0xda, 0xef, 0xb6, 0xcd, 0x07, 0x84,
	0xff, 0xa5, 0x73, 0x91, 0x31, 0xc2, 0xa5, 0x8f, 0xc3, 0x91, 0x5a, 0xf8, 0xf5, 0x36, 0xe1, 0x4f,
	0xab, 0x57, 0xea, 0xbf, 0xa4, 0xb3, 0x8d, 0xab, 0xbb, 0xb7, 0x4f, 0xaf, 0x77, 0xc5, 0x2d, 0xb2,
	0x41, 0x23, 0x23, 0x2c, 0xeb, 0xf3, 0x24, 0x18, 0x9f, 0x20, 0xf4, 0x7a, 0x7e, 0xb3, 0x9b, 0xa3,
	0xe3, 0xc7, 0x69, 0x80, 0x26, 0xd3, 0x00, 0xbd, 0x4c, 0x03, 0x34, 0x9e, 0x05, 0x85, 0xc9, 0x2c,
	0x28, 0x3c, 0xcf, 0x82, 0xc2, 0x79, 0x53, 0x48, 0x3f, 0x88, 0xba, 0x61, 0x0f, 0x2e, 0x68, 0xfe,
	0x2e, 0x58, 0xb1, 0xf8, 0xae, 0x33, 0x63, 0xe8, 0x55, 0x9e, 0xb9, 0x1f, 0x19, 0xee, 0xba, 0xff,
	0xd3, 0xfc, 0x0e, 0xde, 0x02, 0x00, 0x00, 0xff, 0xff, 0xe7, 0x5f, 0x33, 0x08, 0x13, 0x02, 0x00,
	0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// QueryClient is the client API for Query service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type QueryClient interface {
	// VersionTally enables a client to query for the tally of voting power has
	// signalled for a particular version.
	VersionTally(ctx context.Context, in *QueryVersionTallyRequest, opts ...grpc.CallOption) (*QueryVersionTallyResponse, error)
}

type queryClient struct {
	cc grpc1.ClientConn
}

func NewQueryClient(cc grpc1.ClientConn) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) VersionTally(ctx context.Context, in *QueryVersionTallyRequest, opts ...grpc.CallOption) (*QueryVersionTallyResponse, error) {
	out := new(QueryVersionTallyResponse)
	err := c.cc.Invoke(ctx, "/celestia.signal.v1.Query/VersionTally", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
type QueryServer interface {
	// VersionTally enables a client to query for the tally of voting power has
	// signalled for a particular version.
	VersionTally(context.Context, *QueryVersionTallyRequest) (*QueryVersionTallyResponse, error)
}

// UnimplementedQueryServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (*UnimplementedQueryServer) VersionTally(ctx context.Context, req *QueryVersionTallyRequest) (*QueryVersionTallyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method VersionTally not implemented")
}

func RegisterQueryServer(s grpc1.Server, srv QueryServer) {
	s.RegisterService(&_Query_serviceDesc, srv)
}

func _Query_VersionTally_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryVersionTallyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).VersionTally(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/celestia.signal.v1.Query/VersionTally",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).VersionTally(ctx, req.(*QueryVersionTallyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "celestia.signal.v1.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "VersionTally",
			Handler:    _Query_VersionTally_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "celestia/signal/v1/query.proto",
}

func (m *QueryVersionTallyRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryVersionTallyRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryVersionTallyRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Version != 0 {
		i = encodeVarintQuery(dAtA, i, uint64(m.Version))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *QueryVersionTallyResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryVersionTallyResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryVersionTallyResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.TotalVotingPower != 0 {
		i = encodeVarintQuery(dAtA, i, uint64(m.TotalVotingPower))
		i--
		dAtA[i] = 0x18
	}
	if m.ThresholdPower != 0 {
		i = encodeVarintQuery(dAtA, i, uint64(m.ThresholdPower))
		i--
		dAtA[i] = 0x10
	}
	if m.VotingPower != 0 {
		i = encodeVarintQuery(dAtA, i, uint64(m.VotingPower))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintQuery(dAtA []byte, offset int, v uint64) int {
	offset -= sovQuery(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *QueryVersionTallyRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Version != 0 {
		n += 1 + sovQuery(uint64(m.Version))
	}
	return n
}

func (m *QueryVersionTallyResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.VotingPower != 0 {
		n += 1 + sovQuery(uint64(m.VotingPower))
	}
	if m.ThresholdPower != 0 {
		n += 1 + sovQuery(uint64(m.ThresholdPower))
	}
	if m.TotalVotingPower != 0 {
		n += 1 + sovQuery(uint64(m.TotalVotingPower))
	}
	return n
}

func sovQuery(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozQuery(x uint64) (n int) {
	return sovQuery(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *QueryVersionTallyRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryVersionTallyRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryVersionTallyRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Version", wireType)
			}
			m.Version = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Version |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func (m *QueryVersionTallyResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryVersionTallyResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryVersionTallyResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field VotingPower", wireType)
			}
			m.VotingPower = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.VotingPower |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ThresholdPower", wireType)
			}
			m.ThresholdPower = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ThresholdPower |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TotalVotingPower", wireType)
			}
			m.TotalVotingPower = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.TotalVotingPower |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func skipQuery(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowQuery
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
					return 0, ErrIntOverflowQuery
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
					return 0, ErrIntOverflowQuery
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
				return 0, ErrInvalidLengthQuery
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupQuery
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthQuery
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthQuery        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowQuery          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupQuery = fmt.Errorf("proto: unexpected end of group")
)
