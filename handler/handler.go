package handler

import (
	"log"
	"net/http"

	"github.com/kennedysantiagoo/controler"

	"github.com/gorilla/mux"
)

func HandleRequests() {
	router := mux.NewRouter()
	router.HandleFunc("/clock/{hour}", controler.GetAngle).Methods("GET")
	router.HandleFunc("/clock/{hour}/{minute}", controler.GetAngle).Methods("GET")
	log.Println("Listening and serving on port :3000")
	log.Fatal(http.ListenAndServe(":3000", router))
}
