// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ecr

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The AWS::ECR::PublicRepository resource specifies an Amazon Elastic Container Public Registry (Amazon Public ECR) repository, where users can push and pull Docker images. For more information, see https://docs.aws.amazon.com/AmazonECR
func LookupPublicRepository(ctx *pulumi.Context, args *LookupPublicRepositoryArgs, opts ...pulumi.InvokeOption) (*LookupPublicRepositoryResult, error) {
	var rv LookupPublicRepositoryResult
	err := ctx.Invoke("aws-native:ecr:getPublicRepository", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupPublicRepositoryArgs struct {
	// The name to use for the repository. The repository name may be specified on its own (such as nginx-web-app) or it can be prepended with a namespace to group the repository into a category (such as project-a/nginx-web-app). If you don't specify a name, AWS CloudFormation generates a unique physical ID and uses that ID for the repository name. For more information, see https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-name.html.
	RepositoryName string `pulumi:"repositoryName"`
}

type LookupPublicRepositoryResult struct {
	Arn *string `pulumi:"arn"`
	// The CatalogData property type specifies Catalog data for ECR Public Repository. For information about Catalog Data, see <link>
	RepositoryCatalogData *RepositoryCatalogDataProperties `pulumi:"repositoryCatalogData"`
	// The JSON repository policy text to apply to the repository. For more information, see https://docs.aws.amazon.com/AmazonECR/latest/userguide/RepositoryPolicyExamples.html in the Amazon Elastic Container Registry User Guide.
	RepositoryPolicyText interface{} `pulumi:"repositoryPolicyText"`
	// An array of key-value pairs to apply to this resource.
	Tags []PublicRepositoryTag `pulumi:"tags"`
}

func LookupPublicRepositoryOutput(ctx *pulumi.Context, args LookupPublicRepositoryOutputArgs, opts ...pulumi.InvokeOption) LookupPublicRepositoryResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupPublicRepositoryResult, error) {
			args := v.(LookupPublicRepositoryArgs)
			r, err := LookupPublicRepository(ctx, &args, opts...)
			var s LookupPublicRepositoryResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupPublicRepositoryResultOutput)
}

type LookupPublicRepositoryOutputArgs struct {
	// The name to use for the repository. The repository name may be specified on its own (such as nginx-web-app) or it can be prepended with a namespace to group the repository into a category (such as project-a/nginx-web-app). If you don't specify a name, AWS CloudFormation generates a unique physical ID and uses that ID for the repository name. For more information, see https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-name.html.
	RepositoryName pulumi.StringInput `pulumi:"repositoryName"`
}

func (LookupPublicRepositoryOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPublicRepositoryArgs)(nil)).Elem()
}

type LookupPublicRepositoryResultOutput struct{ *pulumi.OutputState }

func (LookupPublicRepositoryResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPublicRepositoryResult)(nil)).Elem()
}

func (o LookupPublicRepositoryResultOutput) ToLookupPublicRepositoryResultOutput() LookupPublicRepositoryResultOutput {
	return o
}

func (o LookupPublicRepositoryResultOutput) ToLookupPublicRepositoryResultOutputWithContext(ctx context.Context) LookupPublicRepositoryResultOutput {
	return o
}

func (o LookupPublicRepositoryResultOutput) Arn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPublicRepositoryResult) *string { return v.Arn }).(pulumi.StringPtrOutput)
}

// The CatalogData property type specifies Catalog data for ECR Public Repository. For information about Catalog Data, see <link>
func (o LookupPublicRepositoryResultOutput) RepositoryCatalogData() RepositoryCatalogDataPropertiesPtrOutput {
	return o.ApplyT(func(v LookupPublicRepositoryResult) *RepositoryCatalogDataProperties { return v.RepositoryCatalogData }).(RepositoryCatalogDataPropertiesPtrOutput)
}

// The JSON repository policy text to apply to the repository. For more information, see https://docs.aws.amazon.com/AmazonECR/latest/userguide/RepositoryPolicyExamples.html in the Amazon Elastic Container Registry User Guide.
func (o LookupPublicRepositoryResultOutput) RepositoryPolicyText() pulumi.AnyOutput {
	return o.ApplyT(func(v LookupPublicRepositoryResult) interface{} { return v.RepositoryPolicyText }).(pulumi.AnyOutput)
}

// An array of key-value pairs to apply to this resource.
func (o LookupPublicRepositoryResultOutput) Tags() PublicRepositoryTagArrayOutput {
	return o.ApplyT(func(v LookupPublicRepositoryResult) []PublicRepositoryTag { return v.Tags }).(PublicRepositoryTagArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupPublicRepositoryResultOutput{})
}