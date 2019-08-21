package main

import (
	mqtt "github.com/eclipse/paho.mqtt.golang"
)

func setupMQTT(c Config) {
	opts := mqtt.NewClientOptions().AddBroker("tcp://" + c.MQTTBrokerHost + ":" + c.MQTTBrokerPort)
	opts.SetClientID(c.ClientID)
}
