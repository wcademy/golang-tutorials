package pkg

import (
	"context"
	"errors"

	"github.com/go-kit/kit/endpoint"
)

var (
	errUnexpected = errors.New("unexpected error")
)

func MakeStatusEndpoint(srv Service) endpoint.Endpoint {
	return func(
		ctx context.Context,
		request interface{},
	) (interface{}, error) {
		res, err := srv.Status(ctx)
		if err != nil {
			return statusResponse{res}, err
		}

		return statusResponse{res}, nil
	}
}

func MakeGetEndpoint(srv Service) endpoint.Endpoint {
	return func(
		ctx context.Context,
		request interface{},
	) (interface{}, error) {
		res, err := srv.Get(ctx)
		if err != nil {
			return getResponse{res, err}, err
		}

		return getResponse{Date: res}, nil
	}
}

func MakeValidateEndpoint(srv Service) endpoint.Endpoint {
	return func(
		ctx context.Context,
		request interface{},
	) (interface{}, error) {
		req, ok := request.(validateRequest)
		if !ok {
			return validateResponse{Err: errUnexpected}, errUnexpected
		}

		res, err := srv.Validate(ctx, req.Date)
		if err != nil {
			return validateResponse{res, err}, nil
		}

		return validateResponse{Valid: res}, nil
	}
}

type Endpoints struct {
	StatusEndpoint   endpoint.Endpoint
	GetEndpoint      endpoint.Endpoint
	ValidateEndpoint endpoint.Endpoint
}

func (e Endpoints) Get(ctx context.Context) (string, error) {
	req := getRequest{}

	resp, err := e.GetEndpoint(ctx, req)
	if err != nil {
		return "", err
	}

	getResp, ok := resp.(getResponse)
	if !ok {
		return "", errUnexpected
	}

	if getResp.Err != nil {
		return "", getResp.Err
	}

	return getResp.Date, nil
}

func (e Endpoints) Status(ctx context.Context) (string, error) {
	req := statusRequest{}

	resp, err := e.StatusEndpoint(ctx, req)
	if err != nil {
		return "", err
	}

	statusResp, ok := resp.(statusResponse)
	if !ok {
		return "", errUnexpected
	}

	return statusResp.Status, nil
}

func (e Endpoints) Validate(ctx context.Context, date string) (bool, error) {
	req := validateRequest{Date: date}

	resp, err := e.ValidateEndpoint(ctx, req)
	if err != nil {
		return false, err
	}

	validateResp, ok := resp.(validateResponse)
	if !ok {
		return false, errUnexpected
	}

	if validateResp.Err != nil {
		return false, validateResp.Err
	}

	return validateResp.Valid, nil
}
