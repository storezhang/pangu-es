package elasticsearch

import (
	`context`

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
