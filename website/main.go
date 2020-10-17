/**
tinyserver
A SIMPLE LIGHT WEIGHT HTTP SERVER BY PROSPER for BRANKAS TEST
*/
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	_ "gitlab.com/Prosp3r/shopr/domainer/srsbt"
)

//Users -
type Users struct {
	Email string `json:"email"`
	//Domain string `json:"domani"`
}

//SystemMessage -
type SystemMessage struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

var allUsers []Users

const (
	port = ":2021"
)

func main() {

	//SITE USER INTERFACE ENDPOINTS
	http.HandleFunc("/", servehome)
	http.HandleFunc("/bt", servebtapp)
	http.HandleFunc("/faq", servefaq)
	http.HandleFunc("/features", features)
	http.HandleFunc("/terms", terms)
	http.HandleFunc("/checkmail", checkMail)

	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))

	log.Println("Listening at port " + port)
	http.ListenAndServe(port, nil)
}

func serveSingle(pattern string, filename string) {
	http.HandleFunc(pattern, func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, filename)
	})
}

//servehome will return the static home page html file
func servehome(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "index.html")
}

//servebtapp will return the static page for bt app
func servebtapp(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "btapp.html")
}

func servefaq(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "btfaq.html")
}

func terms(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "terms.html")
}

func features(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "features.html")
}

//checkMail - will take the parameters received from the front end and communicate with the user service to check for email availability
func checkMail(w http.ResponseWriter, r *http.Request) {

	//User
	var user Users

	//fmt.Println(r.Method)
	//fmt.Println(string(r.RequestURI))
	jsonDcode := json.NewDecoder(r.Body)
	err := jsonDcode.Decode(&user)
	if err != nil {
		panic(err)
	}

	fmt.Println(user.Email)

	//Connect to user service
	resp, err := http.Get("http://localhost:2000/checkmail/" + user.Email)
	if err != nil {
		fmt.Println(err)
	}
	var systemMessage SystemMessage
	resultDcode := json.NewDecoder(resp.Body)
	err = resultDcode.Decode(&systemMessage)
	if err != nil {
		panic(err)
	}
	//fmt.Println(systemMessage)
	json.NewEncoder(w).Encode(systemMessage)
}
