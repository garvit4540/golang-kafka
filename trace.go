package trace

import "fmt"

// Status Codes
const (
	StatusCode400 = 400
	StatusCode500 = 500
)

// Error Codes
const (
	ErrorStartingApp                 = "ERROR_STATRTING_APP"
	ErrorParsingFromContext          = "ERROR_PARSING_FROM_CONTEXT"
	ErrorJSONMarshalUnmarshalError   = "JSON_MARSHAL_UNMARSHAL_ERROR"
	ErrorConnectingProducer          = "ERROR_CONNECTING_PRODUCER"
	ErrorSendingMessageToKafka       = "ERROR_SENDING_MESSAGE_TO_KAFKA"
	ErrorWhileMakingNewSyncProducer  = "ERROR_WHILE_MAKING_NEW_SYNC_PRODUCER"
	ErrorConnectingConsumer          = "ERROR_CONNECTING_CONSUMER"
	ErrorConsumingPartition          = "ERROR_CONSUMING_PARTITION"
	ErrorReceivedFromKafkaToConsumer = "ERROR_RECEIVED_FROM_KAFKA_TO_CONSUMER"
	ErrorClosingWorker               = "ERROR_CLOSING_WORKER"
	ErrorWhileMakingNewSyncConsumer  = "ERROR_WHILE_MAKING_NEW_SYNC_CONSUMER"
)

// Info Codes
const (
	AppStarted                              = "APP_STARTED"
	MessageSuccessFullySentToKafka          = "MESSAGE_SUCCESSFULLY_SENT_TO_KAFKA"
	ConsumerStarted                         = "CONSUMER_STARTED"
	MessagesProcessed                       = "MESSAGES_PROCESSED"
	MessageReceivedFromKafkaToConsumer      = "MSG_RECEIVED_FROM_KAFKA_TO_CONSUMER"
	InterruptionReceivedFromKafkaToConsumer = "INTERRUPTION_RECEIVED_FROM_KAFKA_TO_CONSUMER"
)

func LogInfo(traceCode string, data map[string]interface{}) {
	fmt.Println(traceCode, " data: ", data)
}
