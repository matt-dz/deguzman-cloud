package database

import (
	"context"
	"os"
	"sync"

	"github.com/jackc/pgx/v5/pgxpool"
)

var db *pgxpool.Pool
var once sync.Once

func GetDb() *pgxpool.Pool {
	once.Do(func() {
		var err error

		db, err = pgxpool.New(context.Background(), os.Getenv("DATABASE_URL"))
		if err != nil {
			panic(err)
		}
	})
	return db
}
