package application

import (
	"context"

	domain "sample-grpc-gateway-http-response/service/domain"
	model "sample-grpc-gateway-http-response/service/model"
	pb "sample-grpc-gateway-http-response/service/pb"
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
		return nil, domain.Errors.SampleNotFound
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
