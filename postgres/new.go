package postgres

import (
	"database/sql"
	"fmt"
	"github.com/ljmcclean/knight-hacks-2024/config"
	"github.com/ljmcclean/knight-hacks-2024/services"
)

type postgreSQL struct {
	db *sql.DB
}

func New(cfg *config.Config) (db services.Database, err error) {
	pg, err := connect(cfg)
	if err != nil {
		return nil, err
	}
	return &postgreSQL{
		db: pg,
	}, nil
}

func (pg *postgreSQL) Close() {
	pg.db.Close()
}

func connect(cfg *config.Config) (pg *sql.DB, err error) {
	connStr, err := getConnStr(cfg)
	if err != nil {
		return nil, err
	}
	pg, err = sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}
	err = pg.Ping()
	if err != nil {
		return nil, err
	}
	return pg, nil
}

func getConnStr(cfg *config.Config) (dbURL string, err error) {
	return fmt.Sprintf("postgres://%v:%v@%v:%v/%v?sslmode=disable",
			cfg.Database.User,
			cfg.Database.Password,
			cfg.Database.Host,
			cfg.Database.Port,
			cfg.Database.Name),
		nil
}
