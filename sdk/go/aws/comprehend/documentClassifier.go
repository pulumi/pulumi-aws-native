// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package comprehend

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Document Classifier enables training document classifier models.
type DocumentClassifier struct {
	pulumi.CustomResourceState

	Arn                    pulumi.StringOutput                         `pulumi:"arn"`
	DataAccessRoleArn      pulumi.StringOutput                         `pulumi:"dataAccessRoleArn"`
	DocumentClassifierName pulumi.StringOutput                         `pulumi:"documentClassifierName"`
	InputDataConfig        DocumentClassifierInputDataConfigOutput     `pulumi:"inputDataConfig"`
	LanguageCode           DocumentClassifierLanguageCodeOutput        `pulumi:"languageCode"`
	Mode                   DocumentClassifierModePtrOutput             `pulumi:"mode"`
	ModelKmsKeyId          pulumi.StringPtrOutput                      `pulumi:"modelKmsKeyId"`
	ModelPolicy            pulumi.StringPtrOutput                      `pulumi:"modelPolicy"`
	OutputDataConfig       DocumentClassifierOutputDataConfigPtrOutput `pulumi:"outputDataConfig"`
	Tags                   DocumentClassifierTagArrayOutput            `pulumi:"tags"`
	VersionName            pulumi.StringPtrOutput                      `pulumi:"versionName"`
	VolumeKmsKeyId         pulumi.StringPtrOutput                      `pulumi:"volumeKmsKeyId"`
	VpcConfig              DocumentClassifierVpcConfigPtrOutput        `pulumi:"vpcConfig"`
}

// NewDocumentClassifier registers a new resource with the given unique name, arguments, and options.
func NewDocumentClassifier(ctx *pulumi.Context,
	name string, args *DocumentClassifierArgs, opts ...pulumi.ResourceOption) (*DocumentClassifier, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DataAccessRoleArn == nil {
		return nil, errors.New("invalid value for required argument 'DataAccessRoleArn'")
	}
	if args.InputDataConfig == nil {
		return nil, errors.New("invalid value for required argument 'InputDataConfig'")
	}
	if args.LanguageCode == nil {
		return nil, errors.New("invalid value for required argument 'LanguageCode'")
	}
	var resource DocumentClassifier
	err := ctx.RegisterResource("aws-native:comprehend:DocumentClassifier", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDocumentClassifier gets an existing DocumentClassifier resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDocumentClassifier(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DocumentClassifierState, opts ...pulumi.ResourceOption) (*DocumentClassifier, error) {
	var resource DocumentClassifier
	err := ctx.ReadResource("aws-native:comprehend:DocumentClassifier", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DocumentClassifier resources.
type documentClassifierState struct {
}

type DocumentClassifierState struct {
}

func (DocumentClassifierState) ElementType() reflect.Type {
	return reflect.TypeOf((*documentClassifierState)(nil)).Elem()
}

type documentClassifierArgs struct {
	DataAccessRoleArn      string                              `pulumi:"dataAccessRoleArn"`
	DocumentClassifierName *string                             `pulumi:"documentClassifierName"`
	InputDataConfig        DocumentClassifierInputDataConfig   `pulumi:"inputDataConfig"`
	LanguageCode           DocumentClassifierLanguageCode      `pulumi:"languageCode"`
	Mode                   *DocumentClassifierMode             `pulumi:"mode"`
	ModelKmsKeyId          *string                             `pulumi:"modelKmsKeyId"`
	ModelPolicy            *string                             `pulumi:"modelPolicy"`
	OutputDataConfig       *DocumentClassifierOutputDataConfig `pulumi:"outputDataConfig"`
	Tags                   []DocumentClassifierTag             `pulumi:"tags"`
	VersionName            *string                             `pulumi:"versionName"`
	VolumeKmsKeyId         *string                             `pulumi:"volumeKmsKeyId"`
	VpcConfig              *DocumentClassifierVpcConfig        `pulumi:"vpcConfig"`
}

// The set of arguments for constructing a DocumentClassifier resource.
type DocumentClassifierArgs struct {
	DataAccessRoleArn      pulumi.StringInput
	DocumentClassifierName pulumi.StringPtrInput
	InputDataConfig        DocumentClassifierInputDataConfigInput
	LanguageCode           DocumentClassifierLanguageCodeInput
	Mode                   DocumentClassifierModePtrInput
	ModelKmsKeyId          pulumi.StringPtrInput
	ModelPolicy            pulumi.StringPtrInput
	OutputDataConfig       DocumentClassifierOutputDataConfigPtrInput
	Tags                   DocumentClassifierTagArrayInput
	VersionName            pulumi.StringPtrInput
	VolumeKmsKeyId         pulumi.StringPtrInput
	VpcConfig              DocumentClassifierVpcConfigPtrInput
}

func (DocumentClassifierArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*documentClassifierArgs)(nil)).Elem()
}

type DocumentClassifierInput interface {
	pulumi.Input

	ToDocumentClassifierOutput() DocumentClassifierOutput
	ToDocumentClassifierOutputWithContext(ctx context.Context) DocumentClassifierOutput
}

func (*DocumentClassifier) ElementType() reflect.Type {
	return reflect.TypeOf((**DocumentClassifier)(nil)).Elem()
}

func (i *DocumentClassifier) ToDocumentClassifierOutput() DocumentClassifierOutput {
	return i.ToDocumentClassifierOutputWithContext(context.Background())
}

func (i *DocumentClassifier) ToDocumentClassifierOutputWithContext(ctx context.Context) DocumentClassifierOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DocumentClassifierOutput)
}

type DocumentClassifierOutput struct{ *pulumi.OutputState }

func (DocumentClassifierOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DocumentClassifier)(nil)).Elem()
}

func (o DocumentClassifierOutput) ToDocumentClassifierOutput() DocumentClassifierOutput {
	return o
}

func (o DocumentClassifierOutput) ToDocumentClassifierOutputWithContext(ctx context.Context) DocumentClassifierOutput {
	return o
}

func (o DocumentClassifierOutput) Arn() pulumi.StringOutput {
	return o.ApplyT(func(v *DocumentClassifier) pulumi.StringOutput { return v.Arn }).(pulumi.StringOutput)
}

func (o DocumentClassifierOutput) DataAccessRoleArn() pulumi.StringOutput {
	return o.ApplyT(func(v *DocumentClassifier) pulumi.StringOutput { return v.DataAccessRoleArn }).(pulumi.StringOutput)
}

func (o DocumentClassifierOutput) DocumentClassifierName() pulumi.StringOutput {
	return o.ApplyT(func(v *DocumentClassifier) pulumi.StringOutput { return v.DocumentClassifierName }).(pulumi.StringOutput)
}

func (o DocumentClassifierOutput) InputDataConfig() DocumentClassifierInputDataConfigOutput {
	return o.ApplyT(func(v *DocumentClassifier) DocumentClassifierInputDataConfigOutput { return v.InputDataConfig }).(DocumentClassifierInputDataConfigOutput)
}

func (o DocumentClassifierOutput) LanguageCode() DocumentClassifierLanguageCodeOutput {
	return o.ApplyT(func(v *DocumentClassifier) DocumentClassifierLanguageCodeOutput { return v.LanguageCode }).(DocumentClassifierLanguageCodeOutput)
}

func (o DocumentClassifierOutput) Mode() DocumentClassifierModePtrOutput {
	return o.ApplyT(func(v *DocumentClassifier) DocumentClassifierModePtrOutput { return v.Mode }).(DocumentClassifierModePtrOutput)
}

func (o DocumentClassifierOutput) ModelKmsKeyId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DocumentClassifier) pulumi.StringPtrOutput { return v.ModelKmsKeyId }).(pulumi.StringPtrOutput)
}

func (o DocumentClassifierOutput) ModelPolicy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DocumentClassifier) pulumi.StringPtrOutput { return v.ModelPolicy }).(pulumi.StringPtrOutput)
}

func (o DocumentClassifierOutput) OutputDataConfig() DocumentClassifierOutputDataConfigPtrOutput {
	return o.ApplyT(func(v *DocumentClassifier) DocumentClassifierOutputDataConfigPtrOutput { return v.OutputDataConfig }).(DocumentClassifierOutputDataConfigPtrOutput)
}

func (o DocumentClassifierOutput) Tags() DocumentClassifierTagArrayOutput {
	return o.ApplyT(func(v *DocumentClassifier) DocumentClassifierTagArrayOutput { return v.Tags }).(DocumentClassifierTagArrayOutput)
}

func (o DocumentClassifierOutput) VersionName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DocumentClassifier) pulumi.StringPtrOutput { return v.VersionName }).(pulumi.StringPtrOutput)
}

func (o DocumentClassifierOutput) VolumeKmsKeyId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DocumentClassifier) pulumi.StringPtrOutput { return v.VolumeKmsKeyId }).(pulumi.StringPtrOutput)
}

func (o DocumentClassifierOutput) VpcConfig() DocumentClassifierVpcConfigPtrOutput {
	return o.ApplyT(func(v *DocumentClassifier) DocumentClassifierVpcConfigPtrOutput { return v.VpcConfig }).(DocumentClassifierVpcConfigPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DocumentClassifierInput)(nil)).Elem(), &DocumentClassifier{})
	pulumi.RegisterOutputType(DocumentClassifierOutput{})
}