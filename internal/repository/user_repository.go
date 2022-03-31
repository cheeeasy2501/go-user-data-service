package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
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

func (ur *UserRepository) GetById(id uint64) (m.User, error) {
	var user m.User

	query := fmt.Sprintf("SELECT id, email, firstName, lastName, active FROM users WHERE id=%d", id)
	stmt, err := ur.mysql.Conn.Prepare(query)

	if err != nil {
		return user, err
	}

	row := stmt.QueryRow()

	err = row.Scan(
		&user.Id,
		&user.Email,
		&user.FirstName,
		&user.LastName,
		&user.Active)

	if err != nil && err == sql.ErrNoRows {
		return user, err
	}

	return user, nil
}

func (ur *UserRepository) GetByEmail(email string) (m.User, error) {
	var user m.User

	query := fmt.Sprintf("SELECT id, email, firstName, lastName, active FROM users WHERE email=%s", email)
	stmt, err := ur.mysql.Conn.Prepare(query)

	if err != nil {
		return user, err
	}

	row := stmt.QueryRow()

	err = row.Scan(
		&user.Id,
		&user.Email,
		&user.FirstName,
		&user.LastName,
		&user.Active)

	if err != nil && err == sql.ErrNoRows {
		return user, err
	}

	return user, nil
}

func (ur *UserRepository) Create(m m.User) (m.User, error) {
	//query := fmt.Sprintf("INSERT INTO users VALUES (%d,%s,%s,%s,%s,%d)", m.Id, m.Email, m.Password, m.FirstName, m.LastName, m.Active)
	query := "INSERT INTO users (`email`, `password`, `firstName`, `lastName`, `active`) VALUES (?,?,?,?,?)"

	stmt, err := ur.mysql.Conn.Prepare(query)
	defer stmt.Close()

	if err != nil {
		return m, err
	}

	res, err := stmt.Exec(m.Email, m.Password, m.FirstName, m.LastName, m.Active)

	if err != nil {
		return m, err
	}

	row, err := res.RowsAffected()
	fmt.Println(row)
	if err != nil {
		log.Printf("Error %s when finding rows affected", err)
		return m, err
	}
	log.Printf("%d user created ", row)

	return m, nil
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
