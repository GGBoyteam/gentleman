// Package routes 注册路由
package routes

import (
	"gentleman/app/http/controllers/api/v1/auth"
	"gentleman/app/models/user"
	"gentleman/pkg/database"
	"github.com/gin-gonic/gin"
	"net/http"
)

// RegisterAPIRoutes 注册网页相关路由
func RegisterAPIRoutes(r *gin.Engine) {

	// 测试一个 v1 的路由组，我们所有的 v1 版本的路由都将存放到这里
	v1 := r.Group("/v1")
	{
		// 注册一个路由
		v1.GET("/", func(c *gin.Context) {
			c.HTML(http.StatusOK, "home.html", nil)
		})

		v1.GET("/users", getUser)
		v1.GET("/query_table", func(c *gin.Context) {
			c.HTML(http.StatusOK, "query.html", nil)
		})

		v1.GET("/fly_table", func(c *gin.Context) {
			c.HTML(http.StatusOK, "fly.html", nil)
		})

		authGroup := v1.Group("/auth")
		{
			suc := new(auth.SignupController)
			// 判断 Email 是否已注册
			authGroup.POST("/signup/email/exist", suc.IsEmailExist)

			// 发送验证码
			vcc := new(auth.VerifyCodeController)
			// 图片验证码，需要加限流
			authGroup.POST("/verify-codes/captcha", vcc.ShowCaptcha)
		}

	}
}

func getUser(c *gin.Context) {

	var users []user.User
	db := database.DB
	result := db.Select("*").Find(&users)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": result.Error.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, users)
}
