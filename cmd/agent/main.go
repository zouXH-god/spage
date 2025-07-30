package main

import (
	"fmt"
	"github.com/LiteyukiStudio/spage/agent"
	"github.com/LiteyukiStudio/spage/config"
	"github.com/sirupsen/logrus"
	"net"
)

func main() {
	// 解析参数
	config.Cmd.ParseArgs()
	config.AgentConfig.Server.Token = config.Cmd.GetString("server.token", "")
	config.AgentConfig.Server.Host = config.Cmd.GetString("server.host", "127.0.0.1")
	config.AgentConfig.Server.Port = config.Cmd.GetString("server.port", "9526")
	config.AgentConfig.Service.Host = config.Cmd.GetString("service.host", "")
	config.AgentConfig.Service.Port = config.Cmd.GetString("service.port", "9527")
	config.AgentConfig.Service.Static = config.Cmd.GetString("service.static", "./static")
	config.AgentConfig.Caddy.Point = config.Cmd.GetString("caddy.point", "")
	err := config.AgentConfig.InitAgentConfig()
	if err != nil {
		logrus.Panicf("Failed to initialize agent configuration: %v", err)
		return
	}
	// 注册 grpc 服务
	grpcServer, err := agent.RegisterGrpcApps()
	if err != nil {
		logrus.Panicf("Failed to register grpc apps: %v", err)
		return
	}

	// 创建 TCP 监听器
	servicePoint := fmt.Sprintf("%s:%s", config.AgentConfig.Service.Host, config.AgentConfig.Service.Port)
	lis, err := net.Listen("tcp", servicePoint)
	if err != nil {
		logrus.Fatalf("Failed to listen on %s: %v", servicePoint, err)
	}

	// 启动 gRPC 服务
	logrus.Printf("gRPC server starting on %s\n", servicePoint)
	if err := grpcServer.Serve(lis); err != nil {
		logrus.Fatalf("gRPC server failed: %v", err)
	}

}
