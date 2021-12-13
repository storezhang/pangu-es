package elasticsearch

import (
	`encoding/json`
	`errors`
	`fmt`
	`strings`

	`github.com/json-iterator/go`
)

func (c *Client) getFieldVal(field string, from interface{}) (val interface{}, err error) {
	var data []byte
	if data, err = json.Marshal(from); nil != err {
		return
	}

	_fields := strings.Split(field, `.`)
	var temp jsoniter.Any
	for i, _field := range _fields {
		if i == 0 {
			temp = jsoniter.Get(data, _field)
		} else {
			temp = temp.Get(_field)
		}
		if temp == nil {
			err = errors.New(fmt.Sprintf(`未找到字段%s`, field))
		}
	}

	val = temp.GetInterface()

	return
}
