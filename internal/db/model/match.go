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

// Match holds all the data related to the match between two Superheros.
type Match struct {
	ID                 string `json:"id"`
	SuperheroID        string `json:"superhero_id"`
	MatchedSuperheroID string `json:"matched_superhero_id"`
	CreatedAt          string `json:"created_at"`
}
