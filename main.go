package main

import (
	"fmt"
	"log"
	"net/http"
	"net/response"
	"time"
)

func root(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "root")
}

func sleep(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "sleep")
	time.Sleep(5 * time.Second)
}

func status(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	resp := response.BasicResponse{
		Message:  "success",
		Code:     http.StatusOK,
		Response: r.URL.Path}
	fmt.Fprintln(w, response.JsonResponse(&resp))
}

func main() {
	http.HandleFunc("/", root)
	http.HandleFunc("/sleep", sleep)
	http.HandleFunc("/status", status)
	log.Fatal(http.ListenAndServe(":8000", nil))
}
