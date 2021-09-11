// Package mqtt provides paho mqtt implementation to connect with AWS IoT Gateway
package mqtt

import (
	"crypto/tls"
	"crypto/x509"
	"errors"
	MQTT "github.com/eclipse/paho.mqtt.golang"
	"io/ioutil"
	"time"
)

// AWSIoTConnection holds connection for AWS IoT
type AWSIoTConnection struct {
	client  MQTT.Client
	options *MQTT.ClientOptions
}

// Config struct holds connection parameters for AWSIoTConnection initialisation
// All parameters are compulsory
// KeyPath is path to x.509 Private Key
// CertPath is path to x.509 Public Key
// CAPath is path to CA certificate
// ClientId is the clientId of thing
// Endpoint is Unique AWS IoT endpoint provided by AWS IoT
type Config struct {
	KeyPath  string `json:"keyPath" binding:"required"`
	CertPath string `json:"certPath" binding:"required"`
	CAPath   string `json:"caPath" binding:"required"`
	ClientId string `json:"clientId" binding:"required"`
	Endpoint string `json:"endpoint" binding:"required"`
}

// init It initialises MQTT connection. Its called by NewConnection internally
func (c *AWSIoTConnection) init(config Config) error {
	tlsCert, err := tls.LoadX509KeyPair(
		config.CertPath, config.KeyPath)
	if err != nil {
		return err
	}

	certs := x509.NewCertPool()
	caPem, err := ioutil.ReadFile(config.CAPath)
	if err != nil {
		return err
	}
	certs.AppendCertsFromPEM(caPem)
	tlsConfig := &tls.Config{
		Certificates: []tls.Certificate{tlsCert},
		RootCAs:      certs,
	}

	c.options = MQTT.NewClientOptions()
	c.options.AddBroker("tcps://" + config.Endpoint + ":8883/mqtt")
	c.options.SetMaxReconnectInterval(10 * time.Second)
	c.options.SetClientID(config.ClientId)
	c.options.SetTLSConfig(tlsConfig)
	return nil
}

// NewConnection Creates new MQTT connection with AWS IoT
// It requires config parameter for initialisation
func NewConnection(config Config) (*AWSIoTConnection,error) {
	connection:=AWSIoTConnection{}
	if err := connection.init(config); err != nil {
		return nil,err
	}
	connection.client = MQTT.NewClient(connection.options)
	token := connection.client.Connect()
	token.Wait()
	if err := token.Error(); err != nil {
		return nil,err
	} else {
		return &connection,nil
	}
}

// Disconnect function disconnects the MQTT connection only if its already connected else returns error
func (c *AWSIoTConnection) Disconnect() bool {
	if c.client.IsConnected() {
		c.client.Disconnect(0)
		return true
	} else {
		return false
	}
}

// Subscribe function subscribes on topic with level of qos (Quality of service)
// currently supported 0 & 1 (2 coming in future).
func (c *AWSIoTConnection) Subscribe(topic string, qos byte) error {
	return c.SubscribeWithHandler(topic, qos, nil)
}

// SubscribeWithHandler function subscribes on topic with level of qos (Quality of service)
// currently supported 0 & 1 (2 coming in future)
// & handler function to listen to incoming messages for the topic & qos level.
// It is called every time when message is received
func (c *AWSIoTConnection) SubscribeWithHandler(topic string, qos byte, handler MQTT.MessageHandler) error {
	if !c.client.IsConnected() {
		return errors.New("client not connected")
	} else {
		token := c.client.Subscribe(topic, qos, handler)
		token.Wait()
		if err := token.Error(); err != nil {
			return err
		} else {
			return nil
		}
	}
}

// Unsubscribe function removes subscription for specified topic
func (c *AWSIoTConnection) Unsubscribe(topic string) error {
	if !c.client.IsConnected() {
		return errors.New("client not connected")
	} else {
		token := c.client.Unsubscribe(topic)
		token.Wait()
		if err := token.Error(); err != nil {
			return err
		} else {
			return nil
		}
	}
}

// Publish function publishes data in interface on topic with level of qos (Quality of service)
// currently supported 0 & 1 (2 coming in future)
func (c *AWSIoTConnection) Publish(topic string, data interface{}, qos byte) error {
	token := c.client.Publish(topic, qos, false, data)
	token.Wait()
	if err := token.Error(); err != nil {
		return err
	} else {
		return nil
	}
}
