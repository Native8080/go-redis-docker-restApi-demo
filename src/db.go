package main

import (
	"fmt"
	"encoding/json"
	"strconv"
	"github.com/go-redis/redis"
)

/*
*	Redis Connection
*/

var client *redis.Client

func redisConnnection() {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", 
		DB:       0, 
	})

	pong, err := client.Ping().Result()
	fmt.Println(pong, err)
	if err != nil {
		panic(err)
	}
	
}

/*
*	Initial Data
*/
func initData() {
			   
	createFeed(Feed{
		Subscriber: Subscriber{
			username: "Logan",
			id: "101"
		},
		id: "1"
		publisher: "Design Milk",
		title: "Augmented Reality Startups",
		url: "http://foo.com",
		topic: "Design"
	})

	createFeed(Feed{
		Subscriber: Subscriber{
			username: "Morpheus",
			id: "102"
		},
		id: "2"
		publisher: "Fast Company",
		title: "Sucessful Startups without VC Funding",
		url: "http://foo.com",
		topic: "Business"
	})
}


func findFeed(id int) Feed {
	var feed Feed

	c := redisConnnection()
	defer c.Close()
}

func findAllFeeds() feeds {

	var feeds Feed

	c := redisConnnection()
	defer c.Close()

	

	return feeds
}

func createFeed(f Feed) {
	
	c := RedisConnect()
	defer c.Close()
	
	
	// Save JSON blob to Redis
	reply, err := c.Do("HMSET", )
	HandleError(err)
	
	fmt.Println("HMGETALL ", reply)


}