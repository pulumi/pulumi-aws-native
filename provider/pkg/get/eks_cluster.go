package get

import (
	"context"

	"github.com/aws/aws-sdk-go/service/eks"
)

func (g *Getter) getEKSClusterAttributes(ctx context.Context, physicalResourceID string) (map[string]interface{}, error) {
	resp, err := g.eks.DescribeClusterWithContext(ctx, &eks.DescribeClusterInput{
		// Add params here
	})
	_ = resp
	if err != nil {
		return nil, err
	}

	attrs := map[string]interface{}{
		"Arn": nil,
		"CertificateAuthorityData": nil,
		"Endpoint": nil,
	}
	return attrs, nil
}
