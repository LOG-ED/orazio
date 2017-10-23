package muse

import (
	"log"

	"github.com/garyburd/redigo/redis"
)

func GetInspiratio() string {
	// Connect to localhost, make sure to have redis-server running on the default port
	conn, err := redis.Dial("tcp", ":6379")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	// Ask redigo for []string result for key muse
	strs, err := redis.Strings(conn.Do("MGET", "muse"))
	if err != nil {
		log.Fatal(err)
	}

	if strs == nil {
		// Return default value
		return "127.0.0.1:9090"
	}
}
