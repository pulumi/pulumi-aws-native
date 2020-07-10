package get

import (
	"context"

	"github.com/aws/aws-sdk-go/service/appmesh"
)

func (g *Getter) getAppMeshVirtualServiceAttributes(ctx context.Context, physicalResourceID string) (map[string]interface{}, error) {
	resp, err := g.appmesh.DescribeVirtualServiceWithContext(ctx, &appmesh.DescribeVirtualServiceInput{
		// Add params here
	})
	_ = resp
	if err != nil {
		return nil, err
	}

	attrs := map[string]interface{}{
		"Arn": nil,
		"MeshName": nil,
		"Uid": nil,
		"VirtualServiceName": nil,
	}
	return attrs, nil
}
