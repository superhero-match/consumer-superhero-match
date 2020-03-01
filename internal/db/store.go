/*
  Copyright (C) 2019 - 2020 MWSOFT
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
package db

import (
	"github.com/consumer-superhero-match/internal/db/model"
)

// StoreMatch saves new match.
func(db *DB) StoreMatch (m model.Match) error {
	_, err := db.stmtInsertNewMatch.Exec(
		m.ID,
		m.SuperheroID,
		m.MatchedSuperheroID,
		m.CreatedAt,
	)
	if err != nil {
		return err
	}

	return nil
}
