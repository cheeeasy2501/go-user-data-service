package repository

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"
	m "user-data-service/internal/model"
	"user-data-service/pkg/db"
)

type UserRepository struct {
	mysql *db.Mysql
}

func NewUserRepo(mysql *db.Mysql) *UserRepository {
	return &UserRepository{
		mysql: mysql,
	}
}

func (ur *UserRepository) GetAll() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		var users []m.User

		query := "SELECT id, email, firstName, lastName, active FROM users  LIMIT 20"

		stmt, err := ur.mysql.Conn.Prepare(query)

		if err != nil {
			log.Println(err)
		}

		rows, err := stmt.Query()

		if err != nil {
			log.Println(err)
		}

		for rows.Next() {
			//TODO HOW I CAN NOT SEND PASSWORD ?
			user := m.User{
				Password: "",
			}

			err = rows.Scan(
				&user.Id,
				&user.Email,
				&user.FirstName,
				&user.LastName,
				&user.Active,
			)
			if err != nil {
				log.Println(err)
			}

			users = append(users, user)
		}

		err = json.NewEncoder(writer).Encode(&users)
		if err != nil {
			log.Println(err)
		}
	}
}

func (ur *UserRepository) Get() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		var user m.User

		id, err := strconv.Atoi(request.URL.Query().Get("id"))

		if err != nil {
			// TODO set error into writer
		}

		if id < 0 {
			// TODO set error into writer
		}

		query := fmt.Sprintf("SELECT id, email, firstName, lastName, active FROM users WHERE id=%d", id)

		stmt, err := ur.mysql.Conn.Prepare(query)

		if err != nil {
			// TODO set error into writer
		}

		row := stmt.QueryRow()

		err = row.Scan(
			&user.Id,
			&user.Email,
			&user.FirstName,
			&user.LastName,
			&user.Active)

		if err != nil {
			// TODO set error into writer
		}

		err = json.NewEncoder(writer).Encode(&user)
		if err != nil {
			// TODO set error into writer
		}
	}
}

func (ur *UserRepository) GetByEmail(email string) {
}

func (ur *UserRepository) Create() {
}

func (ur *UserRepository) Update() {
}

func (u *UserRepository) Delete(id uint64) error {

	if id <= 0 {
		fmt.Println("user id = ", id)
		return nil
	}

	return errors.New("Id < 0")
}
