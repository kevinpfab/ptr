package ptr

import (
	"testing"
	"time"
)

func TestBool(t *testing.T) {
	for _, b := range []bool{false, true} {
		p := Bool(b)
		if p == nil {
			t.Fatalf("expected %t, got <nil>", b)
		}
		if b != *p {
			t.Fatalf("expected %t, got %t", b, *p)
		}
	}
}

func TestString(t *testing.T) {
	for _, s := range []string{"", "one", "fish", "two", "fish"} {
		p := String(s)
		if p == nil {
			t.Fatalf("expected %q, got <nil>", s)
		}
		if s != *p {
			t.Fatalf("expected %q, got %q", s, *p)
		}
	}
}

func TestInt(t *testing.T) {
	for _, i := range []int{-1, 0, 1, 1, 2, 3, 5} {
		p := Int(i)
		if p == nil {
			t.Fatalf("expected %v, got <nil>", i)
		}
		if i != *p {
			t.Fatalf("expected %v, got %v", i, *p)
		}
	}
}

func TestInt8(t *testing.T) {
	for _, i := range []int8{-1, 0, 1, 1<<7 - 1, -1 << 7} {
		p := Int8(i)
		if p == nil {
			t.Fatalf("expected %v, got <nil>", i)
		}
		if i != *p {
			t.Fatalf("expected %v, got %v", i, *p)
		}
	}
}

func TestInt16(t *testing.T) {
	for _, i := range []int16{-1, 0, 1, 1<<15 - 1, -1 << 15} {
		p := Int16(i)
		if p == nil {
			t.Fatalf("expected %v, got <nil>", i)
		}
		if i != *p {
			t.Fatalf("expected %v, got %v", i, *p)
		}
	}
}

func TestInt32(t *testing.T) {
	for _, i := range []int32{-1, 0, 1, 1<<31 - 1, -1 << 31} {
		p := Int32(i)
		if p == nil {
			t.Fatalf("expected %v, got <nil>", i)
		}
		if i != *p {
			t.Fatalf("expected %v, got %v", i, *p)
		}
	}
}

func TestInt64(t *testing.T) {
	for _, i := range []int64{-1, 0, 1, 1<<63 - 1, -1 << 63} {
		p := Int64(i)
		if p == nil {
			t.Fatalf("expected %v, got <nil>", i)
		}
		if i != *p {
			t.Fatalf("expected %v, got %v", i, *p)
		}
	}
}

func TestUint(t *testing.T) {
	for _, i := range []uint{0, 1, 100, 1000} {
		p := Uint(i)
		if p == nil {
			t.Fatalf("expected %v, got <nil>", i)
		}
		if i != *p {
			t.Fatalf("expected %v, got %v", i, *p)
		}
	}
}

func TestUint8(t *testing.T) {
	for _, i := range []uint8{0, 1, 1<<8 - 1} {
		p := Uint8(i)
		if p == nil {
			t.Fatalf("expected %v, got <nil>", i)
		}
		if i != *p {
			t.Fatalf("expected %v, got %v", i, *p)
		}
	}
}

func TestUint16(t *testing.T) {
	for _, i := range []uint16{0, 1, 1<<16 - 1} {
		p := Uint16(i)
		if p == nil {
			t.Fatalf("expected %v, got <nil>", i)
		}
		if i != *p {
			t.Fatalf("expected %v, got %v", i, *p)
		}
	}
}

func TestUint32(t *testing.T) {
	for _, i := range []uint32{0, 1, 1<<32 - 1} {
		p := Uint32(i)
		if p == nil {
			t.Fatalf("expected %v, got <nil>", i)
		}
		if i != *p {
			t.Fatalf("expected %v, got %v", i, *p)
		}
	}
}

func TestUint64(t *testing.T) {
	for _, i := range []uint64{0, 1, 1<<64 - 1} {
		p := Uint64(i)
		if p == nil {
			t.Fatalf("expected %v, got <nil>", i)
		}
		if i != *p {
			t.Fatalf("expected %v, got %v", i, *p)
		}
	}
}

func TestUintptr(t *testing.T) {
	for _, i := range []uintptr{0, 1, 10, 100, 1000} {
		p := Uintptr(i)
		if p == nil {
			t.Fatalf("expected %v, got <nil>", i)
		}
		if i != *p {
			t.Fatalf("expected %v, got %v", i, *p)
		}
	}
}

func TestByte(t *testing.T) {
	for _, i := range []byte{0, 1, 1<<8 - 1} {
		p := Byte(i)
		if p == nil {
			t.Fatalf("expected %v, got <nil>", i)
		}
		if i != *p {
			t.Fatalf("expected %v, got %v", i, *p)
		}
	}
}

func TestRune(t *testing.T) {
	for _, i := range []rune{0, 1, 1<<16 - 1} {
		p := Rune(i)
		if p == nil {
			t.Fatalf("expected %v, got <nil>", i)
		}
		if i != *p {
			t.Fatalf("expected %v, got %v", i, *p)
		}
	}
}

func TestFloat32(t *testing.T) {
	for _, f := range []float32{-1.11, 0, 3.14159} {
		p := Float32(f)
		if p == nil {
			t.Fatalf("expected %v, got <nil>", f)
		}
		if f != *p {
			t.Fatalf("expected %v, got %v", f, *p)
		}
	}
}

func TestFloat64(t *testing.T) {
	for _, f := range []float64{-1.11, 0, 3.14159} {
		p := Float64(f)
		if p == nil {
			t.Fatalf("expected %v, got <nil>", f)
		}
		if f != *p {
			t.Fatalf("expected %v, got %v", f, *p)
		}
	}
}

func TestComplex64(t *testing.T) {
	for _, c := range []complex64{0, 1, 12i} {
		p := Complex64(c)
		if p == nil {
			t.Fatalf("expected %v, got <nil>", c)
		}
		if c != *p {
			t.Fatalf("expected %v, got %v", c, *p)
		}
	}
}

func TestComplex128(t *testing.T) {
	for _, c := range []complex128{0, 1, 500i} {
		p := Complex128(c)
		if p == nil {
			t.Fatalf("expected %v, got <nil>", c)
		}
		if c != *p {
			t.Fatalf("expected %v, got %v", c, *p)
		}
	}
}

func TestTime(t *testing.T) {
	for _, tm := range []time.Time{
		time.Now(),
		time.Now().Add(15 * time.Minute),
		time.Now().Add(-40 * time.Hour),
	} {
		p := Time(tm)
		if p == nil {
			t.Fatalf("expected %v, got <nil>", tm)
		}
		if tm != *p {
			t.Fatalf("expected %v, got %v", tm, *p)
		}
	}
}

func TestDuration(t *testing.T) {
	for _, tm := range []time.Duration{
		5 * time.Millisecond,
		10 * time.Second,
		15 * time.Minute,
	} {
		p := Duration(tm)
		if p == nil {
			t.Fatalf("expected %v, got <nil>", tm)
		}
		if tm != *p {
			t.Fatalf("expected %v, got %v", tm, *p)
		}
	}
}
