package elasticsearch

import (
	`github.com/olivere/elastic/v7`
)

// Aggregation 聚合操作信息
type Aggregation struct {
	elastic.Aggregation

	BoolQuery *elastic.BoolQuery
	Field     string
}
