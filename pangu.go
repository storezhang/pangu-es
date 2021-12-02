package elasticsearch

import (
	`github.com/pangum/pangu`
)

func init() {
	if err := pangu.New().Provides(newElasticClient); nil != err {
		panic(err)
	}
}
