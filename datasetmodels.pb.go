// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: datasetmodels.proto

package otsimopb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type Column_Type int32

const (
	Column_STRING      Column_Type = 0
	Column_INTEGER     Column_Type = 1
	Column_REAL        Column_Type = 2
	Column_DATE        Column_Type = 3
	Column_DATE_TIME   Column_Type = 4
	Column_TIME_OF_DAY Column_Type = 5
)

var Column_Type_name = map[int32]string{
	0: "STRING",
	1: "INTEGER",
	2: "REAL",
	3: "DATE",
	4: "DATE_TIME",
	5: "TIME_OF_DAY",
}
var Column_Type_value = map[string]int32{
	"STRING":      0,
	"INTEGER":     1,
	"REAL":        2,
	"DATE":        3,
	"DATE_TIME":   4,
	"TIME_OF_DAY": 5,
}

func (x Column_Type) String() string {
	return proto.EnumName(Column_Type_name, int32(x))
}
func (Column_Type) EnumDescriptor() ([]byte, []int) { return fileDescriptorDatasetmodels, []int{0, 0} }

type Column struct {
	Type Column_Type `protobuf:"varint,1,opt,name=type,proto3,enum=apipb.Column_Type" json:"type,omitempty"`
	Name string      `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (m *Column) Reset()                    { *m = Column{} }
func (m *Column) String() string            { return proto.CompactTextString(m) }
func (*Column) ProtoMessage()               {}
func (*Column) Descriptor() ([]byte, []int) { return fileDescriptorDatasetmodels, []int{0} }

func (m *Column) GetType() Column_Type {
	if m != nil {
		return m.Type
	}
	return Column_STRING
}

func (m *Column) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type TimeOfDay struct {
	Hours        int32 `protobuf:"varint,1,opt,name=hours,proto3" json:"hours,omitempty"`
	Minutes      int32 `protobuf:"varint,2,opt,name=minutes,proto3" json:"minutes,omitempty"`
	Seconds      int32 `protobuf:"varint,3,opt,name=seconds,proto3" json:"seconds,omitempty"`
	Milliseconds int32 `protobuf:"varint,4,opt,name=milliseconds,proto3" json:"milliseconds,omitempty"`
}

func (m *TimeOfDay) Reset()                    { *m = TimeOfDay{} }
func (m *TimeOfDay) String() string            { return proto.CompactTextString(m) }
func (*TimeOfDay) ProtoMessage()               {}
func (*TimeOfDay) Descriptor() ([]byte, []int) { return fileDescriptorDatasetmodels, []int{1} }

func (m *TimeOfDay) GetHours() int32 {
	if m != nil {
		return m.Hours
	}
	return 0
}

func (m *TimeOfDay) GetMinutes() int32 {
	if m != nil {
		return m.Minutes
	}
	return 0
}

func (m *TimeOfDay) GetSeconds() int32 {
	if m != nil {
		return m.Seconds
	}
	return 0
}

func (m *TimeOfDay) GetMilliseconds() int32 {
	if m != nil {
		return m.Milliseconds
	}
	return 0
}

type RowValue struct {
	// Types that are valid to be assigned to Value:
	//	*RowValue_Str
	//	*RowValue_Int
	//	*RowValue_Real
	//	*RowValue_Date
	//	*RowValue_DateOfTime
	//	*RowValue_TimeOfDay
	Value isRowValue_Value `protobuf_oneof:"value"`
}

func (m *RowValue) Reset()                    { *m = RowValue{} }
func (m *RowValue) String() string            { return proto.CompactTextString(m) }
func (*RowValue) ProtoMessage()               {}
func (*RowValue) Descriptor() ([]byte, []int) { return fileDescriptorDatasetmodels, []int{2} }

type isRowValue_Value interface {
	isRowValue_Value()
	MarshalTo([]byte) (int, error)
	Size() int
}

type RowValue_Str struct {
	Str string `protobuf:"bytes,1,opt,name=str,proto3,oneof"`
}
type RowValue_Int struct {
	Int int32 `protobuf:"varint,2,opt,name=int,proto3,oneof"`
}
type RowValue_Real struct {
	Real float32 `protobuf:"fixed32,3,opt,name=real,proto3,oneof"`
}
type RowValue_Date struct {
	Date int64 `protobuf:"varint,4,opt,name=date,proto3,oneof"`
}
type RowValue_DateOfTime struct {
	DateOfTime int64 `protobuf:"varint,5,opt,name=date_of_time,json=dateOfTime,proto3,oneof"`
}
type RowValue_TimeOfDay struct {
	TimeOfDay *TimeOfDay `protobuf:"bytes,6,opt,name=time_of_day,json=timeOfDay,oneof"`
}

func (*RowValue_Str) isRowValue_Value()        {}
func (*RowValue_Int) isRowValue_Value()        {}
func (*RowValue_Real) isRowValue_Value()       {}
func (*RowValue_Date) isRowValue_Value()       {}
func (*RowValue_DateOfTime) isRowValue_Value() {}
func (*RowValue_TimeOfDay) isRowValue_Value()  {}

func (m *RowValue) GetValue() isRowValue_Value {
	if m != nil {
		return m.Value
	}
	return nil
}

func (m *RowValue) GetStr() string {
	if x, ok := m.GetValue().(*RowValue_Str); ok {
		return x.Str
	}
	return ""
}

func (m *RowValue) GetInt() int32 {
	if x, ok := m.GetValue().(*RowValue_Int); ok {
		return x.Int
	}
	return 0
}

func (m *RowValue) GetReal() float32 {
	if x, ok := m.GetValue().(*RowValue_Real); ok {
		return x.Real
	}
	return 0
}

func (m *RowValue) GetDate() int64 {
	if x, ok := m.GetValue().(*RowValue_Date); ok {
		return x.Date
	}
	return 0
}

func (m *RowValue) GetDateOfTime() int64 {
	if x, ok := m.GetValue().(*RowValue_DateOfTime); ok {
		return x.DateOfTime
	}
	return 0
}

func (m *RowValue) GetTimeOfDay() *TimeOfDay {
	if x, ok := m.GetValue().(*RowValue_TimeOfDay); ok {
		return x.TimeOfDay
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*RowValue) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _RowValue_OneofMarshaler, _RowValue_OneofUnmarshaler, _RowValue_OneofSizer, []interface{}{
		(*RowValue_Str)(nil),
		(*RowValue_Int)(nil),
		(*RowValue_Real)(nil),
		(*RowValue_Date)(nil),
		(*RowValue_DateOfTime)(nil),
		(*RowValue_TimeOfDay)(nil),
	}
}

func _RowValue_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*RowValue)
	// value
	switch x := m.Value.(type) {
	case *RowValue_Str:
		_ = b.EncodeVarint(1<<3 | proto.WireBytes)
		_ = b.EncodeStringBytes(x.Str)
	case *RowValue_Int:
		_ = b.EncodeVarint(2<<3 | proto.WireVarint)
		_ = b.EncodeVarint(uint64(x.Int))
	case *RowValue_Real:
		_ = b.EncodeVarint(3<<3 | proto.WireFixed32)
		_ = b.EncodeFixed32(uint64(math.Float32bits(x.Real)))
	case *RowValue_Date:
		_ = b.EncodeVarint(4<<3 | proto.WireVarint)
		_ = b.EncodeVarint(uint64(x.Date))
	case *RowValue_DateOfTime:
		_ = b.EncodeVarint(5<<3 | proto.WireVarint)
		_ = b.EncodeVarint(uint64(x.DateOfTime))
	case *RowValue_TimeOfDay:
		_ = b.EncodeVarint(6<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.TimeOfDay); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("RowValue.Value has unexpected type %T", x)
	}
	return nil
}

func _RowValue_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*RowValue)
	switch tag {
	case 1: // value.str
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.Value = &RowValue_Str{x}
		return true, err
	case 2: // value.int
		if wire != proto.WireVarint {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeVarint()
		m.Value = &RowValue_Int{int32(x)}
		return true, err
	case 3: // value.real
		if wire != proto.WireFixed32 {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeFixed32()
		m.Value = &RowValue_Real{math.Float32frombits(uint32(x))}
		return true, err
	case 4: // value.date
		if wire != proto.WireVarint {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeVarint()
		m.Value = &RowValue_Date{int64(x)}
		return true, err
	case 5: // value.date_of_time
		if wire != proto.WireVarint {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeVarint()
		m.Value = &RowValue_DateOfTime{int64(x)}
		return true, err
	case 6: // value.time_of_day
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(TimeOfDay)
		err := b.DecodeMessage(msg)
		m.Value = &RowValue_TimeOfDay{msg}
		return true, err
	default:
		return false, nil
	}
}

func _RowValue_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*RowValue)
	// value
	switch x := m.Value.(type) {
	case *RowValue_Str:
		n += proto.SizeVarint(1<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(len(x.Str)))
		n += len(x.Str)
	case *RowValue_Int:
		n += proto.SizeVarint(2<<3 | proto.WireVarint)
		n += proto.SizeVarint(uint64(x.Int))
	case *RowValue_Real:
		n += proto.SizeVarint(3<<3 | proto.WireFixed32)
		n += 4
	case *RowValue_Date:
		n += proto.SizeVarint(4<<3 | proto.WireVarint)
		n += proto.SizeVarint(uint64(x.Date))
	case *RowValue_DateOfTime:
		n += proto.SizeVarint(5<<3 | proto.WireVarint)
		n += proto.SizeVarint(uint64(x.DateOfTime))
	case *RowValue_TimeOfDay:
		s := proto.Size(x.TimeOfDay)
		n += proto.SizeVarint(6<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type Row struct {
	Values []*RowValue `protobuf:"bytes,1,rep,name=values" json:"values,omitempty"`
}

func (m *Row) Reset()                    { *m = Row{} }
func (m *Row) String() string            { return proto.CompactTextString(m) }
func (*Row) ProtoMessage()               {}
func (*Row) Descriptor() ([]byte, []int) { return fileDescriptorDatasetmodels, []int{3} }

func (m *Row) GetValues() []*RowValue {
	if m != nil {
		return m.Values
	}
	return nil
}

type DataSet struct {
	Label   string    `protobuf:"bytes,1,opt,name=label,proto3" json:"label,omitempty"`
	Columns []*Column `protobuf:"bytes,5,rep,name=columns" json:"columns,omitempty"`
	Rows    []*Row    `protobuf:"bytes,6,rep,name=rows" json:"rows,omitempty"`
}

func (m *DataSet) Reset()                    { *m = DataSet{} }
func (m *DataSet) String() string            { return proto.CompactTextString(m) }
func (*DataSet) ProtoMessage()               {}
func (*DataSet) Descriptor() ([]byte, []int) { return fileDescriptorDatasetmodels, []int{4} }

func (m *DataSet) GetLabel() string {
	if m != nil {
		return m.Label
	}
	return ""
}

func (m *DataSet) GetColumns() []*Column {
	if m != nil {
		return m.Columns
	}
	return nil
}

func (m *DataSet) GetRows() []*Row {
	if m != nil {
		return m.Rows
	}
	return nil
}

func init() {
	proto.RegisterType((*Column)(nil), "apipb.Column")
	proto.RegisterType((*TimeOfDay)(nil), "apipb.TimeOfDay")
	proto.RegisterType((*RowValue)(nil), "apipb.RowValue")
	proto.RegisterType((*Row)(nil), "apipb.Row")
	proto.RegisterType((*DataSet)(nil), "apipb.DataSet")
	proto.RegisterEnum("apipb.Column_Type", Column_Type_name, Column_Type_value)
}
func (m *Column) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Column) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Type != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintDatasetmodels(dAtA, i, uint64(m.Type))
	}
	if len(m.Name) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintDatasetmodels(dAtA, i, uint64(len(m.Name)))
		i += copy(dAtA[i:], m.Name)
	}
	return i, nil
}

func (m *TimeOfDay) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TimeOfDay) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Hours != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintDatasetmodels(dAtA, i, uint64(m.Hours))
	}
	if m.Minutes != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintDatasetmodels(dAtA, i, uint64(m.Minutes))
	}
	if m.Seconds != 0 {
		dAtA[i] = 0x18
		i++
		i = encodeVarintDatasetmodels(dAtA, i, uint64(m.Seconds))
	}
	if m.Milliseconds != 0 {
		dAtA[i] = 0x20
		i++
		i = encodeVarintDatasetmodels(dAtA, i, uint64(m.Milliseconds))
	}
	return i, nil
}

func (m *RowValue) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RowValue) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Value != nil {
		nn1, err := m.Value.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += nn1
	}
	return i, nil
}

func (m *RowValue_Str) MarshalTo(dAtA []byte) (int, error) {
	i := 0
	dAtA[i] = 0xa
	i++
	i = encodeVarintDatasetmodels(dAtA, i, uint64(len(m.Str)))
	i += copy(dAtA[i:], m.Str)
	return i, nil
}
func (m *RowValue_Int) MarshalTo(dAtA []byte) (int, error) {
	i := 0
	dAtA[i] = 0x10
	i++
	i = encodeVarintDatasetmodels(dAtA, i, uint64(m.Int))
	return i, nil
}
func (m *RowValue_Real) MarshalTo(dAtA []byte) (int, error) {
	i := 0
	dAtA[i] = 0x1d
	i++
	i = encodeFixed32Datasetmodels(dAtA, i, uint32(math.Float32bits(float32(m.Real))))
	return i, nil
}
func (m *RowValue_Date) MarshalTo(dAtA []byte) (int, error) {
	i := 0
	dAtA[i] = 0x20
	i++
	i = encodeVarintDatasetmodels(dAtA, i, uint64(m.Date))
	return i, nil
}
func (m *RowValue_DateOfTime) MarshalTo(dAtA []byte) (int, error) {
	i := 0
	dAtA[i] = 0x28
	i++
	i = encodeVarintDatasetmodels(dAtA, i, uint64(m.DateOfTime))
	return i, nil
}
func (m *RowValue_TimeOfDay) MarshalTo(dAtA []byte) (int, error) {
	i := 0
	if m.TimeOfDay != nil {
		dAtA[i] = 0x32
		i++
		i = encodeVarintDatasetmodels(dAtA, i, uint64(m.TimeOfDay.Size()))
		n2, err := m.TimeOfDay.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n2
	}
	return i, nil
}
func (m *Row) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Row) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Values) > 0 {
		for _, msg := range m.Values {
			dAtA[i] = 0xa
			i++
			i = encodeVarintDatasetmodels(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	return i, nil
}

func (m *DataSet) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *DataSet) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Label) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintDatasetmodels(dAtA, i, uint64(len(m.Label)))
		i += copy(dAtA[i:], m.Label)
	}
	if len(m.Columns) > 0 {
		for _, msg := range m.Columns {
			dAtA[i] = 0x2a
			i++
			i = encodeVarintDatasetmodels(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	if len(m.Rows) > 0 {
		for _, msg := range m.Rows {
			dAtA[i] = 0x32
			i++
			i = encodeVarintDatasetmodels(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	return i, nil
}

func encodeFixed64Datasetmodels(dAtA []byte, offset int, v uint64) int {
	dAtA[offset] = uint8(v)
	dAtA[offset+1] = uint8(v >> 8)
	dAtA[offset+2] = uint8(v >> 16)
	dAtA[offset+3] = uint8(v >> 24)
	dAtA[offset+4] = uint8(v >> 32)
	dAtA[offset+5] = uint8(v >> 40)
	dAtA[offset+6] = uint8(v >> 48)
	dAtA[offset+7] = uint8(v >> 56)
	return offset + 8
}
func encodeFixed32Datasetmodels(dAtA []byte, offset int, v uint32) int {
	dAtA[offset] = uint8(v)
	dAtA[offset+1] = uint8(v >> 8)
	dAtA[offset+2] = uint8(v >> 16)
	dAtA[offset+3] = uint8(v >> 24)
	return offset + 4
}
func encodeVarintDatasetmodels(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *Column) Size() (n int) {
	var l int
	_ = l
	if m.Type != 0 {
		n += 1 + sovDatasetmodels(uint64(m.Type))
	}
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovDatasetmodels(uint64(l))
	}
	return n
}

func (m *TimeOfDay) Size() (n int) {
	var l int
	_ = l
	if m.Hours != 0 {
		n += 1 + sovDatasetmodels(uint64(m.Hours))
	}
	if m.Minutes != 0 {
		n += 1 + sovDatasetmodels(uint64(m.Minutes))
	}
	if m.Seconds != 0 {
		n += 1 + sovDatasetmodels(uint64(m.Seconds))
	}
	if m.Milliseconds != 0 {
		n += 1 + sovDatasetmodels(uint64(m.Milliseconds))
	}
	return n
}

func (m *RowValue) Size() (n int) {
	var l int
	_ = l
	if m.Value != nil {
		n += m.Value.Size()
	}
	return n
}

func (m *RowValue_Str) Size() (n int) {
	var l int
	_ = l
	l = len(m.Str)
	n += 1 + l + sovDatasetmodels(uint64(l))
	return n
}
func (m *RowValue_Int) Size() (n int) {
	var l int
	_ = l
	n += 1 + sovDatasetmodels(uint64(m.Int))
	return n
}
func (m *RowValue_Real) Size() (n int) {
	var l int
	_ = l
	n += 5
	return n
}
func (m *RowValue_Date) Size() (n int) {
	var l int
	_ = l
	n += 1 + sovDatasetmodels(uint64(m.Date))
	return n
}
func (m *RowValue_DateOfTime) Size() (n int) {
	var l int
	_ = l
	n += 1 + sovDatasetmodels(uint64(m.DateOfTime))
	return n
}
func (m *RowValue_TimeOfDay) Size() (n int) {
	var l int
	_ = l
	if m.TimeOfDay != nil {
		l = m.TimeOfDay.Size()
		n += 1 + l + sovDatasetmodels(uint64(l))
	}
	return n
}
func (m *Row) Size() (n int) {
	var l int
	_ = l
	if len(m.Values) > 0 {
		for _, e := range m.Values {
			l = e.Size()
			n += 1 + l + sovDatasetmodels(uint64(l))
		}
	}
	return n
}

func (m *DataSet) Size() (n int) {
	var l int
	_ = l
	l = len(m.Label)
	if l > 0 {
		n += 1 + l + sovDatasetmodels(uint64(l))
	}
	if len(m.Columns) > 0 {
		for _, e := range m.Columns {
			l = e.Size()
			n += 1 + l + sovDatasetmodels(uint64(l))
		}
	}
	if len(m.Rows) > 0 {
		for _, e := range m.Rows {
			l = e.Size()
			n += 1 + l + sovDatasetmodels(uint64(l))
		}
	}
	return n
}

func sovDatasetmodels(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozDatasetmodels(x uint64) (n int) {
	return sovDatasetmodels(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Column) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowDatasetmodels
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
			return fmt.Errorf("proto: Column: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Column: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Type", wireType)
			}
			m.Type = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDatasetmodels
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Type |= (Column_Type(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDatasetmodels
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
				return ErrInvalidLengthDatasetmodels
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipDatasetmodels(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthDatasetmodels
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
func (m *TimeOfDay) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowDatasetmodels
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
			return fmt.Errorf("proto: TimeOfDay: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TimeOfDay: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Hours", wireType)
			}
			m.Hours = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDatasetmodels
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Hours |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Minutes", wireType)
			}
			m.Minutes = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDatasetmodels
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Minutes |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Seconds", wireType)
			}
			m.Seconds = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDatasetmodels
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Seconds |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Milliseconds", wireType)
			}
			m.Milliseconds = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDatasetmodels
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Milliseconds |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipDatasetmodels(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthDatasetmodels
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
func (m *RowValue) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowDatasetmodels
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
			return fmt.Errorf("proto: RowValue: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RowValue: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Str", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDatasetmodels
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
				return ErrInvalidLengthDatasetmodels
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Value = &RowValue_Str{string(dAtA[iNdEx:postIndex])}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Int", wireType)
			}
			var v int32
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDatasetmodels
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Value = &RowValue_Int{v}
		case 3:
			if wireType != 5 {
				return fmt.Errorf("proto: wrong wireType = %d for field Real", wireType)
			}
			var v uint32
			if (iNdEx + 4) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += 4
			v = uint32(dAtA[iNdEx-4])
			v |= uint32(dAtA[iNdEx-3]) << 8
			v |= uint32(dAtA[iNdEx-2]) << 16
			v |= uint32(dAtA[iNdEx-1]) << 24
			m.Value = &RowValue_Real{float32(math.Float32frombits(v))}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Date", wireType)
			}
			var v int64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDatasetmodels
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Value = &RowValue_Date{v}
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field DateOfTime", wireType)
			}
			var v int64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDatasetmodels
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Value = &RowValue_DateOfTime{v}
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TimeOfDay", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDatasetmodels
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
				return ErrInvalidLengthDatasetmodels
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &TimeOfDay{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Value = &RowValue_TimeOfDay{v}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipDatasetmodels(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthDatasetmodels
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
func (m *Row) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowDatasetmodels
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
			return fmt.Errorf("proto: Row: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Row: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Values", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDatasetmodels
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
				return ErrInvalidLengthDatasetmodels
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Values = append(m.Values, &RowValue{})
			if err := m.Values[len(m.Values)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipDatasetmodels(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthDatasetmodels
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
func (m *DataSet) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowDatasetmodels
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
			return fmt.Errorf("proto: DataSet: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: DataSet: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Label", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDatasetmodels
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
				return ErrInvalidLengthDatasetmodels
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Label = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Columns", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDatasetmodels
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
				return ErrInvalidLengthDatasetmodels
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Columns = append(m.Columns, &Column{})
			if err := m.Columns[len(m.Columns)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Rows", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDatasetmodels
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
				return ErrInvalidLengthDatasetmodels
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Rows = append(m.Rows, &Row{})
			if err := m.Rows[len(m.Rows)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipDatasetmodels(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthDatasetmodels
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
func skipDatasetmodels(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowDatasetmodels
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
					return 0, ErrIntOverflowDatasetmodels
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
					return 0, ErrIntOverflowDatasetmodels
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
				return 0, ErrInvalidLengthDatasetmodels
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowDatasetmodels
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
				next, err := skipDatasetmodels(dAtA[start:])
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
	ErrInvalidLengthDatasetmodels = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowDatasetmodels   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("datasetmodels.proto", fileDescriptorDatasetmodels) }

var fileDescriptorDatasetmodels = []byte{
	// 477 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x92, 0xcf, 0x8e, 0xd3, 0x30,
	0x10, 0xc6, 0xeb, 0xcd, 0x9f, 0xb6, 0x53, 0x96, 0x46, 0x66, 0x0f, 0x39, 0x55, 0x55, 0x0e, 0x6c,
	0x2f, 0xe4, 0x50, 0x9e, 0xa0, 0x4b, 0xcb, 0xb6, 0x12, 0x6c, 0x91, 0x1b, 0x21, 0xc1, 0x25, 0x72,
	0x1b, 0x97, 0x8d, 0x94, 0xc4, 0x51, 0xec, 0x52, 0xe5, 0x4d, 0xe0, 0x69, 0xb8, 0x72, 0xe4, 0x11,
	0x50, 0x79, 0x11, 0x34, 0x4e, 0x02, 0xda, 0x53, 0xe6, 0xfb, 0x7e, 0x93, 0xcc, 0x7c, 0x8e, 0xe1,
	0x45, 0xc2, 0x35, 0x57, 0x42, 0xe7, 0x32, 0x11, 0x99, 0x0a, 0xcb, 0x4a, 0x6a, 0x49, 0x1d, 0x5e,
	0xa6, 0xe5, 0x3e, 0xf8, 0x4e, 0xc0, 0x7d, 0x23, 0xb3, 0x53, 0x5e, 0xd0, 0x97, 0x60, 0xeb, 0xba,
	0x14, 0x3e, 0x99, 0x92, 0xd9, 0xf3, 0x39, 0x0d, 0x4d, 0x43, 0xd8, 0xc0, 0x30, 0xaa, 0x4b, 0xc1,
	0x0c, 0xa7, 0x14, 0xec, 0x82, 0xe7, 0xc2, 0xbf, 0x9a, 0x92, 0xd9, 0x90, 0x99, 0x3a, 0xd8, 0x81,
	0x8d, 0x1d, 0x14, 0xc0, 0xdd, 0x45, 0x6c, 0xf3, 0x70, 0xef, 0xf5, 0xe8, 0x08, 0xfa, 0x9b, 0x87,
	0x68, 0x75, 0xbf, 0x62, 0x1e, 0xa1, 0x03, 0xb0, 0xd9, 0x6a, 0xf1, 0xce, 0xbb, 0xc2, 0x6a, 0xb9,
	0x88, 0x56, 0x9e, 0x45, 0xaf, 0x61, 0x88, 0x55, 0x1c, 0x6d, 0xde, 0xaf, 0x3c, 0x9b, 0x8e, 0x61,
	0x84, 0x55, 0xbc, 0x7d, 0x1b, 0x2f, 0x17, 0x9f, 0x3c, 0x27, 0xa8, 0x61, 0x18, 0xa5, 0xb9, 0xd8,
	0x1e, 0x97, 0xbc, 0xa6, 0x37, 0xe0, 0x3c, 0xca, 0x53, 0xa5, 0xcc, 0x7a, 0x0e, 0x6b, 0x04, 0xf5,
	0xa1, 0x9f, 0xa7, 0xc5, 0x49, 0x0b, 0x65, 0xd6, 0x71, 0x58, 0x27, 0x91, 0x28, 0x71, 0x90, 0x45,
	0xa2, 0x7c, 0xab, 0x21, 0xad, 0xa4, 0x01, 0x3c, 0xcb, 0xd3, 0x2c, 0x4b, 0x3b, 0x6c, 0x1b, 0xfc,
	0xc4, 0x0b, 0x7e, 0x10, 0x18, 0x30, 0x79, 0xfe, 0xc8, 0xb3, 0x13, 0x06, 0xb6, 0x94, 0xae, 0xcc,
	0xe0, 0xe1, 0xba, 0xc7, 0x50, 0xa0, 0x97, 0x16, 0xba, 0x19, 0x8a, 0x5e, 0x5a, 0x68, 0x7a, 0x03,
	0x76, 0x25, 0x78, 0x66, 0xe6, 0x5d, 0xad, 0x7b, 0xcc, 0x28, 0x74, 0x13, 0xae, 0x85, 0x19, 0x63,
	0xa1, 0x8b, 0x0a, 0x97, 0xc0, 0x67, 0x2c, 0x8f, 0xb1, 0x4e, 0x73, 0xe1, 0x3b, 0x2d, 0x05, 0x74,
	0xb7, 0x47, 0xcc, 0x4d, 0xe7, 0x30, 0x42, 0x86, 0x3d, 0x09, 0xaf, 0x7d, 0x77, 0x4a, 0x66, 0xa3,
	0xb9, 0xd7, 0xfe, 0x97, 0x7f, 0x27, 0xb3, 0xee, 0xb1, 0xa1, 0xee, 0xc4, 0x5d, 0x1f, 0x9c, 0xaf,
	0xb8, 0x74, 0x10, 0x82, 0xc5, 0xe4, 0x99, 0xde, 0x82, 0x6b, 0x34, 0x9e, 0x9b, 0x35, 0x1b, 0xcd,
	0xc7, 0xed, 0xeb, 0x5d, 0x38, 0xd6, 0xe2, 0xe0, 0x11, 0xfa, 0x4b, 0xae, 0xf9, 0x4e, 0x60, 0x0e,
	0x27, 0xe3, 0x7b, 0x91, 0x35, 0x89, 0x59, 0x23, 0xe8, 0x2d, 0xf4, 0x0f, 0xe6, 0x2e, 0x28, 0xdf,
	0x31, 0x9f, 0xba, 0x7e, 0x72, 0x43, 0x58, 0x47, 0xe9, 0x04, 0xec, 0x4a, 0x9e, 0x95, 0xef, 0x9a,
	0x2e, 0xf8, 0x3f, 0x90, 0x19, 0xff, 0xee, 0xd5, 0xcf, 0xcb, 0x84, 0xfc, 0xba, 0x4c, 0xc8, 0xef,
	0xcb, 0x84, 0x7c, 0xfb, 0x33, 0xe9, 0xc1, 0xf8, 0x20, 0xf3, 0x50, 0x6a, 0x95, 0xe6, 0x32, 0xfc,
	0x52, 0x95, 0x87, 0x0f, 0xe4, 0xf3, 0xa0, 0x91, 0xe5, 0x7e, 0xef, 0x9a, 0xfb, 0xfa, 0xfa, 0x6f,
	0x00, 0x00, 0x00, 0xff, 0xff, 0x92, 0x12, 0x8e, 0x1a, 0xc6, 0x02, 0x00, 0x00,
}
