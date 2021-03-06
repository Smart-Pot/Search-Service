package transport

import (
	"context"
	"encoding/json"
	"net/http"
	"searchservice/endpoints"

	pkghttp "github.com/Smart-Pot/pkg/common/http"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/transport"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
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

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, err
	}

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
