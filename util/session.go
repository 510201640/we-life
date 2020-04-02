package util

import (
	"jaden/we-life/common"
	"time"
)

//生成session
func GenSession() string {

	return EncryptToMD5([]byte(GetRandString(32)))
}

//session保存7天
func SaveSession(session string) {
	if !common.Redis.Exists(session).Val() {
		common.Redis.SetNX(session, 1, time.Hour*24*7)
	}
}

//检查session是否存钻
func CheckSessionExist(session string) bool {
	return common.Redis.Exists(session).Val()
}
