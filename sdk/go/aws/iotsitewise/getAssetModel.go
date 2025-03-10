// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package iotsitewise

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws"
	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource schema for AWS::IoTSiteWise::AssetModel
func LookupAssetModel(ctx *pulumi.Context, args *LookupAssetModelArgs, opts ...pulumi.InvokeOption) (*LookupAssetModelResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupAssetModelResult
	err := ctx.Invoke("aws-native:iotsitewise:getAssetModel", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupAssetModelArgs struct {
	// The ID of the asset model.
	AssetModelId string `pulumi:"assetModelId"`
}

type LookupAssetModelResult struct {
	// The ARN of the asset model, which has the following format.
	AssetModelArn *string `pulumi:"assetModelArn"`
	// The composite asset models that are part of this asset model. Composite asset models are asset models that contain specific properties.
	AssetModelCompositeModels []AssetModelCompositeModel `pulumi:"assetModelCompositeModels"`
	// A description for the asset model.
	AssetModelDescription *string `pulumi:"assetModelDescription"`
	// The external ID of the asset model.
	AssetModelExternalId *string `pulumi:"assetModelExternalId"`
	// The hierarchy definitions of the asset model. Each hierarchy specifies an asset model whose assets can be children of any other assets created from this asset model. You can specify up to 10 hierarchies per asset model.
	AssetModelHierarchies []AssetModelHierarchy `pulumi:"assetModelHierarchies"`
	// The ID of the asset model.
	AssetModelId *string `pulumi:"assetModelId"`
	// A unique, friendly name for the asset model.
	AssetModelName *string `pulumi:"assetModelName"`
	// The property definitions of the asset model. You can specify up to 200 properties per asset model.
	AssetModelProperties []AssetModelProperty `pulumi:"assetModelProperties"`
	// A list of key-value pairs that contain metadata for the asset model.
	Tags []aws.Tag `pulumi:"tags"`
}

func LookupAssetModelOutput(ctx *pulumi.Context, args LookupAssetModelOutputArgs, opts ...pulumi.InvokeOption) LookupAssetModelResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupAssetModelResultOutput, error) {
			args := v.(LookupAssetModelArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("aws-native:iotsitewise:getAssetModel", args, LookupAssetModelResultOutput{}, options).(LookupAssetModelResultOutput), nil
		}).(LookupAssetModelResultOutput)
}

type LookupAssetModelOutputArgs struct {
	// The ID of the asset model.
	AssetModelId pulumi.StringInput `pulumi:"assetModelId"`
}

func (LookupAssetModelOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAssetModelArgs)(nil)).Elem()
}

type LookupAssetModelResultOutput struct{ *pulumi.OutputState }

func (LookupAssetModelResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAssetModelResult)(nil)).Elem()
}

func (o LookupAssetModelResultOutput) ToLookupAssetModelResultOutput() LookupAssetModelResultOutput {
	return o
}

func (o LookupAssetModelResultOutput) ToLookupAssetModelResultOutputWithContext(ctx context.Context) LookupAssetModelResultOutput {
	return o
}

// The ARN of the asset model, which has the following format.
func (o LookupAssetModelResultOutput) AssetModelArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAssetModelResult) *string { return v.AssetModelArn }).(pulumi.StringPtrOutput)
}

// The composite asset models that are part of this asset model. Composite asset models are asset models that contain specific properties.
func (o LookupAssetModelResultOutput) AssetModelCompositeModels() AssetModelCompositeModelArrayOutput {
	return o.ApplyT(func(v LookupAssetModelResult) []AssetModelCompositeModel { return v.AssetModelCompositeModels }).(AssetModelCompositeModelArrayOutput)
}

// A description for the asset model.
func (o LookupAssetModelResultOutput) AssetModelDescription() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAssetModelResult) *string { return v.AssetModelDescription }).(pulumi.StringPtrOutput)
}

// The external ID of the asset model.
func (o LookupAssetModelResultOutput) AssetModelExternalId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAssetModelResult) *string { return v.AssetModelExternalId }).(pulumi.StringPtrOutput)
}

// The hierarchy definitions of the asset model. Each hierarchy specifies an asset model whose assets can be children of any other assets created from this asset model. You can specify up to 10 hierarchies per asset model.
func (o LookupAssetModelResultOutput) AssetModelHierarchies() AssetModelHierarchyArrayOutput {
	return o.ApplyT(func(v LookupAssetModelResult) []AssetModelHierarchy { return v.AssetModelHierarchies }).(AssetModelHierarchyArrayOutput)
}

// The ID of the asset model.
func (o LookupAssetModelResultOutput) AssetModelId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAssetModelResult) *string { return v.AssetModelId }).(pulumi.StringPtrOutput)
}

// A unique, friendly name for the asset model.
func (o LookupAssetModelResultOutput) AssetModelName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAssetModelResult) *string { return v.AssetModelName }).(pulumi.StringPtrOutput)
}

// The property definitions of the asset model. You can specify up to 200 properties per asset model.
func (o LookupAssetModelResultOutput) AssetModelProperties() AssetModelPropertyArrayOutput {
	return o.ApplyT(func(v LookupAssetModelResult) []AssetModelProperty { return v.AssetModelProperties }).(AssetModelPropertyArrayOutput)
}

// A list of key-value pairs that contain metadata for the asset model.
func (o LookupAssetModelResultOutput) Tags() aws.TagArrayOutput {
	return o.ApplyT(func(v LookupAssetModelResult) []aws.Tag { return v.Tags }).(aws.TagArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupAssetModelResultOutput{})
}
