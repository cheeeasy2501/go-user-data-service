package main

import (
	"fmt"
	"net/http"
	r "user-data-service/internal/repository"
	"user-data-service/pkg/db"
	s "user-data-service/pkg/sever"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("User Data Service Started")

	mysqlCnf := db.NewConfig()
	mysql := db.NewMysql(mysqlCnf)

	userRepo := r.NewUserRepo(mysql)

	router := mux.NewRouter()
	router.HandleFunc("api/users", userRepo.GetAll()).Methods("GET")
	router.HandleFunc("api/user", userRepo.Get()).Methods("GET")

	srv := s.NewServer(router)
	err := http.ListenAndServe(":8080", srv)
	if err != nil {
		return
	}
}
