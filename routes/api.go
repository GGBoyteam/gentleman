// Package routes 注册路由
package routes

import (
	"gentleman/app/http/controllers/api/v1/auth"
	"gentleman/app/http/middlewares"
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

		v1.Use(middlewares.LimitIP("60-H"))
		// 注册一个路由
		v1.GET("/", func(c *gin.Context) {
			c.HTML(http.StatusOK, "homepage.html", nil)
		})

		v1.GET("/users", getUser)

		v1.GET("/query_table", func(c *gin.Context) {
			c.HTML(http.StatusOK, "query.html", nil)
		})

		v1.GET("/home", func(c *gin.Context) {
			c.HTML(http.StatusOK, "index.html", nil)
		})

		v1.GET("/fly_table", func(c *gin.Context) {
			c.HTML(http.StatusOK, "fly.html", nil)
		})

		v1.GET("/signup", func(c *gin.Context) {
			c.HTML(http.StatusOK, "signup.html", nil)
		})

		v1.GET("/login", func(c *gin.Context) {
			c.HTML(http.StatusOK, "login.html", nil)
		})

		//v1.GET("/gentleman", func(c *gin.Context) {
		//	c.HTML(http.StatusOK, "gentleman.html", nil)
		//})

		v1.GET("/gostudy", func(c *gin.Context) {
			c.HTML(http.StatusOK, "gostudy.html", nil)
		})

		v1.GET("/country", func(c *gin.Context) {
			c.HTML(http.StatusOK, "country.html", nil)
		})

		v1.GET("/dayPromblem", func(c *gin.Context) {
			c.HTML(http.StatusOK, "dayPromblem.html", nil)
		})

		v1.GET("/weekPromblem", func(c *gin.Context) {
			c.HTML(http.StatusOK, "weekPromblem.html", nil)
		})

		authGroup := v1.Group("/auth")
		authGroup.Use(middlewares.LimitIP("60-H"))
		{
			// 登录
			lgc := new(auth.LoginController)
			authGroup.POST("/login/using-phone", middlewares.GuestJWT(), lgc.LoginByPhone)
			authGroup.POST("/login/using-password", middlewares.GuestJWT(), lgc.LoginByPassword)
			authGroup.POST("/login/refresh-token", middlewares.AuthJWT(), lgc.RefreshToken)

			// 注册
			suc := new(auth.SignupController)
			authGroup.POST("/signup/email/exist", middlewares.GuestJWT(), suc.IsEmailExist)
			authGroup.POST("/signup/using-email", middlewares.GuestJWT(), suc.SignupUsingEmail)
			authGroup.POST("/signup/qq/exist", middlewares.GuestJWT(), middlewares.LimitPerRoute("60-H"), suc.IsQQExist)
			authGroup.POST("/signup/using-qq", middlewares.GuestJWT(), suc.SignupUsingQQ)
			authGroup.GET("/isLogin", middlewares.GuestJWT(), suc.IsLogin)

			// 发送验证码
			vcc := new(auth.VerifyCodeController)
			authGroup.POST("/verify-codes/email", vcc.SendUsingEmail)
			authGroup.POST("/verify-codes/qq", vcc.SendUsingQQ)
			// 图片验证码,加限流
			authGroup.POST("/verify-codes/captcha", middlewares.LimitPerRoute("50-H"), vcc.ShowCaptcha)

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
