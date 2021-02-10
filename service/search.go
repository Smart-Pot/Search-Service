package service

import (
	"context"
	"searchservice/adapter"
	"searchservice/data"
	"time"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
)

type service struct {
	logger log.Logger
}

type Service interface {
	Search(ctx context.Context, query string, pageSize, pageNumber int) ([]*data.Post, error)
}

func NewService(logger log.Logger) Service {
	return &service{
		logger: logger,
	}
}

func (s service) Search(ctx context.Context, query string, pageSize, pageNumber int) (result []*data.Post, err error) {
	defer func(beginTime time.Time) {
		level.Info(s.logger).Log(
			"function", "Search",
			"param:query", query,
			"param:pageSize", pageSize,
			"param:pageNumber", pageNumber,
			"result", result,
			"took", time.Since(beginTime),
			"error", err,
		)
	}(time.Now())
	result, err = adapter.GetSearchResults(ctx, query, pageSize, pageNumber)
	return result, err
}
