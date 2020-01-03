package firebase

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/consumer-superhero-match/internal/firebase/model"
)

// PushNewMatchNotification pushes new match notification to Firebase.
func (f *Firebase) PushNewMatchNotification(req model.Request) error {
	requestBody, err := json.Marshal(map[string]string{
		"token":       req.Token,
		"superheroId": req.SuperheroID,
	})
	if err != nil {
		return err
	}

	resp, err := http.Post(
		f.FunctionAddress,
		f.ContentType,
		bytes.NewBuffer(requestBody),
	)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("firebase request returned status: %d", resp.StatusCode)
	}

	return nil
}
