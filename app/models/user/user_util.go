package user

import (
	"gentleman/pkg/database"
)

// IsEmailExist 判断 Email 已被注册
func IsEmailExist(email string) bool {
	var count int64
	database.DB.Model(User{}).Where("email = ?", email).Count(&count)
	return count > 0
}

// IsQQExist 判断 QQ 已被注册
func IsQQExist(qq string) bool {
	var count int64
	database.DB.Model(User{}).Where("email = ?", qq).Count(&count)
	return count > 0
}
