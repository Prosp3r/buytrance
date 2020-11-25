/*package main

import (
	"bufio"
	"context"
	"encoding/csv"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"sync"
	"time"

	//"tideland.dev/go/net/jwt/token"
	//"github.com/dgrijalva/jwt-go"
	//"github.com/go-net/jwt/token"

	"github.com/dgraph-io/dgo/protos/api"
	"github.com/dgraph-io/dgo/v200"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
)

const (
	port = ":1980"
	//APP_KEY = os.Getenv("BUYTRANCETOKENKEY")
)

var mySigningKey = os.Getenv("BUYTRANCETOKENKEY")

//HB519BUXTICP0QAZEJLT

//SystemMessage -
type SystemMessage struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}


//User type
type User struct {
	UserID         string `json:"userid"`
	Fullname 	   string `json:"fullname"`
	Email          string `json:"email"`
	Phone          string `json:"phone"`
	Datejoin       string `json:"lastlogin,omitempty"`
	Verifycode 	   string `json:"verifycode"`
	Status         string `json:"status"`
}

//failOnError is  single place to handle errors to reduces number of keystrokes per each error handling call.
func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s %s", msg, err)
	}
}

//Useraccounts - hold user account list in memory
//[uint64 - will hold unique userid] string will hold a json record of user account information
//var Useracounts = make(map[uint64]string)
var Useraccounts = make(map[uint64][]map[string]string)

//FiboStore -- Hold fibonacci data
//var FiboStore = make(map[string]uint64)

//mtxx Mutex lock for protecting read and write operations
var mtx = &sync.Mutex{}

//preloadUsers - will load user list from designated source to memory
func preloadUsers() {
	//userListSource - change to env variable for production
	userListSource := "user.csv"

	//temporarily read from csv source - This can be change to badger or sql db
	dataFile, err := os.Open(userListSource)
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
		userID, err := strconv.ParseUint(string(d[0]), 0, 64)
		failOnError(err, "Failed converting line entry to Uint64")
		fmt.Printf("UserID : %v \n", userID)

		//
		userName := string(d[1])
		failOnError(err, "Failed to convert line entry to string")

		userEmail := string(d[2])
		failOnError(err, "Failed to convert line entry to string")

		userPhone := string(d[3])
		failOnError(err, "Failed to convert line entry to string")

		userConfirmCode := string(d[4])
		failOnError(err, "Failed to convert line entry to string")

		userCCodeStatus := string(d[5])
		failOnError(err, "Failed to convert line entry to string")

		userDetail := map[string]string{
			"Name":        userName,
			"Email":       userEmail,
			"Phone":       userPhone,
			"ConfirmCode": userConfirmCode,
			"CCodeStatus": userCCodeStatus,
		}

		//mtx lock for read write concurrency protection
		mtx.Lock()
		Useraccounts[userID] = append(Useraccounts[userID], userDetail)
		mtx.Unlock()
	}
}

func saveToDGraph(c *dgo.Dgraph) {

	//DGraph Handler
	//:::::::::::::::::
	dgcon := DgraphCONN()
	op := &api.Operation{}
	op.Schema = `
		email: string @index(exact) .
		phone: string .
		joindate: datetime .
		verifycode: string .
		verifystatus: string .
	`
	ctx := context.Background()
	err := dgcon.Alter(ctx, op)
	failOnError(err, "Could not alter dgraph db")

	joindate := time.Date(1980, 01, 01, 23, 0, 0, 0, time.UTC) //if needed

	var user = new(User)
	user.UserID =


	//:::::::::::::::::::
	//END DGRAPH HANDLER
}

//saveUserData - will save user data to designate storage location currently supports .csv stored on file
func saveUserData() {
	fileName := "user.csv"
	//fmt.Printf("Saving user data state to %v \n", fileName)

	for {

		users := Useraccounts
		var recordStrings []string
		var byteStrings []byte
		for i, v := range users {
			//fmt.Printf(string(i))
			userID := i

			recordstring := strconv.FormatUint(userID, 10) + "," + v[0]["Name"] + "," + v[0]["Email"] + "," + v[0]["Phone"] + "," + v[0]["ConfirmCode"] + "," + v[0]["CCodeStatus"] + "\n"
			recordStrings = append(recordStrings, recordstring)
			xbyteStrings := []byte(recordstring)

			byteStrings = append(xbyteStrings, byteStrings...)

		}

		//fmt.Printf("%v \n %v \n\n", recordStrings, byteStrings)
		mtx.Lock()
		openFile, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY, 0644)
		failOnError(err, "Failed to open file for saving")

		err = ioutil.WriteFile(fileName, byteStrings, 0644)
		openFile.Close()
		mtx.Unlock()
		time.Sleep(time.Second)
	}

}

//maxUserID - will return the highest number of user id in the list.
func maxUserID() uint64 {
	var x uint64
	for i := range Useraccounts {
		if i > x {
			x = i
		}
	}
	return x
}

//addNewUser - adds  a new user to the list
func addNewUser(userName, userEmail, userPhone string) bool {
	userConfirmCode := generateCode(6)
	userCCodeStatus := "active"

	var userID uint64
	userDetail := map[string]string{
		"Name":        userName,
		"Email":       userEmail,
		"Phone":       userPhone,
		"ConfirmCode": userConfirmCode,
		"CCodeStatus": userCCodeStatus,
	}

	userID = maxUserID() + 1
	//mtx lock for read write concurrency protection
	mtx.Lock()
	Useraccounts[userID] = append(Useraccounts[userID], userDetail)
	mtx.Unlock()
	return false
}

//checkPreviousEmailUse checks submitted email for duplicate returns true if duplicate and false if not
//params email string
//return bool
func checkPreviousEmailUse(email string) bool {
	//checks
	//fmt.Println(email)
	for _, v := range Useraccounts {
		if email == v[0]["Email"] {
			//email already in use
			return true
		}
	}
	//email not in user
	return false
}

//checkUsedMail - check if email was already used
func checkUsedMail(w http.ResponseWriter, r *http.Request) {
	param := mux.Vars(r)
	email := param["email"]
	check := checkPreviousEmailUse(email)
	//var response []SystemMessage
	response := SystemMessage{email, strconv.FormatBool(check)}

	//user := []string{"email":email, status: }
	//json.NewEncoder()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
	return
	//fmt.Println(w, check)
}

//checkMail - Checks if email is in use and valid
func startLogin(w http.ResponseWriter, r *http.Request) {

	//users := Useraccounts
	//
	params := mux.Vars(r)
	//set sent data
	userEmail := params["email"]
	if checkPreviousEmailUse(userEmail) == true {
		//email exists
		//generate and send confirmation code
		confirmCode := generateCode(6)
		//Save code to email profile
		loginUser := setLoginCode(userEmail, confirmCode)
		if loginUser == true {
			//email the code confirmCode to user with email address userEmail
		}
		//fmt.Printf("Email: %v - Code: %v \n", userEmail, confirmCode)

	} else {
		//email does not exist
		//write to file and send confirmation code
		userPhone := ""
		userName := ""
		addUser := addNewUser(userName, userEmail, userPhone)
		if addUser == true {
			//email the code confirmCode to user with email address userEmail
			//return json response to request
		}
		//fmt.Printf("Email: %v - Code: %v \n", userEmail, confirmCode)
	}
}

//1. fetch and load codes and users
//2. loop through codes and verify
//3. return result if any
//4. return false if all fils
//verifyCode - used to completes an initiated login or purchase process.
func verifyCode(w http.ResponseWriter, r *http.Request) {
	param := mux.Vars(r)
	code := param["confirmcode"]
	email := param["email"]
	var response SystemMessage
	for _, v := range Useraccounts {
		if v[0]["Email"] == email && v[0]["ConfirmCode"] == code {
			if v[0]["CCodeStatus"] == "active" {
				//authorised
				response = SystemMessage{email, strconv.FormatBool(true)}
				//fmt.Println(w, "Authorised")
			} else {
				//not authosrised
				response = SystemMessage{email, strconv.FormatBool(false)}
				//fmt.Println(w, "Not Authorised") - 162589
			}
		}
		//not authorized
		//fmt.Println(w, "Not Authorised")
	}

	response = SystemMessage{email, strconv.FormatBool(false)}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
	return
	//fmt.Println(w, check)

}

//generateCode - will generate an alpha numeric string of length len.
//param len (int) length of string to be returned
// kind (string) determine the combination of string types e.g numeric =1-0, alphanumeric=1-0+a-z+A-Z, alpha=a-z+A-Z, complex=a-z+A-Z+1-0+specialchracters
//func generateCode(size int, kind string) string {
func generateCode(size int) string {

	//var lettersRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	var numbersRunes = []rune("1234567890")
	//var xactersRunes = []rune("~!@#$%^&*()+_|}{][")
	r := make([]rune, size)
	for i := range r {
		r[i] = numbersRunes[rand.Intn(len(numbersRunes))]
	}
	return string(r)
}

//setLoginCode - Sets an active login code for the user
func setLoginCode(userEmail, confirmCode string) bool {

	for _, v := range Useraccounts {
		if v[0]["Email"] == userEmail {
			v[0]["ConfirmCode"] = confirmCode
			v[0]["CCodeStatus"] = "active"
			return true
		}
	}
	return false
}

/*
func auth(w http.ResponseWriter, r *http.Request) {
	var user User
	_ = json.NewDecoder(r.Body).Decode(&user)

	token := jwt.NewWithClaims(jwt.SigningMethodES256, jwt.MapClaims{
		"userid":    user.UserId,
		"email":     user.Email,
		"firstname": user.Firstname,
		"lastname":  user.Lastname,
	})
	tokenString, err := token.SignedString([]byte(APP_KEY)) //APP_KEY=os.Getenv("")
	failOnError(err, "Unauthorized")

	json.NewEncoder(w).Encode(JwtToken{Token: tokenString})
}* /

func isAuthorized(endpoint func(http.ResponseWriter, *http.Request)) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		if r.Header["Token"] != nil {
			token, err := jwt.Parse(r.Header["Token"][0], func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("There was an error")
				}
				return mySigningKey, nil
			})

			if err != nil {
				fmt.Fprintf(w, err.Error())
			}

			if token.Valid {
				endpoint(w, r)
			}

		} else {
			fmt.Fprintf(w, "Not Authorized")
		}
	})
}

var (
	//ErrBadFormat -
	ErrBadFormat = errors.New("invalid format")
	//ErrUnresolvableHost -
	//ErrUnresolvableHost = errors.New("unresolvable host")
	emailRegexp = regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
)

//ValidateFormat - validates email format
func ValidateFormat(email string) bool {
	re := emailRegexp
	return re.MatchString(email)
}

//validateEmail - Validate email systax
//param email string
//returns bool
func validateEmail(email string) bool {
	//check for email format
	valid := ValidateFormat(email)
	return valid
}

//signup adds new user to system
func signup(w http.ResponseWriter, r *http.Request) {
	//fmt.

	params := mux.Vars(r)

	//set sent data
	userEmail := params["email"]
	userPhone := params["phone"]
	fmt.Fprintf(w, "Signup page reached \n Email : %v \n Phone: %v \n", userEmail, userPhone)
	//validate sent parameters

	//- is email previously used ?

	//- is email a valid format
	isValid := validateEmail(userEmail)
	if isValid == true {
		fmt.Println("Email is valid")
		//email is properly formed
		for i, d := range Useraccounts {
			if d[0]["Email"] == userEmail || d[0]["Phone"] == userPhone {

				//user already exist - generate login code and send to email
				userid := i
				userName := d[0]["Name"]
				userEmail := d[0]["Email"]
				userPhone := d[0]["Phone"]
				userCreatedDate := d[0]["CreatedDate"]
				confirmCode := d[0]["ConfirmCode"]

				fmt.Println("Account exists")
				fmt.Printf("%v %s %s %s %s %s \n", userid, userName, userEmail, userPhone, userCreatedDate, confirmCode)

				//[x]->				//send login code via message queue				<-[x]
			}
			//user account does not exist - write to list before generating login code and sending to email
			//[x]->				//create user account via message queue				<-[x]
			//[x]->				//send login code via message queue					<-[x]
		}
	} else {
		fmt.Println(w, "Your email is malformed")
		//fmt.Fprintf(w, "\n You sent user email: %v \n Phone: %v \n", userEmail, userPhone)
		//return an error message
	}
	//- save new user
	//- head to confirmation endpoint page
}

/* /accept verification code sent to email address
func verifyCode(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Verify page reached")

	info := mux.Vars(r)
	//retreive request parameters
	code := info["code"]
	fmt.Printf("Received : %v \n", code)
}* /

func homepage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to home page")
}

func profile(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Profile page reached")
}

/*
:::::::::::
:::::::::::
:::::::::::
:::::::::::
* /
func main() {
	//seed random
	rand.Seed(time.Now().UnixNano())

	//Preload user data
	go func() {
		preloadUsers()
		saveUserData()
	}()

	//fmt.Println(Useraccounts)
	//Start Gorrila mux
	router := mux.NewRouter()
	//All crud handlers that will be needed by front ends will be defined and handled here.
	//User activity
	router.HandleFunc("/", homepage).Methods("GET")
	router.HandleFunc("/login/{email}", startLogin).Methods("GET") //receives user email and send out verification code

	router.HandleFunc("/checkmail/{email}", checkUsedMail).Methods("GET") //receives user email and send out verification code
	router.HandleFunc("/signup/{email}/{phone}", signup).Methods("GET")   //receives user access verification code
	router.HandleFunc("/verifycode/{verificationCode}/{email}", verifyCode).Methods("POST")
	//router.HandleFunc("/login", auth).Methods("POST")

	router.Handle("/profile/{token}/{profileid}", isAuthorized(profile))
	//Domain name activity
	/*router.HandleFunc("/searchDomain", searchDomain).Methods("GET")
	router.HandleFunc("/registerDomain", registerDomain).Methods("POST")
	router.HandleFunc("/renewDomain", renewDomain).Methods("POST")
	//Handling shop operations
	router.HandleFunc("/shop", getAllShops).Methods("GET")
	router.HandleFunc("/shop", createAShop).Methods("POST")
	router.HandleFunc("/shop/{shopId}", getAShop).Methods("GET")
	router.HandleFunc("/shop/{shopId}", editAShop).Methods("PUT")
	router.HandleFunc("/shop/{shopId}", deactivateShop).Methods("DELETE")
	//User Subscription starting/renewal upgrade and downgrades
	router.HandleFunc("/subscription", getAllSubscriptions).Methods("GET")
	router.HandleFunc("/subscription/{userId}", getUserSubscriptions).Methods("GET")
	router.HandleFunc("/subscripton", createSubscription).Methods("POST")
	router.HandleFunc("/subscription/{subscriptionId}", updateSubscription).Methods("PUT") //upgrades and downgrades
	router.HandleFunc("/subscription/{subscriptionId}", deleteSubscription).Methods("DELETE")
	//User transactions
	router.HandleFunc("/transaction", getAllTransactions).Methods("GET")
	router.HandleFunc("transactions/{userId}", getAllUserTransactions).Methods("GET")
	* /
	srcx := []byte{'u', 'k', 'w', 'a', 'm', 'e'}
	fmt.Println(srcx)
	//start serving handles
	fmt.Printf("Buytrance-Rest API running at port %v \n", port)
	log.Fatal(http.ListenAndServe(port, router))
}

/*
//MiddleWare
func basicAuthMiddleWare(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("WWW-Authenticate", `Basic realm="Restricted"`)

		s := strings.SplitN(r.Header.Get("Authorization"), " ", 2)
		if len(s) != 2 {
			http.Error(w, "Not authorized", 401)
			return
		}
​
		b, err := base64.StdEncoding.DecodeString(s[1])
		if err != nil {
			http.Error(w, err.Error(), 401)
			return
		}
​
		pair := strings.SplitN(string(b), ":", 2)
		if len(pair) != 2 {
			http.Error(w, "Not authorized", 401)
			return
		}
​
		if pair[0] != "username" || pair[1] != "password" {
			http.Error(w, "Not authorized", 401)
			return
		}​
		h.ServeHTTP(w, r)
	}
}*/
