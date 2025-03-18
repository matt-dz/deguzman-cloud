package database

import (
	"context"
	"log/slog"
	"os"
	"sync"

	"github.com/jackc/pgx/v5/pgxpool"
)

var db *pgxpool.Pool
var once sync.Once

func initDb() {
	var err error

	db, err = pgxpool.New(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		slog.Error("Failed to connect to db", "error", err.Error())
		panic(err)
	}
}

func GetDb() *pgxpool.Pool {
	once.Do(initDb)
	return db
}
