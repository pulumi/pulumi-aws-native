// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package sagemaker

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::SageMaker::Space
func LookupSpace(ctx *pulumi.Context, args *LookupSpaceArgs, opts ...pulumi.InvokeOption) (*LookupSpaceResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupSpaceResult
	err := ctx.Invoke("aws-native:sagemaker:getSpace", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupSpaceArgs struct {
	// The ID of the associated Domain.
	DomainId string `pulumi:"domainId"`
	// A name for the Space.
	SpaceName string `pulumi:"spaceName"`
}

type LookupSpaceResult struct {
	// The space Amazon Resource Name (ARN).
	SpaceArn *string `pulumi:"spaceArn"`
	// The name of the space that appears in the Studio UI.
	SpaceDisplayName *string `pulumi:"spaceDisplayName"`
	// Returns the URL of the space. If the space is created with AWS IAM Identity Center (Successor to AWS Single Sign-On) authentication, users can navigate to the URL after appending the respective redirect parameter for the application type to be federated through AWS IAM Identity Center.
	//
	// The following application types are supported:
	//
	// - Studio Classic: `&redirect=JupyterServer`
	// - JupyterLab: `&redirect=JupyterLab`
	// - Code Editor, based on Code-OSS, Visual Studio Code - Open Source: `&redirect=CodeEditor`
	Url *string `pulumi:"url"`
}

func LookupSpaceOutput(ctx *pulumi.Context, args LookupSpaceOutputArgs, opts ...pulumi.InvokeOption) LookupSpaceResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSpaceResultOutput, error) {
			args := v.(LookupSpaceArgs)
			opts = internal.PkgInvokeDefaultOpts(opts)
			var rv LookupSpaceResult
			secret, err := ctx.InvokePackageRaw("aws-native:sagemaker:getSpace", args, &rv, "", opts...)
			if err != nil {
				return LookupSpaceResultOutput{}, err
			}

			output := pulumi.ToOutput(rv).(LookupSpaceResultOutput)
			if secret {
				return pulumi.ToSecret(output).(LookupSpaceResultOutput), nil
			}
			return output, nil
		}).(LookupSpaceResultOutput)
}

type LookupSpaceOutputArgs struct {
	// The ID of the associated Domain.
	DomainId pulumi.StringInput `pulumi:"domainId"`
	// A name for the Space.
	SpaceName pulumi.StringInput `pulumi:"spaceName"`
}

func (LookupSpaceOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSpaceArgs)(nil)).Elem()
}

type LookupSpaceResultOutput struct{ *pulumi.OutputState }

func (LookupSpaceResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSpaceResult)(nil)).Elem()
}

func (o LookupSpaceResultOutput) ToLookupSpaceResultOutput() LookupSpaceResultOutput {
	return o
}

func (o LookupSpaceResultOutput) ToLookupSpaceResultOutputWithContext(ctx context.Context) LookupSpaceResultOutput {
	return o
}

// The space Amazon Resource Name (ARN).
func (o LookupSpaceResultOutput) SpaceArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSpaceResult) *string { return v.SpaceArn }).(pulumi.StringPtrOutput)
}

// The name of the space that appears in the Studio UI.
func (o LookupSpaceResultOutput) SpaceDisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSpaceResult) *string { return v.SpaceDisplayName }).(pulumi.StringPtrOutput)
}

// Returns the URL of the space. If the space is created with AWS IAM Identity Center (Successor to AWS Single Sign-On) authentication, users can navigate to the URL after appending the respective redirect parameter for the application type to be federated through AWS IAM Identity Center.
//
// The following application types are supported:
//
// - Studio Classic: `&redirect=JupyterServer`
// - JupyterLab: `&redirect=JupyterLab`
// - Code Editor, based on Code-OSS, Visual Studio Code - Open Source: `&redirect=CodeEditor`
func (o LookupSpaceResultOutput) Url() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSpaceResult) *string { return v.Url }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSpaceResultOutput{})
}
