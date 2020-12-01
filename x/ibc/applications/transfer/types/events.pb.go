// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: ibc/applications/transfer/v1/events.proto

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

// EventTransfer is a typed event emitted on ibc transfer
type EventTransfer struct {
	Sender   string `protobuf:"bytes,1,opt,name=sender,proto3" json:"sender,omitempty"`
	Receiver string `protobuf:"bytes,2,opt,name=receiver,proto3" json:"receiver,omitempty"`
}

func (m *EventTransfer) Reset()         { *m = EventTransfer{} }
func (m *EventTransfer) String() string { return proto.CompactTextString(m) }
func (*EventTransfer) ProtoMessage()    {}
func (*EventTransfer) Descriptor() ([]byte, []int) {
	return fileDescriptor_0d04a14abeba54b9, []int{0}
}
func (m *EventTransfer) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EventTransfer) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EventTransfer.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EventTransfer) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventTransfer.Merge(m, src)
}
func (m *EventTransfer) XXX_Size() int {
	return m.Size()
}
func (m *EventTransfer) XXX_DiscardUnknown() {
	xxx_messageInfo_EventTransfer.DiscardUnknown(m)
}

var xxx_messageInfo_EventTransfer proto.InternalMessageInfo

func (m *EventTransfer) GetSender() string {
	if m != nil {
		return m.Sender
	}
	return ""
}

func (m *EventTransfer) GetReceiver() string {
	if m != nil {
		return m.Receiver
	}
	return ""
}

// EventAcknowledgementSuccess is a typed event emitted on packet acknowledgement success
type EventAcknowledgementSuccess struct {
	Success []byte `protobuf:"bytes,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (m *EventAcknowledgementSuccess) Reset()         { *m = EventAcknowledgementSuccess{} }
func (m *EventAcknowledgementSuccess) String() string { return proto.CompactTextString(m) }
func (*EventAcknowledgementSuccess) ProtoMessage()    {}
func (*EventAcknowledgementSuccess) Descriptor() ([]byte, []int) {
	return fileDescriptor_0d04a14abeba54b9, []int{1}
}
func (m *EventAcknowledgementSuccess) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EventAcknowledgementSuccess) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EventAcknowledgementSuccess.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EventAcknowledgementSuccess) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventAcknowledgementSuccess.Merge(m, src)
}
func (m *EventAcknowledgementSuccess) XXX_Size() int {
	return m.Size()
}
func (m *EventAcknowledgementSuccess) XXX_DiscardUnknown() {
	xxx_messageInfo_EventAcknowledgementSuccess.DiscardUnknown(m)
}

var xxx_messageInfo_EventAcknowledgementSuccess proto.InternalMessageInfo

func (m *EventAcknowledgementSuccess) GetSuccess() []byte {
	if m != nil {
		return m.Success
	}
	return nil
}

// EventAcknowledgementError is a typed event emitted on packet acknowledgement error
type EventAcknowledgementError struct {
	Error string `protobuf:"bytes,1,opt,name=error,proto3" json:"error,omitempty"`
}

func (m *EventAcknowledgementError) Reset()         { *m = EventAcknowledgementError{} }
func (m *EventAcknowledgementError) String() string { return proto.CompactTextString(m) }
func (*EventAcknowledgementError) ProtoMessage()    {}
func (*EventAcknowledgementError) Descriptor() ([]byte, []int) {
	return fileDescriptor_0d04a14abeba54b9, []int{2}
}
func (m *EventAcknowledgementError) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EventAcknowledgementError) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EventAcknowledgementError.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EventAcknowledgementError) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventAcknowledgementError.Merge(m, src)
}
func (m *EventAcknowledgementError) XXX_Size() int {
	return m.Size()
}
func (m *EventAcknowledgementError) XXX_DiscardUnknown() {
	xxx_messageInfo_EventAcknowledgementError.DiscardUnknown(m)
}

var xxx_messageInfo_EventAcknowledgementError proto.InternalMessageInfo

func (m *EventAcknowledgementError) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

// EventDenominationTrace is a typed event for denomination trace
type EventDenominationTrace struct {
	TraceHash string `protobuf:"bytes,1,opt,name=trace_hash,json=traceHash,proto3" json:"trace_hash,omitempty" yaml:"trace_hash"`
	Denom     string `protobuf:"bytes,2,opt,name=denom,proto3" json:"denom,omitempty"`
}

func (m *EventDenominationTrace) Reset()         { *m = EventDenominationTrace{} }
func (m *EventDenominationTrace) String() string { return proto.CompactTextString(m) }
func (*EventDenominationTrace) ProtoMessage()    {}
func (*EventDenominationTrace) Descriptor() ([]byte, []int) {
	return fileDescriptor_0d04a14abeba54b9, []int{3}
}
func (m *EventDenominationTrace) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EventDenominationTrace) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EventDenominationTrace.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EventDenominationTrace) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventDenominationTrace.Merge(m, src)
}
func (m *EventDenominationTrace) XXX_Size() int {
	return m.Size()
}
func (m *EventDenominationTrace) XXX_DiscardUnknown() {
	xxx_messageInfo_EventDenominationTrace.DiscardUnknown(m)
}

var xxx_messageInfo_EventDenominationTrace proto.InternalMessageInfo

func (m *EventDenominationTrace) GetTraceHash() string {
	if m != nil {
		return m.TraceHash
	}
	return ""
}

func (m *EventDenominationTrace) GetDenom() string {
	if m != nil {
		return m.Denom
	}
	return ""
}

func init() {
	proto.RegisterType((*EventTransfer)(nil), "ibc.applications.transfer.v1.EventTransfer")
	proto.RegisterType((*EventAcknowledgementSuccess)(nil), "ibc.applications.transfer.v1.EventAcknowledgementSuccess")
	proto.RegisterType((*EventAcknowledgementError)(nil), "ibc.applications.transfer.v1.EventAcknowledgementError")
	proto.RegisterType((*EventDenominationTrace)(nil), "ibc.applications.transfer.v1.EventDenominationTrace")
}

func init() {
	proto.RegisterFile("ibc/applications/transfer/v1/events.proto", fileDescriptor_0d04a14abeba54b9)
}

var fileDescriptor_0d04a14abeba54b9 = []byte{
	// 329 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x91, 0x41, 0x4b, 0xc3, 0x30,
	0x18, 0x86, 0x57, 0xc1, 0xe9, 0x82, 0x1e, 0x2c, 0x73, 0xcc, 0x29, 0x55, 0x7a, 0xd2, 0x83, 0x0d,
	0x43, 0x41, 0xf0, 0x20, 0x38, 0x1d, 0x78, 0x9e, 0x3b, 0x88, 0x17, 0x49, 0xd3, 0xcf, 0x36, 0x6c,
	0x4d, 0x4a, 0xbe, 0xac, 0xba, 0x7f, 0xe1, 0xcf, 0xf2, 0xb8, 0xa3, 0x27, 0x91, 0xed, 0x1f, 0xf8,
	0x0b, 0xa4, 0x69, 0xa7, 0x1e, 0xf4, 0xd4, 0xf7, 0xe1, 0x7b, 0xdf, 0xb7, 0xf9, 0x12, 0x72, 0x24,
	0x42, 0x4e, 0x59, 0x96, 0x8d, 0x05, 0x67, 0x46, 0x28, 0x89, 0xd4, 0x68, 0x26, 0xf1, 0x11, 0x34,
	0xcd, 0xbb, 0x14, 0x72, 0x90, 0x06, 0x83, 0x4c, 0x2b, 0xa3, 0xdc, 0x3d, 0x11, 0xf2, 0xe0, 0xb7,
	0x35, 0x58, 0x5a, 0x83, 0xbc, 0xdb, 0x69, 0xc6, 0x2a, 0x56, 0xd6, 0x48, 0x0b, 0x55, 0x66, 0xfc,
	0x2b, 0xb2, 0xd9, 0x2f, 0x3a, 0x86, 0x95, 0xd3, 0x6d, 0x91, 0x3a, 0x82, 0x8c, 0x40, 0xb7, 0x9d,
	0x03, 0xe7, 0xb0, 0x31, 0xa8, 0xc8, 0xed, 0x90, 0x75, 0x0d, 0x1c, 0x44, 0x0e, 0xba, 0xbd, 0x62,
	0x27, 0xdf, 0xec, 0x9f, 0x91, 0x5d, 0x5b, 0x72, 0xc9, 0x47, 0x52, 0x3d, 0x8d, 0x21, 0x8a, 0x21,
	0x05, 0x69, 0x6e, 0x27, 0x9c, 0x03, 0xa2, 0xdb, 0x26, 0x6b, 0x58, 0x4a, 0xdb, 0xb9, 0x31, 0x58,
	0xa2, 0xdf, 0x25, 0x3b, 0x7f, 0x05, 0xfb, 0x5a, 0x2b, 0xed, 0x36, 0xc9, 0x2a, 0x14, 0xa2, 0x3a,
	0x48, 0x09, 0x7e, 0x44, 0x5a, 0x36, 0x72, 0x0d, 0x52, 0xa5, 0x42, 0xda, 0x45, 0x87, 0x9a, 0x71,
	0x70, 0x4f, 0x09, 0x31, 0x85, 0x78, 0x48, 0x18, 0x26, 0x65, 0xa8, 0xb7, 0xfd, 0xf9, 0xbe, 0xbf,
	0x35, 0x65, 0xe9, 0xf8, 0xdc, 0xff, 0x99, 0xf9, 0x83, 0x86, 0x85, 0x1b, 0x86, 0x49, 0xf1, 0x97,
	0xa8, 0xa8, 0xaa, 0x96, 0x2a, 0xa1, 0x77, 0xf7, 0x3a, 0xf7, 0x9c, 0xd9, 0xdc, 0x73, 0x3e, 0xe6,
	0x9e, 0xf3, 0xb2, 0xf0, 0x6a, 0xb3, 0x85, 0x57, 0x7b, 0x5b, 0x78, 0xb5, 0xfb, 0x8b, 0x58, 0x98,
	0x64, 0x12, 0x06, 0x5c, 0xa5, 0x94, 0x2b, 0x4c, 0x15, 0x56, 0x9f, 0x63, 0x8c, 0x46, 0xf4, 0x99,
	0xfe, 0xff, 0x5c, 0x66, 0x9a, 0x01, 0x86, 0x75, 0x7b, 0xef, 0x27, 0x5f, 0x01, 0x00, 0x00, 0xff,
	0xff, 0xba, 0x67, 0xe5, 0xac, 0xd8, 0x01, 0x00, 0x00,
}

func (m *EventTransfer) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EventTransfer) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EventTransfer) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Receiver) > 0 {
		i -= len(m.Receiver)
		copy(dAtA[i:], m.Receiver)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.Receiver)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Sender) > 0 {
		i -= len(m.Sender)
		copy(dAtA[i:], m.Sender)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.Sender)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *EventAcknowledgementSuccess) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EventAcknowledgementSuccess) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EventAcknowledgementSuccess) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Success) > 0 {
		i -= len(m.Success)
		copy(dAtA[i:], m.Success)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.Success)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *EventAcknowledgementError) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EventAcknowledgementError) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EventAcknowledgementError) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Error) > 0 {
		i -= len(m.Error)
		copy(dAtA[i:], m.Error)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.Error)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *EventDenominationTrace) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EventDenominationTrace) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EventDenominationTrace) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Denom) > 0 {
		i -= len(m.Denom)
		copy(dAtA[i:], m.Denom)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.Denom)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.TraceHash) > 0 {
		i -= len(m.TraceHash)
		copy(dAtA[i:], m.TraceHash)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.TraceHash)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintEvents(dAtA []byte, offset int, v uint64) int {
	offset -= sovEvents(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *EventTransfer) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Sender)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	l = len(m.Receiver)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	return n
}

func (m *EventAcknowledgementSuccess) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Success)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	return n
}

func (m *EventAcknowledgementError) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Error)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	return n
}

func (m *EventDenominationTrace) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.TraceHash)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	l = len(m.Denom)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	return n
}

func sovEvents(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozEvents(x uint64) (n int) {
	return sovEvents(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *EventTransfer) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEvents
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
			return fmt.Errorf("proto: EventTransfer: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EventTransfer: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sender", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Sender = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Receiver", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Receiver = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipEvents(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthEvents
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthEvents
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
func (m *EventAcknowledgementSuccess) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEvents
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
			return fmt.Errorf("proto: EventAcknowledgementSuccess: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EventAcknowledgementSuccess: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Success", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Success = append(m.Success[:0], dAtA[iNdEx:postIndex]...)
			if m.Success == nil {
				m.Success = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipEvents(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthEvents
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthEvents
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
func (m *EventAcknowledgementError) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEvents
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
			return fmt.Errorf("proto: EventAcknowledgementError: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EventAcknowledgementError: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Error", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Error = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipEvents(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthEvents
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthEvents
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
func (m *EventDenominationTrace) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEvents
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
			return fmt.Errorf("proto: EventDenominationTrace: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EventDenominationTrace: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TraceHash", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TraceHash = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Denom", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Denom = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipEvents(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthEvents
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthEvents
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
func skipEvents(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowEvents
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
					return 0, ErrIntOverflowEvents
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
					return 0, ErrIntOverflowEvents
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
				return 0, ErrInvalidLengthEvents
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupEvents
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthEvents
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthEvents        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowEvents          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupEvents = fmt.Errorf("proto: unexpected end of group")
)
