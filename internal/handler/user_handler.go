package handler

import (
	"encoding/json"
	"errors"
	m "github.com/cheeeasy2501/go-user-data-service/internal/model"
	r "github.com/cheeeasy2501/go-user-data-service/internal/repository"
	"net/http"
	"strconv"
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

		id, err, status := checkIdParameter(idParam)

		if err != nil {
			http.Error(writer, err.Error(), status)
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
		var user *m.User

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

func (uh *UserHandler) Update() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		var user *m.User

		err := json.NewDecoder(request.Body).Decode(&user)
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
		}

		user, err = uh.Repo.Update(user)

		if err != nil {
			http.Error(writer, err.Error(), http.StatusBadRequest)
		}

		err = json.NewEncoder(writer).Encode(&user)
		if err != nil {
			http.Error(writer, err.Error(), http.StatusBadRequest)
		}
	}
}

func (uh *UserHandler) Delete() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		idParam := request.URL.Query().Get("id")

		id, err, status := checkIdParameter(idParam)

		if err != nil {
			http.Error(writer, err.Error(), status)
		}

		err = uh.Repo.Delete(&id)

		if err != nil {
			http.Error(writer, err.Error(), http.StatusBadRequest)
		}
	}
}

func checkIdParameter(idParam string) (uint64, error, int) {

	if len(idParam) < 1 {
		return 0, errors.New("Id parameter not found"), http.StatusBadRequest
	}

	id, err := strconv.ParseUint(idParam, 10, 64)

	if err != nil {
		return 0, errors.New("Internal Error"), http.StatusInternalServerError
	}

	if id < 0 {
		return 0, errors.New("User not found"), http.StatusInternalServerError
	}

	return id, nil, 0
}
