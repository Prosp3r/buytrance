package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"

	btuser "github.com/Prosp3r/buytrance/user/pb"
	"github.com/gorilla/mux"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

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
	grpcport := ":8080"
	httpport := ":8081"

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
		return response, nil
	}

	//store in database if not
	//If already exists, return user information and send alert code to email or phone number on file for login
	//response.User = adduser //temporary assignment this will change

	//return response, nil
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
