package app

import (
	"business/pkg/config"
	"business/pkg/db"
	mylog "business/pkg/log"
	"business/pkg/repo"
	"strings"
	"time"
)

const (
	logFileName = "./log/worker_%s.log"
)

//Run ...
func Run() {

	//Initialize config
	{
		config.Init("./configs/config.ini")
	}

	//Initalize logger
	{
		logFileName := strings.Replace(logFileName, "%s", time.Now().Format("2006_01_02___15_04_05"), 1)
		mylog.Init(logFileName, config.GetLogLevel())
	}

	//Load database configs
	{
		dbconf := config.LoadDBConfigs("DB")
		db.SetConfigDatabase(&dbconf)
		db.Init()
		//Init safecity db
		sdbconf := config.LoadDBConfigs("DBS")
		db.SetConfigDatabaseSS(&sdbconf)
		db.InitSS()
	}

	// infinite print loop
	Nsecs := 30000

	for {

		
		time.Sleep(time.Millisecond * time.Duration(Nsecs))
		update()

	}
}

func update() {

	cs := (&repo.Customer{}).GetCustomers()

	(&repo.Customer{}).UpdateViolation(cs)
}
