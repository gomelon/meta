package meta

import (
	"errors"
	"fmt"
	"reflect"
)

//Group a group is consists of multiple Meta with the same name
type Group []*Meta

type Meta struct {
	qualifyName string
	properties  map[string]any
}

func New(qualifyName string) *Meta {
	return &Meta{
		qualifyName: qualifyName,
		properties:  map[string]any{},
	}
}

func (m *Meta) QualifyName() string {
	return m.qualifyName
}

func (m *Meta) Property(key string) any {
	return m.properties[key]
}

func (m *Meta) SetProperty(key, value string) *Meta {
	m.properties[key] = value
	return m
}

func (m *Meta) Properties() map[string]any {
	return m.properties
}

func (m *Meta) SetProperties(properties map[string]any) *Meta {
	for k, v := range properties {
		m.properties[k] = v
	}
	return m
}

func (m *Meta) MapTo(obj any) error {
	objPointerType := reflect.TypeOf(obj)
	if objPointerType.Kind() != reflect.Pointer {
		return errors.New("MapStruct param must be a Pointer")
	}
	objType := objPointerType.Elem()
	objPointerValue := reflect.ValueOf(obj)
	objValue := reflect.Indirect(objPointerValue)
	for i := 0; i < objValue.NumField(); i++ {
		field := objType.Field(i)
		val, ok := m.properties[field.Name]
		if !ok {
			continue
		}
		fieldValue := objValue.Field(i)

		err := setValue(fieldValue, val)
		if err != nil {
			return err
		}
	}
	return nil
}

func setValue(v reflect.Value, val any) error {
	switch v.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		intVal, ok := val.(int64)
		if !ok {
			return fmt.Errorf("expect int value, but not: value=%v", val)
		}
		if v.OverflowInt(intVal) {
			return fmt.Errorf("int value too big: %d", intVal)
		}
		v.SetInt(intVal)
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		intVal, ok := val.(int64)
		if !ok {
			return fmt.Errorf("expect int value, but not: value=%v", val)
		}
		if v.OverflowUint(uint64(intVal)) {
			return fmt.Errorf("int value too big: %d", intVal)
		}
		v.SetUint(uint64(intVal))
	case reflect.Float32, reflect.Float64:
		floatVal, ok := val.(float64)
		if !ok {
			return fmt.Errorf("expect float value, but not: value=%v", val)
		}
		if v.OverflowFloat(floatVal) {
			return fmt.Errorf("float value too big: %f", floatVal)
		}
		v.SetFloat(floatVal)
	case reflect.String:
		strVal, ok := val.(string)
		if !ok {
			return fmt.Errorf("expect string value, but not: value=%v", val)
		}
		v.SetString(strVal)
	case reflect.Bool:
		boolVal, ok := val.(bool)
		if !ok {
			return fmt.Errorf("expect bool value, but not: value=%v", val)
		}
		v.SetBool(boolVal)
	default:
		return errors.New("Unsupported kind: " + v.Kind().String())
	}
	return nil
}
