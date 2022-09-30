package db

import (
	"battles/internal/utils/env"
	"battles/internal/utils/logger"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"os"
	"sync"
)

var db *sqlx.DB = nil
var once sync.Once

func Get() *sqlx.DB {
	once.Do(func() {
		env.InitEnv()
		var err error
		db, err = sqlx.Open(os.Getenv("DB_DRIVER_NAME_SQLX"), os.Getenv("DB_URL_SQLX"))
		if err != nil {
			logger.Get().Warnf(`Get DB error:"%v"`, err)
			panic(err)
		}
	})
	return db
}
