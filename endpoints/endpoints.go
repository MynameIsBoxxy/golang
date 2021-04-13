package endpoints

import (
	"context"
	"gokit/test/service"
	"log"

	"github.com/go-kit/kit/endpoint"
)

type Endpoints struct {
	Validate endpoint.Endpoint
	Fix      endpoint.Endpoint
}

type ValidateReq struct {
	Input string
}

type ValidateResp struct {
	Output string
}

type FixReq struct {
	Input string
}

type FixResp struct {
	Output string
}

func MakeEndpoints(s service.Service) Endpoints {
	return Endpoints{
		Validate: makeValidateEndpoints(s),
		Fix:      makeFixEndpoints(s),
	}
}

func makeValidateEndpoints(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(ValidateReq)
		result, _ := s.Validate(ctx, req.Input)
		return ValidateResp{Output: result}, nil
	}
}

func makeFixEndpoints(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(FixReq)
		result, _ := s.Fix(ctx, req.Input)
		return FixResp{Output: result}, nil
	}
}

func loggingMiddleware(logger log.Logger) Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (interface{}, error) {
			logger.Log("msg", "calling endpoint")
			defer logger.Log("msg", "called endpoint")
			return next(ctx, request)
		}
	}
}
