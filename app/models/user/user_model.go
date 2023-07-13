// Package user 存放用户 Model 相关逻辑
package user

import (
	"gentleman/app/models"
	"gentleman/pkg/database"
	"gentleman/pkg/hash"
)

// User 用户模型
type User struct {
	models.BaseModel

	QQ             string `json:"qq"`
	level          int    `json:"level"`
	Title          string `json:"title"`
	Score          int    `json:"score"`
	Status         string `json:"status"`
	Homepage       string `json:"homepage"`
	Username       string `json:"username"` // 力扣网站的用户名
	AcCount        int    `json:"ac_count"`
	YesterdayCount int    `json:"yesterday_count"`
	LastWeekCount  int    `json:"last_week_count"`
	QQScore        int    `json:"qq_score"`
	Name           string `json:"name"` // 考勤网站的用户名
	Email          string `json:"-"`
	Password       string `json:"-"`
	models.CommonTimestampsField
}

// Create 创建用户，通过 User.ID 来判断是否创建成功
func (userModel *User) Create() {
	database.DB.Create(&userModel)
}

// ComparePassword 密码是否正确
func (userModel *User) ComparePassword(_password string) bool {
	return hash.BcryptCheck(_password, userModel.Password)
}
