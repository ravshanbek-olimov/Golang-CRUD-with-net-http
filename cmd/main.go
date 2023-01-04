package main

import (
	"fmt"
	"log"
	"net/http"

	"ravshanbek-olimov/Golang-CRUD-with-net-http/config"
	"ravshanbek-olimov/Golang-CRUD-with-net-http/controller"
	"ravshanbek-olimov/Golang-CRUD-with-net-http/pkg/db"
)

func main() {

	cfg := config.Load()

	conn, err := db.NewConnectPostgres(cfg)
	if err != nil {
		log.Fatal("error database connection: ", err.Error())
	}

	cont := controller.NewController(conn)

	http.HandleFunc("/movie", cont.Movie)

	fmt.Println("Listening", cfg.HTTPPort)
	err = http.ListenAndServe(cfg.HTTPPort, nil)
	if err != nil {
		log.Fatal("error listening server: ", err.Error())
	}
}
