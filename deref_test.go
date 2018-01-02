package ptr

import (
	"testing"
	"time"
)

func TestDBoolValue(t *testing.T) {
	x := Bool(true)
	def := false
	res := DBool(x, def)
	if res != *x {
		t.Fatalf("expected %t, got %t", *x, res)
	}
}

func TestDBoolDefault(t *testing.T) {
	def := true
	res := DBool(nil, def)
	if res != def {
		t.Fatalf("expected %t, got %t", def, res)
	}
}

func TestDByteValue(t *testing.T) {
	x := Byte(byte(10))
	def := byte(20)
	res := DByte(x, def)
	if res != *x {
		t.Fatalf("expected %t, got %t", *x, res)
	}
}

func TestDByteDefault(t *testing.T) {
	def := byte(20)
	res := DByte(nil, def)
	if res != def {
		t.Fatalf("expected %t, got %t", def, res)
	}
}

func TestDComplex64Value(t *testing.T) {
	x := Complex64(complex64(10))
	def := complex64(20)
	res := DComplex64(x, def)
	if res != *x {
		t.Fatalf("expected %t, got %t", *x, res)
	}
}

func TestDComplex64Default(t *testing.T) {
	def := complex64(20)
	res := DComplex64(nil, def)
	if res != def {
		t.Fatalf("expected %t, got %t", def, res)
	}
}

func TestDComplex128Value(t *testing.T) {
	x := Complex128(complex128(10))
	def := complex128(20)
	res := DComplex128(x, def)
	if res != *x {
		t.Fatalf("expected %t, got %t", *x, res)
	}
}

func TestDComplex128Default(t *testing.T) {
	def := complex128(20)
	res := DComplex128(nil, def)
	if res != def {
		t.Fatalf("expected %t, got %t", def, res)
	}
}

func TestDDurationValue(t *testing.T) {
	x := Duration(time.Duration(5))
	def := time.Duration(10)
	res := DDuration(x, def)
	if res != *x {
		t.Fatalf("expected %t, got %t", *x, res)
	}
}

func TestDDurationDefault(t *testing.T) {
	def := time.Duration(10)
	res := DDuration(nil, def)
	if res != def {
		t.Fatalf("expected %t, got %t", def, res)
	}
}

func TestDFloat32Value(t *testing.T) {
	x := Float32(float32(10))
	def := float32(20)
	res := DFloat32(x, def)
	if res != *x {
		t.Fatalf("expected %t, got %t", *x, res)
	}
}

func TestDFloat32Default(t *testing.T) {
	def := float32(20)
	res := DFloat32(nil, def)
	if res != def {
		t.Fatalf("expected %t, got %t", def, res)
	}
}

func TestDFloat64Value(t *testing.T) {
	x := Float64(float64(10))
	def := float64(20)
	res := DFloat64(x, def)
	if res != *x {
		t.Fatalf("expected %t, got %t", *x, res)
	}
}

func TestDFloat64Default(t *testing.T) {
	def := float64(20)
	res := DFloat64(nil, def)
	if res != def {
		t.Fatalf("expected %t, got %t", def, res)
	}
}

func TestDIntValue(t *testing.T) {
	x := Int(int(10))
	def := int(20)
	res := DInt(x, def)
	if res != *x {
		t.Fatalf("expected %t, got %t", *x, res)
	}
}

func TestDIntDefault(t *testing.T) {
	def := int(20)
	res := DInt(nil, def)
	if res != def {
		t.Fatalf("expected %t, got %t", def, res)
	}
}

func TestDInt8Value(t *testing.T) {
	x := Int8(int8(10))
	def := int8(20)
	res := DInt8(x, def)
	if res != *x {
		t.Fatalf("expected %t, got %t", *x, res)
	}
}

func TestDInt8Default(t *testing.T) {
	def := int8(20)
	res := DInt8(nil, def)
	if res != def {
		t.Fatalf("expected %t, got %t", def, res)
	}
}

func TestDInt16Value(t *testing.T) {
	x := Int16(10)
	def := int16(20)
	res := DInt16(x, def)
	if res != *x {
		t.Fatalf("expected %t, got %t", *x, res)
	}
}

func TestDInt16Default(t *testing.T) {
	def := int16(20)
	res := DInt16(nil, def)
	if res != def {
		t.Fatalf("expected %t, got %t", def, res)
	}
}

func TestDInt32Value(t *testing.T) {
	x := Int32(int32(10))
	def := int32(20)
	res := DInt32(x, def)
	if res != *x {
		t.Fatalf("expected %t, got %t", *x, res)
	}
}

func TestDInt32Default(t *testing.T) {
	def := int32(20)
	res := DInt32(nil, def)
	if res != def {
		t.Fatalf("expected %t, got %t", def, res)
	}
}

func TestDInt64Value(t *testing.T) {
	x := Int64(int64(10))
	def := int64(20)
	res := DInt64(x, def)
	if res != *x {
		t.Fatalf("expected %t, got %t", *x, res)
	}
}

func TestDInt64Default(t *testing.T) {
	def := int64(20)
	res := DInt64(nil, def)
	if res != def {
		t.Fatalf("expected %t, got %t", def, res)
	}
}

func TestDRuneValue(t *testing.T) {
	x := Rune(rune(10))
	def := rune(20)
	res := DRune(x, def)
	if res != *x {
		t.Fatalf("expected %t, got %t", *x, res)
	}
}

func TestDRuneDefault(t *testing.T) {
	def := rune(20)
	res := DRune(nil, def)
	if res != def {
		t.Fatalf("expected %t, got %t", def, res)
	}
}

func TestDStringValue(t *testing.T) {
	x := String("one")
	def := "two"
	res := DString(x, def)
	if res != *x {
		t.Fatalf("expected %t, got %t", *x, res)
	}
}

func TestDStringDefault(t *testing.T) {
	def := "two"
	res := DString(nil, def)
	if res != def {
		t.Fatalf("expected %t, got %t", def, res)
	}
}

func TestDTimeValue(t *testing.T) {
	x := Time(time.Now())
	def := time.Now().Add(10 * time.Minute)
	res := DTime(x, def)
	if res != *x {
		t.Fatalf("expected %t, got %t", *x, res)
	}
}

func TestDTimeDefault(t *testing.T) {
	def := time.Now().Add(10 * time.Minute)
	res := DTime(nil, def)
	if res != def {
		t.Fatalf("expected %t, got %t", def, res)
	}
}

func TestDUintValue(t *testing.T) {
	x := Uint(uint(10))
	def := uint(20)
	res := DUint(x, def)
	if res != *x {
		t.Fatalf("expected %t, got %t", *x, res)
	}
}

func TestDUintDefault(t *testing.T) {
	def := uint(20)
	res := DUint(nil, def)
	if res != def {
		t.Fatalf("expected %t, got %t", def, res)
	}
}

func TestDUint8Value(t *testing.T) {
	x := Uint8(uint8(10))
	def := uint8(20)
	res := DUint8(x, def)
	if res != *x {
		t.Fatalf("expected %t, got %t", *x, res)
	}
}

func TestDUint8Default(t *testing.T) {
	def := uint8(20)
	res := DUint8(nil, def)
	if res != def {
		t.Fatalf("expected %t, got %t", def, res)
	}
}

func TestDUint16Value(t *testing.T) {
	x := Uint16(10)
	def := uint16(20)
	res := DUint16(x, def)
	if res != *x {
		t.Fatalf("expected %t, got %t", *x, res)
	}
}

func TestDUint16Default(t *testing.T) {
	def := uint16(20)
	res := DUint16(nil, def)
	if res != def {
		t.Fatalf("expected %t, got %t", def, res)
	}
}

func TestDUint32Value(t *testing.T) {
	x := Uint32(uint32(10))
	def := uint32(20)
	res := DUint32(x, def)
	if res != *x {
		t.Fatalf("expected %t, got %t", *x, res)
	}
}

func TestDUint32Default(t *testing.T) {
	def := uint32(20)
	res := DUint32(nil, def)
	if res != def {
		t.Fatalf("expected %t, got %t", def, res)
	}
}

func TestDUint64Value(t *testing.T) {
	x := Uint64(uint64(10))
	def := uint64(20)
	res := DUint64(x, def)
	if res != *x {
		t.Fatalf("expected %t, got %t", *x, res)
	}
}

func TestDUint64Default(t *testing.T) {
	def := uint64(20)
	res := DUint64(nil, def)
	if res != def {
		t.Fatalf("expected %t, got %t", def, res)
	}
}

func TestDUintptrValue(t *testing.T) {
	x := Uintptr(uintptr(10))
	def := uintptr(20)
	res := DUintptr(x, def)
	if res != *x {
		t.Fatalf("expected %t, got %t", *x, res)
	}
}

func TestDUintptrDefault(t *testing.T) {
	def := uintptr(20)
	res := DUintptr(nil, def)
	if res != def {
		t.Fatalf("expected %t, got %t", def, res)
	}
}
