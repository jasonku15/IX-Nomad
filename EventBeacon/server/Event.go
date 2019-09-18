package main

// import "time"

type Event struct {
	Name        string `json:"name"`
	Id          string `json:"id"`
	Time        string `json:"date"`
	Description string `json:"description"`
	Location    string `json:"location"`
}

type Events []Event
