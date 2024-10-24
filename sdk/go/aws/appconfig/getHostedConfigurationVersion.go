// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package appconfig

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::AppConfig::HostedConfigurationVersion
func LookupHostedConfigurationVersion(ctx *pulumi.Context, args *LookupHostedConfigurationVersionArgs, opts ...pulumi.InvokeOption) (*LookupHostedConfigurationVersionResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupHostedConfigurationVersionResult
	err := ctx.Invoke("aws-native:appconfig:getHostedConfigurationVersion", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupHostedConfigurationVersionArgs struct {
	// The application ID.
	ApplicationId string `pulumi:"applicationId"`
	// The configuration profile ID.
	ConfigurationProfileId string `pulumi:"configurationProfileId"`
	// Current version number of hosted configuration version.
	VersionNumber string `pulumi:"versionNumber"`
}

type LookupHostedConfigurationVersionResult struct {
	// Current version number of hosted configuration version.
	VersionNumber *string `pulumi:"versionNumber"`
}

func LookupHostedConfigurationVersionOutput(ctx *pulumi.Context, args LookupHostedConfigurationVersionOutputArgs, opts ...pulumi.InvokeOption) LookupHostedConfigurationVersionResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupHostedConfigurationVersionResultOutput, error) {
			args := v.(LookupHostedConfigurationVersionArgs)
			opts = internal.PkgInvokeDefaultOpts(opts)
			var rv LookupHostedConfigurationVersionResult
			secret, err := ctx.InvokePackageRaw("aws-native:appconfig:getHostedConfigurationVersion", args, &rv, "", opts...)
			if err != nil {
				return LookupHostedConfigurationVersionResultOutput{}, err
			}

			output := pulumi.ToOutput(rv).(LookupHostedConfigurationVersionResultOutput)
			if secret {
				return pulumi.ToSecret(output).(LookupHostedConfigurationVersionResultOutput), nil
			}
			return output, nil
		}).(LookupHostedConfigurationVersionResultOutput)
}

type LookupHostedConfigurationVersionOutputArgs struct {
	// The application ID.
	ApplicationId pulumi.StringInput `pulumi:"applicationId"`
	// The configuration profile ID.
	ConfigurationProfileId pulumi.StringInput `pulumi:"configurationProfileId"`
	// Current version number of hosted configuration version.
	VersionNumber pulumi.StringInput `pulumi:"versionNumber"`
}

func (LookupHostedConfigurationVersionOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupHostedConfigurationVersionArgs)(nil)).Elem()
}

type LookupHostedConfigurationVersionResultOutput struct{ *pulumi.OutputState }

func (LookupHostedConfigurationVersionResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupHostedConfigurationVersionResult)(nil)).Elem()
}

func (o LookupHostedConfigurationVersionResultOutput) ToLookupHostedConfigurationVersionResultOutput() LookupHostedConfigurationVersionResultOutput {
	return o
}

func (o LookupHostedConfigurationVersionResultOutput) ToLookupHostedConfigurationVersionResultOutputWithContext(ctx context.Context) LookupHostedConfigurationVersionResultOutput {
	return o
}

// Current version number of hosted configuration version.
func (o LookupHostedConfigurationVersionResultOutput) VersionNumber() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupHostedConfigurationVersionResult) *string { return v.VersionNumber }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupHostedConfigurationVersionResultOutput{})
}
