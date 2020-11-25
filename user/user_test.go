package main

import (
	"fmt"
	"log"
	"os"
	"testing"
	"time"

	btuser "github.com/Prosp3r/buytrance/user/pb"
	"github.com/golang/protobuf/ptypes"
)

func TestMain(m *testing.M) {
	rc := m.Run()

	// rc 0 means we've passed,
	// and CoverMode will be non empty if run with -cover
	if rc == 0 && testing.CoverMode() != "" {
		c := testing.Coverage()
		if c < 0.5 {
			fmt.Println("Tests passed but coverage failed at", c)
			rc = -1
		}
	}
	os.Exit(rc)
}

func TestAddUser(m *testing.T) {
	//func (c *user) AddUser(ctx context.Context, r *btuser.User) (*btuser.UserResponse, error) {
	//timenow := time.Now()
	//tstamp := timenow.Second()

	//tx, err := ptypes.Timestamp()
	tn := time.Now()
	ut, err := ptypes.TimestampProto(tn)
	if err != nil {
		log.Fatalf("Failed to conver to proto's timestamp.Timestamp with the following errot: \n %+v \n", err)
	}

	//st := &timestamp.Timestamp{tstamp, }
	p := []*btuser.User{
		{"", "Prosper", "prosper@email.com", "inactive", "1v0", "pending", ut, ut}
	}

	if !equal(got, exp) {
		t.Fatalf("Test AddUser Test expected: %v \n Got: %v \n", exp, got )
	}
	//in := []user{"", "name1", "user1@fakemail.com", "unverified", "0v1", "pending", timenow, timenow}

}

/*
Uuid                 string               `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
	Username             string               `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	Email                string               `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
	Accountstatus        string               `protobuf:"bytes,4,opt,name=accountstatus,proto3" json:"accountstatus,omitempty"`
	Verificationcode     string               `protobuf:"bytes,5,opt,name=verificationcode,proto3" json:"verificationcode,omitempty"`
	Codestatus           string               `protobuf:"bytes,6,opt,name=codestatus,proto3" json:"codestatus,omitempty"`
	Datejoined           *timestamp.Timestamp `protobuf:"bytes,8,opt,name=datejoined,proto3" json:"datejoined,omitempty"`
	Lastupdated          *timestamp.Timestamp `protobuf:"bytes,9,opt,name=lastupdated,proto3" json:"lastupdated,omitempty"`
*/
