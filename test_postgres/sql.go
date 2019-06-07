package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq" //register its drivers with the database/sql package
)

const (
	host     = "192.168.3.15"
	port     = 5432
	user     = "postgres"
	password = "demopostgrespass"
	dbname   = "demopostgresdb"
)

type UserSummary struct {
	id        int
	age       int
	email     string
	firstName string
	lastName  string
}

func main() {
	usersInfo := querUser(1)
	for _, user := range usersInfo {
		fmt.Println(user)
	}
}

func initDB() *sql.DB {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname) //contains all of the information required to connect to our Postgres database
	db, err := sql.Open("postgres", psqlInfo) //simply validates the arguments provided
	if err != nil {
		panic(err)
	}

	err = db.Ping() //open up a connection to the database and validate connection string
	if err != nil {
		panic(err)
	}
	return db
}

func querUser(companID int) (usersInfo []UserSummary) {
	db := initDB()
	rows, err := db.Query("SELECT * FROM usersummary")
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	for rows.Next() {
		user := UserSummary{}
		err = rows.Scan(
			&user.id,
			&user.age,
			&user.email,
			&user.firstName,
			&user.lastName,
		)
		if err != nil {
			panic(err)
		}
		usersInfo = append(usersInfo, user)
	}

	err = rows.Err()
	if err != nil {
		panic(err)
	}
	return usersInfo
}
