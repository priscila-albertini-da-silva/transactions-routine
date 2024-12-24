package gormfx

import (
	"log"

	"github.com/priscila-albertini-da-silva/transactions-routine/internal/configuration"
	"github.com/priscila-albertini-da-silva/transactions-routine/internal/core/model"
	"go.uber.org/fx"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func initDB() (*gorm.DB, error) {
	host := configuration.Configuration.Database

	gormDB, err := gorm.Open(postgres.New(postgres.Config{
		DSN: host,
	}), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		log.Fatalf("Failed to connect to PostgreSQL database: %v", err)
		return nil, err
	}

	// Executa as migrações
	err = gormDB.AutoMigrate(
		&model.Account{},
		&model.OperationType{},
		&model.Transaction{},
	)
	if err != nil {
		log.Fatalf("Failed to run migrations: %v", err)
		return nil, err
	}

	log.Println("Database migration completed successfully.")
	return gormDB, nil
}

var ModuleGorm = fx.Options(
	fx.Provide(initDB), // Alterado para fx.Provide
)
