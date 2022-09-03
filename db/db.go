package db

import (
	"database/sql"

	_ "github.com/lib/pq" // import implícito (driver do postgres)
)

func ConectaComBanco() *sql.DB {
	conexao := "user=postgres dbname=loja-alura host=localhost password=changeme sslmode=disable"
	db, err := sql.Open("postgres", conexao)

	if err != nil {
		panic(err.Error())
	}
	return db
}
