package models

// VerifyCodeRecord 各类验证码
type VerifyCodeRecord struct {
	ID             uint64 `db:"ID"`             // ID
	Disuse         bool   `db:"Disuse"`         // 弃用
	Symbol         string `db:"Symbol"`         // 标识符
	VerifyCode     string `db:"VerifyCode"`     // 验证码
	CreateTime     string `db:"CreateTime"`     // 创建时间
	ValidityPeriod string `db:"ValidityPeriod"` // 有效期
}

// NoticeRecord 通知 提醒 记录
type NoticeRecord struct {
}
