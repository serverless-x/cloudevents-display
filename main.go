package main

import (
	"context"
	"fmt"
	cloudevents "github.com/cloudevents/sdk-go/v2"
	"log"
)

func display(event cloudevents.Event) {
	fmt.Printf("☁️  cloudevents.Event\n%s", event.String())
}

func main() {
	c, err := cloudevents.NewDefaultClient()
	if err != nil {
		log.Fatal("Failed to create client, ", err)
	}

	log.Fatal(c.StartReceiver(context.Background(), display))
}
