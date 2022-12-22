package consumer

import (
	"context"
	"encoding/json"
	"time"

	"github.com/aws/aws-sdk-go/service/sqs"
)

func (i implConsumer) Run(ctx context.Context, chResponse chan Response) {
	defer close(chResponse)
	for {
		select {
		case <-ctx.Done():
			return
		default:
			output, err := i.sqs.ReceiveMessage(&sqs.ReceiveMessageInput{
				QueueUrl:            i.queueURL,
				MaxNumberOfMessages: i.maxNumberOfMessages,
			})
			if err != nil {
				chResponse <- Response{Error: err}
				continue
			}
			if output == nil || output.Messages == nil || len(output.Messages) == 0 {
				continue
			}
			for _, msg := range output.Messages {
				var body MessageBody
				if err := json.Unmarshal([]byte(*msg.Body), &body); err != nil {
					chResponse <- Response{Error: err}
					continue
				}
				chResponse <- Response{
					Error:       nil,
					Message:     (*Message)(msg),
					MessageBody: &body,
				}
			}
			if i.readMessageDelay > 0 {
				time.Sleep(i.readMessageDelay)
			}
		}
	}
}
