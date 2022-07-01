package main


import (
	"log"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"os"
	"github.com/joho/godotenv"
)


func getConfig() (string, string, string, string) {
	if e := godotenv.Load(); e != nil {
		log.Fatal(fmt.Println(e))	
	}
	username := os.Getenv("db_user")
	password := os.Getenv("db_pass")
	dbName := os.Getenv("db_name")
	dbHost := os.Getenv("db_host")
	return username, password, dbName, dbHost
}


func GetDb() *sql.DB {
	username, password, dbName, dbHost := getConfig()
	db, err := sql.Open("postgres", fmt.Sprintf("user=%s password=%s dbname=%s host=%s sslmode=disable", 
		username, password, dbName, dbHost))
	if err != nil {
        panic(err)
    } 
    return db
}


func GetAmount(db *sql.DB, name string) (int, error) {
	res := db.QueryRow(fmt.Sprintf("select * from account where name='%s';", name))
	var username string
	var amount int
	err := res.Scan(&username, &amount)
	if err != nil {
        return -1, err
    }
	return amount, nil 
}
