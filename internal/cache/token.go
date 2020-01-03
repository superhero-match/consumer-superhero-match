package cache

import (
	"github.com/consumer-superhero-match/internal/cache/model"
	"github.com/go-redis/redis"
)

// GetFirebaseMessagingToken fetches choice(like, dislikes are only in DB) from cache.
func (c *Cache) GetFirebaseMessagingToken(key string) (*model.FirebaseMessagingToken, error) {
	res, err := c.Redis.Get(key).Result()
	if err != nil && err != redis.Nil {
		return nil, err
	}

	if len(res) == 0 {
		return nil, nil
	}

	var token model.FirebaseMessagingToken

	if err := token.UnmarshalBinary([]byte(res)); err != nil {
		return nil, err
	}

	return &token, nil
}
