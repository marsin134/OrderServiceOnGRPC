package config

import (
	"bufio"
	"fmt"
	"github.com/sirupsen/logrus"
	"os"
	"strconv"
	"strings"
)

type DB struct {
	Host    string
	Port    string
	User    string
	Pass    string
	DBName  string
	SSLMode string
}

type Config struct {
	ServerPort int
	DB         DB
	JWTKey     string
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func getEnvPassword(key string) (string, error) {
	if value, ok := os.LookupEnv(key); ok {
		return value, nil
	}
	return "", fmt.Errorf("missing environment variable")
}

func loadDB() (DB, error) {
	password, err := getEnvPassword("DB_PASSWORD")
	if err != nil {
		return DB{}, err
	}
	return DB{
		Host:    getEnv("DB_HOST", "localhost"),
		Port:    getEnv("DB_PORT", "5432"),
		User:    getEnv("DB_USER", "postgres"),
		Pass:    password,
		DBName:  getEnv("DB_NAME", "productService"),
		SSLMode: getEnv("DB_SSLMODE", "disable"),
	}, nil
}

func LoadConfig() (*Config, error) {
	port, err := strconv.Atoi(getEnv("SERVER_PORT", "8080"))
	if err != nil {
		return nil, fmt.Errorf("error receiving the port %w", err)
	}

	db, err := loadDB()

	if err != nil {
		return nil, err
	}

	jwtKey, err := getEnvPassword("DB_PASSWORD")
	if err != nil {
		return nil, err
	}

	return &Config{
		ServerPort: port,
		DB:         db,
		JWTKey:     jwtKey,
	}, nil
}

func LoadEnvFile(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		// Skip comments and blank lines
		if len(line) == 0 || line[0] == '#' {
			continue
		}

		// Separating the key and the value
		parts := strings.SplitN(line, "=", 2)
		if len(parts) != 2 {
			continue
		}

		key := strings.TrimSpace(parts[0])
		value := strings.TrimSpace(parts[1])

		// Setting the environment variable
		// Only if it is not installed yet
		if _, exists := os.LookupEnv(key); !exists {
			os.Setenv(key, value)
		}
	}

	return scanner.Err()
}

func setupLogger() *logrus.Logger {
	logger := logrus.New()
	logger.SetFormatter(&logrus.TextFormatter{
		ForceColors: true,
	})
	return logger
}

var Log = setupLogger()
