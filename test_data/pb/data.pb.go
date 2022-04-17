// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: data.proto

package pb

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/golang/protobuf/proto"
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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type Gender int32

const (
	Gender_men   Gender = 0
	Gender_women Gender = 1
)

var Gender_name = map[int32]string{
	0: "men",
	1: "women",
}

var Gender_value = map[string]int32{
	"men":   0,
	"women": 1,
}

func (x Gender) String() string {
	return proto.EnumName(Gender_name, int32(x))
}

func (Gender) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_871986018790d2fd, []int{0}
}

type Person struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty" idx:"normal"`
	Phone                string   `protobuf:"bytes,2,opt,name=phone,proto3" json:"phone,omitempty" idx:"unique"`
	Age                  int32    `protobuf:"varint,3,opt,name=age,proto3" json:"age,omitempty" idx:"normal"`
	BirthDay             int32    `protobuf:"varint,4,opt,name=birthDay,proto3" json:"birthDay,omitempty"`
	Gender               Gender   `protobuf:"varint,5,opt,name=gender,proto3,enum=pb.Gender" json:"gender,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Person) Reset()         { *m = Person{} }
func (m *Person) String() string { return proto.CompactTextString(m) }
func (*Person) ProtoMessage()    {}
func (*Person) Descriptor() ([]byte, []int) {
	return fileDescriptor_871986018790d2fd, []int{0}
}
func (m *Person) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Person) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Person.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Person) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Person.Merge(m, src)
}
func (m *Person) XXX_Size() int {
	return m.Size()
}
func (m *Person) XXX_DiscardUnknown() {
	xxx_messageInfo_Person.DiscardUnknown(m)
}

var xxx_messageInfo_Person proto.InternalMessageInfo

func (m *Person) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Person) GetPhone() string {
	if m != nil {
		return m.Phone
	}
	return ""
}

func (m *Person) GetAge() int32 {
	if m != nil {
		return m.Age
	}
	return 0
}

func (m *Person) GetBirthDay() int32 {
	if m != nil {
		return m.BirthDay
	}
	return 0
}

func (m *Person) GetGender() Gender {
	if m != nil {
		return m.Gender
	}
	return Gender_men
}

func init() {
	proto.RegisterEnum("pb.Gender", Gender_name, Gender_value)
	proto.RegisterType((*Person)(nil), "pb.Person")
}

func init() { proto.RegisterFile("data.proto", fileDescriptor_871986018790d2fd) }

var fileDescriptor_871986018790d2fd = []byte{
	// 251 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4a, 0x49, 0x2c, 0x49,
	0xd4, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x2a, 0x48, 0x92, 0xd2, 0x4d, 0xcf, 0x2c, 0xc9,
	0x28, 0x4d, 0xd2, 0x4b, 0xce, 0xcf, 0xd5, 0x4f, 0xcf, 0x4f, 0xcf, 0xd7, 0x07, 0x4b, 0x25, 0x95,
	0xa6, 0x81, 0x79, 0x60, 0x0e, 0x98, 0x05, 0xd1, 0xa2, 0xb4, 0x8b, 0x91, 0x8b, 0x2d, 0x20, 0xb5,
	0xa8, 0x38, 0x3f, 0x4f, 0x48, 0x85, 0x8b, 0x25, 0x2f, 0x31, 0x37, 0x55, 0x82, 0x51, 0x81, 0x51,
	0x83, 0xd3, 0x49, 0xe0, 0xd3, 0x3d, 0x79, 0x9e, 0xcc, 0x94, 0x0a, 0x2b, 0xa5, 0xbc, 0xfc, 0xa2,
	0xdc, 0xc4, 0x1c, 0xa5, 0x20, 0xb0, 0xac, 0x90, 0x1a, 0x17, 0x6b, 0x41, 0x46, 0x7e, 0x5e, 0xaa,
	0x04, 0x13, 0x9a, 0xb2, 0xd2, 0xbc, 0xcc, 0xc2, 0xd2, 0x54, 0xa5, 0x20, 0x88, 0xb4, 0x90, 0x12,
	0x17, 0x73, 0x62, 0x7a, 0xaa, 0x04, 0xb3, 0x02, 0xa3, 0x06, 0x2b, 0x16, 0xc3, 0x40, 0x92, 0x42,
	0x52, 0x5c, 0x1c, 0x49, 0x99, 0x45, 0x25, 0x19, 0x2e, 0x89, 0x95, 0x12, 0x2c, 0x20, 0x85, 0x41,
	0x70, 0xbe, 0x90, 0x12, 0x17, 0x5b, 0x7a, 0x6a, 0x5e, 0x4a, 0x6a, 0x91, 0x04, 0xab, 0x02, 0xa3,
	0x06, 0x9f, 0x11, 0x97, 0x5e, 0x41, 0x92, 0x9e, 0x3b, 0x58, 0x24, 0x08, 0x2a, 0xa3, 0x25, 0xc3,
	0xc5, 0x06, 0x11, 0x11, 0x62, 0xe7, 0x62, 0xce, 0x4d, 0xcd, 0x13, 0x60, 0x10, 0xe2, 0xe4, 0x62,
	0x2d, 0xcf, 0x07, 0x31, 0x19, 0x9d, 0x04, 0x4e, 0x3c, 0x92, 0x63, 0xbc, 0xf0, 0x48, 0x8e, 0xf1,
	0xc1, 0x23, 0x39, 0xc6, 0x19, 0x8f, 0xe5, 0x18, 0x92, 0xd8, 0xc0, 0x7e, 0x36, 0x06, 0x04, 0x00,
	0x00, 0xff, 0xff, 0x61, 0x81, 0x52, 0x98, 0x34, 0x01, 0x00, 0x00,
}

func (m *Person) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Person) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Person) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.Gender != 0 {
		i = encodeVarintData(dAtA, i, uint64(m.Gender))
		i--
		dAtA[i] = 0x28
	}
	if m.BirthDay != 0 {
		i = encodeVarintData(dAtA, i, uint64(m.BirthDay))
		i--
		dAtA[i] = 0x20
	}
	if m.Age != 0 {
		i = encodeVarintData(dAtA, i, uint64(m.Age))
		i--
		dAtA[i] = 0x18
	}
	if len(m.Phone) > 0 {
		i -= len(m.Phone)
		copy(dAtA[i:], m.Phone)
		i = encodeVarintData(dAtA, i, uint64(len(m.Phone)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintData(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintData(dAtA []byte, offset int, v uint64) int {
	offset -= sovData(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Person) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovData(uint64(l))
	}
	l = len(m.Phone)
	if l > 0 {
		n += 1 + l + sovData(uint64(l))
	}
	if m.Age != 0 {
		n += 1 + sovData(uint64(m.Age))
	}
	if m.BirthDay != 0 {
		n += 1 + sovData(uint64(m.BirthDay))
	}
	if m.Gender != 0 {
		n += 1 + sovData(uint64(m.Gender))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovData(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozData(x uint64) (n int) {
	return sovData(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Person) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowData
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
			return fmt.Errorf("proto: Person: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Person: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowData
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
				return ErrInvalidLengthData
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthData
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Phone", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowData
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
				return ErrInvalidLengthData
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthData
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Phone = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Age", wireType)
			}
			m.Age = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowData
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Age |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field BirthDay", wireType)
			}
			m.BirthDay = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowData
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.BirthDay |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Gender", wireType)
			}
			m.Gender = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowData
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Gender |= Gender(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipData(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthData
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
func skipData(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowData
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
					return 0, ErrIntOverflowData
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
					return 0, ErrIntOverflowData
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
				return 0, ErrInvalidLengthData
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupData
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthData
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthData        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowData          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupData = fmt.Errorf("proto: unexpected end of group")
)
