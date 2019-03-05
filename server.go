package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"
	"strconv"

	"./src/handlers"

	_ "github.com/go-sql-driver/mysql"
)

var (
	db      *sql.DB
	handler handlers.Handler
)

func dbConnect(status chan<- bool) {
	var err error
	db, err = sql.Open("mysql", "root:root@tcp(localhost:3302)/test?charset=utf8")
	if err != nil {
		fmt.Println(err)
		status <- false
		return
	}

	err = db.Ping()
	if err != nil {
		fmt.Println(err)
		status <- false
		return
	}

	status <- true
}

func main() {
	params := os.Args[1:]
	port := ":8081"
	status := make(chan bool)

	go dbConnect(status)

	if len(params) > 0 {
		if _, ok := strconv.Atoi(params[0]); ok != nil {
			println("Given arg is not a number")
			fmt.Scanln()
			return
		}
		port = ":" + params[0]
	}

	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("css/"))))
	http.Handle("/js/", http.StripPrefix("/js/", http.FileServer(http.Dir("js/"))))

	handler = handlers.Handler{}
	if !<-status {
		println("Something wrong")
		return
	}

	handler.DB = db

	http.HandleFunc("/", handler.HandleIndex)

	http.HandleFunc("/getlist", handler.HandleGetList)

	http.ListenAndServe(port, nil)

}
