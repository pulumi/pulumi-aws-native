package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCategorize(t *testing.T) {
	assert.Equal(t,
		RefReturnsArn,
		Categorize("\nThe Amazon Resource Name \\(ARN\\) of the certificate authority\\.\n\n"))

	assert.Equal(t,
		RefReturnsArn,
		Categorize("When you pass the logical ID of this resource to the intrinsic `Ref`function, `Ref`returns the ARN of the analyzer created\\.\n\nFor more information about using the `Ref`function, see [Ref](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/intrinsic-function-reference-ref.html"))

	assert.Equal(t,
		ShouldNotBeUsed,
		Categorize("\nThis reference should not be used in CloudFormation templates\\. Instead, use `AWS::ACMPCA::Certificate.Arn` to identify a certificate, and use `AWS::ACMPCA::Certificate.CertificateAuthorityArn` to identify a certificate authority\\.\n\n"))

	assert.Equal(t,
		RefReturnsArn,
		Categorize("When you pass the logical ID of this resource to the intrinsic `Ref`function, `Ref`returns the Arn of the Cost Category that is created by the template\\."))

	assert.Equal(t,
		RefReturnsName,
		Categorize("When you pass the logical ID of this resource to the intrinsic `Ref`function, `Ref`returns the value of `DashboardName`\\."))
}
