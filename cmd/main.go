package main

import (
	"fmt"
	"log"
	"net/http"
	h "user-data-service/internal/handler"
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
	userHandler := h.NewUserHandler(userRepo)
	router := mux.NewRouter()
	router.HandleFunc("/api/users", userHandler.GetAll()).Methods("GET")
	router.HandleFunc("/api/user", userHandler.Get()).Methods("GET")
	router.HandleFunc("/api/user", userHandler.Create()).Methods("POST")

	srv := s.NewServer(router)
	err = http.ListenAndServe(":8080", srv)
	if err != nil {
		fmt.Println(err)
	}
}
