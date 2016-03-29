// Code generated by protoc-gen-gogo.
// source: watch.proto
// DO NOT EDIT!

package otsimopb

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type WatchEvent_EventType int32

const (
	PROFILE_UPDATED     WatchEvent_EventType = 0
	CHILD_UPDATED       WatchEvent_EventType = 1
	CHILD_GAMES_UPDATED WatchEvent_EventType = 2
	CHILD_SOUND_UPDATED WatchEvent_EventType = 3
)

var WatchEvent_EventType_name = map[int32]string{
	0: "PROFILE_UPDATED",
	1: "CHILD_UPDATED",
	2: "CHILD_GAMES_UPDATED",
	3: "CHILD_SOUND_UPDATED",
}
var WatchEvent_EventType_value = map[string]int32{
	"PROFILE_UPDATED":     0,
	"CHILD_UPDATED":       1,
	"CHILD_GAMES_UPDATED": 2,
	"CHILD_SOUND_UPDATED": 3,
}

func (x WatchEvent_EventType) String() string {
	return proto.EnumName(WatchEvent_EventType_name, int32(x))
}
func (WatchEvent_EventType) EnumDescriptor() ([]byte, []int) { return fileDescriptorWatch, []int{3, 0} }

type EmitRequest struct {
	ProfileId string      `protobuf:"bytes,1,opt,name=profile_id,json=profileId,proto3" json:"profile_id,omitempty"`
	Event     *WatchEvent `protobuf:"bytes,2,opt,name=event" json:"event,omitempty"`
}

func (m *EmitRequest) Reset()                    { *m = EmitRequest{} }
func (m *EmitRequest) String() string            { return proto.CompactTextString(m) }
func (*EmitRequest) ProtoMessage()               {}
func (*EmitRequest) Descriptor() ([]byte, []int) { return fileDescriptorWatch, []int{0} }

type EmitResponse struct {
}

func (m *EmitResponse) Reset()                    { *m = EmitResponse{} }
func (m *EmitResponse) String() string            { return proto.CompactTextString(m) }
func (*EmitResponse) ProtoMessage()               {}
func (*EmitResponse) Descriptor() ([]byte, []int) { return fileDescriptorWatch, []int{1} }

type WatchRequest struct {
	// profile id is for Create request
	ProfileId string `protobuf:"bytes,2,opt,name=profile_id,json=profileId,proto3" json:"profile_id,omitempty"`
}

func (m *WatchRequest) Reset()                    { *m = WatchRequest{} }
func (m *WatchRequest) String() string            { return proto.CompactTextString(m) }
func (*WatchRequest) ProtoMessage()               {}
func (*WatchRequest) Descriptor() ([]byte, []int) { return fileDescriptorWatch, []int{2} }

type WatchEvent struct {
	Type      WatchEvent_EventType `protobuf:"varint,1,opt,name=type,proto3,enum=apipb.WatchEvent_EventType" json:"type,omitempty"`
	ProfileId string               `protobuf:"bytes,2,opt,name=profile_id,json=profileId,proto3" json:"profile_id,omitempty"`
	ChildId   string               `protobuf:"bytes,3,opt,name=child_id,json=childId,proto3" json:"child_id,omitempty"`
	GameId    string               `protobuf:"bytes,4,opt,name=game_id,json=gameId,proto3" json:"game_id,omitempty"`
}

func (m *WatchEvent) Reset()                    { *m = WatchEvent{} }
func (m *WatchEvent) String() string            { return proto.CompactTextString(m) }
func (*WatchEvent) ProtoMessage()               {}
func (*WatchEvent) Descriptor() ([]byte, []int) { return fileDescriptorWatch, []int{3} }

type WatchResponse struct {
	Created  bool        `protobuf:"varint,1,opt,name=created,proto3" json:"created,omitempty"`
	Canceled bool        `protobuf:"varint,2,opt,name=canceled,proto3" json:"canceled,omitempty"`
	Event    *WatchEvent `protobuf:"bytes,3,opt,name=event" json:"event,omitempty"`
}

func (m *WatchResponse) Reset()                    { *m = WatchResponse{} }
func (m *WatchResponse) String() string            { return proto.CompactTextString(m) }
func (*WatchResponse) ProtoMessage()               {}
func (*WatchResponse) Descriptor() ([]byte, []int) { return fileDescriptorWatch, []int{4} }

func init() {
	proto.RegisterType((*EmitRequest)(nil), "apipb.EmitRequest")
	proto.RegisterType((*EmitResponse)(nil), "apipb.EmitResponse")
	proto.RegisterType((*WatchRequest)(nil), "apipb.WatchRequest")
	proto.RegisterType((*WatchEvent)(nil), "apipb.WatchEvent")
	proto.RegisterType((*WatchResponse)(nil), "apipb.WatchResponse")
	proto.RegisterEnum("apipb.WatchEvent_EventType", WatchEvent_EventType_name, WatchEvent_EventType_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// Client API for WatchService service

type WatchServiceClient interface {
	Emit(ctx context.Context, in *EmitRequest, opts ...grpc.CallOption) (*EmitResponse, error)
	Watch(ctx context.Context, in *WatchRequest, opts ...grpc.CallOption) (WatchService_WatchClient, error)
}

type watchServiceClient struct {
	cc *grpc.ClientConn
}

func NewWatchServiceClient(cc *grpc.ClientConn) WatchServiceClient {
	return &watchServiceClient{cc}
}

func (c *watchServiceClient) Emit(ctx context.Context, in *EmitRequest, opts ...grpc.CallOption) (*EmitResponse, error) {
	out := new(EmitResponse)
	err := grpc.Invoke(ctx, "/apipb.WatchService/Emit", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *watchServiceClient) Watch(ctx context.Context, in *WatchRequest, opts ...grpc.CallOption) (WatchService_WatchClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_WatchService_serviceDesc.Streams[0], c.cc, "/apipb.WatchService/Watch", opts...)
	if err != nil {
		return nil, err
	}
	x := &watchServiceWatchClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type WatchService_WatchClient interface {
	Recv() (*WatchResponse, error)
	grpc.ClientStream
}

type watchServiceWatchClient struct {
	grpc.ClientStream
}

func (x *watchServiceWatchClient) Recv() (*WatchResponse, error) {
	m := new(WatchResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for WatchService service

type WatchServiceServer interface {
	Emit(context.Context, *EmitRequest) (*EmitResponse, error)
	Watch(*WatchRequest, WatchService_WatchServer) error
}

func RegisterWatchServiceServer(s *grpc.Server, srv WatchServiceServer) {
	s.RegisterService(&_WatchService_serviceDesc, srv)
}

func _WatchService_Emit_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(EmitRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(WatchServiceServer).Emit(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _WatchService_Watch_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(WatchRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(WatchServiceServer).Watch(m, &watchServiceWatchServer{stream})
}

type WatchService_WatchServer interface {
	Send(*WatchResponse) error
	grpc.ServerStream
}

type watchServiceWatchServer struct {
	grpc.ServerStream
}

func (x *watchServiceWatchServer) Send(m *WatchResponse) error {
	return x.ServerStream.SendMsg(m)
}

var _WatchService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "apipb.WatchService",
	HandlerType: (*WatchServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Emit",
			Handler:    _WatchService_Emit_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Watch",
			Handler:       _WatchService_Watch_Handler,
			ServerStreams: true,
		},
	},
}

func (m *EmitRequest) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *EmitRequest) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.ProfileId) > 0 {
		data[i] = 0xa
		i++
		i = encodeVarintWatch(data, i, uint64(len(m.ProfileId)))
		i += copy(data[i:], m.ProfileId)
	}
	if m.Event != nil {
		data[i] = 0x12
		i++
		i = encodeVarintWatch(data, i, uint64(m.Event.Size()))
		n1, err := m.Event.MarshalTo(data[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	return i, nil
}

func (m *EmitResponse) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *EmitResponse) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	return i, nil
}

func (m *WatchRequest) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *WatchRequest) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.ProfileId) > 0 {
		data[i] = 0x12
		i++
		i = encodeVarintWatch(data, i, uint64(len(m.ProfileId)))
		i += copy(data[i:], m.ProfileId)
	}
	return i, nil
}

func (m *WatchEvent) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *WatchEvent) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Type != 0 {
		data[i] = 0x8
		i++
		i = encodeVarintWatch(data, i, uint64(m.Type))
	}
	if len(m.ProfileId) > 0 {
		data[i] = 0x12
		i++
		i = encodeVarintWatch(data, i, uint64(len(m.ProfileId)))
		i += copy(data[i:], m.ProfileId)
	}
	if len(m.ChildId) > 0 {
		data[i] = 0x1a
		i++
		i = encodeVarintWatch(data, i, uint64(len(m.ChildId)))
		i += copy(data[i:], m.ChildId)
	}
	if len(m.GameId) > 0 {
		data[i] = 0x22
		i++
		i = encodeVarintWatch(data, i, uint64(len(m.GameId)))
		i += copy(data[i:], m.GameId)
	}
	return i, nil
}

func (m *WatchResponse) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *WatchResponse) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Created {
		data[i] = 0x8
		i++
		if m.Created {
			data[i] = 1
		} else {
			data[i] = 0
		}
		i++
	}
	if m.Canceled {
		data[i] = 0x10
		i++
		if m.Canceled {
			data[i] = 1
		} else {
			data[i] = 0
		}
		i++
	}
	if m.Event != nil {
		data[i] = 0x1a
		i++
		i = encodeVarintWatch(data, i, uint64(m.Event.Size()))
		n2, err := m.Event.MarshalTo(data[i:])
		if err != nil {
			return 0, err
		}
		i += n2
	}
	return i, nil
}

func encodeFixed64Watch(data []byte, offset int, v uint64) int {
	data[offset] = uint8(v)
	data[offset+1] = uint8(v >> 8)
	data[offset+2] = uint8(v >> 16)
	data[offset+3] = uint8(v >> 24)
	data[offset+4] = uint8(v >> 32)
	data[offset+5] = uint8(v >> 40)
	data[offset+6] = uint8(v >> 48)
	data[offset+7] = uint8(v >> 56)
	return offset + 8
}
func encodeFixed32Watch(data []byte, offset int, v uint32) int {
	data[offset] = uint8(v)
	data[offset+1] = uint8(v >> 8)
	data[offset+2] = uint8(v >> 16)
	data[offset+3] = uint8(v >> 24)
	return offset + 4
}
func encodeVarintWatch(data []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		data[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	data[offset] = uint8(v)
	return offset + 1
}
func (m *EmitRequest) Size() (n int) {
	var l int
	_ = l
	l = len(m.ProfileId)
	if l > 0 {
		n += 1 + l + sovWatch(uint64(l))
	}
	if m.Event != nil {
		l = m.Event.Size()
		n += 1 + l + sovWatch(uint64(l))
	}
	return n
}

func (m *EmitResponse) Size() (n int) {
	var l int
	_ = l
	return n
}

func (m *WatchRequest) Size() (n int) {
	var l int
	_ = l
	l = len(m.ProfileId)
	if l > 0 {
		n += 1 + l + sovWatch(uint64(l))
	}
	return n
}

func (m *WatchEvent) Size() (n int) {
	var l int
	_ = l
	if m.Type != 0 {
		n += 1 + sovWatch(uint64(m.Type))
	}
	l = len(m.ProfileId)
	if l > 0 {
		n += 1 + l + sovWatch(uint64(l))
	}
	l = len(m.ChildId)
	if l > 0 {
		n += 1 + l + sovWatch(uint64(l))
	}
	l = len(m.GameId)
	if l > 0 {
		n += 1 + l + sovWatch(uint64(l))
	}
	return n
}

func (m *WatchResponse) Size() (n int) {
	var l int
	_ = l
	if m.Created {
		n += 2
	}
	if m.Canceled {
		n += 2
	}
	if m.Event != nil {
		l = m.Event.Size()
		n += 1 + l + sovWatch(uint64(l))
	}
	return n
}

func sovWatch(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozWatch(x uint64) (n int) {
	return sovWatch(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *EmitRequest) Unmarshal(data []byte) error {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowWatch
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: EmitRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EmitRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ProfileId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWatch
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthWatch
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ProfileId = string(data[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Event", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWatch
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthWatch
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Event == nil {
				m.Event = &WatchEvent{}
			}
			if err := m.Event.Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipWatch(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthWatch
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
func (m *EmitResponse) Unmarshal(data []byte) error {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowWatch
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: EmitResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EmitResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipWatch(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthWatch
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
func (m *WatchRequest) Unmarshal(data []byte) error {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowWatch
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: WatchRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: WatchRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ProfileId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWatch
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthWatch
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ProfileId = string(data[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipWatch(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthWatch
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
func (m *WatchEvent) Unmarshal(data []byte) error {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowWatch
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: WatchEvent: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: WatchEvent: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Type", wireType)
			}
			m.Type = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWatch
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				m.Type |= (WatchEvent_EventType(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ProfileId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWatch
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthWatch
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ProfileId = string(data[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ChildId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWatch
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthWatch
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ChildId = string(data[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field GameId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWatch
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthWatch
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.GameId = string(data[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipWatch(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthWatch
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
func (m *WatchResponse) Unmarshal(data []byte) error {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowWatch
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: WatchResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: WatchResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Created", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWatch
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				v |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Created = bool(v != 0)
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Canceled", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWatch
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				v |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Canceled = bool(v != 0)
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Event", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWatch
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthWatch
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Event == nil {
				m.Event = &WatchEvent{}
			}
			if err := m.Event.Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipWatch(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthWatch
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
func skipWatch(data []byte) (n int, err error) {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowWatch
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
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
					return 0, ErrIntOverflowWatch
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if data[iNdEx-1] < 0x80 {
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
					return 0, ErrIntOverflowWatch
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthWatch
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowWatch
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := data[iNdEx]
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
				next, err := skipWatch(data[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
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
	ErrInvalidLengthWatch = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowWatch   = fmt.Errorf("proto: integer overflow")
)

var fileDescriptorWatch = []byte{
	// 430 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x84, 0x52, 0xdd, 0x8e, 0xd2, 0x40,
	0x18, 0xa5, 0xed, 0xb2, 0x94, 0x8f, 0xdd, 0xd5, 0x9d, 0x9a, 0x88, 0x18, 0x89, 0xe9, 0x8d, 0x7b,
	0xb3, 0xc5, 0xa0, 0x2f, 0xb0, 0x4a, 0x55, 0x92, 0x55, 0x36, 0x2d, 0xc4, 0xc4, 0x1b, 0xd2, 0x96,
	0xa1, 0x4c, 0x42, 0x3b, 0x95, 0x4e, 0x31, 0xbe, 0x85, 0xcf, 0xe0, 0xd3, 0x70, 0xc9, 0x23, 0xf8,
	0xf3, 0x10, 0xde, 0x3a, 0x9d, 0x29, 0x2d, 0xa8, 0x71, 0x2f, 0xa6, 0xe9, 0x39, 0x67, 0xce, 0x77,
	0xbe, 0x99, 0x6f, 0xa0, 0xf5, 0xc9, 0x63, 0xc1, 0xc2, 0x4a, 0x56, 0x94, 0x51, 0x54, 0xf7, 0x12,
	0x92, 0xf8, 0x9d, 0xcb, 0x90, 0xb0, 0x45, 0xe6, 0x5b, 0x01, 0x8d, 0x7a, 0x21, 0x0d, 0x69, 0x4f,
	0xa8, 0x7e, 0x36, 0x17, 0x48, 0x00, 0xf1, 0x27, 0x5d, 0xe6, 0x04, 0x5a, 0x76, 0x44, 0x98, 0x83,
	0x3f, 0x66, 0x38, 0x65, 0xe8, 0x11, 0x00, 0xe7, 0xe7, 0x64, 0x89, 0xa7, 0x64, 0xd6, 0x56, 0x1e,
	0x2b, 0x17, 0x4d, 0xa7, 0x59, 0x30, 0xc3, 0x19, 0x7a, 0x02, 0x75, 0xbc, 0xc6, 0x31, 0x6b, 0xab,
	0x5c, 0x69, 0xf5, 0xcf, 0x2d, 0x91, 0x69, 0xbd, 0xcf, 0xdb, 0xb0, 0x73, 0xc1, 0x91, 0xba, 0x79,
	0x06, 0x27, 0xb2, 0x6c, 0x9a, 0xd0, 0x38, 0xc5, 0xe6, 0x25, 0x9c, 0x88, 0x4d, 0xff, 0xce, 0x51,
	0xff, 0xc8, 0x31, 0x7f, 0x29, 0x00, 0x55, 0x51, 0xd4, 0x83, 0x23, 0xf6, 0x39, 0xc1, 0xa2, 0x9f,
	0xb3, 0xfe, 0xc3, 0xbf, 0x52, 0x2d, 0xf1, 0x1d, 0xf3, 0x2d, 0x8e, 0xd8, 0x78, 0x4b, 0x79, 0xf4,
	0x00, 0xf4, 0x60, 0x41, 0x96, 0xb3, 0x5c, 0xd4, 0x84, 0xd8, 0x10, 0x98, 0x4b, 0xf7, 0xa1, 0x11,
	0x7a, 0x91, 0xb0, 0x1d, 0x09, 0xe5, 0x38, 0x87, 0xbc, 0x25, 0x0c, 0xcd, 0x32, 0x05, 0x19, 0x70,
	0xe7, 0xc6, 0x19, 0xbd, 0x1a, 0x5e, 0xdb, 0xd3, 0xc9, 0xcd, 0xe0, 0x6a, 0x6c, 0x0f, 0xee, 0xd6,
	0xd0, 0x39, 0x9c, 0xbe, 0x7c, 0x33, 0xbc, 0x1e, 0x94, 0x94, 0xc2, 0xab, 0x19, 0x92, 0x7a, 0x7d,
	0xf5, 0xd6, 0x76, 0x4b, 0x41, 0xad, 0x04, 0x77, 0x34, 0x79, 0x57, 0x39, 0x34, 0x33, 0x86, 0xd3,
	0xe2, 0xa2, 0xe4, 0xcd, 0xa1, 0x36, 0x34, 0x82, 0x15, 0xf6, 0x18, 0x96, 0xe3, 0xd0, 0x9d, 0x1d,
	0x44, 0x1d, 0x7e, 0x0a, 0x2f, 0x0e, 0xf0, 0x12, 0xcb, 0x23, 0xea, 0x4e, 0x89, 0xab, 0x41, 0x69,
	0xff, 0x1f, 0x54, 0x3f, 0x2b, 0x06, 0xe3, 0xe2, 0xd5, 0x9a, 0x04, 0x38, 0xbf, 0xea, 0x7c, 0x70,
	0x08, 0x15, 0x8e, 0xbd, 0xc7, 0xd1, 0x31, 0x0e, 0xb8, 0xa2, 0xbf, 0xe7, 0x50, 0x17, 0x05, 0x90,
	0xb1, 0x9f, 0xb1, 0xb3, 0xdc, 0x3b, 0x24, 0xa5, 0xe7, 0xa9, 0xf2, 0xe2, 0x62, 0xf3, 0xbd, 0x5b,
	0xdb, 0xf2, 0xb5, 0xf9, 0xd1, 0x55, 0xb6, 0x7c, 0x7d, 0xe3, 0xeb, 0xcb, 0xcf, 0x6e, 0xed, 0x83,
	0x4e, 0x59, 0x4a, 0x22, 0x9a, 0xf8, 0x5f, 0x55, 0x6d, 0x34, 0x76, 0xfd, 0x63, 0xf1, 0x4e, 0x9f,
	0xfd, 0x0e, 0x00, 0x00, 0xff, 0xff, 0x5d, 0xaf, 0x7c, 0x4a, 0xec, 0x02, 0x00, 0x00,
}
