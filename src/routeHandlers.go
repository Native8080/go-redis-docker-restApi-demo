package main

import (
	"encoding/json"
	"strconv"
	"github.com/gorilla/mux"
)


/*
*	Requuest Handlers
*/
func getAllFeeds(w http.ResponseWriter, req *http.Request){

	var feeds Feeds

	// call to func findAllFeeds() feeds
	feeds := findAllFeeds()
	// do proper error checking & logging / sent proper
	// status codes in json for error
	// Switch to Marshall instead of encode
	if err := json.NewEncoder(w).Encode(feeds); err != nil {
		panic(err)
	}

}


func getFeed(w http.ResponseWriter, req *http.Request){

	params := mux.Vars(req)
	id, err := strconv.Atoi(params["id"])

	

	var feed Feed

	// call findFeed(id int) Feed
	feed := findFeed(id int)

}

func createFeed(w http.ResponseWriter, req *http.Request){
	var feed Feed
	
	params := mux.Vars(req)

	

	// call func createFeed(f Feed)
	createFeed(feed)

}

func updateFeed(w http.ResponseWriter, req *http.Request){
	params := mux.Vars(req)
}

func removeFeed(w http.ResponseWriter, req *http.Request){
	params := mux.Vars(req)
}