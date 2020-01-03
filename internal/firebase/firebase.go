package firebase

import (
	"github.com/consumer-superhero-match/internal/config"
)

// Firebase holds all the Firebase related data.
type Firebase struct {
	FunctionAddress string
	ContentType     string
	ParameterName   string
}

// NewFirebase creates new value of type Firebase.
func NewFirebase(cfg *config.Config) *Firebase {
	return &Firebase{
		FunctionAddress: cfg.Firebase.FunctionAddress,
		ContentType:     cfg.Firebase.ContentType,
		ParameterName:   cfg.Firebase.ParameterName,
	}
}
