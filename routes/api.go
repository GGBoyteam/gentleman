// Package routes 注册路由
package routes

import (
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
			c.HTML(http.StatusOK, "index.html", nil)
		})

		v1.GET("/users", getUser)

	}
}

func getUser(c *gin.Context) {

	var users []user.User
	db := database.DB
	result := db.Select("id, qq, title, score, status").Find(&users)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": result.Error.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, users)
}
