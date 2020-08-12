package testutil

import (
	"kirby/config"
	"os"
)

// LoadEnv sets up the test environment variables
func LoadEnv() {
	os.Setenv("KIRBY_ENV", "test")
	config.LoadEnv()
}
