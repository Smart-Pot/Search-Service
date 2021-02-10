package endpoints

import (
	"searchservice/data"
	"searchservice/service"

	"github.com/go-kit/kit/endpoint"
)

type Endpoints struct {
	Search endpoint.Endpoint
}

type SearchResponse struct {
	Posts   []*data.Post
	Success int32
	Message string
}

type SearchRequest struct {
	Query      string `json:"query"`
	PageSize   int    `json:"pageSize"`
	PageNumber int    `json:"pageNumber"`
}

func MakeEndpoints(s service.Service) Endpoints {
	return Endpoints{
		Search: makeSearchEndpoint(s),
	}
}
