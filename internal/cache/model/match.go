/*
  Copyright (C) 2019 - 2022 MWSOFT
  This program is free software: you can redistribute it and/or modify
  it under the terms of the GNU General Public License as published by
  the Free Software Foundation, either version 3 of the License, or
  (at your option) any later version.
  This program is distributed in the hope that it will be useful,
  but WITHOUT ANY WARRANTY; without even the implied warranty of
  MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
  GNU General Public License for more details.
  You should have received a copy of the GNU General Public License
  along with this program.  If not, see <http://www.gnu.org/licenses/>.
*/
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
