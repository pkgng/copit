package copit

import (
	"database/sql"
	"errors"
	"reflect"
)

// Copy copy things
func Copy(toValue interface{}, fromValue interface{}) (err error) {
	var (
		from = indirect(reflect.ValueOf(fromValue))
		to   = indirect(reflect.ValueOf(toValue))
	)

	if !to.CanAddr() {
		return errors.New("copy to value is unaddressable")
	}

	// Return is from value is invalid
	if !from.IsValid() {
		return
	}

	fromType := indirectType(from.Type())
	toType := indirectType(to.Type())

	// Just set it if possible to assign
	// And need to do copy anyway if the type is struct
	if fromType.Kind() != reflect.Struct && from.Type().AssignableTo(to.Type()) {
		to.Set(from)
		return
	}

	// slice -> slice Section
	if to.Kind() == reflect.Slice {
		if from.Kind() == reflect.Slice {
			for i := 0; i < from.Len(); i++ {
				if indirect(from.Index(i)).IsValid() {
					Copy(toValue, indirect(from.Index(i)).Interface())
				}
			}
		} else if fromType.Kind() == reflect.Struct {
			dest := indirect(reflect.New(toType).Elem())
			if err := Copy(dest.Addr().Interface(), indirect(from).Interface()); err != nil {
				return err
			}

			if dest.Addr().Type().AssignableTo(to.Type().Elem()) {
				to.Set(reflect.Append(to, dest.Addr()))
			} else if dest.Type().AssignableTo(to.Type().Elem()) {
				to.Set(reflect.Append(to, dest))
			}
		}
		return
	}

	// ---------------------  struct -> struct only
	if fromType.Kind() != reflect.Struct || toType.Kind() != reflect.Struct {
		return
	}

	if true {
		toTypeFields := deepFields(toType)
		for _, field := range toTypeFields {
			name := field.Name
			if tagName := field.Tag.Get("copit"); tagName != "" {
				name = tagName
			}

			if fromField := from.FieldByName(name); fromField.IsValid() {
				if toField := to.FieldByName(field.Name); toField.IsValid() {
					if toField.CanSet() {
						if !set(toField, fromField) {
							if err := Copy(toField.Addr().Interface(), fromField.Interface()); err != nil {
								return err
							}
						}
					}
				}
			}
		}
	}

	if true {
		fromTypeFields := deepFields(fromType)
		for _, field := range fromTypeFields {
			name := field.Name

			if fromField := from.FieldByName(name); fromField.IsValid() {

				var toMethod reflect.Value
				if to.CanAddr() {
					toMethod = to.Addr().MethodByName(name)
				} else {
					toMethod = to.MethodByName(name)
				}

				if toMethod.IsValid() && toMethod.Type().NumIn() == 1 && fromField.Type().AssignableTo(toMethod.Type().In(0)) {
					toMethod.Call([]reflect.Value{fromField})
				}
			}
		}

	}

	return
}

func deepFields(reflectType reflect.Type) []reflect.StructField {
	var fields []reflect.StructField

	if reflectType = indirectType(reflectType); reflectType.Kind() == reflect.Struct {
		for i := 0; i < reflectType.NumField(); i++ {
			v := reflectType.Field(i)
			if v.Anonymous {
				fields = append(fields, deepFields(v.Type)...)
			} else {
				fields = append(fields, v)
			}
		}
	}

	return fields
}

func indirect(reflectValue reflect.Value) reflect.Value {
	for reflectValue.Kind() == reflect.Ptr {
		reflectValue = reflectValue.Elem()
	}
	return reflectValue
}

func indirectType(reflectType reflect.Type) reflect.Type {
	for reflectType.Kind() == reflect.Ptr || reflectType.Kind() == reflect.Slice {
		reflectType = reflectType.Elem()
	}
	return reflectType
}

func set(to, from reflect.Value) bool {
	if from.IsValid() {
		if to.Kind() == reflect.Ptr {
			//set `to` to nil if from is nil
			if from.Kind() == reflect.Ptr && from.IsNil() {
				to.Set(reflect.Zero(to.Type()))
				return true
			} else if to.IsNil() {
				to.Set(reflect.New(to.Type().Elem()))
			}
			to = to.Elem()
		}

		if from.Type().ConvertibleTo(to.Type()) {
			to.Set(from.Convert(to.Type()))
		} else if scanner, ok := to.Addr().Interface().(sql.Scanner); ok {
			err := scanner.Scan(from.Interface())
			if err != nil {
				return false
			}
		} else if from.Kind() == reflect.Ptr {
			return set(to, from.Elem())
		} else {
			return false
		}
	}
	return true
}
