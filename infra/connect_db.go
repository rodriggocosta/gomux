package infra

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

func Connect() (*sql.DB, error) {
	conn := "user=postgres dbname=vendas password=root@06 port=5432 host=localhost sslmode=disable"

	db, err := sql.Open("postgres", conn)

	if err != nil {
		panic(err)
	}

	err = db.Ping()

	if err != nil {
		panic(err)
	}

	fmt.Println("connect ok")

	return db, nil
}
