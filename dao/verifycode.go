package dao

import (
	"fmt"
	"math"
	"strings"
	"time"

	"github.com/grpcbrick/sender/models"
	"github.com/yinxulai/goutils/sqldb"
)

const verifyCodeTableName = "verify-code"
const verifyCodeHistoryTableName = "verify-code-history"

func createVerifyCodeTable() error {
	var err error
	// 主表
	masterStmp := sqldb.CreateStmt(strings.Join([]string{
		"CREATE TABLE IF NOT EXISTS",
		"`" + verifyCodeTableName + "`",
		"(",
		"`ID` int(11) NOT NULL AUTO_INCREMENT COMMENT 'ID',",
		"`Code` varchar(128) NOT NULL COMMENT '验证码',",
		"`Operation` varchar(128) DEFAULT '' COMMENT '操作',",
		"`ExpireTime` datetime DEFAULT '' COMMENT '过期时间',",
		"`DeletedTime` datetime DEFAULT NULL COMMENT '删除时间',",
		"`CreatedTime` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',",
		"`UpdatedTime` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',",
		"PRIMARY KEY (`ID`,`Code`,`ExpireTime`),",
		"UNIQUE KEY `ID` (`ID`)",
		")",
		"ENGINE=InnoDB AUTO_INCREMENT=0 DEFAULT CHARSET=utf8mb4;",
	}, " "))
	_, err = masterStmp.Exec()
	if err != nil {
		return err
	}

	// 历史表
	historyStmp := sqldb.CreateStmt(strings.Join([]string{
		"CREATE TABLE IF NOT EXISTS",
		"`" + verifyCodeHistoryTableName + "`",
		"(",
		"`Index` int(11) NOT NULL AUTO_INCREMENT COMMENT 'Index',",
		"`ID` int(11) COMMENT 'ID',",
		"`Code` varchar(128) COMMENT '验证码',",
		"`Operation` varchar(128) COMMENT '操作',",
		"`ExpireTime` datetime COMMENT '过期时间',",
		"`DeletedTime` datetime COMMENT '删除时间',",
		"`CreatedTime` datetime COMMENT '创建时间',",
		"`UpdatedTime` datetime COMMENT '更新时间',",
		"PRIMARY KEY (`Index`)",
		")",
		"ENGINE=InnoDB AUTO_INCREMENT=0 DEFAULT CHARSET=utf8mb4;",
	}, " "))
	_, err = historyStmp.Exec()
	if err != nil {
		return err
	}

	return nil
}

func truncateVerifyCodeTable() error {
	var err error
	masterStmp := sqldb.CreateStmt("truncate table `" + verifyCodeTableName + "`")
	_, err = masterStmp.Exec()
	if err != nil {
		return err
	}

	historyStmp := sqldb.CreateStmt("truncate table `" + verifyCodeHistoryTableName + "`")
	_, err = historyStmp.Exec()
	if err != nil {
		return err
	}
	return nil
}

// CreateVerifyCodeHistory 对指定数据创建一条历史快照
func CreateVerifyCodeHistory(id int64) error {
	var err error
	namedData := map[string]interface{}{
		"ID": id,
	}

	// 插入一条更新历史
	historyStmp := sqldb.CreateNamedStmt(strings.Join([]string{
		"INSERT INTO",
		"`" + verifyCodeHistoryTableName + "`",
		"(`ID`,`Code`,`Operation`,`ExpireTime`,`DeletedTime`,`CreatedTime`,`UpdatedTime`)",
		"SELECT",
		"`ID`,`Code`,`Operation`,`ExpireTime`,`DeletedTime`,`CreatedTime`,`UpdatedTime`",
		"FROM",
		"`" + verifyCodeTableName + "`",
		"WHERE",
		"`ID`=:ID",
	}, " "))
	_, err = historyStmp.Exec(namedData)
	if err != nil {
		return err
	}
	return nil
}

// CreateVerifyCode 一个验证码
func CreateVerifyCode(code, operation, expireTime string) (int64, error) {
	stmp := sqldb.CreateNamedStmt(strings.Join([]string{
		"INSERT INTO",
		"`" + verifyCodeTableName + "`",
		"(`Code`, `Operation`, `ExpireTime`)",
		"VALUES",
		"(:Code, :Operation, :ExpireTime)",
	}, " "))

	data := map[string]interface{}{
		"Code":       code,
		"Operation":  operation,
		"ExpireTime": expireTime,
	}
	result, err := stmp.Exec(data)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	return id, err
}

// CountVerifyCodeByID 根据 id 统计
func CountVerifyCodeByID(id int64) (int64, error) {
	stmp := sqldb.CreateNamedStmt(strings.Join([]string{
		"SELECT COUNT(*) as Count FROM",
		"`" + verifyCodeTableName + "`",
		"WHERE",
		"`ID`=:ID",
	}, " "))

	result := struct{ Count int64 }{}
	namedData := map[string]interface{}{
		"ID": id,
	}
	err := stmp.Get(&result, namedData)
	if err != nil {
		return 0, err
	}
	return result.Count, nil
}

// QueryVerifyCodes 查询组
func QueryVerifyCodes(page, limit int64) (totalPage, currentPage int64, verifyCodes []*models.VerifyCode, err error) {
	currentPage = page // 固定当前页

	// 查询数据长度
	countStmp := sqldb.CreateStmt(strings.Join([]string{
		"SELECT COUNT(*) as `Count` FROM",
		"`" + verifyCodeTableName + "`",
	}, " "))

	countResult := struct{ Count int64 }{}
	err = countStmp.Get(&countResult)
	if err != nil {
		return totalPage, currentPage, verifyCodes, err
	}

	count := countResult.Count
	// 计算总页码数
	totalPage = int64(math.Ceil(float64(count) / float64(limit)))

	// 超出数据总页数
	if page > totalPage {
		// 返回当前页、空数据（当前页确实不存在数据）
		return totalPage, page, verifyCodes, err
	}

	// 计算偏移
	offSet := (page - 1) * limit
	stmp := sqldb.CreateNamedStmt(strings.Join([]string{
		"SELECT * FROM",
		"`" + verifyCodeTableName + "`",
		"LIMIT :Limit",
		"OFFSET :Offset",
	}, " "))

	verifyCodes = []*models.VerifyCode{}
	namedData := map[string]interface{}{
		"Limit":  limit,
		"Offset": offSet,
	}

	err = stmp.Select(&verifyCodes, namedData)
	if err != nil {
		return totalPage, currentPage, verifyCodes, err
	}

	return totalPage, currentPage, verifyCodes, err
}

// QueryVerifyCodeByID 根据 HashCode 查询
func QueryVerifyCodeByID(id int64) (*models.VerifyCode, error) {
	var err error
	namedData := map[string]interface{}{"ID": id}
	stmp := sqldb.CreateNamedStmt(strings.Join([]string{
		"SELECT * FROM",
		"`" + verifyCodeTableName + "`",
		"WHERE",
		"`ID`=:ID",
	}, " "))

	verifyCodes := new(models.VerifyCode)
	err = stmp.Get(verifyCodes, namedData)
	if err != nil {
		return nil, err
	}

	return verifyCodes, nil
}

// DeleteVerifyCodeByID 删除标签
func DeleteVerifyCodeByID(id int64) error {
	var err error
	namedData := map[string]interface{}{"ID": id}
	stmp := sqldb.CreateNamedStmt(strings.Join([]string{
		"DELETE FROM",
		"`" + verifyCodeTableName + "`",
		"WHERE",
		"`ID`=:ID",
	}, " "))

	// 更新
	err = updataVerifyCodeFieldByID(id, map[string]interface{}{"DeletedTime": time.Now()})
	if err != nil {
		return err
	}

	_, err = stmp.Exec(namedData)
	return err
}

// updataVerifyCodeFieldByID 更新标签类型
func UpdateVerifyCodeCodeByID(id int64, name string) error {
	return updataVerifyCodeFieldByID(id, map[string]interface{}{"Name": name})
}

// UpdateVerifyCodeOperationByID 更新标签状态
func UpdateVerifyCodeOperationByID(id int64, category string) error {
	return updataVerifyCodeFieldByID(id, map[string]interface{}{"Category": category})
}

// UpdateVerifyCodeExpireTimeByID 更新标签值
func UpdateVerifyCodeExpireTimeByID(id int64, description string) error {
	return updataVerifyCodeFieldByID(id, map[string]interface{}{"Description": description})
}

// 根据 ID 更新标签
func updataVerifyCodeFieldByID(id int64, field map[string]interface{}) error {
	var err error

	fieldSQL := []string{}
	for name := range field {
		fieldSQL = append(fieldSQL, fmt.Sprintf("`%s`=:%s", name, name))
	}

	stmp := sqldb.CreateNamedStmt(strings.Join([]string{
		"UPDATE",
		"`" + verifyCodeTableName + "`",
		"SET",
		strings.Join(fieldSQL, ","),
		"WHERE",
		"`ID`=:ID",
	}, " "))

	// 更新
	field["ID"] = id
	_, err = stmp.Exec(field)
	return err
}
