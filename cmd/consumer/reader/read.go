package reader

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/consumer-superhero-match/internal/consumer/model"
	dbm "github.com/consumer-superhero-match/internal/db/model"
)

// Read consumes the Kafka topic and stores the match to DB.
func (r *Reader) Read() error {
	ctx := context.Background()

	for {
		fmt.Print("before FetchMessage")
		m, err := r.Consumer.Consumer.FetchMessage(ctx)
		fmt.Print("after FetchMessage")
		if err != nil {
			err = r.Consumer.Consumer.Close()
			if err != nil {
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

		var match model.Match

		if err := json.Unmarshal(m.Value, &match); err != nil {
			_ = r.Consumer.Consumer.Close()
			if err != nil {
				fmt.Println("Unmarshal")
				fmt.Println(err)
				err = r.Consumer.Consumer.Close()
				if err != nil {
					return err
				}

				return err
			}
		}

		err = r.DB.StoreMatch(dbm.Match{
			ID:                 match.ID,
			SuperheroID:        match.SuperheroID,
			MatchedSuperheroID: match.MatchedSuperheroID,
			CreatedAt:          match.CreatedAt,
		}, )
		if err != nil {
			fmt.Println("DB")
			fmt.Println(err)
			err = r.Consumer.Consumer.Close()
			if err != nil {
				return err
			}

			return err
		}

		//keys := make([]string, 0)
		//keys = append(keys, fmt.Sprintf("%s.%s", match.SuperheroID, match.MatchedSuperheroID))
		//keys = append(keys, fmt.Sprintf("%s.%s", match.MatchedSuperheroID, match.SuperheroID))
		//
		//err = r.Cache.DeleteChoice(keys)
		//if err != nil {
		//	fmt.Println("Cache")
		//	fmt.Println(err)
		//	err = r.Consumer.Consumer.Close()
		//	if err != nil {
		//		return err
		//	}
		//
		//	return err
		//}

		err = r.Consumer.Consumer.CommitMessages(ctx, m)
		if err != nil {
			err = r.Consumer.Consumer.Close()
			if err != nil {
				return err
			}

			return err
		}
	}
}
