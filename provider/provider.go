package provider

import (
	"context"
	"os"
	"strconv"
	"time"

	"github.com/yinxulai/goutils/random"
	tencentSMS "github.com/yinxulai/goutils/tencent/sms"
	"github.com/yinxulai/grpc-services/sender/standard"
)

// NewService NewService
func NewService() *Service {
	service := new(Service)
	return service
}

type SMSProvider interface {
	SendVerifyCode(code string) error
}

type CallProvider interface {
	SendVerifyCode(code string) error
}

type EmailProvider interface {
	SendVerifyCode(code string) error
}

// Service Service
type Service struct {
	SMS   []*SMSProvider
	Call  []*CallProvider
	Email []*EmailProvider
}

// CheckVerifyCode 验证 验证码
func (srv *Service) CheckVerifyCode(ctx context.Context, req *standard.CheckVerifyCodeRequest) (resp *standard.StateResponse, err error) {
	var err error
	var count uint64
	response := new(standard.StateResponse)

	err = checkVerifyCodeRecordBySymbolWithVerifyCodeStmt.GetContext(ctx, &count, req.Symbol, req.VerifyCode)
	if err != nil {
		response.State = standard.State_DB_OPERATION_FATLURE
		response.Message = err.Error()
		return response, nil
	}

	// 验证码不存在
	if count <= 0 {
		response.State = standard.State_FAILURE
		response.Message = "验证码错误"
		return response, nil
	}

	response.State = standard.State_SUCCESS
	response.Message = "验证码正确"
	return response, nil
}

// DiscardVerifyCode 报废一个验证码
func (srv *Service) DiscardVerifyCode(ctx context.Context, req *standard.DiscardVerifyCodeRequest) (*standard.StateResponse, error) {
	var err error
	// 创建结构体
	response := new(standard.StateResponse)
	_, err = discardVerifyCodeRecordBySymbolWithVerifyCodeStmt.ExecContext(ctx, req.Symbol, req.VerifyCode)
	if err != nil {
		response.State = standard.State_DB_OPERATION_FATLURE
		response.Message = err.Error()
		return response, nil
	}

	response.State = standard.State_SUCCESS
	response.Message = "完成"
	return response, nil
}

// SendVerifyCodeByEmail SendVerifyCodeByEmail
func (srv *Service) SendVerifyCodeByEmail(ctx context.Context, req *standard.SendVerifyCodeByEmailRequest) (*standard.StateResponse, error) {
}

// SendVerifyCodeByPhoneSms 发送短信验证码
func (srv *Service) SendVerifyCodeByPhoneSms(ctx context.Context, req *standard.SendSmsVerifyCodeRequest) (*standard.StateResponse, error) {
	var err error
	response := new(standard.StateResponse)
	verifyCode := strconv.FormatInt(random.Number(1000, 9999), 10)
	sendParamstplid, _ := strconv.ParseUint(os.Getenv("TENCENT_VERIFY_CODE_TEMPLATE_ID"), 10, 64)
	validityPeriod := time.Now().Add(time.Second * time.Duration(req.ValidityPeriod)).Format("2006-01-02 15:04:05")

	_, err = insertVerifyCodeRecordStmt.ExecContext(ctx, 0, req.Symbol, verifyCode, validityPeriod)
	if err != nil {
		response.State = standard.State_DB_OPERATION_FATLURE
		response.Message = err.Error()
		return response, nil
	}

	// 发送短信
	sendParams := new(tencentSMS.SendParams)
	sendParams.Tel.Mobile = req.PhoneNumber
	sendParams.Tel.Nationcode = req.CountryCode
	sendParams.Time = uint64(time.Now().Unix())
	sendParams.TplID = sendParamstplid
	sendParams.Params = []interface{}{req.Action, verifyCode, req.ValidityPeriod}
	sendParams.Sig = srv.tencentSMS.Sig(sendParams) // 签名
	_, err = srv.tencentSMS.Send(sendParams)        // 发送
	if err != nil {
		response.State = standard.State_FAILURE
		response.Message = "验证码发送失败"
		return response, nil
	}

	response.State = standard.State_SUCCESS
	response.Message = "发送成功"
	return response, nil
}

// SendVerifyCodeByCallPhone SendVerifyCodeByCallPhone
func (srv *Service) SendVerifyCodeByCallPhone(ctx context.Context, req *standard.SendVerifyCodeByCallPhoneRequest) (*standard.StateResponse, error) {
}
