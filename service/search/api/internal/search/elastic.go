package search

import (
	"context"
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/zeromicro/go-zero/core/logx"
	"log"
)

/*
* @Author: chuang
* @Date:   2023/2/2 21:08
 */

type ElasticClient struct {
	Client *elasticsearch.Client
	Ctx    context.Context
	Index  string
}

func NewElasticClient(config elasticsearch.Config, index string) *ElasticClient {
	client, err := elasticsearch.NewClient(config)
	if err != nil {
		logx.Errorf("ElasticSearch启动异常, cause: %s", err.Error())
		return nil
	}
	elastic := &ElasticClient{
		Client: client,
		Ctx:    context.Background(),
		Index:  index,
	}
	elastic.getClusterInfo()
	return elastic
}

//Get cluster info
func (elastic *ElasticClient) getClusterInfo() {
	// Check response status
	res, err := elastic.Client.Info()
	if err != nil {
		log.Fatalf("Error getting response: %s", err)
	}
	defer res.Body.Close()
	// Check response status
	if res.IsError() {
		log.Fatalf("Error: %s", res.String())
	}
	// Print client and server version numbers.
	log.Printf("Client: %s", elasticsearch.Version)
	log.Println("-----ElasticSearch start success-----")
}
