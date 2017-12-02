package store

import (
	log "github.com/sirupsen/logrus"

	"github.com/JunkieJI/go-boilerplate/config"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3" // driver for sqlx
)

// Store interface
type Store interface {
	Ping() bool
}

type store struct {
	db *sqlx.DB
}

// NewStore creates a sqlx connection
func NewStore(dsn config.DSN) Store {
	if dsn.Driver == "" {
		config.Init()
		dsn = config.GetDSN()
	}
	s := &store{}
	s.connect(dsn)

	return s
}

func (s *store) connect(dsn config.DSN) error {
	log.Println("Connecting...")
	var err error
	s.db, err = sqlx.Connect(dsn.Driver, dsn.ConnectionString())
	if err != nil {
		return err
	}

	return nil
}

func (s *store) Ping() bool {
	return s.db.Ping() == nil
}
