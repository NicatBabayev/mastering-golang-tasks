package server

import (
	"net/http"
	"session-22/config"
	"session-22/handlers"
	"session-22/middleware"

	_ "gorm.io/driver/postgres"
	"gorm.io/gorm"
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
	bookHandler := handlers.NewBooksHandler(s.DB)
	loginHandler := handlers.NewLoginHandler()
	http.HandleFunc("/books/", middleware.Authenticate(bookHandler.BooksHandler))
	http.HandleFunc("/login/", loginHandler.LoginHandler)
	srvPort := srvConfig["SRV_PORT"]
	err = http.ListenAndServe(":"+srvPort, nil)
	if err != nil {
		return err
	}
	return nil
}
