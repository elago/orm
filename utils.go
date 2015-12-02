package orm

import (
	"fmt"
	"reflect"
	"strings"
	// "time"
)

func camel2Snake(s string) string {
	data := make([]byte, 0, len(s)*2)
	j := false
	num := len(s)
	for i := 0; i < num; i++ {
		d := s[i]
		if i > 0 && d >= 'A' && d <= 'Z' && j {
			data = append(data, '_')
		}
		if d != '_' {
			j = true
		}
		data = append(data, d)
	}
	return strings.ToLower(string(data[:]))
}

func assertType(data interface{}) interface{} {
	if valStr, ok := data.([]byte); ok {
		return string(valStr)
	} else if valInt, ok := data.(int); ok {
		return int(valInt)
	} else if valInt8, ok := data.(int8); ok {
		return int8(valInt8)
	} else if valInt16, ok := data.(int16); ok {
		return int16(valInt16)
	} else if valInt32, ok := data.(int32); ok {
		return int32(valInt32)
	} else if valInt64, ok := data.(int64); ok {
		return int64(valInt64)
	} else if valUint, ok := data.(uint); ok {
		return uint(valUint)
	} else if valUint8, ok := data.(uint8); ok {
		return uint8(valUint8)
	} else if valUint16, ok := data.(uint16); ok {
		return uint16(valUint16)
	} else if valUint32, ok := data.(uint32); ok {
		return uint32(valUint32)
	} else if valUint64, ok := data.(uint64); ok {
		return uint64(valUint64)
	} else if valFloat32, ok := data.(float32); ok {
		return float32(valFloat32)
	} else if valFloat64, ok := data.(float64); ok {
		return float64(valFloat64)
	} else if valBool, ok := data.(bool); ok {
		return bool(valBool)
	} else {
		return data
	}
}

func assertWideType(data interface{}) interface{} {
	if valStr, ok := data.([]byte); ok {
		return string(valStr)
	} else if valInt, ok := data.(int); ok {
		return int64(valInt)
	} else if valInt8, ok := data.(int8); ok {
		return int64(valInt8)
	} else if valInt16, ok := data.(int16); ok {
		return int64(valInt16)
	} else if valInt32, ok := data.(int32); ok {
		return int64(valInt32)
	} else if valInt64, ok := data.(int64); ok {
		return int64(valInt64)
	} else if valUint, ok := data.(uint); ok {
		return uint64(valUint)
	} else if valUint8, ok := data.(uint8); ok {
		return uint64(valUint8)
	} else if valUint16, ok := data.(uint16); ok {
		return uint64(valUint16)
	} else if valUint32, ok := data.(uint32); ok {
		return uint64(valUint32)
	} else if valUint64, ok := data.(uint64); ok {
		return uint64(valUint64)
	} else if valFloat32, ok := data.(float32); ok {
		return float64(valFloat32)
	} else if valFloat64, ok := data.(float64); ok {
		return float64(valFloat64)
	} else if valBool, ok := data.(bool); ok {
		return bool(valBool)
	} else {
		return data
	}
}

func assignField(f *reflect.Value, value interface{}) {
	switch f.Type().Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		f.SetInt(assertWideType(value).(int64))
	case reflect.Float32, reflect.Float64:
		f.SetFloat(value.(float64))
	case reflect.String:
		f.SetString(value.(string))
	case reflect.Struct:
		f.SetBytes(value.([]byte))
	case reflect.Bool:
		f.SetBool(value.(bool))
	default:
		f.SetBytes(value.([]byte))
	}
}

func checkErr(err error) {
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
}
