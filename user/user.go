package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net"
	"net/http"

	btuser "github.com/Prosp3r/buytrance/user/pb"
	"github.com/dgraph-io/dgo/protos/api"
	"github.com/gorilla/mux"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/dgraph-io/dgo"
)

var dgServer = "localhost:9080"

//DgraphCONN - Opens connection to Dgraph database.
func DgraphCONN() *dgo.Dgraph {
	conn, err := grpc.Dial(dgServer, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect to DGraph gRPC: \n %+v \n", err)
	}
	defer conn.Close()
	dc := api.NewDgraphClient(conn)
	dg := dgo.NewDgraphClient(dc)

	return dg
}

/*/User type
type Usertype struct {
	UserUID    string `json:"useruid,omitempty"`
	Fullname   string `json:"fullname,omitempty"`
	Email      string `json:"email,omitempty"`
	Phone      string `json:"phone,omitempty"`
	Datejoin   string `json:"lastlogin,omitempty"`
	Verifycode string `json:"verifycode,omitempty"`
	Status     string `json:"status,omitempty"`
}*/

//User - holds data structure for user type
type user struct{}

//
func main() {
	grpcport := ":8090"
	httpport := ":8091"

	go func() {
		startJSONAPI(httpport)
	}()

	lis, err := net.Listen("tcp", grpcport)
	if err != nil {
		log.Fatalf("Failed to listen on port %v with error: %v \n", grpcport, err)
	}
	fmt.Printf("BTUser Server started on port %v \n", grpcport)

	s := grpc.NewServer()
	btuser.RegisterUSERServiceServer(s, &user{})

	//bt.RegisterORIServiceServer(s, &server{})
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed serving User with error : %v \n", err)
	}
}

//startJSONAPI - Will start an http api that accepts json payload
func startJSONAPI(port string) {

	//Start Gorrila mux
	router := mux.NewRouter()

	router.HandleFunc("/", Infopage).Methods("GET")
	router.HandleFunc("/adduser", AddUser).Methods("POST")
	router.HandleFunc("/getuser", GetUser).Methods("GET")
	router.HandleFunc("/getusers", GetUsers).Methods("GET")
	router.HandleFunc("/updateUser", UpdateUser).Methods("PUT")
	router.HandleFunc("/deactivateuser", DeactivateUser).Methods("PUT")

	fmt.Printf("Starting server at port %s \n", port)
	log.Fatal(http.ListenAndServe(port, router))

}

//Infopage - Displays welcome information about the API to users
func Infopage(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Hi, you've reached the User service page \n\n This service helps manage users of the app \n\n  The following end points require JSON Payloads \n /adduser to add a new user to the system \n /getuser to fetch a single user's information \n /getusers with an s to fetch all user's information in a list \n /updateUser to update a user's information \n /deactivateuser (administrative) to flag a user's profile as inactive \n ")
}

//AddUser - Method will add a new user to the system
//Return - string UniqueUserID
func (c *user) AddUser(ctx context.Context, r *btuser.User) (*btuser.UserResponse, error) {

	var response *btuser.UserResponse
	//check if user exists
	userEmail := r.Email
	if IsUserSignedUp(userEmail) == false {
		//User is already in the system
		//response.User, err = AddUserToDGraph(r)

		adduser, err := AddUserToDGraph(r)
		if err != nil {
			log.Fatalf("Failed adding new user with error: \n %+v \n", err)
			return nil, err
		}
		response.User = adduser
		//return response, nil
	}

	//store in database if not
	//If already exists, return user information and send alert code to email or phone number on file for login
	//response.User = adduser //temporary assignment this will change

	return response, nil
}

//AddUser - for Handling JSON API Requests mapped to AddUser gRPC method
//param - http.Request with JSON Payload [{username, email, phone, }]
//returns json response http response writer
func AddUser(w http.ResponseWriter, r *http.Request) {
	//set status
	//w.WriteHeader(http.StatusOK)
}

//GetUser - will return user's detailed information
func (c *user) GetUser(ctx context.Context, r *btuser.IDString) (*btuser.UserResponse, error) {
	var response *btuser.UserResponse
	return response, nil
}

//GetUser - for Handling JSON API Requests mapped to GetUser gRPC method
//param - http.Request with JSON Payload
//returns JSON Response to http response writer
func GetUser(w http.ResponseWriter, r *http.Request) {
	//set status
	//w.WriteHeader(http.StatusOK)
}

//GetUsers -  will return a list of users
func (c *user) GetUsers(ctx context.Context, r *btuser.Emptyentry) (*btuser.Users, error) {
	var response *btuser.Users
	return response, nil
}

//GetUsers - for Handling JSON API Requests mapped to GetUsers gRPC method
//param - http.Request with JSON Payload
//returns JSON Response to http response writer
func GetUsers(w http.ResponseWriter, r *http.Request) {
	//set status
	//w.WriteHeader(http.StatusOK)
}

//UpdateUser - will update a user's information
func (c *user) UpdateUser(ctx context.Context, r *btuser.User) (*btuser.UserResponse, error) {
	var response *btuser.UserResponse
	return response, nil
}

//UpdateUser - for Handling JSON API Requests mapped to UpdateUser gRPC method
//param - http.Request with JSON Payload
//returns JSON Response to http response writer
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	//set status
	//w.WriteHeader(http.StatusOK)
}

//DeactivateUser - Will flag a user's profile/account deactivated
func (c *user) DeactivateUser(ctx context.Context, r *btuser.User) (*btuser.BOOLValue, error) {
	var response *btuser.BOOLValue
	return response, nil
}

//DeactivateUser - for Handling JSON API Requests mapped to DeactivateUser gRPC method
//param - http.Request with JSON Payload
//returns JSON Response to http response writer
func DeactivateUser(w http.ResponseWriter, r *http.Request) {
	//set status
	//w.WriteHeader(http.StatusOK)
}

//UpdateUser - will update a user's information
//This is a temporary fix for cache issues resulting from typo in method name UpdateUser in proto file
func (c *user) UdateUser(ctx context.Context, r *btuser.User) (*btuser.UserResponse, error) {
	var response *btuser.UserResponse
	return response, nil
}

//:::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::
//::::::::::::::::::::::USER UTILITY ::::::::::::::::::::::::::::::::
//:::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::

//UserSchemas -
var UserSchemas = `username: string .
email: string @index(exact) .
accountstatus: string .
verificationcode string .
codestatus string .
datejoined datetime .
lastupdated datetime .

type User struct {
	username
	email
	accountstatus
	verificationcode
	codestatus
	datejoin
	lastupdated
}`

//AddUserToDGraph - adds a new user node to the DGraph database
//Takes in a variable *btuser.User struct and returns *api.Response or error
func AddUserToDGraph(userInfoJSON *btuser.User) (*btuser.User, error) {
	dgraph := DgraphCONN()
	op := &api.Operation{}
	op.Schema = `
	username: string .
	email: string @index(exact) .
	accountstatus: string .
	verificationcode string .
	codestatus string .
	datejoined datetime .
	lastupdated datetime .

	type User {
		username
		email
		accountstatus
		verificationcode
		codestatus
		datejoin
		lastupdated
	}`
	ctx := context.Background()
	err := dgraph.Alter(ctx, op)
	if err != nil {
		log.Fatal(err)
	}
	//datejoined := time.Now()
	//dob := time.Date(1980, 01, 01, 23, 0, 0, 0, time.UTC)
	mu := &api.Mutation{
		CommitNow: true,
	}
	pb, err := json.Marshal(userInfoJSON)
	if err != nil {
		log.Fatal(err)
	}

	mu.SetJson = pb
	assigned, err := dgraph.NewTxn().Mutate(ctx, mu)
	if err != nil {
		log.Fatalf("Failed with error: \n %+v \n", err)
	}

	variables := map[string]string{"$email": assigned.Uids[userInfoJSON.Email]}
	query := `query me($email: string){
		me(func: email($email)){
			uid
			username
			email
			accountstatus
			verificationcode
			codestatus
			datejoin
			lastupdated
		}
	}`
	resp, err := dgraph.NewTxn().QueryWithVars(ctx, query, variables)
	if err != nil {
		log.Fatalf("Failed with error: \n %+v \n", err)
	}
	var person *btuser.User
	err = json.Unmarshal(resp.Json, &person)
	if err != nil {
		log.Fatalf("Failed with error : \n %v \n", err)
	}

	return person, nil
}

//GetUserFromDGraph - Returns User
func GetUserFromDGraph(email string) (*btuser.User, error) {
	dgraph := DgraphCONN()
	variables := map[string]string{"$email": email}

	query := `query all($email: string){
		all(func: q($email)){
			uid
			username
			email
			accountstatus
			verificationcode
			codestatus
			datejoin
			lastupdated
		}
	}`

	//
	var usr *btuser.User
	ctx := context.Background()
	resp, err := dgraph.NewTxn().QueryWithVars(ctx, query, variables)
	if err != nil {
		log.Fatalf("Could not query DGraph. The following error occured: \n %+v \n", err.Error())
		//return nil, err.Error()
	}

	//fmt.Println(string(resp.Json))
	err = json.Unmarshal(resp.Json, &usr)
	if err != nil {
		log.Fatal(err)
	}

	return usr, nil
}

//IsUserSignedUp - Will check if a user with the given email address is already signed up or not -
//Param string email
//return bool true if user is signedup, false if not
func IsUserSignedUp(email string) bool {
	userInfo, err := GetUserFromDGraph(email)
	if err != nil {
		log.Fatalf("Could not get user information from DGraph the following error occured : \n %+v \n", err)
	}

	//check string length of returned useruid
	if len(userInfo.Uuid) > 2 {
		//user is signed up and has a uuid
		return true
	}
	return false
}
