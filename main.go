package main

import (
	"flag"
	"fmt"
	"github.com/garyburd/redigo/redis"
	"log"
	"net/http"
)

var (
	httpAddr  = "0.0.0.0:8080"
	redisAddr = flag.String("redis-address", "localhost:6379", "Adress to the Redis Serv")
	c, err    = redis.Dial("tcp", *redisAddr)
)

func http_handler(w http.ResponseWriter, r *http.Request) {
	c.Send("INCR", "hits")
	v, err := redis.Int(c.Do("GET", "hits"))
	log.Println("count up ", v)
	log.Println("request is ", r)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Fprintf(w, "%d", v)
}

func main() {
	c.Send("SET", "hits", 0)
	log.Println("start")
	http.HandleFunc("/", http_handler)
	http.ListenAndServe(httpAddr, nil)
}
