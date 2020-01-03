package cache

import (
	"github.com/go-redis/redis"
)

// DeleteChoice deletes choice from Redis cache because Superheros have matched.
func (c *Cache) DeleteChoice(keys []string) error {
	err := c.Redis.Del(keys...).Err()
	if err != nil && err != redis.Nil {
		return err
	}

	return nil
}