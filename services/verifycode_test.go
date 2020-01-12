package services

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/grpcbrick/sender/dao"
	"github.com/grpcbrick/sender/standard"
	"github.com/yinxulai/goutils/config"
	"github.com/yinxulai/goutils/sqldb"
)

func TestMain(m *testing.M) {
	fmt.Println("准备测试环境") // 日志
	config.Set("email_host", "encrypt_password")
	config.Set("email_port", "encrypt_password")
	config.Set("email_account", "encrypt_password")
	config.Set("email_password", "encrypt_password")
	sqldb.Init("mysql", "root:root@tcp(localhost:3306)/default?charset=utf8mb4&parseTime=true") // 测试数据库
	dao.InitTables()                                                                            // 初始化测试数据库                                                                        // 预插入一条用户数据
	fmt.Println("开始执行测试")                                                                       // 日志
	exitCode := m.Run()                                                                         // 执行测试
	fmt.Println("测试执行完成,清理测试数据")                                                                // 日志
	dao.TruncateTables()                                                                        // 重置测试数据库
	os.Exit(exitCode)                                                                           // 推出
}

func TestService_SendVerifyCodeByEmail(t *testing.T) {
	srv := New()
	tests := []struct {
		name      string
		args      *standard.SendVerifyCodeByEmailRequest
		wantState standard.State
		wantErr   bool
	}{
		{"正常邮件测试", &standard.SendVerifyCodeByEmailRequest{
			EmailAddress: "test@yinxilai.com", Operation: "test1", ValidityPeriod: 10,
		}, standard.State_SUCCESS, false},
		{"正常邮件测试", &standard.SendVerifyCodeByEmailRequest{
			EmailAddress: "test@yinxilai.com", Operation: "test2", ValidityPeriod: 10,
		}, standard.State_SUCCESS, false},
		{"正常邮件测试", &standard.SendVerifyCodeByEmailRequest{
			EmailAddress: "test@yinxilai.com", Operation: "test3", ValidityPeriod: 10,
		}, standard.State_SUCCESS, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResp, err := srv.SendVerifyCodeByEmail(context.Background(), tt.args)
			if (err != nil) != tt.wantErr {
				t.Errorf("Service.SendVerifyCodeByEmail() 邮件发送错误,得到：%v, 期望得到：%v", err, tt.wantErr)
				return
			}

			if gotResp.State.String() != tt.wantState.String() {
				t.Errorf("Service.SendVerifyCodeByEmail() 返回状态不符合预期，得到：%v, 期望得到：%v", gotResp, tt.wantState)
				return
			}

			// 如果目标是成功、就去测一下查询接口
			if gotResp.State == standard.State_SUCCESS {

			}
		})
	}
}
