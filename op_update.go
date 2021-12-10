package elasticsearch

import (
	`context`
	`errors`
	`fmt`
	`net/http`
	`reflect`
	`strings`

	`github.com/olivere/elastic/v7`
)

func (c *Client) UpdateByScript(
	index string, docId string, script *elastic.Script,
) (result *elastic.UpdateResponse, err error) {
	_update := elastic.NewUpdateService(c.Client)
	result, err = _update.
		Index(index).Id(docId).
		Script(script).
		RetryOnConflict(5).
		Refresh(`true`).
		Do(context.Background())

	if nil != err {
		if elasticErr, ok := err.(*elastic.Error); ok {
			if http.StatusNotFound == elasticErr.Status {
				err = nil
			}
		}
	}

	return
}

func (c *Client) UpdateByQueryAndScript(
	index string, query elastic.Query, script *elastic.Script,
) (res *elastic.BulkIndexByScrollResponse, err error) {
	_update := elastic.NewUpdateByQueryService(c.Client)
	res, err = _update.
		Index(index).
		Query(query).
		Script(script).
		ProceedOnVersionConflict().
		Refresh(`true`).
		Do(context.Background())

	if nil != err {
		if elasticErr, ok := err.(*elastic.Error); ok {
			if http.StatusNotFound == elasticErr.Status {
				err = nil
			}
		}
	}

	return
}

func (c *Client) UpdateByFields(
	index, docId string, cond interface{}, fields ...string,
) (result *elastic.UpdateResponse, err error) {
	_fields := make(map[string]interface{}, len(fields))
	for _, field := range fields {
		if _fields[field], err = c.getFieldVal(field, cond); nil != err {
			return
		}
	}
	result, err = c.Update().
		Index(index).
		Id(docId).
		Doc(_fields).
		Refresh(`true`).
		Do(context.Background())

	return
}

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
			val = v.FieldByName(t.Field(i).Name)
			break
		}
	}

	if nil == val {
		err = errors.New(fmt.Sprintf(`未找到字段%s`, field))
	}

	return
}
