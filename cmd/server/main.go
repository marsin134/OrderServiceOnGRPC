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

	//err = db.RunMigrations("migrations\\001_init.sql")
	//if err != nil {
	//	config.Log.WithFields(logrus.Fields{
	//		"func":  "main",
	//		"error": err}).Error("Failed to run migrations")
	//}

	//repo := repository.NewRepository(repository.NewOrderRepository(db))
	//ctx := context.Background()
	//
	//err = repo.Order.DeleteOrder(ctx, "2a28de9f-a55c-46bb-9c10-082c27e9038f")
	//if err != nil {
	//	fmt.Println(err)
	//}
	//err = repo.Order.DeleteOrder(ctx, "cfd4c16f-0d1d-4f79-9e13-696b989df97c")
	//if err != nil {
	//	fmt.Println(err)
	//}
}
