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
