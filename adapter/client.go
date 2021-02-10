package adapter

import (
	"context"
	"encoding/json"
	"errors"
	"searchservice/data"

	"github.com/olivere/elastic/v7"
)

const URL = "localhost:9200/smart-pot_opt/posts/_search?pretty"

func GetSearchResults(ctx context.Context, query string, pageSize, pageNumber int) ([]*data.Post, error) {
	client, err := elastic.NewClient()
	if err != nil {
		return nil, err
	}

	exist, err := client.IndexExists("smart-pot_opt").Do(ctx)
	if err != nil {
		return nil, err
	}
	if !exist {
		return nil, errors.New("index not exist")
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
