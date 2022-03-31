package handler

import (
	"encoding/json"
	"net/http"
	"strconv"
	m "user-data-service/internal/model"
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
			http.Error(writer, err.Error(), http.StatusInternalServerError)
		}

		err = json.NewEncoder(writer).Encode(&users)
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
		}
	}
}

func (uh *UserHandler) Get() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {

		idParam := request.URL.Query().Get("id")
		//emailParam := request.URL.Query().Get("email")

		if len(idParam) < 1 {
			http.Error(writer, "Id parameter not found", http.StatusBadRequest)
			return
		}

		id, err := strconv.ParseUint(idParam, 10, 64)

		if err != nil {
			http.Error(writer, "Internal Error", http.StatusInternalServerError)
			return
		}

		if id < 0 {
			http.Error(writer, "User not found", http.StatusInternalServerError)
			return
		}

		user, err := uh.Repo.GetById(id)

		if err != nil {
			http.Error(writer, "User not found", http.StatusNotFound)
			return
		}

		err = json.NewEncoder(writer).Encode(&user)

		if err != nil {
			http.Error(writer, "Internal Error", http.StatusInternalServerError)
			return
		}
	}
}

func (uh *UserHandler) Create() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		var user m.User

		err := json.NewDecoder(request.Body).Decode(&user)
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
		}

		user, err = uh.Repo.Create(user)

		if err != nil {
			http.Error(writer, err.Error(), http.StatusBadRequest)
		}

		err = json.NewEncoder(writer).Encode(&user)
		if err != nil {
			http.Error(writer, err.Error(), http.StatusBadRequest)
		}
	}
}
