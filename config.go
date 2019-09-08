package envconfig

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"reflect"
	"strconv"
)

func ParseConfig(configFile string, unmarshalfunc func([]byte)) {
	if fileInfo, err := os.Stat(configFile); err != nil {
		if os.IsNotExist(err) {
			panic(fmt.Sprintf("configuration file [%s] does not exist.", configFile))
		} else {
			panic(fmt.Sprintf("configuration file [%s] can not be stated. %v", configFile, err))
		}
	} else {
		if fileInfo.IsDir() {
			panic(fmt.Sprintf("%s is a directory name", configFile))
		}
	}
	content, err := ioutil.ReadFile(configFile)
	if err != nil {
		panic(fmt.Sprintf("read configuration file error. %v", err))
	}
	content = bytes.TrimSpace(content)
	unmarshalfunc(content)
}

const tagName = `env`

func ParseEnvironmentConfig(obj interface{}) error {
	objT := reflect.TypeOf(obj)
	objV := reflect.ValueOf(obj)
	switch {
	case isStruct(objT):
	case isStructPtr(objT):
		objT = objT.Elem()
		objV = objV.Elem()
	default:
		return fmt.Errorf("%v must be a struct or a struct pointer", obj)
	}
	for i := 0; i < objT.NumField(); i++ {
		field := objT.Field(i)
		tag := field.Tag.Get(tagName)
		envValue := os.Getenv(tag)
		fieldValue := reflect.ValueOf(obj).Elem()
		switch field.Type.Kind() {
		case reflect.String:
			fieldValue.FieldByName(field.Name).SetString(envValue)
		case reflect.Int:
			i, err := strconv.Atoi(envValue)
			if err != nil {
				return fmt.Errorf("strconv.Atoi(%s) error. %v", envValue, err)
			}
			fieldValue.FieldByName(field.Name).SetInt(int64(i))
		case reflect.Int64:
			i, err := strconv.ParseInt(envValue, 10, 64)
			if err != nil {
				return fmt.Errorf("strconv.ParseInt(%v, 10, 64) error. %v", envValue, err)
			}
			fieldValue.FieldByName(field.Name).SetInt(i)
		case reflect.Int32:
			i, err := strconv.ParseInt(envValue, 10, 32)
			if err != nil {
				return fmt.Errorf("strconv.ParseInt(%v, 10, 64) error. %v", envValue, err)
			}
			fieldValue.FieldByName(field.Name).SetInt(i)
		case reflect.Int16:
			i, err := strconv.ParseInt(envValue, 10, 16)
			if err != nil {
				return fmt.Errorf("strconv.ParseInt(%v, 10, 64) error. %v", envValue, err)
			}
			fieldValue.FieldByName(field.Name).SetInt(i)
		case reflect.Int8:
			i, err := strconv.ParseInt(envValue, 10, 8)
			if err != nil {
				return fmt.Errorf("strconv.ParseInt(%v, 10, 64) error. %v", envValue, err)
			}
			fieldValue.FieldByName(field.Name).SetInt(i)
		case reflect.Bool:
			b, err := strconv.ParseBool(envValue)
			if err != nil {
				return fmt.Errorf("strconv.ParseInt(%v, 10, 64) error. %v", envValue, err)
			}
			fieldValue.FieldByName(field.Name).SetBool(b)
		case reflect.Float64:
			f, err := strconv.ParseFloat(envValue, 64)
			if err != nil {
				return fmt.Errorf("strconv.ParseInt(%v, 10, 64) error. %v", envValue, err)
			}
			fieldValue.FieldByName(field.Name).SetFloat(f)
		case reflect.Float32:
			f, err := strconv.ParseFloat(envValue, 32)
			if err != nil {
				return fmt.Errorf("strconv.ParseInt(%v, 10, 64) error. %v", envValue, err)
			}
			fieldValue.FieldByName(field.Name).SetFloat(f)
		default:
			return fmt.Errorf("not support %s", field.Type.Kind().String())
		}
	}
	return nil
}

func isStruct(t reflect.Type) bool {
	return t.Kind() == reflect.Struct
}

func isStructPtr(t reflect.Type) bool {
	return t.Kind() == reflect.Ptr && t.Elem().Kind() == reflect.Struct
}
