package repository

import (
	"database/sql"
	"fmt"
)

func listDrivers() {
	for _, driver := range sql.Drivers() {
		fmt.Printf("Driver: %w", driver)
	}
}

type Config struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
	SSLMode  string
}

func GetPgConnection(cfg Config) (*sql.DB, error) {
	var connectionStr = fmt.Sprintf("dbname=%s user=%s password=%s host=%s port=%s sslmode=disable",
		cfg.DBName, cfg.Username, cfg.Password, cfg.Host, cfg.SSLMode)

	var db, err = sql.Open("postgres", connectionStr)
	if err != nil {
		fmt.Println(err)
	}

	return db, nil
}

//type NewDBConnection interface {
//	getPgConnection()
//}
