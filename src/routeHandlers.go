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
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

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
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(req)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		panic(err)
	}
	


	var feed Feed

	// call findFeed(id int) Feed
	feed := findFeed(id int)

}

func createFeed(w http.ResponseWriter, req *http.Request){
	var feed Feed
	
	//params := mux.Vars(req)
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))

	

	// call func createFeed(f Feed)
	createFeed(feed)

}

func updateFeed(w http.ResponseWriter, req *http.Request){
	params := mux.Vars(req)
}

func removeFeed(w http.ResponseWriter, req *http.Request){
	params := mux.Vars(req)
}