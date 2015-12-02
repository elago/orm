package orm

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gogather/com/log"
	"reflect"
	"strings"
)

var db *sql.DB
var models *Model

type Model struct {
	name   string
	fields map[string]Field
}

type Field struct {
	sqlType string
}

func init() {
	var err error
	db, err = sql.Open("mysql", "root:lijun@tcp(127.0.0.1:3306)/blog")

	if err != nil {
		log.Warnln(err)
	}

	db.SetMaxOpenConns(2000)
	db.SetMaxIdleConns(1000)
	db.Ping()

	// data := query("users", 8)
	// log.Pinkln(data)
}

// query the bean
func query(tableName string, id int64) map[string]interface{} {
	rows, err := db.Query("SELECT * FROM `"+tableName+"` where id=? limit 1", id)
	defer rows.Close()
	checkErr(err)

	columns, _ := rows.Columns()
	scanArgs := make([]interface{}, len(columns))
	values := make([]interface{}, len(columns))
	for j := range values {
		scanArgs[j] = &values[j]
	}

	record := make(map[string]interface{})
	for rows.Next() {
		err = rows.Scan(scanArgs...)
		for i, col := range values {
			if col != nil {
				record[columns[i]] = assertType(col)
			}
		}
	}

	return record
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

func checkErr(err error) {
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
}

func RegisterModel(model interface{}) {
	value := reflect.ValueOf(model)
	indir := reflect.Indirect(value)
	typ := indir.Type()

	m := &Model{}
	m.name = camel2Snake(typ.Name())
	m.fields = make(map[string]Field)

	for i := 0; i < indir.NumField(); i++ {
		var f Field
		f.sqlType = ""
		sqlField := camel2Snake(typ.Field(i).Name)

		m.fields[sqlField] = f
	}

	log.Pinkln(m)
}

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
