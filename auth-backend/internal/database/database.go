package database

import (
	"context"
	"os"
	"sync"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
)

var db *pgxpool.Pool
var once sync.Once

// Any custom DB types made with CREATE TYPE need to be registered with pgx.
// https://github.com/kyleconroy/sqlc/issues/2116
// https://stackoverflow.com/questions/75658429/need-to-update-psql-row-of-a-composite-type-in-golang-with-jack-pgx
// https://pkg.go.dev/github.com/jackc/pgx/v5/pgtype
func getCustomDateTypes(ctx context.Context, pool *pgxpool.Pool) ([]*pgtype.Type, error) {
	// Get a single connection just to load type information.
	conn, err := pool.Acquire(ctx)
	defer conn.Release()
	if err != nil {
		return nil, err
	}

	dataTypeNames := []string{
		"role",
		// An underscore prefix is an array type in pgtypes.
		"_role",
	}

	var typesToRegister []*pgtype.Type
	for _, typeName := range dataTypeNames {
		dataType, err := conn.Conn().LoadType(ctx, typeName)
		if err != nil {
			return nil, err
		}
		// You need to register only for this connection too, otherwise the array type will look for the register element type.
		conn.Conn().TypeMap().RegisterType(dataType)
		typesToRegister = append(typesToRegister, dataType)
	}
	return typesToRegister, nil
}

func GetDb() *pgxpool.Pool {
	once.Do(func() {
		var err error

		config, err := pgxpool.ParseConfig(os.Getenv("DATABASE_URL"))
		if err != nil {
			panic(err)
		}
		dbpool, err := pgxpool.NewWithConfig(context.Background(), config)
		if err != nil {
			panic(err)
		}

		customTypes, err := getCustomDateTypes(context.Background(), dbpool)
		if err != nil {
			panic(err)
		}

		config.AfterConnect = func(ctx context.Context, conn *pgx.Conn) error {
			for _, t := range customTypes {
				conn.TypeMap().RegisterType(t)
			}
			return nil
		}
		dbpool.Close()

		db, err = pgxpool.NewWithConfig(context.Background(), config)
		if err != nil {
			panic(err)
		}
	})
	return db
}
