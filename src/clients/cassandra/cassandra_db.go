package cassandra

import "github.com/gocql/gocql"

var (
	Session *gocql.Session
)

func init() {
	cluster := gocql.NewCluster("127.0.0.1")
	cluster.Keyspace = "oauth"
	cluster.Consistency = gocql.Quorum
	var err error
	Session, err = cluster.CreateSession()
	if err != nil {
		panic(err)
	}
}

func GetSession() *gocql.Session {
	return Session
}
