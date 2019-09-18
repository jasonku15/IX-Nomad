package main

// import "time"

type Event struct {
	Name        string `json:"name"`
	Time        string `json:"date"`
	Description string `json:"description"`
	Location    string `json:"location"`
}

type Events []Event
