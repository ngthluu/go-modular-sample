package db

import (
	"database/sql"
)

var _ Db = (*PostgreSqlDb)(nil)

type PostgreSqlDb struct {
	Conn *sql.DB
}

func NewPostgreSqlDb() *PostgreSqlDb {
	return nil
}

func (db *PostgreSqlDb) GetConn() *sql.DB {
	return db.Conn
}
