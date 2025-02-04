package server

import (
	_ "gorm.io/driver/postgres"
	"gorm.io/gorm"
	"net/http"
	"session-22/config"
	"session-22/handlers"
)

type Server struct {
	DB *gorm.DB
}

func NewServer(db *gorm.DB) *Server {
	return &Server{DB: db}
}

func (s *Server) Start() error {
	srvConfig, err := config.ParseSrvConfig()
	if err != nil {
		return err
	}
	handler := handlers.NewHandler(s.DB)
	http.HandleFunc("/books", handler.BooksHandler)
	srvPort := srvConfig["SRV_PORT"]
	err = http.ListenAndServe(":"+srvPort, nil)
	if err != nil {
		return err
	}
	return nil
}
