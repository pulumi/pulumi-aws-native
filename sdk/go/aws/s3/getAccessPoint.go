// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package s3

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The AWS::S3::AccessPoint resource is an Amazon S3 resource type that you can use to access buckets.
func LookupAccessPoint(ctx *pulumi.Context, args *LookupAccessPointArgs, opts ...pulumi.InvokeOption) (*LookupAccessPointResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupAccessPointResult
	err := ctx.Invoke("aws-native:s3:getAccessPoint", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupAccessPointArgs struct {
	// The name you want to assign to this Access Point. If you don't specify a name, AWS CloudFormation generates a unique ID and uses that ID for the access point name.
	Name string `pulumi:"name"`
}

type LookupAccessPointResult struct {
	// The alias of this Access Point. This alias can be used for compatibility purposes with other AWS services and third-party applications.
	Alias *string `pulumi:"alias"`
	// The Amazon Resource Name (ARN) of the specified accesspoint.
	Arn *string `pulumi:"arn"`
	// Indicates whether this Access Point allows access from the public Internet. If VpcConfiguration is specified for this Access Point, then NetworkOrigin is VPC, and the Access Point doesn't allow access from the public Internet. Otherwise, NetworkOrigin is Internet, and the Access Point allows access from the public Internet, subject to the Access Point and bucket access policies.
	NetworkOrigin *AccessPointNetworkOrigin `pulumi:"networkOrigin"`
	// The Access Point Policy you want to apply to this access point.
	//
	// Search the [CloudFormation User Guide](https://docs.aws.amazon.com/cloudformation/) for `AWS::S3::AccessPoint` for more information about the expected schema for this property.
	Policy interface{} `pulumi:"policy"`
	// The PublicAccessBlock configuration that you want to apply to this Access Point. You can enable the configuration options in any combination. For more information about when Amazon S3 considers a bucket or object public, see https://docs.aws.amazon.com/AmazonS3/latest/dev/access-control-block-public-access.html#access-control-block-public-access-policy-status 'The Meaning of Public' in the Amazon Simple Storage Service Developer Guide.
	PublicAccessBlockConfiguration *AccessPointPublicAccessBlockConfiguration `pulumi:"publicAccessBlockConfiguration"`
}

func LookupAccessPointOutput(ctx *pulumi.Context, args LookupAccessPointOutputArgs, opts ...pulumi.InvokeOption) LookupAccessPointResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupAccessPointResultOutput, error) {
			args := v.(LookupAccessPointArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("aws-native:s3:getAccessPoint", args, LookupAccessPointResultOutput{}, options).(LookupAccessPointResultOutput), nil
		}).(LookupAccessPointResultOutput)
}

type LookupAccessPointOutputArgs struct {
	// The name you want to assign to this Access Point. If you don't specify a name, AWS CloudFormation generates a unique ID and uses that ID for the access point name.
	Name pulumi.StringInput `pulumi:"name"`
}

func (LookupAccessPointOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAccessPointArgs)(nil)).Elem()
}

type LookupAccessPointResultOutput struct{ *pulumi.OutputState }

func (LookupAccessPointResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAccessPointResult)(nil)).Elem()
}

func (o LookupAccessPointResultOutput) ToLookupAccessPointResultOutput() LookupAccessPointResultOutput {
	return o
}

func (o LookupAccessPointResultOutput) ToLookupAccessPointResultOutputWithContext(ctx context.Context) LookupAccessPointResultOutput {
	return o
}

// The alias of this Access Point. This alias can be used for compatibility purposes with other AWS services and third-party applications.
func (o LookupAccessPointResultOutput) Alias() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAccessPointResult) *string { return v.Alias }).(pulumi.StringPtrOutput)
}

// The Amazon Resource Name (ARN) of the specified accesspoint.
func (o LookupAccessPointResultOutput) Arn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAccessPointResult) *string { return v.Arn }).(pulumi.StringPtrOutput)
}

// Indicates whether this Access Point allows access from the public Internet. If VpcConfiguration is specified for this Access Point, then NetworkOrigin is VPC, and the Access Point doesn't allow access from the public Internet. Otherwise, NetworkOrigin is Internet, and the Access Point allows access from the public Internet, subject to the Access Point and bucket access policies.
func (o LookupAccessPointResultOutput) NetworkOrigin() AccessPointNetworkOriginPtrOutput {
	return o.ApplyT(func(v LookupAccessPointResult) *AccessPointNetworkOrigin { return v.NetworkOrigin }).(AccessPointNetworkOriginPtrOutput)
}

// The Access Point Policy you want to apply to this access point.
//
// Search the [CloudFormation User Guide](https://docs.aws.amazon.com/cloudformation/) for `AWS::S3::AccessPoint` for more information about the expected schema for this property.
func (o LookupAccessPointResultOutput) Policy() pulumi.AnyOutput {
	return o.ApplyT(func(v LookupAccessPointResult) interface{} { return v.Policy }).(pulumi.AnyOutput)
}

// The PublicAccessBlock configuration that you want to apply to this Access Point. You can enable the configuration options in any combination. For more information about when Amazon S3 considers a bucket or object public, see https://docs.aws.amazon.com/AmazonS3/latest/dev/access-control-block-public-access.html#access-control-block-public-access-policy-status 'The Meaning of Public' in the Amazon Simple Storage Service Developer Guide.
func (o LookupAccessPointResultOutput) PublicAccessBlockConfiguration() AccessPointPublicAccessBlockConfigurationPtrOutput {
	return o.ApplyT(func(v LookupAccessPointResult) *AccessPointPublicAccessBlockConfiguration {
		return v.PublicAccessBlockConfiguration
	}).(AccessPointPublicAccessBlockConfigurationPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupAccessPointResultOutput{})
}
