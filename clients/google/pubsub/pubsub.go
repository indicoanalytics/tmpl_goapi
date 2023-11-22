package pubsub

import (
	"context"

	"api.default.indicoinnovation.pt/adapters/logging"
	"api.default.indicoinnovation.pt/app/appinstance"
	"api.default.indicoinnovation.pt/entity"
	"api.default.indicoinnovation.pt/pkg/helpers"
	"cloud.google.com/go/pubsub"
)

func New() (context.Context, *pubsub.Client) {
	ctx := context.Background()

	client, err := pubsub.NewClient(ctx, appinstance.Data.Config.GcpProjectID)
	if err != nil {
		logging.Log(&entity.LogDetails{
			Message: "error to create new pubsub client",
			Reason:  err.Error(),
		}, "critical", nil)

		panic(err)
	}

	return ctx, client
}

func Publish(topicID string, message interface{}) {
	byteMessage, err := helpers.Marshal(message)
	if err != nil {
		logging.Log(&entity.LogDetails{
			Message: "error to marshal message before publishing it to messaging service",
			Reason:  err.Error(),
			Request: map[string]interface{}{
				"queue_name": topicID,
				"message":    message,
			},
		}, "critical", nil)

		return
	}

	ctx, client := New()
	defer client.Close()

	topic := client.Topic(topicID)

	result := topic.Publish(ctx, &pubsub.Message{
		Data: byteMessage,
	})

	serverID, err := result.Get(ctx)
	if err != nil {
		logging.Log(&entity.LogDetails{
			Message: "error to publish message to google pubsub",
			Reason:  err.Error(),
			Request: map[string]interface{}{
				"message":  message,
				"topic_id": topicID,
			},
			Response: map[string]interface{}{"server_id": serverID},
		}, "critical", nil)

		panic(err)
	}

	logging.Log(&entity.LogDetails{
		Message: "message successfully published to google pubsub",
		Request: map[string]interface{}{
			"message":  message,
			"topic_id": topicID,
		},
		Response: map[string]interface{}{"server_id": serverID},
	}, "debug", nil)
}
