package model

// Match holds all the data related to the match between two Superheros.
type Match struct {
	ID                 string `json:"id"`
	SuperheroID        string `json:"superhero_id"`
	MatchedSuperheroID string `json:"matched_superhero_id"`
	CreatedAt          string `json:"created_at"`
}
