package pkg

import (
	"context"
	"encoding/json"
	"net/http"
)

// Сначала описываем "модели" запросов и ответов нашего сервиса

type getRequest struct{}

type getResponse struct {
	Date string `json:"date"`
	Err  error  `json:"err,omitempty"`
}

type validateRequest struct {
	Date string `json:"date"`
}

type validateResponse struct {
	Valid bool  `json:"valid"`
	Err   error `json:"err,omitempty"`
}

type statusRequest struct{}

type statusResponse struct {
	Status string `json:"status"`
}

// Затем описываем "декодеры" для входящих запросов

func decodeGetRequest(
	_ context.Context,
	_ *http.Request,
) (interface{}, error) {
	return getRequest{}, nil
}

func decodeValidateRequest(
	_ context.Context,
	r *http.Request,
) (interface{}, error) {
	var req validateRequest

	err := json.NewDecoder(r.Body).Decode(&req)

	if err != nil {
		return nil, err
	}

	return req, nil
}

func decodeStatusRequest(
	_ context.Context,
	_ *http.Request,
) (interface{}, error) {
	return statusRequest{}, nil
}

func encodeResponse(
	_ context.Context,
	w http.ResponseWriter,
	response interface{},
) error {
	return json.NewEncoder(w).Encode(response)
}
