package elasticsearch

import (
	`context`
	`encoding/json`
	`net/http`
	`reflect`

	`github.com/olivere/elastic/v7`
)

func (c *Client) Save(index string, docId string, bean interface{}) (err error) {
	_, err = c.Index().Index(index).Id(docId).BodyJson(bean).Refresh(`true`).Do(context.Background())

	return
}

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

func (c *Client) GetsByQuery(index string, _ elastic.BoolQuery, resultType reflect.Type) (results []interface{}, err error) {
	var res *elastic.SearchResult
	if res, err = c.Search(index).Do(context.Background()); nil != err {
		return
	}

	results = make([]interface{}, 0, res.TotalHits())
	for _, item := range res.Each(resultType) {
		results = append(results, item)
	}

	return
}
