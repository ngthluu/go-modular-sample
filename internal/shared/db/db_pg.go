package db

import (
	"database/sql"
	"fmt"
	"go-modular-sample/internal/shared/constant"
	"log"

	_ "github.com/lib/pq"
)

var _ Db = (*PostgreSqlDb)(nil)

type PostgreSqlDb struct {
	Conn *sql.DB
}

func NewPostgreSqlDb() *PostgreSqlDb {
	connStr := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=verify-full",
		constant.EnvDbUser,
		constant.EnvDbPassword,
		constant.EnvDbHost,
		constant.EnvPort,
		constant.EnvDbName,
	)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	return &PostgreSqlDb{
		Conn: db,
	}
}

func (db *PostgreSqlDb) GetConn() *sql.DB {
	return db.Conn
}
