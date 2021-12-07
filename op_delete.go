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
