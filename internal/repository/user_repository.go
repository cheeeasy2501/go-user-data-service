package repository

import (
	"errors"
	"fmt"
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

func (ur *UserRepository) GetAll() ([]m.User, error) {
	var users []m.User

	query := "SELECT id, email, firstName, lastName, active FROM users  LIMIT 20"

	stmt, err := ur.mysql.Conn.Prepare(query)

	if err != nil {
		return nil, err
	}

	rows, err := stmt.Query()

	if err != nil {
		return nil, err
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
			return nil, err
		}

		users = append(users, user)
	}

	return users, nil
}

func (ur *UserRepository) GetById(id uint64) m.User {
	var user m.User

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

	return user
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
