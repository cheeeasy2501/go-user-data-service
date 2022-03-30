package db

import (
	"database/sql"
)

type Mysql struct {
	cnf  *Config
	Conn *sql.DB
}

func NewMysql(cnf *Config) *Mysql {
	return &Mysql{
		cnf:  cnf,
		Conn: nil,
	}
}

func (mysql *Mysql) Open() error {

	conn, err := sql.Open("mysql", mysql.cnf.DatabaseURL)

	if err != nil {
		return err
	}

	err = conn.Ping()

	if err != nil {
		return err
	}

	mysql.Conn = conn

	return nil
}

func (mysql *Mysql) Close() error {
	err := mysql.Conn.Close()

	if err != nil {
		return err
	}

	return nil
}
