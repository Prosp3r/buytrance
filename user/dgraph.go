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

