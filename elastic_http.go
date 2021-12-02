package elasticsearch

import (
	`time`
)

type elasticHttp struct {
	Dial struct {
		TimeOut   time.Duration `default:"10s" yaml:"timeout"`
		KeepAlive time.Duration `default:"30s" yaml:"keepalive"`
	} `yaml:"dial"`

	// Connections 连接配置
	Connections struct {
		// 单个主机最大连接数
		MaxPerHost int `default:"300" yaml:"maxPerHost"`
		// ExpectContinueTimeout
		Wait time.Duration `default:"1s" yaml:"wait"`
		// 强制尝试Http2
		Http2 bool `default:"true" yaml:"http2"`
		// 空闲连接配置
		Idle struct {
			// 最大空闲连接数
			Max int `default:"300" yaml:"max"`
			// 单个主机最大空闲连接数
			PerHost int `default:"150" yaml:"perHost"`
			// 连接超时时间
			TimeOut time.Duration `default:"60s" yaml:"timeout"`
		} `yaml:"idle"`

		// TLS 安全传输配置
		TLS struct {
			// 握手超时时间
			TimeOut time.Duration `default:"60s" yaml:"timeout"`
		} `yaml:"tls"`
	} `yaml:"connections"`
}
