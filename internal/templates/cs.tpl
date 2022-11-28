// this code has been generated by sbmf, do not change it manually
// @formatter:off
using System;
using System.Collections;
using System.IO;

namespace {{ .Namespace }}
{
{{- range $name, $type := .Types }}
    using {{ $type.Name }} = {{ $type.Type }};
{{- end }}

{{- range $name, $values := .Enums }}
    public enum {{ $name }} {
    {{- range $values }}

        {{ .Name }} = {{ .Value }},
    {{- end }}
    }
{{- end }}
{{- range $name, $fields := .Messages }}
    public struct {{ $name }} {
    {{- range $fields }}
        public {{ .Type }} {{ .Name }};
    {{- end }}
    }
{{- end }}
}

namespace {{ .Namespace }}.Extensions
{
    public static class BinaryExtensions
    {
    {{- range $name, $fields := .Messages }}
        public static byte[] MarshalBinary(this {{ $name }} o)
        {
            MemoryStream ms = new MemoryStream();
            BinaryWriter writer = new BinaryWriter(ms);
        {{- range $fields }}
            {{- if isStringList .Type }}
            writer.WriteList(o.{{ .Name }});
            {{- else if isString .Type }}
            WriteString(writer, o.{{ .Name }});
            {{- else if isPrimitive .Type }}
            writer.Write(o.{{ .Name }});
            {{- else if isList .Type }}
            writer.WriteList(o.{{ .Name }});
            {{- else if isEnum .Type }}
            writer.Write((int)o.{{ .Name }});
            {{- else if isMessage .Type }}
            writer.Write(o.{{ .Name }}.MarshalBinary());
            {{- else }}
            writer.Write(o.{{ .Name }});
            {{- end}}
        {{- end }}
            writer.Flush();

            return ms.ToArray();
        }

        public static void UnmarshalBinary(ref this {{ $name }} o,BinaryReader reader)
        {
        {{- range $fields }}
            {{- if isStringList .Type }}
            o.{{ .Name }} = reader.ReadList<{{ findPrimitiveType .Type }}>();
            {{- else if isString .Type }}
            o.{{ .Name }} = ReadString(reader);
            {{- else if isPrimitive .Type }}
            o.{{ .Name }} = reader.{{ readFunc .Type }}();
            {{- else if isList .Type }}
            o.{{ .Name }} = reader.ReadList<{{ findPrimitiveType .Type }}>();
            {{- else if isEnum .Type }}
            o.{{ .Name }} = ({{ .Type }})reader.ReadInt32();
            {{- else if isMessage .Type }}
            o.{{ .Name }}.UnmarshalBinary(reader);
            {{- else }}
            o.{{ .Name }} = reader.{{ readFunc .Type }}();
            {{- end}}
        {{- end }}
        }

    {{- end }}


    public static void WriteList(this BinaryWriter writer, IEnumerable list)
    {
        if(list == null)
        {
            writer.Write(0);
            return;
        }

        var length = ((Array)list).Length;
        writer.Write(length);
        foreach (var item in list)
        {
            if(item is int)
            {
                writer.Write((int)item);
            }
            else if(item is long)
            {
                writer.Write((long)item);
            }
            else if(item is float)
            {
                writer.Write((float)item);
            }
            else if(item is double)
            {
                writer.Write((double)item);
            }
            else if(item is string)
            {
                WriteString(writer, (string)item);
            }
            else if(item is bool)
            {
                writer.Write((bool)item);
            }
        {{- range $name, $values := .Enums }}
            else if(item is {{ $name }})
            {
                writer.Write((int)item);
            }
        {{- end }}
            else if(item.GetType().IsArray)
         {
                writer.WriteList((IEnumerable)item);
            }
            else
            {
                throw new Exception("Unknown type");
            }
        }
    }


    public static T[] ReadList<T>(this BinaryReader reader)
    {
        var length = reader.ReadInt32();
        var result = new T[length];

        for (var i = 0; i < length; i++)
        {
            if(typeof(T) == typeof(int))
            {
                result[i] = (T)(object)reader.ReadInt32();
            }
            else if (typeof(T) == typeof(long))
            {
                result[i] = (T)(object)reader.ReadInt64();
            }
            else if (typeof(T) == typeof(float))
            {
                result[i] = (T)(object)reader.ReadSingle();
            }
            else if (typeof(T) == typeof(double))
            {
                result[i] = (T)(object)reader.ReadDouble();
            }
            else if (typeof(T) == typeof(string))
            {
                result[i] = (T)(object)ReadString(reader);
            }
            else if (typeof(T) == typeof(bool))
            {
                result[i] = (T)(object)reader.ReadBoolean();
            }
        {{- range $name, $values := .Enums }}
            else if (typeof(T) == typeof({{ $name }}))
            {
                result[i] = (T)(object)reader.ReadInt32();
            }
        {{- end }}
            else if (typeof(T).IsArray)
            {
                var method = typeof(BinaryExtensions).GetMethod(nameof(ReadList));
                var generic = method.MakeGenericMethod(typeof(T).GetElementType());
                result[i] = (T)generic.Invoke(null, new object[] { reader });
            }
            else
            {
                throw new Exception("Unknown type");
            }
        }

            return result;
    }

    public static void WriteString(BinaryWriter writer, string value)
    {
        writer.Write(value.Length);
        writer.Write(value.ToCharArray());
    }

    public static string ReadString(BinaryReader reader)
    {
        return new string(reader.ReadChars(reader.ReadInt32()));
    }

    public static byte GetMessageId(object message)
    {
        switch (message.GetType())
        {
    {{- range $name, $id := .MessageIDs }}
            case Type t when t == typeof({{ $name }}):
                return {{ $id }};
    {{- end }}
            default:
                throw new Exception("Unknown message type " + message.GetType());
        }
    }

    public static void WriteMessage(BinaryWriter writer, object message)
    {
        writer.Write(GetMessageId(message));
        switch (message.GetType())
        {
    {{- range $name, $message := .Messages }}
        case Type t when t == typeof({{ $name }}):
            writer.Write((({{ $name }})message).MarshalBinary());
            break;
    {{- end }}
        default:
            throw new Exception("Unknown message type " + message.GetType());
        }
    }

    public static object ReadMessage(BinaryReader reader){
        var messageId = reader.ReadByte();
        switch (messageId)
        {
        {{- range $name, $id := .MessageIDs }}
            case {{ $id }}:
            var msg{{ $name }} = new {{ $name }}();
            msg{{ $name }}.UnmarshalBinary(reader);
            return msg{{ $name }};
        {{- end }}
            default:
                throw new Exception("Unknown message id " + messageId);
        }
    }
}


public class PacketReader
{
    private readonly MemoryStream _buffer = new MemoryStream();
    private int _nextPacketSize = 0;

    public object Read(byte[] data)
    {
        _buffer.Write(data, 0, data.Length);

        if (_nextPacketSize == 0 && _buffer.Length >= 4)
        {
            _buffer.Position = 0;
            _nextPacketSize = new BinaryReader(_buffer).ReadInt32();
        }

        if (_nextPacketSize > 0 && _buffer.Length >= _nextPacketSize)
        {
            var result = BinaryExtensions.ReadMessage(new BinaryReader(_buffer));
            _nextPacketSize = 0;

            return result;
        }

        return null;
    }
}
}