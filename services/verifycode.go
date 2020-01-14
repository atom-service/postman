package services

import (
	"context"
	"math/rand"
	"time"

	"github.com/grpcbrick/sender/dao"
	"github.com/grpcbrick/sender/provider"
	"github.com/grpcbrick/sender/standard"
)

// SMSProvider 短信服务
type SMSProvider interface {
	Weights() int
	SendVerifyCode(operation, code string) error
}

// CallProvider 电话服务
type CallProvider interface {
	Weights() int
	SendVerifyCode(operation, code string) error
}

// EmailProvider 邮件服务
type EmailProvider interface {
	Weights() int
	SendVerifyCode(email, operation, code string) error
}

// New New
func New() *Service {
	service := new(Service)
	service.Email = provider.NewEmail()
	return service
}

// Service Service
type Service struct {
	SMS   SMSProvider   // 不直接使用,使用对应的 next 方法获取
	Call  CallProvider  // 不直接使用,使用对应的 next 方法获取
	Email EmailProvider // 不直接使用,使用对应的 next 方法获取
}

func (srv *Service) randomCode(l int) string {
	str := "0123456789"
	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < l; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}

func (srv *Service) nextSMSProvider() SMSProvider {
	return srv.SMS
}

func (srv *Service) nextCallProvider() CallProvider {
	return srv.Call
}

func (srv *Service) nextEmailProvider() EmailProvider {
	return srv.Email
}

// CheckVerifyCode CheckVerifyCode
func (srv *Service) CheckVerifyCode(ctx context.Context, req *standard.CheckVerifyCodeRequest) (resp *standard.CheckVerifyCodeeResponse, err error) {
	resp = new(standard.CheckVerifyCodeeResponse)
	record, err := dao.QueryVerifyCodeByKey(req.Key)
	if err != nil {
		resp.State = standard.State_DB_OPERATION_FATLURE
		resp.Message = err.Error()
	}

	if record.OutProtoStruct().Code != req.VerifyCode {
		resp.State = standard.State_FAILURE
		resp.Message = err.Error()
	}

	resp.State = standard.State_SUCCESS
	return resp, err
}

// DestroyVerifyCodeByKey DestroyVerifyCodeByKey
func (srv *Service) DestroyVerifyCodeByKey(ctx context.Context, req *standard.DestroyVerifyCodeByKeyRequest) (resp *standard.DestroyVerifyCodeByKeyResponse, err error) {
	resp = new(standard.DestroyVerifyCodeByKeyResponse)
	err = dao.UpdateVerifyCodeExpireTimeByKey(req.Key, time.Now())
	if err != nil {
		resp.State = standard.State_DB_OPERATION_FATLURE
		resp.Message = err.Error()
	}

	resp.State = standard.State_SUCCESS
	return resp, err
}

// SendVerifyCodeBySms SendVerifyCodeBySms
func (srv *Service) SendVerifyCodeBySms(ctx context.Context, req *standard.SendVerifyCodeBySmsRequest) (resp *standard.SendVerifyCodeBySmsResponse, err error) {
	resp = new(standard.SendVerifyCodeBySmsResponse)
	smsProvider := srv.nextSMSProvider()
	if smsProvider == nil {
		resp.State = standard.State_NO_SERVICE_AVAILABLE
		return resp, nil
	}

	verifyCode := srv.randomCode(6)
	err = smsProvider.SendVerifyCode(req.Operation, verifyCode)
	if err != nil {
		resp.State = standard.State_FAILURE
		resp.Message = err.Error()
	}

	expireTime := time.Now().Add(time.Minute * 5)
	key, err := dao.CreateVerifyCode(req.Operation, verifyCode, expireTime)
	if err != nil {
		resp.State = standard.State_DB_OPERATION_FATLURE
		resp.Message = err.Error()
	}

	resp.State = standard.State_SUCCESS
	resp.Key = key
	return resp, err
}

// SendVerifyCodeByEmail SendVerifyCodeByEmail
func (srv *Service) SendVerifyCodeByEmail(ctx context.Context, req *standard.SendVerifyCodeByEmailRequest) (resp *standard.SendVerifyCodeByEmailResponse, err error) {
	resp = new(standard.SendVerifyCodeByEmailResponse)
	emailProvider := srv.nextEmailProvider()
	if emailProvider == nil {
		resp.State = standard.State_NO_SERVICE_AVAILABLE
		return resp, nil
	}

	verifyCode := srv.randomCode(6)
	err = emailProvider.SendVerifyCode(req.EmailAddress, req.Operation, srv.randomCode(6))
	if err != nil {
		resp.State = standard.State_FAILURE
		resp.Message = err.Error()
	}

	expireTime := time.Now().Add(time.Minute * 5)
	key, err := dao.CreateVerifyCode(req.Operation, verifyCode, expireTime)
	if err != nil {
		resp.State = standard.State_DB_OPERATION_FATLURE
		resp.Message = err.Error()
	}

	resp.State = standard.State_SUCCESS
	resp.Key = key
	return resp, err
}

// SendVerifyCodeByCallPhone SendVerifyCodeByCallPhone
func (srv *Service) SendVerifyCodeByCallPhone(ctx context.Context, req *standard.SendVerifyCodeByCallPhoneRequest) (resp *standard.SendVerifyCodeByCallPhoneResponse, err error) {
	resp = new(standard.SendVerifyCodeByCallPhoneResponse)
	callProvider := srv.nextCallProvider()
	if callProvider == nil {
		resp.State = standard.State_NO_SERVICE_AVAILABLE
		return resp, nil
	}

	verifyCode := srv.randomCode(6)
	err = callProvider.SendVerifyCode(req.Operation, srv.randomCode(6))
	if err != nil {
		resp.State = standard.State_FAILURE
		resp.Message = err.Error()
	}

	expireTime := time.Now().Add(time.Minute * 5)
	key, err := dao.CreateVerifyCode(req.Operation, verifyCode, expireTime)
	if err != nil {
		resp.State = standard.State_DB_OPERATION_FATLURE
		resp.Message = err.Error()
	}

	resp.State = standard.State_SUCCESS
	resp.Key = key
	return resp, err
}
