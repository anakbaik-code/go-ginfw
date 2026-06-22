package config

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"time"
)

func NewMySQL(cfg *Config) (*sql.DB, func(), error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?parseTime=true",
		cfg.DBUser,
		cfg.DBPassword,
		cfg.DBHost,
		cfg.DBName,
	)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, nil, err
	}

	// pool config
	db.SetMaxOpenConns(cfg.DBMaxOpenConns)
	db.SetMaxIdleConns(cfg.DBMaxIdleConns)

	lifetime, _ := time.ParseDuration(cfg.DBConnMaxLifetime)
	db.SetConnMaxLifetime(lifetime)

	if err := db.Ping(); err != nil {
		return nil, nil, err
	}

	cleanup := func() {
		log.Println("stopping database connection safe ... ")
		db.Close()
	}

	return db, cleanup, nil
}
