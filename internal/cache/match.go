package cache

import (
	"fmt"
	"github.com/consumer-superhero-match/internal/cache/model"
)

// SetMatch stores matches into Redis cache.
func (c *Cache) SetMatch(match model.Match) error {
	err := c.Redis.Set(
		fmt.Sprintf("match.%s.%s", match.SuperheroID, match.MatchedSuperheroID),
		match,
		0,
	).Err()
	if err != nil {
		return err
	}

	return nil
}