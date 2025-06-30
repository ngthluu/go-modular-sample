package db

import (
	"database/sql"
)

type Db interface {
	GetConn() *sql.DB
}
