package main

import (
	"fmt"
	"strconv"
	"net/http"
	"encoding/json"
	"gopkg.in/redis.v3"
	"github.com/gorilla/mux"
)

type UrlShortener struct {
	redisClient *redis.Client
}

type JsonPost struct {
	url	string `json:"url"`
	custom	string `json:"custom"`
}

type JsonResponse struct {
	status		string `json:"status"`
	shortlink	string `json:"shortlink"`
}

func InitUrlShortener() *UrlShortener {
	// Initialize redis client
	shortener := &UrlShortener{
		redisClient : redis.NewClient(&redis.Options{
			Addr:     SERVER_IP+":"+DB_PORT,
			Password: DB_PASSWORD,
			DB:       0,
		}),
	}
	
	// Test redis connection
	//pong, err := shortener.redisClient.Ping().Result()
	//fmt.Println(pong, err)
	
	return shortener
}

func (shortener *UrlShortener) CreateShortlink(rw http.ResponseWriter, req *http.Request) {
	
	// First decode JSON sent by POST
	var jp JsonPost
	encoder := json.NewEncoder(rw)
	err := json.NewDecoder(req.Body).Decode(&jp)
	if err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		err := encoder.Encode(&JsonResponse{status: err.Error(), shortlink: ""})
		if err != nil {
			fmt.Fprintf(rw, "Error: %v", err.Error())
		}
		return
	}
	
	// Get original URL sent
	url := jp.url
	
	// In case of empty JSON (or no URL sent), default is Google URL (for test purpose)
	if len(url) == 0 {
		url = "http://www.google.com"
	}
	
	// Check if a custom parameter has been sent
	var shortlink string
	if len(jp.custom) > 0 {
		if(len(jp.custom) > MAX_CODE_LENTH) {
			// Limit size of custom to MAX_CODE_LENTH
			shortlink = jp.custom[:MAX_CODE_LENTH]
		} else {
			shortlink = jp.custom
		}
	} else {
		// Otherwise, use a random string as shortlink
		shortlink = randomString(MAX_CODE_LENTH)
	}
	
	// Check if the key (shortlink) is already used
	val, err := shortener.redisClient.Get(shortlink).Result()
	if err != redis.Nil || len(val) > 0 {
		// If so, we append a random number (converted into string) to the shortlink 
		shortlink += strconv.Itoa(randomNumber(MAX_RANDOM_NUMBER))
	}
	
	// In Reddis: key = shortlink; val = original URL
	err = shortener.redisClient.Set(shortlink, url, VALIDITY_DURATION).Err()
	if err != nil {
		rw.WriteHeader(http.StatusConflict)
		err := encoder.Encode(&JsonResponse{status: err.Error(), shortlink: ""})
		if err != nil {
			fmt.Fprintf(rw, "Error: %v", err.Error())
		}
		return
	}
	
	// Display shortlink for test
	fmt.Println(shortlink)
	
	encoder.Encode(&JsonResponse{status: "OK", shortlink: shortlink})
}

func (shortener *UrlShortener) Redirect(rw http.ResponseWriter, req *http.Request) {
	// Get the shortlink
	vars := mux.Vars(req)
	shortlink := vars["shortlink"]
	
	if len(shortlink) > 0 {
		// Retrieve original URL
		url, err := shortener.redisClient.Get(shortlink).Result()
		
		if len(url) > 0 {
			
			//TODO: save information in a log file
			
			// Note: for proper redirection, 'url' MUST start with "http" or "https"
			// Otherwise the path is considered relative instead of absolute
			http.Redirect(rw, req, url, http.StatusFound)
			
		} else if err == redis.Nil {
			fmt.Fprintf(rw, "Shortlink '%s' does not exist", shortlink)
			return
		} else if err != nil {
			fmt.Fprintf(rw, "Error: %v", err)
			return
		}
	}
}

func (shortener *UrlShortener) Monitor(rw http.ResponseWriter, req *http.Request) {
	//TODO: parse the log file and show result
}
