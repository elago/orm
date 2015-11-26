package orm

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"reflect"
)

func RegisterModel(model interface{}) {
	value := reflect.ValueOf(model)
	typ := value.Type()
	for i := 0; i < value.NumMethod(); i++ {
		fmt.Printf("method[%d]%s\n", i, typ.Method(i).Name)
	}
}
