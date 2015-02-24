package configr

import (
	"flag"
	"fmt"
	"os"
	"reflect"
	"strconv"
)

//UnmarshalFromEnv will marshal values tagged with env to the passed in struct
func UnmarshalFromEnv(c interface{}) {
	topType := reflect.TypeOf(c).Elem()
	topValue := reflect.ValueOf(c)
	for i := 0; i < topType.NumField(); i++ {
		field := topType.Field(i)
		if field.Tag.Get("env") != "" {
			envVar := os.Getenv(field.Tag.Get("env"))
			if envVar == "" {
				continue
			}
			switch field.Type.Kind() {
			case reflect.Bool:
				b, err := strconv.ParseBool(envVar)
				if err != nil {
					fmt.Println("didn't set from ", field.Tag.Get("env"), " due to ", err)
					continue
				}
				f := topValue.Elem().Field(i)
				f.SetBool(b)
			case reflect.Int64:
				integer, err := strconv.ParseInt(envVar, 0, 64)
				if err != nil {
					fmt.Println("didn't set from ", field.Tag.Get("env"), " due to ", err)
					continue
				}
				f := topValue.Elem().Field(i)
				f.SetInt(integer)
			case reflect.String:
				f := topValue.Elem().Field(i)
				f.SetString(envVar)
			case reflect.Float64:
				float, err := strconv.ParseFloat(envVar, 64)
				if err != nil {
					fmt.Println("didn't set from ", field.Tag.Get("env"), " due to ", err)
					continue
				}
				f := topValue.Elem().Field(i)
				f.SetFloat(float)
			}
		}
	}
}

//UnmarshalFromFlags will register a cli flag when a struct element is tagged with cli and describe it as it is tagged in desc
func UnmarshalFromFlags(c interface{}) {
	topType := reflect.TypeOf(c).Elem()
	topValue := reflect.ValueOf(c)
	for i := 0; i < topType.NumField(); i++ {
		field := topType.Field(i)
		if field.Tag.Get("cli") != "" {
			desc := field.Tag.Get("desc")
			if field.Tag.Get("env") != "" {
				desc = desc + "; This variable is set by ENV \"" + field.Tag.Get("env") + "\""
			}
			switch field.Type.Kind() {
			case reflect.Bool:
				temp := flag.Bool(field.Tag.Get("cli"), topValue.Elem().Field(i).Bool(), desc)
				f := topValue.Elem().Field(i)
				defer func() { f.SetBool(*temp) }()
			case reflect.Int64:
				temp := flag.Int64(field.Tag.Get("cli"), topValue.Elem().Field(i).Int(), desc)
				f := topValue.Elem().Field(i)
				defer func() { f.SetInt(*temp) }()
			case reflect.String:
				temp := flag.String(field.Tag.Get("cli"), topValue.Elem().Field(i).String(), desc)
				f := topValue.Elem().Field(i)
				defer func() { f.SetString(*temp) }()
			case reflect.Float64:
				temp := flag.Float64(field.Tag.Get("cli"), topValue.Elem().Field(i).Float(), desc)
				f := topValue.Elem().Field(i)
				defer func() { f.SetFloat(*temp) }()
			}
		}
	}
	flag.Parse()
}
