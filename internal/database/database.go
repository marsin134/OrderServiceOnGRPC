package database

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"orderServiceGRPC/internal/config"
	"os"
	"time"
)

type MethodsDB interface {
	Close()
	RunMigrations(migrationFilePath string) error
	HealthCheck() error
	GetDB() *DB
}

type DB struct {
	*sqlx.DB
}

func ConnectedDB(cfg *config.Config) (*DB, error) {
	config.Log.WithFields(logrus.Fields{
		"port":    cfg.DB.Port,
		"DB name": cfg.DB.DBName,
	}).Info("Connect to db")

	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		cfg.DB.Host, cfg.DB.Port, cfg.DB.User, cfg.DB.Pass, cfg.DB.DBName, cfg.DB.SSLMode)

	db, err := sqlx.Connect("postgres", connStr)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("error checking the connection to the database: %w", err)
	}

	db.SetConnMaxLifetime(5 * time.Minute)

	config.Log.Info("Successfully connected to db")

	return &DB{db}, nil
}

func (db *DB) RunMigrations(migrationFilePath string) error {
	if _, err := os.Stat(migrationFilePath); os.IsNotExist(err) {
		return fmt.Errorf("file %s does not exist", migrationFilePath)
	}

	migration, err := os.ReadFile(migrationFilePath)
	if err != nil {
		return fmt.Errorf("error reading file %s: %w", migrationFilePath, err)
	}

	config.Log.WithFields(logrus.Fields{
		"path": migrationFilePath}).Info("Attempting to run migration")

	_, err = db.Exec(string(migration))
	if err != nil {
		return fmt.Errorf("error executing migration %s: %w", migrationFilePath, err)
	}

	config.Log.Info("Successfully run migration")
	return nil
}

func (db *DB) Close() {
	db.DB.Close()
}

func (db *DB) GetDB() *DB {
	return db
}

// psql -h localhost -U postgres -d product_service -f migrations/001_init.sql
