package model

import "encoding/json"

// Match holds all the data related to the match between two Superheros.
type Match struct {
	ID                 string `json:"id"`
	SuperheroID        string `json:"superheroId"`
	MatchedSuperheroID string `json:"matchedSuperheroId"`
	CreatedAt          string `json:"createdAt"`
}

// MarshalBinary ...
func (m Match) MarshalBinary() ([]byte, error) {
	return json.Marshal(m)
}

// UnmarshalBinary ...
func (m *Match) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, &m)
}
