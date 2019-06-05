package main

import (
	"encoding/json"
	"log"
	"net/http"
	"github.com/gorilla/mux"
)


func assembleRequestHandlers() {
	http.handleFunc("/feed/", getAllFeeds)
	http.handleFunc("/feed/{id}", getFeed)
	http.handleFunc("/feed/{id}", createFeed)
	http.handleFunc("/feed/{id}", updateFeed)
	http.handleFunc("/feed/{id}", removeFeed)
}


func main() {

	assembleRequestHandlers()
	log.Fatal(http.ListenAndServe(":8080", nil))
}