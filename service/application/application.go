package application

import (
	"context"

	model "sample-grpc-gateway-http-response/service/model"
	pb "sample-grpc-gateway-http-response/service/pb"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type SampleService struct {
}

type GetSampleParams struct {
	Id int64
}

type ApplicationInterface interface {
	GetSample(context.Context, *GetSampleParams) (*pb.GetSampleResponse, error)
}

func (s *SampleService) GetSample(ctx context.Context, params *pb.GetSampleRequest) (*pb.GetSampleResponse, error) {
	data, err := getSampleData(ctx, params.Id)
	if err != nil {
		return nil, err
	}
	if data == nil {
		return nil, status.New(codes.NotFound, "sample not found").Err()
	}
	return &pb.GetSampleResponse{
		Id:   data.Id,
		Name: data.Name,
	}, nil
}

/**
 * 本来はDBからデータをとってくる処理などが必要
 * サンプル実装のためデータをnilで返す
 */
func getSampleData(ctx context.Context, id int64) (*model.Sample, error) {
	// do something
	return nil, nil
}
