package elasticsearch

type panguConfig struct {
	// ES配置
	ElasticSearch config `json:"elasticsearch" yaml:"elasticsearch" validate:"required"`
}
