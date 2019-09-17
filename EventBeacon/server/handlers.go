package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func CreateEvent(w http.ResponseWriter, r *http.Request) {
	stmt, err := db.Prepare("INSERT INTO events(title) VALUES(?)")
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err.Error())
	}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err.Error())
	}

	keyVal := Event()
	json.Unmarshal(body, &keyVal)
	title := keyVal["title"]
	_, err = stmt.Exec(title)
	if err != nil {
		panic(err.Error())
	}
	fmt.Fprintln(w, "CreateEvent!")
}

func EditEvent(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "EditEvent!")
	vars := mux.Vars(r)
	eventID := vars["EventId"]
	fmt.Fprintln(w, "EventId:", eventID)
}

func GetEvent(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "GetEvent")
	vars := mux.Vars(r)
	eventID := vars["EventId"]
	fmt.Fprintln(w, "EventId:", eventID)
	var event Event
	result, err := db.Query("SELECT id, title from Events")
	if err != nil {
		panic(err.Error())
	}
	defer result.Close()
	for result.Next() {
		var post Post
		err := result.Scan(&post.ID, &post.Title)
		if err != nil {
			panic(err.Error())
		}
		posts = append(posts, post)
	}
	json.NewEncoder(w).Encode(posts)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	//if err := json.NewEncoder(w).Encode(event); err != nil {
	//	panic(err)
	//}
}

func getPosts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var posts []Post
	result, err := db.Query("SELECT id, title from posts")
	if err != nil {
		panic(err.Error())
	}
	defer result.Close()
	for result.Next() {
		var post Post
		err := result.Scan(&post.ID, &post.Title)
		if err != nil {
			panic(err.Error())
		}
		posts = append(posts, post)
	}
	json.NewEncoder(w).Encode(posts)
}
