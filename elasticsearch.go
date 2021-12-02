package elasticsearch

import (
	`net`
	`net/http`

	`github.com/olivere/elastic/v7`
	`github.com/pangum/pangu`
)

func newElasticClient(config *pangu.Config) (client *Client, err error) {
	_panguConfig := new(panguConfig)
	if err = config.Load(_panguConfig); nil != err {
		return
	}

	_conf := _panguConfig.ElasticSearch
	// 配置http client
	httpClient := new(http.Client)
	httpClient.Transport = &http.Transport{
		DialContext: (&net.Dialer{
			Timeout:   _conf.Http.Dial.TimeOut,
			KeepAlive: _conf.Http.Dial.KeepAlive,
		}).DialContext,
		MaxIdleConns:          _conf.Http.Connections.Idle.Max,
		MaxIdleConnsPerHost:   _conf.Http.Connections.Idle.PerHost,
		MaxConnsPerHost:       _conf.Http.Connections.MaxPerHost,
		IdleConnTimeout:       _conf.Http.Connections.Idle.TimeOut,
		TLSHandshakeTimeout:   _conf.Http.Connections.TLS.TimeOut,
		ExpectContinueTimeout: _conf.Http.Connections.Wait,
		ForceAttemptHTTP2:     _conf.Http.Connections.Http2,
	}
	
	client = new(Client)
	if client.Client, err = elastic.NewClient(
		elastic.SetHttpClient(httpClient),
		elastic.SetBasicAuth(_conf.Username, _conf.Password),
		elastic.SetURL(_conf.Address),
		elastic.SetSniff(_conf.Sniff),
	); nil != err {
		return
	}

	return
}
