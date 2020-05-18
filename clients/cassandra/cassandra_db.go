package cassandra

import (
	"fmt"
	"github.com/gocql/gocql"
)

var (
	session *gocql.Session
)

func init() {
	// connect to Cassandra cluster
	cluster := gocql.NewCluster("127.0.0.1")
	cluster.Keyspace = "oauth"
	cluster.Consistency = gocql.Quorum

	var err error
	if session, err = cluster.CreateSession(); err != nil {
		panic("couldn't connect to Cassandra DB")
	}

	fmt.Println("connected to Cassandra DB")
	defer session.Close()
}

func GetSession() *gocql.Session {
	return session
}
