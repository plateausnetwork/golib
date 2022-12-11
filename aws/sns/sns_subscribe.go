package sns

import "github.com/aws/aws-sdk-go/service/sns"

func (i implSNS) Subscribe(input *SubscribeInput) (*SubscribeOutput, error) {
	res, err := i.sns.Subscribe((*sns.SubscribeInput)(input))
	return (*SubscribeOutput)(res), err
}
