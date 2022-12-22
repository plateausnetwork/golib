package consumer

import "github.com/aws/aws-sdk-go/service/sqs"

func (i implConsumer) DeleteMessage(messageReceipt *string) error {
	_, err := i.sqs.DeleteMessage(&sqs.DeleteMessageInput{
		QueueUrl:      i.queueURL,
		ReceiptHandle: messageReceipt,
	})
	return err
}
