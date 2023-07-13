// Package user 存放用户 Model 相关逻辑
package user

import (
	"gentleman/app/models"
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
	Username       string `json:"username"`
	AcCount        int    `json:"ac_count"`
	YesterdayCount int    `json:"yesterday_count"`
	LastWeekCount  int    `json:"last_week_count"`
	QQScore        int    `json:"qq_score"`
	GGname         string `json:"GGname"`
	GGemail        string `json:"-"`
	GGpassword     string `json:"-"`
	models.CommonTimestampsField
}

func (User) TableName() string {
	return "user"
}
