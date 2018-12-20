// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: example.proto

package protos

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

// Response represents a rpc message
type Player struct {
	PrivateID    string `protobuf:"bytes,1,opt,name=PrivateID,proto3" json:"PrivateID,omitempty"`
	AccessToken  string `protobuf:"bytes,2,opt,name=AccessToken,proto3" json:"AccessToken,omitempty"`
	Name         string `protobuf:"bytes,3,opt,name=Name,proto3" json:"Name,omitempty"`
	SoftCurrency int32  `protobuf:"varint,4,opt,name=SoftCurrency,proto3" json:"SoftCurrency,omitempty"`
	Trophies     int32  `protobuf:"varint,5,opt,name=Trophies,proto3" json:"Trophies,omitempty"`
}

func (m *Player) Reset()         { *m = Player{} }
func (m *Player) String() string { return proto.CompactTextString(m) }
func (*Player) ProtoMessage()    {}
func (*Player) Descriptor() ([]byte, []int) {
	return fileDescriptor_example_4c07b9927c24f8f0, []int{0}
}
func (m *Player) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Player) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Player.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *Player) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Player.Merge(dst, src)
}
func (m *Player) XXX_Size() int {
	return m.Size()
}
func (m *Player) XXX_DiscardUnknown() {
	xxx_messageInfo_Player.DiscardUnknown(m)
}

var xxx_messageInfo_Player proto.InternalMessageInfo

func (m *Player) GetPrivateID() string {
	if m != nil {
		return m.PrivateID
	}
	return ""
}

func (m *Player) GetAccessToken() string {
	if m != nil {
		return m.AccessToken
	}
	return ""
}

func (m *Player) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Player) GetSoftCurrency() int32 {
	if m != nil {
		return m.SoftCurrency
	}
	return 0
}

func (m *Player) GetTrophies() int32 {
	if m != nil {
		return m.Trophies
	}
	return 0
}

type AuthResponse struct {
	Code   string  `protobuf:"bytes,1,opt,name=Code,proto3" json:"Code,omitempty"`
	Msg    string  `protobuf:"bytes,2,opt,name=Msg,proto3" json:"Msg,omitempty"`
	Player *Player `protobuf:"bytes,3,opt,name=Player" json:"Player,omitempty"`
}

func (m *AuthResponse) Reset()         { *m = AuthResponse{} }
func (m *AuthResponse) String() string { return proto.CompactTextString(m) }
func (*AuthResponse) ProtoMessage()    {}
func (*AuthResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_example_4c07b9927c24f8f0, []int{1}
}
func (m *AuthResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *AuthResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_AuthResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *AuthResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AuthResponse.Merge(dst, src)
}
func (m *AuthResponse) XXX_Size() int {
	return m.Size()
}
func (m *AuthResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_AuthResponse.DiscardUnknown(m)
}

var xxx_messageInfo_AuthResponse proto.InternalMessageInfo

func (m *AuthResponse) GetCode() string {
	if m != nil {
		return m.Code
	}
	return ""
}

func (m *AuthResponse) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

func (m *AuthResponse) GetPlayer() *Player {
	if m != nil {
		return m.Player
	}
	return nil
}

type AuthArg struct {
	AccessToken string `protobuf:"bytes,1,opt,name=accessToken,proto3" json:"accessToken,omitempty"`
}

func (m *AuthArg) Reset()         { *m = AuthArg{} }
func (m *AuthArg) String() string { return proto.CompactTextString(m) }
func (*AuthArg) ProtoMessage()    {}
func (*AuthArg) Descriptor() ([]byte, []int) {
	return fileDescriptor_example_4c07b9927c24f8f0, []int{2}
}
func (m *AuthArg) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *AuthArg) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_AuthArg.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *AuthArg) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AuthArg.Merge(dst, src)
}
func (m *AuthArg) XXX_Size() int {
	return m.Size()
}
func (m *AuthArg) XXX_DiscardUnknown() {
	xxx_messageInfo_AuthArg.DiscardUnknown(m)
}

var xxx_messageInfo_AuthArg proto.InternalMessageInfo

func (m *AuthArg) GetAccessToken() string {
	if m != nil {
		return m.AccessToken
	}
	return ""
}

type FindMatchArg struct {
	RoomType string `protobuf:"bytes,1,opt,name=RoomType,proto3" json:"RoomType,omitempty"`
}

func (m *FindMatchArg) Reset()         { *m = FindMatchArg{} }
func (m *FindMatchArg) String() string { return proto.CompactTextString(m) }
func (*FindMatchArg) ProtoMessage()    {}
func (*FindMatchArg) Descriptor() ([]byte, []int) {
	return fileDescriptor_example_4c07b9927c24f8f0, []int{3}
}
func (m *FindMatchArg) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *FindMatchArg) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_FindMatchArg.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *FindMatchArg) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FindMatchArg.Merge(dst, src)
}
func (m *FindMatchArg) XXX_Size() int {
	return m.Size()
}
func (m *FindMatchArg) XXX_DiscardUnknown() {
	xxx_messageInfo_FindMatchArg.DiscardUnknown(m)
}

var xxx_messageInfo_FindMatchArg proto.InternalMessageInfo

func (m *FindMatchArg) GetRoomType() string {
	if m != nil {
		return m.RoomType
	}
	return ""
}

type FindMatchResponse struct {
	Code string `protobuf:"bytes,1,opt,name=Code,proto3" json:"Code,omitempty"`
	Msg  string `protobuf:"bytes,2,opt,name=Msg,proto3" json:"Msg,omitempty"`
}

func (m *FindMatchResponse) Reset()         { *m = FindMatchResponse{} }
func (m *FindMatchResponse) String() string { return proto.CompactTextString(m) }
func (*FindMatchResponse) ProtoMessage()    {}
func (*FindMatchResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_example_4c07b9927c24f8f0, []int{4}
}
func (m *FindMatchResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *FindMatchResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_FindMatchResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *FindMatchResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FindMatchResponse.Merge(dst, src)
}
func (m *FindMatchResponse) XXX_Size() int {
	return m.Size()
}
func (m *FindMatchResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_FindMatchResponse.DiscardUnknown(m)
}

var xxx_messageInfo_FindMatchResponse proto.InternalMessageInfo

func (m *FindMatchResponse) GetCode() string {
	if m != nil {
		return m.Code
	}
	return ""
}

func (m *FindMatchResponse) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

type FindMatchPush struct {
	Code string `protobuf:"bytes,1,opt,name=Code,proto3" json:"Code,omitempty"`
	IP   string `protobuf:"bytes,2,opt,name=IP,proto3" json:"IP,omitempty"`
	Port int32  `protobuf:"varint,3,opt,name=Port,proto3" json:"Port,omitempty"`
}

func (m *FindMatchPush) Reset()         { *m = FindMatchPush{} }
func (m *FindMatchPush) String() string { return proto.CompactTextString(m) }
func (*FindMatchPush) ProtoMessage()    {}
func (*FindMatchPush) Descriptor() ([]byte, []int) {
	return fileDescriptor_example_4c07b9927c24f8f0, []int{5}
}
func (m *FindMatchPush) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *FindMatchPush) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_FindMatchPush.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *FindMatchPush) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FindMatchPush.Merge(dst, src)
}
func (m *FindMatchPush) XXX_Size() int {
	return m.Size()
}
func (m *FindMatchPush) XXX_DiscardUnknown() {
	xxx_messageInfo_FindMatchPush.DiscardUnknown(m)
}

var xxx_messageInfo_FindMatchPush proto.InternalMessageInfo

func (m *FindMatchPush) GetCode() string {
	if m != nil {
		return m.Code
	}
	return ""
}

func (m *FindMatchPush) GetIP() string {
	if m != nil {
		return m.IP
	}
	return ""
}

func (m *FindMatchPush) GetPort() int32 {
	if m != nil {
		return m.Port
	}
	return 0
}

func init() {
	proto.RegisterType((*Player)(nil), "protos.Player")
	proto.RegisterType((*AuthResponse)(nil), "protos.AuthResponse")
	proto.RegisterType((*AuthArg)(nil), "protos.AuthArg")
	proto.RegisterType((*FindMatchArg)(nil), "protos.FindMatchArg")
	proto.RegisterType((*FindMatchResponse)(nil), "protos.FindMatchResponse")
	proto.RegisterType((*FindMatchPush)(nil), "protos.FindMatchPush")
}
func (m *Player) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Player) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.PrivateID) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintExample(dAtA, i, uint64(len(m.PrivateID)))
		i += copy(dAtA[i:], m.PrivateID)
	}
	if len(m.AccessToken) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintExample(dAtA, i, uint64(len(m.AccessToken)))
		i += copy(dAtA[i:], m.AccessToken)
	}
	if len(m.Name) > 0 {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintExample(dAtA, i, uint64(len(m.Name)))
		i += copy(dAtA[i:], m.Name)
	}
	if m.SoftCurrency != 0 {
		dAtA[i] = 0x20
		i++
		i = encodeVarintExample(dAtA, i, uint64(m.SoftCurrency))
	}
	if m.Trophies != 0 {
		dAtA[i] = 0x28
		i++
		i = encodeVarintExample(dAtA, i, uint64(m.Trophies))
	}
	return i, nil
}

func (m *AuthResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *AuthResponse) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Code) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintExample(dAtA, i, uint64(len(m.Code)))
		i += copy(dAtA[i:], m.Code)
	}
	if len(m.Msg) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintExample(dAtA, i, uint64(len(m.Msg)))
		i += copy(dAtA[i:], m.Msg)
	}
	if m.Player != nil {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintExample(dAtA, i, uint64(m.Player.Size()))
		n1, err := m.Player.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	return i, nil
}

func (m *AuthArg) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *AuthArg) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.AccessToken) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintExample(dAtA, i, uint64(len(m.AccessToken)))
		i += copy(dAtA[i:], m.AccessToken)
	}
	return i, nil
}

func (m *FindMatchArg) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *FindMatchArg) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.RoomType) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintExample(dAtA, i, uint64(len(m.RoomType)))
		i += copy(dAtA[i:], m.RoomType)
	}
	return i, nil
}

func (m *FindMatchResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *FindMatchResponse) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Code) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintExample(dAtA, i, uint64(len(m.Code)))
		i += copy(dAtA[i:], m.Code)
	}
	if len(m.Msg) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintExample(dAtA, i, uint64(len(m.Msg)))
		i += copy(dAtA[i:], m.Msg)
	}
	return i, nil
}

func (m *FindMatchPush) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *FindMatchPush) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Code) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintExample(dAtA, i, uint64(len(m.Code)))
		i += copy(dAtA[i:], m.Code)
	}
	if len(m.IP) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintExample(dAtA, i, uint64(len(m.IP)))
		i += copy(dAtA[i:], m.IP)
	}
	if m.Port != 0 {
		dAtA[i] = 0x18
		i++
		i = encodeVarintExample(dAtA, i, uint64(m.Port))
	}
	return i, nil
}

func encodeVarintExample(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *Player) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.PrivateID)
	if l > 0 {
		n += 1 + l + sovExample(uint64(l))
	}
	l = len(m.AccessToken)
	if l > 0 {
		n += 1 + l + sovExample(uint64(l))
	}
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovExample(uint64(l))
	}
	if m.SoftCurrency != 0 {
		n += 1 + sovExample(uint64(m.SoftCurrency))
	}
	if m.Trophies != 0 {
		n += 1 + sovExample(uint64(m.Trophies))
	}
	return n
}

func (m *AuthResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Code)
	if l > 0 {
		n += 1 + l + sovExample(uint64(l))
	}
	l = len(m.Msg)
	if l > 0 {
		n += 1 + l + sovExample(uint64(l))
	}
	if m.Player != nil {
		l = m.Player.Size()
		n += 1 + l + sovExample(uint64(l))
	}
	return n
}

func (m *AuthArg) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.AccessToken)
	if l > 0 {
		n += 1 + l + sovExample(uint64(l))
	}
	return n
}

func (m *FindMatchArg) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.RoomType)
	if l > 0 {
		n += 1 + l + sovExample(uint64(l))
	}
	return n
}

func (m *FindMatchResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Code)
	if l > 0 {
		n += 1 + l + sovExample(uint64(l))
	}
	l = len(m.Msg)
	if l > 0 {
		n += 1 + l + sovExample(uint64(l))
	}
	return n
}

func (m *FindMatchPush) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Code)
	if l > 0 {
		n += 1 + l + sovExample(uint64(l))
	}
	l = len(m.IP)
	if l > 0 {
		n += 1 + l + sovExample(uint64(l))
	}
	if m.Port != 0 {
		n += 1 + sovExample(uint64(m.Port))
	}
	return n
}

func sovExample(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozExample(x uint64) (n int) {
	return sovExample(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Player) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowExample
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Player: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Player: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PrivateID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowExample
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthExample
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PrivateID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AccessToken", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowExample
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthExample
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AccessToken = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowExample
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthExample
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field SoftCurrency", wireType)
			}
			m.SoftCurrency = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowExample
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.SoftCurrency |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Trophies", wireType)
			}
			m.Trophies = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowExample
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Trophies |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipExample(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthExample
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
func (m *AuthResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowExample
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: AuthResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: AuthResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Code", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowExample
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthExample
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Code = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Msg", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowExample
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthExample
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Msg = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Player", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowExample
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthExample
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Player == nil {
				m.Player = &Player{}
			}
			if err := m.Player.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipExample(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthExample
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
func (m *AuthArg) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowExample
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: AuthArg: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: AuthArg: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AccessToken", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowExample
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthExample
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AccessToken = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipExample(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthExample
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
func (m *FindMatchArg) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowExample
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: FindMatchArg: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: FindMatchArg: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RoomType", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowExample
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthExample
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.RoomType = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipExample(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthExample
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
func (m *FindMatchResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowExample
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: FindMatchResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: FindMatchResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Code", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowExample
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthExample
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Code = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Msg", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowExample
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthExample
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Msg = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipExample(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthExample
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
func (m *FindMatchPush) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowExample
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: FindMatchPush: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: FindMatchPush: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Code", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowExample
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthExample
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Code = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field IP", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowExample
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthExample
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.IP = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Port", wireType)
			}
			m.Port = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowExample
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Port |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipExample(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthExample
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
func skipExample(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowExample
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
					return 0, ErrIntOverflowExample
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
					return 0, ErrIntOverflowExample
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
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthExample
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowExample
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
				next, err := skipExample(dAtA[start:])
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
	ErrInvalidLengthExample = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowExample   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("example.proto", fileDescriptor_example_4c07b9927c24f8f0) }

var fileDescriptor_example_4c07b9927c24f8f0 = []byte{
	// 327 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x51, 0x3d, 0x4f, 0x02, 0x41,
	0x14, 0x64, 0xf9, 0x12, 0x1e, 0x1f, 0xd1, 0xad, 0x2e, 0xc4, 0x5c, 0xc8, 0x16, 0x86, 0x68, 0x42,
	0xa1, 0x95, 0x25, 0x62, 0x34, 0x14, 0x98, 0xcd, 0x4a, 0x69, 0xb3, 0x1e, 0x4f, 0x20, 0xc2, 0xed,
	0x65, 0x77, 0x31, 0xf2, 0x2f, 0xfc, 0x01, 0xfe, 0x20, 0x4b, 0x4a, 0x4b, 0x03, 0x7f, 0xc4, 0xdc,
	0xde, 0x71, 0x62, 0x62, 0x63, 0x75, 0xf3, 0x66, 0xde, 0xbc, 0x9b, 0xc9, 0x42, 0x03, 0x5f, 0xe5,
	0x22, 0x9a, 0x63, 0x37, 0xd2, 0xca, 0x2a, 0x5a, 0x76, 0x1f, 0xc3, 0xde, 0x09, 0x94, 0xf9, 0x5c,
	0xae, 0x50, 0xd3, 0x63, 0xa8, 0x72, 0x3d, 0x7b, 0x91, 0x16, 0x07, 0xd7, 0x1e, 0x69, 0x93, 0x4e,
	0x55, 0xfc, 0x10, 0xb4, 0x0d, 0xb5, 0x5e, 0x10, 0xa0, 0x31, 0x23, 0xf5, 0x8c, 0xa1, 0x97, 0x77,
	0xfa, 0x3e, 0x45, 0x29, 0x14, 0xef, 0xe4, 0x02, 0xbd, 0x82, 0x93, 0x1c, 0xa6, 0x0c, 0xea, 0xf7,
	0xea, 0xc9, 0xf6, 0x97, 0x5a, 0x63, 0x18, 0xac, 0xbc, 0x62, 0x9b, 0x74, 0x4a, 0xe2, 0x17, 0x47,
	0x5b, 0x50, 0x19, 0x69, 0x15, 0x4d, 0x67, 0x68, 0xbc, 0x92, 0xd3, 0xb3, 0x99, 0x3d, 0x40, 0xbd,
	0xb7, 0xb4, 0x53, 0x81, 0x26, 0x52, 0xa1, 0xc1, 0xf8, 0x1f, 0x7d, 0x35, 0xc6, 0x34, 0x9e, 0xc3,
	0xf4, 0x10, 0x0a, 0x43, 0x33, 0x49, 0x13, 0xc5, 0x90, 0x9e, 0xec, 0x3a, 0xb9, 0x2c, 0xb5, 0xf3,
	0x66, 0x52, 0xda, 0x74, 0x13, 0x56, 0xa4, 0x2a, 0x3b, 0x83, 0x83, 0xf8, 0x7a, 0x4f, 0x4f, 0xe2,
	0x7a, 0x72, 0xaf, 0x5e, 0x72, 0x7f, 0x9f, 0x62, 0xa7, 0x50, 0xbf, 0x99, 0x85, 0xe3, 0xa1, 0xb4,
	0x81, 0x73, 0xb4, 0xa0, 0x22, 0x94, 0x5a, 0x8c, 0x56, 0xd1, 0x2e, 0x4e, 0x36, 0xb3, 0x4b, 0x38,
	0xca, 0x76, 0xff, 0x97, 0x9d, 0xdd, 0x42, 0x23, 0xb3, 0xf2, 0xa5, 0x99, 0xfe, 0x69, 0x6b, 0x42,
	0x7e, 0xc0, 0x53, 0x57, 0x7e, 0xc0, 0xe3, 0x1d, 0xae, 0xb4, 0x75, 0x75, 0x4b, 0xc2, 0xe1, 0x2b,
	0xef, 0x63, 0xe3, 0x93, 0xf5, 0xc6, 0x27, 0x5f, 0x1b, 0x9f, 0xbc, 0x6d, 0xfd, 0xdc, 0x7a, 0xeb,
	0xe7, 0x3e, 0xb7, 0x7e, 0xee, 0x31, 0x79, 0xfb, 0x8b, 0xef, 0x00, 0x00, 0x00, 0xff, 0xff, 0x71,
	0x44, 0x40, 0x9d, 0x13, 0x02, 0x00, 0x00,
}
