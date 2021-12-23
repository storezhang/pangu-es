package elasticsearch

import (
	`context`

	`github.com/olivere/elastic/v7`
)

func (c *Client) DeleteByDocId(index string, ids ...string) (err error) {
	if 1 == len(ids) {
		_, err = c.Delete().Index(index).Id(ids[0]).Refresh(`true`).Do(context.Background())
	} else {
		err = c.DeleteByQuery(index, elastic.NewIdsQuery().Ids(ids...))
	}

	return
}

func (c *Client) DeleteByQuery(index string, query elastic.Query) (err error) {
	_, err = elastic.NewDeleteByQueryService(c.Client).
		Index(index).
		ProceedOnVersionConflict().
		Refresh("true").
		Query(query).
		Do(context.Background())
	if nil != err {
		return
	}

	return
}

func (c *Client) DeleteByEqFields(index string, cond interface{}, fields ...string) (err error) {
	boolQ := elastic.NewBoolQuery()

	for _, field := range fields {
		var val interface{}
		if val, err = c.getFieldVal(field, cond); nil != err {
			return
		}
		boolQ.Must(elastic.NewMatchQuery(field, val))
	}

	err = c.DeleteByQuery(index, boolQ)

	return
}
