package endpoints

import (
	"context"
	"searchservice/service"

	"github.com/go-kit/kit/endpoint"
)

func makeSearchEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(SearchRequest)
		result, err := s.Search(ctx, req.Query, req.PageSize, req.PageNumber)
		response := SearchResponse{Posts: result, Success: 1, Message: "Posts found!"}
		if err != nil {
			response.Success = 0
			response.Message = err.Error()
		}
		return response, nil
	}
}
