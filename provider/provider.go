package provider

import (
	"context"
	"github.com/grpcbrick/sender/standard"
)

// SMSProvider 短信服务
type SMSProvider interface {
	Weights() int
	SendVerifyCode(code string) error
}

// CallProvider 电话服务
type CallProvider interface {
	Weights() int
	SendVerifyCode(code string) error
}

// EmailProvider 邮件服务
type EmailProvider interface {
	Weights() int
	SendVerifyCode(code string) error
}

// Service Service
type Service struct {
	SMS   []*SMSProvider
	Call  []*CallProvider
	Email []*EmailProvider
}

// NewService NewService
func NewService() *Service {
	service := new(Service)
	return service
}

// CheckVerifyCode CheckVerifyCode
func (srv Service) CheckVerifyCode(ctx context.Context, req *standard.CheckVerifyCodeRequest) (resp *standard.CheckVerifyCodeeResponse, err error) {
	resp = new(standard.CheckVerifyCodeeResponse)
	return resp, err
}

// DiscardVerifyCode DiscardVerifyCode
func (srv Service) DiscardVerifyCode(ctx context.Context, req *standard.DiscardVerifyCodeRequest) (resp *standard.DiscardVerifyCodeResponse, err error) {
	resp = new(standard.DiscardVerifyCodeResponse)
	return resp, err
}

// SendVerifyCodeBySms SendVerifyCodeBySms
func (srv Service) SendVerifyCodeBySms(ctx context.Context, req *standard.SendVerifyCodeBySmsRequest) (resp *standard.SendVerifyCodeBySmsResponse, err error) {
	resp = new(standard.SendVerifyCodeBySmsResponse)
	return resp, err
}

// SendVerifyCodeByEmail SendVerifyCodeByEmail
func (srv Service) SendVerifyCodeByEmail(ctx context.Context, req *standard.SendVerifyCodeByEmailRequest) (resp *standard.SendVerifyCodeByEmailResponse, err error) {
	resp = new(standard.SendVerifyCodeByEmailResponse)
	return resp, err
}

// SendVerifyCodeByCallPhone SendVerifyCodeByCallPhone
func (srv Service) SendVerifyCodeByCallPhone(ctx context.Context, req *standard.SendVerifyCodeByCallPhoneRequest) (resp *standard.SendVerifyCodeByCallPhoneResponse, err error) {
	resp = new(standard.SendVerifyCodeByCallPhoneResponse)
	return resp, err
}
