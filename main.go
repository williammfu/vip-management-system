package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/williammfu/vip-management-system/controller"
	"github.com/williammfu/vip-management-system/utils"
)

func main() {
	log.Println("Starting server. . .")

	router := mux.NewRouter()
	router.HandleFunc("/api/vips", controller.RetrieveAllVips).Methods("GET")
	router.HandleFunc("/api/vips", controller.StoreVip).Methods("POST")
	router.HandleFunc("/api/vips/{id}", controller.RetrieveVip).Methods("GET")
	router.HandleFunc("/api/vips/{id}", controller.UpdateVip).Methods("PUT")
	router.HandleFunc("/api/vips/{id}/arrived", controller.ArrivedVip).Methods("PATCH")
	router.Use(utils.BasicAuth)

	log.Fatal(http.ListenAndServe(":8000", router))
}
