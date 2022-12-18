package connection

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/Sakagam1/DBMS_TASK/internal/config"

	_ "github.com/jackc/pgx/v5/stdlib"
)

var Connection *sql.DB = nil

func GetConnectionToDB() (DB *sql.DB, err error) {
	if Connection == nil {
		conf := config.GetConfig()
		dbName := conf.Database.DbDBName
		dbHost := conf.Database.DbHost
		dbUserName := conf.Database.DbUserName
		dbPassword := conf.Database.DbPassword
		connection_information := fmt.Sprintf("host=%s dbname=%s user=%s password=%s", dbHost, dbName, dbUserName, dbPassword)
		Connection, err = sql.Open("pgx", connection_information)
		defer Connection.Close()
		if err != nil {
			log.Fatal(err)
			return nil, err
		}
		err = DB.Ping()
		if err != nil {
			log.Fatal("Connection Error:", err)
			return nil, err
		}
		return Connection, err
	}
	return Connection, nil
}
