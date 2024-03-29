//go:generate mockgen -source=${GOFILE} -package=${GOPACKAGE} -destination=${GOPACKAGE}_mock.go
package consumer

import (
	"context"
	"time"

	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/rhizomplatform/golib/aws/session"
)

type (
	Options struct {
		QueueURL            *string
		MaxNumberOfMessages *int64
		Session             session.Session
		ReadMessageDelay    time.Duration
	}
	Consumer interface {
		Run(ctx context.Context, chResponse chan Response)
		DeleteMessage(messageReceipt *string) error
	}
	implConsumer struct {
		queueURL            *string
		maxNumberOfMessages *int64
		sqs                 *sqs.SQS
		readMessageDelay    time.Duration
	}
)

func New(opts Options) Consumer {
	return implConsumer{
		sqs:                 sqs.New(opts.Session.GetSession()),
		queueURL:            opts.QueueURL,
		maxNumberOfMessages: opts.MaxNumberOfMessages,
		readMessageDelay:    opts.ReadMessageDelay,
	}
}
