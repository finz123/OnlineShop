package database

import (
	"database/sql"
	"fmt"
	"log"
	"online-shop/infra/config"

	_ "github.com/lib/pq"
)

var (
	db  *sql.DB
	err error
)

func handleDatabaseConnection() {
	appConfig := config.GetAppConfig()
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		appConfig.DBHost, appConfig.DBPort, appConfig.DBUser, appConfig.DBPassword, appConfig.DBName,
	)

	db, err = sql.Open(appConfig.DBDialect, psqlInfo)

	if err != nil {
		log.Panic("Data Connection Wrong", err)
	}

	err = db.Ping()

	if err != nil {
		log.Panic("koneksi ke database gagal", err)
	}
}

func createDatabaseTable() {
	productTable := `
	CREATE TABLE IF NOT EXISTS "product" (
		id SERIAL PRIMARY KEY,
		nama VARCHAR(255) NOT NULL,
		harga INT NOT NULL,
		deskripsi TEXT,
		gambar VARCHAR(255)
	);
`

	createTableQuery := fmt.Sprintf("%s", productTable)

	_, err = db.Exec(createTableQuery)

	if err != nil {
		log.Panic("error occured while trying to create required tables:", err)
	}
}

func InitializeDatabase() {
	handleDatabaseConnection()
	createDatabaseTable()
}

func GetDatabaseInstance() *sql.DB {
	return db
}
