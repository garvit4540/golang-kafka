package consumers

import (
	"github.com/IBM/sarama"
	trace "github.com/garvit4540/golang-kafka"
	"os"
	"os/signal"
	"syscall"
)

func main() {

	topic := "comments"
	brokersUrl := []string{"localhost:29092"}
	worker, err := connectConsumer(brokersUrl)
	if err != nil {
		trace.LogInfo(trace.ErrorConnectingConsumer, map[string]interface{}{
			"error": err,
		})
		panic(err)
	}

	consumer, err := worker.ConsumePartition(topic, 0, sarama.OffsetOldest)
	if err != nil {
		trace.LogInfo(trace.ErrorConsumingPartition, map[string]interface{}{
			"error": err,
		})
		panic(err)
	}
	trace.LogInfo(trace.ConsumerStarted, map[string]interface{}{
		"error": err,
	})

	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, syscall.SIGINT, syscall.SIGTERM)
	msgCount := 0

	doneCh := make(chan struct{})
	go func() {
		for {
			select {
			case err := <-consumer.Errors():
				trace.LogInfo(trace.ErrorReceivedFromKafkaToConsumer, map[string]interface{}{
					"error": err,
				})
			case msg := <-consumer.Messages():
				msgCount++
				trace.LogInfo(trace.MessageReceivedFromKafkaToConsumer, map[string]interface{}{
					"Topic":           string(msg.Topic),
					"Current Message": string(msg.Value),
					"Message Count":   msgCount,
				})
			case <-sigchan:
				trace.LogInfo(trace.InterruptionReceivedFromKafkaToConsumer, nil)
				doneCh <- struct{}{}
			}
		}
	}()

	<-doneCh
	trace.LogInfo(trace.MessagesProcessed, map[string]interface{}{
		"Processed Messages Count": msgCount,
	})

	if err := worker.Close(); err != nil {
		trace.LogInfo(trace.ErrorClosingWorker, map[string]interface{}{
			"error": err,
		})
		panic(err)
	}

}

// Helper Functions
func connectConsumer(brokersUrl []string) (sarama.Consumer, error) {

	config := sarama.NewConfig()
	config.Consumer.Return.Errors = true

	conn, err := sarama.NewConsumer(brokersUrl, config)
	if err != nil {
		trace.LogInfo(trace.ErrorWhileMakingNewSyncConsumer, map[string]interface{}{
			"error": err,
		})
		return nil, err
	}

	return conn, nil
}
