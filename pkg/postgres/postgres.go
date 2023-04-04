package postgres

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "abcd"
	dbname   = "postgres"
)

var PC *sql.DB

func Init() {
	// connection string
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	// open database
	db, err := sql.Open("postgres", psqlconn)
	if err != nil {
		panic(err)
	}
	PC = db

	// check db
	err = PC.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Connected!")
}

func InsertPage(id string, data string) error {
	insertDynStmt := `INSERT INTO "pages"("id", "data") VALUES($1, $2)`
	_, err := PC.Exec(insertDynStmt, id, data)
	return err
}

func GetDBPage(id string) ([]byte, error) {
	rows, err := PC.Query(`SELECT "id", "data" FROM "pages" WHERE "id"=$1`, id)
	var val []byte
	if err != nil {
		return val, err
	}

	defer rows.Close()
	for rows.Next() {
		var name string
		var roll string

		err = rows.Scan(&name, &roll)
		if err != nil {
			return val, err
		}
		val = []byte(roll)
		// fmt.Println(name, roll)
	}
	return val, err
}
