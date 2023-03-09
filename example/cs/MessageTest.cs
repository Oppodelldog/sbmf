using System;
using System.Collections.Generic;
using System.IO;
using Messages;
using Messages.Extensions;
using NUnit.Framework;

namespace cs
{
    [TestFixture]
    public class MessageTest
    {
        [Test]
        public void TestPrimitive()
        {
            var p = new Primitive
            {
                I32 = 1,
                I64 = 2,
                F32 = 3.3f,
                F64 = 4.4f,
                S = "hello",
                B = true
            };

            var p2 = new Primitive();
            var data = p.MarshalBinary();
            WriteFile("primitive.bin", data);
            var reader = new BinaryReader(new MemoryStream(data));
            p2.UnmarshalBinary(reader);

            Assert.AreEqual(p.I32, p2.I32);
            Assert.AreEqual(p.I64, p2.I64);
            Assert.AreEqual(p.F32, p2.F32);
            Assert.AreEqual(p.F64, p2.F64);
            Assert.AreEqual(p.S, p2.S);
            Assert.AreEqual(p.B, p2.B);
        }

        [Test]
        public void TestPrimitiveLists()
        {
            var p = new PrimitiveLists
            {
                I32 = new[] { 1, 2, 3 },
                I64 = new[] { 4L, 5L, 6L },
                II32 = new[] { new[] { 7, 8, 9 }, new[] { 10, 11, 12 } },
                II64 = new[] { new[] { 13L, 14L, 15L }, new[] { 16L, 17L, 18L } },
                F32 = new[] { Single.MinValue, 0, 0.11f, 0.22f, 0.33f, 0.44f, 0.55f, Single.MaxValue },
                F64 = new[] { Double.MinValue, 0, 0.11d, 0.22d, 0.33d, 0.44d, 0.55d, Double.MaxValue },
                S = new[] { "hello", "world" },
                S2 = new[] { new[] { "hello", "world" }, new[] { "you", "are", "wonderful" } },
                B = new[] { true, false, true }
            };

            var p2 = new PrimitiveLists();
            var data = p.MarshalBinary();
            WriteFile("primitive-lists.bin", data);
            var reader = new BinaryReader(new MemoryStream(data));
            p2.UnmarshalBinary(reader);

            Assert.AreEqual(p.I32, p2.I32);
            Assert.AreEqual(p.I64, p2.I64);
            Assert.AreEqual(p.II32, p2.II32);
            Assert.AreEqual(p.II64, p2.II64);
            Assert.AreEqual(p.F32, p2.F32);
            Assert.AreEqual(p.F64, p2.F64);
            Assert.AreEqual(p.S, p2.S);
            Assert.AreEqual(p.S2, p2.S2);
            Assert.AreEqual(p.B, p2.B);
        }

        [Test]
        public void TestPrimitiveMaps()
        {
            var p = new PrimitiveMaps
            {
                I32 = new Dictionary<int, int> { { 1, 2 }, { 3, 4 } },
                I64 = new Dictionary<long, long> { { 5L, 6L }, { 7L, 8L } },
                F32 = new Dictionary<float, float> { { 9.9f, 10.1f }, { 11.1f, 12.2f } },
                F64 = new Dictionary<double, double> { { 13.3d, 14.4d }, { 15.5d, 16.6d } },
                S = new Dictionary<string, string> { { "hello", "world" }, { "you", "are" } },
                B = new Dictionary<bool, bool> { { true, false }, { false, true } },
                SI32 = new Dictionary<string, int> { { "hello", 1 }, { "world", 2 } },
                SII32 = new Dictionary<string, int[]> { { "twenties", new[] { 20, 21, 22 } }, { "thirties", new[] { 30, 31, 32 } } },
                SOF = new Dictionary<string, OneField>() { { "one", new OneField { S = "field" } }, { "two", new OneField { S = "fields" } } }
            };

            var p2 = new PrimitiveMaps();
            var data = p.MarshalBinary();
            WriteFile("primitive-maps.bin", data);
            var reader = new BinaryReader(new MemoryStream(data));
            p2.UnmarshalBinary(reader);

            Assert.AreEqual(p.I32, p2.I32);
            Assert.AreEqual(p.I64, p2.I64);
            Assert.AreEqual(p.F32, p2.F32);
            Assert.AreEqual(p.F64, p2.F64);
            Assert.AreEqual(p.S, p2.S);
            Assert.AreEqual(p.B, p2.B);
            Assert.AreEqual(p.SI32, p2.SI32);
            Assert.AreEqual(p.SII32, p2.SII32);
            Assert.AreEqual(p.SOF, p2.SOF);
        }

        [Test]
        public void TestAlias()
        {
            var p = new Alias
            {
                MI32 = MyInteger32.Value1,
                MI64 = 2,
                MF32 = 3.3f,
                MF64 = 4.4f,
                MS = "hello",
                E = TestEnum.Value2,
                B = true
            };

            var p2 = new Alias();
            var data = p.MarshalBinary();
            WriteFile("alias.bin", data);
            var reader = new BinaryReader(new MemoryStream(data));
            p2.UnmarshalBinary(reader);

            Assert.AreEqual(p.MI32, p2.MI32);
            Assert.AreEqual(p.MI64, p2.MI64);
            Assert.AreEqual(p.MF32, p2.MF32);
            Assert.AreEqual(p.MF64, p2.MF64);
            Assert.AreEqual(p.MS, p2.MS);
            Assert.AreEqual(p.E, p2.E);
            Assert.AreEqual(p.B, p2.B);
        }

        [Test]
        public void TestAliasLists()
        {
            var p = new AliasLists
            {
                MI32 = new[] { (MyInteger32)1, (MyInteger32)2, (MyInteger32)3 },
                MI64 = new[] { 4L, 5L, 6L },
                MF32 = new[] { Single.MinValue, 0, 0.6f, 0.7f, 0.8f, 0.9f, 1.0f, Single.MaxValue },
                MF64 = new[] { Double.MinValue, 0, 0.6f, 0.7f, 0.8f, 0.9f, 1.0f, Double.MaxValue },
                MS = new[] { "hello", "world" },
                E = new[] { TestEnum.Value1, TestEnum.Value2 },
                B = new[] { true, false, true }
            };

            var p2 = new AliasLists();
            var data = p.MarshalBinary();
            WriteFile("alias-lists.bin", data);
            var reader = new BinaryReader(new MemoryStream(data));
            p2.UnmarshalBinary(reader);

            Assert.AreEqual(p.MI32, p2.MI32);
            Assert.AreEqual(p.MI64, p2.MI64);
            Assert.AreEqual(p.MF32, p2.MF32);
            Assert.AreEqual(p.MF64, p2.MF64);
            Assert.AreEqual(p.MS, p2.MS);
            Assert.AreEqual(p.E, p2.E);
            Assert.AreEqual(p.B, p2.B);
        }

        [Test]
        public void TestFoobar()
        {
            var fb = new Foobar()
            {
                PL = new PrimitiveLists()
                {
                    I32 = new[] { 1, 2, 3 },
                    I64 = new[] { 4L, 5L, 6L },
                    F32 = new[] { 7.7f, 8.8f, 9.9f },
                    F64 = new[] { 10.1d, 11.1d, 12.1d },
                    S = new[] { "hello", "world" },
                    B = new[] { true, false, true }
                },
                P = new Primitive()
                {
                    I32 = 1,
                    I64 = 2,
                    F32 = 3.3f,
                    F64 = 4.4f,
                    S = "hello",
                    B = true
                },
                A = new Alias()
                {
                    MI32 = MyInteger32.Value1,
                    MI64 = 2,
                    MF32 = 3.3f,
                    MF64 = 4.4f,
                    MS = "hello",
                    E = TestEnum.Value2,
                    B = true
                },
                AL = new AliasLists()
                {
                    MI32 = new[] { (MyInteger32)1, (MyInteger32)2, (MyInteger32)3 },
                    MI64 = new[] { 4L, 5L, 6L },
                    MF32 = new[] { 7.7f, 8.8f, 9.9f },
                    MF64 = new[] { 10.1d, 11.1d, 12.1d },
                    MS = new[] { "hello", "world" },
                    E = new[] { TestEnum.Value1, TestEnum.Value2 },
                    B = new[] { true, false, true }
                },
            };

            var data = fb.MarshalBinary();

            System.Console.Write(Directory.GetCurrentDirectory());
            WriteFile("foobar.bin", data);
            var fb2 = new Foobar();
            var reader = new BinaryReader(new MemoryStream(data));
            fb2.UnmarshalBinary(reader);

            Assert.AreEqual(fb.PL.B, fb2.PL.B);
            Assert.AreEqual(fb.PL.F32, fb2.PL.F32);
            Assert.AreEqual(fb.PL.F64, fb2.PL.F64);
            Assert.AreEqual(fb.PL.I32, fb2.PL.I32);
            Assert.AreEqual(fb.PL.I64, fb2.PL.I64);
            Assert.AreEqual(fb.PL.S, fb2.PL.S);

            Assert.AreEqual(fb.P.B, fb2.P.B);
            Assert.AreEqual(fb.P.F32, fb2.P.F32);
            Assert.AreEqual(fb.P.F64, fb2.P.F64);
            Assert.AreEqual(fb.P.I32, fb2.P.I32);
            Assert.AreEqual(fb.P.I64, fb2.P.I64);
            Assert.AreEqual(fb.P.S, fb2.P.S);

            Assert.AreEqual(fb.A.B, fb2.A.B);
            Assert.AreEqual(fb.A.E, fb2.A.E);
            Assert.AreEqual(fb.A.MF32, fb2.A.MF32);
            Assert.AreEqual(fb.A.MF64, fb2.A.MF64);
            Assert.AreEqual(fb.A.MI32, fb2.A.MI32);
            Assert.AreEqual(fb.A.MI64, fb2.A.MI64);
            Assert.AreEqual(fb.A.MS, fb2.A.MS);

            Assert.AreEqual(fb.AL.B, fb2.AL.B);
            Assert.AreEqual(fb.AL.E, fb2.AL.E);
            Assert.AreEqual(fb.AL.MF32, fb2.AL.MF32);
            Assert.AreEqual(fb.AL.MF64, fb2.AL.MF64);
            Assert.AreEqual(fb.AL.MI32, fb2.AL.MI32);
            Assert.AreEqual(fb.AL.MI64, fb2.AL.MI64);
            Assert.AreEqual(fb.AL.MS, fb2.AL.MS);
        }

        [Test]
        public void TestOneFieldList()
        {
            var p = new OneFieldList()
            {
                Fields = new []
                {
                    new OneField{ S = "hello" },
                    new OneField{ S = "world" },
                }
            };

            var p2 = new OneFieldList();
            var data = p.MarshalBinary();

            WriteFile("one-field-list.bin", data);

            var reader = new BinaryReader(new MemoryStream(data));
            p2.UnmarshalBinary(reader);

            Assert.AreEqual(p.Fields, p2.Fields);
        }

        [Test]
        public void TestCrossLanguageReadPrimitive()
        {
            byte[] fileBytes = File.ReadAllBytes("../../../../go/out-primitive.bin");
            var reader = new BinaryReader(new MemoryStream(fileBytes));
            var p = new Primitive();
            p.UnmarshalBinary(reader);

            Assert.AreEqual(1, p.I32);
            Assert.AreEqual(2, p.I64);
            Assert.AreEqual(3.3f, p.F32);
            Assert.AreEqual(4.4d, p.F64);
            Assert.AreEqual("hello", p.S);
            Assert.AreEqual(true, p.B);
        }

        [Test]
        public void TestCrossLanguageReadAlias()
        {
            byte[] fileBytes = File.ReadAllBytes("../../../../go/out-alias.bin");
            var reader = new BinaryReader(new MemoryStream(fileBytes));
            var p = new Alias();
            p.UnmarshalBinary(reader);

            Assert.AreEqual((MyInteger32)1, p.MI32);
            Assert.AreEqual(2, p.MI64);
            Assert.AreEqual(3.3f, p.MF32);
            Assert.AreEqual(4.4d, p.MF64);
            Assert.AreEqual("hello", p.MS);
            Assert.AreEqual(TestEnum.Value2, p.E);
            Assert.AreEqual(true, p.B);
        }

        [Test]
        public void TestCrossLanguageReadPrimitiveLists()
        {
            byte[] fileBytes = File.ReadAllBytes("../../../../go/out-primitive-lists.bin");
            var reader = new BinaryReader(new MemoryStream(fileBytes));
            var p = new PrimitiveLists();
            p.UnmarshalBinary(reader);

            Assert.AreEqual(new[] { Int32.MinValue, 0, Int32.MaxValue }, p.I32);
            Assert.AreEqual(new[] { Int64.MinValue, 0, Int64.MaxValue, }, p.I64);
            Assert.AreEqual(new[] { new[] { Int32.MinValue }, new[] { 42, Int32.MaxValue } }, p.II32);
            Assert.AreEqual(new[] { new[] { Int64.MinValue }, new[] { 42, Int64.MaxValue } }, p.II64);
            Assert.AreEqual(new[] { Single.MinValue, 0, 0.6f, 0.7f, 0.8f, 0.9f, 1.0f, Single.MaxValue }, p.F32);
            Assert.AreEqual(new[] { Double.MinValue, 0, 0.6d, 0.7d, 0.8d, 0.9d, 1.0d, Double.MaxValue }, p.F64);
            Assert.AreEqual(new[] { "hello", "world" }, p.S);
            Assert.AreEqual(new[] { new[] { "hello", "world" }, new[] { "you", "are", "wonderful" } }, p.S2);
            Assert.AreEqual(new[] { true, false }, p.B);
        }

        [Test]
        public void TestCrossLanguageReadPrimitiveMaps()
        {
            byte[] fileBytes = File.ReadAllBytes("../../../../go/out-primitive-maps.bin");
            var reader = new BinaryReader(new MemoryStream(fileBytes));
            var p = new PrimitiveMaps();
            p.UnmarshalBinary(reader);

            Assert.AreEqual(new Dictionary<int, int> { { Int32.MinValue, 1 }, { 2, Int32.MaxValue } }, p.I32);
            Assert.AreEqual(new Dictionary<long, long> { { Int64.MinValue, 3 }, { 4, Int64.MaxValue } }, p.I64);
            Assert.AreEqual(new Dictionary<float, float> { { Single.MinValue, 5.5f }, { 6.6f, Single.MaxValue } }, p.F32);
            //Assert.AreEqual(new Dictionary<double, double> { { Double.MinValue, 7.7f }, { 8.8f, Double.MaxValue } }, p.F64);
            Assert.AreEqual(new Dictionary<string, string> { { "hello", "world" }, { "you", "are" } }, p.S);
            Assert.AreEqual(new Dictionary<bool, bool> { { true, false }, { false, true } }, p.B);
            Assert.AreEqual(new Dictionary<string, int> { { "one", 1 }, { "two", 2 }, { "three", 3 } }, p.SI32);
            Assert.AreEqual(new Dictionary<string, int[]> { { "twenties", new[] { 20, 21, 22 } }, { "thirties", new[] { 30, 31, 32 } } }, p.SII32);
        }

        [Test]
        public void TestCrossLanguageReadAliasLists()
        {
            byte[] fileBytes = File.ReadAllBytes("../../../../go/out-alias-lists.bin");
            var reader = new BinaryReader(new MemoryStream(fileBytes));
            var p = new AliasLists();
            p.UnmarshalBinary(reader);

            Assert.AreEqual(new[] { (MyInteger32)Int32.MinValue, (MyInteger32)0, (MyInteger32)Int32.MaxValue }, p.MI32);
            Assert.AreEqual(new[] { Int64.MinValue, 0, Int64.MaxValue, }, p.MI64);
            Assert.AreEqual(new[] { Single.MinValue, 0, 0.11f, 0.22f, 0.33f, 0.44f, 0.55f, Single.MaxValue }, p.MF32);
            Assert.AreEqual(new[] { Double.MinValue, 0, 0.11d, 0.22d, 0.33d, 0.44d, 0.55d, Double.MaxValue }, p.MF64);
            Assert.AreEqual(new[] { "hello", "world" }, p.MS);
            Assert.AreEqual(new[] { TestEnum.Value1, TestEnum.Value2 }, p.E);
            Assert.AreEqual(new[] { true, false }, p.B);
        }

        [Test]
        public void TestCrossLanguageReadFoobar()
        {
            byte[] fileBytes = File.ReadAllBytes("../../../../go/out-foobar.bin");
            var reader = new BinaryReader(new MemoryStream(fileBytes));
            var p = new Foobar();
            p.UnmarshalBinary(reader);

            Assert.AreEqual(1, p.P.I32);
            Assert.AreEqual(2, p.P.I64);
            Assert.AreEqual(3.3f, p.P.F32);
            Assert.AreEqual(4.4d, p.P.F64);
            Assert.AreEqual("hello", p.P.S);
            Assert.AreEqual(true, p.P.B);

            Assert.AreEqual(new[] { Int32.MinValue, 0, Int32.MaxValue }, p.PL.I32);
            Assert.AreEqual(new[] { Int64.MinValue, 0, Int64.MaxValue, }, p.PL.I64);
            Assert.AreEqual(new[] { Single.MinValue, 0, 0.678912345f, 0.77f, 0.88f, 0.99f, 1.1f, Single.MaxValue }, p.PL.F32);
            Assert.AreEqual(new[] { Double.MinValue, 0, 0.678912345d, 0.77d, 0.88d, 0.99d, 1.1d, Double.MaxValue }, p.PL.F64);
            Assert.AreEqual(new[] { "hello", "world" }, p.PL.S);
            Assert.AreEqual(new[] { true, false }, p.PL.B);

            Assert.AreEqual((MyInteger32)1, p.A.MI32);
            Assert.AreEqual(2, p.A.MI64);
            Assert.AreEqual(3.3f, p.A.MF32);
            Assert.AreEqual(4.4d, p.A.MF64);
            Assert.AreEqual("hello", p.A.MS);
            Assert.AreEqual(TestEnum.Value2, p.A.E);
            Assert.AreEqual(true, p.A.B);

            Assert.AreEqual(new[] { (MyInteger32)Int32.MinValue, (MyInteger32)0, (MyInteger32)Int32.MaxValue }, p.AL.MI32);
            Assert.AreEqual(new[] { Int64.MinValue, 0, Int64.MaxValue, }, p.AL.MI64);
            Assert.AreEqual(new[] { Single.MinValue, 0, 0.111111111112f, 0.22f, 0.33f, 0.44f, 0.55f, Single.MaxValue }, p.AL.MF32);
            Assert.AreEqual(new[] { Double.MinValue, 0, 0.111111111112d, 0.22d, 0.33d, 0.44d, 0.55d, Double.MaxValue }, p.AL.MF64);
            Assert.AreEqual(new[] { "hello", "world" }, p.AL.MS);
            Assert.AreEqual(new[] { TestEnum.Value1, TestEnum.Value2 }, p.AL.E);
            Assert.AreEqual(new[] { true, false }, p.AL.B);
        }

        [Test]
        public void TestCrossLanguageOneFieldList()
        {
            byte[] fileBytes = File.ReadAllBytes("../../../../go/out-one-field-list.bin");
            var reader = new BinaryReader(new MemoryStream(fileBytes));
            var p = new OneFieldList();
            p.UnmarshalBinary(reader);

            Assert.AreEqual(new[] { new OneField { S = "hello" }, new OneField { S = "world" } }, p.Fields);
        }

        [Test]
        public void TestOneFieldWriteMessage()
        {
            var m = new OneField() { S = "hello-world" };
            var writer = new BinaryWriter(new FileStream("../../../out-one-field.bin", FileMode.Create));
            BinaryExtensions.WriteMessage(writer, m);
            writer.Flush();
            writer.Close();
        }

        [Test]
        public void TestOneFieldReadMessage()
        {
            var reader = new BinaryReader(new FileStream("../../../out-one-field.bin", FileMode.Open));
            var m = BinaryExtensions.ReadMessage(reader);
            Assert.AreEqual("hello-world", ((OneField)m).S);
            reader.Close();
        }

        private static void WriteFile(string name, byte[] data)
        {
            File.WriteAllBytes("../../../out-" + name, data);
        }

        [Test]
        public void TestPacketReader()
        {
            var fileBytes = File.ReadAllBytes("../../../out-2packets-one-field.bin");
            var reader = new MemoryStream(fileBytes);

            var pr = new PacketReader();

            var messages = new List<object>();
            foreach (var b in fileBytes)
            {
                var m = pr.Read(new[] { b });
                if (m != null)
                {
                    messages.Add(m);
                }
            }

            Assert.AreEqual(2, messages.Count);
            Assert.AreEqual("hello", ((OneField)messages[0]).S);
            Assert.AreEqual("world", ((OneField)messages[1]).S);
        }

        [Test]
        public void TestWritePackets()
        {
            var m1 = new OneField() { S = "hello" };
            var m2 = new OneField() { S = "world" };
            var writer = new BinaryWriter(new FileStream("../../../out-2packets-one-field.bin", FileMode.Create));
            BinaryExtensions.WritePacket(writer, m1);
            BinaryExtensions.WritePacket(writer, m2);
            writer.Close();
        }
    }
}