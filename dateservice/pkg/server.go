package pkg

import (
	"net/http"

	httptransport "github.com/go-kit/kit/transport/http"
)

func NewHTTPServer(endpoints Endpoints) *http.ServeMux {
	router := http.NewServeMux()

	// создадим простой middleware
	// он будет устанавливать для всех запросов,
	// зарегистрированных через него, тип ответа "application/json"
	handle := func(pattern string, handler http.Handler) {
		router.Handle(pattern, http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
			writer.Header().Add("Content-Type", "application/json; charset=utf-8")
			handler.ServeHTTP(writer, request)
		}))
	}

	handle(
		"/status",
		httptransport.NewServer(
			endpoints.StatusEndpoint,
			decodeStatusRequest,
			encodeResponse,
		),
	)

	handle(
		"/get",
		httptransport.NewServer(
			endpoints.GetEndpoint,
			decodeGetRequest,
			encodeResponse,
		),
	)

	handle(
		"/validate",
		httptransport.NewServer(
			endpoints.ValidateEndpoint,
			decodeValidateRequest,
			encodeResponse,
		),
	)

	return router
}
