// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package emr

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws"
	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource schema for AWS::EMR::Studio
func LookupStudio(ctx *pulumi.Context, args *LookupStudioArgs, opts ...pulumi.InvokeOption) (*LookupStudioResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupStudioResult
	err := ctx.Invoke("aws-native:emr:getStudio", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupStudioArgs struct {
	// The ID of the EMR Studio.
	StudioId string `pulumi:"studioId"`
}

type LookupStudioResult struct {
	// The Amazon Resource Name (ARN) of the EMR Studio.
	Arn *string `pulumi:"arn"`
	// The default Amazon S3 location to back up EMR Studio Workspaces and notebook files. A Studio user can select an alternative Amazon S3 location when creating a Workspace.
	DefaultS3Location *string `pulumi:"defaultS3Location"`
	// A detailed description of the Studio.
	Description *string `pulumi:"description"`
	// Your identity provider's authentication endpoint. Amazon EMR Studio redirects federated users to this endpoint for authentication when logging in to a Studio with the Studio URL.
	IdpAuthUrl *string `pulumi:"idpAuthUrl"`
	// The name of relay state parameter for external Identity Provider.
	IdpRelayStateParameterName *string `pulumi:"idpRelayStateParameterName"`
	// A descriptive name for the Amazon EMR Studio.
	Name *string `pulumi:"name"`
	// The ID of the EMR Studio.
	StudioId *string `pulumi:"studioId"`
	// A list of up to 5 subnet IDs to associate with the Studio. The subnets must belong to the VPC specified by VpcId. Studio users can create a Workspace in any of the specified subnets.
	SubnetIds []string `pulumi:"subnetIds"`
	// A list of tags to associate with the Studio. Tags are user-defined key-value pairs that consist of a required key string with a maximum of 128 characters, and an optional value string with a maximum of 256 characters.
	Tags []aws.Tag `pulumi:"tags"`
	// The unique Studio access URL.
	Url *string `pulumi:"url"`
}

func LookupStudioOutput(ctx *pulumi.Context, args LookupStudioOutputArgs, opts ...pulumi.InvokeOption) LookupStudioResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupStudioResultOutput, error) {
			args := v.(LookupStudioArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("aws-native:emr:getStudio", args, LookupStudioResultOutput{}, options).(LookupStudioResultOutput), nil
		}).(LookupStudioResultOutput)
}

type LookupStudioOutputArgs struct {
	// The ID of the EMR Studio.
	StudioId pulumi.StringInput `pulumi:"studioId"`
}

func (LookupStudioOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupStudioArgs)(nil)).Elem()
}

type LookupStudioResultOutput struct{ *pulumi.OutputState }

func (LookupStudioResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupStudioResult)(nil)).Elem()
}

func (o LookupStudioResultOutput) ToLookupStudioResultOutput() LookupStudioResultOutput {
	return o
}

func (o LookupStudioResultOutput) ToLookupStudioResultOutputWithContext(ctx context.Context) LookupStudioResultOutput {
	return o
}

// The Amazon Resource Name (ARN) of the EMR Studio.
func (o LookupStudioResultOutput) Arn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupStudioResult) *string { return v.Arn }).(pulumi.StringPtrOutput)
}

// The default Amazon S3 location to back up EMR Studio Workspaces and notebook files. A Studio user can select an alternative Amazon S3 location when creating a Workspace.
func (o LookupStudioResultOutput) DefaultS3Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupStudioResult) *string { return v.DefaultS3Location }).(pulumi.StringPtrOutput)
}

// A detailed description of the Studio.
func (o LookupStudioResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupStudioResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

// Your identity provider's authentication endpoint. Amazon EMR Studio redirects federated users to this endpoint for authentication when logging in to a Studio with the Studio URL.
func (o LookupStudioResultOutput) IdpAuthUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupStudioResult) *string { return v.IdpAuthUrl }).(pulumi.StringPtrOutput)
}

// The name of relay state parameter for external Identity Provider.
func (o LookupStudioResultOutput) IdpRelayStateParameterName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupStudioResult) *string { return v.IdpRelayStateParameterName }).(pulumi.StringPtrOutput)
}

// A descriptive name for the Amazon EMR Studio.
func (o LookupStudioResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupStudioResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

// The ID of the EMR Studio.
func (o LookupStudioResultOutput) StudioId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupStudioResult) *string { return v.StudioId }).(pulumi.StringPtrOutput)
}

// A list of up to 5 subnet IDs to associate with the Studio. The subnets must belong to the VPC specified by VpcId. Studio users can create a Workspace in any of the specified subnets.
func (o LookupStudioResultOutput) SubnetIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupStudioResult) []string { return v.SubnetIds }).(pulumi.StringArrayOutput)
}

// A list of tags to associate with the Studio. Tags are user-defined key-value pairs that consist of a required key string with a maximum of 128 characters, and an optional value string with a maximum of 256 characters.
func (o LookupStudioResultOutput) Tags() aws.TagArrayOutput {
	return o.ApplyT(func(v LookupStudioResult) []aws.Tag { return v.Tags }).(aws.TagArrayOutput)
}

// The unique Studio access URL.
func (o LookupStudioResultOutput) Url() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupStudioResult) *string { return v.Url }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupStudioResultOutput{})
}
