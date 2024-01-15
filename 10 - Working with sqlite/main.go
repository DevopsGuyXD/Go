package main

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)
	
func main() {
	
	database, err := sql.Open("sqlite3", "./nraboy.db"); if err != nil{
		fmt.Println(err)
	}

	statement, err := database.Prepare("CREATE TABLE IF NOT EXISTS people(id INTEGER PRIMARY KEY AUTOINCREMENT, firstname TEXT, lastname TEXT)"); if err != nil{
		fmt.Println(err)
	}
	statement.Exec()

	statement, err = database.Prepare("INSERT INTO people(firstname, lastname) VALUES('Jim','Sanders')"); if err != nil{
		fmt.Println(err)
	}
	statement.Exec()

	rows, _ := database.Query("SELECT id, firstname, lastname FROM people")

	var id int
	var firstname string
	var lastname string

	for rows.Next(){
		rows.Scan(&id, &firstname, &lastname)
		fmt.Println(id)
		fmt.Println(firstname)
		fmt.Println(lastname)
	}
}