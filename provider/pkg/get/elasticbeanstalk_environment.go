package get

import (
	"context"

	"github.com/aws/aws-sdk-go/service/elasticbeanstalk"
)

func (g *Getter) getElasticBeanstalkEnvironmentAttributes(ctx context.Context, physicalResourceID string) (map[string]interface{}, error) {
	resp, err := g.elasticbeanstalk.DescribeEnvironmentsWithContext(ctx, &elasticbeanstalk.DescribeEnvironmentsInput{
		// Add params here
	})
	_ = resp
	if err != nil {
		return nil, err
	}

	attrs := map[string]interface{}{
		"EndpointURL": nil,
	}
	return attrs, nil
}
