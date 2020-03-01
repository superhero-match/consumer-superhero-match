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
