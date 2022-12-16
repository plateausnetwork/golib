package sns

import "github.com/aws/aws-sdk-go/service/sns"

func (i implSNS) CreateTopic(input CreateTopicInput) (*CreateTopicOutput, error) {
	res, err := i.sns.CreateTopic((*sns.CreateTopicInput)(&input))
	return (*CreateTopicOutput)(res), err
}
