package DB_Comunicate

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/mem"
	log "github.com/sirupsen/logrus"
	"strconv"
	"time"
)

const (
	host    = "localhost"
	port    = 5432
	dbname  = "postgres"
	logPath = "C:\\Program Files\\PostgreSQL\\13\\data\\log"
)

func Connect(userName string, password string) *sql.DB {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, userName, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Error(err)
		return nil
	}

	err = db.Ping()
	if err != nil {
		log.Error(err)
		return nil
	}
	fmt.Println("User connection ...\n Success")
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

	/*if _, err := os.Stat(logPath + "\\logs.txt"); errors.Is(err, os.ErrNotExist) {
		log.Error(err)
	}
	*/
	end := time.Now()
	delta := end.Sub(start)
	info := hardware(delta)
	defer rows.Close()
	return info, nil
}

func hardware(deltaTime time.Duration) (aboutSys string) {
	v, err := mem.VirtualMemory()
	if err != nil {
		log.Error(err)
	} else if int(v.UsedPercent) >= 95 {
		log.Error("Error: more memory has used")
	}

	aboutSys = "RAM: " + "Total: " + convert(v.Total) + " Free:" + convert(v.Free) +
		" Used:" + strconv.Itoa(int(v.UsedPercent)) + "%"

	state, err := cpu.Percent(deltaTime, true)
	if err != nil {
		log.Error(err)
	}

	fmt.Println("CPU")
	for i, j := range state {
		fmt.Println("Core", i, ":", j, "%")
		sum := +int(j)
		if sum/(i+1) > 95 {
			log.Error("Error: more CPU resources has used")
		}
	}
	return aboutSys
}

func convert(value uint64) string {
	mb := strconv.Itoa(int(int64(value / 1024 / 1024)))
	return mb + "Mb"
}
