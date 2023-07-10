package main

import (
	"database/sql"
	"encoding/json"
	"html/template"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	QQ     string `json:"qq"`
	Title  string `json:"title"`
	Score  int    `json:"score"`
	Status string `json:"status"`
}

var db *sql.DB

func getUsers(w http.ResponseWriter, r *http.Request) {
	rows, err := db.Query("SELECT qq, title, score, status FROM user")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var user User
		err := rows.Scan(&user.QQ, &user.Title, &user.Score, &user.Status)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		users = append(users, user)
	}
	err = rows.Err()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response, err := json.Marshal(users)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}

func index(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "index.html", nil)
}

func renderTemplate(w http.ResponseWriter, tmpl string, data interface{}) {
	t, err := template.ParseFiles(tmpl)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = t.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func main() {
	// Replace the connection details with your MySQL database configuration
	//dsn := "root:wxl7756212@tcp(localhost:3306)/leetcode" # 服务器上的代码

	dsn := "root:123456@tcp(localhost:3306)/leetcode"
	var err error
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Println("connect failed")
		log.Fatal(err)
	}
	defer db.Close()
	http.HandleFunc("/", index)
	http.HandleFunc("/users", getUsers)

	log.Println("Server started on port 9090")
	log.Fatal(http.ListenAndServe(":9090", nil))
}
