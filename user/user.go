package main

import (
	"context"
	"fmt"
	"log"
	"net"

	bt "github.com/Prosp3r/buytrance/user/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

//User - holds data structure for user type
type user struct {
}

//
func main() {
	port := ":8080"

	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Failed to listen on port %v with error: %v \n", port, err)
	}
	fmt.Printf("OriCalc Server started on port %v \n", port)

	s := grpc.NewServer()
	bt.RegisterUSERServiceServer(s, &user{})

	//ori.RegisterORIServiceServer(s, &server{})
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed serving ORI with error : %v \n", err)
	}
}

//AddUser - Method will add a new user to the system
func (c *user) AddUser(ctx context.Context, r *bt.User) (*bt.UserResponse, error) {

	var response *bt.UserResponse
	return response, nil
}

//GetUser - will return user's detailed information
func (c *user) GetUser(ctx context.Context, r *bt.IDString) (*bt.UserResponse, error) {
	var response *bt.UserResponse
	return response, nil
}

//GetUsers -  will return a list of users
func (c *user) GetUsers(ctx context.Context, r *bt.Emptyentry) (*bt.Users, error) {
	var response *bt.Users
	return response, nil
}

//UpdateUser - will update a user's information
func (c *user) UpdateUser(ctx context.Context, r *bt.User) (*bt.UserResponse, error) {
	var response *bt.UserResponse
	return response, nil
}

//DeactivateUser - Will flag a user's profile/account deactivated
func (c *user) DeactivateUser(ctx context.Context, r *bt.User) (*bt.BOOLValue, error) {
	var response *bt.BOOLValue
	return response, nil
}
