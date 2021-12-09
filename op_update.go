package elasticsearch

import (
	`context`
	`net/http`

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
		Refresh("true").
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
		Refresh("true").
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
