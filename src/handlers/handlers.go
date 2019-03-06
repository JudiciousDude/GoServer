package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
	"strconv"
	"strings"

	"../dbObjects"
)

type Handler struct {
	res []dbObjects.Resourse
	DB  *sql.DB
}

var (
	strictChecker *regexp.Regexp
	textChecker   *regexp.Regexp
)

func init() {
	strictChecker = regexp.MustCompile("[\\W]")
	textChecker = regexp.MustCompile("['|\"](\\s*)$")
}

func (handler Handler) HandleIndex(writer http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodGet:
		file, err := ioutil.ReadFile("index.html")
		if err != nil {
			log.Println(err)
			writer.WriteHeader(http.StatusInternalServerError)
			return
		}

		writer.WriteHeader(200)
		writer.Write(file)

	case http.MethodPost:
		body := json.NewDecoder(req.Body)
		res := dbObjects.Resourse{}

		err := body.Decode(&res)
		if err != nil {
			writer.WriteHeader(http.StatusBadRequest)
			log.Println("HandleIndex:", err.Error())
			return
		}

		if strictChecker.MatchString(res.Name) {
			writer.WriteHeader(http.StatusForbidden)
			return
		}

		if strings.ContainsRune(res.Conditions, ';') || textChecker.MatchString(res.Conditions) {
			writer.WriteHeader(http.StatusForbidden)
			return
		}

		err = res.AddToDb(handler.DB)
		if err != nil {
			fmt.Println(err)
			writer.WriteHeader(http.StatusInternalServerError)
			return
		}

		writer.WriteHeader(200)
	}

	log.Println("HandleIndex caught:", req.URL.Path, "with method", req.Method)
}

func (handler Handler) HandleGetList(writer http.ResponseWriter, req *http.Request) {
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

	log.Println("HandleGetList caught:", req.URL.Path, "with method", req.Method)
}

func (handler Handler) HandleDelete(writer http.ResponseWriter, req *http.Request) {
	number := strings.TrimPrefix(req.URL.Path, "/delete/")

	if _, ok := strconv.Atoi(number); ok != nil {
		writer.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	_, err := handler.DB.Exec("DELETE FROM Stock WHERE id=?", number)
	if err != nil {
		writer.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	writer.WriteHeader(http.StatusOK)

	log.Println("HandleDelete caught:", req.URL.Path, "with method", req.Method)
}
