package model

// Request holds new match request body data that is going to be sent to Firebase.
// This request will trigger Firebase function that will notify the matched user that
// they got new match.
type Request struct {
	Token       string `json:"token"`
	SuperheroID string `json:"superheroId"`
}
