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
		RefReturnsArn,
		Categorize("When you pass the logical ID of this resource to the intrinsic `Ref`function, `Ref`returns the resource ARN\\. For example:"))

	assert.Equal(t,
		RefReturnsName,
		Categorize("When you pass the logical ID of this resource to the intrinsic `Ref`function, `Ref`returns the value of `DashboardName`\\."))

	assert.Equal(t,
		RefReturnsName,
		Categorize("When you pass the logical ID of this resource to the intrinsic `Ref`function, `Ref`returns the name of the resource\\."))

	assert.Equal(t,
		RefReturnsName,
		Categorize("When you pass the logical ID of this resource to the intrinsic `Ref`function, `Ref`returns the DB instance name\\."))

	assert.Equal(t,
		RefReturnsArn,
		Categorize("When you pass the logical ID of this resource to the intrinsic `Ref`function, `Ref`returns the scheduling policy ARN, such as `arn:aws:batch:us-east-1:111122223333:scheduling-policy/HighPriority`\\."))

	assert.Equal(t,
		RefReturnsID,
		Categorize("When you pass the logical ID of this resource to the intrinsic `Ref`function, `Ref`returns the wireless device ID\\."))

	assert.Equal(t,
		RefReturnsArn,
		Categorize("When you pass the logical ID of this resource to the intrinsic `Ref`function, `Ref`returns the Amazon Resource Name \\(ARN\\) of the endpoint configuration, such as `arn:aws:sagemaker:us-west-2:012345678901:notebook-instance-lifecycle-config/mylifecycleconfig`"))
}
