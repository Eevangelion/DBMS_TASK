package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Sakagam1/DBMS_TASK/internal/config"
	"github.com/Sakagam1/DBMS_TASK/internal/db"
	in_memory "github.com/Sakagam1/DBMS_TASK/internal/db/in-memory"
	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

func main() {
	log.Printf("Server started")
	conf := config.New()
	port := conf.Server.Port
	db.DB = in_memory.SetupDB()
	router := NewRouter()

	log.Fatal(
		http.ListenAndServe(
			fmt.Sprintf("%s:%d", conf.Server.Address, port), router,
		),
	)
}
