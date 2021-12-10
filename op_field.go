package elasticsearch

import (
	`errors`
	`fmt`
	`reflect`
	`strings`
)

func (c *Client) getFieldVal(field string, from interface{}) (val interface{}, err error) {
	t := reflect.TypeOf(from)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	if t.Kind() != reflect.Struct {
		return
	}

	fieldNum := t.NumField()
	for i := 0; i < fieldNum; i++ {
		if strings.ToUpper(t.Field(i).Name) == strings.ToUpper(field) {
			v := reflect.Indirect(reflect.ValueOf(from))
			val = v.FieldByName(t.Field(i).Name).Interface()
			break
		}
	}

	if nil == val {
		err = errors.New(fmt.Sprintf(`未找到字段%s`, field))
	}

	return
}
