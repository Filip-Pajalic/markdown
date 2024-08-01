package markdown

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

func Decode(markdown string, v interface{}) error {
	lines := strings.Split(markdown, "\n")
	val := reflect.ValueOf(v).Elem()
	typ := val.Type()

	for _, line := range lines {
		if strings.HasPrefix(line, "## ") {
			header := strings.TrimPrefix(line, "## ")
			for i := 0; i < val.NumField(); i++ {
				fieldType := typ.Field(i)
				tag := fieldType.Tag.Get("markdown")
				if tag == "header" {
					val.Field(i).SetString(header)
					break
				}
			}
		} else if strings.HasPrefix(line, "- **") {
			line = strings.TrimPrefix(line, "- **")
			parts := strings.SplitN(line, "**: ", 2)
			if len(parts) != 2 {
				continue
			}
			key := parts[0]
			value := parts[1]
			for i := 0; i < val.NumField(); i++ {
				fieldType := typ.Field(i)
				tag := fieldType.Tag.Get("markdown")
				tagParts := strings.Split(tag, ",")
				if len(tagParts) > 1 && tagParts[1] == key {
					field := val.Field(i)
					switch field.Kind() {
					case reflect.String:
						field.SetString(value)
					case reflect.Int:
						num, err := strconv.Atoi(value)
						if err != nil {
							return err
						}
						field.SetInt(int64(num))
					}
					break
				}
			}
		}
	}

	return nil
}

func Encode(v interface{}) (string, error) {
	var sb strings.Builder
	val := reflect.ValueOf(v)
	typ := val.Type()

	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		fieldType := typ.Field(i)
		tag := fieldType.Tag.Get("markdown")
		if tag == "" {
			continue
		}
		parts := strings.Split(tag, ",")
		tagType := parts[0]
		tagName := fieldType.Name
		if len(parts) > 1 {
			tagName = parts[1]
		}

		switch tagType {
		case "header":
			sb.WriteString(fmt.Sprintf("## %v\n", field.Interface()))
		case "item":
			sb.WriteString(fmt.Sprintf("- **%s**: %v\n", tagName, field.Interface()))
		}
	}

	return sb.String(), nil
}
