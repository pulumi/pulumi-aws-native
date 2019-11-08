package get

import (
	"context"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/service/iam"
)

func (g *Getter) getIAMInstanceProfileAttributes(ctx context.Context, physicalResourceID string) (map[string]interface{}, error) {
	resp, err := g.iam.GetInstanceProfileWithContext(ctx, &iam.GetInstanceProfileInput{
		InstanceProfileName: aws.String(physicalResourceID),
	})
	if err != nil {
		if awsErr, ok := err.(awserr.Error); ok && awsErr.Code() == "NoSuchEntity" {
			return nil, NewNotFoundError(err)
		}
		return nil, err
	}
	instanceProfile := resp.InstanceProfile

	arn := aws.StringValue(instanceProfile.Arn)

	attrs := map[string]interface{}{
		"Arn": arn,
	}
	return attrs, nil
}
