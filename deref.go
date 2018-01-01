package ptr

import "time"

// DBool returns the dereferenced pointer to the input *bool, or the provided default if nil.
func DBool(b *bool, def bool) bool {
	if b == nil {
		return def
	}
	return *b
}

// DByte returns the dereferenced pointer to the input *byte, or the provided default if nil.
func DByte(b *byte, def byte) byte {
	if b == nil {
		return def
	}
	return *b
}

// DComplex64 returns the dereferenced pointer to the input *complex64, or the provided default if nil.
func DComplex64(c *complex64, def complex64) complex64 {
	if c == nil {
		return def
	}
	return *c
}

// DComplex128 returns the dereferenced pointer to the input *complex128, or the provided default if nil.
func DComplex128(c *complex128, def complex128) complex128 {
	if c == nil {
		return def
	}
	return *c
}

// DDuration returns the dereferenced pointer to the input *time.Duration, or the provided default if nil.
func DDuration(t *time.Duration, def time.Duration) time.Duration {
	if t == nil {
		return def
	}
	return *t
}

// DFloat32 returns the dereferenced pointer to the input *float32, or the provided default if nil.
func DFloat32(f *float32, def float32) float32 {
	if f == nil {
		return def
	}
	return *f
}

// DFloat64 returns the dereferenced pointer to the input *float64, or the provided default if nil.
func DFloat64(f *float64, def float64) float64 {
	if f == nil {
		return def
	}
	return *f
}

// DInt returns the dereferenced pointer to the input *int, or the provided default if nil.
func DInt(i *int, def int) int {
	if i == nil {
		return def
	}
	return *i
}

// DInt8 returns the dereferenced pointer to the input *int8, or the provided default if nil.
func DInt8(i *int8, def int8) int8 {
	if i == nil {
		return def
	}
	return *i
}

// DInt16 returns the dereferenced pointer to the input *int16, or the provided default if nil.
func DInt16(i *int16, def int16) int16 {
	if i == nil {
		return def
	}
	return *i
}

// DInt32 returns the dereferenced pointer to the input *int32, or the provided default if nil.
func DInt32(i *int32, def int32) int32 {
	if i == nil {
		return def
	}
	return *i
}

// DInt64 returns the dereferenced pointer to the input *int64
func DInt64(i *int64, def int64) int64 {
	if i == nil {
		return def
	}
	return *i
}

// DRune returns the dereferenced pointer to the input *rune, or the provided default if nil.
func DRune(r *rune, def rune) rune {
	if r == nil {
		return def
	}
	return *r
}

// DString returns the dereferenced pointer to the input *string, or the provided default if nil.
func DString(s *string, def string) string {
	if s == nil {
		return def
	}
	return *s
}

// DTime returns the dereferenced pointer to the input *time.Time, or the provided default if nil.
func DTime(t *time.Time, def time.Time) time.Time {
	if t == nil {
		return def
	}
	return *t
}

// DUint returns the dereferenced pointer to the input *uint, or the provided default if nil.
func DUint(u *uint, def uint) uint {
	if u == nil {
		return def
	}
	return *u
}

// DUint8 returns the dereferenced pointer to the input *uint8, or the provided default if nil.
func DUint8(u *uint8, def uint8) uint8 {
	if u == nil {
		return def
	}
	return *u
}

// DUint16 returns the dereferenced pointer to the input *uint16, or the provided default if nil.
func DUint16(u *uint16, def uint16) uint16 {
	if u == nil {
		return def
	}
	return *u
}

// DUint32 returns the dereferenced pointer to the input *uint32, or the provided default if nil.
func DUint32(u *uint32, def uint32) uint32 {
	if u == nil {
		return def
	}
	return *u
}

// DUint64 returns the dereferenced pointer to the input *uint64, or the provided default if nil.
func DUint64(u *uint64, def uint64) uint64 {
	if u == nil {
		return def
	}
	return *u
}

// DUintptr returns the dereferenced pointer to the input *uintptr, or the provided default if nil.
func DUintptr(u *uintptr, def uintptr) uintptr {
	if u == nil {
		return def
	}
	return *u
}
