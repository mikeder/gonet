package main

import (
  "fmt"
  "log"
  "net/http"
  "net/response"
)

func handler(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("content-type", "application/json")
  resp := &response.BasicResponse{
    Message: "success",
    Code: http.StatusOK,
    Response: r.URL.Path}
  fmt.Fprintln(w, response.JsonResponse(resp))
}

func main()  {
  http.HandleFunc("/", handler) // each request calls handler
	log.Fatal(http.ListenAndServe(":8000", nil))
}
