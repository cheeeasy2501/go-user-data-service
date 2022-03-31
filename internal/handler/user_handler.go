package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	r "user-data-service/internal/repository"
)

type UserHandler struct {
	Repo *r.UserRepository
}

func NewUserHandler(r *r.UserRepository) *UserHandler {
	return &UserHandler{Repo: r}
}

func (uh *UserHandler) GetAll() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		users, err := uh.Repo.GetAll()

		if err != nil {
			log.Println(err)
		}

		err = json.NewEncoder(writer).Encode(&users)
		if err != nil {
			log.Println(err)
		}
	}
}

func (uh *UserHandler) Get() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		id, err := strconv.ParseUint(request.URL.Query().Get("id"), 10, 64)

		if err != nil {
			// TODO set error into writer
		}

		if id < 0 {
			// TODO set error into writer
		}

		user := uh.Repo.GetById(id)

		err = json.NewEncoder(writer).Encode(&user)

		if err != nil {
			// TODO set error into writer
		}
	}
}
