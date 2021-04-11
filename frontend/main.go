package main

import (
	"net/http"
	"log"
	"time"
	"os"
	"encoding/json"

	. "ongkyle.com/reading-list/frontend/spa"
	. "ongkyle.com/reading-list/frontend/services"
	. "ongkyle.com/reading-list/common"

	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
)

var store = sessions.NewCookieStore([]byte(os.Getenv("psuedo_random")))

func main() {
	r := mux.NewRouter()
    // It's important that this is before your catch-all route ("/")
    api := r.PathPrefix("/api/").Subrouter()
	api.HandleFunc("/readinglist", serveRequests)
	log.Println("Finished setting up routers")

	staticFiles := http.FileServer(http.Dir("./public/static"))
	r.PathPrefix("/js/").Handler(staticFiles)
	r.PathPrefix("/css/").Handler(staticFiles)
	log.Println("Finished setting up static content")

	spa := SpaHandler("public", "index.html")
	r.PathPrefix("/").Handler(spa)
	log.Println("Finished setting up spa")

	srv := &http.Server{
		Handler: r,
		Addr:    "127.0.0.1:8080",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Println("Listening on port 8080...")
	log.Fatalln(srv.ListenAndServe())
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
	w.Header().Set("Content-Type", "application/json")
	session, _ := store.Get(r, "session-name")
	log.Println(session.ID)
	response := NewDataService("backend:8888").Get(session.ID)
	w.WriteHeader(response.Code)
	json.NewEncoder(w).Encode(response)

	log.Println(r)
	log.Println(w)
	log.Println(response)

}

func servePost(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	session, _ := store.Get(r, "session-name")
	var items []Item
	err := json.NewDecoder(r.Body).Decode(&items)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	log.Println("Parsed these items...")
	log.Println(items)
	response := NewDataService("backend:8888").Save(session.ID, items)
	w.WriteHeader(response.Code)

	log.Println(r)
	log.Println(w)
	log.Println(response)
}

