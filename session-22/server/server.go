package server

import (
	"net/http"
	"session-22/config"
	"session-22/handlers"
)

func Init() error {
	srvConf, err := config.ParseSrvConfig()
	port := srvConf["SRV_PORT"]
	if err != nil {
		return err
	}
	http.HandleFunc("/books/", handlers.BooksHandler)
	err = http.ListenAndServe(":"+port, nil)
	if err != nil {
		return err
	}
	return nil
}
