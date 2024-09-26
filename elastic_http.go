package elasticsearch

type elasticHttp struct {
	Dial *dial `yaml:"dial"`
	// Connections 连接配置
	Connections *connections `yaml:"connections"`
}
