package orm

import (
	"database/sql"
	"github.com/gogather/com/log"
	"reflect"
)

var db *sql.DB
var models map[string]*Model

type Model struct {
	typ    reflect.Type
	fields map[string]Field
}

type Field struct {
	sqlType reflect.Type
}

func PrintModels() {
	// log.Blueln(models["user_log"])
	log.Bluef("\n%s\n", models)
}
