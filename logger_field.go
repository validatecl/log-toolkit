package log_toolkit

import (
	"fmt"
	"reflect"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

//FieldMarshaller Interface de marshaller
type FieldMarshaller interface {
	MarshalFields(args ...interface{}) []zap.Field
}

type fieldMarshaller struct {
}

// NewFieldMarshaller Crea un nuevo Field Marshaller
func NewFieldMarshaller() FieldMarshaller {
	return &fieldMarshaller{}
}

//MarshalFields function
func (f *fieldMarshaller) MarshalFields(args ...interface{}) []zap.Field {
	zapFields := make([]zap.Field, 0)
	for _, element := range args {
		zapFields = marshallRecursive(reflect.ValueOf(element), "", zapFields)
	}

	return zapFields

}

func marshallRecursive(value reflect.Value, key string, zapFields []zap.Field) []zap.Field {
	if value.IsValid() {
		switch value.Kind() {
		case reflect.Ptr:
			zapFields = marshallRecursive(reflect.Indirect(value), "", zapFields)
		case reflect.Map:
			for _, key := range value.MapKeys() {
				fieldName := fmt.Sprintf("%v", key)
				zapFields = marshallRecursive(value.MapIndex(key), fieldName, zapFields)
			}
		case reflect.Slice:
			for i := 0; i < value.Len(); i++ {
				zapFields = marshallRecursive(value.Index(i), "", zapFields)
			}
		case reflect.Struct:
			for i := 0; i < value.NumField(); i++ {
				reflectType := value.Type()
				field := value.Field(i)
				zapFields = marshallRecursive(field, reflectType.Field(i).Name, zapFields)
			}
		case reflect.Interface:
			errorInterface := reflect.TypeOf((*error)(nil)).Elem()
			if value.Elem().Type().Implements(errorInterface) {
				zapFields = append(zapFields, zap.Field{Key: key, Interface: value, Type: zapcore.StringType, String: fmt.Sprintf("Error: %v", reflect.Indirect(value.Elem()))})
				break
			} else {
				zapFields = marshallRecursive(reflect.Indirect(value), "", zapFields)
			}
		default:
			zapFields = append(zapFields, zap.Field{Key: key, Interface: value, Type: zapcore.StringType, String: fmt.Sprintf("%v", value)})

		}
	}

	return zapFields
}
