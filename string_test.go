package ptr

import (
	"testing"
	"time"
)

type scase struct {
	in  interface{}
	out string
}

var scases = []scase{
	{struct{}{}, "{}"},
	{true, "true"},
	{"happy", "happy"},
	{int(-1), "-1"},
	{int8(-2), "-2"},
	{int16(-3), "-3"},
	{int32(-4), "-4"},
	{int64(-5), "-5"},
	{uint(1), "1"},
	{uint8(2), "2"},
	{uint16(3), "3"},
	{uint32(4), "4"},
	{uint64(5), "5"},
	{uintptr(6), "6"},
	{byte(10), "10"},
	{rune(11), "11"},
	{float32(4.5), "4.5"},
	{float64(9000.1), "9000.1"},
	{complex64(12i), "(0+12i)"},
	{complex128(13i), "(0+13i)"},
	{time.Time{}, time.Time{}.String()},
}

func TestS(t *testing.T) {
	for _, c := range scases {
		var res string
		switch v := c.in.(type) {
		case bool:
			res = S(&v)
		case string:
			res = S(&v)
		case int:
			res = S(&v)
		case int8:
			res = S(&v)
		case int16:
			res = S(&v)
		case int32:
			res = S(&v)
		case int64:
			res = S(&v)
		case uint:
			res = S(&v)
		case uint8:
			res = S(&v)
		case uint16:
			res = S(&v)
		case uint32:
			res = S(&v)
		case uint64:
			res = S(&v)
		case uintptr:
			res = S(&v)
		case float32:
			res = S(&v)
		case float64:
			res = S(&v)
		case complex64:
			res = S(&v)
		case complex128:
			res = S(&v)
		case time.Time:
			res = S(&v)
		default:
			res = S(&v)
		}

		if res != c.out {
			t.Errorf("var item %T = %v; S(&item) = \"%v\", want \"%v\"", c.in, c.in, res, c.out)
		}
	}
}

func TestS_Nil(t *testing.T) {
	if S(nil) != "<nil>" {
		t.Errorf("S(nil) = \"%v\", want \"<nil>\"", S(nil))
	}
	
	var s *string
	if S(s) != "<nil>" {
		t.Errorf("S(<nil string>) = \"%v\", want \"<nil>\"", S(s))
	}
}
