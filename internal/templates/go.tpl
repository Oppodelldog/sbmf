// this code has been generated by sbmf (https://github.com/Oppodelldog/sbmf), do not change it manually
// @formatter:off
package {{.Package}}

import (
"bytes"
"encoding/binary"
"errors"
"fmt"
"io"
"reflect"
)

var ErrUnknownMessage = errors.New("unknown message")

const (
    // Messages
{{- range $name, $id := .MessageIDs }}
    {{ messageIdName $name }} = int8({{ $id }})
{{- end }}
)
const (
    // Enums
{{- range $name, $values := .Enums }}
    {{- range $values }}
        {{$name}}{{.Name}} {{$name}} = {{.Value}}
    {{- end }}
{{- end }}
)

type (
    // CustomTypes
{{- range $name, $type := .CustomTypes }}
    {{ .Name }} {{ typeDef . }}
{{- end }}

    // Enums
{{- range $name, $values := .Enums }}
    {{- if typeDoesNotExist $name }}
        {{ $name }} int32
    {{- end }}
{{- end }}

    // Messages
{{- range $name, $fields := .Messages }}
    {{ $name }} struct {
    {{- range $fields }}
        {{ .Name }} {{ typeDef . }}
    {{- end }}
    }
{{- end }}
)

func unmarshal(v interface{}, r io.Reader) error {
    switch v := v.(type) {
    case *string:
        return unmarshalString(r, v)

    // Messages
{{- range $name, $fields := .Messages }}
    case *{{ $name }}:
    return v.UnmarshalBinary(r)
{{- end }}

    // Enums
{{- range $name, $values := .Enums }}
{{- if typeDoesNotExist $name }}
    case *{{$name}}:
    var i int32
    e := binary.Read(r, binary.LittleEndian, &i)
    if e != nil {
        return fmt.Errorf("err unmarshal {{$name}}: %w", e)
    }
    *v = {{$name}}(i)

    return nil
{{- end }}
{{- end }}

    // CustomTypes
{{- range .CustomTypes }}
    case *{{ .Name }}:
    var t {{ typeDef . }}
    var e=unmarshal(&t,r)
    if e != nil {
    return fmt.Errorf("err unmarshal {{ .Name }}: %w", e)
    }
    *v = {{ .Name }}(t)
    return nil
{{- end }}

    // ListTypes
{{- range $type := listTypes }}
    case *{{ typeDef $type }}:
        return unmarshalSlice(r,v)
{{- end }}

    // MapTypes
{{- range $type := mapTypes }}
    case *{{ typeDef $type }}:
    return unmarshalMap(r,v)
{{- end }}

    default:
     return binary.Read(r, binary.LittleEndian, v)
    }
}

func marshal(v interface{}, w io.Writer) error {
    switch v := v.(type) {
    case string:
        return marshalString(w, v)

    // Messages
{{- range $name, $fields := .Messages }}
    case {{ $name }}:
    d, e := v.MarshalBinary()
    if e != nil {
        return fmt.Errorf("err marshal {{ $name }}: %w", e)
    }
    _, e = w.Write(d)
    if e != nil {
        return fmt.Errorf("err write {{ $name }}: %w", e)
    }
    return nil
{{- end }}

    // Enums
{{- range $name, $values := .Enums }}
{{- if typeDoesNotExist $name }}
    case {{ $name }}:
    return binary.Write(w, binary.LittleEndian, int32(v))
{{- end }}
{{- end }}

    // CustomTypes
{{- range  .CustomTypes }}
    case {{ .Name }}:
    return marshal({{ typeDef . }}(v),w)
{{- end }}

    // ListTypes
{{- range $type := listTypes }}
    case {{ typeDef $type }}:
    return marshalSlice(w,v)
{{- end }}

    // MapTypes
{{- range $type := mapTypes }}
    case {{ typeDef $type }}:
    return marshalMap(w,v)
{{- end }}

    default:
        return binary.Write(w, binary.LittleEndian, v)
    }
}

// Messages
{{- range $name, $fields := .Messages }}
    func (m *{{ $name }}) UnmarshalBinary(r io.Reader) error{
    var e error

    {{- range $fields }}
    if e=unmarshal(&m.{{ .Name }}, r); e != nil {
        return fmt.Errorf("err unmarshal m.{{ .Name }}: %w", e)
    }
    {{- end }}

    return nil
    }

    func (m *{{ $name }}) MarshalBinary() ([]byte, error) {
    var data []byte
    w := bytes.NewBuffer(data)
    var e error

    {{- range $fields }}
    if e=marshal(m.{{ .Name }},w); e != nil {
        return nil, fmt.Errorf("err marshal m.{{ .Name }}: %w", e)
    }
    {{- end }}

    return w.Bytes(),nil
    }
{{- end }}


func marshalSlice[T any](w io.Writer, v []T) error {
    var length = int32(len(v))
    if e := binary.Write(w, binary.LittleEndian, length); e != nil {
        return e
    }
    for i := int32(0); i < length; i++ {
        switch vt := any(v[i]).(type) {
        case string:
            if e:= marshalString(w, vt); e != nil {
                return fmt.Errorf("err marshal string: %w", e)
            }
        default:
            if e := marshal(vt, w); e != nil {
                return fmt.Errorf("err marshal slice: %w", e)
            }
        }
    }

    return nil
}

func unmarshalSlice[T any](r io.Reader, v *[]T) error {
    var length int32
    if e := binary.Read(r, binary.LittleEndian, &length); e != nil {
        return fmt.Errorf("err unmarshal slice length: %w", e)
    }
    *v = make([]T, length)
    for i := 0; i < int(length); i++ {
        if e := unmarshal(&(*v)[i], r); e != nil {
            return fmt.Errorf("err unmarshal slice (len=%d): %w", length, e)
        }
    }

    return nil
}

func unmarshalString(r io.Reader, v *string) error {
    var l int32
    if e := binary.Read(r, binary.LittleEndian, &l); e != nil {
        return  fmt.Errorf("err read string length: %w", e)
    }

    var b = make([]byte, l)
    e := binary.Read(r, binary.LittleEndian, &b)
    if e != nil {
        return  fmt.Errorf("err read string value (len=%v): %w",l, e)
    }
    *v = string(b)

    return nil
}

func marshalString(w io.Writer, v string) error {
    if e := binary.Write(w, binary.LittleEndian, int32(len(v))); e != nil {
        return fmt.Errorf("err write string length: %w", e)
    }
    _, e := w.Write([]byte(v))
    if e != nil {
        return fmt.Errorf("err write string value (len=%v): %w",len(v), e)
    }

    return nil
}

func WriteMessage(w io.Writer, m interface{}) error {
var messageID int8
    switch m.(type) {
    {{- range $name, $fields := .Messages }}
        case {{ $name }}:
        messageID = {{ messageIdName $name }}
    {{- end }}
    }

    if messageID == 0 {
        return fmt.Errorf("%w: %T", ErrUnknownMessage, m)
    }

    if e :=binary.Write(w, binary.LittleEndian, messageID); e != nil {
        return fmt.Errorf("err write message id: %w", e)
    }

    return marshal(m, w)
}

func ReadMessage(r io.Reader) (interface{}, error) {
    var id int8
    if e := binary.Read(r, binary.LittleEndian, &id); e != nil {
        return nil, fmt.Errorf("err read message id: %w", e)
    }

    switch id {
    {{- range $name, $fields := .Messages }}
        case {{ messageIdName $name }}:
        var m {{ $name }}
        if e := unmarshal(&m, r); e != nil {
            return nil, fmt.Errorf("err unmarshal {{ $name }}: %w", e)
        }
        return m, nil
    {{- end }}
    }

    return nil, fmt.Errorf("%w: %d", ErrUnknownMessage,id)
}

func ReadPacket(r io.Reader) (interface{}, error) {
    var length int32
    if e := binary.Read(r, binary.LittleEndian, &length); e != nil {
        return nil, fmt.Errorf("err read packet length: %w", e)
    }

    var data = make([]byte, length)
    if e := binary.Read(r, binary.LittleEndian, &data); e != nil {
        return nil, fmt.Errorf("err read packet data (len=%v): %w",length, e)
    }

    return ReadMessage(bytes.NewReader(data))
}

func WritePacket(w io.Writer, m interface{}) error {
    var data []byte
    var buffer = bytes.NewBuffer(data)
    if e := WriteMessage(buffer, m); e != nil {
        return fmt.Errorf("err write message: %w", e)
    }

    var length = int32(buffer.Len())
    if e := binary.Write(w, binary.LittleEndian, length); e != nil {
        return fmt.Errorf("err write packet length: %w", e)
    }

    _, e := w.Write(buffer.Bytes())
    if e != nil{
        return fmt.Errorf("err write packet data (len=%v): %w",length, e)
    }

    return nil
}

type PacketReader struct {
    bytes.Buffer
    nextPacketLength int32
}

func (pr *PacketReader) Read(r io.Reader) (interface{}, error) {
    data, err := io.ReadAll(r)
    if err != nil {
        return nil, fmt.Errorf("err read packet: %w", err)
    }

    _, err = pr.Write(data)
    if err != nil {
        return nil, fmt.Errorf("err write packet: %w", err)
    }

    if pr.nextPacketLength == 0 && pr.Len() >= 4 {
        var length int32
        if err = binary.Read(&pr.Buffer, binary.LittleEndian, &length); err != nil {
            return nil, fmt.Errorf("err read packet length: %w", err)
        }
        pr.nextPacketLength = length
    }

    if pr.nextPacketLength > 0 && int32(pr.Len()) >= pr.nextPacketLength {
        var messageData = make([]byte, pr.nextPacketLength)
        if e := binary.Read(&pr.Buffer, binary.LittleEndian, &messageData); e != nil {
            return nil, fmt.Errorf("err read packet data (len=%v): %w",pr.nextPacketLength, e)
        }
        pr.nextPacketLength = 0
        return ReadMessage(bytes.NewReader(messageData))
    }

    return nil, nil
}


func marshalMap(w io.Writer, mapValue interface{}) error {
    var mapValueReflect = reflect.ValueOf(mapValue)
    var mapLength = mapValueReflect.Len()
    if e := binary.Write(w, binary.LittleEndian, int32(mapLength)); e != nil {
        return fmt.Errorf("err write map length: %w", e)
    }

    for _, key := range mapValueReflect.MapKeys() {
        var value = mapValueReflect.MapIndex(key)
        if e := marshal(key.Interface(), w); e != nil {
            return fmt.Errorf("err marshal map key: %w", e)
        }
        if e := marshal(value.Interface(), w); e != nil {
            return fmt.Errorf("err marshal map value: %w", e)
        }
    }

    return nil
}

func unmarshalMap(r io.Reader, mp interface{}) error {
    var rmp = reflect.ValueOf(mp)
    var mapReflect = reflect.ValueOf(rmp.Elem().Interface())
    var mapType = mapReflect.Type()

    mapValue := reflect.MakeMapWithSize(mapReflect.Type(), 0)

    var mapLength int32
    if e := binary.Read(r, binary.LittleEndian, &mapLength); e != nil {
        return fmt.Errorf("err read map length: %w", e)
    }

    for i := 0; i < int(mapLength); i++ {
        var key = reflect.New(mapType.Key()).Interface()
        if e := unmarshal(key, r); e != nil {
            return fmt.Errorf("err unmarshal map key: %w", e)
        }

        var value = reflect.New(mapType.Elem()).Interface()
        if e := unmarshal(value, r); e != nil {
            return fmt.Errorf("err unmarshal map value: %w", e)
        }
        mapValue.SetMapIndex(reflect.ValueOf(key).Elem(), reflect.ValueOf(value).Elem())
    }

    rmp.Elem().Set(mapValue)

    return nil
}