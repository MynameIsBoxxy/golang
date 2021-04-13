package transports

import (
	"context"
	"gokit/test/endpoints"
	"gokit/test/pb"

	"github.com/go-kit/kit/log"
	gt "github.com/go-kit/kit/transport/grpc"
)

type gRPCServer struct {
	validate gt.Handler
	fix      gt.Handler
	pb.UnimplementedTestServiceServer
}

func NewGRPCServer(endpoints endpoints.Endpoints, logger log.Logger) pb.TestServiceServer {
	return &gRPCServer{
		validate: gt.NewServer(
			endpoints.Validate,
			decodeValidateReq,
			encodeValidateResp,
		),
		fix: gt.NewServer(
			endpoints.Fix,
			decodeFixReq,
			encodeFixResp,
		),
	}
}

func (s *gRPCServer) Validate(ctx context.Context, req *pb.Input) (*pb.Output, error) {
	_, resp, err := s.validate.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}

	return resp.(*pb.Output), nil
}

func (s *gRPCServer) Fix(ctx context.Context, req *pb.Input) (*pb.Output, error) {
	_, resp, err := s.fix.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}

	return resp.(*pb.Output), nil
}

func decodeValidateReq(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.Input)
	return endpoints.ValidateReq{Input: req.Str}, nil
}

func encodeValidateResp(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(endpoints.ValidateResp)
	return &pb.Output{Str: resp.Output}, nil
}

func decodeFixReq(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.Input)
	return endpoints.FixReq{Input: req.Str}, nil
}

func encodeFixResp(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(endpoints.FixResp)
	return &pb.Output{Str: resp.Output}, nil
}
