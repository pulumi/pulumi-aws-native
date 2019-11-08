package get

import (
	"context"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/service/iam"
)

func (g *Getter) getIAMRoleAttributes(ctx context.Context, physicalResourceID string) (map[string]interface{}, error) {
	resp, err := g.iam.GetRoleWithContext(ctx, &iam.GetRoleInput{
		RoleName: aws.String(physicalResourceID),
	})
	if err != nil {
		if awsErr, ok := err.(awserr.Error); ok && awsErr.Code() == iam.ErrCodeNoSuchEntityException {
			return nil, NewNotFoundError(err)
		}
		return nil, err
	}
	role := resp.Role

	arn := aws.StringValue(role.Arn)
	roleID := aws.StringValue(role.RoleId)

	attrs := map[string]interface{}{
		"Arn":    arn,
		"RoleId": roleID,
	}
	return attrs, nil
}
