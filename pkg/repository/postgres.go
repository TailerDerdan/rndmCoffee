package postgres

import (
	"database/sql"
	"fmt"

	"github.com/jmoiron/sqlx"
	//"golang.org/x/tools/go/cfg"
	
)

type Config struct {
	Host     string
	Port     string
	Password string
	Username string
	DBName   string
	SSLMode  string
}

func NewPostgresDB(cfg Config) (*sql.DB, error) {
	db, err := sqlx.Open("postgres", fmt.Sprintf("host: %s, port: %s, username: %s, passowrd: %s, dbname: %s, sslmode: %s",
		cfg.Host, cfg.Port, cfg.Username, cfg.Password, cfg.DBName, cfg.SSLMode))

	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
