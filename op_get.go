package elasticsearch

import (
	`context`
	`encoding/json`
	`net/http`
	`reflect`

	`github.com/olivere/elastic/v7`
)

func (c *Client) GetByDocId(index string, docId string, result interface{}) (exists bool, err error) {
	var rsp *elastic.GetResult
	rsp, err = c.Get().Index(index).Id(docId).Do(context.Background())
	if nil != err {
		if elasticErr, ok := err.(*elastic.Error); ok {
			if http.StatusNotFound == elasticErr.Status {
				err = nil
			}
		}
	}
	if nil == rsp || !rsp.Found || nil == rsp.Source {
		return
	}
	if err = json.Unmarshal(rsp.Source, result); nil != err {
		return
	}

	exists = true

	return
}

func (c *Client) GetsByQuery(index string, query elastic.Query, resultType reflect.Type) (results []interface{}, err error) {
	var res *elastic.SearchResult
	if res, err = c.Search(index).Query(query).Do(context.Background()); nil != err {
		return
	}

	results = make([]interface{}, 0, res.TotalHits())
	for _, item := range res.Each(resultType) {
		results = append(results, item)
	}

	return
}

func (c *Client) GetsByFields(index string, fieldsQuery *FieldsQuery) (results []interface{}, err error) {
	sourceContext := c.Search().
		Index(index).
		Query(fieldsQuery.Query).
		FetchSourceContext(elastic.NewFetchSourceContext(true).Include(fieldsQuery.Fields...))
	if fieldsQuery.Size > 0 {
		sourceContext.Size(fieldsQuery.Size)
	}

	var res *elastic.SearchResult
	res, err = sourceContext.Do(context.Background())
	if nil != err {
		if elasticErr, ok := err.(*elastic.Error); ok {
			if http.StatusNotFound == elasticErr.Status {
				err = nil
			}
		}
	}

	results = make([]interface{}, 0, res.TotalHits())
	for _, item := range res.Each(fieldsQuery.ResultType) {
		results = append(results, item)
	}
	return
}

func (c *Client) GetsByPaging(index string, page *Paging) (results []interface{}, total int64, err error) {
	var res *elastic.SearchResult
	res, err = c.Search().
		Index(index).
		Query(page.BoolQ).
		From(page.From).
		Size(page.Size).
		SortBy(page.Sorters...).
		Do(context.Background())

	if nil != err {
		if elasticErr, ok := err.(*elastic.Error); ok {
			if http.StatusNotFound == elasticErr.Status {
				err = nil
				results = make([]interface{}, 0, 0)

				return
			}
		}
	}

	if nil != err {
		return
	}

	total = res.TotalHits()
	results = make([]interface{}, 0, page.Size)
	for _, item := range res.Each(page.ResultType) {
		results = append(results, item)
	}

	return
}
