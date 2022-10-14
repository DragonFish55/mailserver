package main

import (
	"net/http"
	"fmt"
	"io/ioutil"
	"encoding/json"
)

func createserver(serveraddr string){

	http.HandleFunc("/signup", regUser);
	http.HandleFunc("/signin", loginUser);
	http.HandleFunc("/logout", logoutUser);
	http.HandleFunc("/logout", logoutUser);
	http.ListenAndServe(saddress,nil);
}

func logoutUser(w http.ResponseWriter, r *http.Request){

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

