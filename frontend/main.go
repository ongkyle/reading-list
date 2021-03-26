package main

import (
	"net/http"
	"fmt"
	"log"
	"time"

	. "ongkyle.com/reading-list/frontend/spa"

    "github.com/gorilla/mux"
)


func main() {
	r := mux.NewRouter()
    // It's important that this is before your catch-all route ("/")
    api := r.PathPrefix("/api/").Subrouter()
	api.HandleFunc("/readinglist", serveRequests)
	fmt.Println("Finished setting up routers")

	staticFiles := http.FileServer(http.Dir("./public/static"))
	r.PathPrefix("/js/").Handler(staticFiles)
	r.PathPrefix("/css/").Handler(staticFiles)
	fmt.Println("Finished setting up static content")

	spa := SpaHandler("public", "index.html")
	r.PathPrefix("/").Handler(spa)
	fmt.Println("Finished setting up spa")

	srv := &http.Server{
		Handler: r,
		Addr:    "127.0.0.1:8000",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Fatal(srv.ListenAndServe())
	fmt.Println("Listening on port 8080...")
}

func serveRequests(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	switch r.Method {
	case "GET":
		serveGet(w, r)
	case "POST":
		servePost(w, r)
	default:
        w.WriteHeader(http.StatusNotFound)
        w.Write([]byte(`{"message": "not found"}`))
	}
}

func serveGet(w http.ResponseWriter, r *http.Request){
	w.WriteHeader(http.StatusOK)
	log.Println()
	log.Println(w)
	log.Println(r)
	log.Println()

}

func servePost(w http.ResponseWriter, r *http.Request){
	w.WriteHeader(http.StatusCreated)
	log.Println()
	log.Println(w)
	log.Println(r)
	log.Println()

} 