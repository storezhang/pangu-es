package elasticsearch

import (
	`time`
)

// tls 安全传输配置
type tls struct {
	// 握手超时时间
	Timeout time.Duration `default:"60s" yaml:"timeout"`
}
