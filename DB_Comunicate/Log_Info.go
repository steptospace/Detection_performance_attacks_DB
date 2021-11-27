package DB_Comunicate

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	log "github.com/sirupsen/logrus"
)

// при запуске проги

func Formation(userName string, password string, request string) error {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, userName, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)

	if err != nil {
		log.Error(err)
	}

	err = db.Ping()
	if err != nil {
		log.Error(err)
	}
	fmt.Println("Successfully connected!")

	tx, err := db.Begin()
	if err != nil {
		return err
	}
	defer tx.Commit()

	res, err := db.Exec(request) // send command to database
	if err != nil {
		log.Error(err)
	}

	fmt.Println(res)

	return nil
}

func InitLogs() string {
	return "alter system set log_min_duration_statement = 0;" +
		"alter system set log_statement = 'all';" +
		"alter system set log_filename = 'logs.txt';" +
		"select pg_reload_conf();"
}

func CloseLog() string {
	return "alert system set log_filename = 'postgresql-%Y-%m-%d_%H%M%S.log';" +
		"select pg_reload_conf();"
}
