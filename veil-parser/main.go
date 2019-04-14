package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"

	"github.com/go-redis/redis"
	"github.com/sergivb01/veil-parser/config"
	"github.com/sergivb01/veil-parser/message"
)

var fileConfig = "config.yaml"

func main() {
	if err := config.LoadFromFile(fileConfig); err != nil {
		log.Fatalf("error while loading config %s: %v\n", fileConfig, err)
	}
	signals := make(chan os.Signal)
	signal.Notify(signals, os.Interrupt)

	client := redis.NewClient(&redis.Options{
		Addr:     config.Config.Redis.Host,
		Password: config.Config.Redis.Password,
	})

	if _, err := client.Ping().Result(); err != nil {
		fmt.Printf("could not connect to redis: %v\n", err)
	}

	sub := client.Subscribe("mc")
	c := sub.Channel()
	listenChannels(signals, c)
}

func listenChannels(signals <-chan os.Signal, mc <-chan *redis.Message) {
	for {
		select {
		case <-signals:
			fmt.Fprint(os.Stderr, "quitting")
			return

		case redisMessage := <-mc:
			m := redisMessage.Payload

			if message.ChatRegex.MatchString(m) {
				fmt.Printf("IS MESSAGE!!! %s\n", m)
			} else if message.KillRegex.MatchString(m) {
				fmt.Printf("IS KILL!!! %s\n", m)
			} else {
				fmt.Println(m)
			}

		}
	}
}
