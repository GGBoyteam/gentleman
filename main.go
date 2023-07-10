package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

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

func getUsers(c *gin.Context) {
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

func index(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}

func main() {
	// Replace the connection details with your MySQL database configuration
	//dsn := "root:wxl7756212@tcp(localhost:3306)/leetcode" # 服务器上的代码

	dsn := "root:123456@tcp(localhost:3306)/leetcode"
	var err error
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Println("connect failed")
		log.Fatal(err)
	}

	// new 一个 Gin Engine 实例
	r := gin.Default()
	// 注册一个路由
	r.LoadHTMLGlob("templates/*")
	r.GET("/", index)

	r.GET("/users", getUsers)

	err = r.Run(":9090")
	if err != nil {
		panic(err)
	}
}
