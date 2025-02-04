package app

import (
	"session-22/db"
	"session-22/server"
)

func Init() error {
	//	init db
	dbCon, err := db.NewConnection()
	if err != nil {
		return err
	}
	//	init server
	srv := server.NewServer(dbCon)
	err = srv.Start()
	if err != nil {
		return err
	}
	return nil
}
