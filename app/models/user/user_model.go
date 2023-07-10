// Package user 存放用户 Model 相关逻辑
package user

import (
	"gentleman/app/models"
)

// User 用户模型
type User struct {
	models.BaseModel

	QQ     string `json:"qq"`
	level  int    `json:"-"`
	Title  string `json:"title"`
	Score  int    `json:"score"`
	Status string `json:"status"`

	models.CommonTimestampsField
}

func (User) TableName() string {
	return "user"
}
