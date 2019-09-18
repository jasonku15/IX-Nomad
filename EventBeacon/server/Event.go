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

type User struct {
	UserId string `json:"user_id"`
}

type Event_subscription struct {
	EventId string `json:"event_id"`
	Users   []User `json:"user_ids"`
}
