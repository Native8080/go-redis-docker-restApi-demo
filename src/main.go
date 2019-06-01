package main

import (
	"encoding/json"
	"log"
	"net/http"
	"github.com/gorilla/mux"
)

type Subscriber struct {
	id				int		`json:"id,omitempty"`
	username 		string  `json:"username,omitempty"`
	name 			string  `json:"name,omitempty"`
	
}

type Feed struct {
	id				int		`json:"id,omitempty`
	publisher		string	`json:"publisher,omitempty"`
	title			string	`json:"title,omitempty"`
	url				string	`json:"url,omitempty"`
	topic			string	`json:topic,omitempty`
}

var Subscribers		[]Subscriber
var Feeds 			[]Feed



/*
*	Utility Functions
*/


/*
*	Requuest Handlers
*/
func getAllFeeds(w http.ResponseWriter, r *http.Request){

}


func getFeed(w http.ResponseWriter, r *http.Request){

}

func createFeed(w http.ResponseWriter, r *http.Request){

}

func updateFeed(w http.ResponseWriter, r *http.Request){

}

func removeFeed(w http.ResponseWriter, r *http.Request){

}

/*
*
*/
func assembleRequestHandlers() {
	http.handleFunc("/feed/all", getAllFeeds)
	http.handleFunc("/feed/{id}", getFeed)
	http.handleFunc("/feed/create", createFeed)
	http.handleFunc("/feed/update", updateFeed)
	http.handleFunc("/feed/remove", removeFeed)
	//log.Fatal(http.ListenAndServe(":8080", nil))
}


func main() {
	/*
	*
	*/
	assembleRequestHandlers()
	log.Fatal(http.ListenAndServe(":8080", nil))
}