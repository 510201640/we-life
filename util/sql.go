package util

import (
	"errors"
	"fmt"
	"sdyxmall/business-examine/wrapper/page"
	"strings"
)

// 生成相同条件查询的统计数量SQL
func GetSameConditionCountSQL(sqlStr string) (string, error) {
	index := strings.Index(sqlStr, "FROM")
	if index == -1 {
		index = strings.Index(strings.ToUpper(sqlStr), "FROM")
	}
	if index == -1 {
		return "", errors.New("invalid sql")
	}
	return "SELECT count(*) AS total " + sqlStr[index:], nil
}

// 从page对象生成limit 条件SQL语句
func GetLimitSQLFromPage(pageMeta *page.Page) string {
	if pageMeta == nil || !pageMeta.Valid() {
		return ""
	}
	pageNum := pageMeta.PageNum
	pageSize := pageMeta.PageSize
	return fmt.Sprintf(" LIMIT %d, %d", (pageNum-1)*pageSize, pageSize)
}

func AndIn(columnName string, data ...interface{}) string {
	if len(data) == 0 {
		return ""
	}
	tags := strings.Repeat("?,", len(data))
	tags = strings.TrimRight(tags, ",")
	sql := fmt.Sprintf(" and %s in(%s)", columnName, tags)
	return sql
}
