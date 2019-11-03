package get

import (
	"context"

	"github.com/aws/aws-sdk-go/service/route53resolver"
)

func (g *Getter) getRoute53ResolverResolverRuleAssociationAttributes(ctx context.Context, physicalResourceID string) (map[string]interface{}, error) {
	resp, err := g.route53resolver.GetResolverRuleAssociationWithContext(ctx, &route53resolver.GetResolverRuleAssociationInput{
		// Add params here
	})
	_ = resp
	if err != nil {
		return nil, err
	}

	attrs := map[string]interface{}{
		"Name": nil,
		"ResolverRuleAssociationId": nil,
		"ResolverRuleId": nil,
		"VPCId": nil,
	}
	return attrs, nil
}
