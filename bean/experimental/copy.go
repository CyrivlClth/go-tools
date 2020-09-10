// package experimental
package experimental

import (
	"errors"
	"reflect"
)

func CopyProperties(s, t interface{}) error {
	tType := reflect.TypeOf(t)
	if tType.Kind() != reflect.Ptr {
		return errors.New("accept pointer only")
	}
	tValue := reflect.ValueOf(t).Elem()
	sType := reflect.TypeOf(s)
	if sType.Kind() == reflect.Ptr {
		sType = sType.Elem()
	}
	sValue := reflect.ValueOf(s)
	if sValue.Kind() == reflect.Ptr {
		sValue = sValue.Elem()
	}
	for i := 0; i < sType.NumField(); i++ {
		t := sType.Field(i)
		h := sValue.Field(i)
		f := tValue.FieldByName(t.Name)
		f.Set(reflect.ValueOf(h.Interface()).Convert(f.Type()))
	}

	return nil
}
