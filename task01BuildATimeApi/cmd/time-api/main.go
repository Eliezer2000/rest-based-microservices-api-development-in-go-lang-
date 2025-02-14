package main

import (
	"log"
	"net/http"

	"github.com/eliezer2000/task01/internal/adapters/time_http"
	"github.com/eliezer2000/task01/internal/domain"
	"github.com/eliezer2000/task01/internal/ports"
)

func main() {

	var timeService ports.TimeProvider = &domain.TimeService{}

	timeHandler := &time_http.TimeHandler{TimeService: timeService}

	http.Handle("/api/time", timeHandler)
	log.Println("Server listening on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
