// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cloudformation

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The Custom Resource Emulator allows you to use AWS CloudFormation Custom Resources directly in your Pulumi programs. It provides a way to invoke AWS Lambda functions that implement custom provisioning logic following the CloudFormation Custom Resource protocol.
//
// > **Note**: Currently, only Lambda-backed Custom Resources are supported. SNS-backed Custom Resources are not supported at this time.
//
// ## Example Usage
//
// A full example of creating a CloudFormation Custom Resource Lambda function and using it in Pulumi can be found [here](https://github.com/pulumi/pulumi-aws-native/tree/master/examples/cfn-custom-resource).
//
// ## About CloudFormation Custom Resources
//
// CloudFormation Custom Resources allow you to write custom provisioning logic for resources that aren't directly available as AWS CloudFormation resource types. Common use cases include:
//
// - Implementing complex provisioning logic
// - Performing custom validations or transformations
// - Integrating with third-party services
// - Implementing organization-specific infrastructure patterns
//
// For more information about CloudFormation Custom Resources, see [Custom Resources](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/template-custom-resources.html) in the AWS CloudFormation User Guide.
//
// ## Permissions
//
// The IAM principal used by your Pulumi program must have the following permissions:
//
// 1. `lambda:InvokeFunction` on the Lambda function specified in `serviceToken`
// 2. S3 permissions on the bucket specified in `bucketName`:
//   - `s3:PutObject`
//   - `s3:GetObject`
//   - `s3:HeadObject`
//
// ## Lambda Function Requirements
//
// The Lambda function specified in `serviceToken` must implement the CloudFormation Custom Resource lifecycle.
// For detailed information about implementing Lambda-backed Custom Resources, see [AWS Lambda-backed Custom Resources](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/template-custom-resources-lambda.html) in the AWS CloudFormation User Guide.
//
// ## Timeouts
//
// Custom Resources have a default timeout of 60 minutes, matching the CloudFormation timeout for custom resource operations. You can customize it using the [`customTimeouts`](https://www.pulumi.com/docs/iac/concepts/options/customtimeouts/) resource option.
type CustomResourceEmulator struct {
	pulumi.CustomResourceState

	// The name of the S3 bucket to use for storing the response from the Custom Resource.
	Bucket pulumi.StringOutput `pulumi:"bucket"`
	// The response data returned by invoking the Custom Resource.
	Data pulumi.MapOutput `pulumi:"data"`
	// Whether the response data contains sensitive information that should be marked as secret and not logged.
	NoEcho pulumi.BoolOutput `pulumi:"noEcho"`
	// The name or unique identifier that corresponds to the `PhysicalResourceId` included in the Custom Resource response. If no `PhysicalResourceId` is provided in the response, a random ID will be generated.
	PhysicalResourceId pulumi.StringOutput `pulumi:"physicalResourceId"`
	// The CloudFormation type of the Custom Resource provider. For example, `Custom::MyCustomResource`.
	ResourceType pulumi.StringOutput `pulumi:"resourceType"`
	// The service token, such as a Lambda function ARN, that is invoked when the resource is created, updated, or deleted.
	ServiceToken pulumi.StringOutput `pulumi:"serviceToken"`
	// A stand-in value for the CloudFormation stack ID.
	StackId pulumi.StringOutput `pulumi:"stackId"`
}

// NewCustomResourceEmulator registers a new resource with the given unique name, arguments, and options.
func NewCustomResourceEmulator(ctx *pulumi.Context,
	name string, args *CustomResourceEmulatorArgs, opts ...pulumi.ResourceOption) (*CustomResourceEmulator, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.BucketKeyPrefix == nil {
		return nil, errors.New("invalid value for required argument 'BucketKeyPrefix'")
	}
	if args.BucketName == nil {
		return nil, errors.New("invalid value for required argument 'BucketName'")
	}
	if args.CustomResourceProperties == nil {
		return nil, errors.New("invalid value for required argument 'CustomResourceProperties'")
	}
	if args.ResourceType == nil {
		return nil, errors.New("invalid value for required argument 'ResourceType'")
	}
	if args.ServiceToken == nil {
		return nil, errors.New("invalid value for required argument 'ServiceToken'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource CustomResourceEmulator
	err := ctx.RegisterResource("aws-native:cloudformation:CustomResourceEmulator", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCustomResourceEmulator gets an existing CustomResourceEmulator resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCustomResourceEmulator(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CustomResourceEmulatorState, opts ...pulumi.ResourceOption) (*CustomResourceEmulator, error) {
	var resource CustomResourceEmulator
	err := ctx.ReadResource("aws-native:cloudformation:CustomResourceEmulator", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering CustomResourceEmulator resources.
type customResourceEmulatorState struct {
}

type CustomResourceEmulatorState struct {
}

func (CustomResourceEmulatorState) ElementType() reflect.Type {
	return reflect.TypeOf((*customResourceEmulatorState)(nil)).Elem()
}

type customResourceEmulatorArgs struct {
	// The prefix to use for the bucket key when storing the response from the Custom Resource provider.
	BucketKeyPrefix string `pulumi:"bucketKeyPrefix"`
	// The name of the S3 bucket to use for storing the response from the Custom Resource.
	//
	// The IAM principal configured for the provider must have `s3:PutObject`, `s3:HeadObject` and `s3:GetObject` permissions on this bucket.
	BucketName string `pulumi:"bucketName"`
	// The properties to pass as an input to the Custom Resource.
	// The properties are passed as a map of key-value pairs whereas all primitive values (number, boolean) are converted to strings for CloudFormation interoperability.
	CustomResourceProperties map[string]interface{} `pulumi:"customResourceProperties"`
	// The CloudFormation type of the Custom Resource. For example, `Custom::MyCustomResource`.
	// This is required for CloudFormation interoperability.
	ResourceType string `pulumi:"resourceType"`
	// The service token to use for the Custom Resource. The service token is invoked when the resource is created, updated, or deleted.
	// This can be a Lambda Function ARN with optional version or alias identifiers.
	//
	// The IAM principal configured for the provider must have `lambda:InvokeFunction` permissions on this service token.
	ServiceToken string `pulumi:"serviceToken"`
	// A stand-in value for the CloudFormation stack ID. This is required for CloudFormation interoperability.
	// If not provided, the Pulumi Stack ID is used.
	StackId *string `pulumi:"stackId"`
}

// The set of arguments for constructing a CustomResourceEmulator resource.
type CustomResourceEmulatorArgs struct {
	// The prefix to use for the bucket key when storing the response from the Custom Resource provider.
	BucketKeyPrefix pulumi.StringInput
	// The name of the S3 bucket to use for storing the response from the Custom Resource.
	//
	// The IAM principal configured for the provider must have `s3:PutObject`, `s3:HeadObject` and `s3:GetObject` permissions on this bucket.
	BucketName pulumi.StringInput
	// The properties to pass as an input to the Custom Resource.
	// The properties are passed as a map of key-value pairs whereas all primitive values (number, boolean) are converted to strings for CloudFormation interoperability.
	CustomResourceProperties pulumi.MapInput
	// The CloudFormation type of the Custom Resource. For example, `Custom::MyCustomResource`.
	// This is required for CloudFormation interoperability.
	ResourceType pulumi.StringInput
	// The service token to use for the Custom Resource. The service token is invoked when the resource is created, updated, or deleted.
	// This can be a Lambda Function ARN with optional version or alias identifiers.
	//
	// The IAM principal configured for the provider must have `lambda:InvokeFunction` permissions on this service token.
	ServiceToken pulumi.StringInput
	// A stand-in value for the CloudFormation stack ID. This is required for CloudFormation interoperability.
	// If not provided, the Pulumi Stack ID is used.
	StackId pulumi.StringPtrInput
}

func (CustomResourceEmulatorArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*customResourceEmulatorArgs)(nil)).Elem()
}

type CustomResourceEmulatorInput interface {
	pulumi.Input

	ToCustomResourceEmulatorOutput() CustomResourceEmulatorOutput
	ToCustomResourceEmulatorOutputWithContext(ctx context.Context) CustomResourceEmulatorOutput
}

func (*CustomResourceEmulator) ElementType() reflect.Type {
	return reflect.TypeOf((**CustomResourceEmulator)(nil)).Elem()
}

func (i *CustomResourceEmulator) ToCustomResourceEmulatorOutput() CustomResourceEmulatorOutput {
	return i.ToCustomResourceEmulatorOutputWithContext(context.Background())
}

func (i *CustomResourceEmulator) ToCustomResourceEmulatorOutputWithContext(ctx context.Context) CustomResourceEmulatorOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CustomResourceEmulatorOutput)
}

type CustomResourceEmulatorOutput struct{ *pulumi.OutputState }

func (CustomResourceEmulatorOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CustomResourceEmulator)(nil)).Elem()
}

func (o CustomResourceEmulatorOutput) ToCustomResourceEmulatorOutput() CustomResourceEmulatorOutput {
	return o
}

func (o CustomResourceEmulatorOutput) ToCustomResourceEmulatorOutputWithContext(ctx context.Context) CustomResourceEmulatorOutput {
	return o
}

// The name of the S3 bucket to use for storing the response from the Custom Resource.
func (o CustomResourceEmulatorOutput) Bucket() pulumi.StringOutput {
	return o.ApplyT(func(v *CustomResourceEmulator) pulumi.StringOutput { return v.Bucket }).(pulumi.StringOutput)
}

// The response data returned by invoking the Custom Resource.
func (o CustomResourceEmulatorOutput) Data() pulumi.MapOutput {
	return o.ApplyT(func(v *CustomResourceEmulator) pulumi.MapOutput { return v.Data }).(pulumi.MapOutput)
}

// Whether the response data contains sensitive information that should be marked as secret and not logged.
func (o CustomResourceEmulatorOutput) NoEcho() pulumi.BoolOutput {
	return o.ApplyT(func(v *CustomResourceEmulator) pulumi.BoolOutput { return v.NoEcho }).(pulumi.BoolOutput)
}

// The name or unique identifier that corresponds to the `PhysicalResourceId` included in the Custom Resource response. If no `PhysicalResourceId` is provided in the response, a random ID will be generated.
func (o CustomResourceEmulatorOutput) PhysicalResourceId() pulumi.StringOutput {
	return o.ApplyT(func(v *CustomResourceEmulator) pulumi.StringOutput { return v.PhysicalResourceId }).(pulumi.StringOutput)
}

// The CloudFormation type of the Custom Resource provider. For example, `Custom::MyCustomResource`.
func (o CustomResourceEmulatorOutput) ResourceType() pulumi.StringOutput {
	return o.ApplyT(func(v *CustomResourceEmulator) pulumi.StringOutput { return v.ResourceType }).(pulumi.StringOutput)
}

// The service token, such as a Lambda function ARN, that is invoked when the resource is created, updated, or deleted.
func (o CustomResourceEmulatorOutput) ServiceToken() pulumi.StringOutput {
	return o.ApplyT(func(v *CustomResourceEmulator) pulumi.StringOutput { return v.ServiceToken }).(pulumi.StringOutput)
}

// A stand-in value for the CloudFormation stack ID.
func (o CustomResourceEmulatorOutput) StackId() pulumi.StringOutput {
	return o.ApplyT(func(v *CustomResourceEmulator) pulumi.StringOutput { return v.StackId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*CustomResourceEmulatorInput)(nil)).Elem(), &CustomResourceEmulator{})
	pulumi.RegisterOutputType(CustomResourceEmulatorOutput{})
}
