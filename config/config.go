package config

import (
	"github.com/joho/godotenv"
	"os"
	"strconv"
	"strings"

	"github.com/rs/zerolog/log"
)

func Init(confPath ...string) {
	if len(confPath) == 0 {
		confPath = append(confPath, ".env")
	}
	for _, path := range confPath {
		if err := godotenv.Load(path); err != nil {
			log.Warn().Msgf("Error loading %s file", path)
			continue
		}
		log.Info().Msgf("Loaded %s config file", path)
	}
}

type DatabaseConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	Database string
}

func NewDatabaseConfig() *DatabaseConfig {
	return &DatabaseConfig{
		Host:     getStrEnv("DB_HOST", "localhost"),
		Port:     getIntEnv("DB_PORT", 5432),
		User:     getStrEnv("DB_USER", "root"),
		Password: getStrEnv("DB_PASSWORD", ""),
		Database: getStrEnv("DB_NAME", ""),
	}
}

type LogConfig struct {
	Level  int8
	Format string
}

func NewLogConfig() *LogConfig {
	levelStr := getStrEnv("LOG_LEVEL", "info")
	levelStr = strings.ToLower(levelStr)
	var level int8 = 1
	switch levelStr {
	case "trace":
		level = -1
	case "debug":
		level = 0
	case "info":
		level = 1
	case "warn":
		level = 2
	case "error":
		level = 3
	case "fatal":
		level = 4
	case "panic":
		level = 5
	default:
		log.Info().Msgf("Log level %s not found, set log level INFO", levelStr)
		level = 1
	}
	return &LogConfig{
		Level:  level,
		Format: getStrEnv("LOG_FORMAT", "json"),
	}
}

func getStrEnv(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return fallback
}

func getIntEnv(key string, fallback int) int {
	if value := os.Getenv(key); value != "" {
		val, err := strconv.Atoi(value)
		if err == nil {
			return val
		}
	}
	return fallback
}

//func getBoolEnv(key string, fallback bool) bool {
//	if value := os.Getenv(key); value != "" {
//		val, err := strconv.ParseBool(value)
//		if err == nil {
//			return val
//		}
//	}
//	return fallback
//}
