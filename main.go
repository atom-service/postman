package main

import (
	"net"

	"github.com/grpcbrick/sender/dao"
	"github.com/grpcbrick/sender/services"
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

	config.SetStandard("email_host", "", true, "邮箱服务的地址")
	config.SetStandard("email_port", "", true, "邮箱服务的端口")
	config.SetStandard("email_account", "", true, "邮箱服务的账号")
	config.SetStandard("email_password", "", true, "邮箱服务的密码")
	config.SetStandard("email_verify_code_template", "{{.Operation}}:{{.Code}}", true, "邮箱验证码模板")

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
	standard.RegisterSenderServer(grpcServer, services.New())
	panic(grpcServer.Serve(lis))
}
