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

	subscribers := findAllFeeds()

	if err := json.NewEncoder(w).Encode(subscribers); err != nil {
		panic(err)
	}

}


func getFeed(w http.ResponseWriter, req *http.Request){

	params := mux.Vars(req)
	sub, err := params["subscriber"]
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		panic(err)
	}

	subscrbr := findFeed(sub, id)

	if err := json.NewEncoder(w).Encode(subscrbr); err != nil {
		panic(err)
	}

}

func postFeed(w http.ResponseWriter, req *http.Request){
	var subscriber Subscriber
	
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}
	
	defer client.Body.Close()

	if err := json.UnMarshal(body, &subscriber); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422)
		
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)		
		}

	}

	createFeed(subscriber)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
}


func updateFeed(w http.ResponseWriter, req *http.Request){
	params := mux.Vars(req)
}

func removeFeed(w http.ResponseWriter, req *http.Request){
	params := mux.Vars(req)
}