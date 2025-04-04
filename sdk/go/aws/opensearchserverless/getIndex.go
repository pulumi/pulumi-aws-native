// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package opensearchserverless

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// An OpenSearch Serverless index resource
func LookupIndex(ctx *pulumi.Context, args *LookupIndexArgs, opts ...pulumi.InvokeOption) (*LookupIndexResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupIndexResult
	err := ctx.Invoke("aws-native:opensearchserverless:getIndex", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupIndexArgs struct {
	// The endpoint for the collection.
	CollectionEndpoint string `pulumi:"collectionEndpoint"`
	// The name of the OpenSearch Serverless index.
	IndexName string `pulumi:"indexName"`
}

type LookupIndexResult struct {
	// Index Mappings
	Mappings *MappingsProperties `pulumi:"mappings"`
	// Index settings
	Settings *IndexSettings `pulumi:"settings"`
	// The unique identifier for the index.
	Uuid *string `pulumi:"uuid"`
}

func LookupIndexOutput(ctx *pulumi.Context, args LookupIndexOutputArgs, opts ...pulumi.InvokeOption) LookupIndexResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupIndexResultOutput, error) {
			args := v.(LookupIndexArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("aws-native:opensearchserverless:getIndex", args, LookupIndexResultOutput{}, options).(LookupIndexResultOutput), nil
		}).(LookupIndexResultOutput)
}

type LookupIndexOutputArgs struct {
	// The endpoint for the collection.
	CollectionEndpoint pulumi.StringInput `pulumi:"collectionEndpoint"`
	// The name of the OpenSearch Serverless index.
	IndexName pulumi.StringInput `pulumi:"indexName"`
}

func (LookupIndexOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupIndexArgs)(nil)).Elem()
}

type LookupIndexResultOutput struct{ *pulumi.OutputState }

func (LookupIndexResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupIndexResult)(nil)).Elem()
}

func (o LookupIndexResultOutput) ToLookupIndexResultOutput() LookupIndexResultOutput {
	return o
}

func (o LookupIndexResultOutput) ToLookupIndexResultOutputWithContext(ctx context.Context) LookupIndexResultOutput {
	return o
}

// Index Mappings
func (o LookupIndexResultOutput) Mappings() MappingsPropertiesPtrOutput {
	return o.ApplyT(func(v LookupIndexResult) *MappingsProperties { return v.Mappings }).(MappingsPropertiesPtrOutput)
}

// Index settings
func (o LookupIndexResultOutput) Settings() IndexSettingsPtrOutput {
	return o.ApplyT(func(v LookupIndexResult) *IndexSettings { return v.Settings }).(IndexSettingsPtrOutput)
}

// The unique identifier for the index.
func (o LookupIndexResultOutput) Uuid() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupIndexResult) *string { return v.Uuid }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupIndexResultOutput{})
}
