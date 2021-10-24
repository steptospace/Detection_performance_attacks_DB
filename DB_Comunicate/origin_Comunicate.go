package DB_Comunicate

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

const (
	host   = "localhost"
	port   = 5432
	dbname = "postgres"
)

func ErrorEx(err error) {
	recover()
	fmt.Println(err)
}

//Input: User data (PostgreSQL script)
//Output: "Ok (all data about request)" or "Error: ..."
func StartCommunicate(userName string, password string, textRequest string) string {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, userName, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		defer ErrorEx(err)
		panic(err)
		//Не забыть про RollBack!!!
	}

	defer db.Close()
	err = db.Ping()
	if err != nil {
		defer ErrorEx(err)
		panic(err)
	}

	fmt.Println("Successfully connected!")

	rows, err := db.Query(textRequest)
	if err != nil {
		defer ErrorEx(err)
		panic(err)
	}

	fmt.Println(rows)
	defer rows.Close()

	return "Success"
}
