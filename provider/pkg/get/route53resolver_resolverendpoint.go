package get

import (
	"context"

	"github.com/aws/aws-sdk-go/service/route53resolver"
)

func (g *Getter) getRoute53ResolverResolverEndpointAttributes(ctx context.Context, physicalResourceID string) (map[string]interface{}, error) {
	resp, err := g.route53resolver.GetResolverEndpointWithContext(ctx, &route53resolver.GetResolverEndpointInput{
		// Add params here
	})
	_ = resp
	if err != nil {
		return nil, err
	}

	attrs := map[string]interface{}{
		"Arn": nil,
		"Direction": nil,
		"HostVPCId": nil,
		"IpAddressCount": nil,
		"Name": nil,
		"ResolverEndpointId": nil,
	}
	return attrs, nil
}
