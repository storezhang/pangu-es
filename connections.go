package elasticsearch

import (
	`time`
)

type connections struct {
	// 单个主机最大连接数
	MaxPerHost int `default:"300" yaml:"maxPerHost"`
	// ExpectContinueTimeout
	Wait time.Duration `default:"1s" yaml:"wait"`
	// 强制尝试Http2
	Http2 bool `default:"true" yaml:"http2"`
	// 空闲连接配置
	Idle *idle `yaml:"idle"`
	// TLS 安全传输配置
	TLS *tls `yaml:"tls"`
}
