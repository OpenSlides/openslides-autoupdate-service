package datastore

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

// GetStructJsonNames gets all json-tagged names from struct. Includes replacing
// struct fields by replacements annotated in struct tag
func GetStructJsonNames(value interface{}) ([]string, []int, error) {
	v := reflect.ValueOf(value)
	t := reflect.TypeOf(v.Interface())
	var names []string
	var indices []int
	for i := 0; i < v.NumField(); i++ {

		f := t.Field(i)
		tag := f.Tag.Get("json")
		if tag == "" {
			continue
		}

		commaIndex := strings.Index(tag, ",")
		if commaIndex >= 0 {
			tag = tag[:commaIndex]
		}

		if strings.Contains(tag, "$") {
			replacement := f.Tag.Get("replacement")
			if replacement == "" {
				return nil, nil, fmt.Errorf("not implemented error in getstructjsonnames for field %s", f.Name)
			}
			rfield := v.FieldByName(replacement)
			vrepl := ""
			if rfield.Kind() == reflect.String {
				vrepl = rfield.String()
			} else if rfield.Kind() == reflect.Int {
				vrepl = strconv.FormatInt(rfield.Int(), 10)
			} else {
				return nil, nil, fmt.Errorf("the value of the replacement '%s' value must be a string or int", f.Name)
			}
			if vrepl == "" {
				return nil, nil, fmt.Errorf("there is no value in the struct for replacement '%s'", replacement)
			}
			tag = strings.Replace(tag, "$", "$"+vrepl, 1)
		}
		indices = append(indices, i)
		names = append(names, tag)
	}
	return names, indices, nil
}
