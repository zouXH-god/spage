package agent

import (
	"context"
	"fmt"
	"github.com/LiteyukiStudio/spage/pkg/config"
	pb "github.com/LiteyukiStudio/spage/protos/result/protos/source"
	"google.golang.org/grpc"
	"os"
	"path/filepath"
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

/*
站点保存路径规则：
站点保存路径：<static_path>/<OwnerName>/<ProjectName>/<Name>
*/
func getSitePath(ownerName, projectName, siteName string) (string, error) {
	sitePath := filepath.Join(config.AgentConfig.Service.Static, ownerName, projectName, siteName)
	// 判断路径是否存在
	if _, err := os.Stat(sitePath); os.IsNotExist(err) {
		// 创建目录
		if err = os.MkdirAll(sitePath, os.ModePerm); err != nil {
			return "", err
		}
	}
	return sitePath, nil
}

func (ServerVisit) CreateSite(ctx context.Context, request *pb.CreateSiteRequest) (response *pb.CreateSiteResponse, err error) {
	sitePath, err := getSitePath(request.OwnerName, request.ProjectName, request.Name)
	if err != nil {
		response.Success = false
		response.Message = err.Error()
		return
	}
	response.Success = true
	response.Message = fmt.Sprintf("Path Created Successfully 【%s】", sitePath)
	return
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
