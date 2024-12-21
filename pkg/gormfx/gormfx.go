package gormfx

import (
	"log"

	"github.com/priscila-albertini-da-silva/transactions-routine/internal/configuration"
	"go.uber.org/fx"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func initDB() {
	host := configuration.Configuration.Database

	gormDB, err := gorm.Open(postgres.New(postgres.Config{
		DSN: host,
	}), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatalf("Failed to connect to PostgreSQL database: %v", err)
	}

	gormDB.AutoMigrate(
	// &Account{},
	// &OperationType{},
	// &Transaction{},
	)

	log.Println("Database migration completed successfully.")
}

var ModuleGorm = fx.Module("gorm", fx.Invoke(initDB))
