package main

import (
	"encoding/json"
	"log"
	"net/http"
	"github.com/gorilla/mux"
)

/*
*	Utility Functions
*/


/*
*
*/
func assembleRequestHandlers() {
	http.handleFunc("/feed/all", getAllFeeds)
	http.handleFunc("/feed/{id}", getFeed)
	http.handleFunc("/feed/create", createFeed)
	http.handleFunc("/feed/update", updateFeed)
	http.handleFunc("/feed/remove", removeFeed)
}


func main() {
	/*
	*
	*/
	assembleRequestHandlers()
	log.Fatal(http.ListenAndServe(":8080", nil))
}