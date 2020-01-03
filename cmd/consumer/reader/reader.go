package reader

import (
	"github.com/consumer-superhero-match/internal/cache"
	"github.com/consumer-superhero-match/internal/config"
	"github.com/consumer-superhero-match/internal/consumer"
	"github.com/consumer-superhero-match/internal/db"
	"github.com/consumer-superhero-match/internal/firebase"
)

// Reader holds all the data relevant.
type Reader struct {
	DB       *db.DB
	Consumer *consumer.Consumer
	Cache    *cache.Cache
	Firebase *firebase.Firebase
}

// NewReader configures Reader.
func NewReader(cfg *config.Config) (r *Reader, err error) {
	dbs, err := db.NewDB(cfg)
	if err != nil {
		return nil, err
	}

	ch, err := cache.NewCache(cfg)
	if err != nil {
		return nil, err
	}

	cs := consumer.NewConsumer(cfg)

	f := firebase.NewFirebase(cfg)

	return &Reader{
		DB:       dbs,
		Consumer: cs,
		Cache:    ch,
		Firebase: f,
	}, nil
}
