package utils

import (
	"fmt"
	"os"
)

// ConnectionURLBuilder func for building URL connection.
func ConnectionURLBuilder(n string) (string, error) {
	// Define URL to connection.
	var url string

	// Switch given names.
	switch n {
	case "mysql":
		// URL for PostgreSQL connection.
		// %s:%s@tcp(%s:%s)/%s", dbUser, dbPassword, dbHost, dbPort, dbName
		url = "root:123456@tcp(localhost:3306)/echo-fiber?parseTime=True"
		// url = fmt.Sprintf(
		// 	"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		// 	os.Getenv("DB_HOST"),
		// 	os.Getenv("DB_PORT"),
		// 	os.Getenv("DB_USER"),
		// 	os.Getenv("DB_PASSWORD"),
		// 	os.Getenv("DB_NAME"),
		// 	os.Getenv("DB_SSL_MODE"),
		// )
	case "postgres":
		// URL for PostgreSQL connection.
		url = fmt.Sprintf(
			"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
			"localhost",
			"5432",
			"root",
			"123456",
			"echo-sqlx",
			"disable",
			// os.Getenv("DB_HOST"),
			// os.Getenv("DB_PORT"),
			// os.Getenv("DB_USER"),
			// os.Getenv("DB_PASSWORD"),
			// os.Getenv("DB_NAME"),
			// os.Getenv("DB_SSL_MODE"),
		)
	case "redis":
		// URL for Redis connection.
		url = fmt.Sprintf(
			"%s:%s",
			os.Getenv("REDIS_HOST"),
			os.Getenv("REDIS_PORT"),
		)
	case "fiber":
		// URL for Fiber connection.
		url = fmt.Sprintf(
			"%s:%s",
			"localhost",
			"8080",
			// os.Getenv("SERVER_HOST"),
			// os.Getenv("SERVER_PORT"),
		)
	default:
		// Return error message.
		return "", fmt.Errorf("connection name '%v' is not supported", n)
	}

	// Return connection URL.
	return url, nil
}
