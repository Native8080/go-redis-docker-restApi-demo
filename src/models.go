package main


/*
*	Subscriber Model
*/
type Subscriber struct {
	Username 		string  `json:"username,omitempty"`
	Id				int		`json:"id,omitempty"`
	Feed			*Feed 	`json:"feed,omitempty"`
}

/*
*	Feed Model
*/
type Feed struct {
	Id			int		`json:"id,omitempty`
	Publisher	string	`json:"publisher,omitempty"`
	Title		string	`json:"title,omitempty"`
	Url			string	`json:"url,omitempty"`
	Topic		string	`json:topic,omitempty`
}

type Subscribers		[]Subscriber
type Feeds 				[]Feed
