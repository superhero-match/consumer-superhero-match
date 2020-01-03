package config

// Firebase holds all the Firebase related data.
type Firebase struct {
	FunctionAddress string `env:"FIREBASE_NEW_MATCH_FUNCTION_ADDRESS" default:"https://us-central1-superheromatch.cloudfunctions.net/newMatch"`
	ContentType     string `env:"FIREBASE_NEW_MATCH_CONTENT_TYPE" default:"application/json"`
}
