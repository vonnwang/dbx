package dbx

import (
	"database/sql"
	"database/sql/driver"
	"imooc.com/tietang/red-envelope/infra/dbx/mapping"
)

type Mapper func(model interface{}, rows *sql.Rows) (interface{}, error)
type RowsMapper func(rows *sql.Rows) (interface{}, error)
type RowMapper func(row *sql.Row) (interface{}, error)

type TxRunner struct {
	*Runner
	driver.Tx
}

type Runner struct {
	SqlExecutor
	mapping.EntityMapper
	ILogger
	LoggerSettings
}

func NewTxRunner(tx *sql.Tx) *TxRunner {
	r := &TxRunner{}
	r.Runner = &Runner{}
	r.SqlExecutor = tx
	r.Tx = tx
	return r
}

func NewRunner(se SqlExecutor, em mapping.EntityMapper) *Runner {
	r := &Runner{}
	r.SqlExecutor = se
	r.EntityMapper = em
	return r
}