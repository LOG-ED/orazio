package muse

import (
	"log"

	"github.com/garyburd/redigo/redis"
)

func GetInspiratio() []string {
	// Connect to localhost, make sure to have redis-server running on the default port
	conn, err := redis.Dial("tcp", "redis:6379")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	// Ask redigo for []string result for key muse
	muse, err := redis.Strings(conn.Do("MGET", "muse"))
	if err != nil {
		log.Fatal(err)
	}

	if muse == nil {
		// Return default value
		return []string{"http://calliope:9090"}
	}
	return muse
}
