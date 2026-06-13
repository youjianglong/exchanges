package common

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

var (
	typeFloat64    = reflect.TypeFor[Float64]()
	typeInt64      = reflect.TypeFor[Int64]()
	typeMixed      = reflect.TypeFor[Mixed]()
	typeRawMessage = reflect.TypeFor[json.RawMessage]()
)

// StrictDecode 严格大小写敏感的 JSON 反序列化统一入口。
// struct 走 StrictUnmarshal，slice 逐元素 StrictUnmarshal，其余 fallback json.Unmarshal。
func StrictDecode(data []byte, target any) error {
	v := reflect.ValueOf(target)
	if v.Kind() != reflect.Ptr || v.IsNil() {
		return fmt.Errorf("target must be a non-nil pointer")
	}
	switch v.Elem().Kind() {
	case reflect.Struct:
		return StrictUnmarshal(data, target)
	case reflect.Slice:
		elemType := v.Elem().Type().Elem()
		if elemType.Kind() == reflect.Pointer {
			elemType = elemType.Elem()
		}
		if elemType.Kind() != reflect.Struct {
			return json.Unmarshal(data, target)
		}
		return strictUnmarshalSlice(data, v)
	default:
		return json.Unmarshal(data, target)
	}
}

// StrictUnmarshal 严格大小写敏感的 JSON 反序列化
// 自动从结构体的 json tag 中提取字段映射，无需手动写字段映射
func StrictUnmarshal(data []byte, target any) error {
	var raw map[string]json.RawMessage
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	v := reflect.ValueOf(target)
	if v.Kind() != reflect.Ptr || v.IsNil() {
		return fmt.Errorf("target must be a non-nil pointer")
	}
	v = v.Elem()
	t := v.Type()

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		jsonTag := field.Tag.Get("json")
		jsonField := parseJSONTag(jsonTag)
		if jsonField == "" || jsonField == "-" {
			continue
		}

		if rawValue, ok := raw[jsonField]; ok {
			structField := v.Field(i)
			if structField.CanSet() {
				if err := setFieldValue(structField, rawValue, jsonTag); err != nil {
					return fmt.Errorf("set field %s: %v", field.Name, err)
				}
			}
		}
	}
	return nil
}

func strictUnmarshalSlice(data []byte, v reflect.Value) error {
	var items []json.RawMessage
	if err := json.Unmarshal(data, &items); err != nil {
		return err
	}

	sliceVal := v.Elem()
	sliceType := sliceVal.Type()
	elemType := sliceType.Elem()
	isPtr := elemType.Kind() == reflect.Ptr
	if isPtr {
		elemType = elemType.Elem()
	}

	result := reflect.MakeSlice(sliceType, 0, len(items))
	for _, item := range items {
		var elem reflect.Value
		if isPtr {
			elem = reflect.New(elemType)
			if err := StrictUnmarshal(item, elem.Interface()); err != nil {
				return err
			}
		} else {
			elem = reflect.New(elemType).Elem()
			if err := StrictUnmarshal(item, elem.Addr().Interface()); err != nil {
				return err
			}
		}
		result = reflect.Append(result, elem)
	}
	sliceVal.Set(result)
	return nil
}

func parseJSONTag(tag string) string {
	if tag == "" {
		return ""
	}
	if idx := strings.Index(tag, ","); idx != -1 {
		return tag[:idx]
	}
	return tag
}

func hasJSONTagOptions(tag string) bool {
	for _, opt := range jsonTagOptions(tag) {
		if opt != "" && opt != "omitempty" {
			return true
		}
	}
	return false
}

func jsonTagOptions(tag string) []string {
	idx := strings.Index(tag, ",")
	if idx == -1 {
		return nil
	}
	return strings.Split(tag[idx+1:], ",")
}

func jsonTagHasOption(tag, name string) bool {
	for _, opt := range jsonTagOptions(tag) {
		if opt == name {
			return true
		}
	}
	return false
}

func setFieldValue(field reflect.Value, rawValue json.RawMessage, jsonTag string) error {
	if jsonTagHasOption(jsonTag, "string") {
		return setFieldFromJSONString(field, rawValue)
	}
	if hasJSONTagOptions(jsonTag) {
		return json.Unmarshal(rawValue, field.Addr().Interface())
	}
	switch field.Type() {
	case typeFloat64:
		var f Float64
		if err := f.UnmarshalJSON(rawValue); err != nil {
			return err
		}
		field.Set(reflect.ValueOf(f))
		return nil
	case typeInt64:
		var i Int64
		if err := i.UnmarshalJSON(rawValue); err != nil {
			return err
		}
		field.Set(reflect.ValueOf(i))
		return nil
	case typeMixed:
		field.Set(reflect.ValueOf(Mixed(unquote(rawValue))))
		return nil
	case typeRawMessage:
		field.Set(reflect.ValueOf(append(json.RawMessage(nil), rawValue...)))
		return nil
	}

	switch field.Kind() {
	case reflect.String:
		var s string
		if err := json.Unmarshal(rawValue, &s); err != nil {
			return err
		}
		field.SetString(s)
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		var n float64
		if err := json.Unmarshal(rawValue, &n); err != nil {
			return err
		}
		field.SetInt(int64(n))
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		var n float64
		if err := json.Unmarshal(rawValue, &n); err != nil {
			return err
		}
		field.SetUint(uint64(n))
	case reflect.Float32, reflect.Float64:
		var n float64
		if err := json.Unmarshal(rawValue, &n); err != nil {
			return err
		}
		field.SetFloat(n)
	case reflect.Bool:
		var b bool
		if err := json.Unmarshal(rawValue, &b); err != nil {
			return err
		}
		field.SetBool(b)
	default:
		return json.Unmarshal(rawValue, field.Addr().Interface())
	}
	return nil
}

func setFieldFromJSONString(field reflect.Value, rawValue json.RawMessage) error {
	var s string
	if err := json.Unmarshal(rawValue, &s); err != nil {
		return err
	}
	switch field.Kind() {
	case reflect.String:
		field.SetString(s)
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		n, err := strconv.ParseInt(s, 10, 64)
		if err != nil {
			return err
		}
		field.SetInt(n)
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		n, err := strconv.ParseUint(s, 10, 64)
		if err != nil {
			return err
		}
		field.SetUint(n)
	case reflect.Float32, reflect.Float64:
		n, err := strconv.ParseFloat(s, 64)
		if err != nil {
			return err
		}
		field.SetFloat(n)
	case reflect.Bool:
		b, err := strconv.ParseBool(s)
		if err != nil {
			return err
		}
		field.SetBool(b)
	default:
		return fmt.Errorf("unsupported type for json string option: %s", field.Type())
	}
	return nil
}
