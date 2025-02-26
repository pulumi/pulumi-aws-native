// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package kendra

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-aws-native/sdk/go/aws"
	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// A Kendra index
type Index struct {
	pulumi.CustomResourceState

	// The Amazon Resource Name (ARN) of the index. For example: `arn:aws:kendra:us-west-2:111122223333:index/0123456789abcdef` .
	Arn pulumi.StringOutput `pulumi:"arn"`
	// The identifier for the index. For example: `f4aeaa10-8056-4b2c-a343-522ca0f41234` .
	AwsId pulumi.StringOutput `pulumi:"awsId"`
	// Capacity units
	CapacityUnits IndexCapacityUnitsConfigurationPtrOutput `pulumi:"capacityUnits"`
	// A description for the index
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Document metadata configurations
	DocumentMetadataConfigurations IndexDocumentMetadataConfigurationArrayOutput `pulumi:"documentMetadataConfigurations"`
	// Indicates whether the index is a Enterprise Edition index, a Developer Edition index, or a GenAI Enterprise Edition index.
	Edition IndexEditionOutput `pulumi:"edition"`
	// The name of the index.
	Name pulumi.StringOutput `pulumi:"name"`
	// An IAM role that gives Amazon Kendra permissions to access your Amazon CloudWatch logs and metrics. This is also the role used when you use the [BatchPutDocument](https://docs.aws.amazon.com/kendra/latest/dg/BatchPutDocument.html) operation to index documents from an Amazon S3 bucket.
	RoleArn pulumi.StringOutput `pulumi:"roleArn"`
	// Server side encryption configuration
	ServerSideEncryptionConfiguration IndexServerSideEncryptionConfigurationPtrOutput `pulumi:"serverSideEncryptionConfiguration"`
	// Tags for labeling the index
	Tags aws.TagArrayOutput `pulumi:"tags"`
	// The user context policy.
	//
	// ATTRIBUTE_FILTER
	//
	// - All indexed content is searchable and displayable for all users. If you want to filter search results on user context, you can use the attribute filters of `_user_id` and `_group_ids` or you can provide user and group information in `UserContext` .
	//
	// USER_TOKEN
	//
	// - Enables token-based user access control to filter search results on user context. All documents with no access control and all documents accessible to the user will be searchable and displayable.
	UserContextPolicy IndexUserContextPolicyPtrOutput `pulumi:"userContextPolicy"`
	// Defines the type of user token used for the index.
	UserTokenConfigurations IndexUserTokenConfigurationArrayOutput `pulumi:"userTokenConfigurations"`
}

// NewIndex registers a new resource with the given unique name, arguments, and options.
func NewIndex(ctx *pulumi.Context,
	name string, args *IndexArgs, opts ...pulumi.ResourceOption) (*Index, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Edition == nil {
		return nil, errors.New("invalid value for required argument 'Edition'")
	}
	if args.RoleArn == nil {
		return nil, errors.New("invalid value for required argument 'RoleArn'")
	}
	replaceOnChanges := pulumi.ReplaceOnChanges([]string{
		"edition",
		"serverSideEncryptionConfiguration",
	})
	opts = append(opts, replaceOnChanges)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Index
	err := ctx.RegisterResource("aws-native:kendra:Index", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetIndex gets an existing Index resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetIndex(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *IndexState, opts ...pulumi.ResourceOption) (*Index, error) {
	var resource Index
	err := ctx.ReadResource("aws-native:kendra:Index", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Index resources.
type indexState struct {
}

type IndexState struct {
}

func (IndexState) ElementType() reflect.Type {
	return reflect.TypeOf((*indexState)(nil)).Elem()
}

type indexArgs struct {
	// Capacity units
	CapacityUnits *IndexCapacityUnitsConfiguration `pulumi:"capacityUnits"`
	// A description for the index
	Description *string `pulumi:"description"`
	// Document metadata configurations
	DocumentMetadataConfigurations []IndexDocumentMetadataConfiguration `pulumi:"documentMetadataConfigurations"`
	// Indicates whether the index is a Enterprise Edition index, a Developer Edition index, or a GenAI Enterprise Edition index.
	Edition IndexEdition `pulumi:"edition"`
	// The name of the index.
	Name *string `pulumi:"name"`
	// An IAM role that gives Amazon Kendra permissions to access your Amazon CloudWatch logs and metrics. This is also the role used when you use the [BatchPutDocument](https://docs.aws.amazon.com/kendra/latest/dg/BatchPutDocument.html) operation to index documents from an Amazon S3 bucket.
	RoleArn string `pulumi:"roleArn"`
	// Server side encryption configuration
	ServerSideEncryptionConfiguration *IndexServerSideEncryptionConfiguration `pulumi:"serverSideEncryptionConfiguration"`
	// Tags for labeling the index
	Tags []aws.Tag `pulumi:"tags"`
	// The user context policy.
	//
	// ATTRIBUTE_FILTER
	//
	// - All indexed content is searchable and displayable for all users. If you want to filter search results on user context, you can use the attribute filters of `_user_id` and `_group_ids` or you can provide user and group information in `UserContext` .
	//
	// USER_TOKEN
	//
	// - Enables token-based user access control to filter search results on user context. All documents with no access control and all documents accessible to the user will be searchable and displayable.
	UserContextPolicy *IndexUserContextPolicy `pulumi:"userContextPolicy"`
	// Defines the type of user token used for the index.
	UserTokenConfigurations []IndexUserTokenConfiguration `pulumi:"userTokenConfigurations"`
}

// The set of arguments for constructing a Index resource.
type IndexArgs struct {
	// Capacity units
	CapacityUnits IndexCapacityUnitsConfigurationPtrInput
	// A description for the index
	Description pulumi.StringPtrInput
	// Document metadata configurations
	DocumentMetadataConfigurations IndexDocumentMetadataConfigurationArrayInput
	// Indicates whether the index is a Enterprise Edition index, a Developer Edition index, or a GenAI Enterprise Edition index.
	Edition IndexEditionInput
	// The name of the index.
	Name pulumi.StringPtrInput
	// An IAM role that gives Amazon Kendra permissions to access your Amazon CloudWatch logs and metrics. This is also the role used when you use the [BatchPutDocument](https://docs.aws.amazon.com/kendra/latest/dg/BatchPutDocument.html) operation to index documents from an Amazon S3 bucket.
	RoleArn pulumi.StringInput
	// Server side encryption configuration
	ServerSideEncryptionConfiguration IndexServerSideEncryptionConfigurationPtrInput
	// Tags for labeling the index
	Tags aws.TagArrayInput
	// The user context policy.
	//
	// ATTRIBUTE_FILTER
	//
	// - All indexed content is searchable and displayable for all users. If you want to filter search results on user context, you can use the attribute filters of `_user_id` and `_group_ids` or you can provide user and group information in `UserContext` .
	//
	// USER_TOKEN
	//
	// - Enables token-based user access control to filter search results on user context. All documents with no access control and all documents accessible to the user will be searchable and displayable.
	UserContextPolicy IndexUserContextPolicyPtrInput
	// Defines the type of user token used for the index.
	UserTokenConfigurations IndexUserTokenConfigurationArrayInput
}

func (IndexArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*indexArgs)(nil)).Elem()
}

type IndexInput interface {
	pulumi.Input

	ToIndexOutput() IndexOutput
	ToIndexOutputWithContext(ctx context.Context) IndexOutput
}

func (*Index) ElementType() reflect.Type {
	return reflect.TypeOf((**Index)(nil)).Elem()
}

func (i *Index) ToIndexOutput() IndexOutput {
	return i.ToIndexOutputWithContext(context.Background())
}

func (i *Index) ToIndexOutputWithContext(ctx context.Context) IndexOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IndexOutput)
}

type IndexOutput struct{ *pulumi.OutputState }

func (IndexOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Index)(nil)).Elem()
}

func (o IndexOutput) ToIndexOutput() IndexOutput {
	return o
}

func (o IndexOutput) ToIndexOutputWithContext(ctx context.Context) IndexOutput {
	return o
}

// The Amazon Resource Name (ARN) of the index. For example: `arn:aws:kendra:us-west-2:111122223333:index/0123456789abcdef` .
func (o IndexOutput) Arn() pulumi.StringOutput {
	return o.ApplyT(func(v *Index) pulumi.StringOutput { return v.Arn }).(pulumi.StringOutput)
}

// The identifier for the index. For example: `f4aeaa10-8056-4b2c-a343-522ca0f41234` .
func (o IndexOutput) AwsId() pulumi.StringOutput {
	return o.ApplyT(func(v *Index) pulumi.StringOutput { return v.AwsId }).(pulumi.StringOutput)
}

// Capacity units
func (o IndexOutput) CapacityUnits() IndexCapacityUnitsConfigurationPtrOutput {
	return o.ApplyT(func(v *Index) IndexCapacityUnitsConfigurationPtrOutput { return v.CapacityUnits }).(IndexCapacityUnitsConfigurationPtrOutput)
}

// A description for the index
func (o IndexOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Index) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// Document metadata configurations
func (o IndexOutput) DocumentMetadataConfigurations() IndexDocumentMetadataConfigurationArrayOutput {
	return o.ApplyT(func(v *Index) IndexDocumentMetadataConfigurationArrayOutput { return v.DocumentMetadataConfigurations }).(IndexDocumentMetadataConfigurationArrayOutput)
}

// Indicates whether the index is a Enterprise Edition index, a Developer Edition index, or a GenAI Enterprise Edition index.
func (o IndexOutput) Edition() IndexEditionOutput {
	return o.ApplyT(func(v *Index) IndexEditionOutput { return v.Edition }).(IndexEditionOutput)
}

// The name of the index.
func (o IndexOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Index) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// An IAM role that gives Amazon Kendra permissions to access your Amazon CloudWatch logs and metrics. This is also the role used when you use the [BatchPutDocument](https://docs.aws.amazon.com/kendra/latest/dg/BatchPutDocument.html) operation to index documents from an Amazon S3 bucket.
func (o IndexOutput) RoleArn() pulumi.StringOutput {
	return o.ApplyT(func(v *Index) pulumi.StringOutput { return v.RoleArn }).(pulumi.StringOutput)
}

// Server side encryption configuration
func (o IndexOutput) ServerSideEncryptionConfiguration() IndexServerSideEncryptionConfigurationPtrOutput {
	return o.ApplyT(func(v *Index) IndexServerSideEncryptionConfigurationPtrOutput {
		return v.ServerSideEncryptionConfiguration
	}).(IndexServerSideEncryptionConfigurationPtrOutput)
}

// Tags for labeling the index
func (o IndexOutput) Tags() aws.TagArrayOutput {
	return o.ApplyT(func(v *Index) aws.TagArrayOutput { return v.Tags }).(aws.TagArrayOutput)
}

// The user context policy.
//
// ATTRIBUTE_FILTER
//
// - All indexed content is searchable and displayable for all users. If you want to filter search results on user context, you can use the attribute filters of `_user_id` and `_group_ids` or you can provide user and group information in `UserContext` .
//
// USER_TOKEN
//
// - Enables token-based user access control to filter search results on user context. All documents with no access control and all documents accessible to the user will be searchable and displayable.
func (o IndexOutput) UserContextPolicy() IndexUserContextPolicyPtrOutput {
	return o.ApplyT(func(v *Index) IndexUserContextPolicyPtrOutput { return v.UserContextPolicy }).(IndexUserContextPolicyPtrOutput)
}

// Defines the type of user token used for the index.
func (o IndexOutput) UserTokenConfigurations() IndexUserTokenConfigurationArrayOutput {
	return o.ApplyT(func(v *Index) IndexUserTokenConfigurationArrayOutput { return v.UserTokenConfigurations }).(IndexUserTokenConfigurationArrayOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*IndexInput)(nil)).Elem(), &Index{})
	pulumi.RegisterOutputType(IndexOutput{})
}
