package trace

import "fmt"

// Status Codes
const (
	StatusCode400 = 400
	StatusCode500 = 500
)

// Error Codes
const (
	ErrorStartingApp                = "ERROR_STATRTING_APP"
	ErrorParsingFromContext         = "ERROR_PARSING_FROM_CONTEXT"
	ErrorJSONMarshalUnmarshalError  = "JSON_MARSHAL_UNMARSHAL_ERROR"
	ErrorConnectingWithProducer     = "ERROR_CONNECTING_WITH_PRODUCER"
	ErrorSendingMessageToKafka      = "ERROR_SENDING_MESSAGE_TO_KAFKA"
	ErrorWhileMakingNewSyncProducer = "ERROR_WHILE_MAKING_NEW_SYNC_PRODUCER"
)

// Info Codes
const (
	AppStarted                     = "APP_STARTED"
	MessageSuccessFullySentToKafka = "MESSAGE_SUCCESSFULLY_SENT_TO_KAFKA"
)

func LogInfo(traceCode string, data map[string]interface{}) {
	fmt.Println(traceCode, " data: ", data)
}
