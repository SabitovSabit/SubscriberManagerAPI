package sqlserver

import (
	"database/sql"
	"fmt"

	_ "github.com/denisenkom/go-mssqldb"
)

var Db *sql.DB

func Init() {
	server := "DESKTOP-LG63R1M"
	port := 1433
	database := "SubscriberDb"

	conStr := fmt.Sprintf("server=%s;port=%d;database=%s;integrated security=true;", server, port, database)
	var err error
	Db, err = sql.Open("mssql", conStr)

	if err != nil {
		fmt.Println(err.Error())
	}
}
