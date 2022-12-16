package sns

import "github.com/aws/aws-sdk-go/service/sns"

type (
	ConfirmSubscriptionInput        sns.ConfirmSubscriptionInput
	ConfirmSubscriptionOutput       sns.ConfirmSubscriptionOutput
	CreatePlatformApplicationInput  sns.CreatePlatformApplicationInput
	CreatePlatformApplicationOutput sns.CreatePlatformApplicationOutput
	CreateTopicInput                sns.CreateTopicInput
	CreateTopicOutput               sns.CreateTopicOutput
	PublishInput                    sns.PublishInput
	PublishOutput                   sns.PublishOutput
	SubscribeInput                  sns.SubscribeInput
	SubscribeOutput                 sns.SubscribeOutput
	UnsubscribeInput                sns.UnsubscribeInput
	UnsubscribeOutput               sns.UnsubscribeOutput
)
