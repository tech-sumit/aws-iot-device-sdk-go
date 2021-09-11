// Package lifecycle provides topics to intercept thing lifecycle & subscription events
package lifecycle

const (
	// Presence events: Connect / Disconnect
	// connected: AWS IoT presence connected topic
	connected       string = "$aws/events/lifecycle/connected/scv/"
	// AnyConnected : When subscribed to this topic gives events for all devices connecting to the iot gateway
	AnyConnected    string = connected + "+"
	// disconnected: AWS IoT presence disconnected topic
	disconnected    string = "$aws/events/lifecycle/disconnected/scv/"
	// AnyDisconnected : When subscribed to this topic gives events for all devices disconnected from the iot gateway
	AnyDisconnected string = disconnected + "+"

	// Subscribed / Unsubscribed
	// subscribed: AWS IoT client subscription topic
	subscribed      string = "$aws/events/subscriptions/subscribed/"
	// AnySubscribed : When subscribed to this topic gives events for all subscriptions being performed on iot gateway
	AnySubscribed   string = subscribed + "+"
	// unsubscribed: AWS IoT client unsubscription topic
	unsubscribed    string = "$aws/events/subscriptions/unsubscribed/"
	// AnyUnsubscribed : When subscribed to this topic gives events for all unsubscribe operations being performed on iot gateway
	AnyUnsubscribed string = unsubscribed + "+"
)

// ClientConnected builds topic which will listen to connection event for clientId
func ClientConnected(clientId string) string {
	return connected + clientId
}

// ClientDisconnected builds topic which will listen to disconnection event for clientId
func ClientDisconnected(clientId string) string {
	return disconnected + clientId
}

// OnSubscribed builds topic which will listen to all subscription events for specified topic
func OnSubscribed(topic string) string {
	return subscribed + topic
}

// OnUnsubscribed builds topic which will listen to all unsubscribe events for specified topic
func OnUnsubscribed(topic string) string {
	return unsubscribed + topic
}
