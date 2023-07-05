package store

import (
	"bytes"
	"context"
	"log"

	"github.com/elastic/go-elasticsearch/esapi"
)

func (ec *ElasticClient) GetContentByKeyword(searchIndex string, jsonQuery []byte) (*esapi.Response, error) {
	// Set up the search request
	result := esapi.SearchRequest{
		Index: []string{searchIndex},
		Body:  bytes.NewReader(jsonQuery),
	}

	res, err := result.Do(context.Background(), ec)
	if err != nil {
		log.Printf("Error executing search request: %s", err)
		return nil, err
	}

	if res.IsError() {
		log.Printf("Search request failed: %s", res.String())
		return nil, err
	}

	return res, nil
}
