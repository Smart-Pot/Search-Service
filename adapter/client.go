package adapter

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"searchservice/data"

	elastic "github.com/olivere/elastic/v7"
)

const url = "localhost:9200/smart-pot_opt/posts/_search?pretty"

// GetSearchResults gets search results from Elasticsearch client.
func GetSearchResults(ctx context.Context, query string, pageSize, pageNumber int) ([]*data.Post, error) {
	client, err := elastic.NewClient()
	if err != nil {
		return nil, err
	}

	err = checkIndexExist(ctx, client, "smart-pot_opt")

	if err != nil {
		return nil, err
	}

	termQuery := elastic.NewTermQuery("plant", query)
	searchReasult, err := client.Search().
		Index("smart-pot_opt").
		Query(termQuery).
		From(pageSize * (pageNumber - 1)).Size(pageSize).
		Pretty(true).
		Do(ctx)

	if err != nil {
		return nil, err
	}

	var posts []*data.Post

	if searchReasult.Hits.TotalHits.Value > 0 {
		for _, hit := range searchReasult.Hits.Hits {
			var p *data.Post
			err := json.Unmarshal(hit.Source, &p)
			if err != nil {
				return nil, errors.New("deserialize search results failed")
			}

			posts = append(posts, p)

		}
	} else {
		return nil, errors.New("found no posts")
	}

	return posts, err

}

func checkIndexExist(ctx context.Context, client *elastic.Client, indexName string) error {
	exist, err := client.IndexExists(indexName).Do(ctx)
	if err != nil {
		return err
	}
	if !exist {
		return createIndex(ctx, client, indexName)
	}

	return nil
}

func createIndex(ctx context.Context, client *elastic.Client, indexName string) error {
	_, err := client.CreateIndex(indexName).BodyString(data.PostMapping).Do(ctx)
	if err != nil {
		panic(fmt.Sprintf("index not created, err: %s", err))
	}
	return nil
}
