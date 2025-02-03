package app

import (
	"session-22/db"
	"session-22/server"
)

func Init() error {
	//	init db
	_, err := db.NewConnection()
	if err != nil {
		return err
	}
	//	init server
	err = server.Init()
	if err != nil {
		return err
	}
	return nil
}
