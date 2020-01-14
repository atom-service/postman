package dao

import (
	"fmt"
	"math"
	"math/rand"
	"strings"
	"time"

	"github.com/grpcbrick/sender/models"
	"github.com/yinxulai/goutils/sqldb"
)

const verifyCodeTableName = "verify-code"

func createVerifyCodeTable() error {
	var err error
	// 主表
	masterStmp := sqldb.CreateStmt(strings.Join([]string{
		"CREATE TABLE IF NOT EXISTS",
		"`" + verifyCodeTableName + "`",
		"(",
		"`Key` VARCHAR(128) NOT NULL COMMENT 'Key',",
		"`Code` varchar(128) NOT NULL COMMENT '验证码',",
		"`Operation` varchar(128) DEFAULT '' COMMENT '操作',",
		"`ExpireTime` datetime DEFAULT NULL COMMENT '过期时间',",
		"`DeletedTime` datetime DEFAULT NULL COMMENT '删除时间',",
		"`CreatedTime` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',",
		"`UpdatedTime` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',",
		"PRIMARY KEY (`Key`,`Code`),",
		"UNIQUE KEY (`Key`)",
		")",
		"ENGINE=InnoDB AUTO_INCREMENT=0 DEFAULT CHARSET=utf8mb4;",
	}, " "))

	_, err = masterStmp.Exec()
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

	return nil
}

// 生成一个为一个 key
func generateUniqueKey() (string, error) {
	randomString := func(l int) string {
		str := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
		bytes := []byte(str)
		result := []byte{}
		r := rand.New(rand.NewSource(time.Now().UnixNano()))
		for i := 0; i < l; i++ {
			result = append(result, bytes[r.Intn(len(bytes))])
		}
		return string(result)
	}

	for {
		newKey := randomString(128)
		count, err := CountVerifyCodeByKey(newKey)
		if err != nil {
			return "", err
		}

		if count <= 0 {
			return newKey, nil
		}
	}
}

func clearExpiredVerifyCode() error {
	stmp := sqldb.CreateNamedStmt(strings.Join([]string{
		"DELETE FROM",
		"`" + verifyCodeTableName + "`",
		"WHERE",
		"`ExpireTime` <= :CurrentTime",
	}, " "))

	namedData := map[string]interface{}{
		"CurrentTime": time.Now(),
	}

	_, err := stmp.Exec(namedData)
	if err != nil {
		return err
	}

	return nil
}

// CreateVerifyCode 一个验证码
func CreateVerifyCode(code, operation string, expireTime time.Time) (string, error) {
	key, err := generateUniqueKey()
	if err != nil {
		return "", err
	}

	stmp := sqldb.CreateNamedStmt(strings.Join([]string{
		"INSERT INTO",
		"`" + verifyCodeTableName + "`",
		"(`Key`, `Code`, `Operation`, `ExpireTime`)",
		"VALUES",
		"(:Key, :Code, :Operation, :ExpireTime)",
	}, " "))

	data := map[string]interface{}{
		"Key":        key,
		"Code":       code,
		"Operation":  operation,
		"ExpireTime": expireTime,
	}

	_, err = stmp.Exec(data)
	if err != nil {
		return key, err
	}
	return key, err
}

// CountVerifyCodeByKey 根据 id 统计
func CountVerifyCodeByKey(key string) (int64, error) {
	var err error
	err = clearExpiredVerifyCode()
	if err != nil {
		return 0, err
	}

	stmp := sqldb.CreateNamedStmt(strings.Join([]string{
		"SELECT COUNT(*) as Count FROM",
		"`" + verifyCodeTableName + "`",
		"WHERE",
		"`Key`=:Key",
	}, " "))

	result := struct{ Count int64 }{}
	namedData := map[string]interface{}{
		"Key": key,
	}

	err = stmp.Get(&result, namedData)
	if err != nil {
		return 0, err
	}
	return result.Count, nil
}

// QueryVerifyCodes 查询组
func QueryVerifyCodes(page, limit int64) (totalPage, currentPage int64, verifyCodes []*models.VerifyCode, err error) {
	err = clearExpiredVerifyCode()
	if err != nil {
		return totalPage, currentPage, verifyCodes, err
	}

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

// QueryVerifyCodeByKey 根据 HashCode 查询
func QueryVerifyCodeByKey(key string) (*models.VerifyCode, error) {
	var err error
	err = clearExpiredVerifyCode()
	if err != nil {
		return nil, err
	}

	namedData := map[string]interface{}{"Key": key}
	stmp := sqldb.CreateNamedStmt(strings.Join([]string{
		"SELECT * FROM",
		"`" + verifyCodeTableName + "`",
		"WHERE",
		"`Key`=:Key",
	}, " "))

	verifyCode := new(models.VerifyCode)
	err = stmp.Get(verifyCode, namedData)
	if err != nil {
		return nil, err
	}

	return verifyCode, nil
}

// UpdateVerifyCodeExpireTimeByKey 更新标签值
func UpdateVerifyCodeExpireTimeByKey(key string, expireTime time.Time) error {
	return updataVerifyCodeFieldByKey(key, map[string]interface{}{"ExpireTime": expireTime})
}

// 根据 ID 更新标签
func updataVerifyCodeFieldByKey(key string, field map[string]interface{}) error {
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
		"`Key`=:Key",
	}, " "))

	// 更新
	field["ID"] = key
	_, err = stmp.Exec(field)
	return err
}
