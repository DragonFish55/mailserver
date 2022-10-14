package main

import (
	
    _ "github.com/lib/pq"
)

//remote backend server address
var port string  = ":8080"
var saddress string = "http://127.0.0.1" + port

//remote backend database
var dbtype string = "postgres"
var dbaddr string = "localhost"
var dbport string = "5432"
var dbuser string = ""
var dbpass string = ""