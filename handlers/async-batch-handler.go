package handlers

import (
	"bitbucket.org/ryanair/gofrlib/log"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/life4/genesis/slices"
)

// HandleAsync
// Example recordHandler function:
//
//	func (h *Handler) handleSqsMessage(ctx context.Context) func(record events.SQSMessage) string {
//	 	return func(record events.SQSMessage) string {
//				// Handle the record
//			}
//		}
func HandleAsync(sqs events.SQSEvent, recordHandler func(record events.SQSMessage) string) (events.SQSEventResponse, error) {
	failedMessageIds := slices.MapAsync(sqs.Records, len(sqs.Records), recordHandler)

	var failedMessages []events.SQSBatchItemFailure
	for _, failedMessageId := range slices.Filter(failedMessageIds, func(el string) bool { return el != "" }) {
		failedMessages = append(failedMessages, events.SQSBatchItemFailure{ItemIdentifier: failedMessageId})
	}

	if len(failedMessages) > 0 {
		log.ErrorW(fmt.Sprintf("Some of the messages failed: %s", log.ToString(failedMessages)), log.EventBody, log.ToString(sqs))
	}
	return events.SQSEventResponse{
		BatchItemFailures: failedMessages,
	}, nil
}
