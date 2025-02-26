// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package opensearchserverless

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Amazon OpenSearchServerless collection resource
func LookupCollection(ctx *pulumi.Context, args *LookupCollectionArgs, opts ...pulumi.InvokeOption) (*LookupCollectionResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupCollectionResult
	err := ctx.Invoke("aws-native:opensearchserverless:getCollection", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupCollectionArgs struct {
	// The identifier of the collection
	Id string `pulumi:"id"`
}

type LookupCollectionResult struct {
	// The Amazon Resource Name (ARN) of the collection.
	Arn *string `pulumi:"arn"`
	// The endpoint for the collection.
	CollectionEndpoint *string `pulumi:"collectionEndpoint"`
	// The OpenSearch Dashboards endpoint for the collection.
	DashboardEndpoint *string `pulumi:"dashboardEndpoint"`
	// The description of the collection
	Description *string `pulumi:"description"`
	// The identifier of the collection
	Id *string `pulumi:"id"`
	// Indicates whether to use standby replicas for the collection. You can't update this property after the collection is already created. If you attempt to modify this property, the collection continues to use the original value.
	StandbyReplicas *CollectionStandbyReplicas `pulumi:"standbyReplicas"`
}

func LookupCollectionOutput(ctx *pulumi.Context, args LookupCollectionOutputArgs, opts ...pulumi.InvokeOption) LookupCollectionResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupCollectionResultOutput, error) {
			args := v.(LookupCollectionArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("aws-native:opensearchserverless:getCollection", args, LookupCollectionResultOutput{}, options).(LookupCollectionResultOutput), nil
		}).(LookupCollectionResultOutput)
}

type LookupCollectionOutputArgs struct {
	// The identifier of the collection
	Id pulumi.StringInput `pulumi:"id"`
}

func (LookupCollectionOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupCollectionArgs)(nil)).Elem()
}

type LookupCollectionResultOutput struct{ *pulumi.OutputState }

func (LookupCollectionResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupCollectionResult)(nil)).Elem()
}

func (o LookupCollectionResultOutput) ToLookupCollectionResultOutput() LookupCollectionResultOutput {
	return o
}

func (o LookupCollectionResultOutput) ToLookupCollectionResultOutputWithContext(ctx context.Context) LookupCollectionResultOutput {
	return o
}

// The Amazon Resource Name (ARN) of the collection.
func (o LookupCollectionResultOutput) Arn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupCollectionResult) *string { return v.Arn }).(pulumi.StringPtrOutput)
}

// The endpoint for the collection.
func (o LookupCollectionResultOutput) CollectionEndpoint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupCollectionResult) *string { return v.CollectionEndpoint }).(pulumi.StringPtrOutput)
}

// The OpenSearch Dashboards endpoint for the collection.
func (o LookupCollectionResultOutput) DashboardEndpoint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupCollectionResult) *string { return v.DashboardEndpoint }).(pulumi.StringPtrOutput)
}

// The description of the collection
func (o LookupCollectionResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupCollectionResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

// The identifier of the collection
func (o LookupCollectionResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupCollectionResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

// Indicates whether to use standby replicas for the collection. You can't update this property after the collection is already created. If you attempt to modify this property, the collection continues to use the original value.
func (o LookupCollectionResultOutput) StandbyReplicas() CollectionStandbyReplicasPtrOutput {
	return o.ApplyT(func(v LookupCollectionResult) *CollectionStandbyReplicas { return v.StandbyReplicas }).(CollectionStandbyReplicasPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupCollectionResultOutput{})
}
