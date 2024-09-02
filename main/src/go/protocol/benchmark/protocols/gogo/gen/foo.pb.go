// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: foo.proto

package gen

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	github_com_gogo_protobuf_sortkeys "github.com/gogo/protobuf/sortkeys"
	io "io"
	math "math"
	math_bits "math/bits"
	reflect "reflect"
	strconv "strconv"
	strings "strings"
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

type TestEnum int32

const (
	UNKNOWN TestEnum = 0
	ONE     TestEnum = 1
	TWO     TestEnum = 2
)

var TestEnum_name = map[int32]string{
	0: "UNKNOWN",
	1: "ONE",
	2: "TWO",
}

var TestEnum_value = map[string]int32{
	"UNKNOWN": 0,
	"ONE":     1,
	"TWO":     2,
}

func (TestEnum) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_7ce1e2eec643ca48, []int{0}
}

type TestStruct struct {
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Age  int32  `protobuf:"varint,2,opt,name=age,proto3" json:"age,omitempty"`
}

func (m *TestStruct) Reset()      { *m = TestStruct{} }
func (*TestStruct) ProtoMessage() {}
func (*TestStruct) Descriptor() ([]byte, []int) {
	return fileDescriptor_7ce1e2eec643ca48, []int{0}
}
func (m *TestStruct) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *TestStruct) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_TestStruct.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *TestStruct) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TestStruct.Merge(m, src)
}
func (m *TestStruct) XXX_Size() int {
	return m.Size()
}
func (m *TestStruct) XXX_DiscardUnknown() {
	xxx_messageInfo_TestStruct.DiscardUnknown(m)
}

var xxx_messageInfo_TestStruct proto.InternalMessageInfo

func (m *TestStruct) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *TestStruct) GetAge() int32 {
	if m != nil {
		return m.Age
	}
	return 0
}

type TestArray struct {
	Numbers []int32 `protobuf:"varint,1,rep,packed,name=numbers,proto3" json:"numbers,omitempty"`
}

func (m *TestArray) Reset()      { *m = TestArray{} }
func (*TestArray) ProtoMessage() {}
func (*TestArray) Descriptor() ([]byte, []int) {
	return fileDescriptor_7ce1e2eec643ca48, []int{1}
}
func (m *TestArray) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *TestArray) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_TestArray.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *TestArray) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TestArray.Merge(m, src)
}
func (m *TestArray) XXX_Size() int {
	return m.Size()
}
func (m *TestArray) XXX_DiscardUnknown() {
	xxx_messageInfo_TestArray.DiscardUnknown(m)
}

var xxx_messageInfo_TestArray proto.InternalMessageInfo

func (m *TestArray) GetNumbers() []int32 {
	if m != nil {
		return m.Numbers
	}
	return nil
}

type TestMap struct {
	Entries map[string]int32 `protobuf:"bytes,1,rep,name=entries,proto3" json:"entries,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
}

func (m *TestMap) Reset()      { *m = TestMap{} }
func (*TestMap) ProtoMessage() {}
func (*TestMap) Descriptor() ([]byte, []int) {
	return fileDescriptor_7ce1e2eec643ca48, []int{2}
}
func (m *TestMap) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *TestMap) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_TestMap.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *TestMap) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TestMap.Merge(m, src)
}
func (m *TestMap) XXX_Size() int {
	return m.Size()
}
func (m *TestMap) XXX_DiscardUnknown() {
	xxx_messageInfo_TestMap.DiscardUnknown(m)
}

var xxx_messageInfo_TestMap proto.InternalMessageInfo

func (m *TestMap) GetEntries() map[string]int32 {
	if m != nil {
		return m.Entries
	}
	return nil
}

func init() {
	proto.RegisterEnum("gen.TestEnum", TestEnum_name, TestEnum_value)
	proto.RegisterType((*TestStruct)(nil), "gen.TestStruct")
	proto.RegisterType((*TestArray)(nil), "gen.TestArray")
	proto.RegisterType((*TestMap)(nil), "gen.TestMap")
	proto.RegisterMapType((map[string]int32)(nil), "gen.TestMap.EntriesEntry")
}

func init() { proto.RegisterFile("foo.proto", fileDescriptor_7ce1e2eec643ca48) }

var fileDescriptor_7ce1e2eec643ca48 = []byte{
	// 290 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4c, 0xcb, 0xcf, 0xd7,
	0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x4e, 0x4f, 0xcd, 0x53, 0x32, 0xe2, 0xe2, 0x0a, 0x49,
	0x2d, 0x2e, 0x09, 0x2e, 0x29, 0x2a, 0x4d, 0x2e, 0x11, 0x12, 0xe2, 0x62, 0xc9, 0x4b, 0xcc, 0x4d,
	0x95, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x0c, 0x02, 0xb3, 0x85, 0x04, 0xb8, 0x98, 0x13, 0xd3, 0x53,
	0x25, 0x98, 0x14, 0x18, 0x35, 0x58, 0x83, 0x40, 0x4c, 0x25, 0x55, 0x2e, 0x4e, 0x90, 0x1e, 0xc7,
	0xa2, 0xa2, 0xc4, 0x4a, 0x21, 0x09, 0x2e, 0xf6, 0xbc, 0xd2, 0xdc, 0xa4, 0xd4, 0xa2, 0x62, 0x09,
	0x46, 0x05, 0x66, 0x0d, 0xd6, 0x20, 0x18, 0x57, 0xa9, 0x8a, 0x8b, 0x1d, 0xa4, 0xcc, 0x37, 0xb1,
	0x40, 0xc8, 0x98, 0x8b, 0x3d, 0x35, 0xaf, 0xa4, 0x28, 0x33, 0x15, 0xa2, 0x88, 0xdb, 0x48, 0x52,
	0x2f, 0x3d, 0x35, 0x4f, 0x0f, 0x2a, 0xad, 0xe7, 0x0a, 0x91, 0x03, 0x51, 0x95, 0x41, 0x30, 0x95,
	0x52, 0x56, 0x5c, 0x3c, 0xc8, 0x12, 0x20, 0x87, 0x64, 0xa7, 0x56, 0x42, 0xdd, 0x06, 0x62, 0x0a,
	0x89, 0x70, 0xb1, 0x96, 0x25, 0xe6, 0x94, 0xc2, 0x1c, 0x07, 0xe1, 0x58, 0x31, 0x59, 0x30, 0x6a,
	0x69, 0x72, 0x71, 0x80, 0x0c, 0x77, 0xcd, 0x2b, 0xcd, 0x15, 0xe2, 0xe6, 0x62, 0x0f, 0xf5, 0xf3,
	0xf6, 0xf3, 0x0f, 0xf7, 0x13, 0x60, 0x10, 0x62, 0xe7, 0x62, 0xf6, 0xf7, 0x73, 0x15, 0x60, 0x04,
	0x31, 0x42, 0xc2, 0xfd, 0x05, 0x98, 0x9c, 0x4c, 0x2e, 0x3c, 0x94, 0x63, 0xb8, 0xf1, 0x50, 0x8e,
	0xe1, 0xc3, 0x43, 0x39, 0xc6, 0x86, 0x47, 0x72, 0x8c, 0x2b, 0x1e, 0xc9, 0x31, 0x9e, 0x78, 0x24,
	0xc7, 0x78, 0xe1, 0x91, 0x1c, 0xe3, 0x83, 0x47, 0x72, 0x8c, 0x2f, 0x1e, 0xc9, 0x31, 0x7c, 0x78,
	0x24, 0xc7, 0x38, 0xe1, 0xb1, 0x1c, 0xc3, 0x85, 0xc7, 0x72, 0x0c, 0x37, 0x1e, 0xcb, 0x31, 0x24,
	0xb1, 0x81, 0xc3, 0xd0, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0x75, 0x80, 0xa5, 0xa4, 0x50, 0x01,
	0x00, 0x00,
}

func (x TestEnum) String() string {
	s, ok := TestEnum_name[int32(x)]
	if ok {
		return s
	}
	return strconv.Itoa(int(x))
}
func (this *TestStruct) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*TestStruct)
	if !ok {
		that2, ok := that.(TestStruct)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Name != that1.Name {
		return false
	}
	if this.Age != that1.Age {
		return false
	}
	return true
}
func (this *TestArray) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*TestArray)
	if !ok {
		that2, ok := that.(TestArray)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if len(this.Numbers) != len(that1.Numbers) {
		return false
	}
	for i := range this.Numbers {
		if this.Numbers[i] != that1.Numbers[i] {
			return false
		}
	}
	return true
}
func (this *TestMap) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*TestMap)
	if !ok {
		that2, ok := that.(TestMap)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if len(this.Entries) != len(that1.Entries) {
		return false
	}
	for i := range this.Entries {
		if this.Entries[i] != that1.Entries[i] {
			return false
		}
	}
	return true
}
func (this *TestStruct) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 6)
	s = append(s, "&gen.TestStruct{")
	s = append(s, "Name: "+fmt.Sprintf("%#v", this.Name)+",\n")
	s = append(s, "Age: "+fmt.Sprintf("%#v", this.Age)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *TestArray) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 5)
	s = append(s, "&gen.TestArray{")
	s = append(s, "Numbers: "+fmt.Sprintf("%#v", this.Numbers)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *TestMap) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 5)
	s = append(s, "&gen.TestMap{")
	keysForEntries := make([]string, 0, len(this.Entries))
	for k, _ := range this.Entries {
		keysForEntries = append(keysForEntries, k)
	}
	github_com_gogo_protobuf_sortkeys.Strings(keysForEntries)
	mapStringForEntries := "map[string]int32{"
	for _, k := range keysForEntries {
		mapStringForEntries += fmt.Sprintf("%#v: %#v,", k, this.Entries[k])
	}
	mapStringForEntries += "}"
	if this.Entries != nil {
		s = append(s, "Entries: "+mapStringForEntries+",\n")
	}
	s = append(s, "}")
	return strings.Join(s, "")
}
func valueToGoStringFoo(v interface{}, typ string) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("func(v %v) *%v { return &v } ( %#v )", typ, typ, pv)
}
func (m *TestStruct) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TestStruct) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *TestStruct) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Age != 0 {
		i = encodeVarintFoo(dAtA, i, uint64(m.Age))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintFoo(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *TestArray) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TestArray) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *TestArray) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Numbers) > 0 {
		dAtA2 := make([]byte, len(m.Numbers)*10)
		var j1 int
		for _, num1 := range m.Numbers {
			num := uint64(num1)
			for num >= 1<<7 {
				dAtA2[j1] = uint8(uint64(num)&0x7f | 0x80)
				num >>= 7
				j1++
			}
			dAtA2[j1] = uint8(num)
			j1++
		}
		i -= j1
		copy(dAtA[i:], dAtA2[:j1])
		i = encodeVarintFoo(dAtA, i, uint64(j1))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *TestMap) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TestMap) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *TestMap) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Entries) > 0 {
		for k := range m.Entries {
			v := m.Entries[k]
			baseI := i
			i = encodeVarintFoo(dAtA, i, uint64(v))
			i--
			dAtA[i] = 0x10
			i -= len(k)
			copy(dAtA[i:], k)
			i = encodeVarintFoo(dAtA, i, uint64(len(k)))
			i--
			dAtA[i] = 0xa
			i = encodeVarintFoo(dAtA, i, uint64(baseI-i))
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintFoo(dAtA []byte, offset int, v uint64) int {
	offset -= sovFoo(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *TestStruct) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovFoo(uint64(l))
	}
	if m.Age != 0 {
		n += 1 + sovFoo(uint64(m.Age))
	}
	return n
}

func (m *TestArray) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Numbers) > 0 {
		l = 0
		for _, e := range m.Numbers {
			l += sovFoo(uint64(e))
		}
		n += 1 + sovFoo(uint64(l)) + l
	}
	return n
}

func (m *TestMap) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Entries) > 0 {
		for k, v := range m.Entries {
			_ = k
			_ = v
			mapEntrySize := 1 + len(k) + sovFoo(uint64(len(k))) + 1 + sovFoo(uint64(v))
			n += mapEntrySize + 1 + sovFoo(uint64(mapEntrySize))
		}
	}
	return n
}

func sovFoo(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozFoo(x uint64) (n int) {
	return sovFoo(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *TestStruct) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&TestStruct{`,
		`Name:` + fmt.Sprintf("%v", this.Name) + `,`,
		`Age:` + fmt.Sprintf("%v", this.Age) + `,`,
		`}`,
	}, "")
	return s
}
func (this *TestArray) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&TestArray{`,
		`Numbers:` + fmt.Sprintf("%v", this.Numbers) + `,`,
		`}`,
	}, "")
	return s
}
func (this *TestMap) String() string {
	if this == nil {
		return "nil"
	}
	keysForEntries := make([]string, 0, len(this.Entries))
	for k, _ := range this.Entries {
		keysForEntries = append(keysForEntries, k)
	}
	github_com_gogo_protobuf_sortkeys.Strings(keysForEntries)
	mapStringForEntries := "map[string]int32{"
	for _, k := range keysForEntries {
		mapStringForEntries += fmt.Sprintf("%v: %v,", k, this.Entries[k])
	}
	mapStringForEntries += "}"
	s := strings.Join([]string{`&TestMap{`,
		`Entries:` + mapStringForEntries + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringFoo(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *TestStruct) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowFoo
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
			return fmt.Errorf("proto: TestStruct: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TestStruct: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFoo
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
				return ErrInvalidLengthFoo
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthFoo
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Age", wireType)
			}
			m.Age = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFoo
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
		default:
			iNdEx = preIndex
			skippy, err := skipFoo(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthFoo
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
func (m *TestArray) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowFoo
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
			return fmt.Errorf("proto: TestArray: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TestArray: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType == 0 {
				var v int32
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowFoo
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					v |= int32(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				m.Numbers = append(m.Numbers, v)
			} else if wireType == 2 {
				var packedLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowFoo
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					packedLen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if packedLen < 0 {
					return ErrInvalidLengthFoo
				}
				postIndex := iNdEx + packedLen
				if postIndex < 0 {
					return ErrInvalidLengthFoo
				}
				if postIndex > l {
					return io.ErrUnexpectedEOF
				}
				var elementCount int
				var count int
				for _, integer := range dAtA[iNdEx:postIndex] {
					if integer < 128 {
						count++
					}
				}
				elementCount = count
				if elementCount != 0 && len(m.Numbers) == 0 {
					m.Numbers = make([]int32, 0, elementCount)
				}
				for iNdEx < postIndex {
					var v int32
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowFoo
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						v |= int32(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					m.Numbers = append(m.Numbers, v)
				}
			} else {
				return fmt.Errorf("proto: wrong wireType = %d for field Numbers", wireType)
			}
		default:
			iNdEx = preIndex
			skippy, err := skipFoo(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthFoo
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
func (m *TestMap) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowFoo
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
			return fmt.Errorf("proto: TestMap: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TestMap: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Entries", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFoo
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
				return ErrInvalidLengthFoo
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthFoo
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Entries == nil {
				m.Entries = make(map[string]int32)
			}
			var mapkey string
			var mapvalue int32
			for iNdEx < postIndex {
				entryPreIndex := iNdEx
				var wire uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowFoo
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
				if fieldNum == 1 {
					var stringLenmapkey uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowFoo
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						stringLenmapkey |= uint64(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					intStringLenmapkey := int(stringLenmapkey)
					if intStringLenmapkey < 0 {
						return ErrInvalidLengthFoo
					}
					postStringIndexmapkey := iNdEx + intStringLenmapkey
					if postStringIndexmapkey < 0 {
						return ErrInvalidLengthFoo
					}
					if postStringIndexmapkey > l {
						return io.ErrUnexpectedEOF
					}
					mapkey = string(dAtA[iNdEx:postStringIndexmapkey])
					iNdEx = postStringIndexmapkey
				} else if fieldNum == 2 {
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowFoo
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						mapvalue |= int32(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
				} else {
					iNdEx = entryPreIndex
					skippy, err := skipFoo(dAtA[iNdEx:])
					if err != nil {
						return err
					}
					if (skippy < 0) || (iNdEx+skippy) < 0 {
						return ErrInvalidLengthFoo
					}
					if (iNdEx + skippy) > postIndex {
						return io.ErrUnexpectedEOF
					}
					iNdEx += skippy
				}
			}
			m.Entries[mapkey] = mapvalue
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipFoo(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthFoo
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
func skipFoo(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowFoo
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
					return 0, ErrIntOverflowFoo
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
					return 0, ErrIntOverflowFoo
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
				return 0, ErrInvalidLengthFoo
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupFoo
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthFoo
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthFoo        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowFoo          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupFoo = fmt.Errorf("proto: unexpected end of group")
)