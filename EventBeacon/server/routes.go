package main

import (
	"net/http"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{
		"CreateEvent",
		"POST",
		"/createevent",
		CreateEvent,
	},
	Route{
		"GetEvent",
		"GET",
		"/getevent/{eventid}",
		GetEventDetails,
	},
	Route{
		"EditEvent",
		"GET",
		"/editevent/{eventid}",
		EditEvent,
	},
	Route{
		"AcceptEvent",
		"GET",
		"/acceptevent/{eventid}/{userid}",
		AcceptEvent,
	},
	Route{
		"QuitEvent",
		"GET",
		"/quitevent/{eventid}/{userid}",
		QuitEvent,
	},
	Route{
		"GetAtendees",
		"GET",
		"/getAttendees/{eventid}",
		QuitEvent,
	},
}
