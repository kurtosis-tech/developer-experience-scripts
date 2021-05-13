package main

import (
	"fmt"

	"github.com/gocql/gocql"
)

func main() {

	// Define object used to represent local Cassandra cluster
	cluster := gocql.NewCluster("127.0.0.1")
	cluster.Consistency = gocql.One
	cluster.ProtoVersion = 4

	// Define object used to send queries to local Cassandra cluster
	session, err := cluster.CreateSession()
	if err != nil {
		panic(err)
	}
	defer session.Close()

	// Create a keyspace "test" to use for testing purposes
	err = session.Query(`CREATE KEYSPACE test WITH REPLICATION = { 'class' : 'SimpleStrategy', 'replication_factor' : 3 }`).Exec()
	if err != nil {
		panic(err)
	}

	// Create a table "tweet" to use for testing purposes
	err = session.Query(`CREATE TABLE test.tweet(timeline text, id timeuuid PRIMARY KEY, text text)`).Exec()
	if err != nil {
		panic(err)
	}

	// Insert a tweet into "tweet" table
	err = session.Query(`INSERT INTO test.tweet (timeline, id, text) VALUES (?, ?, ?)`, "me", gocql.TimeUUID(), "hello world").Exec()
	if err != nil {
		panic(err)
	}

	// Read a tweet from "tweet" table
	var (
		id   gocql.UUID
		text string
	)
	err = session.Query(`SELECT id, text FROM test.tweet WHERE timeline = ? LIMIT 1 ALLOW FILTERING`, "me").Scan(&id, &text)
	if err != nil {
		panic(err)
	}

	// Print the contents of the tweet
	fmt.Printf("Tweet: %v, %v\n", id, text)
	// Print results of test condition
	fmt.Printf("Test passed?: %v\n", text == "hello world")
}
