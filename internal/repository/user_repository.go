package repository

import (
	"context"
	"database/sql"
	"fmt"
	m "github.com/cheeeasy2501/go-user-data-service/internal/model"
	"github.com/cheeeasy2501/go-user-data-service/pkg/db"
	"log"
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

	stmt, err := ur.mysql.Conn.PrepareContext(context.TODO(), query)

	if err != nil {
		return nil, err
	}

	rows, err := stmt.QueryContext(context.TODO())

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

func (ur *UserRepository) GetById(id uint64) (*m.User, error) {
	var user m.User

	query := "SELECT id, email, password, firstName, lastName, active FROM users WHERE id=?"
	stmt, err := ur.mysql.Conn.PrepareContext(context.TODO(), query)

	if err != nil {
		return &user, err
	}

	row := stmt.QueryRowContext(context.TODO(), id)
	err = row.Scan(
		&user.Id,
		&user.Email,
		&user.Password,
		&user.FirstName,
		&user.LastName,
		&user.Active)

	if err != nil && err == sql.ErrNoRows {
		return &user, err
	}

	stmt.Close()

	return &user, nil
}

func (ur *UserRepository) GetByEmail(email string) (*m.User, error) {
	var user m.User

	query := fmt.Sprintf("SELECT id, email, firstName, lastName, active FROM users WHERE email=%s", email)
	stmt, err := ur.mysql.Conn.PrepareContext(context.TODO(), query)

	if err != nil {
		return &user, err
	}

	row := stmt.QueryRowContext(context.TODO())

	err = row.Scan(
		&user.Id,
		&user.Email,
		&user.FirstName,
		&user.LastName,
		&user.Active)

	if err != nil && err == sql.ErrNoRows {
		return &user, err
	}

	return &user, nil
}

func (ur *UserRepository) Create(m *m.User) (*m.User, error) {
	//query := fmt.Sprintf("INSERT INTO users VALUES (%d,%s,%s,%s,%s,%d)", m.Id, m.Email, m.Password, m.FirstName, m.LastName, m.Active)
	query := "INSERT INTO users (`email`, `password`, `firstName`, `lastName`, `active`) VALUES (?,?,?,?,?)"

	stmt, err := ur.mysql.Conn.PrepareContext(context.TODO(), query)
	defer stmt.Close()

	if err != nil {
		return m, err
	}

	res, err := stmt.ExecContext(context.TODO(), m.Email, m.Password, m.FirstName, m.LastName, m.Active)

	if err != nil {
		return m, err
	}

	rows, err := res.RowsAffected()

	if err != nil {
		log.Printf("Error %s when finding rows affected", err)
		return m, err
	}
	log.Printf("%d user created ", rows)

	return m, nil
}

func (ur *UserRepository) Update(m *m.User) (*m.User, error) {
	query := "UPDATE users SET firstName=?, lastName=? WHERE id=?"
	stmt, err := ur.mysql.Conn.PrepareContext(context.TODO(), query)

	if err != nil {
		return m, err
	}

	res, err := stmt.ExecContext(
		context.Background(),
		m.FirstName,
		m.LastName,
		m.Id,
	)
	if err != nil {
		return nil, err
	}

	rows, err := res.RowsAffected()

	if err != nil {
		log.Printf("Error %s when finding rows affected", err)
		return m, err
	}
	log.Printf("%d user updated ", rows)

	return m, nil
}

func (ur *UserRepository) Delete(id *uint64) error {
	query := "DELETE FROM users WHERE id=?"
	stmt, err := ur.mysql.Conn.PrepareContext(context.TODO(), query)

	if err != nil {
		return err
	}

	res, err := stmt.ExecContext(context.TODO(), id)
	if err != nil {
		return err
	}

	rows, err := res.RowsAffected()

	if err != nil {
		log.Printf("Error %s when finding rows affected", err)
		return err
	}
	log.Printf("%d user deleted ", rows)

	return nil
}
