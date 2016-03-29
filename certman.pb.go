// Code generated by protoc-gen-gogo.
// source: certman.proto
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

type ServiceInfo struct {
	DomainName string `protobuf:"bytes,1,opt,name=domain_name,json=domainName,proto3" json:"domain_name,omitempty"`
}

func (m *ServiceInfo) Reset()                    { *m = ServiceInfo{} }
func (m *ServiceInfo) String() string            { return proto.CompactTextString(m) }
func (*ServiceInfo) ProtoMessage()               {}
func (*ServiceInfo) Descriptor() ([]byte, []int) { return fileDescriptorCertman, []int{0} }

type Certificate struct {
	DomainName string `protobuf:"bytes,1,opt,name=domain_name,json=domainName,proto3" json:"domain_name,omitempty"`
	Cert       []byte `protobuf:"bytes,2,opt,name=cert,proto3" json:"cert,omitempty"`
	Key        []byte `protobuf:"bytes,3,opt,name=key,proto3" json:"key,omitempty"`
	ExpiresAt  int64  `protobuf:"varint,4,opt,name=expires_at,json=expiresAt,proto3" json:"expires_at,omitempty"`
}

func (m *Certificate) Reset()                    { *m = Certificate{} }
func (m *Certificate) String() string            { return proto.CompactTextString(m) }
func (*Certificate) ProtoMessage()               {}
func (*Certificate) Descriptor() ([]byte, []int) { return fileDescriptorCertman, []int{1} }

func init() {
	proto.RegisterType((*ServiceInfo)(nil), "apipb.ServiceInfo")
	proto.RegisterType((*Certificate)(nil), "apipb.Certificate")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// Client API for CertificateManager service

type CertificateManagerClient interface {
	Get(ctx context.Context, in *ServiceInfo, opts ...grpc.CallOption) (*Certificate, error)
}

type certificateManagerClient struct {
	cc *grpc.ClientConn
}

func NewCertificateManagerClient(cc *grpc.ClientConn) CertificateManagerClient {
	return &certificateManagerClient{cc}
}

func (c *certificateManagerClient) Get(ctx context.Context, in *ServiceInfo, opts ...grpc.CallOption) (*Certificate, error) {
	out := new(Certificate)
	err := grpc.Invoke(ctx, "/apipb.CertificateManager/Get", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for CertificateManager service

type CertificateManagerServer interface {
	Get(context.Context, *ServiceInfo) (*Certificate, error)
}

func RegisterCertificateManagerServer(s *grpc.Server, srv CertificateManagerServer) {
	s.RegisterService(&_CertificateManager_serviceDesc, srv)
}

func _CertificateManager_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(ServiceInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(CertificateManagerServer).Get(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

var _CertificateManager_serviceDesc = grpc.ServiceDesc{
	ServiceName: "apipb.CertificateManager",
	HandlerType: (*CertificateManagerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _CertificateManager_Get_Handler,
		},
	},
	Streams: []grpc.StreamDesc{},
}

func (m *ServiceInfo) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *ServiceInfo) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.DomainName) > 0 {
		data[i] = 0xa
		i++
		i = encodeVarintCertman(data, i, uint64(len(m.DomainName)))
		i += copy(data[i:], m.DomainName)
	}
	return i, nil
}

func (m *Certificate) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *Certificate) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.DomainName) > 0 {
		data[i] = 0xa
		i++
		i = encodeVarintCertman(data, i, uint64(len(m.DomainName)))
		i += copy(data[i:], m.DomainName)
	}
	if m.Cert != nil {
		if len(m.Cert) > 0 {
			data[i] = 0x12
			i++
			i = encodeVarintCertman(data, i, uint64(len(m.Cert)))
			i += copy(data[i:], m.Cert)
		}
	}
	if m.Key != nil {
		if len(m.Key) > 0 {
			data[i] = 0x1a
			i++
			i = encodeVarintCertman(data, i, uint64(len(m.Key)))
			i += copy(data[i:], m.Key)
		}
	}
	if m.ExpiresAt != 0 {
		data[i] = 0x20
		i++
		i = encodeVarintCertman(data, i, uint64(m.ExpiresAt))
	}
	return i, nil
}

func encodeFixed64Certman(data []byte, offset int, v uint64) int {
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
func encodeFixed32Certman(data []byte, offset int, v uint32) int {
	data[offset] = uint8(v)
	data[offset+1] = uint8(v >> 8)
	data[offset+2] = uint8(v >> 16)
	data[offset+3] = uint8(v >> 24)
	return offset + 4
}
func encodeVarintCertman(data []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		data[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	data[offset] = uint8(v)
	return offset + 1
}
func (m *ServiceInfo) Size() (n int) {
	var l int
	_ = l
	l = len(m.DomainName)
	if l > 0 {
		n += 1 + l + sovCertman(uint64(l))
	}
	return n
}

func (m *Certificate) Size() (n int) {
	var l int
	_ = l
	l = len(m.DomainName)
	if l > 0 {
		n += 1 + l + sovCertman(uint64(l))
	}
	if m.Cert != nil {
		l = len(m.Cert)
		if l > 0 {
			n += 1 + l + sovCertman(uint64(l))
		}
	}
	if m.Key != nil {
		l = len(m.Key)
		if l > 0 {
			n += 1 + l + sovCertman(uint64(l))
		}
	}
	if m.ExpiresAt != 0 {
		n += 1 + sovCertman(uint64(m.ExpiresAt))
	}
	return n
}

func sovCertman(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozCertman(x uint64) (n int) {
	return sovCertman(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ServiceInfo) Unmarshal(data []byte) error {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCertman
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
			return fmt.Errorf("proto: ServiceInfo: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ServiceInfo: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DomainName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCertman
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
				return ErrInvalidLengthCertman
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DomainName = string(data[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipCertman(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthCertman
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
func (m *Certificate) Unmarshal(data []byte) error {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCertman
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
			return fmt.Errorf("proto: Certificate: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Certificate: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DomainName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCertman
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
				return ErrInvalidLengthCertman
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DomainName = string(data[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Cert", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCertman
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthCertman
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Cert = append(m.Cert[:0], data[iNdEx:postIndex]...)
			if m.Cert == nil {
				m.Cert = []byte{}
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Key", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCertman
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthCertman
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Key = append(m.Key[:0], data[iNdEx:postIndex]...)
			if m.Key == nil {
				m.Key = []byte{}
			}
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ExpiresAt", wireType)
			}
			m.ExpiresAt = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCertman
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				m.ExpiresAt |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipCertman(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthCertman
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
func skipCertman(data []byte) (n int, err error) {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowCertman
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
					return 0, ErrIntOverflowCertman
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
					return 0, ErrIntOverflowCertman
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
				return 0, ErrInvalidLengthCertman
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowCertman
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
				next, err := skipCertman(data[start:])
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
	ErrInvalidLengthCertman = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowCertman   = fmt.Errorf("proto: integer overflow")
)

var fileDescriptorCertman = []byte{
	// 258 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0xe2, 0x4d, 0x4e, 0x2d, 0x2a,
	0xc9, 0x4d, 0xcc, 0xd3, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x4d, 0x2c, 0xc8, 0x2c, 0x48,
	0x92, 0xd2, 0x4d, 0xcf, 0x2c, 0xc9, 0x28, 0x4d, 0xd2, 0x4b, 0xce, 0xcf, 0xd5, 0x4f, 0xcf, 0x4f,
	0xcf, 0xd7, 0x07, 0xcb, 0x26, 0x95, 0xa6, 0x81, 0x79, 0x60, 0x0e, 0x98, 0x05, 0xd1, 0xa5, 0xa4,
	0xc7, 0xc5, 0x1d, 0x9c, 0x5a, 0x54, 0x96, 0x99, 0x9c, 0xea, 0x99, 0x97, 0x96, 0x2f, 0x24, 0xcf,
	0xc5, 0x9d, 0x92, 0x9f, 0x9b, 0x98, 0x99, 0x17, 0x9f, 0x97, 0x98, 0x9b, 0x2a, 0xc1, 0xa8, 0xc0,
	0xa8, 0xc1, 0x19, 0xc4, 0x05, 0x11, 0xf2, 0x03, 0x8a, 0x28, 0x15, 0x73, 0x71, 0x3b, 0x03, 0xad,
	0xcd, 0x4c, 0xcb, 0x4c, 0x4e, 0x2c, 0x49, 0x25, 0xa8, 0x5e, 0x48, 0x88, 0x8b, 0x05, 0xe4, 0x4c,
	0x09, 0x26, 0xa0, 0x0c, 0x4f, 0x10, 0x98, 0x2d, 0x24, 0xc0, 0xc5, 0x9c, 0x9d, 0x5a, 0x29, 0xc1,
	0x0c, 0x16, 0x02, 0x31, 0x85, 0x64, 0xb9, 0xb8, 0x52, 0x2b, 0x0a, 0x32, 0x8b, 0x52, 0x8b, 0xe3,
	0x13, 0x4b, 0x24, 0x58, 0x80, 0x12, 0xcc, 0x41, 0x9c, 0x50, 0x11, 0xc7, 0x12, 0x23, 0x67, 0x2e,
	0x21, 0x24, 0x4b, 0x7d, 0x13, 0xf3, 0x12, 0xd3, 0x53, 0x8b, 0x84, 0x74, 0xb9, 0x98, 0xdd, 0x53,
	0x4b, 0x84, 0x84, 0xf4, 0xc0, 0x1e, 0xd7, 0x43, 0xf2, 0x86, 0x14, 0x4c, 0x0c, 0x49, 0x97, 0x93,
	0xd2, 0x89, 0x87, 0x72, 0x0c, 0x17, 0x80, 0xf8, 0xc4, 0x23, 0x39, 0xc6, 0x0b, 0x40, 0xfc, 0x00,
	0x88, 0x27, 0x3c, 0x96, 0x63, 0x88, 0xe2, 0xc8, 0x2f, 0x29, 0xce, 0xcc, 0xcd, 0x2f, 0x48, 0x4a,
	0x62, 0x03, 0x07, 0x8a, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0x65, 0x9c, 0xb0, 0x30, 0x5b, 0x01,
	0x00, 0x00,
}
