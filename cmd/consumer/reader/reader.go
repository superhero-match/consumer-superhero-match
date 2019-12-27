package reader

import (
	"github.com/consumer-superhero-match/internal/config"
	"github.com/consumer-superhero-match/internal/consumer"
	"github.com/consumer-superhero-match/internal/db"
)

// Reader holds all the data relevant.
type Reader struct {
	DB       *db.DB
	Consumer *consumer.Consumer
}

// NewReader configures Reader.
func NewReader(cfg *config.Config) (r *Reader, err error) {
	dbs, err := db.NewDB(cfg)
	if err != nil {
		return nil, err
	}

	cs := consumer.NewConsumer(cfg)

	return &Reader{
		DB:       dbs,
		Consumer: cs,
	}, nil
}