package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func getHome(res http.ResponseWriter, req *http.Request) {
	fmt.Println("Inside getHome route", req.Method)
	switch req.Method {
	case "GET":
		result := "Home Page\n"
		io.WriteString(res, result)
		return
	case "POST":
		if err := req.ParseForm(); err != nil {
			fmt.Fprintf(res, "ParseForm() err: %v", err)
			return
		}
		fmt.Fprintf(res, "Post from website! r.PostFrom = %v\n", req.PostForm)
		io.WriteString(res, req.FormValue("name"))
		return
	default:
		result := "Only suppoet Get and Post method"
		fmt.Println(result)
		io.WriteString(res, result)
		return
	}
}

func main() {
	fmt.Println("Starting server")
	http.HandleFunc("/", getHome)
	err := http.ListenAndServe(":3333", nil)
	if err != nil {
		log.Fatal(err)
	}
}
