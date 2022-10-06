package database

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func ConnectWithDatabase() *sql.DB {
	connection := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
					"0.0.0.0", 5051, "postgres", "qwerty", "alura_loja")
	
	db, err := sql.Open("postgres", connection)
	
	if err != nil {
		panic(err.Error())
	}
	
	return db
}
