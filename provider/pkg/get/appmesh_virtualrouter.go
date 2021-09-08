package get

import (
	"context"

	"github.com/aws/aws-sdk-go/service/appmesh"
)

func (g *Getter) getAppMeshVirtualRouterAttributes(ctx context.Context, physicalResourceID string) (map[string]interface{}, error) {
	resp, err := g.appmesh.DescribeVirtualRouterWithContext(ctx, &appmesh.DescribeVirtualRouterInput{
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
		"VirtualRouterName": nil,
	}
	return attrs, nil
}
