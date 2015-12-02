package orm

import (
	"database/sql"
	// "fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gogather/com/log"
	"reflect"
)

var typeMap map[string]interface{}

func init() {
	var err error
	db, err = sql.Open("mysql", "root:lijun@tcp(127.0.0.1:3306)/blog")

	if err != nil {
		log.Warnln(err)
	}

	db.SetMaxOpenConns(2000)
	db.SetMaxIdleConns(1000)
	db.Ping()

	models = make(map[string]*Model)

	// data := query("users", 8)
	// log.Pinkln(data)
}

func TestQuery() {
	query("users", 8)
}

// query the bean
func query(tableName string, id int64) {
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

	model := models[tableName]

	val := reflect.New(model.typ).Elem()
	f := val.Field(0)
	f1 := val.Field(1)
	// TODO
	// log.Yellowln(f.Type().Name())
	// f.SetInt(8)
	assignField(&f, 8)
	assignField(&f1, "hello")
	// f.Set(8)
	log.Greenln(f)
	log.Pinkln(val)

}

func RegisterModel(model interface{}) {
	value := reflect.ValueOf(model)
	indir := reflect.Indirect(value)
	typ := indir.Type()

	m := &Model{}
	m.typ = typ
	m.fields = make(map[string]Field)

	for i := 0; i < indir.NumField(); i++ {
		var f Field
		f.sqlType = typ.Field(i).Type
		sqlField := camel2Snake(typ.Field(i).Name)

		m.fields[sqlField] = f
	}

	// log.Pinkln(m)

	// log.Pinkf("put [%s] into [%s] \n", m.name, m)
	models[camel2Snake(typ.Name())] = m
}

func getTypeName(model interface{}) string {
	value := reflect.ValueOf(model)
	indir := reflect.Indirect(value)
	typ := indir.Type()
	return typ.Name()
}
