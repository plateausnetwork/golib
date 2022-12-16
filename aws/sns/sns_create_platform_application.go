package sns

import "github.com/aws/aws-sdk-go/service/sns"

func (i implSNS) CreatePlatformApplication(input *CreatePlatformApplicationInput) (
	*CreatePlatformApplicationOutput, error,
) {
	res, err := i.sns.CreatePlatformApplication((*sns.CreatePlatformApplicationInput)(input))
	return (*CreatePlatformApplicationOutput)(res), err
}
