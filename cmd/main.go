package main

import (
	"fmt"
	"log"
	"net/http"
	h "user-data-service/internal/handler"
	m "user-data-service/internal/middleware"
	r "user-data-service/internal/repository"
	"user-data-service/pkg/db"
	s "user-data-service/pkg/sever"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("User Data Service Started")

	mysqlCnf := db.NewConfig()
	mysql := db.NewMysql(mysqlCnf)
	err := mysql.Open()
	if err != nil {
		log.Fatal(err)
	}

	userRepo := r.NewUserRepo(mysql)
	userH := h.NewUserHandler(userRepo)
	router := mux.NewRouter()
	RegisterUserRoutes(router, userH)

	srv := s.NewServer(router)
	err = http.ListenAndServe(":8080", srv)
	if err != nil {
		fmt.Println(err)
	}
}

func RegisterUserRoutes(router *mux.Router, h *h.UserHandler) {
	router.HandleFunc("/api/users", m.DefaultHeaders(h.GetAll())).Methods("GET")
	router.HandleFunc("/api/user", m.DefaultHeaders(h.Get())).Methods("GET")
	router.HandleFunc("/api/user", m.DefaultHeaders(h.Create())).Methods("POST")
	router.HandleFunc("/api/user", m.DefaultHeaders(h.Update())).Methods("PATCH")
	router.HandleFunc("/api/user", m.DefaultHeaders(h.Delete())).Methods("DELETE")
}
