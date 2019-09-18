package main

// import "time"

type Event struct {
	Id 			int		  `json:"id"`
	Name        string    `json:"name"`
	// Date        time.Time `json:"date"`
	Description string    `json:"description"`
	Location    string    `json:"location"`
}

type Events []Event
