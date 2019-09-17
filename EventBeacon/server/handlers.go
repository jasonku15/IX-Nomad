package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"github.com/gorilla/mux"
)

func CreateEvent(w http.ResponseWriter, r *http.Request) {
	stmt,err := db.Prepare("INSERT INTO events(name, description, location) VALUES('HACKATHON', 'hackathon', 'mtcc')")
	// body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err.Error())
	}
	// body, err := ioutil.ReadAll(r.Body)
	// if err != nil {
	// 	panic(err.Error())
	// }

	// keyVal := Event()
	// json.Unmarshal(body, &keyVal)
	// title := keyVal["title"]
	stmt.Exec()
	// if err != nil {
	// 	panic(err.Error())
	// }
	fmt.Fprintln(w, "CreateEvent!")
}

func EditEvent(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintln(w, "EditEvent!")
	// vars := mux.Vars(r)
	// eventID := vars["EventId"]
	// newDescription = vars["newDescription"]
	// timeSlot = vars["newTime"]
	// stmt, err := db.Prepare("UPDATE Events set description = newDescription, timeSlot = timeSlot where EventId = eventID")
	// fmt.Fprintln(w, "EventId:", eventID)
}

func GetEventDetails(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "GetEvent")
	vars := mux.Vars(r)
	eventID := vars["EventId"]
	fmt.Fprintln(w, "EventId:", eventID)
	result, err := db.Query("SELECT * from Events where eventID = " + eventID)
	if err != nil {
		panic(err)
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	reqBody, err := ioutil.ReadAll(r.Body)

	var ev Event
	json.Unmarshal(reqBody, &newEvent)
	err = result.Scan(&ev.Name, &ev.Description, &ev.Location)
	if err != nil {
		panic(err)
	}
}

func getPosts(w http.ResponseWriter, r *http.Request) {
	// w.Header().Set("Content-Type", "application/json")
	// var posts []Post
	// result, err := db.Query("SELECT id, title from posts")
	// if err != nil {
	// 	panic(err.Error())
	// }
	// defer result.Close()
	// for result.Next() {
	// 	var post Post
	// 	err := result.Scan(&post.ID, &post.Title)
	// 	if err != nil {
	// 		panic(err.Error())
	// 	}
	// 	posts = append(posts, post)
	// }
	// json.NewEncoder(w).Encode(posts)
}

func cancelEvent(w http.ResponseWriter, r *http.Request){
	// fmt.Fprintln(w, "Cancel Event!")
	// vars := mux.Vars(r)
	// eventID := vars["EventId"]
	// personName := vars["person"]
	// // check if event exists
	// if(false){
	// 	fmt.Fprintln(w, "EventId:", eventID)
	// }
	// // delte from both attendee and event tables
	// // stmt, err := db.Prepare("insert into Attendee values(eventID, personName)")
	// fmt.Fprintln(w, "EventId:", eventID)
}

func acceptEvent(w http.ResponseWriter, r *http.Request){
	// fmt.Fprintln(w, "Accept Event!")
	// vars := mux.Vars(r)
	// eventID := vars["EventId"]
	// personName := vars["person"]
	// // check if person already attends
	// if(true){
	// 	fmt.Fprintln(w, "EventId:", eventID)
	// }

	// stmt, err := db.Prepare("insert into Attendee values(eventID, personName)")
	// fmt.Fprintln(w, "EventId:", eventID)
}

func quitEvent(w http.ResponseWriter, r *http.Request){
	// fmt.Fprintln(w, "Quit Event!")
	// vars := mux.Vars(r)
	// eventID := vars["EventId"]
	// personName := vars["person"]
	// // check if person already quits
	// if(false){
	// }
	// stmt, err := db.Prepare("Delete From Attendee where criteria")
	// fmt.Fprintln(w, "EventId:", eventID)
}

func getAttendees(w http.ResponseWriter, r *http.Request){
	// w.Header().Set("Content-Type", "application/json")
	// var posts []Post
	// result, err := db.Query("SELECT id, names from attendees")
	// if err != nil {
	// 	panic(err.Error())
	// }
	// defer result.Close()
	// for result.Next() {

	// 	var post Post
	// 	err := result.Scan(&post.ID, &post.Title)
	// 	if err != nil {
	// 		panic(err.Error())
	// 	}
	// 	posts = append(posts, post)
	// }
	// json.NewEncoder(w).Encode(posts)
}
