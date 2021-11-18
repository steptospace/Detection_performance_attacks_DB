package DB_Comunicate

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/mem"
	log "github.com/sirupsen/logrus"
	"time"
)

const (
	host   = "localhost"
	port   = 5432
	dbname = "postgres"
)

func Connect(userName string, password string) *sql.DB {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, userName, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Error(err)
		return nil
		//Не забыть про RollBack!!!
	}

	err = db.Ping()
	if err != nil {
		log.Error(err)
		return nil
	}

	fmt.Println("User connection ... Success")

	return db
}

func Close(db *sql.DB) {
	db.Close()
}

//Input: User data (PostgreSQL script)
//Output: "Ok (all data about request)" or "Error: ..."
func StartCommunicate(db *sql.DB, textRequest string) (string, error) {
	start := time.Now()
	rows, err := db.Query(textRequest) // send command to database And execute
	if err != nil {
		log.Error(err)
	}
	end := time.Now()
	delta := end.Sub(start)
	hardware(delta)
	defer rows.Close()
	return "Success", nil
}

func hardware(deltaTime time.Duration) {
	v, _ := mem.VirtualMemory()
	fmt.Println("RAM:")
	fmt.Println("Total: ", toMb(v.Total), "\nFree:", toMb(v.Free), "\nUsed:", int(v.UsedPercent), "%")
	state, err := cpu.Percent(deltaTime, true)
	if err != nil {
		log.Error(err)
	}
	fmt.Println("CPU")
	for i, j := range state {
		fmt.Println("Core", i, ":", j, "%")

	}
	if int(v.UsedPercent) >= 95 {
		//report in program and print text in Logs
		// Error: more memory has used
	}

}

func toMb(value uint64) uint64 {
	return value / 1024 / 1024
}
