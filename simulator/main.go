package main

import (
	"log"
)

func init() {
	err = godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	producer = kafka.NewKafkaProducer()
	kafka.Publish(msgOla, topicreadteste, producer)

	route = route2.Route{
		ID       '1',
		ClientID '1',
	}

	route.LoadPositions()
	stringsjon, _ = route.ExportJsonPositions()
	fmt.Println(stringsjon[0])
}
