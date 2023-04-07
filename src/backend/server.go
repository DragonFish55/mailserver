package main

import (
	"net/http"
	"fmt"
	"io/ioutil"
	"encoding/json"
)

func create_server(address: string) {
	
	handler := http.NewServeMux()
	s := &http.Server{
		Addr:address,
		Handler: handler
	}

	servers[address] = s

	http.HandleFunc("/signup", regUser);
	http.HandleFunc("/signin", loginUser);
	http.HandleFunc("/logout", logoutUser);
	http.HandleFunc("/logout", logout_user);
	http.HandleFunc("/shutdown", shutdown_server(s));
	http.HandleFunc("/newserver", create_server());
	
	http.ListenAndServe(saddress,nil);
}




func shutdown_server(w http.ResponseWriter, r *http.Request, server http.Server){
	w.Write([]byte("OK"))
    go func() {
        if err := s.Shutdown(context.Background()); err != nil {
            log.Fatal(err)
        }
    }()
}

func check_server_status(server http.Server) bool {

}

func logout_user(w http.ResponseWriter, r *http.Request){

}

//login user account
func loginUser(w http.ResponseWriter, r *http.Request){
	
	userInfo, err := ioutil.ReadAll(r.Body);
	if err != nil{ fmt.Println("Error add user reading json data"); }
	var user_reg UserLogin = UserLogin{Username:"", Password:"",};
	var key string;
	var pass string;
	err = json.Unmarshal([]byte(userInfo), &user_reg);
	if err != nil { 
		fmt.Println("Error extracting json add user data");
		w.Write([]byte(`{"valid":"true"}`));
	}
	sqlconn, err := createDbConnect("Users");
	if err != nil { 
		fmt.Println("Could not connect to db");
		w.Write([]byte(`{"valid":"true"}`));
	};
	querRows, err := sqlconn.Query("SELECT password, enckey FROM USER WHERE username = ?", user_reg.Username)
	if err != nil { 
		fmt.Println("Could not query User table");
		w.Write([]byte(`{"valid":"true"}`));
	};
	
	defer querRows.Close();

	for querRows.Next() {
        err := querRows.Scan(&pass, &key)
        if err != nil {
			panic(err)
		} else {
			if checkUserValid(pass, key) {
				fmt.Println("Password is Valid");
				w.Write([]byte(`{"valid":"true"}`))
				break;
			}
			fmt.Println("Password is Invalid");
			w.Write([]byte(`{"valid":"false"}`))
			break;
		}
    }
}

// Add new user account to email database
func regUser(w http.ResponseWriter, r *http.Request){
	var user_reg UserReg = UserReg{
		Username: "",
		Password: "",
		Confpass: "",
		Fname: "",
		Lname: "",
		Bday: "",
	};

	userInfo, err := ioutil.ReadAll(r.Body);
	if err != nil{ fmt.Println("Error add user reading json data"); }
	err = json.Unmarshal(userInfo, &user_reg);
	if err != nil { 
		fmt.Println("Error extracting json add user data");
		w.Write([]byte(`{"valid":"true"}`));
	}
}

