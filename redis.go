package main

import (
	"github.com/gomodule/redigo/redis"
	"log"
	"strconv"
)

type hacker struct {
	Name string `json:"name"`
	Score int `json:"score"`
}

type hackers []hacker

func (h *hackers)getHacker() {
	conn, err := redis.Dial("tcp", "localhost:6379")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	reaData, err := redis.Strings(conn.Do("ZRANGE", "hackers", 0, -1, "withscores") )
	if err != nil {
		log.Fatal(err)
	}

	count := len(reaData)-1

	for i := 0; i < count; i+=2 {
		name := reaData[i]
		score,_ := strconv.Atoi( reaData[i+1] )
		*h = append(*h,hacker{
			Score: score,
			Name: name,
		} )
	}
}

