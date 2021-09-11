package main

import (
	MQTT "github.com/eclipse/paho.mqtt.golang"
	mqtt "github.com/tech-sumit/aws-iot-device-sdk-go"
)

func main(){
	connection,err:=mqtt.NewConnection(mqtt.Config{
		KeyPath:  "<KEY_PATH>",
		CertPath: "<CERT_PATH>",
		CAPath:   "<CA_PATH>",
		ClientId: "ping_client",
		Endpoint: "<MQTT_GATEWAY_ENDPOINT>",
	})
	if err!=nil {
		panic(err)
	}
	go func() {
		err=connection.SubscribeWithHandler("ping",0, func(client MQTT.Client, message MQTT.Message) {
			print(string(message.Payload()))
		})
	}()
	if err!=nil {
		panic(err)
	}
	err=connection.Publish("ping","pong",0)
	if err!=nil {
		panic(err)
	}
}
