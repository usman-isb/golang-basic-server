package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Contact struct {
	Name string `json:Name`
	Age  int64  `json:Age`
}

func getContacts(res http.ResponseWriter, req *http.Request) {
	fmt.Println("Inside index method")
	result := "Get Contact method\n"
	io.WriteString(res, result)
}

func createContact(res http.ResponseWriter, req *http.Request) {
	reqBody, _ := ioutil.ReadAll(req.Body)
	var contact Contact
	json.Unmarshal(reqBody, &contact)
	json.NewEncoder(res).Encode(contact)
	newContact, err := json.Marshal(reqBody)
	if err != nil {
		log.Fatal(err)
		fmt.Println(err)
	} else {
		fmt.Println(string(newContact))
	}
}

func main() {
	fmt.Println("Server starting")
	r := mux.NewRouter().StrictSlash(true)

	// routes
	r.HandleFunc("/getContacts", getContacts).Methods("GET")
	r.HandleFunc("/createContact", createContact).Methods("POST")

	// start server
	err := http.ListenAndServe(":3333", r)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("Server started")
	}
}
