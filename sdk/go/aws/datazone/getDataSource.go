// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package datazone

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// A data source is used to import technical metadata of assets (data) from the source databases or data warehouses into Amazon DataZone.
func LookupDataSource(ctx *pulumi.Context, args *LookupDataSourceArgs, opts ...pulumi.InvokeOption) (*LookupDataSourceResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupDataSourceResult
	err := ctx.Invoke("aws-native:datazone:getDataSource", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupDataSourceArgs struct {
	// The ID of the Amazon DataZone domain where the data source is created.
	DomainId string `pulumi:"domainId"`
	// The unique identifier of the data source.
	Id string `pulumi:"id"`
}

type LookupDataSourceResult struct {
	// The unique identifier of a connection used to fetch relevant parameters from connection during Datasource run
	ConnectionId *string `pulumi:"connectionId"`
	// The timestamp of when the data source was created.
	CreatedAt *string `pulumi:"createdAt"`
	// The description of the data source.
	Description *string `pulumi:"description"`
	// The ID of the Amazon DataZone domain where the data source is created.
	DomainId *string `pulumi:"domainId"`
	// Specifies whether the data source is enabled.
	EnableSetting *DataSourceEnableSetting `pulumi:"enableSetting"`
	// The unique identifier of the Amazon DataZone environment to which the data source publishes assets.
	EnvironmentId *string `pulumi:"environmentId"`
	// The unique identifier of the data source.
	Id *string `pulumi:"id"`
	// The number of assets created by the data source during its last run.
	LastRunAssetCount *float64 `pulumi:"lastRunAssetCount"`
	// The timestamp that specifies when the data source was last run.
	LastRunAt *string `pulumi:"lastRunAt"`
	// The status of the last run of this data source.
	LastRunStatus *string `pulumi:"lastRunStatus"`
	// The name of the data source.
	Name *string `pulumi:"name"`
	// The ID of the Amazon DataZone project to which the data source is added.
	ProjectId *string `pulumi:"projectId"`
	// Specifies whether the assets that this data source creates in the inventory are to be also automatically published to the catalog.
	PublishOnImport *bool `pulumi:"publishOnImport"`
	// Specifies whether the business name generation is to be enabled for this data source.
	Recommendation *DataSourceRecommendationConfiguration `pulumi:"recommendation"`
	// The schedule of the data source runs.
	Schedule *DataSourceScheduleConfiguration `pulumi:"schedule"`
	// The status of the data source.
	Status *DataSourceStatus `pulumi:"status"`
	// The timestamp of when this data source was updated.
	UpdatedAt *string `pulumi:"updatedAt"`
}

func LookupDataSourceOutput(ctx *pulumi.Context, args LookupDataSourceOutputArgs, opts ...pulumi.InvokeOption) LookupDataSourceResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupDataSourceResultOutput, error) {
			args := v.(LookupDataSourceArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("aws-native:datazone:getDataSource", args, LookupDataSourceResultOutput{}, options).(LookupDataSourceResultOutput), nil
		}).(LookupDataSourceResultOutput)
}

type LookupDataSourceOutputArgs struct {
	// The ID of the Amazon DataZone domain where the data source is created.
	DomainId pulumi.StringInput `pulumi:"domainId"`
	// The unique identifier of the data source.
	Id pulumi.StringInput `pulumi:"id"`
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

// The unique identifier of a connection used to fetch relevant parameters from connection during Datasource run
func (o LookupDataSourceResultOutput) ConnectionId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDataSourceResult) *string { return v.ConnectionId }).(pulumi.StringPtrOutput)
}

// The timestamp of when the data source was created.
func (o LookupDataSourceResultOutput) CreatedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDataSourceResult) *string { return v.CreatedAt }).(pulumi.StringPtrOutput)
}

// The description of the data source.
func (o LookupDataSourceResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDataSourceResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

// The ID of the Amazon DataZone domain where the data source is created.
func (o LookupDataSourceResultOutput) DomainId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDataSourceResult) *string { return v.DomainId }).(pulumi.StringPtrOutput)
}

// Specifies whether the data source is enabled.
func (o LookupDataSourceResultOutput) EnableSetting() DataSourceEnableSettingPtrOutput {
	return o.ApplyT(func(v LookupDataSourceResult) *DataSourceEnableSetting { return v.EnableSetting }).(DataSourceEnableSettingPtrOutput)
}

// The unique identifier of the Amazon DataZone environment to which the data source publishes assets.
func (o LookupDataSourceResultOutput) EnvironmentId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDataSourceResult) *string { return v.EnvironmentId }).(pulumi.StringPtrOutput)
}

// The unique identifier of the data source.
func (o LookupDataSourceResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDataSourceResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

// The number of assets created by the data source during its last run.
func (o LookupDataSourceResultOutput) LastRunAssetCount() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v LookupDataSourceResult) *float64 { return v.LastRunAssetCount }).(pulumi.Float64PtrOutput)
}

// The timestamp that specifies when the data source was last run.
func (o LookupDataSourceResultOutput) LastRunAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDataSourceResult) *string { return v.LastRunAt }).(pulumi.StringPtrOutput)
}

// The status of the last run of this data source.
func (o LookupDataSourceResultOutput) LastRunStatus() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDataSourceResult) *string { return v.LastRunStatus }).(pulumi.StringPtrOutput)
}

// The name of the data source.
func (o LookupDataSourceResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDataSourceResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

// The ID of the Amazon DataZone project to which the data source is added.
func (o LookupDataSourceResultOutput) ProjectId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDataSourceResult) *string { return v.ProjectId }).(pulumi.StringPtrOutput)
}

// Specifies whether the assets that this data source creates in the inventory are to be also automatically published to the catalog.
func (o LookupDataSourceResultOutput) PublishOnImport() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupDataSourceResult) *bool { return v.PublishOnImport }).(pulumi.BoolPtrOutput)
}

// Specifies whether the business name generation is to be enabled for this data source.
func (o LookupDataSourceResultOutput) Recommendation() DataSourceRecommendationConfigurationPtrOutput {
	return o.ApplyT(func(v LookupDataSourceResult) *DataSourceRecommendationConfiguration { return v.Recommendation }).(DataSourceRecommendationConfigurationPtrOutput)
}

// The schedule of the data source runs.
func (o LookupDataSourceResultOutput) Schedule() DataSourceScheduleConfigurationPtrOutput {
	return o.ApplyT(func(v LookupDataSourceResult) *DataSourceScheduleConfiguration { return v.Schedule }).(DataSourceScheduleConfigurationPtrOutput)
}

// The status of the data source.
func (o LookupDataSourceResultOutput) Status() DataSourceStatusPtrOutput {
	return o.ApplyT(func(v LookupDataSourceResult) *DataSourceStatus { return v.Status }).(DataSourceStatusPtrOutput)
}

// The timestamp of when this data source was updated.
func (o LookupDataSourceResultOutput) UpdatedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDataSourceResult) *string { return v.UpdatedAt }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupDataSourceResultOutput{})
}
