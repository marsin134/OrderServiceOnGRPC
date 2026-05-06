package main

import (
	"github.com/sirupsen/logrus"
	"orderServiceGRPC/internal/config"
	"orderServiceGRPC/internal/database"
)

func main() {
	if err := config.LoadEnvFile(".env"); err != nil {
		config.Log.WithFields(logrus.Fields{
			"func":  "main",
			"error": err}).Error("Failed to load .env")
	}

	cfg, err := config.LoadConfig()
	if err != nil {
		config.Log.WithFields(logrus.Fields{
			"func":  "main",
			"error": err}).Error("Failed to read .env")
	}

	db, err := database.ConnectedDB(cfg)
	defer db.Close()
	if err != nil {
		config.Log.WithFields(logrus.Fields{
			"func":  "main",
			"error": err}).Error("Failed to connect to database")
	}

	//err = db.RunMigrations("C:\\Users\\lampe\\GolandProjects\\orderServiceGRPC\\migrations\\001_init.sql")
	//if err != nil {
	//	config.Log.WithFields(logrus.Fields{
	//		"func":  "main",
	//		"error": err}).Error("Failed to run migrations")
	//}

}
