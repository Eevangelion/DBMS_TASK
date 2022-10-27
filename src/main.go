package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Sakagam1/DBMS_TASK/src/config"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

func CreateRouter() *mux.Router {
	router := mux.NewRouter()
	// router.HandleFunc("/").Methods("GET")
	// router.HandleFunc("/").Methods("POST")
	http.Handle("/", router)
	return router
}

func main() {
	log.Printf("Server started")
	conf := config.New()
	router := CreateRouter()
	log.Fatal(
		http.ListenAndServe(
			fmt.Sprintf("%s:%s", conf.Server.Address, conf.Server.Port), router,
		),
	)
}
