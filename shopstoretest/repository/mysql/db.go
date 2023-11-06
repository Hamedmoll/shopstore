package mysql

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"shopstoretest/cfg"
	"time"
)

type MySQLDB struct {
	DB *sql.DB
}

func New(cfg cfg.DataBaseConfig) *MySQLDB {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@%s(%s:%s)/%s",
		cfg.DataBaseUser, cfg.DataBasePassword, cfg.DataBaseProtocol, cfg.DataBaseHost, cfg.DataBasePort, cfg.DataBaseName))
	//db, err := sql.Open("mysql", "shopstore:shopstoret0lk2o20@tcp(localhost:19999)/shopstore_db")

	if err != nil {
		panic(fmt.Errorf("cant open mysql db %v", err))
	}

	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	return &MySQLDB{DB: db}
}
