// Package requests 处理请求数据和表单验证
package requests

import (
	"gentleman/app/requests/validators"
	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
)

type SignupEmailExistRequest struct {
	Email string `json:"email,omitempty"`
}

func SignupEmailExist(data interface{}, c *gin.Context) map[string][]string {
	// 自定义验证规则
	rules := govalidator.MapData{
		"email": []string{"required", "min:4", "max:30", "email"},
	}
	// 自定义验证出错时的提示
	messages := govalidator.MapData{
		"email": []string{
			"required:Email 为必填项",
			"min:Email 长度需大于 4",
			"max:Email 长度需小于 30",
			"email:Email 格式不正确，请提供有效的邮箱地址",
		},
	}
	return validate(data, rules, messages)
}

// SignupUsingEmailRequest 通过邮箱注册的请求信息
type SignupUsingEmailRequest struct {
	Email           string `json:"email,omitempty" valid:"email"`
	VerifyCode      string `json:"verify_code,omitempty" valid:"verify_code"`
	Name            string `valid:"name" json:"name"`
	Password        string `valid:"password" json:"password,omitempty"`
	PasswordConfirm string `valid:"password_confirm" json:"password_confirm,omitempty"`
}

func SignupUsingEmail(data interface{}, c *gin.Context) map[string][]string {

	rules := govalidator.MapData{
		"email":            []string{"required", "min:4", "max:30", "email", "not_exists:users,email"},
		"name":             []string{"required", "alpha_num", "between:3,20", "not_exists:users,name"},
		"password":         []string{"required", "min:6"},
		"password_confirm": []string{"required"},
		"verify_code":      []string{"required", "digits:6"},
	}

	messages := govalidator.MapData{
		"email": []string{
			"required:Email 为必填项",
			"min:Email 长度需大于 4",
			"max:Email 长度需小于 30",
			"email:Email 格式不正确，请提供有效的邮箱地址",
			"not_exists:Email 已被占用",
		},
		"name": []string{
			"required:用户名为必填项",
			"alpha_num:用户名格式错误，只允许数字和英文",
			"between:用户名长度需在 3~20 之间",
		},
		"password": []string{
			"required:密码为必填项",
			"min:密码长度需大于 6",
		},
		"password_confirm": []string{
			"required:确认密码框为必填项",
		},
		"verify_code": []string{
			"required:验证码答案必填",
			"digits:验证码长度必须为 6 位的数字",
		},
	}

	errs := validate(data, rules, messages)

	_data := data.(*SignupUsingEmailRequest)
	errs = validators.ValidatePasswordConfirm(_data.Password, _data.PasswordConfirm, errs)
	errs = validators.ValidateVerifyCode(_data.Email, _data.VerifyCode, errs)

	return errs
}

type SignupQQExistRequest struct {
	QQ string `json:"qq,omitempty" valid:"qq"`
}

func SignupQQExist(data interface{}, c *gin.Context) map[string][]string {
	// 自定义验证规则
	rules := govalidator.MapData{
		"qq": []string{"required", "min:4", "max:10"},
	}
	// 自定义验证出错时的提示
	messages := govalidator.MapData{
		"qq": []string{
			"required:QQ 为必填项",
			"min:QQ 长度需大于 4",
			"max:QQ 长度需小于 10",
		},
	}
	return validate(data, rules, messages)
}

// SignupUsingQQRequest 通过邮箱注册的请求信息
type SignupUsingQQRequest struct {
	QQ              string `json:"qq,omitempty" valid:"qq"`
	Name            string `valid:"name" json:"name"`
	Password        string `valid:"password" json:"password,omitempty"`
	PasswordConfirm string `valid:"password_confirm" json:"password_confirm,omitempty"`
	LCUserName      string `valid:"lcusername" json:"lcusername,omitempty"`
	CaptchaID       string `json:"captcha_id,omitempty" valid:"captcha_id"`
	CaptchaAnswer   string `json:"captcha_answer,omitempty" valid:"captcha_answer"`
}

func SignupUsingQQ(data interface{}, c *gin.Context) map[string][]string {

	rules := govalidator.MapData{
		"qq":               []string{"required", "min:4", "max:10", "not_exists:users,qq"},
		"name":             []string{"required", "alpha_num", "between:3,20", "not_exists:users,name"},
		"password":         []string{"required", "min:6"},
		"password_confirm": []string{"required"},
		"lcusername":       []string{"required", "check-lc:users,homepage"},
		"captcha_id":       []string{"required"},
		"captcha_answer":   []string{"required", "digits:6"},
	}

	messages := govalidator.MapData{
		"qq": []string{
			"required:QQ 为必填项",
			"min:QQ 长度需大于 4",
			"max:QQ 长度需小于 10",
			"not_exists:QQ 已被占用",
		},
		"name": []string{
			"required:用户名为必填项",
			"alpha_num:用户名格式错误，只允许数字和英文",
			"between:用户名长度需在 3~20 之间",
		},
		"password": []string{
			"required:密码为必填项",
			"min:密码长度需大于 6",
		},
		"password_confirm": []string{
			"required:确认密码框为必填项",
		},
		"lcusername": []string{
			"required:力扣用户域名为必填项",
			"check-lc: 域名不存在或者被占用",
		},
		"captcha_id": []string{
			"required:图片验证码的 ID 为必填",
		},
		"captcha_answer": []string{
			"required:图片验证码答案必填",
			"digits:图片验证码长度必须为 6 位的数字",
		},
	}

	errs := validate(data, rules, messages)
	_data := data.(*SignupUsingQQRequest)
	errs = validators.ValidatePasswordConfirm(_data.Password, _data.PasswordConfirm, errs)

	return errs
}
