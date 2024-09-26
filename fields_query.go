package elasticsearch

import (
	`reflect`

	`github.com/olivere/elastic/v7`
)

type FieldsQuery struct {
	// 搜索条件
	Condition interface{}
	// 匹配字段
	FieldsEq []string
	// 模糊字段
	FieldsLike []string
	// 开始位置
	From int
	// 查询数量
	Size int
	// 排序字段
	Sorters []elastic.Sorter
	// 结果类型
	ResultType reflect.Type
}
