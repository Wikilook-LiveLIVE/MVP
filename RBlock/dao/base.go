package dao

import (
	"github.com/artjoma/flog"
	"github.com/jackc/pgx"
	"net"
	"time"
)

var (
	pool *pgx.ConnPool
	log  *flog.Logger
)

//New connection pool
func NewConnectionPool(_log *flog.Logger, host string, port uint16, user string, pass string,


	return err
}

func Close() {

}

func GetPool() *pgx.ConnPool{
	return pool
}

//Use for insert(PK increment), etc.
func InsertReturning(sqlName string, values ...interface{}) (int64, error) {

	return rowId, tx.Commit()
}

func UpdateReturning(sqlName string, values ...interface{}) (int64, error) {


	//Commit
	return rowId, tx.Commit()
}

func SelectForInt(sqlName string, values ...interface{}) (result int64, err error) {

	return result, nil
}

func SelectForString(sqlName string, values ...interface{}) (result string, err error) {

	return result, nil
}
