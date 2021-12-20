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

func (c *Client) GetByEqFields(
	index string, cond interface{}, resultType reflect.Type, fields ...string,
) (results []interface{}, err error) {
	boolQ := elastic.NewBoolQuery()

	for _, field := range fields {
		var val interface{}
		if val, err = c.getFieldVal(field, cond); nil != err {
			return
		}
		boolQ.Must(elastic.NewMatchQuery(field, val))
	}

	results, err = c.GetsByQuery(index, boolQ, resultType)

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

func (c *Client) GetsByPaging(index string, paging *Paging) (results []interface{}, total int64, err error) {
	sourceContext := c.Search().
		Index(index).
		Query(paging.Query).
		SortBy(paging.Sorters...)

	if nil != paging.Fields && 0 < len(paging.Fields) {
		sourceContext.FetchSourceContext(elastic.NewFetchSourceContext(true).Include(paging.Fields...))
	}
	if paging.Size > 0 {
		sourceContext.Size(paging.Size)
	} else {
		sourceContext.Size(300)
	}
	if paging.From > 0 {
		sourceContext.From(paging.From)
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

	total = res.TotalHits()
	results = make([]interface{}, 0, res.TotalHits())
	for _, item := range res.Each(paging.ResultType) {
		results = append(results, item)
	}
	return
}
