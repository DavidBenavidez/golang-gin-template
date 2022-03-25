package utils

import (
	"fmt"
	"os"
)

func GetEnvVariable(key, defaultValue string) string {
	envValue, exist := os.LookupEnv(key)
	if exist {
		fmt.Printf("ENV VARIABLE %s FOUND: %s\n", key, envValue)
		return envValue
	} else {
		fmt.Printf("ENV VARIABLE %s NOT FOUND. DEFAULTING TO %s\n", key, defaultValue)
		return defaultValue
	}
}
