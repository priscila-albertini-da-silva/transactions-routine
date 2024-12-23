package database

import (
	"fmt"
	"log"

	"github.com/priscila-albertini-da-silva/transactions-routine/internal/core/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DBExpect struct {
	DB *gorm.DB
}

func NewDbExpect(conn *gorm.DB) *DBExpect {
	return &DBExpect{conn}
}

func NewDatabaseConnection(host string) *DBExpect {
	db := getDatabaseConnection(host)
	err := db.AutoMigrate(&model.Account{}, &model.Transaction{})
	if err != nil {
		log.Fatalf("failed to run migrations: %v", err)
	}
	return &DBExpect{
		DB: db,
	}
}

func getDatabaseConnection(host string) *gorm.DB {
	gormDB, err := gorm.Open(postgres.New(postgres.Config{
		DSN: host,
	}), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})

	if err != nil {
		log.Panicf("error to connect database: %v", err)
	}

	log.Printf("Database connected")

	return gormDB
}

func (db DBExpect) GetAll(tableName string) (result [][]string) {
	rows, err := db.DB.Table(tableName).Select("*").Order("id").Rows()
	defer rows.Close()
	if err != nil {
		log.Fatal(err)
	}

	cols, err := rows.Columns()

	if err != nil {
		log.Fatal(err)
	}

	rawResult := make([][]byte, len(cols))
	resultRow := make([]string, len(cols))

	dest := make([]interface{}, len(cols))
	for i := range rawResult {
		dest[i] = &rawResult[i]
	}

	for rows.Next() {
		err = rows.Scan(dest...)
		if err != nil {
			fmt.Println("Failed to scan row", err)

			return
		}

		for i, raw := range rawResult {
			if raw == nil {
				resultRow[i] = "\\N"
			} else {
				resultRow[i] = string(raw)
			}
		}

		cloneResultRow := make([]string, len(resultRow))

		copy(cloneResultRow, resultRow)
		result = append(result, cloneResultRow)
	}

	if err != nil {
		log.Fatal(err)
	}

	return result
}
