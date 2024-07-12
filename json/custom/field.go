package main

import (
	"bytes"
	"fmt"
	"math"
	"strconv"
)

type field struct {
	i32 int32
	i64 int64
	f64 float64
	str string
}

func (f *field) UnmarshalJSON(dat []byte) error {
	datLen := len(dat)
	switch {
	case datLen >= 2 && dat[0] == '"' && dat[datLen-1] == '"':
		s, ok := unquote(dat)
		if !ok {
			return fmt.Errorf("无法解析: %s", dat)
		}
		f.str = s
		return nil
	case bytes.ContainsRune(dat, '.'):
		v, exx := strconv.ParseFloat(string(dat), 10)
		if exx != nil {
			return exx
		}
		f.f64 = v
		return nil
	case datLen > 0:
		v, exx := strconv.ParseInt(string(dat), 10, 64)
		if exx != nil {
			return exx
		}
		switch {
		case v >= math.MinInt32 && v <= math.MaxInt32:
			f.i32 = int32(v)
		default:
			f.i64 = v
		}
		return nil
	default:
		return fmt.Errorf("无法解析: %s", dat)
	}
}
