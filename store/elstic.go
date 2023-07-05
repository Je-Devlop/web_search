package store

import (
	"log"

	"github.com/elastic/go-elasticsearch/v8"
)

type ElasticClient struct {
	*elasticsearch.Client
}

func NewElasticClient(es_end_point string) (*ElasticClient, error) {
	cfg := elasticsearch.Config{
		Addresses: []string{
			es_end_point,
		},
	}

	//Connect to Elasticsearch
	es, err := elasticsearch.NewClient(cfg)
	if err != nil {
		log.Fatalf("Error creating the client: %s", err)
		return nil, err
	}

	res, err := es.Info()
	if err != nil {
		log.Fatalf("Error getting response: %s", err)
		return nil, err
	}

	defer res.Body.Close()
	log.Println(res)

	return &ElasticClient{es}, nil
}
