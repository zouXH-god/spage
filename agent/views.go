package agent

import (
	"context"
	pb "github.com/LiteyukiStudio/spage/protos/result/protos/source"
	"google.golang.org/grpc"
)

type ServerVisit struct {
	pb.UnimplementedAgentServiceServer
}

// RegisterGrpcApps 注册grpc服务
func RegisterGrpcApps() (*grpc.Server, error) {
	// 创建gRPC服务器并添加拦截器
	s := grpc.NewServer(
		grpc.UnaryInterceptor(UnaryInterceptor),
	)
	pb.RegisterAgentServiceServer(s, &ServerVisit{})
	return s, nil
}

func (ServerVisit) CreateSite(ctx context.Context, request *pb.CreateSiteRequest) (*pb.CreateSiteResponse, error) {
	return &pb.CreateSiteResponse{
		Message: "ok",
	}, nil
}

func (ServerVisit) UpdateSite(context.Context, *pb.UpdateSiteRequest) (*pb.UpdateSiteResponse, error) {
	return &pb.UpdateSiteResponse{
		Message: "ok",
	}, nil
}

func (ServerVisit) DeleteSite(context.Context, *pb.DeleteSiteRequest) (*pb.DeleteSiteResponse, error) {
	return &pb.DeleteSiteResponse{
		Message: "ok",
	}, nil
}

func (ServerVisit) UploadRelease(grpc.ClientStreamingServer[pb.UploadReleaseRequest, pb.UploadReleaseResponse]) error {
	return nil
}
