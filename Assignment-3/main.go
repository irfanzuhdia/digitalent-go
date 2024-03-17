package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"math/rand"
	"net/http"
	"time"
)

type Status struct {
	Water int `json:"water"`
	Wind  int `json:"wind"`
}

type StatusResponse struct {
	Status      Status `json:"status"`
	WaterStatus string `json:"waterStatus"`
	WindStatus  string `json:"windStatus"`
}

var status Status

func updateStatus() {
	for {
		status.Water = rand.Intn(100) + 1
		status.Wind = rand.Intn(100) + 1
		time.Sleep(2 * time.Second)
	}
}

func getStatus(w http.ResponseWriter, r *http.Request) {
	waterStatus := ""
	windStatus := ""

	if status.Water < 6 {
		waterStatus = "Aman"
	} else if status.Water >= 6 && status.Water <= 8 {
		waterStatus = "Siaga"
	} else {
		waterStatus = "Bahaya"
	}

	if status.Wind < 7 {
		windStatus = "Aman"
	} else if status.Wind >= 7 && status.Wind <= 15 {
		windStatus = "Siaga"
	} else {
		windStatus = "Bahaya"
	}

	response := StatusResponse{
		Status:      status,
		WaterStatus: waterStatus,
		WindStatus:  windStatus,
	}

	jsonResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	tpl, err := template.ParseFiles("index.html")
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	tpl.Execute(w, status)
}

func main() {
	rand.Seed(time.Now().UnixNano())

	go updateStatus()

	http.HandleFunc("/status", getStatus)
	http.HandleFunc("/", indexHandler)

	fmt.Println("Server is running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
