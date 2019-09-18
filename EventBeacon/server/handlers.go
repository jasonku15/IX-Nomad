package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"github.com/gorilla/mux"
	"encoding/json"
	"net/http/httputil"
	"strings"
	"log"

	//"github.com/nlopes/slack"
)

type SlackResponse struct {
	response_type string `json:"response_type"`
	text string `json:"text"`
}

type SlackRequest struct {
	response_url string
}

func SetupEvent(w http.ResponseWriter, r *http.Request) {

	/*scp, err := slack.SlashCommandParse(r)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}*/
	
	/*mux.Vars(r)
	responseURL := vars["response_url"]
	if responseURL == "" {
		fmt.Println("error response URL")
		http.Error(w, "No response URL", http.StatusInternalServerError)
		return
	}
	fmt.Println("responseURL: ", responseURL)*/

	/*decoder := json.NewDecoder(r.Body)
	var slackRequest SlackRequest
	err := decoder.Decode(&slackRequest)
	if err != nil || slackRequest.response_url == "" {
		fmt.Println("Error:", err.Error())
		fmt.Println("error response URL")
		http.Error(w, "No response URL", http.StatusInternalServerError)
		return
	}
	responseURL := slackRequest.response_url
	fmt.Println("response URL: ", slackRequest.response_url)*/

	requestDump, err := httputil.DumpRequest(r, true)
	if err != nil {
	fmt.Println(err)
	}
	fmt.Println(string(requestDump))

	r.ParseForm()
	if r.FormValue("response_url") == "" {
		fmt.Println("Missing response URL")
		return
	}
	responseURL := r.FormValue("response_url")
	fmt.Println("response URL: ", responseURL)
	
	/*slackResponse := SlackResponse {
		response_type: "in_channel",
		text: "It's 80 degrees right now.",
	}
	sr, err := json.Marshal(slackResponse)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}*/

	jsonData := `{
		"text": "It's 80 degrees right now.",
		"attachments": [
			{
				"text":"Partly cloudy today and tomorrow"
			}
		]
	}`
	jdReader := strings.NewReader(jsonData)
	//_, err = http.Post(responseURL, "application/json", bytes.NewBuffer(sr))
	_, err = http.Post(responseURL, "application/json", jdReader)
	if err != nil {
		log.Fatalln(err)
	}

	//json.NewEncoder(w).Encode(slackResponse)

	//w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	//w.Header().Add("Content-Type", "application/json")
	//fmt.Fprintln(w, "Hello World")
	//w.Write([]byte("Hello World"))

}

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
