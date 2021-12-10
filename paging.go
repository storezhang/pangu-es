package elasticsearch

import (
	`reflect`

	`github.com/olivere/elastic/v7`
)

// Paging 分页查询信息
type Paging struct {
	// 查询条件
	boolQ *elastic.BoolQuery
	// 开始位置
	from int
	// 查询数量
	size int
	// 排序字段
	sorters []elastic.Sorter
	// 结果类型
	resultType reflect.Type
}
