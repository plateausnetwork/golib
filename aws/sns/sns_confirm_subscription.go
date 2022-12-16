package sns

import "github.com/aws/aws-sdk-go/service/sns"

func (i implSNS) ConfirmSubscription(input *ConfirmSubscriptionInput) (*ConfirmSubscriptionOutput, error) {
	res, err := i.sns.ConfirmSubscription((*sns.ConfirmSubscriptionInput)(input))
	return (*ConfirmSubscriptionOutput)(res), err
}
