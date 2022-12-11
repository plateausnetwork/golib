package sns

import "github.com/aws/aws-sdk-go/service/sns"

func (i implSNS) Publish(input PublishInput) (*PublishOutput, error) {
	res, err := i.sns.Publish((*sns.PublishInput)(&input))
	return (*PublishOutput)(res), err
}
