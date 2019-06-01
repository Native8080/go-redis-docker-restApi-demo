package main


/*
*	Subscriber Model
*/
type Subscriber struct {
	id				int		`json:"id,omitempty"`
	username 		string  `json:"username,omitempty"`
	name 			string  `json:"name,omitempty"`
	
}

/*
*	Feed Model
*/
type Feed struct {
	id				int		`json:"id,omitempty`
	publisher		string	`json:"publisher,omitempty"`
	title			string	`json:"title,omitempty"`
	url				string	`json:"url,omitempty"`
	topic			string	`json:topic,omitempty`
}

var Subscribers		[]Subscriber
var Feeds 			[]Feed
