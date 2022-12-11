package sns

import "github.com/aws/aws-sdk-go/service/sns"

type (
	CreateTopicInput  sns.CreateTopicInput
	CreateTopicOutput sns.CreateTopicOutput
	PublishInput      sns.PublishInput
	PublishOutput     sns.PublishOutput
	SubscribeInput    sns.SubscribeInput
	SubscribeOutput   sns.SubscribeOutput
	UnsubscribeInput  sns.UnsubscribeInput
	UnsubscribeOutput sns.UnsubscribeOutput
)
