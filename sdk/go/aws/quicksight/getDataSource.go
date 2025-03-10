// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package quicksight

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws"
	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Definition of the AWS::QuickSight::DataSource Resource Type.
func LookupDataSource(ctx *pulumi.Context, args *LookupDataSourceArgs, opts ...pulumi.InvokeOption) (*LookupDataSourceResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupDataSourceResult
	err := ctx.Invoke("aws-native:quicksight:getDataSource", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupDataSourceArgs struct {
	// The AWS account ID.
	AwsAccountId string `pulumi:"awsAccountId"`
	// An ID for the data source. This ID is unique per AWS Region for each AWS account.
	DataSourceId string `pulumi:"dataSourceId"`
}

type LookupDataSourceResult struct {
	// <p>A set of alternate data source parameters that you want to share for the credentials
	//             stored with this data source. The credentials are applied in tandem with the data source
	//             parameters when you copy a data source by using a create or update request. The API
	//             operation compares the <code>DataSourceParameters</code> structure that's in the request
	//             with the structures in the <code>AlternateDataSourceParameters</code> allow list. If the
	//             structures are an exact match, the request is allowed to use the credentials from this
	//             existing data source. If the <code>AlternateDataSourceParameters</code> list is null,
	//             the <code>Credentials</code> originally used with this <code>DataSourceParameters</code>
	//             are automatically allowed.</p>
	AlternateDataSourceParameters []DataSourceParameters `pulumi:"alternateDataSourceParameters"`
	// <p>The Amazon Resource Name (ARN) of the data source.</p>
	Arn *string `pulumi:"arn"`
	// <p>The time that this data source was created.</p>
	CreatedTime *string `pulumi:"createdTime"`
	// The parameters that Amazon QuickSight uses to connect to your underlying source.
	DataSourceParameters *DataSourceParameters `pulumi:"dataSourceParameters"`
	// Error information from the last update or the creation of the data source.
	ErrorInfo *DataSourceErrorInfo `pulumi:"errorInfo"`
	// <p>The last time that this data source was updated.</p>
	LastUpdatedTime *string `pulumi:"lastUpdatedTime"`
	// A display name for the data source.
	Name *string `pulumi:"name"`
	// A list of resource permissions on the data source.
	Permissions []DataSourceResourcePermission `pulumi:"permissions"`
	// Secure Socket Layer (SSL) properties that apply when Amazon QuickSight connects to your underlying source.
	SslProperties *DataSourceSslProperties `pulumi:"sslProperties"`
	// The HTTP status of the request.
	Status *DataSourceResourceStatus `pulumi:"status"`
	// Contains a map of the key-value pairs for the resource tag or tags assigned to the data source.
	Tags []aws.Tag `pulumi:"tags"`
	// Use this parameter only when you want Amazon QuickSight to use a VPC connection when connecting to your underlying source.
	VpcConnectionProperties *DataSourceVpcConnectionProperties `pulumi:"vpcConnectionProperties"`
}

func LookupDataSourceOutput(ctx *pulumi.Context, args LookupDataSourceOutputArgs, opts ...pulumi.InvokeOption) LookupDataSourceResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupDataSourceResultOutput, error) {
			args := v.(LookupDataSourceArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("aws-native:quicksight:getDataSource", args, LookupDataSourceResultOutput{}, options).(LookupDataSourceResultOutput), nil
		}).(LookupDataSourceResultOutput)
}

type LookupDataSourceOutputArgs struct {
	// The AWS account ID.
	AwsAccountId pulumi.StringInput `pulumi:"awsAccountId"`
	// An ID for the data source. This ID is unique per AWS Region for each AWS account.
	DataSourceId pulumi.StringInput `pulumi:"dataSourceId"`
}

func (LookupDataSourceOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDataSourceArgs)(nil)).Elem()
}

type LookupDataSourceResultOutput struct{ *pulumi.OutputState }

func (LookupDataSourceResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDataSourceResult)(nil)).Elem()
}

func (o LookupDataSourceResultOutput) ToLookupDataSourceResultOutput() LookupDataSourceResultOutput {
	return o
}

func (o LookupDataSourceResultOutput) ToLookupDataSourceResultOutputWithContext(ctx context.Context) LookupDataSourceResultOutput {
	return o
}

// <p>A set of alternate data source parameters that you want to share for the credentials
//
//	stored with this data source. The credentials are applied in tandem with the data source
//	parameters when you copy a data source by using a create or update request. The API
//	operation compares the <code>DataSourceParameters</code> structure that's in the request
//	with the structures in the <code>AlternateDataSourceParameters</code> allow list. If the
//	structures are an exact match, the request is allowed to use the credentials from this
//	existing data source. If the <code>AlternateDataSourceParameters</code> list is null,
//	the <code>Credentials</code> originally used with this <code>DataSourceParameters</code>
//	are automatically allowed.</p>
func (o LookupDataSourceResultOutput) AlternateDataSourceParameters() DataSourceParametersArrayOutput {
	return o.ApplyT(func(v LookupDataSourceResult) []DataSourceParameters { return v.AlternateDataSourceParameters }).(DataSourceParametersArrayOutput)
}

// <p>The Amazon Resource Name (ARN) of the data source.</p>
func (o LookupDataSourceResultOutput) Arn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDataSourceResult) *string { return v.Arn }).(pulumi.StringPtrOutput)
}

// <p>The time that this data source was created.</p>
func (o LookupDataSourceResultOutput) CreatedTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDataSourceResult) *string { return v.CreatedTime }).(pulumi.StringPtrOutput)
}

// The parameters that Amazon QuickSight uses to connect to your underlying source.
func (o LookupDataSourceResultOutput) DataSourceParameters() DataSourceParametersPtrOutput {
	return o.ApplyT(func(v LookupDataSourceResult) *DataSourceParameters { return v.DataSourceParameters }).(DataSourceParametersPtrOutput)
}

// Error information from the last update or the creation of the data source.
func (o LookupDataSourceResultOutput) ErrorInfo() DataSourceErrorInfoPtrOutput {
	return o.ApplyT(func(v LookupDataSourceResult) *DataSourceErrorInfo { return v.ErrorInfo }).(DataSourceErrorInfoPtrOutput)
}

// <p>The last time that this data source was updated.</p>
func (o LookupDataSourceResultOutput) LastUpdatedTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDataSourceResult) *string { return v.LastUpdatedTime }).(pulumi.StringPtrOutput)
}

// A display name for the data source.
func (o LookupDataSourceResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDataSourceResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

// A list of resource permissions on the data source.
func (o LookupDataSourceResultOutput) Permissions() DataSourceResourcePermissionArrayOutput {
	return o.ApplyT(func(v LookupDataSourceResult) []DataSourceResourcePermission { return v.Permissions }).(DataSourceResourcePermissionArrayOutput)
}

// Secure Socket Layer (SSL) properties that apply when Amazon QuickSight connects to your underlying source.
func (o LookupDataSourceResultOutput) SslProperties() DataSourceSslPropertiesPtrOutput {
	return o.ApplyT(func(v LookupDataSourceResult) *DataSourceSslProperties { return v.SslProperties }).(DataSourceSslPropertiesPtrOutput)
}

// The HTTP status of the request.
func (o LookupDataSourceResultOutput) Status() DataSourceResourceStatusPtrOutput {
	return o.ApplyT(func(v LookupDataSourceResult) *DataSourceResourceStatus { return v.Status }).(DataSourceResourceStatusPtrOutput)
}

// Contains a map of the key-value pairs for the resource tag or tags assigned to the data source.
func (o LookupDataSourceResultOutput) Tags() aws.TagArrayOutput {
	return o.ApplyT(func(v LookupDataSourceResult) []aws.Tag { return v.Tags }).(aws.TagArrayOutput)
}

// Use this parameter only when you want Amazon QuickSight to use a VPC connection when connecting to your underlying source.
func (o LookupDataSourceResultOutput) VpcConnectionProperties() DataSourceVpcConnectionPropertiesPtrOutput {
	return o.ApplyT(func(v LookupDataSourceResult) *DataSourceVpcConnectionProperties { return v.VpcConnectionProperties }).(DataSourceVpcConnectionPropertiesPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupDataSourceResultOutput{})
}
