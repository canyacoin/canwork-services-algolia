package main

import (
	"fmt"
	"os"
)

func getEnv(key, fallback string) string {
	returnVal := fallback
	if value, ok := os.LookupEnv(key); ok {
		returnVal = value
	}
	if returnVal == "" {
		panic(fmt.Sprintf("Unable to retrieve key: %s", key))
	}
	return returnVal
}
