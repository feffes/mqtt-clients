package main

import (
	"fmt"
	"os"
	//mqtt "github.com/eclipse/paho.mqtt.golang"
)

var conf Config

// Config contains user-defined configs
type Config struct {
	MQTTBrokerHost string
	MQTTBrokerPort string
	ClientID       string
	Topic          string
	Message        string
	gpioPin        string
}

func main() {
	ch := make(chan os.Signal, 1)
	conf = loadEnv()
	fmt.Println("Loaded env:")
	fmt.Printf("%v", conf)
	fmt.Println("HERE")
	sensor := pinConnect(conf.gpioPin)
	tr := make(chan string)
	go poll(sensor, tr)
	fmt.Println(<-tr)
	setupMQTT(conf)
	<-ch
}

func loadEnv() Config {
	brokerHost, exists := os.LookupEnv("MQTT_BROKER_HOST")
	if !exists {
		brokerHost = "localhost"
	}
	brokerPort, exists := os.LookupEnv("MQTT_BROKER_PORT")
	if !exists {
		brokerPort = "1883"
	}
	clid, exists := os.LookupEnv("CLIENT_ID")
	if !exists {
		clid = "motion-tracker"
	}
	topic, exists := os.LookupEnv("MQTT_TOPIC")
	if !exists {
		topic = "test"
	}
	triggerMessage, exists := os.LookupEnv("TRIGGER_MESSAGE")
	if !exists {
		triggerMessage = "TRIGGERED >:("
	}
	gpioPin, exists := os.LookupEnv("GPIO_PIN")
	if !exists {
		gpioPin = "14"
	}

	return Config{brokerHost, brokerPort, clid, topic, triggerMessage, gpioPin}

}
