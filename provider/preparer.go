package provider

import (
	"bytes"
	"os"

	_ "github.com/go-sql-driver/mysql" // mysql 驱动
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
)

var (
	createVerifyCodeRecordStmt                        *sqlx.Stmt
	insertVerifyCodeRecordStmt                        *sqlx.Stmt
	discardVerifyCodeRecordExpiredStmt                *sqlx.Stmt
	discardVerifyCodeRecordBySymbolStmt               *sqlx.Stmt
	discardVerifyCodeRecordBySymbolWithVerifyCodeStmt *sqlx.Stmt
	checkVerifyCodeRecordBySymbolWithVerifyCodeStmt   *sqlx.Stmt
)

func init() {
	var err error
	godotenv.Load()

	database, err := sqlx.Connect("mysql", os.Getenv("MYSQL_DB_URI"))
	if err != nil {
		panic(err)
	}

	// 创建验证码表
	createVerifyCodeRecordStmt = MustPreparex(
		database,
		" CREATE TABLE IF NOT EXISTS `verify_code_records` (",
		" `ID` int(11) NOT NULL AUTO_INCREMENT COMMENT 'ID',",
		" `Disuse` tinyint(1) NOT NULL COMMENT '是否已经废弃',",
		" `Symbol` varchar(128) NOT NULL COMMENT '唯一标识符',",
		" `VerifyCode` varchar(128) NOT NULL COMMENT '验证码',",
		" `ValidityPeriod` timestamp NOT NULL COMMENT '到期时间',",
		" `CreateTime` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',",
		" PRIMARY KEY (`ID`, `Disuse`, `Symbol`, `VerifyCode`, `CreateTime`)",
		" )ENGINE=InnoDB AUTO_INCREMENT=0 DEFAULT CHARSET=utf8mb4;",
	)

	_, err = createVerifyCodeRecordStmt.Exec()
	if err != nil {
		panic(err)
	}

	// 插入验证码记录
	insertVerifyCodeRecordStmt = MustPreparex(
		database,
		" INSERT INTO `verify_code_records`",
		" (`Disuse`, `Symbol`, `VerifyCode`,`ValidityPeriod`)",
		" VALUES",
		" (?, ?, ?, ?);",
	)

	// 清除所有过期的
	discardVerifyCodeRecordExpiredStmt = MustPreparex(
		database,
		"UPDATE `verify_code_records`",
		" SET `Disuse` = 1",
		" WHERE`Disuse` = 0",
		" AND`ValidityPeriod` < NOW();",
	)

	// 报废指定 symbol 的发送记录
	discardVerifyCodeRecordBySymbolStmt = MustPreparex(
		database,
		" UPDATE `verify_code_records`",
		" SET `Disuse` = 1",
		" WHERE `Disuse` = 0",
		" AND `symbol` LIKE ? ;",
	)

	// 报废指定 symbol 的发送记录
	discardVerifyCodeRecordBySymbolWithVerifyCodeStmt = MustPreparex(
		database,
		"UPDATE `verify_code_records`",
		" SET `Disuse` = 1",
		" WHERE `Disuse` = 0",
		" AND `symbol` LIKE ?",
		" AND `symbol` = ? ;",
	)

	// 通过 symbol 查询指定的发送记录
	checkVerifyCodeRecordBySymbolWithVerifyCodeStmt = MustPreparex(
		database,
		"SELECT COUNT(*) FROM `verify_code_records`",
		" WHERE `Disuse` = 1",
		" AND `Symbol` LIKE ?",
		" AND `VerifyCode` LIKE ?;",
	)

}

// MustPreparex 解析 query
func MustPreparex(database *sqlx.DB, querys ...string) *sqlx.Stmt {
	var queryBuf bytes.Buffer

	for _, s := range querys {
		queryBuf.WriteString(s)
	}

	stmp, err := database.Preparex(queryBuf.String())
	if err != nil {
		panic(err)
	}
	return stmp
}

// MustPreparexNamed 解析 query
func MustPreparexNamed(database *sqlx.DB, querys ...string) *sqlx.NamedStmt {
	var queryBuf bytes.Buffer

	for _, s := range querys {
		queryBuf.WriteString(s)
	}

	stmp, err := database.PrepareNamed(queryBuf.String())
	if err != nil {
		panic(err)
	}
	return stmp
}
