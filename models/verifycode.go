package models

import "database/sql"

// VerifyCode 各类验证码
type VerifyCode struct {
	ID          sql.NullInt64
	Code        sql.NullString // 验证码
	Operation   sql.NullString // 操作
	ExpireTime  sql.NullTime   // 过期时间
	DeletedTime sql.NullTime
	CreatedTime sql.NullTime
	UpdatedTime sql.NullTime
}

// NoticeRecord 通知 提醒 记录
type NoticeRecord struct {
}
