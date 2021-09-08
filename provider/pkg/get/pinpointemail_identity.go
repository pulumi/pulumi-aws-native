package get

import (
	"context"

	_ "github.com/aws/aws-sdk-go/service/pinpointemail"
)

func (g *Getter) getPinpointEmailIdentityAttributes(ctx context.Context, physicalResourceID string) (map[string]interface{}, error) {
	// resp, err := g.pinpointemail.
	var err error
	if err != nil {
		return nil, err
	}

	attrs := map[string]interface{}{
		"IdentityDNSRecordName1": nil,
		"IdentityDNSRecordName2": nil,
		"IdentityDNSRecordName3": nil,
		"IdentityDNSRecordValue1": nil,
		"IdentityDNSRecordValue2": nil,
		"IdentityDNSRecordValue3": nil,
	}
	return attrs, nil
}
