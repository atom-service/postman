package main

import (
	"net"

	"github.com/grpcbrick/sender/dao"
	"github.com/grpcbrick/sender/provider"
	"github.com/grpcbrick/sender/standard"
	"github.com/yinxulai/goutils/config"
	"github.com/yinxulai/goutils/grpc/interceptor"
	"github.com/yinxulai/goutils/sqldb"
	"google.golang.org/grpc"
)

func init() {
	config.SetStandard("rpc_port", ":3000", true, "RPC 服务监听的端口")
	config.SetStandard("mysql_url", "", true, "RPC 使用的 MYSQL 数据库配置")
	config.SetStandard("encrypt_password", "encrypt_password", false, "作为一些数据加密的密钥")
	config.LoadFlag()
}

func main() {
	var err error
	sqldb.Init("mysql", config.MustGet("mysql_url"))
	dao.MustInitTables()
	lis, err := net.Listen("tcp", config.MustGet("rpc_port"))
	if err != nil {
		panic(err)
	}
	grpcServer := grpc.NewServer(interceptor.NewCalllogs()...)
	standard.RegisterSenderServer(grpcServer, provider.NewService())
	panic(grpcServer.Serve(lis))
}
