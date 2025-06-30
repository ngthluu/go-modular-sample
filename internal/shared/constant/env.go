package constant

import "os"

var (
	EnvBaseUrl    = os.Getenv("BASE_URL")
	EnvPort       = os.Getenv("PORT")
	EnvDbHost     = os.Getenv("DB_HOST")
	EnvDbPort     = os.Getenv("DB_PORT")
	EnvDbName     = os.Getenv("DB_NAME")
	EnvDbUser     = os.Getenv("DB_USER")
	EnvDbPassword = os.Getenv("DB_PASSWORD")
)
