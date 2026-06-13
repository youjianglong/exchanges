package common

import (
	"errors"
	"fmt"
	"math"
	"strconv"
	"strings"
)

func unquote(data []byte) string {
	str := string(data)
	if str == "null" || str == "\"\"" {
		return ""
	}
	if data[0] == '"' && data[len(data)-1] == '"' {
		data = data[1 : len(data)-1]
	}
	str = string(data)
	return str
}

func quote(data string) string {
	b := make([]byte, 0, len(data)+2)
	b = append(b, '"')
	b = append(b, data...)
	b = append(b, '"')
	return string(b)
}

var ErrEmptyString = errors.New("empty string")

type Float64 string

func NewFloat64(f float64, decimals ...int) Float64 {
	dec := 9
	if len(decimals) > 0 {
		dec = decimals[0]
	}
	return Float64(fmt.Sprintf("%."+strconv.Itoa(dec)+"f", f))
}

func (f Float64) String() string {
	return string(f)
}

func (f Float64) Float64() (float64, error) {
	if f == "" {
		return 0, ErrEmptyString
	}
	return strconv.ParseFloat(string(f), 64)
}

func (f Float64) Value() float64 {
	if f == "" {
		return 0
	}
	v, _ := f.Float64()
	return v
}

func (f Float64) Round(precision int) Float64 {
	return NewFloat64(f.Value(), precision)
}

func (f Float64) Floor(precision int) Float64 {
	v := f.Value()
	v = math.Floor(v*math.Pow10(precision)) / math.Pow10(precision) // 向下取整
	return NewFloat64(v, precision)
}

func (f Float64) Ceil(precision int) Float64 {
	v := f.Value()
	if precision == 0 {
		return NewFloat64(math.Ceil(v), 0)
	}
	v = math.Ceil(v*math.Pow10(precision)) / math.Pow10(precision) // 向上取整
	return NewFloat64(v, precision)
}

func (f Float64) Add(other Float64) Float64 {
	return NewFloat64(f.Value() + other.Value())
}

func (f Float64) Sub(other Float64) Float64 {
	return NewFloat64(f.Value() - other.Value())
}

func (f Float64) IsZero() bool {
	return f == "" || f == "0" || f == "0.0" || strings.TrimRight(string(f), "0") == "0."
}

func (f *Float64) UnmarshalJSON(data []byte) error {
	if len(data) == 0 {
		return nil
	}
	str := unquote(data)
	if str == "" {
		return nil
	}
	_, err := strconv.ParseFloat(str, 64)
	if err != nil {
		return fmt.Errorf("invalid float64 value: %s, error: %w, data: %s", str, err, string(data))
	}
	*f = Float64(str)
	return nil
}

func (f Float64) MarshalJSON() ([]byte, error) {
	return []byte(quote(string(f))), nil
}

type Int64 string

func NewInt64(i int64) Int64 {
	return Int64(strconv.FormatInt(i, 10))
}

func (i Int64) String() string {
	return string(i)
}

func (i Int64) Int64() (int64, error) {
	return strconv.ParseInt(string(i), 10, 64)
}

func (i Int64) Value() int64 {
	if i == "" {
		return 0
	}
	v, _ := i.Int64()
	return v
}

func (i Int64) IsZero() bool {
	return i == "0" || i == ""
}

func (i *Int64) UnmarshalJSON(data []byte) error {
	if len(data) == 0 {
		return nil
	}
	str := unquote(data)
	if str == "" {
		return nil
	}
	_, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		return fmt.Errorf("invalid int64 value: %s, error: %w, data: %s", str, err, string(data))
	}
	*i = Int64(str)
	return nil
}

func (i Int64) MarshalJSON() ([]byte, error) {
	return []byte(quote(string(i))), nil
}

type Mixed string

func (i *Mixed) UnmarshalJSON(data []byte) error {
	*i = Mixed(unquote(data))
	return nil
}

func (i Mixed) MarshalJSON() ([]byte, error) {
	return []byte(quote(string(i))), nil
}

func (i Mixed) String() string {
	return string(i)
}

func (i Mixed) Value() string {
	return string(i)
}

func (i Mixed) Float64() Float64 {
	return Float64(i)
}

func (i Mixed) Int64() Int64 {
	return Int64(i)
}

func (i Mixed) IsZero() bool {
	return i == ""
}

// Abs 取绝对值
func Abs[T Float64 | Int64 | Mixed](val T) T {
	if len(val) > 0 && val[0] == '-' {
		return val[1:]
	}
	return val
}
