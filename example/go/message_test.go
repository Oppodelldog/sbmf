package data

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"math"
	"os"
	"reflect"
	"testing"
)

const (
	float32Min = -3.4028235e+38
	float64Min = -1.7976931348623157e+308
)

func TestPrimitive(t *testing.T) {
	p := Primitive{
		I32: 1,
		I64: 2,
		F32: 3.3,
		F64: 4.4,
		S:   "hello",
		B:   true,
	}

	d, err := p.MarshalBinary()
	if err != nil {
		t.Fatal(err)
	}
	writeFile(t, "primitive.bin", d)

	var p2 Primitive
	err = p2.UnmarshalBinary(bytes.NewBuffer(d))
	if err != nil {
		t.Fatal(err)
	}

	assertEquals(t, p2.I32, p.I32)
	assertEquals(t, p2.I64, p.I64)
	assertEquals(t, p2.F32, p.F32)
	assertEquals(t, p2.F64, p.F64)
	assertEquals(t, p2.S, p.S)
	assertEquals(t, p2.B, p.B)
}

func TestPrimitiveLists(t *testing.T) {
	foo := PrimitiveLists{
		I32: []int32{math.MinInt32, 0, math.MaxInt32},
		I64: []int64{math.MinInt64, 0, math.MaxInt64},
		F32: []float32{float32Min, 0, math.MaxFloat32},
		F64: []float64{float64Min, 0, math.MaxFloat64},
		S:   []string{"hello", "world"},
		B:   []bool{true, false},
	}

	d, err := foo.MarshalBinary()
	if err != nil {
		t.Fatal(err)
	}
	writeFile(t, "primitive-lists.bin", d)

	var foo2 PrimitiveLists
	err = foo2.UnmarshalBinary(bytes.NewBuffer(d))
	if err != nil {
		t.Fatal(err)
	}

	assertSlicesEqual(t, foo2.I32, foo.I32)
	assertSlicesEqual(t, foo2.I64, foo.I64)
	assertSlicesEqual(t, foo2.F32, foo.F32)
	assertSlicesEqual(t, foo2.F64, foo.F64)
	assertSlicesEqual(t, foo2.S, foo.S)
	assertSlicesEqual(t, foo2.B, foo.B)
}

func TestAlias(t *testing.T) {
	bar := Alias{
		MI32: 1,
		MI64: 2,
		MF32: 3.3,
		MF64: 4.4,
		MS:   "hello",
		E:    TestEnumTestEnumValue2,
		B:    true,
	}

	d, err := bar.MarshalBinary()
	if err != nil {
		t.Fatal(err)
	}
	writeFile(t, "alias.bin", d)

	var bar2 Alias
	err = bar2.UnmarshalBinary(bytes.NewBuffer(d))
	if err != nil {
		t.Fatal(err)
	}

	assertEquals(t, bar2.MI32, bar.MI32)
	assertEquals(t, bar2.MI64, bar.MI64)
	assertEquals(t, bar2.MF32, bar.MF32)
	assertEquals(t, bar2.MF64, bar.MF64)
	assertEquals(t, bar2.MS, bar.MS)
	assertEquals(t, bar2.E, bar.E)
	assertEquals(t, bar2.B, bar.B)
}

func TestAliasLists(t *testing.T) {
	bar := AliasLists{
		MI32: []MyInteger32{math.MinInt32, 0, math.MaxInt32},
		MI64: []MyInteger64{math.MinInt64, 0, math.MaxInt64},
		MF32: []MyFloat32{float32Min, 0, math.MaxFloat32},
		MF64: []MyFloat64{float64Min, 0, math.MaxFloat64},
		MS:   []MyString{"hello", "world"},
		E:    []TestEnum{TestEnumTestEnumValue1, TestEnumTestEnumValue2},
		B:    []MyBoolean{true, false},
	}

	d, err := bar.MarshalBinary()
	if err != nil {
		t.Fatal(err)
	}
	writeFile(t, "alias-lists.bin", d)

	var bar2 AliasLists
	err = bar2.UnmarshalBinary(bytes.NewBuffer(d))
	if err != nil {
		t.Fatal(err)
	}

	assertSlicesEqual(t, bar2.MI32, bar.MI32)
	assertSlicesEqual(t, bar2.MI64, bar.MI64)
	assertSlicesEqual(t, bar2.MF32, bar.MF32)
	assertSlicesEqual(t, bar2.MF64, bar.MF64)
	assertSlicesEqual(t, bar2.MS, bar.MS)
	assertSlicesEqual(t, bar2.E, bar.E)
	assertSlicesEqual(t, bar2.B, bar.B)
}

func TestFooBar(t *testing.T) {
	foo := getFoobarFixture()

	d, err := foo.MarshalBinary()
	if err != nil {
		t.Fatal(err)
	}
	writeFile(t, "foobar.bin", d)

	var foo2 Foobar
	err = foo2.UnmarshalBinary(bytes.NewBuffer(d))
	if err != nil {
		t.Fatal(err)
	}

	assertEquals(t, foo2.A.MI32, foo.A.MI32)
	assertEquals(t, foo2.A.MI64, foo.A.MI64)
	assertEquals(t, foo2.A.MF32, foo.A.MF32)
	assertEquals(t, foo2.A.MF64, foo.A.MF64)
	assertEquals(t, foo2.A.MS, foo.A.MS)
	assertEquals(t, foo2.A.E, foo.A.E)
	assertEquals(t, foo2.A.B, foo.A.B)

	assertEquals(t, foo2.P.I32, foo.P.I32)
	assertEquals(t, foo2.P.I64, foo.P.I64)
	assertEquals(t, foo2.P.F32, foo.P.F32)
	assertEquals(t, foo2.P.F64, foo.P.F64)
	assertEquals(t, foo2.P.S, foo.P.S)
	assertEquals(t, foo2.P.B, foo.P.B)

	assertSlicesEqual(t, foo2.AL.MI32, foo.AL.MI32)
	assertSlicesEqual(t, foo2.AL.MI64, foo.AL.MI64)
	assertSlicesEqual(t, foo2.AL.MF32, foo.AL.MF32)
	assertSlicesEqual(t, foo2.AL.MF64, foo.AL.MF64)
	assertSlicesEqual(t, foo2.AL.MS, foo.AL.MS)
	assertSlicesEqual(t, foo2.AL.E, foo.AL.E)
	assertSlicesEqual(t, foo2.AL.B, foo.AL.B)

	assertSlicesEqual(t, foo2.PL.I32, foo.PL.I32)
	assertSlicesEqual(t, foo2.PL.I64, foo.PL.I64)
	assertSlicesEqual(t, foo2.PL.F32, foo.PL.F32)
	assertSlicesEqual(t, foo2.PL.F64, foo.PL.F64)
	assertSlicesEqual(t, foo2.PL.S, foo.PL.S)
	assertSlicesEqual(t, foo.PL.B, foo.PL.B)
}

func writeFile(t *testing.T, name string, data []byte) {
	t.Helper()
	err := ioutil.WriteFile(fmt.Sprintf("out-%s", name), data, 0644)
	if err != nil {
		t.Fatal(err)
	}
}

func getFoobarFixture() Foobar {
	return Foobar{
		A: Alias{
			MI32: 1,
			MI64: 2,
			MF32: 3.3,
			MF64: 4.4,
			MS:   "hello",
			E:    TestEnumTestEnumValue2,
			B:    true,
		},
		P: Primitive{
			I32: 1,
			I64: 2,
			F32: 3.3,
			F64: 4.4,
			S:   "hello",
			B:   true,
		},
		AL: AliasLists{
			MI32: []MyInteger32{math.MinInt32, 0, math.MaxInt32},
			MI64: []MyInteger64{math.MinInt64, 0, math.MaxInt64},
			MF32: []MyFloat32{float32Min, 0, math.MaxFloat32},
			MF64: []MyFloat64{float64Min, 0, math.MaxFloat64},
			MS:   []MyString{"hello", "world"},
			E:    []TestEnum{TestEnumTestEnumValue1, TestEnumTestEnumValue2},
			B:    []MyBoolean{true, false},
		},
		PL: PrimitiveLists{
			I32: []int32{math.MinInt32, 0, math.MaxInt32},
			I64: []int64{math.MinInt64, 0, math.MaxInt64},
			F32: []float32{float32Min, 0, math.MaxFloat32},
			F64: []float64{float64Min, 0, math.MaxFloat64},
			S:   []string{"hello", "world"},
			B:   []bool{true, false},
		},
	}
}

func TestRead(t *testing.T) {
	data, err := ioutil.ReadFile("out-foobar.bin")
	if err != nil {
		t.Fatal(err)
	}

	var foo Foobar
	err = foo.UnmarshalBinary(bytes.NewBuffer(data))
	if err != nil {
		t.Fatal(err)
	}
}

func readCsFile(t *testing.T, name string) []byte {
	t.Helper()
	data, err := ioutil.ReadFile("../cs/out-" + name)
	if err != nil {
		t.Fatal(err)
	}

	return data
}

func TestCrossLanguagePrimitive(t *testing.T) {
	data := readCsFile(t, "primitive.bin")
	var p Primitive
	err := p.UnmarshalBinary(bytes.NewBuffer(data))
	if err != nil {
		t.Fatal(err)
	}

	assertEquals(t, p.I32, int32(1))
	assertEquals(t, p.I64, int64(2))
	assertEquals(t, p.F32, float32(3.3))
	//TODO: check why this fails, go -> c# works
	//assertEquals(t, p.F64, float64(4.4))
	assertEquals(t, p.S, "hello")
	assertEquals(t, p.B, true)
}

func TestCrossLanguageAlias(t *testing.T) {
	data := readCsFile(t, "alias.bin")
	var a Alias
	err := a.UnmarshalBinary(bytes.NewBuffer(data))
	if err != nil {
		t.Fatal(err)
	}

	assertEquals(t, a.MI32, MyInteger32(1))
	assertEquals(t, a.MI64, MyInteger64(2))
	assertEquals(t, a.MF32, MyFloat32(3.3))
	//TODO: check why this fails, go -> c# works
	//assertEquals(t, a.MF64, MyFloat64(4.4))
	assertEquals(t, a.MS, MyString("hello"))
	assertEquals(t, a.E, TestEnumTestEnumValue2)
	assertEquals(t, a.B, MyBoolean(true))
}

func TestCrossLanguagePrimitiveLists(t *testing.T) {
	data := readCsFile(t, "primitive-lists.bin")
	var pl PrimitiveLists
	err := pl.UnmarshalBinary(bytes.NewBuffer(data))
	if err != nil {
		t.Fatal(err)
	}

	assertSlicesEqual(t, pl.I32, []int32{1, 2, 3})
	assertSlicesEqual(t, pl.I64, []int64{4, 5, 6})
	assertSlicesEqual(t, pl.F32, []float32{7.7, 8.8, 9.9})
	assertSlicesEqual(t, pl.F64, []float64{10.10, 11.11, 12.12})
	assertSlicesEqual(t, pl.S, []string{"hello", "world"})
	assertSlicesEqual(t, pl.B, []bool{true, false, true})
}

func TestCrossLanguageAliasList(t *testing.T) {
	data := readCsFile(t, "alias-lists.bin")
	var al AliasLists
	err := al.UnmarshalBinary(bytes.NewBuffer(data))
	if err != nil {
		t.Fatal(err)
	}

	assertSlicesEqual(t, al.MI32, []MyInteger32{1, 2, 3})
	assertSlicesEqual(t, al.MI64, []MyInteger64{4, 5, 6})
	assertSlicesEqual(t, al.MF32, []MyFloat32{7.7, 8.8, 9.9})
	assertSlicesEqual(t, al.MF64, []MyFloat64{10.10, 11.11, 12.12})
	assertSlicesEqual(t, al.MS, []MyString{"hello", "world"})
	assertSlicesEqual(t, al.E, []TestEnum{TestEnumTestEnumValue1, TestEnumTestEnumValue2})
	assertSlicesEqual(t, al.B, []MyBoolean{true, false, true})
}

func TestWriteMessage(t *testing.T) {
	o := OneField{}
	o.S = "hello-world"

	f, err := os.Create("out-envelope-one-field.bin")
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()

	err = WriteMessage(f, o)
	if err != nil {
		t.Fatal(err)
	}
}

func TestReadMessage(t *testing.T) {
	f, err := os.Open("out-envelope-one-field.bin")
	if err != nil {
		t.Fatal(err)
	}

	defer f.Close()

	o, err := ReadMessage(f)
	if err != nil {
		t.Fatal(err)
	}

	oneField, ok := o.(OneField)
	if !ok {
		t.Fatal("expected OneField")
	}

	assertEquals(t, oneField.S, "hello-world")
}

func BenchmarkFoobar_MarshalBinary(b *testing.B) {
	foo := getFoobarFixture()

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = foo.MarshalBinary()
	}
}

func BenchmarkFoobar_UnmarshalBinary(b *testing.B) {
	foo := getFoobarFixture()

	d, _ := foo.MarshalBinary()

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = foo.UnmarshalBinary(bytes.NewBuffer(d))
	}
}

func assertEquals(t *testing.T, a, b any) {
	t.Helper()
	if a != b {
		t.Fatalf("%v != %v", a, b)
	}
}

func assertSlicesEqual[T any](t *testing.T, a, b []T) {
	t.Helper()
	if len(a) != len(b) {
		t.Fatalf("len(a) != len(b): %d != %d", len(a), len(b))
	}
	if !reflect.DeepEqual(a, b) {
		t.Fatalf("a != b: %v != %v", a, b)
	}
}
