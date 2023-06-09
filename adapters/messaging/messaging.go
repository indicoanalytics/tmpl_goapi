package messaging

import (
	"api.default.indicoinnovation.pt/clients/google/pubsub"
)

func Publish(queueName string, message interface{}) {
	pubsub.Publish(queueName, message)
}
