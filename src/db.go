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
	})

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


func findFeed(s string, id int) Feed {
	
	var subscriber Subscriber

	key := fmt.Sprintf("%d:%d", subcriber, id)

	defer client.Close()

	val, err := client.HGetAll(key).Result()
	//HandleError(err)

	json.Unmarshall([]byte(val), &subscriber)

	return subscriber


}

func findAllFeeds() feeds {

	var feeds Feeds

	defer client.Close()

	
	keys, err := client.Keys("username:*").Result()
	if err != nil {
		return err
	}

	for _, key := range keys {
		var feed Feed

		val, err := client.Get(key).Result()
		if err != nil {
			return err
		}

		if err := json.Unmarshall([]byte(val), &feed); err != nil {
			panic(err)
		}

		feeds = append(feeds, feed)
	}

	return feeds
}

func createFeed(subcriber Subscriber) {

	//c := RedisConnect()
	defer client.Close()	
	
	bytes, error := json.Marshall(subcriber)
	if err != nil {
		return err
	}

	// redis key format
	// subscriberUsername:subscriberId

	key := fmt.Sprintf("%s:%s", s.Username, s.Id)
	val := string(bytes)
	err := client.Set(key, val, 0).Err()
	if err != nil {
		return err
	}

	return nil

}

func removeFeed(s Subscriber) Subscribers {

}

func updateFeed(s Subscriber) Subscribers {

}