package main

import (
	"fmt"
	grpcUser "github.com/cheeeasy2501/go-user-data-service/grpc/user"
	h "github.com/cheeeasy2501/go-user-data-service/internal/handler"
	m "github.com/cheeeasy2501/go-user-data-service/internal/middleware"
	r "github.com/cheeeasy2501/go-user-data-service/internal/repository"
	"github.com/cheeeasy2501/go-user-data-service/pkg/db"
	s "github.com/cheeeasy2501/go-user-data-service/pkg/sever"
	"google.golang.org/grpc"
	"log"
	"net"
	"net/http"

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

	grpcS := grpc.NewServer()
	grpcUserSrv := grpcUser.NewGRPCServer(userRepo)
	grpcUser.RegisterUserServiceServer(grpcS, grpcUserSrv)

	//GRPC START
	go func() {
		l, err := net.Listen("tcp", ":8081")

		if err != nil {
			log.Fatal(err)
		}

		if err := grpcS.Serve(l); err != nil {
			fmt.Println("GRPC User Service not started!", err)
			log.Fatal(err)
		}
	}()

	//HTTP START
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
