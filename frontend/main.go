package main

import (
	"net/http"
	. "ongkyle.com/reading-list/frontend/spa"
)


func main() {
    http.ListenAndServe(":8080", SpaHandler("public", "index.html"))
}