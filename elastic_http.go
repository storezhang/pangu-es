package elasticsearch

type elasticHttp struct {
	TimeOut               int64 `default:"10" yaml:"timeOut"`
	KeepAlive             int64 `default:"30" yaml:"keepAlive"`
	MaxIdleConns          int64 `default:"300" yaml:"maxIdleConns"`
	MaxIdleConnsPerHost   int64 `default:"150" yaml:"maxIdleConnsPerHost"`
	MaxConnsPerHost       int64 `default:"300" yaml:"maxConnsPerHost"`
	IdleConnTimeout       int64 `default:"60" yaml:"idleConnTimeout"`
	TLSHandshakeTimeout   int64 `default:"60" yaml:"tlsHandshakeTimeout"`
	ExpectContinueTimeout int64 `default:"1" yaml:"expectContinueTimeout"`
	// 强制尝试Http2
	ForceAttemptHTTP2 bool `default:"true" yaml:"forceAttemptHTTP2"`
}
