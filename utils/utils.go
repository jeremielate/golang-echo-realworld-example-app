package utils

import (
	"log"
	"os"
)

func LookupRequiredEnv(key string) string {
	value, ok := os.LookupEnv(key)
	if !ok {
		log.Fatalf("environment variable %v not set.")
	}
	return value
}
