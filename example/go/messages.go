// this code has been generated by sbmf, do not change it manually
// @formatter:off
package data

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"io"
)

var ErrUnknownMessage = errors.New("unknown message")

const (
	// Messages
	AliasID          = int8(1)
	AliasListsID     = int8(2)
	FoobarID         = int8(3)
	OneFieldID       = int8(4)
	PrimitiveID      = int8(5)
	PrimitiveListsID = int8(6)
)
const (
	// Enums
	TestEnumTestEnumValue1 TestEnum = 1
	TestEnumTestEnumValue2 TestEnum = 2
)

type (
	// Types
	MyFloat64   float64
	MyInteger32 int32
	MyInteger64 int64
	MyString    string
	MyBoolean   bool
	MyFloat32   float32

	// Enums
	TestEnum int

	// Messages
	Alias struct {
		MI32 MyInteger32
		MI64 MyInteger64
		MF32 MyFloat32
		MF64 MyFloat64
		MS   MyString
		E    TestEnum
		B    MyBoolean
	}
	AliasLists struct {
		MI32 []MyInteger32
		MI64 []MyInteger64
		MF32 []MyFloat32
		MF64 []MyFloat64
		MS   []MyString
		E    []TestEnum
		B    []MyBoolean
	}
	Foobar struct {
		A  Alias
		P  Primitive
		AL AliasLists
		PL PrimitiveLists
	}
	OneField struct {
		S string
	}
	Primitive struct {
		I32 int32
		I64 int64
		F32 float32
		F64 float64
		S   string
		B   bool
	}
	PrimitiveLists struct {
		I32 []int32
		I64 []int64
		F32 []float32
		F64 []float64
		S   []string
		B   []bool
	}
)

func unmarshal(v interface{}, r io.Reader) error {
	switch v := v.(type) {
	case *string:
		return unmarshalString(r, v)

		// Messages
	case *Alias:
		return v.UnmarshalBinary(r)
	case *AliasLists:
		return v.UnmarshalBinary(r)
	case *Foobar:
		return v.UnmarshalBinary(r)
	case *OneField:
		return v.UnmarshalBinary(r)
	case *Primitive:
		return v.UnmarshalBinary(r)
	case *PrimitiveLists:
		return v.UnmarshalBinary(r)

		// Enums
	case *TestEnum:
		var i int32
		e := binary.Read(r, binary.LittleEndian, &i)
		if e != nil {
			return e
		}
		*v = TestEnum(i)

		return nil

		// Types
	case *MyFloat64:
		var t float64
		var e = unmarshal(&t, r)
		if e != nil {
			return e
		}
		*v = MyFloat64(t)
		return nil
	case *MyInteger32:
		var t int32
		var e = unmarshal(&t, r)
		if e != nil {
			return e
		}
		*v = MyInteger32(t)
		return nil
	case *MyInteger64:
		var t int64
		var e = unmarshal(&t, r)
		if e != nil {
			return e
		}
		*v = MyInteger64(t)
		return nil
	case *MyString:
		var t string
		var e = unmarshal(&t, r)
		if e != nil {
			return e
		}
		*v = MyString(t)
		return nil
	case *MyBoolean:
		var t bool
		var e = unmarshal(&t, r)
		if e != nil {
			return e
		}
		*v = MyBoolean(t)
		return nil
	case *MyFloat32:
		var t float32
		var e = unmarshal(&t, r)
		if e != nil {
			return e
		}
		*v = MyFloat32(t)
		return nil

		// ListTypes
	case *[]MyBoolean:
		return unmarshalSlice(r, v)
	case *[]MyFloat32:
		return unmarshalSlice(r, v)
	case *[]MyFloat64:
		return unmarshalSlice(r, v)
	case *[]MyInteger32:
		return unmarshalSlice(r, v)
	case *[]MyInteger64:
		return unmarshalSlice(r, v)
	case *[]MyString:
		return unmarshalSlice(r, v)
	case *[]TestEnum:
		return unmarshalSlice(r, v)
	case *[]bool:
		return unmarshalSlice(r, v)
	case *[]float32:
		return unmarshalSlice(r, v)
	case *[]float64:
		return unmarshalSlice(r, v)
	case *[]int32:
		return unmarshalSlice(r, v)
	case *[]int64:
		return unmarshalSlice(r, v)
	case *[]string:
		return unmarshalSlice(r, v)

	default:
		return binary.Read(r, binary.LittleEndian, v)
	}
}

func marshal(v interface{}, w io.Writer) error {
	switch v := v.(type) {
	case string:
		return marshalString(w, v)

		// Messages
	case Alias:
		d, e := v.MarshalBinary()
		if e != nil {
			return e
		}
		_, e = w.Write(d)
		return e
	case AliasLists:
		d, e := v.MarshalBinary()
		if e != nil {
			return e
		}
		_, e = w.Write(d)
		return e
	case Foobar:
		d, e := v.MarshalBinary()
		if e != nil {
			return e
		}
		_, e = w.Write(d)
		return e
	case OneField:
		d, e := v.MarshalBinary()
		if e != nil {
			return e
		}
		_, e = w.Write(d)
		return e
	case Primitive:
		d, e := v.MarshalBinary()
		if e != nil {
			return e
		}
		_, e = w.Write(d)
		return e
	case PrimitiveLists:
		d, e := v.MarshalBinary()
		if e != nil {
			return e
		}
		_, e = w.Write(d)
		return e

		// Enums
	case TestEnum:
		return binary.Write(w, binary.LittleEndian, int32(v))

		// Types
	case MyFloat64:
		return marshal(float64(v), w)
	case MyInteger32:
		return marshal(int32(v), w)
	case MyInteger64:
		return marshal(int64(v), w)
	case MyString:
		return marshal(string(v), w)
	case MyBoolean:
		return marshal(bool(v), w)
	case MyFloat32:
		return marshal(float32(v), w)

		// ListTypes
	case []MyBoolean:
		return marshalSlice(w, v)
	case []MyFloat32:
		return marshalSlice(w, v)
	case []MyFloat64:
		return marshalSlice(w, v)
	case []MyInteger32:
		return marshalSlice(w, v)
	case []MyInteger64:
		return marshalSlice(w, v)
	case []MyString:
		return marshalSlice(w, v)
	case []TestEnum:
		return marshalSlice(w, v)
	case []bool:
		return marshalSlice(w, v)
	case []float32:
		return marshalSlice(w, v)
	case []float64:
		return marshalSlice(w, v)
	case []int32:
		return marshalSlice(w, v)
	case []int64:
		return marshalSlice(w, v)
	case []string:
		return marshalSlice(w, v)

	default:
		return binary.Write(w, binary.LittleEndian, v)
	}
}

// Messages
func (m *Alias) UnmarshalBinary(r io.Reader) error {
	var e error
	if e = unmarshal(&m.MI32, r); e != nil {
		return e
	}
	if e = unmarshal(&m.MI64, r); e != nil {
		return e
	}
	if e = unmarshal(&m.MF32, r); e != nil {
		return e
	}
	if e = unmarshal(&m.MF64, r); e != nil {
		return e
	}
	if e = unmarshal(&m.MS, r); e != nil {
		return e
	}
	if e = unmarshal(&m.E, r); e != nil {
		return e
	}
	if e = unmarshal(&m.B, r); e != nil {
		return e
	}

	return nil
}

func (m *Alias) MarshalBinary() ([]byte, error) {
	var data []byte
	w := bytes.NewBuffer(data)
	var e error
	if e = marshal(m.MI32, w); e != nil {
		return nil, e
	}
	if e = marshal(m.MI64, w); e != nil {
		return nil, e
	}
	if e = marshal(m.MF32, w); e != nil {
		return nil, e
	}
	if e = marshal(m.MF64, w); e != nil {
		return nil, e
	}
	if e = marshal(m.MS, w); e != nil {
		return nil, e
	}
	if e = marshal(m.E, w); e != nil {
		return nil, e
	}
	if e = marshal(m.B, w); e != nil {
		return nil, e
	}

	return w.Bytes(), nil
}
func (m *AliasLists) UnmarshalBinary(r io.Reader) error {
	var e error
	if e = unmarshal(&m.MI32, r); e != nil {
		return e
	}
	if e = unmarshal(&m.MI64, r); e != nil {
		return e
	}
	if e = unmarshal(&m.MF32, r); e != nil {
		return e
	}
	if e = unmarshal(&m.MF64, r); e != nil {
		return e
	}
	if e = unmarshal(&m.MS, r); e != nil {
		return e
	}
	if e = unmarshal(&m.E, r); e != nil {
		return e
	}
	if e = unmarshal(&m.B, r); e != nil {
		return e
	}

	return nil
}

func (m *AliasLists) MarshalBinary() ([]byte, error) {
	var data []byte
	w := bytes.NewBuffer(data)
	var e error
	if e = marshal(m.MI32, w); e != nil {
		return nil, e
	}
	if e = marshal(m.MI64, w); e != nil {
		return nil, e
	}
	if e = marshal(m.MF32, w); e != nil {
		return nil, e
	}
	if e = marshal(m.MF64, w); e != nil {
		return nil, e
	}
	if e = marshal(m.MS, w); e != nil {
		return nil, e
	}
	if e = marshal(m.E, w); e != nil {
		return nil, e
	}
	if e = marshal(m.B, w); e != nil {
		return nil, e
	}

	return w.Bytes(), nil
}
func (m *Foobar) UnmarshalBinary(r io.Reader) error {
	var e error
	if e = unmarshal(&m.A, r); e != nil {
		return e
	}
	if e = unmarshal(&m.P, r); e != nil {
		return e
	}
	if e = unmarshal(&m.AL, r); e != nil {
		return e
	}
	if e = unmarshal(&m.PL, r); e != nil {
		return e
	}

	return nil
}

func (m *Foobar) MarshalBinary() ([]byte, error) {
	var data []byte
	w := bytes.NewBuffer(data)
	var e error
	if e = marshal(m.A, w); e != nil {
		return nil, e
	}
	if e = marshal(m.P, w); e != nil {
		return nil, e
	}
	if e = marshal(m.AL, w); e != nil {
		return nil, e
	}
	if e = marshal(m.PL, w); e != nil {
		return nil, e
	}

	return w.Bytes(), nil
}
func (m *OneField) UnmarshalBinary(r io.Reader) error {
	var e error
	if e = unmarshal(&m.S, r); e != nil {
		return e
	}

	return nil
}

func (m *OneField) MarshalBinary() ([]byte, error) {
	var data []byte
	w := bytes.NewBuffer(data)
	var e error
	if e = marshal(m.S, w); e != nil {
		return nil, e
	}

	return w.Bytes(), nil
}
func (m *Primitive) UnmarshalBinary(r io.Reader) error {
	var e error
	if e = unmarshal(&m.I32, r); e != nil {
		return e
	}
	if e = unmarshal(&m.I64, r); e != nil {
		return e
	}
	if e = unmarshal(&m.F32, r); e != nil {
		return e
	}
	if e = unmarshal(&m.F64, r); e != nil {
		return e
	}
	if e = unmarshal(&m.S, r); e != nil {
		return e
	}
	if e = unmarshal(&m.B, r); e != nil {
		return e
	}

	return nil
}

func (m *Primitive) MarshalBinary() ([]byte, error) {
	var data []byte
	w := bytes.NewBuffer(data)
	var e error
	if e = marshal(m.I32, w); e != nil {
		return nil, e
	}
	if e = marshal(m.I64, w); e != nil {
		return nil, e
	}
	if e = marshal(m.F32, w); e != nil {
		return nil, e
	}
	if e = marshal(m.F64, w); e != nil {
		return nil, e
	}
	if e = marshal(m.S, w); e != nil {
		return nil, e
	}
	if e = marshal(m.B, w); e != nil {
		return nil, e
	}

	return w.Bytes(), nil
}
func (m *PrimitiveLists) UnmarshalBinary(r io.Reader) error {
	var e error
	if e = unmarshal(&m.I32, r); e != nil {
		return e
	}
	if e = unmarshal(&m.I64, r); e != nil {
		return e
	}
	if e = unmarshal(&m.F32, r); e != nil {
		return e
	}
	if e = unmarshal(&m.F64, r); e != nil {
		return e
	}
	if e = unmarshal(&m.S, r); e != nil {
		return e
	}
	if e = unmarshal(&m.B, r); e != nil {
		return e
	}

	return nil
}

func (m *PrimitiveLists) MarshalBinary() ([]byte, error) {
	var data []byte
	w := bytes.NewBuffer(data)
	var e error
	if e = marshal(m.I32, w); e != nil {
		return nil, e
	}
	if e = marshal(m.I64, w); e != nil {
		return nil, e
	}
	if e = marshal(m.F32, w); e != nil {
		return nil, e
	}
	if e = marshal(m.F64, w); e != nil {
		return nil, e
	}
	if e = marshal(m.S, w); e != nil {
		return nil, e
	}
	if e = marshal(m.B, w); e != nil {
		return nil, e
	}

	return w.Bytes(), nil
}

func marshalSlice[T any](w io.Writer, v []T) error {
	var length = int32(len(v))
	if e := binary.Write(w, binary.LittleEndian, length); e != nil {
		return e
	}
	for i := int32(0); i < length; i++ {
		switch vt := any(v[i]).(type) {
		case string:
			e := marshal(int32(len(vt)), w)
			if e != nil {
				return e
			}
			_, e = w.Write([]byte(vt))
			if e != nil {
				return e
			}
		default:
			if e := marshal(vt, w); e != nil {
				return e
			}
		}
	}

	return nil
}

func unmarshalSlice[T any](r io.Reader, v *[]T) error {
	var length int32
	if e := binary.Read(r, binary.LittleEndian, &length); e != nil {
		return e
	}
	*v = make([]T, length)
	for i := 0; i < int(length); i++ {
		if e := unmarshal(&(*v)[i], r); e != nil {
			return e
		}
	}

	return nil
}

func unmarshalString(r io.Reader, v *string) error {
	var l int32
	if e := binary.Read(r, binary.LittleEndian, &l); e != nil {
		return e
	}

	var b = make([]byte, l)
	e := binary.Read(r, binary.LittleEndian, &b)
	if e != nil {
		return e
	}
	*v = string(b)

	return nil
}

func marshalString(w io.Writer, v string) error {
	if e := binary.Write(w, binary.LittleEndian, int32(len(v))); e != nil {
		return e
	}
	_, e := w.Write([]byte(v))
	if e != nil {
		return e
	}

	return nil
}

func WriteMessage(w io.Writer, m interface{}) error {
	var messageID int8
	switch m.(type) {
	case Alias:
		messageID = AliasID
	case AliasLists:
		messageID = AliasListsID
	case Foobar:
		messageID = FoobarID
	case OneField:
		messageID = OneFieldID
	case Primitive:
		messageID = PrimitiveID
	case PrimitiveLists:
		messageID = PrimitiveListsID
	}

	if messageID == 0 {
		return fmt.Errorf("%w: %T", ErrUnknownMessage, m)
	}

	if e := binary.Write(w, binary.LittleEndian, messageID); e != nil {
		return e
	}

	return marshal(m, w)
}

func ReadMessage(r io.Reader) (interface{}, error) {
	var id int8
	if e := binary.Read(r, binary.LittleEndian, &id); e != nil {
		return nil, e
	}

	switch id {
	case AliasID:
		var m Alias
		if e := unmarshal(&m, r); e != nil {
			return nil, e
		}
		return m, nil
	case AliasListsID:
		var m AliasLists
		if e := unmarshal(&m, r); e != nil {
			return nil, e
		}
		return m, nil
	case FoobarID:
		var m Foobar
		if e := unmarshal(&m, r); e != nil {
			return nil, e
		}
		return m, nil
	case OneFieldID:
		var m OneField
		if e := unmarshal(&m, r); e != nil {
			return nil, e
		}
		return m, nil
	case PrimitiveID:
		var m Primitive
		if e := unmarshal(&m, r); e != nil {
			return nil, e
		}
		return m, nil
	case PrimitiveListsID:
		var m PrimitiveLists
		if e := unmarshal(&m, r); e != nil {
			return nil, e
		}
		return m, nil
	}

	return nil, fmt.Errorf("%w: %d", ErrUnknownMessage, id)
}

func ReadPacket(r io.Reader) (interface{}, error) {
	var length int32
	if e := binary.Read(r, binary.LittleEndian, &length); e != nil {
		return nil, e
	}

	var data = make([]byte, length)
	if e := binary.Read(r, binary.LittleEndian, &data); e != nil {
		return nil, e
	}

	return ReadMessage(bytes.NewReader(data))
}

func WritePacket(w io.Writer, m interface{}) error {
	var data []byte
	var buffer = bytes.NewBuffer(data)
	if e := WriteMessage(buffer, m); e != nil {
		return e
	}

	var length = int32(buffer.Len())
	if e := binary.Write(w, binary.LittleEndian, length); e != nil {
		return e
	}

	_, e := w.Write(buffer.Bytes())
	return e
}

type PacketReader struct {
	bytes.Buffer
	nextPacketLength int32
}

func (pr *PacketReader) Read(r io.Reader) (interface{}, error) {
	data, err := io.ReadAll(r)
	if err != nil {
		return nil, err
	}

	_, err = pr.Write(data)
	if err != nil {
		return nil, err
	}

	if pr.nextPacketLength == 0 && pr.Len() >= 4 {
		var length int32
		if err = binary.Read(&pr.Buffer, binary.LittleEndian, &length); err != nil {
			return nil, err
		}
		pr.nextPacketLength = length
	}

	if pr.nextPacketLength > 0 && int32(pr.Len()) >= pr.nextPacketLength {
		var data = make([]byte, pr.nextPacketLength)
		if e := binary.Read(&pr.Buffer, binary.LittleEndian, &data); e != nil {
			return nil, e
		}
		pr.nextPacketLength = 0
		return ReadMessage(bytes.NewReader(data))
	}

	return nil, nil
}
