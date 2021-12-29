package elasticsearch

type CountQuery struct {
	// 搜索条件
	Condition interface{}
	// 匹配字段
	FieldsEq []string
	// 模糊字段
	FieldsLike []string
}
