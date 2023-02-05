package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Sakagam1/DBMS_TASK/internal/config"
	connection "github.com/Sakagam1/DBMS_TASK/internal/db/db_connection"
	_ "github.com/jackc/pgx/v5/stdlib"
)

func main() {
	log.Printf("Server started")
	conf := config.GetConfig()
	port := conf.Server.Port
	router := NewRouter()
	_, err := connection.GetConnectionToDB()
	if err != nil {
		log.Print("Error while connecting", err)
	}
	log.Println(conf.Server.Port)
	defer connection.Connection.Close()
	log.Fatal(
		http.ListenAndServe(
			fmt.Sprintf("%s:%d", conf.Server.Address, port), router,
		),
	)

}
