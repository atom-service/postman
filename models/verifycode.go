package models

import (
	"database/sql"
	"github.com/grpcbrick/sender/standard"
)

// VerifyCode 各类验证码
type VerifyCode struct {
	Key         sql.NullString
	Code        sql.NullString // 验证码
	Operation   sql.NullString // 操作
	ExpireTime  sql.NullTime   // 过期时间
	DeletedTime sql.NullTime
	CreatedTime sql.NullTime
	UpdatedTime sql.NullTime
}

// LoadProtoStruct LoadProtoStruct
func (srv *VerifyCode) LoadProtoStruct(Data *standard.VerifyCode) {
	srv.Key.Scan(Data.Key)
	srv.Code.Scan(Data.Code)
	srv.Operation.Scan(Data.Operation)
	srv.ExpireTime.Scan(Data.ExpireTime)
	srv.DeletedTime.Scan(Data.DeletedTime)
	srv.CreatedTime.Scan(Data.CreatedTime)
	srv.UpdatedTime.Scan(Data.UpdatedTime)
}

// OutProtoStruct OutProtoStruct
func (srv *VerifyCode) OutProtoStruct() *standard.VerifyCode {
	Data := new(standard.VerifyCode)
	Data.Key = srv.Key.String
	Data.Code = srv.Code.String
	Data.Operation = srv.Operation.String
	Data.ExpireTime = srv.ExpireTime.Time.String()
	Data.DeletedTime = srv.DeletedTime.Time.String()
	Data.CreatedTime = srv.CreatedTime.Time.String()
	Data.UpdatedTime = srv.UpdatedTime.Time.String()
	return Data
}

// NoticeRecord 通知 提醒 记录
type NoticeRecord struct {
}
