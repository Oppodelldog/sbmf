// this code has been generated by sbmf (https://github.com/Oppodelldog/sbmf), do not change it manually
// @formatter:off
using System;
using System.Collections;
using System.Collections.Generic;
using System.IO;

namespace Messages
{
    public enum MyInteger32 {

        Value1 = 1,

        Value2 = 2,
    }
    public enum TestEnum {

        Value1 = 1,

        Value2 = 2,
    }
    public struct Alias {
        public MyInteger32 MI32;
        public long MI64;
        public float MF32;
        public double MF64;
        public string MS;
        public TestEnum E;
        public System.Boolean B;
    }
    public struct AliasLists {
        public MyInteger32[] MI32;
        public long[] MI64;
        public float[] MF32;
        public double[] MF64;
        public string[] MS;
        public TestEnum[] E;
        public System.Boolean[] B;
    }
    public struct Foobar {
        public Alias A;
        public Primitive P;
        public AliasLists AL;
        public PrimitiveLists PL;
    }
    public struct OneField {
        public string S;
    }
    public struct Primitive {
        public int I32;
        public long I64;
        public float F32;
        public double F64;
        public string S;
        public System.Boolean B;
    }
    public struct PrimitiveLists {
        public int[] I32;
        public long[] I64;
        public float[] F32;
        public double[] F64;
        public int[][] II32;
        public long[][] II64;
        public string[] S;
        public string[][] S2;
        public System.Boolean[] B;
    }
    public struct PrimitiveMaps {
        public Dictionary<System.Int32, int> I32;
        public Dictionary<System.Int64, long> I64;
        public Dictionary<System.Single, float> F32;
        public Dictionary<System.Double, double> F64;
        public Dictionary<System.String, string> S;
        public Dictionary<System.Boolean, System.Boolean> B;
        public Dictionary<System.String, int> SI32;
        public Dictionary<System.String, int[]> SII32;
    }
}

namespace Messages.Extensions
{
    public static class BinaryExtensions
    {
        public static byte[] MarshalBinary(this Alias o)
        {
            MemoryStream ms = new MemoryStream();
            BinaryWriter writer = new BinaryWriter(ms);
            writer.Write((int)o.MI32);
            writer.Write(o.MI64);
            writer.Write(o.MF32);
            writer.Write(o.MF64);
            writer.WriteStringSbmf(o.MS);
            writer.Write((int)o.E);
            writer.Write(o.B);
            writer.Flush();

            return ms.ToArray();
        }

        public static void UnmarshalBinary(ref this Alias o,BinaryReader reader)
        {
            o.MI32 = (MyInteger32)reader.ReadInt32();
            o.MI64 = reader.ReadInt64();
            o.MF32 = reader.ReadSingle();
            o.MF64 = reader.ReadDouble();
            o.MS = reader.ReadStringSbmf();
            o.E = (TestEnum)reader.ReadInt32();
            o.B = reader.ReadBoolean();
        }
        public static byte[] MarshalBinary(this AliasLists o)
        {
            MemoryStream ms = new MemoryStream();
            BinaryWriter writer = new BinaryWriter(ms);
            writer.WriteList(o.MI32);
            writer.WriteList(o.MI64);
            writer.WriteList(o.MF32);
            writer.WriteList(o.MF64);
            writer.WriteList(o.MS);
            writer.WriteList(o.E);
            writer.WriteList(o.B);
            writer.Flush();

            return ms.ToArray();
        }

        public static void UnmarshalBinary(ref this AliasLists o,BinaryReader reader)
        {
            o.MI32 = reader.ReadList<MyInteger32>();
            o.MI64 = reader.ReadList<long>();
            o.MF32 = reader.ReadList<float>();
            o.MF64 = reader.ReadList<double>();
            o.MS = reader.ReadList<string>();
            o.E = reader.ReadList<TestEnum>();
            o.B = reader.ReadList<System.Boolean>();
        }
        public static byte[] MarshalBinary(this Foobar o)
        {
            MemoryStream ms = new MemoryStream();
            BinaryWriter writer = new BinaryWriter(ms);
            writer.Write(o.A.MarshalBinary());
            writer.Write(o.P.MarshalBinary());
            writer.Write(o.AL.MarshalBinary());
            writer.Write(o.PL.MarshalBinary());
            writer.Flush();

            return ms.ToArray();
        }

        public static void UnmarshalBinary(ref this Foobar o,BinaryReader reader)
        {
            o.A.UnmarshalBinary(reader);
            o.P.UnmarshalBinary(reader);
            o.AL.UnmarshalBinary(reader);
            o.PL.UnmarshalBinary(reader);
        }
        public static byte[] MarshalBinary(this OneField o)
        {
            MemoryStream ms = new MemoryStream();
            BinaryWriter writer = new BinaryWriter(ms);
            writer.WriteStringSbmf(o.S);
            writer.Flush();

            return ms.ToArray();
        }

        public static void UnmarshalBinary(ref this OneField o,BinaryReader reader)
        {
            o.S = reader.ReadStringSbmf();
        }
        public static byte[] MarshalBinary(this Primitive o)
        {
            MemoryStream ms = new MemoryStream();
            BinaryWriter writer = new BinaryWriter(ms);
            writer.Write(o.I32);
            writer.Write(o.I64);
            writer.Write(o.F32);
            writer.Write(o.F64);
            writer.WriteStringSbmf(o.S);
            writer.Write(o.B);
            writer.Flush();

            return ms.ToArray();
        }

        public static void UnmarshalBinary(ref this Primitive o,BinaryReader reader)
        {
            o.I32 = reader.ReadInt32();
            o.I64 = reader.ReadInt64();
            o.F32 = reader.ReadSingle();
            o.F64 = reader.ReadDouble();
            o.S = reader.ReadStringSbmf();
            o.B = reader.ReadBoolean();
        }
        public static byte[] MarshalBinary(this PrimitiveLists o)
        {
            MemoryStream ms = new MemoryStream();
            BinaryWriter writer = new BinaryWriter(ms);
            writer.WriteList(o.I32);
            writer.WriteList(o.I64);
            writer.WriteList(o.F32);
            writer.WriteList(o.F64);
            writer.WriteList(o.II32);
            writer.WriteList(o.II64);
            writer.WriteList(o.S);
            writer.WriteList(o.S2);
            writer.WriteList(o.B);
            writer.Flush();

            return ms.ToArray();
        }

        public static void UnmarshalBinary(ref this PrimitiveLists o,BinaryReader reader)
        {
            o.I32 = reader.ReadList<int>();
            o.I64 = reader.ReadList<long>();
            o.F32 = reader.ReadList<float>();
            o.F64 = reader.ReadList<double>();
            o.II32 = reader.ReadList<int[]>();
            o.II64 = reader.ReadList<long[]>();
            o.S = reader.ReadList<string>();
            o.S2 = reader.ReadList<string[]>();
            o.B = reader.ReadList<System.Boolean>();
        }
        public static byte[] MarshalBinary(this PrimitiveMaps o)
        {
            MemoryStream ms = new MemoryStream();
            BinaryWriter writer = new BinaryWriter(ms);
            writer.WriteMap(o.I32,writer.Write,writer.Write);
            writer.WriteMap(o.I64,writer.Write,writer.Write);
            writer.WriteMap(o.F32,writer.Write,writer.Write);
            writer.WriteMap(o.F64,writer.Write,writer.Write);
            writer.WriteMap(o.S,writer.WriteStringSbmf,writer.WriteStringSbmf);
            writer.WriteMap(o.B,writer.Write,writer.Write);
            writer.WriteMap(o.SI32,writer.WriteStringSbmf,writer.Write);
            writer.WriteMap(o.SII32,writer.WriteStringSbmf,writer.WriteList);
            writer.Flush();

            return ms.ToArray();
        }

        public static void UnmarshalBinary(ref this PrimitiveMaps o,BinaryReader reader)
        {
            o.I32 = ReadMap(reader, reader.ReadInt32,reader.ReadInt32);
            o.I64 = ReadMap(reader, reader.ReadInt64,reader.ReadInt64);
            o.F32 = ReadMap(reader, reader.ReadSingle,reader.ReadSingle);
            o.F64 = ReadMap(reader, reader.ReadDouble,reader.ReadDouble);
            o.S = ReadMap(reader, reader.ReadStringSbmf,reader.ReadStringSbmf);
            o.B = ReadMap(reader, reader.ReadBoolean,reader.ReadBoolean);
            o.SI32 = ReadMap(reader, reader.ReadStringSbmf,reader.ReadInt32);
            o.SII32 = ReadMap(reader, reader.ReadStringSbmf,reader.ReadList<int>);
        }

    public static Dictionary<TKey,TValue> ReadMap<TKey,TValue>(BinaryReader reader, Func<TKey> readKey,Func<TValue> readValue)
    {
            Dictionary<TKey,TValue> m = new Dictionary<TKey,TValue>();
            int count = reader.ReadInt32();
            for (int i = 0; i < count; i++)
            {
                TKey key = readKey();
                TValue value = readValue();
                m.Add(key, value);
            }

            return m;
    }

    public static void WriteMap<TKey, TValue>(this BinaryWriter writer, Dictionary<TKey, TValue> map, Action<TKey> writeKey,Action<TValue> writeValue)
    {
        writer.Write(map.Count);
        foreach (var item in map)
        {
            writeKey(item.Key);
            writeValue(item.Value);
        }
    }

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
                writer.WriteStringSbmf((string)item);
            }
            else if(item is bool)
            {
                writer.Write((bool)item);
            }
            else if(item is MyInteger32)
            {
                writer.Write((int)item);
            }
            else if(item is TestEnum)
            {
                writer.Write((int)item);
            }
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
                result[i] = (T)(object)reader.ReadStringSbmf();
            }
            else if (typeof(T) == typeof(bool))
            {
                result[i] = (T)(object)reader.ReadBoolean();
            }
            else if (typeof(T) == typeof(MyInteger32))
            {
                result[i] = (T)(object)reader.ReadInt32();
            }
            else if (typeof(T) == typeof(TestEnum))
            {
                result[i] = (T)(object)reader.ReadInt32();
            }
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

    public static void WriteStringSbmf(this BinaryWriter writer, string value)
    {
        writer.Write(value.Length);
        writer.Write(value.ToCharArray());
    }

    public static string ReadStringSbmf(this BinaryReader reader)
    {
        return new string(reader.ReadChars(reader.ReadInt32()));
    }

    public static byte GetMessageId(object message)
    {
        switch (message.GetType())
        {
            case Type t when t == typeof(Alias):
                return 1;
            case Type t when t == typeof(AliasLists):
                return 2;
            case Type t when t == typeof(Foobar):
                return 3;
            case Type t when t == typeof(OneField):
                return 4;
            case Type t when t == typeof(Primitive):
                return 5;
            case Type t when t == typeof(PrimitiveLists):
                return 6;
            case Type t when t == typeof(PrimitiveMaps):
                return 7;
            default:
                throw new Exception("Unknown message type " + message.GetType());
        }
    }

    public static void WriteMessage(BinaryWriter writer, object message)
    {
        writer.Write(GetMessageId(message));
        switch (message.GetType())
        {
        case Type t when t == typeof(Alias):
            writer.Write(((Alias)message).MarshalBinary());
            break;
        case Type t when t == typeof(AliasLists):
            writer.Write(((AliasLists)message).MarshalBinary());
            break;
        case Type t when t == typeof(Foobar):
            writer.Write(((Foobar)message).MarshalBinary());
            break;
        case Type t when t == typeof(OneField):
            writer.Write(((OneField)message).MarshalBinary());
            break;
        case Type t when t == typeof(Primitive):
            writer.Write(((Primitive)message).MarshalBinary());
            break;
        case Type t when t == typeof(PrimitiveLists):
            writer.Write(((PrimitiveLists)message).MarshalBinary());
            break;
        case Type t when t == typeof(PrimitiveMaps):
            writer.Write(((PrimitiveMaps)message).MarshalBinary());
            break;
        default:
            throw new Exception("Unknown message type " + message.GetType());
        }
    }

    public static object ReadMessage(BinaryReader reader){
        var messageId = reader.ReadByte();
        switch (messageId)
        {
            case 1:
            var msgAlias = new Alias();
            msgAlias.UnmarshalBinary(reader);
            return msgAlias;
            case 2:
            var msgAliasLists = new AliasLists();
            msgAliasLists.UnmarshalBinary(reader);
            return msgAliasLists;
            case 3:
            var msgFoobar = new Foobar();
            msgFoobar.UnmarshalBinary(reader);
            return msgFoobar;
            case 4:
            var msgOneField = new OneField();
            msgOneField.UnmarshalBinary(reader);
            return msgOneField;
            case 5:
            var msgPrimitive = new Primitive();
            msgPrimitive.UnmarshalBinary(reader);
            return msgPrimitive;
            case 6:
            var msgPrimitiveLists = new PrimitiveLists();
            msgPrimitiveLists.UnmarshalBinary(reader);
            return msgPrimitiveLists;
            case 7:
            var msgPrimitiveMaps = new PrimitiveMaps();
            msgPrimitiveMaps.UnmarshalBinary(reader);
            return msgPrimitiveMaps;
            default:
                throw new Exception("Unknown message id " + messageId);
        }
    }

    public static void WritePacket(BinaryWriter writer, object message)
    {
        var ms = new MemoryStream();
        var bw = new BinaryWriter(ms);
        WriteMessage(bw, message);
        writer.Write((int)ms.Length);
        writer.Write(ms.ToArray());
    }
}


public class PacketReader
{
    private readonly MemoryStream _buffer = new MemoryStream();
    private int _nextPacketSize = 0;
    public object Read(byte[] data)
    {
        _buffer.Write(data,0,data.Length);

        if (_nextPacketSize == 0 && _buffer.Length >= 4)
        {
            _buffer.Position = 0;
            _nextPacketSize = new BinaryReader(_buffer).ReadInt32();
        }

        if (_nextPacketSize > 0 && _buffer.Length-4 >= _nextPacketSize)
        {
            _buffer.Position = 4;
            var result = BinaryExtensions.ReadMessage(new BinaryReader(_buffer));
            _nextPacketSize = 0;

            if (_buffer.Length > 0)
            {
            var pos = (int)_buffer.Position;
            var cutlen = (int)_buffer.Length-pos;
            var tmp = new byte[cutlen];
            _buffer.Read(tmp, 0, cutlen);
            _buffer.SetLength(0);
            _buffer.Write(tmp, 0, cutlen);
            _buffer.Position = 0;
        }

        return result;
    }

    return null;
}
}
}