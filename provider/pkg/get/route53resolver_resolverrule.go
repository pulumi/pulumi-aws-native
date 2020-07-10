package get

import (
	"context"

	"github.com/aws/aws-sdk-go/service/route53resolver"
)

func (g *Getter) getRoute53ResolverResolverRuleAttributes(ctx context.Context, physicalResourceID string) (map[string]interface{}, error) {
	resp, err := g.route53resolver.GetResolverRuleWithContext(ctx, &route53resolver.GetResolverRuleInput{
		// Add params here
	})
	_ = resp
	if err != nil {
		return nil, err
	}

	attrs := map[string]interface{}{
		"Arn": nil,
		"DomainName": nil,
		"Name": nil,
		"ResolverEndpointId": nil,
		"ResolverRuleId": nil,
		"TargetIps": nil,
	}
	return attrs, nil
}
