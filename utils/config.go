package utils

import (
	"log"
	"os"
)

var DefaultIotEndpoint = "a1b2c3d4e5f6g7.iot.us-west-2.amazonaws.com"
var DefaultRestEndpoint = "https://api.3dp.cloud/v1"

type Config struct {
	AccessToken  string
	IotEndpoint  string
	RestEndpoint string
}

func LoadConfig() Config {
	res := Config{}

	res.AccessToken = os.Getenv("3DP_CLOUD_COMPANION_TOKEN")
	if res.AccessToken == "" {
		log.Fatal("3DP_CLOUD_COMPANION_TOKEN environment variable not set")
	}

	res.IotEndpoint = os.Getenv("3DP_CLOUD_COMPANION_IOT_ENDPOINT")
	if res.IotEndpoint == "" {
		res.IotEndpoint = DefaultIotEndpoint
	}

	res.RestEndpoint = os.Getenv("3DP_CLOUD_COMPANION_REST_ENDPOINT")
	if res.RestEndpoint == "" {
		res.RestEndpoint = DefaultRestEndpoint
	}

	return res
}
