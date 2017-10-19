package main

import (
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/gocql/gocql"
)

func main() {
	cluster := gocql.NewCluster("cassandra", "cassandra_cassandra_node_1")
	cluster.Keyspace = "test"
	cluster.Consistency = gocql.Quorum
	session, err := cluster.CreateSession()
	if err != nil {
		panic(err)
	}
	defer session.Close()

	start := time.Now()
	log.Println("start")
	wg := new(sync.WaitGroup)
	semaphore := make(chan int, 50)
	for i := 0; i < 100; i++ {
		wg.Add(1)
		semaphore <- 1
		go func(i int) {
			defer wg.Done()
			for j := 0; j < 100; j++ {
				if err := session.Query("insert into test_table (id, test_value) values (?, ?)", fmt.Sprintf("%d-%d", i, j), fmt.Sprintf("test%d-%d", i, j)).Exec(); err != nil {
					log.Fatalf(err.Error())
				}
			}
			<-semaphore
		}(i)
	}
	wg.Wait()
	log.Println("end", time.Since(start))
}
