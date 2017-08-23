package db

import (
	"strings"
	"sync"

	"github.com/maknahar/go-web-skeleton/internal/utils/config"
	"github.com/maknahar/go-web-skeleton/internal/utils/logger"
	"gopkg.in/jackc/pgx.v2"
)

//Pool of database connection
var ConnPool *pgx.ConnPool
var once sync.Once

//Init the connection to DB
func Init(url string) error {
	l := logger.GetLogger()
	if ConnPool == nil {
		once.Do(func() {
			connConfig, err := pgx.ParseURI(url)
			if err != nil {
				l.Fatalf("Invalid Database URL, Err: %+v", err)
				return
			}
			poolConfig := pgx.ConnPoolConfig{
				ConnConfig:     connConfig,
				MaxConnections: config.MaxDBConnections,
			}
			ConnPool, err = pgx.NewConnPool(poolConfig)
			if err != nil {
				l.Fatalf("Unable to connect to Database, Err: %+v", err)
				return
			}
			l.Println("Successfully established database connection to", poolConfig.Database)
		})
	}
	return nil
}

// ExecuteInTransaction executes the given function in DB transaction
// only if there is no error otherwise it is rolled back.
func ExecuteInTransaction(f func(tx *pgx.Tx) error) (err error) {
	tx, err := ConnPool.Begin()
	if err != nil {
		return err
	}

	defer func() {
		if r := recover(); r != nil {
			logger.GetLogger().Println("RECOVERED: function", r)
			rollbackTransaction(tx)
		}
	}()

	err = f(tx)
	if err != nil {
		rollbackTransaction(tx)
		return err
	}

	err = tx.Commit()
	if err != nil {
		rollbackTransaction(tx)
		return err
	}
	return nil
}

type DBConn struct {
	conn          *pgx.ConnPool
	tx            *pgx.Tx
	isTransaction bool
}

// Interface to abstract the queryer(dbconnection or transaction)
type Queryer interface {
	Exec(sql string, arguments ...interface{}) (pgx.CommandTag, error)
	Query(sql string, args ...interface{}) (*pgx.Rows, error)
	QueryRow(sql string, args ...interface{}) *pgx.Row
}

// Initialize the DB connection and assign the existing db connection
func (db *DBConn) Init() {
	db.conn = ConnPool
}

func (db *DBConn) GetQueryer() Queryer {
	if db.isTransaction {
		return db.tx
	} else {
		return db.conn
	}
}

// ExecuteInTransaction executes the given function in DB transaction, i.e. It commits
// only if there is not error otherwise it is rolledback.
func (db *DBConn) ExecuteInTransaction(f func() error) (err error) {
	tx, err := db.conn.Begin()
	if err != nil {
		return err
	}
	db.tx = tx
	db.isTransaction = true

	defer func() {
		if r := recover(); r != nil {
			logger.GetLogger().Println("Recovered in function ", r)
			db.rollbackTransaction(tx)
		}
		db.isTransaction = false
	}()

	err = f()
	if err != nil {
		db.rollbackTransaction(tx)
		return err
	}
	err = tx.Commit()
	if err != nil {
		db.rollbackTransaction(tx)
		return err
	}
	return nil
}

func (db *DBConn) rollbackTransaction(tx *pgx.Tx) {
	err := tx.Rollback()
	if err != nil {
		logger.GetLogger().Println("ERROR: While rollback, Err: ", err)
	}
}

func rollbackTransaction(tx *pgx.Tx) {
	err := tx.Rollback()
	if err != nil {
		logger.GetLogger().Println("ERROR: While rolling back", err)
	}
}

func StringArrayToString(s []string) string {
	return "{" + strings.Join(s, ",") + "}"
}
