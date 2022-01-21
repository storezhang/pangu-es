package elasticsearch

import (
	`github.com/olivere/elastic/v7`
)

type queryOpt func(query elastic.Query)
