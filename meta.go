package meta

import (
	"errors"
	"reflect"
	"strconv"
)

//Group a group is consists of multiple Meta with the same name
type Group []*Meta

type Meta struct {
	name       string
	properties map[string]string
}

func New(name string) *Meta {
	return &Meta{
		name:       name,
		properties: map[string]string{},
	}
}

func (m *Meta) Name() string {
	return m.name
}

func (m *Meta) Property(key string) string {
	return m.properties[key]
}

func (m *Meta) SetProperty(key, value string) *Meta {
	m.properties[key] = value
	return m
}

func (m *Meta) Properties() map[string]string {
	return m.properties
}

func (m *Meta) SetProperties(properties map[string]string) *Meta {
	for k, v := range properties {
		m.properties[k] = v
	}
	return m
}

func (m *Meta) MapStruct(obj any) error {
	objPointerType := reflect.TypeOf(obj)
	if objPointerType.Kind() != reflect.Pointer {
		return errors.New("MapStruct param must be a Pointer")
	}
	objType := objPointerType.Elem()
	objPointerValue := reflect.ValueOf(obj)
	objValue := reflect.Indirect(objPointerValue)
	for i := 0; i < objValue.NumField(); i++ {
		field := objType.Field(i)
		strVal := m.Property(field.Name)
		if strVal == "" {
			continue
		}
		fieldValue := objValue.Field(i)

		err := setValueFromString(field, fieldValue, strVal)
		if err != nil {
			return err
		}
	}
	return nil
}

func setValueFromString(field reflect.StructField, v reflect.Value, strVal string) error {
	switch field.Type.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		val, err := strconv.ParseInt(strVal, 0, 64)
		if err != nil {
			return err
		}
		if v.OverflowInt(val) {
			return errors.New("Int value too big: " + strVal)
		}
		v.SetInt(val)
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		val, err := strconv.ParseUint(strVal, 0, 64)
		if err != nil {
			return err
		}
		if v.OverflowUint(val) {
			return errors.New("UInt value too big: " + strVal)
		}
		v.SetUint(val)
	case reflect.Float32:
		val, err := strconv.ParseFloat(strVal, 32)
		if err != nil {
			return err
		}
		v.SetFloat(val)
	case reflect.Float64:
		val, err := strconv.ParseFloat(strVal, 64)
		if err != nil {
			return err
		}
		v.SetFloat(val)
	case reflect.String:
		v.SetString(strVal)
	case reflect.Bool:
		if field.Name == strVal {
			v.SetBool(true)
			return nil
		}
		val, err := strconv.ParseBool(strVal)
		if err != nil {
			return err
		}
		v.SetBool(val)
	default:
		return errors.New("Unsupported kind: " + v.Kind().String())
	}
	return nil
}
