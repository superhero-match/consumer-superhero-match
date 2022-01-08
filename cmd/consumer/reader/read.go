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
package reader

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"go.uber.org/zap"

	cm "github.com/superhero-match/consumer-superhero-match/internal/cache/model"
	"github.com/superhero-match/consumer-superhero-match/internal/consumer/model"
	dbm "github.com/superhero-match/consumer-superhero-match/internal/db/model"
	fm "github.com/superhero-match/consumer-superhero-match/internal/firebase/model"
)

// Read consumes the Kafka topic and stores the match to DB.
func (r *reader) Read() error {
	ctx := context.Background()

	for {
		fmt.Println("before FetchMessage")
		m, err := r.Consumer.FetchMessage(ctx)
		fmt.Print("after FetchMessage")
		if err != nil {
			r.Logger.Error(
				"failed to fetch message",
				zap.String("err", err.Error()),
				zap.String("time", time.Now().UTC().Format(r.TimeFormat)),
			)

			err = r.Consumer.Close()
			if err != nil {
				r.Logger.Error(
					"failed to close consumer",
					zap.String("err", err.Error()),
					zap.String("time", time.Now().UTC().Format(r.TimeFormat)),
				)

				return err
			}

			return err
		}

		fmt.Printf(
			"message at topic/partition/offset \n%v/\n%v/\n%v: \n%s = \n%s\n",
			m.Topic,
			m.Partition,
			m.Offset,
			string(m.Key),
			string(m.Value),
		)
		fmt.Print()

		var match model.Match

		if err := json.Unmarshal(m.Value, &match); err != nil {
			r.Logger.Error(
				"failed to unmarshal JSON into consumer Match model",
				zap.String("err", err.Error()),
				zap.String("time", time.Now().UTC().Format(r.TimeFormat)),
			)

			err = r.Consumer.Close()
			if err != nil {
				r.Logger.Error(
					"failed to close consumer",
					zap.String("err", err.Error()),
					zap.String("time", time.Now().UTC().Format(r.TimeFormat)),
				)

				return err
			}

			return err
		}

		err = r.DB.StoreMatch(dbm.Match{
			ID:                 match.ID,
			SuperheroID:        match.SuperheroID,
			MatchedSuperheroID: match.MatchedSuperheroID,
			CreatedAt:          match.CreatedAt,
		})
		if err != nil {
			r.Logger.Error(
				"failed to store match in database",
				zap.String("err", err.Error()),
				zap.String("time", time.Now().UTC().Format(r.TimeFormat)),
			)

			err = r.Consumer.Close()
			if err != nil {
				r.Logger.Error(
					"failed to close consumer",
					zap.String("err", err.Error()),
					zap.String("time", time.Now().UTC().Format(r.TimeFormat)),
				)

				return err
			}

			return err
		}

		// Construct cache keys to be deleted after the match has occurred.
		keys := make([]string, 0)
		keys = append(keys, fmt.Sprintf(r.ChoiceKeyFormat, match.SuperheroID, match.MatchedSuperheroID))
		keys = append(keys, fmt.Sprintf(r.ChoiceKeyFormat, match.MatchedSuperheroID, match.SuperheroID))

		// Delete likes form the cache as the two users have matched.
		err = r.Cache.DeleteChoice(keys)
		if err != nil {
			r.Logger.Error(
				"failed to delete choice from cache",
				zap.String("err", err.Error()),
				zap.String("time", time.Now().UTC().Format(r.TimeFormat)),
			)

			err = r.Consumer.Close()
			if err != nil {
				r.Logger.Error(
					"failed to close consumer",
					zap.String("err", err.Error()),
					zap.String("time", time.Now().UTC().Format(r.TimeFormat)),
				)

				return err
			}

			return err
		}

		token, err := r.Cache.GetFirebaseMessagingToken(fmt.Sprintf(r.TokenKeyFormat, match.MatchedSuperheroID))
		if err != nil || token == nil {
			r.Logger.Error(
				"failed to fetch Firebase messaging token from cache",
				zap.String("err", err.Error()),
				zap.String("time", time.Now().UTC().Format(r.TimeFormat)),
			)

			err = r.Consumer.Close()
			if err != nil {
				r.Logger.Error(
					"failed to close consumer",
					zap.String("err", err.Error()),
					zap.String("time", time.Now().UTC().Format(r.TimeFormat)),
				)

				return err
			}

			return err
		}

		err = r.Firebase.PushNewMatchNotification(fm.Request{
			Token:       token.Token,
			SuperheroID: match.SuperheroID,
		})
		if err != nil {
			r.Logger.Error(
				"failed to push new match notification to Firebase",
				zap.String("err", err.Error()),
				zap.String("time", time.Now().UTC().Format(r.TimeFormat)),
			)

			err = r.Consumer.Close()
			if err != nil {
				r.Logger.Error(
					"failed to close consumer",
					zap.String("err", err.Error()),
					zap.String("time", time.Now().UTC().Format(r.TimeFormat)),
				)

				return err
			}

			return err
		}

		err = r.Cache.SetMatch(fmt.Sprintf(r.MatchKeyFormat, match.SuperheroID, match.MatchedSuperheroID), cm.Match{
			ID:                 match.ID,
			SuperheroID:        match.SuperheroID,
			MatchedSuperheroID: match.MatchedSuperheroID,
			CreatedAt:          match.CreatedAt,
		})
		if err != nil {
			r.Logger.Error(
				"failed to store match in cache",
				zap.String("err", err.Error()),
				zap.String("time", time.Now().UTC().Format(r.TimeFormat)),
			)

			err = r.Consumer.Close()
			if err != nil {
				r.Logger.Error(
					"failed to close consumer",
					zap.String("err", err.Error()),
					zap.String("time", time.Now().UTC().Format(r.TimeFormat)),
				)

				return err
			}

			return err
		}

		err = r.Consumer.CommitMessages(ctx, m)
		if err != nil {
			r.Logger.Error(
				"failed to commit messages",
				zap.String("err", err.Error()),
				zap.String("time", time.Now().UTC().Format(r.TimeFormat)),
			)

			err = r.Consumer.Close()
			if err != nil {
				r.Logger.Error(
					"failed to close consumer",
					zap.String("err", err.Error()),
					zap.String("time", time.Now().UTC().Format(r.TimeFormat)),
				)

				return err
			}

			return err
		}
	}
}
