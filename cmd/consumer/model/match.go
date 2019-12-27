package model

// Match holds all the data related to the match between two Superheros.
type Match struct {
	ID                 string `json:"id"`
	SuperheroID        string `json:"superheroId"`
	MatchedSuperheroID string `json:"matchedSuperheroId"`
	CreatedAt          string `json:"createdAt"`
}