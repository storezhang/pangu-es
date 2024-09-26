package elasticsearch

type config struct {
	// Local 是否本地
	Local bool `default:"false" yaml:"local"`

	// Address 地址
	Address string `yaml:"address" validate:"required"`
	// Username 用户名
	Username string `yaml:"username"`
	// Password 密码
	Password string `yaml:"password"`

	// 是否启用Sniff
	Sniff bool `default:"false" yaml:"sniff"`

	// Http httpClient配置
	Http *elasticHttp `yaml:"http"`
}
