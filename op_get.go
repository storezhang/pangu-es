package elasticsearch

import (
	`context`
	`encoding/json`
	`fmt`
	`reflect`

	`github.com/olivere/elastic/v7`
)

func (c *Client) GetByDocId(index string, docId string, result interface{}) (exists bool, err error) {
	var rsp *elastic.GetResult
	if rsp, err = c.Get().Index(index).Id(docId).Do(context.Background()); nil != err {
		return
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

func (c *Client) GetByFields(index string, query *FieldsQuery) (results []interface{}, err error) {
	sourceContext := c.Search().
		Index(index).
		SortBy(query.Sorters...)
	if query.Size > 0 {
		sourceContext.Size(query.Size)
	} else {
		sourceContext.Size(300)
	}
	if query.From > 0 {
		sourceContext.From(query.From)
	}

	boolQ := elastic.NewBoolQuery()
	for _, field := range query.FieldsEq {
		var val interface{}
		if val, err = c.getFieldVal(field, query.Condition); nil != err {
			return
		}
		boolQ.Must(elastic.NewMatchQuery(field, val))
	}
	for _, field := range query.FieldsLike {
		var val interface{}
		if val, err = c.getFieldVal(field, query.Condition); nil != err {
			return
		}
		bq := elastic.NewBoolQuery()
		q := elastic.NewQueryStringQuery(fmt.Sprintf("*%s*", val))
		q.AllowLeadingWildcard(true)
		q.DefaultField(field)
		q.AnalyzeWildcard(true)
		bq.Should(q)
		bq.MinimumNumberShouldMatch(1)

		boolQ.Must(bq)
	}

	sourceContext.Query(boolQ)

	var res *elastic.SearchResult
	if res, err = sourceContext.Do(context.Background()); nil != err {
		return
	}

	results = make([]interface{}, 0)
	for _, item := range res.Each(query.ResultType) {
		results = append(results, item)
	}

	return
}

func (c *Client) GetsByQuery(
	index string, query elastic.Query, resultType reflect.Type, sorters ...elastic.Sorter,
) (results []interface{}, err error) {
	var res *elastic.SearchResult
	if res, err = c.Search(index).
		Query(query).
		SortBy(sorters...).
		Do(context.Background()); nil != err {
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
	if res, err = sourceContext.Do(context.Background()); nil != err {
		return
	}

	total = res.TotalHits()
	results = make([]interface{}, 0)
	for _, item := range res.Each(paging.ResultType) {
		results = append(results, item)
	}
	return
}

func (c *Client) Count(index string, query *CountQuery) (total int64, err error) {
	sourceContext := c.Search().
		Index(index).
		Size(0)

	boolQ := elastic.NewBoolQuery()
	for _, field := range query.FieldsEq {
		var val interface{}
		if val, err = c.getFieldVal(field, query.Condition); nil != err {
			return
		}
		boolQ.Must(elastic.NewMatchQuery(field, val))
	}
	for _, field := range query.FieldsLike {
		var val interface{}
		if val, err = c.getFieldVal(field, query.Condition); nil != err {
			return
		}
		bq := elastic.NewBoolQuery()
		q := elastic.NewQueryStringQuery(fmt.Sprintf("*%s*", val))
		q.AllowLeadingWildcard(true)
		q.DefaultField(field)
		q.AnalyzeWildcard(true)
		bq.Should(q)
		bq.MinimumNumberShouldMatch(1)

		boolQ.Must(bq)
	}
	sourceContext.Query(boolQ)

	var res *elastic.SearchResult
	if res, err = sourceContext.Do(context.Background()); nil != err {
		return
	}
	total = res.TotalHits()

	return
}
