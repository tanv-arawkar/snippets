package postgres_conn

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

const (
	host = "localhost"
	port = 5432
	user = "postgres"
	password = "root"
	dbname = "contact_list_management"
)

type DB struct {
	*sql.DB
}

//type postgres_conner interface {
//	
//	connect_to_db() (db *sql.DB)
//    register_new_user(db *sql.DB, name string,username string,password string,phone_number string,dob string) 
//}

func Connect_to_db() (*DB) {
	postgresqlDbInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
    host, port, user, password, dbname)

	db ,err := sql.Open("postgres",postgresqlDbInfo)

	if err != nil {
		  panic(err)
	}

	err = db.Ping()
  	if err != nil {
    	panic(err)
  	}
	
  fmt.Println("Established a successful connection!")

  return &DB{db}

}

func Register_new_user(db *DB, name string,username string,password string,phone_number int,dob string) {
  sqlStatement := "INSERT INTO contact_list (name, username, password, phone_number, dob) VALUES ($1, $2, $3, $4, $5) RETURNING id"
  id := -1
  err := db.QueryRow(sqlStatement, name, username,password,phone_number,dob).Scan(&id)
  if err != nil {
    panic(err)
  }
  fmt.Println("New user registered. ID is:", id)
}

func Validate_existing_user(db *DB, username string,password string) (bool,int){
//	id:= -1
//	sqlStatement := "SELECT name FROM contact_list where id = 2;"
//
//	err := db.QueryRow(sqlStatement,username).Scan(&id)
//	 if err != nil {
//            return false,-2
//     } else {
//        return true,id
//	 }	
//	 if err != nil {
//	  panic(err)
//    }

//	if id != -1 {
//		return true
//	} else {
//		return false
//	}	
	//TODO
	return true,2
}

func update_existing_data() {
	//TODO
}

func delete_existing_user() {
	//TODO
}
