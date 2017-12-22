// Package ptr provides small helper functions for working with pointers of basic types inline.
//
// For example, you can write...
//     phone := struct{Price *float32}{
//         Price: ptr.Float32(5.00),
//     }
// instead of having to set a basic literal to a variable...
//     price := 5.00
//     phone := struct{Price *float32}{
//         Price: &price,
//     }
//
// It also exposes a function to safely get the string representation of the underlying value of a pointer.
//     var name *string
//     fmt.Println(ptr.S(name))
//     // Output: <nil>
//
//     michael := "Michael Scott"
//     name = &michael
//     fmt.Println(ptr.S(name))
//     // Output: Michael Scott
package ptr

import "time"

// Bool returns a pointer to the input bool.
func Bool(b bool) *bool {
	return &b
}

// String returns a pointer to the input string.
func String(s string) *string {
	return &s
}

// Int returns a pointer to the input int.
func Int(i int) *int {
	return &i
}

// Int8 returns a pointer to the input int8.
func Int8(i int8) *int8 {
	return &i
}

// Int16 returns a pointer to the input int16.
func Int16(i int16) *int16 {
	return &i
}

// Int32 returns a pointer to the input int32.
func Int32(i int32) *int32 {
	return &i
}

// Int64 returns a pointer to the input int64
func Int64(i int64) *int64 {
	return &i
}

// Uint returns a pointer to the input uint.
func Uint(u uint) *uint {
	return &u
}

// Uint8 returns a pointer to the input uint8.
func Uint8(u uint8) *uint8 {
	return &u
}

// Uint16 returns a pointer to the input uint16.
func Uint16(u uint16) *uint16 {
	return &u
}

// Uint32 returns a pointer to the input uint32.
func Uint32(u uint32) *uint32 {
	return &u
}

// Uint64 returns a pointer to the input uint64.
func Uint64(u uint64) *uint64 {
	return &u
}

// Uintptr returns a pointer to the input uintptr.
func Uintptr(u uintptr) *uintptr {
	return &u
}

// Byte returns a pointer to the input byte.
func Byte(b byte) *byte {
	return &b
}

// Rune returns a pointer to the input rune.
func Rune(r rune) *rune {
	return &r
}

// Float32 returns a pointer to the input float32.
func Float32(f float32) *float32 {
	return &f
}

// Float64 returns a pointer to the input float64.
func Float64(f float64) *float64 {
	return &f
}

// Complex64 returns a pointer to the input complex64.
func Complex64(c complex64) *complex64 {
	return &c
}

// Complex128 returns a pointer to the input complex128.
func Complex128(c complex128) *complex128 {
	return &c
}

// Time returns a pointer to the input time.Time.
func Time(t time.Time) *time.Time {
	return &t
}
