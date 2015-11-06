package main

import "time"

/*
 * This file should be a yaml file instead 
 */

 // Server and Database config
const SERVER_IP = "localhost"
const SERVER_PORT = "8080"
const DB_PORT = "6379"
const DB_PASSWORD = ""

// Validity duration for shortened urls: 3 months
var VALIDITY_DURATION = time.Hour * 24 * 30 * 3

// Max code length (used as a shortlink)
const MAX_CODE_LENTH = 6

// Max random number (a 3-digit number appended to shortlink if already used)
const MAX_RANDOM_NUMBER = 999
