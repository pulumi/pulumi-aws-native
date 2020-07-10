package get

import (
	"context"
	"fmt"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/endpoints"
	"github.com/aws/aws-sdk-go/service/s3"
)

const hyphenatedS3Website = "%s.s3-website-%s.amazonaws.com"
const dottedS3Website = "%s.s3-website.%s.amazonaws.com"
const dottedChinaS3Website = "%s.s3-website.%s.amazonaws.com.cn"

var s3WebsiteFormats = map[string]string{
	"ap-east-1":      dottedS3Website,
	"ap-northeast-1": hyphenatedS3Website,
	"ap-northeast-2": dottedS3Website,
	"ap-northeast-3": dottedS3Website,
	"ap-south-1":     dottedS3Website,
	"ap-southeast-1": hyphenatedS3Website,
	"ap-southeast-2": hyphenatedS3Website,
	"ca-central-1":   dottedS3Website,
	"cn-northwest-1": dottedChinaS3Website,
	"eu-central-1":   dottedS3Website,
	"eu-north-1":     dottedS3Website,
	"eu-west-1":      hyphenatedS3Website,
	"eu-west-2":      dottedS3Website,
	"eu-west-3":      dottedS3Website,
	"me-south-1":     dottedS3Website,
	"sa-east-1":      hyphenatedS3Website,
	"us-east-1":      hyphenatedS3Website,
	"us-east-2":      dottedS3Website,
	"us-west-1":      hyphenatedS3Website,
	"us-west-2":      hyphenatedS3Website,
}

func (g *Getter) getS3BucketAttributes(ctx context.Context, physicalResourceID string) (map[string]interface{}, error) {
	_, err := g.s3.HeadBucketWithContext(ctx, &s3.HeadBucketInput{
		Bucket: aws.String(physicalResourceID),
	})
	if err != nil {
		if awsErr, ok := err.(awserr.RequestFailure); ok && awsErr.StatusCode() == 404 {
			return nil, NewNotFoundError(err)
		}
		return nil, err
	}

	locationResp, err := g.s3.GetBucketLocationWithContext(ctx, &s3.GetBucketLocationInput{
		Bucket: aws.String(physicalResourceID),
	})
	if err != nil {
		return nil, err
	}

	// Get the region or default to us-east-1 (https://docs.aws.amazon.com/en_pv/AmazonS3/latest/API/API_GetBucketLocation.html)
	region := aws.StringValue(locationResp.LocationConstraint)
	if region == "" {
		region = "us-east-1"
	}
	endpoint, err := endpoints.DefaultResolver().EndpointFor(endpoints.S3ServiceID, region)
	if err != nil {
		return nil, err
	}

	websiteFormat, ok := s3WebsiteFormats[region]
	if !ok {
		websiteFormat = dottedS3Website
	}

	arn := fmt.Sprintf("arn:aws:s3:::%s", physicalResourceID)
	domainName := fmt.Sprintf("%s.s3.amazonaws.com", physicalResourceID)
	dualStackDomainName := fmt.Sprintf("%s.s3.dualstack.%s.amazonaws.com", physicalResourceID, region)
	regionalDomainName := fmt.Sprintf("%s.%s", physicalResourceID, strings.TrimPrefix(endpoint.URL, "https://"))
	websiteURL := fmt.Sprintf(websiteFormat, physicalResourceID, region)

	attrs := map[string]interface{}{
		"Arn":                 arn,
		"DomainName":          domainName,
		"DualStackDomainName": dualStackDomainName,
		"RegionalDomainName":  regionalDomainName,
		"WebsiteURL":          websiteURL,
	}
	return attrs, nil
}
