package elasticsearch

type elasticHttp struct {
	Dial *Dial `yaml:"dial"`
	// Connections 连接配置
	Connections *Connections `yaml:"connections"`
}
