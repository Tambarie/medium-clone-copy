package Database

import (
	"database/sql"
	"fmt"
)

const(
	host = "localhost"
	port = 5432
	user = "decagon"
	dbname = "mini_blog"
)


func ConnectDB() *sql.DB {
	var psqInfo = fmt.Sprintf("host=%s port=%d user=%s dbname=%s sslmode=disable",host,port,user,dbname)
	Db, err := sql.Open("postgres",psqInfo)
	if err != nil{
		panic(err)
	}
	return Db
}
