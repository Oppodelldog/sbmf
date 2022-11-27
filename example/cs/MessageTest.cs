using System;
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
                F32 = new[] { 7.7f, 8.8f, 9.9f },
                F64 = new[] { 10.10d, 11.11d, 12.12d },
                S = new[] { "hello", "world" },
                B = new[] { true, false, true }
            };

            var p2 = new PrimitiveLists();
            var data = p.MarshalBinary();
            WriteFile("primitive-lists.bin", data);
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
        public void TestAlias()
        {
            var p = new Alias
            {
                MI32 = 1,
                MI64 = 2,
                MF32 = 3.3f,
                MF64 = 4.4f,
                MS = "hello",
                E = TestEnum.TestEnumValue2,
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
                MI32 = new[] { 1, 2, 3 },
                MI64 = new[] { 4L, 5L, 6L },
                MF32 = new[] { 7.7f, 8.8f, 9.9f },
                MF64 = new[] { 10.10d, 11.11d, 12.12d },
                MS = new[] { "hello", "world" },
                E = new[] { TestEnum.TestEnumValue1, TestEnum.TestEnumValue2 },
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
                    MI32 = 1,
                    MI64 = 2,
                    MF32 = 3.3f,
                    MF64 = 4.4f,
                    MS = "hello",
                    E = TestEnum.TestEnumValue2,
                    B = true
                },
                AL = new AliasLists()
                {
                    MI32 = new[] { 1, 2, 3 },
                    MI64 = new[] { 4L, 5L, 6L },
                    MF32 = new[] { 7.7f, 8.8f, 9.9f },
                    MF64 = new[] { 10.1d, 11.1d, 12.1d },
                    MS = new[] { "hello", "world" },
                    E = new[] { TestEnum.TestEnumValue1, TestEnum.TestEnumValue2 },
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

            Assert.AreEqual(1, p.MI32);
            Assert.AreEqual(2, p.MI64);
            Assert.AreEqual(3.3f, p.MF32);
            Assert.AreEqual(4.4d, p.MF64);
            Assert.AreEqual("hello", p.MS);
            Assert.AreEqual(TestEnum.TestEnumValue2, p.E);
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
            Assert.AreEqual(new[] { Single.MinValue, 0, Single.MaxValue }, p.F32);
            Assert.AreEqual(new[] { Double.MinValue, 0, Double.MaxValue, }, p.F64);
            Assert.AreEqual(new[] { "hello", "world" }, p.S);
            Assert.AreEqual(new[] { true, false }, p.B);
        }

        [Test]
        public void TestCrossLanguageReadAliasLists()
        {
            byte[] fileBytes = File.ReadAllBytes("../../../../go/out-alias-lists.bin");
            var reader = new BinaryReader(new MemoryStream(fileBytes));
            var p = new AliasLists();
            p.UnmarshalBinary(reader);

            Assert.AreEqual(new[] { Int32.MinValue, 0, Int32.MaxValue }, p.MI32);
            Assert.AreEqual(new[] { Int64.MinValue, 0, Int64.MaxValue, }, p.MI64);
            Assert.AreEqual(new[] { Single.MinValue, 0, Single.MaxValue }, p.MF32);
            Assert.AreEqual(new[] { Double.MinValue, 0, Double.MaxValue, }, p.MF64);
            Assert.AreEqual(new[] { "hello", "world" }, p.MS);
            Assert.AreEqual(new[] { TestEnum.TestEnumValue1, TestEnum.TestEnumValue2 }, p.E);
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
            Assert.AreEqual(new[] { Single.MinValue, 0, Single.MaxValue }, p.PL.F32);
            Assert.AreEqual(new[] { Double.MinValue, 0, Double.MaxValue, }, p.PL.F64);
            Assert.AreEqual(new[] { "hello", "world" }, p.PL.S);
            Assert.AreEqual(new[] { true, false }, p.PL.B);

            Assert.AreEqual(1, p.A.MI32);
            Assert.AreEqual(2, p.A.MI64);
            Assert.AreEqual(3.3f, p.A.MF32);
            Assert.AreEqual(4.4d, p.A.MF64);
            Assert.AreEqual("hello", p.A.MS);
            Assert.AreEqual(TestEnum.TestEnumValue2, p.A.E);
            Assert.AreEqual(true, p.A.B);

            Assert.AreEqual(new[] { Int32.MinValue, 0, Int32.MaxValue }, p.AL.MI32);
            Assert.AreEqual(new[] { Int64.MinValue, 0, Int64.MaxValue, }, p.AL.MI64);
            Assert.AreEqual(new[] { Single.MinValue, 0, Single.MaxValue }, p.AL.MF32);
            Assert.AreEqual(new[] { Double.MinValue, 0, Double.MaxValue, }, p.AL.MF64);
            Assert.AreEqual(new[] { "hello", "world" }, p.AL.MS);
            Assert.AreEqual(new[] { TestEnum.TestEnumValue1, TestEnum.TestEnumValue2 }, p.AL.E);
            Assert.AreEqual(new[] { true, false }, p.AL.B);
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
            var buffer = new MemoryStream();
            var fileBytes = File.ReadAllBytes("../../../out-one-field.bin");
            var binaryWriter = new BinaryWriter(buffer);
            binaryWriter.Write(fileBytes.Length);
            binaryWriter.Write(fileBytes);
            
            var pr = new PacketReader();

            var m = pr.Read(buffer.ToArray());
            
            Assert.AreEqual("hello-world", ((OneField)m).S);
        }
    }
}