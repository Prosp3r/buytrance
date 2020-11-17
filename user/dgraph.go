/*
DGraph wrapper that can be dropped into any of the distributed packages to manage connection and queries to the DGraph database.
*/

package main

import (
	"log"

	"github.com/dgraph-io/dgo/protos/api"
	"github.com/dgraph-io/dgo/v200"
	"google.golang.org/grpc"
)

var dgServer = "localhost:9080"

//DgraphCONN - Opens connection to Dgraph database.
func DgraphCONN() *dgo.Dgraph {
	conn, err := grpc.Dial(dgServer, grpc.WithInsecure())
	if err != nil {
		log.Fatal("Could not connect to DGraph gRPC")
	}
	defer conn.Close()
	dc := api.NewDgraphClient(conn)
	dg := dgo.NewDgraphClient(dc)

	return dg
}
