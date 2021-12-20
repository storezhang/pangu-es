package elasticsearch

import (
	`reflect`

	`github.com/olivere/elastic/v7`
)

// FieldsQuery 查询条件
type FieldsQuery struct {
	elastic.Query

	// 查询结果包含的字段
	Fields []string
	// 查询大小
	Size int
	// 插叙结果类型
	// 结果类型
	ResultType reflect.Type
}
