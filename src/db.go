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
	if err != nil {
		panic(err)
	}
	
}

/*
*	Initial Data
*/
func initData() {
			   
	createFeed(Subscriber{
		username: "Morpheus",
		id: "101",	
		Feed: Feed{
			id: "2",
			publisher: "Fast Company",
			title: "Sucessful Startups without VC Funding",
			url: "http://foo.com",
			topic: "Business",	
		},
	}),
	createFeed(Subscriber{
		username: "Logan",
		id: "102",	
		Feed: Feed{
			id: "2",
			publisher: "Design Milk",
			title: "Augmented Reality Startups",
			url: "http://foo.com",
			topic: "Design",	
		},
	})

}





func findFeed(id int) Feed {
	
	var feed Feed

	//client := redisConnnection()
	defer client.Close()

	//keys, err := client.Do("KEYS", "post:*")
	//HandleError(err)


}





func findAllFeeds() feeds {

	var feeds Feeds

	defer client.Close()

	
	hashMap, err := client.HgetAll("username:*").Result()
	if err != nil {
		return err
	}

	for key, value := range hashMap {
		var feed Feed

		reply, err := client.Hget("",)
		// Handle Error

		if err := json.Unmarshall(reply, &feed); err != nil {
			panic(err)
		}

		feeds = append(feeds, feed)
	}

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