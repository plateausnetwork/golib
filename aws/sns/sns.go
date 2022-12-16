//go:generate mockgen -source=${GOFILE} -package=${GOPACKAGE} -destination=${GOPACKAGE}_mock.go
package sns

import (
	"github.com/aws/aws-sdk-go/service/sns"
	"github.com/rhizomplatform/golib/aws/session"
)

type (
	Options struct {
		Session session.Session
	}
	SNS interface {
		CreateTopic(create CreateTopicInput) (*CreateTopicOutput, error)
		Publish(input PublishInput) (*PublishOutput, error)
		Subscribe(input *SubscribeInput) (*SubscribeOutput, error)
		Unsubscribe(input *UnsubscribeInput) (*UnsubscribeOutput, error)
	}
	implSNS struct {
		sns *sns.SNS
	}
)

func New(opts Options) SNS {
	return implSNS{
		sns: sns.New(opts.Session.GetSession()),
	}
}
