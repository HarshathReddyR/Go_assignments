package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func main() {
	http.HandleFunc("/user", GetStudent)
	panic(http.ListenAndServe(":8080", nil))

}

type Student struct {
	// Id       uint64
	Name     string
	Password string
	Email    string
}

var S1 = map[uint64]Student{
	1: {
		Name:     "Abhi",
		Password: "someSecretPassword",
		Email:    "abhi@gmail.com",
	},
	2: {
		Name:     "Giri",
		Password: "someSecretPassword",
		Email:    "giri@gmail.com",
	},
	3: {
		Name:     "Madhan",
		Password: "someSecretPassword",
		Email:    "madhan@gmail.com",
	},
}

func FetchStudent(S1Id uint64) (Student, error) {
	value, ok := S1[S1Id]
	if !ok {
		return Student{}, errors.New("Data is not found")
	}
	return value, nil
}
func GetStudent(rw http.ResponseWriter, rq *http.Request) {
	StudentIdString := rw.URL.Query().Get("S1_id")
	StudentId, err := strconv.ParseUint(StudentIdString, 10, 64)
	if err != nil {
		log.Println("conversion error", err)
		conerr := map[string]string{"msg": "not a valid number"}
		jsonData, err := json.Marshal(conerr)
		if err != nil {
			log.Println("error while converting to json", err)
			rw.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintln(rw, http.StatusText(http.StatusInternalServerError))
			return
		}
		rw.WriteHeader(http.StatusBadRequest)
		rw.Write(jsonData)
		return
	}
	sData, err := FetchStudent(StudentId)
	if err != nil {
		log.Println("fetching data error", err)
		fetcherr := map[string]string{"msg": "error while fetching"}
		jsonData, err := json.Marshal(fetcherr)
		if err != nil {
			log.Println("error while fetchdata conversion to json  ", err)
			rw.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintln(rw, http.StatusText(http.StatusInternalServerError))
			return
		}
		rw.WriteHeader(http.StatusBadRequest)
		rw.Write(jsonData)
		return
	}
	studentData, err := json.Marshal(sData)
	if err != nil {
		log.Println("Error while converting user data to json", err)
		rw.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(rw, http.StatusText(http.StatusInternalServerError))
		return
	}
	rw.Write(studentData)
}
