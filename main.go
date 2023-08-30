package main

import (
	"fmt"
	"log"

	"github.com/gocql/gocql"
)

func main() {
	// Connect to Cassandra
	cluster := gocql.NewCluster("127.0.0.1")
	cluster.Keyspace = "demo"
	cluster.Consistency = gocql.Quorum
	session, err := cluster.CreateSession()
	if err != nil {
		log.Fatalf("Failed to connect to Cassandra: %v", err)
	}
	defer session.Close()

	// Insert a user
	userID := gocql.TimeUUID()
	if err := session.Query(`INSERT INTO users (id, name, age) VALUES (?, ?, ?)`, userID, "Alice", 30).Exec(); err != nil {
		log.Fatalf("Failed to insert user: %v", err)
	}

	// Query the user
	var name string
	var age int
	if err := session.Query(`SELECT name, age FROM users WHERE id = ? LIMIT 1`, userID).Scan(&name, &age); err != nil {
		log.Fatalf("Failed to query user: %v", err)
	}
	fmt.Printf("Retrieved user %s, age %d\n", name, age)

	// Query all users
	//iter := session.Query(`SELECT name, age FROM users`).Iter()

	// var name string
	// var age int
	// for iter.Scan(&name, &age) {
	// 	fmt.Printf("Retrieved user %s, age %d\n", name, age)
	// }

	// if err := iter.Close(); err != nil {
	// 	log.Fatalf("Error during query iteration: %v", err)
	// }
}
