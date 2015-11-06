package main

import "net/http"

func main() {
	// Creation of an URL shortener
	shortener := InitUrlShortener()
	
	// Creation of a personalized router
	router := MyRouter(shortener)
	
	// Start web server
	http.ListenAndServe(SERVER_IP+":"+SERVER_PORT, router)
}
