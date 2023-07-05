package search

import (
	"encoding/json"
	"log"

	"github.com/elastic/go-elasticsearch/esapi"
)

type elasticStore interface {
	GetContentByKeyword(searchIndex string, jsonQuery []byte) (*esapi.Response, error)
}

type SearchRequest struct {
	KeyWord string `form:"keyword" binding:"required"`
}

type SearchResponse struct {
	Url         string `json:"url"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

func setQuery(req SearchRequest) ([]byte, error) {

	query := map[string]interface{}{
		"query": map[string]interface{}{
			"multi_match": map[string]interface{}{
				"query":  req.KeyWord,
				"fields": []string{"keyword", "title", "description"},
			},
		},
	}

	jsonQuery, err := json.Marshal(query)
	if err != nil {
		log.Printf("Can not marshal query: %s", err)
		return []byte{}, err
	}

	return jsonQuery, nil
}

func convertResponse(res *esapi.Response) ([]SearchResponse, error) {
	searchResults := make([]SearchResponse, 0)

	var esapiResponse struct {
		Hits struct {
			Hits []struct {
				Source SearchResponse `json:"_source"`
			} `json:"hits"`
		} `json:"hits"`
	}
	if err := json.NewDecoder(res.Body).Decode(&esapiResponse); err != nil {
		log.Fatalf("Error Decoder response: %s", err)
	}

	for _, hit := range esapiResponse.Hits.Hits {
		result := SearchResponse{
			Url:         hit.Source.Url,
			Title:       hit.Source.Title,
			Description: hit.Source.Description,
		}

		searchResults = append(searchResults, result)
	}

	return searchResults, nil
}
