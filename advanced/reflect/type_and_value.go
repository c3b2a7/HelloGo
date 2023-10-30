package reflect

import (
	"fmt"
	reflect "reflect"
	"strings"
)

func TypeAndValue(e interface{}) {
	t := reflect.TypeOf(e)
	v := reflect.ValueOf(e)
	fmt.Println("type: ", t)
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		fmt.Println(field.Anonymous, field.PkgPath)
		fmt.Println("FieldName:", field.Name, "FiledType:", field.Type, "FiledValue:", v.Field(i))
		fmt.Println()
		v.Field(i).Interface()
	}
}

func CreateInsert(e interface{}) string {
	t, v := reflectTypeAndValue(e)

	var columns []string
	var values []string

	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		validKind(f.Type.Kind())
		columns = append(columns, f.Name)
		values = append(values, SprintAny(v.Field(i).Interface()))
	}

	tableName := t.Name()

	return fmt.Sprintf("INSERT INTO %s (%s); VALUES (%s)", tableName,
		strings.Join(columns, ", "), strings.Join(values, ", "))
}

func Invoke(fun interface{}, params ...interface{}) []interface{} {
	t := reflect.TypeOf(fun)
	if t.Kind() != reflect.Func {
		panic("invalid kind: " + t.Kind().String())
	}

	if t.NumIn() != len(params) {
		panic("invalid input")
	}

	var values []reflect.Value
	for _, param := range params {
		values = append(values, reflect.ValueOf(param))
	}

	var out []interface{}
	ret := reflect.ValueOf(fun).Call(values)
	//errInterface := reflect.TypeOf((*error)(nil)).Elem()
	for i := 0; i < len(ret); i++ {
		out = append(out, ret[i].Interface())
	}
	return out
}

func reflectTypeAndValue(e interface{}) (reflect.Type, reflect.Value) {
	t := reflect.TypeOf(e)
	switch t.Kind() {
	case reflect.Pointer:
		//return reflectTypeAndValue(reflect.ValueOf(e).Elem().Interface())
		return reflectTypeAndValue(reflect.ValueOf(e).Elem().Interface())
	case reflect.Struct:
		return t, reflect.ValueOf(e)
	}
	panic("e must be struct")
}

func validKind(k reflect.Kind) {
	switch k {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		fallthrough
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		fallthrough
	case reflect.Bool:
		fallthrough
	case reflect.String:
		return
	}
	panic("kind " + k.String() + " is invalid")
}
