package envvars

import (
	"os"
)

// LogLevel - setting the logging level for the application
var LogLevel = getEnv("LOGGING_LEVEL", "debug")

// ServerPort - port the server runs on
var ServerPort = getEnv("SERVER_PORT", "9090")

// Scheme - HTTP or HTTPS
var Scheme = getEnv("URL_SCHEME", "https")

// Env - Deployed ENV
var Env = getEnv("ENV", "dev")

// AWSRegion - deployment region
var AWSRegion = getEnv("AWS_REGION", "us-west-2")

func getEnv(key string, fallback string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return fallback
	}
	return value
}
