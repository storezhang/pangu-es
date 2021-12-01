package elasticsearch

import (
	"github.com/olivere/elastic/v7"
)

type Client struct {
	*elastic.Client
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
