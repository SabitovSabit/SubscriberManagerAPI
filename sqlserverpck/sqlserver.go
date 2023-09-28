package sqlserver

import (
	"database/sql"
	"fmt"
     helper "subscriptionApi/helperpck"
	_ "github.com/denisenkom/go-mssqldb"
)

var Db *sql.DB

func Init() {
	constr := helper.GetValue().SqlConnStr
	var err error
	Db, err = sql.Open("mssql", constr)
    helper.File.Close()
	if err != nil {
		fmt.Println(err.Error())
	}
}
