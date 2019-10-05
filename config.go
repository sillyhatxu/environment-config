package envconfig

import (
	"bufio"
	"bytes"
	"fmt"
	"github.com/BurntSushi/toml"
	"io/ioutil"
	"os"
	"reflect"
	"strconv"
	"strings"
)

func ParseConfig(input interface{}, opts ...Option) error {
	//default
	config := &Config{
		configFile:  defaultConfigFile,
		environment: defaultLocalEnvironment,
	}
	for _, opt := range opts {
		opt(config)
	}
	if fileInfo, err := os.Stat(config.configFile); err != nil {
		if os.IsNotExist(err) {
			return fmt.Errorf("configuration file [%s] does not exist.", config.configFile)
		} else {
			return fmt.Errorf("configuration file [%s] can not be stated. %v", config.configFile, err)
		}
	} else {
		if fileInfo.IsDir() {
			return fmt.Errorf("%s is a directory name", config.configFile)
		}
	}
	content, err := ioutil.ReadFile(config.configFile)
	if err != nil {
		return fmt.Errorf("read configuration file error. %v", err)
	}
	content = bytes.TrimSpace(content)
	err = toml.Unmarshal(content, input)
	if err != nil {
		return fmt.Errorf("unmarshal toml object error. %v", err)
	}
	if config.environment == defaultLocalEnvironment {
		err := loadEnv(defaultLocalEnvFile)
		if err != nil {
			return err
		}
	}
	err = parseEnvironmentConfig(input)
	if err != nil {
		return err
	}
	return nil
}

func loadEnv(fileName string) error {
	file, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if scanner.Text() == "" {
			continue
		}
		array := strings.Split(scanner.Text(), "=")
		if array == nil || len(array) != 2 {
			return fmt.Errorf("split error. %s", scanner.Text())
		}
		err := os.Setenv(array[0], array[1])
		if err != nil {
			return err
		}
	}
	return nil
}

const tagName = `env`

//func parseEnvironmentConfig(obj interface{}) error {
//	objT := reflect.TypeOf(obj)
//	objV := reflect.ValueOf(obj)
//	switch {
//	case isStruct(objT):
//	case isStructPtr(objT):
//		objT = objT.Elem()
//		objV = objV.Elem()
//	default:
//		return fmt.Errorf("%v must be a struct or a struct pointer", obj)
//	}
//	return parse(objT, objV)
//}
//
//func parse(t reflect.Type, v reflect.Value) error {
//	fmt.Println("Type is", t.Name(), "and kind is", t.Kind())
//	for i := 0; i < t.NumField(); i++ {
//		f := t.Field(i)
//		fmt.Println(fmt.Sprintf("Name : %s; Type : %s; Kind : %s; Tag : %s", f.Name, f.Type.Name(), f.Type.Kind(), f.Tag))
//		tag := f.Tag.Get(tagName)
//		if tag == "" {
//			continue
//		} else if tag == "-" {
//			fieldValue := v.Elem()
//			return parseEnvironmentConfig(fieldValue)
//		}
//		if f.Tag != "" {
//			fmt.Println("Tag is", f.Tag)
//			fmt.Println("tag1 is", f.Tag.Get("tag1"), "tag2 is", f.Tag.Get("tag2"))
//		}
//	}
//	return nil
//}

func parseEnvironmentConfig(obj interface{}) error {
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
		if tag == "" {
			continue
		}
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
