// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package securitylake

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::SecurityLake::AwsLogSource
type AwsLogSource struct {
	pulumi.CustomResourceState

	// AWS account where you want to collect logs from.
	Accounts pulumi.StringArrayOutput `pulumi:"accounts"`
	// The ARN for the data lake.
	DataLakeArn pulumi.StringOutput `pulumi:"dataLakeArn"`
	// The name for a AWS source. This must be a Regionally unique value.
	SourceName pulumi.StringOutput `pulumi:"sourceName"`
	// The version for a AWS source. This must be a Regionally unique value.
	SourceVersion pulumi.StringOutput `pulumi:"sourceVersion"`
}

// NewAwsLogSource registers a new resource with the given unique name, arguments, and options.
func NewAwsLogSource(ctx *pulumi.Context,
	name string, args *AwsLogSourceArgs, opts ...pulumi.ResourceOption) (*AwsLogSource, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DataLakeArn == nil {
		return nil, errors.New("invalid value for required argument 'DataLakeArn'")
	}
	if args.SourceVersion == nil {
		return nil, errors.New("invalid value for required argument 'SourceVersion'")
	}
	replaceOnChanges := pulumi.ReplaceOnChanges([]string{
		"dataLakeArn",
		"sourceName",
		"sourceVersion",
	})
	opts = append(opts, replaceOnChanges)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource AwsLogSource
	err := ctx.RegisterResource("aws-native:securitylake:AwsLogSource", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAwsLogSource gets an existing AwsLogSource resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAwsLogSource(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AwsLogSourceState, opts ...pulumi.ResourceOption) (*AwsLogSource, error) {
	var resource AwsLogSource
	err := ctx.ReadResource("aws-native:securitylake:AwsLogSource", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AwsLogSource resources.
type awsLogSourceState struct {
}

type AwsLogSourceState struct {
}

func (AwsLogSourceState) ElementType() reflect.Type {
	return reflect.TypeOf((*awsLogSourceState)(nil)).Elem()
}

type awsLogSourceArgs struct {
	// AWS account where you want to collect logs from.
	Accounts []string `pulumi:"accounts"`
	// The ARN for the data lake.
	DataLakeArn string `pulumi:"dataLakeArn"`
	// The name for a AWS source. This must be a Regionally unique value.
	SourceName *string `pulumi:"sourceName"`
	// The version for a AWS source. This must be a Regionally unique value.
	SourceVersion string `pulumi:"sourceVersion"`
}

// The set of arguments for constructing a AwsLogSource resource.
type AwsLogSourceArgs struct {
	// AWS account where you want to collect logs from.
	Accounts pulumi.StringArrayInput
	// The ARN for the data lake.
	DataLakeArn pulumi.StringInput
	// The name for a AWS source. This must be a Regionally unique value.
	SourceName pulumi.StringPtrInput
	// The version for a AWS source. This must be a Regionally unique value.
	SourceVersion pulumi.StringInput
}

func (AwsLogSourceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*awsLogSourceArgs)(nil)).Elem()
}

type AwsLogSourceInput interface {
	pulumi.Input

	ToAwsLogSourceOutput() AwsLogSourceOutput
	ToAwsLogSourceOutputWithContext(ctx context.Context) AwsLogSourceOutput
}

func (*AwsLogSource) ElementType() reflect.Type {
	return reflect.TypeOf((**AwsLogSource)(nil)).Elem()
}

func (i *AwsLogSource) ToAwsLogSourceOutput() AwsLogSourceOutput {
	return i.ToAwsLogSourceOutputWithContext(context.Background())
}

func (i *AwsLogSource) ToAwsLogSourceOutputWithContext(ctx context.Context) AwsLogSourceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AwsLogSourceOutput)
}

type AwsLogSourceOutput struct{ *pulumi.OutputState }

func (AwsLogSourceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AwsLogSource)(nil)).Elem()
}

func (o AwsLogSourceOutput) ToAwsLogSourceOutput() AwsLogSourceOutput {
	return o
}

func (o AwsLogSourceOutput) ToAwsLogSourceOutputWithContext(ctx context.Context) AwsLogSourceOutput {
	return o
}

// AWS account where you want to collect logs from.
func (o AwsLogSourceOutput) Accounts() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *AwsLogSource) pulumi.StringArrayOutput { return v.Accounts }).(pulumi.StringArrayOutput)
}

// The ARN for the data lake.
func (o AwsLogSourceOutput) DataLakeArn() pulumi.StringOutput {
	return o.ApplyT(func(v *AwsLogSource) pulumi.StringOutput { return v.DataLakeArn }).(pulumi.StringOutput)
}

// The name for a AWS source. This must be a Regionally unique value.
func (o AwsLogSourceOutput) SourceName() pulumi.StringOutput {
	return o.ApplyT(func(v *AwsLogSource) pulumi.StringOutput { return v.SourceName }).(pulumi.StringOutput)
}

// The version for a AWS source. This must be a Regionally unique value.
func (o AwsLogSourceOutput) SourceVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *AwsLogSource) pulumi.StringOutput { return v.SourceVersion }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AwsLogSourceInput)(nil)).Elem(), &AwsLogSource{})
	pulumi.RegisterOutputType(AwsLogSourceOutput{})
}