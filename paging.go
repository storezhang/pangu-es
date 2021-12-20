package elasticsearch

import (
	`reflect`

	`github.com/olivere/elastic/v7`
)

// Paging 分页查询信息
type Paging struct {
	// 查询条件
	elastic.Query

	// 查询结果包含的字段
	Fields []string
	// 开始位置
	From int
	// 查询数量
	Size int
	// 排序字段
	Sorters []elastic.Sorter
	// 结果类型
	ResultType reflect.Type
}
