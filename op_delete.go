package elasticsearch

import (
	`context`
	`net/http`

	`github.com/olivere/elastic/v7`
)

func (c *Client) DeleteByDocId(index string, docId string) (err error) {
	var resp *elastic.DeleteResponse
	resp, err = c.Delete().Index(index).Id(docId).Refresh(`true`).Do(context.Background())
	if nil != resp && http.StatusNotFound == resp.Status {
		err = nil
	}

	if nil != err {
		if elasticErr, ok := err.(*elastic.Error); ok {
			if http.StatusNotFound == elasticErr.Status {
				err = nil
			}
		}
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
		if elasticErr, ok := err.(*elastic.Error); ok {
			if http.StatusNotFound == elasticErr.Status {
				err = nil
			}
		}
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
