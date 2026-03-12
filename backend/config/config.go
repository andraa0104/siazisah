package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

type Config struct {
	DBHost            string
	DBPort            string
	DBUser            string
	DBPassword        string
	DBName            string
	DBMaxOpenConns    int
	DBMaxIdleConns    int
	DBConnMaxLifetime time.Duration
	DBConnMaxIdleTime time.Duration
	ServerPort        string
	JWTSecret         string
}

var AppConfig *Config

func LoadConfig() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using environment variables")
	}

	AppConfig = &Config{
		DBHost:            getEnv("DB_HOST", "localhost"),
		DBPort:            getEnv("DB_PORT", "3306"),
		DBUser:            getEnv("DB_USER", "root"),
		DBPassword:        getEnv("DB_PASSWORD", ""),
		DBName:            getEnv("DB_NAME", "siazisah"),
		DBMaxOpenConns:    getEnvInt("DB_MAX_OPEN_CONNS", 30),
		DBMaxIdleConns:    getEnvInt("DB_MAX_IDLE_CONNS", 15),
		DBConnMaxLifetime: time.Duration(getEnvInt("DB_CONN_MAX_LIFETIME_MINUTES", 30)) * time.Minute,
		DBConnMaxIdleTime: time.Duration(getEnvInt("DB_CONN_MAX_IDLE_TIME_MINUTES", 10)) * time.Minute,
		ServerPort:        getEnv("SERVER_PORT", "8080"),
		JWTSecret:         getEnv("JWT_SECRET", "your-secret-key"),
	}
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

func getEnvInt(key string, defaultValue int) int {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}

	parsed, err := strconv.Atoi(value)
	if err != nil {
		return defaultValue
	}

	return parsed
}

func InitDB() (*sql.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true&loc=Local&timeout=5s&readTimeout=10s&writeTimeout=10s&interpolateParams=true",
		AppConfig.DBUser,
		AppConfig.DBPassword,
		AppConfig.DBHost,
		AppConfig.DBPort,
		AppConfig.DBName,
	)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	// Connection pool tuning: reuse existing DB connections across requests.
	db.SetMaxOpenConns(AppConfig.DBMaxOpenConns)
	db.SetMaxIdleConns(AppConfig.DBMaxIdleConns)
	db.SetConnMaxLifetime(AppConfig.DBConnMaxLifetime)
	db.SetConnMaxIdleTime(AppConfig.DBConnMaxIdleTime)

	if err = db.Ping(); err != nil {
		return nil, err
	}

	log.Printf("Database connected successfully (pool: open=%d idle=%d lifetime=%s idleTime=%s)",
		AppConfig.DBMaxOpenConns,
		AppConfig.DBMaxIdleConns,
		AppConfig.DBConnMaxLifetime,
		AppConfig.DBConnMaxIdleTime,
	)
	return db, nil
}
