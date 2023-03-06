package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/go-co-op/gocron"

	"github.com/Sakagam1/DBMS_TASK/internal/config"
	"github.com/Sakagam1/DBMS_TASK/internal/db"
	connection "github.com/Sakagam1/DBMS_TASK/internal/db/db_connection"
	_ "github.com/jackc/pgx/v5/stdlib"
)

func main() {
	// s.Every(1).Monday().At("00:00").Do(db.UserRepo.SetRemainingReports())

	log.Printf("Server started")
	conf := config.GetConfig()
	port := conf.Server.Port
	router := NewRouter()
	_, err := connection.GetConnectionToDB()
	if err != nil {
		log.Print("Error while connecting", err)
	}
	log.Println(conf.Server.Address, conf.Server.Port)
	defer connection.Connection.Close()

	s := gocron.NewScheduler(time.Local)
	s.Every(1).Monday().At("00:00").Do(db.UserRepo.SetRemainingReports)

	log.Fatal(
		http.ListenAndServe(
			fmt.Sprintf("%s:%d", conf.Server.Address, port), router,
		),
	)
}
