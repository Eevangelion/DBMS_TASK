package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Sakagam1/DBMS_TASK/internal/config"
)

func main() {
	log.Printf("Server started")
	conf := config.GetConfig()
	port := conf.Server.Port

	router := NewRouter()

	log.Fatal(
		http.ListenAndServe(
			fmt.Sprintf("%s:%d", conf.Server.Address, port), router,
		),
	)
}
