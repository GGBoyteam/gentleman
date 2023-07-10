// Package routes 注册路由
package routes

import (
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
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

type User struct {
	ID     uint   `gorm:"primarykey"`
	QQ     string `json:"qq"`
	Title  string `json:"title"`
	Score  int    `json:"score"`
	Status string `json:"status"`
}

func (User) TableName() string {
	return "user"
}

var db *gorm.DB

func getUser(c *gin.Context) {
	// Replace the connection details with your MySQL database configuration
	//dsn := "root:wxl7756212@tcp(localhost:3306)/leetcode" # 服务器上的代码
	dsn := "root:123456@tcp(localhost:3306)/leetcode"
	var err error
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Println("connect failed")
		log.Fatal(err)
	}
	var users []User
	result := db.Select("id, qq, title, score, status").Find(&users)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": result.Error.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, users)
}
