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
package firebase

import (
	"github.com/superhero-match/consumer-superhero-match/internal/config"
	"github.com/superhero-match/consumer-superhero-match/internal/firebase/model"
)

// Firebase interface defines Firebase methods.
type Firebase interface {
	PushNewMatchNotification(req model.Request) error
}

// firebase holds all the Firebase related data.
type firebase struct {
	FunctionAddress string
	ContentType     string
}

// NewFirebase creates new value of type Firebase.
func NewFirebase(cfg *config.Config) Firebase {
	return &firebase{
		FunctionAddress: cfg.Firebase.FunctionAddress,
		ContentType:     cfg.Firebase.ContentType,
	}
}
