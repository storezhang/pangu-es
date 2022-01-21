package elasticsearch

import (
	`time`
)

// Idle 空闲连接配置
type Idle struct {
	// 最大空闲连接数
	Max int `default:"300" yaml:"max"`
	// 单个主机最大空闲连接数
	PerHost int `default:"150" yaml:"perHost"`
	// 连接超时时间
	Timeout time.Duration `default:"60s" yaml:"timeout"`
}
