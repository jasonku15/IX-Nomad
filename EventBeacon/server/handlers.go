package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
)

func CreateEvent(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	var event Event
	json.Unmarshal(reqBody, &event)
	stmt, err := db.Prepare("INSERT INTO events(name, description, location, time) VALUES(?, ?, ?, ?)")
	if err != nil {
		panic(err.Error())
	}
	_, e := stmt.Exec(event.Name, event.Description, event.Location, event.Time)
	if e != nil {
		panic(err.Error())
	}

	fmt.Printf("Created Event!")
}

func EditEvent(w http.ResponseWriter, r *http.Request) {
	//fmt.Printf("Editing Event!")
	//reqBody, _ := ioutil.ReadAll(r.Body)
	//var event Event
	//json.Unmarshal(reqBody, &event)
	//vars := mux.Vars(r)
	//stmt, err := db.Prepare("UPDATE events set name = ?, description = ?, time = ?, location = ?, status = ? where event_id = ?")
	//if err != nil {
	//	panic(err.Error())
	//}
	//_, e := stmt.Exec(event.Name, event.Description, event.Time, event.Location, event.Id)
	//if e != nil {
	//	panic(err.Error())
	//}
}

func GetEventDetails(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "GetEvent")
	vars := mux.Vars(r)
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	eventID := vars["eventid"]
	fmt.Fprintln(w, "EventId:", eventID)
	str := "SELECT * from events where event_id = " + eventID
	fmt.Fprintln(w, str)

	result, err := db.Query(str)
	if err != nil {
		panic(err)
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	// w.WriteHeader(http.StatusOK)

	var ev Event
	json.Unmarshal(reqBody, &ev)
	for result.Next() {
		err = result.Scan(&ev.Id, &ev.Name, &ev.Description, &ev.Location)
		if err != nil {
			panic(err)
		}
	}
	fmt.Fprintln(w, ev.Name)

	json.NewEncoder(w).Encode(ev)

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

func cancelEvent(w http.ResponseWriter, r *http.Request) {
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

func acceptEvent(w http.ResponseWriter, r *http.Request) {
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

func quitEvent(w http.ResponseWriter, r *http.Request) {
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

func getAttendees(w http.ResponseWriter, r *http.Request) {
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
