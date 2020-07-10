package get

import (
	"context"

	"github.com/aws/aws-sdk-go/service/route53"
)

func (g *Getter) getRoute53HostedZoneAttributes(ctx context.Context, physicalResourceID string) (map[string]interface{}, error) {
	resp, err := g.route53.GetHostedZoneWithContext(ctx, &route53.GetHostedZoneInput{
		// Add params here
	})
	_ = resp
	if err != nil {
		return nil, err
	}

	attrs := map[string]interface{}{
		"NameServers": nil,
	}
	return attrs, nil
}
