package database

import (
	"database/sql"
	"fmt"

	"github.com/ilmsg/ideal-octo-sniffle/config"
)

func GetDatabase(dbconfig *config.DBConfig) (*sql.DB, error) {
	dataSourceName := fmt.Sprintf("%s:%s@%s/%s?parseTime=true",
		dbconfig.Username,
		dbconfig.Passwod,
		dbconfig.Host,
		dbconfig.DBName,
	)

	db, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		return nil, err
	}
	return db, nil
}
