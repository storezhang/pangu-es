package elasticsearch

import (
	`context`
)

func (c *Client) Save(index string, docId string, bean interface{}) (err error) {
	_, err = c.Index().Index(index).Id(docId).BodyJson(bean).Refresh(`true`).Do(context.Background())

	return
}
