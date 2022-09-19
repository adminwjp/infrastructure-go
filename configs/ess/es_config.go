package ess

import "github.com/olivere/elastic/v7"
var EsConfigInstance=&EsConfig{}
type EsConfig struct {
	Client *elastic.Client
}
