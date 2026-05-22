package api

import (
	"fmt"
	"log"

	"github.com/3dp-cloud/companion/utils"
	mqtt "github.com/eclipse/paho.mqtt.golang"
)

type Client struct {
	mqtt.Client
}

func Connect(cfg utils.Config) Client {
	opts := mqtt.NewClientOptions()
	opts.AddBroker(fmt.Sprintf("wss://%s/mqtt", cfg.IotEndpoint))
	opts.SetClientID("companion")
	opts.SetUsername(cfg.AccessToken)
	opts.SetCleanSession(true)

	client := mqtt.NewClient(opts)
	token := client.Connect()
	token.Wait()
	if token.Error() != nil {
		log.Fatalf("Error connecting to AWS IoT: %v", token.Error())
	}
	fmt.Println("Connected to AWS IoT")
	return Client{client}
}
