package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"../dbObjects"
)

type Handler struct {
	res []dbObjects.Resourse
	DB  *sql.DB
}

func (handler Handler) HandleIndex(writer http.ResponseWriter, req *http.Request) {
	fmt.Println(req.URL.Path)

	file, err := ioutil.ReadFile("index.html")
	if err != nil {
		log.Println(err)
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	writer.WriteHeader(200)
	writer.Write(file)
}

func (handler Handler) HandleGetList(writer http.ResponseWriter, req *http.Request) {
	fmt.Println(req.URL.Path)
	tmp := dbObjects.Resourse{}

	res, err := tmp.LoadListFromDB(handler.DB)
	if err != nil {
		log.Println(err)
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	jsoned, err := json.Marshal(res)
	if err != nil {
		log.Println(err)
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	writer.WriteHeader(http.StatusOK)
	writer.Header().Set("Content-Type", "application/json")
	writer.Write(jsoned)
}
