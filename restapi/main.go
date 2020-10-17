package main

import (
	"bufio"
	"encoding/csv"
	"errors"
	"fmt"
	"io"
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
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
)

const (
	port = ":1980"
	//APP_KEY = os.Getenv("BUYTRANCETOKENKEY")

)

var mySigningKey = os.Getenv("BUYTRANCETOKENKEY")

//HB519BUXTICP0QAZEJLT

/*
//User type
type User struct {
	UserID         string `json:"userid"`
	Email          string `json:"email"`
	Firstname      string `json:"firstname"`
	Lastname       string `json:"lastname"`
	Phone          string `json:"phone"`
	Usertype       string `json:"usertype"`
	Password       string `json:"password"`
	PasswordExpiry string `json:"passwordexpiry"`
	Datejoined     string `json:"datejoined,omitempty"`
	Lastlogin      string `json:"lastlogin,omitempty"`
	Status         string `json:"status"`
}*/

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

		//
		userName := string(d[1])
		failOnError(err, "Failed to convert line entry to string")

		userEmail := string(d[2])
		failOnError(err, "Failed to convert line entry to string")

		userPhone := string(d[3])
		failOnError(err, "Failed to convert line entry to string")

		userDetail := map[string]string{
			"Name":  userName,
			"Email": userEmail,
			"Phone": userPhone,
		}

		//mtx lock for read write concurrency protection
		mtx.Lock()
		Useraccounts[userID] = append(Useraccounts[userID], userDetail)
		mtx.Unlock()
	}
}

//checkPreviousEmailUse checks submitted email for duplicate returns true if duplicate and false if not
//params email string
//return bool
func checkPreviousEmailUse(email string) bool {
	//checks
	for _, v := range Useraccounts {
		if email == v[0]["Email"] {
			//email already in use
			return true
		}
	}
	//email not in user
	return false
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
}*/

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
				fmt.Println("Account exists")
				fmt.Printf("%v %s %s %s \n", userid, userName, userEmail, userPhone)

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

//accept verification code sent to email address
func verifyCode(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Verify page reached")

	info := mux.Vars(r)
	//retreive request parameters
	code := info["code"]
	fmt.Printf("Received : %v \n", code)
}

func homepage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to home page")
}

func profile(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Profile page reached")
}

func main() {
	//seed random
	rand.Seed(time.Now().UnixNano())

	//Preload user data
	go func() {
		preloadUsers()
	}()

	//fmt.Println(Useraccounts)
	//Start Gorrila mux
	router := mux.NewRouter()
	//All crud handlers that will be needed by front ends will be defined and handled here.
	//User activity
	router.HandleFunc("/", homepage).Methods("GET")
	router.HandleFunc("/signup/{email}/{phone}", signup).Methods("GET") //receives user access verification code
	router.HandleFunc("/verifycode/{verificationCode}", verifyCode).Methods("POST")
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
	*/
	//start serving handles
	fmt.Printf("Server running at port %v \n", port)
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
