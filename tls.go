package elasticsearch

import (
	`time`
)

// TLS 安全传输配置
type TLS struct {
	// 握手超时时间
	Timeout time.Duration `default:"60s" yaml:"timeout"`
}
