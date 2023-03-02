package main

import ( "postgres_conn"
		 "fmt"
)

func main() {

	db := postgres_conn.Connect_to_db()
    //postgres_conn.Register_new_user(db,"Tanu","tan","qwerty123",9228,"06-04-1996")
	ans,id:= postgres_conn.Validate_existing_user(db,"tan","qwerty123")

	fmt.Println("%v %v",ans,id)
}
