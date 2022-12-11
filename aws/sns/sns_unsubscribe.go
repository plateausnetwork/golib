package sns

import "github.com/aws/aws-sdk-go/service/sns"

func (i implSNS) Unsubscribe(input *UnsubscribeInput) (*UnsubscribeOutput, error) {
	res, err := i.sns.Unsubscribe((*sns.UnsubscribeInput)(input))
	return (*UnsubscribeOutput)(res), err
}
