package adapter

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"searchservice/data"

	elastic "github.com/olivere/elastic/v7"
)

// GetSearchResults gets search results from Elasticsearch client.
func GetSearchResults(ctx context.Context, query string, pageSize, pageNumber int) ([]*data.Post, error) {
	url := elastic.SetURL("http://elasticsearch:9200")
	cl, err := elastic.NewClient(url)
	if err != nil {
		return nil, err
	}

	termQuery := elastic.NewTermQuery("plant", query)
	searchReasult, err := cl.Search().
		Index("smart-pot_opt").
		Type("posts").
		Query(termQuery).
		From(pageSize * (pageNumber - 1)).Size(pageSize).
		Pretty(true).
		Do(ctx)

	if err != nil {
		return nil, err
	}

	var posts []*data.Post
	fmt.Println(searchReasult.Hits.TotalHits.Value)
	if searchReasult.Hits.TotalHits.Value > 0 {
		fmt.Println(len(searchReasult.Hits.Hits))
		fmt.Println(searchReasult.Hits.Hits)
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
