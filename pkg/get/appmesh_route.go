package get

import (
	"context"

	"github.com/aws/aws-sdk-go/service/appmesh"
)

func (g *Getter) getAppMeshRouteAttributes(ctx context.Context, physicalResourceID string) (map[string]interface{}, error) {
	resp, err := g.appmesh.DescribeRouteWithContext(ctx, &appmesh.DescribeRouteInput{
		// Add params here
	})
	_ = resp
	if err != nil {
		return nil, err
	}

	attrs := map[string]interface{}{
		"Arn": nil,
		"MeshName": nil,
		"RouteName": nil,
		"Uid": nil,
		"VirtualRouterName": nil,
	}
	return attrs, nil
}
