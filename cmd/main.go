package main

import (
	"fmt"
	r "user-data-service/internal/repository"
	"user-data-service/pkg/db"

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
}
