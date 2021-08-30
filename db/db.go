package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func ConectaComDB() *sql.DB {
	connDB := "user=danfmsouza dbname=aluraDani password='IPai6xk3!!' host=192.168.15.187 sslmode=disable"
	db, err := sql.Open("postgres", connDB)
	if err != nil {
		fmt.Println("Deu ruim" + err.Error())
		panic(err.Error())
	}
	return db
}
