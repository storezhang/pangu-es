package elasticsearch

import (
	`time`
)

type dial struct {
	Timeout   time.Duration `default:"10s" yaml:"timeout"`
	KeepAlive time.Duration `default:"30s" yaml:"keepalive"`
}
