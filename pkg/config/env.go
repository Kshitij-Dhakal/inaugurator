package config

import "os"

func IsDevEnv() bool {
	return os.Getenv("INAUGURATOR_ENV") == "dev"
}