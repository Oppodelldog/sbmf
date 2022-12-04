// this code has been generated by sbmf, do not change it manually
// @formatter:off
package {{.Package}}

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
{{- range $name, $id := .MessageIDs }}
    {{ $name }}ID = int8({{ $id }})
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
// Types
{{- range $name, $type := .Types }}
    {{ $type.Name }} {{ $type.Type }}
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
        {{ .Name }} {{range loop .Dim }}[]{{end}}{{ .Type }}
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
    return e
    }
    *v = {{$name}}(i)

    return nil
{{- end }}
{{- end }}

// Types
{{- range $name, $type := .Types }}
    case *{{ $type.Name }}:
    var t {{ $type.Type }}
    var e=unmarshal(&t,r)
    if e != nil {
    return e
    }
    *v = {{ $type.Name }}(t)
    return nil
{{- end }}

// ListTypes
{{- range $type, $dims := .ListTypes }}
    {{- range $dims }}
    case *{{range loop . }}[]{{end}}{{ $type }}:
        return unmarshalSlice(r,v)
    {{- end}}
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
    return e
    }
    _, e = w.Write(d)
    return e
{{- end }}

// Enums
{{- range $name, $values := .Enums }}
{{- if typeDoesNotExist $name }}
    case {{ $name }}:
    return binary.Write(w, binary.LittleEndian, int32(v))
{{- end }}
{{- end }}

// Types
{{- range $name, $type := .Types }}
    case {{ $type.Name }}:
    return marshal({{ $type.Type }}(v),w)
{{- end }}

// ListTypes
{{- range  $type, $dims := .ListTypes }}
    {{- range $dims }}
    case {{ range loop . }}[]{{ end}}{{ $type }}:
        return marshalSlice(w,v)
    {{- end}}
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
        return e
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
        return nil,e
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
e := marshal(int32(len(vt)),w)
if e != nil {
return e
}
_,e = w.Write([]byte(vt))
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
    {{- range $name, $fields := .Messages }}
        case {{ $name }}:
        messageID = {{ $name }}ID
    {{- end }}
    }

    if messageID == 0 {
        return fmt.Errorf("%w: %T", ErrUnknownMessage, m)
    }

    if e :=binary.Write(w, binary.LittleEndian, messageID); e != nil {
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
    {{- range $name, $fields := .Messages }}
        case {{ $name }}ID:
        var m {{ $name }}
        if e := unmarshal(&m, r); e != nil {
            return nil, e
        }
        return m, nil
    {{- end }}
    }

    return nil, fmt.Errorf("%w: %d", ErrUnknownMessage,id)
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
        var messageData = make([]byte, pr.nextPacketLength)
        if e := binary.Read(&pr.Buffer, binary.LittleEndian, &messageData); e != nil {
            return nil, e
        }
        pr.nextPacketLength = 0
        return ReadMessage(bytes.NewReader(messageData))
    }

    return nil, nil
}

