package main

import (
	"context"
	"encoding/json"
	"log"

	//"time"

	btuser "github.com/Prosp3r/buytrance/user/pb"
	"github.com/dgraph-io/dgo/v200/protos/api"
)

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
func AddUserToDGraph(userInfoJSON *btuser.User) (*api.Response, error) {
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
		log.Fatal(err)
	}
	return assigned, nil
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
