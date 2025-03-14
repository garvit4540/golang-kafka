package main

import (
	"encoding/json"
	"github.com/IBM/sarama"
	trace "github.com/garvit4540/golang-kafka"
	fiber "github.com/gofiber/fiber/v2"
)

func main() {

	app := fiber.New()
	api := app.Group("/api/v1")
	api.Post("/comments", createComment)

	// Start the app
	startApp(app)
}

// Structs
type Comment struct {
	Text string `form:"text" json:"text"`
}

// Controller Functions
func createComment(ctx *fiber.Ctx) error {

	// Initialise new comment
	cmt := new(Comment)

	// Get and parse comment from ctx
	if err := ctx.BodyParser(cmt); err != nil {
		trace.LogInfo(trace.ErrorParsingFromContext, map[string]interface{}{
			"error": err,
		})
		ctx.Status(trace.StatusCode400).JSON(map[string]interface{}{
			"success": false,
			"error":   err,
		})
		return err
	}

	// Convert body to bytes
	cmdInBytes, err := json.Marshal(cmt)
	if err != nil {
		trace.LogInfo(trace.ErrorJSONMarshalUnmarshalError, map[string]interface{}{
			"error": err,
		})
	}

	pushCommentToQueue("comments", cmdInBytes)

	// Return Comment in JSON format
	err = ctx.JSON(map[string]interface{}{
		"success": true,
		"message": "Comment Pushed Successfully",
		"comment": cmt,
	})
	if err != nil {
		ctx.Status(trace.StatusCode500).JSON(map[string]interface{}{
			"success": false,
			"error":   "Error Creating Product",
		})
	}

	return nil
}

// Helper Functions

func startApp(app *fiber.App) {
	err := app.Listen(":3000")
	if err != nil {
		trace.LogInfo(trace.ErrorStartingApp, map[string]interface{}{
			"port": 3000,
			"err":  err,
		})
	}
	trace.LogInfo(trace.AppStarted, map[string]interface{}{
		"port": 3000,
	})
}

func pushCommentToQueue(topic string, message []byte) error {
	brokersUrl := []string{"localhost:29092"}

	producer, err := connectProducer(brokersUrl)
	if err != nil {
		trace.LogInfo(trace.ErrorConnectingWithProducer, map[string]interface{}{
			"error": err,
		})
		return err
	}

	defer producer.Close()

	msg := &sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.StringEncoder(message),
	}

	partition, offset, err := producer.SendMessage(msg)
	if err != nil {
		trace.LogInfo(trace.ErrorSendingMessageToKafka, map[string]interface{}{
			"error": err,
		})
		return err
	}

	trace.LogInfo(trace.MessageSuccessFullySentToKafka, map[string]interface{}{
		"topic":     topic,
		"partition": partition,
		"offset":    offset,
	})

	return nil

}

func connectProducer(brokersUrl []string) (sarama.SyncProducer, error) {

	config := sarama.NewConfig()
	config.Producer.Return.Successes = true
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Retry.Max = 5

	conn, err := sarama.NewSyncProducer(brokersUrl, config)
	if err != nil {
		trace.LogInfo(trace.ErrorWhileMakingNewSyncProducer, map[string]interface{}{
			"error": err,
		})
		return nil, err
	}

	return conn, nil

}
