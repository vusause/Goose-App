package main

import (
	"fmt"
    "log"
    "net/http"

    "github.com/gorilla/mux"
)

var (
	gooseCount = 1
)

func main() {
    router := NewRouter()
    log.Fatal(http.ListenAndServe(":8080", router))
}

func NewRouter() *mux.Router {
	router := mux.NewRouter()
	
	router.HandleFunc("/", defaultHandler)
	router.HandleFunc("/add", addHandler)
	router.HandleFunc("/remove", removeHandler)
	
    return router
}

func defaultHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	gooseImage := "<img src=\"https://www.nps.gov/miss/learn/nature/images/GordonDietzman-20160716-4541.jpg?maxwidth=650&autorotate=false\">"
	htmlPage := "<html><head><title>Page Title</title></head><body>"
	for i := 0; i < gooseCount; i++ {
		htmlPage += gooseImage
	}
	htmlPage += "</body></html>"
	fmt.Fprintf(w, htmlPage)
}

func addHandler(w http.ResponseWriter, r *http.Request) {
	gooseCount++
	w.WriteHeader(200)
}

func removeHandler(w http.ResponseWriter, r *http.Request) {
	if (gooseCount > 0) {
		gooseCount--
	}
	w.WriteHeader(200)
}
