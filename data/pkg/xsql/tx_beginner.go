package xsql

import (
	"database/sql"
)

type TxBeginner interface {
	Begin() (*sql.Tx, error)
}
