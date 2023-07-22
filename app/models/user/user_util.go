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

// GetByPhone 通过手机号来获取用户
func GetByPhone(phone string) (userModel User) {
	database.DB.Where("phone = ?", phone).First(&userModel)
	return
}

// GetByMulti 通过 手机号/Email/用户名 来获取用户
func GetByMulti(loginID string) (userModel User) {
	database.DB.
		Where("qq = ?", loginID).
		Or("name = ?", loginID).
		First(&userModel)
	return
}
