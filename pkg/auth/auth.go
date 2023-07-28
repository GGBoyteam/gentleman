// Package auth 授权相关逻辑
package auth

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"gentleman/app/models/user"
	"gentleman/pkg/logger"
	"github.com/gin-gonic/gin"
)

func showstruct(data interface{}) string {
	bs, _ := json.Marshal(data)
	var out bytes.Buffer
	json.Indent(&out, bs, "", "\t")

	return fmt.Sprintf("你要查看的结构体=%v\n", out.String())
}

// Attempt 尝试登录
func Attempt(email string, password string) (user.User, error) {
	userModel := user.GetByMulti(email)
	fmt.Println(showstruct(userModel))
	if userModel.ID == 0 {
		return user.User{}, errors.New("账号不存在")
	}

	if !userModel.ComparePassword(password) {
		return user.User{}, errors.New("密码错误")
	}

	return userModel, nil
}

// LoginByPhone 登录指定用户
func LoginByPhone(phone string) (user.User, error) {
	userModel := user.GetByPhone(phone)
	if userModel.ID == 0 {
		return user.User{}, errors.New("手机号未注册")
	}

	return userModel, nil
}

// CurrentUser 从 gin.context 中获取当前登录用户
func CurrentUser(c *gin.Context) user.User {
	userModel, ok := c.MustGet("current_user").(user.User)
	if !ok {
		logger.LogIf(errors.New("无法获取用户"))
		return user.User{}
	}
	// db is now a *DB value
	return userModel
}

// CurrentUID 从 gin.context 中获取当前登录用户 ID
func CurrentUID(c *gin.Context) string {
	return c.GetString("current_user_id")
}
