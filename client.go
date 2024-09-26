package elasticsearch

import (
	`github.com/olivere/elastic/v7`
)

type Client struct {
	*elastic.Client
}
