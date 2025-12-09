package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Response struct {
	Message string `json:"message"`
	Status  string `json:"status"`
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		resp := Response{
			Message: "API Operativa - Hackademy - Bustos Anahi - Prof Santiago Abastante",
			Status:  "Secure",
		}

		json.NewEncoder(w).Encode(resp)
	})

	fmt.Println("Servidor Backend corriendo en puerto 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
