package jsonrender

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
)

func Iterate(obj map[string]interface{}, indent int) {
	for key, value := range obj {
		padding := strings.Repeat("  ", indent)
		if reflect.TypeOf(value).Kind() == reflect.Map {
			fmt.Printf("%s%s: {\n", padding, key)
			Iterate(value.(map[string]interface{}), indent+1)
			fmt.Printf("%s}\n", padding)
		} else if reflect.TypeOf(value).Kind() == reflect.Slice {
			fmt.Printf("%s%s: [\n", padding, key)
			slice := reflect.ValueOf(value)
			for i := 0; i < slice.Len(); i++ {
				item, _ := json.Marshal(slice.Index(i).Interface())
				fmt.Printf("%s  %s,\n", padding, item)
			}
			fmt.Printf("%s]\n", padding)
		} else {
			item, _ := json.Marshal(value)
			fmt.Printf("%s%s: %s,\n", padding, key, item)
		}
	}
}