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
    //#http.HandleFunc("/", root)
    //#http.HandleFunc("/sleep", sleep)
    //#http.HandleFunc("/status", status)
    //#log.Fatal(http.ListenAndServe(":8000", nil))
    request()
}

func request() {
    resp, err := http.Get("https://google.com")
    if err != nil {
        log.Fatal(err)
    }
    fmt.Print("Content Length: ")
    fmt.Print(resp.ContentLength)
    fmt.Print("\n")
    fmt.Print("TLS Version: ")
    fmt.Print(resp.TLS.Version)
    fmt.Print("\n")
    fmt.Print("Handshake Complete: ")
    fmt.Print(resp.TLS.HandshakeComplete)
    fmt.Print("\n")
    fmt.Print("Did Resume: ")
    fmt.Print(resp.TLS.DidResume)
    fmt.Print("\n")
    fmt.Print("Cipher Suite: ")
    fmt.Print(resp.TLS.CipherSuite)
    fmt.Print("\n")
    fmt.Print("Peer Certificates: ")
    fmt.Print(resp.TLS.PeerCertificates)
    fmt.Print("\n")
    fmt.Print("Verified Chains: ")
    fmt.Print(resp.TLS.VerifiedChains)
    fmt.Print("\n")
}
