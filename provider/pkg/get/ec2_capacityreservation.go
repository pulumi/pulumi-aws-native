package get

import (
	"context"

	"github.com/aws/aws-sdk-go/service/ec2"
)

func (g *Getter) getEC2CapacityReservationAttributes(ctx context.Context, physicalResourceID string) (map[string]interface{}, error) {
	resp, err := g.ec2.DescribeCapacityReservationsWithContext(ctx, &ec2.DescribeCapacityReservationsInput{
		// Add params here
	})
	_ = resp
	if err != nil {
		return nil, err
	}

	attrs := map[string]interface{}{
		"AvailabilityZone": nil,
		"AvailableInstanceCount": nil,
		"InstanceType": nil,
		"Tenancy": nil,
		"TotalInstanceCount": nil,
	}
	return attrs, nil
}
