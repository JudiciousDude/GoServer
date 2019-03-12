package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"./src/handlers"

	_ "github.com/go-sql-driver/mysql"
)

var (
	db      *sql.DB
	handler handlers.Handler
	config  Config
)

type Config struct {
	User       string `json:"user"`
	Password   string `json:"password"`
	ConnType   string `json:"conntype"`
	Path       string `json:"dbpath"`
	Dbport     string `json:"dbport"`
	Dbname     string `json:"dbname"`
	Serverport string `json:"serverport"`
}

//Connects to database
func dbConnect() {
	var err error
	pathStr := config.User
	if config.Password != "" {
		pathStr += ":" + config.Password
	}
	pathStr += "@" + config.ConnType + "(" + config.Path + ":" + config.Dbport + ")/" + config.Dbname + "?charset=utf8"
	db, err = sql.Open("mysql", pathStr)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = db.Ping()
	if err != nil {
		fmt.Println(err)
		return
	}
}

func main() {
	file, err := ioutil.ReadFile("config.json")
	if err != nil {
		fmt.Println(err)
		return
	}
	json.Unmarshal(file, &config)
	dbConnect()
	handler.DB = db

	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("css/"))))
	http.Handle("/js/", http.StripPrefix("/js/", http.FileServer(http.Dir("js/"))))

	http.HandleFunc("/", handler.HandleIndex)

	http.HandleFunc("/getlist", handler.HandleGetList)
	http.HandleFunc("/delete/", handler.HandleDelete)
	http.ListenAndServe(":"+config.Serverport, nil)

}
