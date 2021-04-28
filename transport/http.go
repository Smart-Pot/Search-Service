package transport

import (
	"context"
	"encoding/json"
	"net/http"
	"searchservice/endpoints"
	"strconv"

	pkghttp "github.com/Smart-Pot/pkg/common/http"
	"github.com/Smart-Pot/pkg/common/perrors"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/transport"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
)

var (
	ParamMissing = perrors.New("Parametre is missing.", http.StatusBadRequest)
)

const userIDTag = "x-user-id"

func MakeHTTPHandlers(e endpoints.Endpoints, logger log.Logger) http.Handler {
	r := mux.NewRouter().PathPrefix("/search").Subrouter()

	options := []httptransport.ServerOption{
		httptransport.ServerErrorHandler(transport.NewLogErrorHandler(logger)),
		httptransport.ServerErrorEncoder(encodeError),
	}

	r.Methods("GET").Path("/find").Handler(httptransport.NewServer(
		e.Search,
		decodeSearchHTTPRequest,
		encodeHTTPResponse,
		options...,
	))

	return pkghttp.EnableCORS(r)
}

func encodeHTTPResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	return json.NewEncoder(w).Encode(response)
}

func decodeSearchHTTPRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req endpoints.SearchRequest

	query, ok := r.URL.Query()["query"]
	if !ok {
		return nil, ParamMissing
	}

	pageSize, ok := r.URL.Query()["pageSize"]
	if !ok {
		return nil, ParamMissing
	}

	size, err := strconv.Atoi(pageSize[0])

	if err != nil {
		return nil, err
	}

	pageNumber, ok := r.URL.Query()["pageNumber"]
	if !ok {
		return nil, ParamMissing
	}
	number, err := strconv.Atoi(pageNumber[0])

	if err != nil {
		return nil, err
	}
	req.Query = query[0]
	req.PageSize = size
	req.PageNumber = number

	return req, nil

}

func encodeError(_ context.Context, err error, w http.ResponseWriter) {
	if err == nil {
		panic("encodeError with nil error")
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusBadRequest)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"error": err.Error(),
	})
}
