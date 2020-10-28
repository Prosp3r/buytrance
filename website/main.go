/**
tinyserver
A SIMPLE LIGHT WEIGHT HTTP SERVER BY PROSPER for BRANKAS TEST
*/
package main

import (
	"bufio"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"sync"

	"github.com/nats-io/go-nats"
	_ "gitlab.com/Prosp3r/shopr/domainer/srsbt"
)

//NATSHost - the connection string for NATS
var NATSHost = "https://localhost:4222"

//NATS_URL = os.Getenv("NATS_URL")

//Users -
type Users struct {
	Email string `json:"email"`
	//Domain string `json:"domani"`
}

//Sites -
type Sites struct {
	Email  string `json:"email"`
	Domain string `json:"domain"`
}

//Subdomains -
type Subdomains struct {
	DomainID     uint64 `json:"domainid"`
	Email        string `json:"email"`
	Domain       string `json:"domain"`
	CreationTime int64  `json:"creationtime"`
	Status       string `json:"status"`
}

//Searchdomains -
type Searchdomains struct {
	Domain string `json:"domain"`
}

//SubdomainList -
var SubdomainList []Subdomains

//var SubdomainList = make(map[uint64]map[string]string)

//AllUsers -
var AllUsers = make(map[uint64][]map[string]string)

//SystemMessage -
type SystemMessage struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

var allUsers []Users

const (
	port = ":2021"
)

var mtx = &sync.Mutex{}

//failOnError is  single place to handle errors to reduces number of keystrokes per each error handling call.
func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s %s", msg, err)
	}
}

/* ::::::::::::::::::::::::::::
:::::::::::::::::::::::::::::::*/
func main() {

	//SITE USER INTERFACE ENDPOINTS
	http.HandleFunc("/", servehome)
	http.HandleFunc("/bt", servebtapp)
	http.HandleFunc("/faq", servefaq)
	http.HandleFunc("/features", features)
	http.HandleFunc("/terms", terms)
	http.HandleFunc("/checkmail", checkMail)
	http.HandleFunc("/createsite", createSite)
	http.HandleFunc("/checkdomain", checkSubdomain)
	//http.HandleFunc("/checkdomainname", checkMail)

	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))

	log.Println("Website running on port " + port)
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
	//fmt.Println("here here mail...")
	//loadSubDomains()
	//User
	var user Users
	jsonDcode := json.NewDecoder(r.Body)
	err := jsonDcode.Decode(&user)
	if err != nil {
		panic(err)
	}

	//Connect to user service
	resp, err := http.Get("http://localhost:1980/checkmail/" + user.Email)
	if err != nil {
		fmt.Println(err)
	}
	var systemMessage SystemMessage
	//fmt.Printf("ResulDecode: %v \n", resp.Body)
	resultDcode := json.NewDecoder(resp.Body)
	err = resultDcode.Decode(&systemMessage)
	if err != nil {
		panic(err)
	}
	json.NewEncoder(w).Encode(systemMessage)
	return
}

//Check if subdomain is already used
func checkSubdomain(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Here here here domain ...")
	//refresh subdomain list
	loadSubDomains()

	/*var subD Searchdomains
	jsonDcode := json.NewDecoder(r.Body)
	err := jsonDcode.Decode(&subD)
	if err != nil {
		panic(err)
	}
	//domain
	domain := subD.Domain
	fmt.Println(domain)

	for _, v := range SubdomainList {
		if v.Domain == domain {
			fmt.Println("Used Domain")
			//domain is already used
			response := SystemMessage{
				Code:    domain,
				Message: "notavailable",
			}
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(response)
			return
		}
	}

	//domain is not yet used
	response := SystemMessage{
		Code:    domain,
		Message: "available",
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
	return*/
}

//createSite
func createSite(w http.ResponseWriter, r *http.Request) {

	var site Sites
	jsonDcode := json.NewDecoder(r.Body)
	err := jsonDcode.Decode(&site)
	if err != nil {
		panic(err)
	}

	//connect to nats
	nc, err := nats.Connect(NATSHost)
	failOnError(err, "Could not connect to NATS Streaming Server")
	payload := struct {
		Useremail    string
		Domainname   string
		Subscription string
		Status       string
	}{
		Useremail:    site.Email,
		Domainname:   site.Domain,
		Subscription: "basic",
		Status:       "creating",
	}

	natsMsg, err := json.Marshal(payload)
	failOnError(err, "Failed to create Json message for NATS")

	nc.Publish("createsite", natsMsg)
	nc.Publish("createsubdomain", natsMsg)
	fmt.Printf("Published event for %v \n", natsMsg)
	log.Println("Published site creation event")

	//resp, err := http.Get("http://localhost:1980/createsite/" + site.Email + "/" + site.Domain)
	//if err != nil {
	//	fmt.Println(err)
	//}
	response := SystemMessage{
		site.Domain,
		"sent",
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
	return
}

//writesSubdomain to file
func writeSubdomains() {

}

//loadsSubDomains from file
func loadSubDomains() {
	//domainListSource - change to env variable for production
	domainListSource := "domains.csv"

	//temporarily read from csv source - This can be change to badger or sql db
	dataFile, err := os.Open(domainListSource)
	defer dataFile.Close()
	failOnError(err, "Attempt to open user list source failed")

	reader := csv.NewReader(bufio.NewReader(dataFile))
	lines, err := reader.ReadAll()

	//Checking for empty file error type
	if err == io.EOF {
		failOnError(err, "Empty source file detected while reading line.")
	}

	//store user profile information
	for _, d := range lines {
		fmt.Println(d)
		domainID, err := strconv.ParseUint(string(d[0]), 0, 64)
		failOnError(err, "Failed converting line entry to Uint64")
		fmt.Printf("domainID : %v \n", domainID)

		userEmail := string(d[1])
		failOnError(err, "Failed to convert line entry to string")

		//
		domainName := string(d[2])
		failOnError(err, "Failed to convert line entry to string")

		domainCreationTime, err := strconv.ParseUint(string(d[3]), 0, 64)
		failOnError(err, "Failed to convert line entry to string")

		domainStatus := string(d[4])
		failOnError(err, "Failed to convert line entry to string")

		//now := time.Now()
		//nanos := uint64(now.UnixNano())

		//domainDetails
		dDetails := new(Subdomains)
		dDetails.DomainID = domainID
		dDetails.Email = userEmail
		dDetails.CreationTime = int64(domainCreationTime)
		dDetails.Domain = domainName
		dDetails.Status = domainStatus

		//mtx lock for read write concurrency protection
		mtx.Lock()
		SubdomainList = append(SubdomainList, *dDetails)
		mtx.Unlock()
	}
}
