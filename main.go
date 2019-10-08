package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

const (
	host = "localhost"
	post = 5432
	user = "postgres"
	password = "postgres"
	dbname = "postgres"

)

func main(){
	datasource := fmt.Sprintf("host=%s port =%d user=%s password=%s dbname=%s sslmode=disable", 
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", datasource)
	defer db.Close() // when function bout to returmn, hshove this last command
	if err != nil{
		panic(err)
	}

}
func ping(db *sql.DB){

	err = db.Ping()
	if err != nil{
		panic(err)
		//log.Fatal(err)
	}
	fmt.Println("Successfully connected!")

	//db.Close()
}
}

func searchByName(db *sql.DB, searchvalue string)	{
	
	searchvalue := "Ivysaur"
row := db.QueryRow("SELECT * FROM pokemon WHERE name = $1", searchvalue)
var id int
var name string
row.Scan(&id, &name)
fmt.Println(id, name)

/*
rows, _ := db.Query("SELECT * FROM POKEMON")
for rows.Next(){
	var id int
	var name string
	rows.Scan(&id, &name)
	fmt.Println(id, name)*/