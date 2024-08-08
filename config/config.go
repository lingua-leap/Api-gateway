package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/cast"

	_ "github.com/lib/pq"
)

type Config struct {
	SECRET_KEY_ACCESS  string
	HOST               string
	GIN_SERVER_PORT    string
	GRPC_USER_PORT     string
	GRPC_LEARNING_PORT string
	GRPC_PROGRESS_PORT string
}

func Load() Config {
	err := godotenv.Load(".env")
	if err != nil {
		log.Println("No .env file found")
	}

	config := Config{}

	config.GIN_SERVER_PORT = cast.ToString(coalesce("GIN_SERVER_PORT", ":8080"))
	config.GRPC_USER_PORT = cast.ToString(coalesce("GRPC_USER_PORT", ":50050"))
	config.GRPC_LEARNING_PORT = cast.ToString(coalesce("GRPC_LEARNING_PORT", ":50051"))
	config.GRPC_PROGRESS_PORT = cast.ToString(coalesce("GRPC_PROGRESS_PORT", ":50052"))
	config.HOST = cast.ToString(coalesce("HOST", "localhost"))
	config.SECRET_KEY_ACCESS = cast.ToString(coalesce("SECRET_KEY_ACCESS", "secret-key"))

	return config
}

func coalesce(env string, defaultValue interface{}) interface{} {
	value, exists := os.LookupEnv(env)
	if !exists {
		return defaultValue
	}
	return value
}
