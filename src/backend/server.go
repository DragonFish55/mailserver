package main

import (
	"net/http"
)

func addUser(w http.ResponseWriter, r *http.Request){
	
}

func server(ipAddr string){
	
	http.HandleFunc("/addUser", addUser);

	http.ListenAndServe(ipAddr,nil);
}

