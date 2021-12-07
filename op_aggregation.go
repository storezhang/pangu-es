package elasticsearch

import (
	`context`
	`encoding/json`
	`net/http`

	`github.com/olivere/elastic/v7`
)

func (c *Client) AggregationMetrics(index string, agg *Aggregation) (result elastic.Aggregations, err error) {
	var _result *elastic.SearchResult

	result = make(map[string]json.RawMessage)
	if _result, err = c.Search().Index(index).Query(agg.BoolQuery).Size(0).Aggregation(agg.Field, agg.Aggregation).
		Do(context.Background()); nil != err {
		if elasticErr, ok := err.(*elastic.Error); ok {
			if http.StatusNotFound == elasticErr.Status {
				err = nil
			}
		}
	}
	if nil == _result {
		return
	}
	result = _result.Aggregations

	return
}
