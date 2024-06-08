package env

import (
	"os"
)

var (
	MigrationPath = GetOrDefault("MIGRATION_PATH", "file://db/migrations")
	DbHost        = GetOrDefault("DB_HOST", "localhost")
	DbPort        = GetOrDefault("DB_PORT", "5432")
	DbName        = GetOrDefault("DB_NAME", "contacts")
	DbUser        = GetOrDefault("DB_USER", "docker")
	DbPassword    = GetOrDefault("DB_PASSWORD", "docker")
	JwtKey        = GetOrDefault("JWT_KEY", "some_secret_key")
)

func GetOrDefault(name string, defaultValue string) string {

	value := os.Getenv(name)
	if value != "" {
		return value
	}

	return defaultValue
}
