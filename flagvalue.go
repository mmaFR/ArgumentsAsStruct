package Parser

import (
	"errors"
	"strconv"
)

func numError(err error) error {
	ne, ok := err.(*strconv.NumError)
	if !ok {
		return err
	}
	if ne.Err == strconv.ErrSyntax {
		return errors.New("parse error")
	}
	if ne.Err == strconv.ErrRange {
		return errors.New("value out of range")
	}
	return err
}

type intXBitsValue struct {
	p8  *int8
	p16 *int16
	p32 *int32
	pN  uint8
}

func (i *intXBitsValue) Set(s string) error {
	var v int64
	var err error
	v, err = strconv.ParseInt(s, 0, int(i.pN))
	if err != nil {
		err = numError(err)
	}
	switch i.pN {
	case 8:
		*i.p8 = int8(v)
	case 16:
		*i.p16 = int16(v)
	case 32:
		*i.p32 = int32(v)
	}
	return err
}
func (i *intXBitsValue) Get() interface{} {
	switch i.pN {
	case 8:
		return int8(*i.p8)
	case 16:
		return int16(*i.p16)
	case 32:
		return int32(*i.p32)
	default:
		return nil
	}
}
func (i *intXBitsValue) String() string {
	var val int64
	switch i.pN {
	case 8:
		val = int64(*i.p8)
	case 16:
		val = int64(*i.p16)
	case 32:
		val = int64(*i.p32)
	}
	return strconv.FormatInt(val, 10)
}
func newIntXBitsValue(val interface{}, p interface{}, bitsSize uint8) *intXBitsValue {
	var value *intXBitsValue = new(intXBitsValue)
	value.pN = bitsSize
	switch bitsSize {
	case 8:
		*p.(*int8) = val.(int8)
		value.p8 = p.(*int8)
	case 16:
		*p.(*int16) = val.(int16)
		value.p16 = p.(*int16)
	case 32:
		*p.(*int32) = val.(int32)
		value.p32 = p.(*int32)
	default:
		return nil
	}
	return value
}

type uintXBitsValue struct {
	p8  *uint8
	p16 *uint16
	p32 *uint32
	pN  uint8
}

func (ui *uintXBitsValue) Set(s string) error {
	var v uint64
	var err error
	v, err = strconv.ParseUint(s, 0, int(ui.pN))
	if err != nil {
		err = numError(err)
	}
	switch ui.pN {
	case 8:
		*ui.p8 = uint8(v)
	case 16:
		*ui.p16 = uint16(v)
	case 32:
		*ui.p32 = uint32(v)
	}
	return err
}
func (ui *uintXBitsValue) Get() interface{} {
	switch ui.pN {
	case 8:
		return uint8(*ui.p8)
	case 16:
		return uint16(*ui.p16)
	case 32:
		return uint32(*ui.p32)
	default:
		return nil
	}
}
func (ui *uintXBitsValue) String() string {
	var val uint64
	switch ui.pN {
	case 8:
		val = uint64(*ui.p8)
	case 16:
		val = uint64(*ui.p16)
	case 32:
		val = uint64(*ui.p32)
	}
	return strconv.FormatUint(val, 10)
}
func newUintXBitsValue(val interface{}, p interface{}, bitsSize uint8) *uintXBitsValue {
	var value *uintXBitsValue = new(uintXBitsValue)
	value.pN = bitsSize
	switch bitsSize {
	case 8:
		*p.(*uint8) = val.(uint8)
		value.p8 = p.(*uint8)
	case 16:
		*p.(*uint16) = val.(uint16)
		value.p16 = p.(*uint16)
	case 32:
		*p.(*uint32) = val.(uint32)
		value.p32 = p.(*uint32)
	default:
		return nil
	}
	return value
}

type float32Value float32

func (i *float32Value) Set(s string) error {
	v, err := strconv.ParseFloat(s, 32)
	if err != nil {
		err = numError(err)
	}
	*i = float32Value(v)
	return err
}
func (i *float32Value) Get() interface{} { return float32(*i) }
func (i *float32Value) String() string   { return strconv.FormatFloat(float64(*i), 'g', -1, 32) }
func newFloat32Value(val float32, p *float32) *float32Value {
	*p = val
	return (*float32Value)(p)
}
