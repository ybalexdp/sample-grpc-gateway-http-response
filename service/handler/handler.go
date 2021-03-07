package handler

import (
	"context"
	"errors"

	app "sample-grpc-gateway-http-response/service/application"
	"sample-grpc-gateway-http-response/service/domain"
	"sample-grpc-gateway-http-response/service/pb"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Handler struct {
	Application app.ApplicationInterface
}

func (h *Handler) GetSample(ctx context.Context, req *pb.GetSampleRequest) (*pb.GetSampleResponse, error) {
	data, err := h.Application.GetSample(ctx, &app.GetSampleParams{
		Id: req.Id,
	})
	if err != nil {
		if errors.Is(err, domain.Errors.SampleNotFound) {
			return nil, status.New(codes.NotFound, err.Error()).Err()
		}
		return nil, status.New(codes.Internal, err.Error()).Err()
	}
	return data, nil
}
