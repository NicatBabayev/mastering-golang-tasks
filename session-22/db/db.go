package db

import (
	"gorm.io/driver/postgres"
	_ "gorm.io/driver/postgres"
	"gorm.io/gorm"
	"session-22/config"
)

type Postgres struct{}

type Connector interface {
	Connect(string) (*gorm.DB, error)
}

func (p Postgres) Connect(dsn string) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}

func NewConnection() (*gorm.DB, error) {
	// get configuration
	dsn, err := config.GenerateDSN("postgres")
	if err != nil {
		return nil, err
	}
	// connect to db
	var dbConnector Connector
	postgre := Postgres{}
	dbConnector = postgre
	db, err := dbConnector.Connect(dsn)
	if err != nil {
		return nil, err
	}
	return db, nil
	// return the result of connection(either error or db instance)

}
