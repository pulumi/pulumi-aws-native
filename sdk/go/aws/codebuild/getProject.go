// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package codebuild

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::CodeBuild::Project
func LookupProject(ctx *pulumi.Context, args *LookupProjectArgs, opts ...pulumi.InvokeOption) (*LookupProjectResult, error) {
	var rv LookupProjectResult
	err := ctx.Invoke("aws-native:codebuild:getProject", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupProjectArgs struct {
	Id string `pulumi:"id"`
}

type LookupProjectResult struct {
	Arn                     *string                     `pulumi:"arn"`
	Artifacts               *ProjectArtifacts           `pulumi:"artifacts"`
	BadgeEnabled            *bool                       `pulumi:"badgeEnabled"`
	BuildBatchConfig        *ProjectBuildBatchConfig    `pulumi:"buildBatchConfig"`
	Cache                   *ProjectCache               `pulumi:"cache"`
	ConcurrentBuildLimit    *int                        `pulumi:"concurrentBuildLimit"`
	Description             *string                     `pulumi:"description"`
	EncryptionKey           *string                     `pulumi:"encryptionKey"`
	Environment             *ProjectEnvironment         `pulumi:"environment"`
	FileSystemLocations     []ProjectFileSystemLocation `pulumi:"fileSystemLocations"`
	Id                      *string                     `pulumi:"id"`
	LogsConfig              *ProjectLogsConfig          `pulumi:"logsConfig"`
	QueuedTimeoutInMinutes  *int                        `pulumi:"queuedTimeoutInMinutes"`
	ResourceAccessRole      *string                     `pulumi:"resourceAccessRole"`
	SecondaryArtifacts      []ProjectArtifacts          `pulumi:"secondaryArtifacts"`
	SecondarySourceVersions []ProjectSourceVersion      `pulumi:"secondarySourceVersions"`
	SecondarySources        []ProjectSource             `pulumi:"secondarySources"`
	ServiceRole             *string                     `pulumi:"serviceRole"`
	Source                  *ProjectSource              `pulumi:"source"`
	SourceVersion           *string                     `pulumi:"sourceVersion"`
	Tags                    []ProjectTag                `pulumi:"tags"`
	TimeoutInMinutes        *int                        `pulumi:"timeoutInMinutes"`
	Triggers                *ProjectTriggers            `pulumi:"triggers"`
	Visibility              *string                     `pulumi:"visibility"`
	VpcConfig               *ProjectVpcConfig           `pulumi:"vpcConfig"`
}

func LookupProjectOutput(ctx *pulumi.Context, args LookupProjectOutputArgs, opts ...pulumi.InvokeOption) LookupProjectResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupProjectResult, error) {
			args := v.(LookupProjectArgs)
			r, err := LookupProject(ctx, &args, opts...)
			var s LookupProjectResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupProjectResultOutput)
}

type LookupProjectOutputArgs struct {
	Id pulumi.StringInput `pulumi:"id"`
}

func (LookupProjectOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupProjectArgs)(nil)).Elem()
}

type LookupProjectResultOutput struct{ *pulumi.OutputState }

func (LookupProjectResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupProjectResult)(nil)).Elem()
}

func (o LookupProjectResultOutput) ToLookupProjectResultOutput() LookupProjectResultOutput {
	return o
}

func (o LookupProjectResultOutput) ToLookupProjectResultOutputWithContext(ctx context.Context) LookupProjectResultOutput {
	return o
}

func (o LookupProjectResultOutput) Arn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupProjectResult) *string { return v.Arn }).(pulumi.StringPtrOutput)
}

func (o LookupProjectResultOutput) Artifacts() ProjectArtifactsPtrOutput {
	return o.ApplyT(func(v LookupProjectResult) *ProjectArtifacts { return v.Artifacts }).(ProjectArtifactsPtrOutput)
}

func (o LookupProjectResultOutput) BadgeEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupProjectResult) *bool { return v.BadgeEnabled }).(pulumi.BoolPtrOutput)
}

func (o LookupProjectResultOutput) BuildBatchConfig() ProjectBuildBatchConfigPtrOutput {
	return o.ApplyT(func(v LookupProjectResult) *ProjectBuildBatchConfig { return v.BuildBatchConfig }).(ProjectBuildBatchConfigPtrOutput)
}

func (o LookupProjectResultOutput) Cache() ProjectCachePtrOutput {
	return o.ApplyT(func(v LookupProjectResult) *ProjectCache { return v.Cache }).(ProjectCachePtrOutput)
}

func (o LookupProjectResultOutput) ConcurrentBuildLimit() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupProjectResult) *int { return v.ConcurrentBuildLimit }).(pulumi.IntPtrOutput)
}

func (o LookupProjectResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupProjectResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o LookupProjectResultOutput) EncryptionKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupProjectResult) *string { return v.EncryptionKey }).(pulumi.StringPtrOutput)
}

func (o LookupProjectResultOutput) Environment() ProjectEnvironmentPtrOutput {
	return o.ApplyT(func(v LookupProjectResult) *ProjectEnvironment { return v.Environment }).(ProjectEnvironmentPtrOutput)
}

func (o LookupProjectResultOutput) FileSystemLocations() ProjectFileSystemLocationArrayOutput {
	return o.ApplyT(func(v LookupProjectResult) []ProjectFileSystemLocation { return v.FileSystemLocations }).(ProjectFileSystemLocationArrayOutput)
}

func (o LookupProjectResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupProjectResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o LookupProjectResultOutput) LogsConfig() ProjectLogsConfigPtrOutput {
	return o.ApplyT(func(v LookupProjectResult) *ProjectLogsConfig { return v.LogsConfig }).(ProjectLogsConfigPtrOutput)
}

func (o LookupProjectResultOutput) QueuedTimeoutInMinutes() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupProjectResult) *int { return v.QueuedTimeoutInMinutes }).(pulumi.IntPtrOutput)
}

func (o LookupProjectResultOutput) ResourceAccessRole() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupProjectResult) *string { return v.ResourceAccessRole }).(pulumi.StringPtrOutput)
}

func (o LookupProjectResultOutput) SecondaryArtifacts() ProjectArtifactsArrayOutput {
	return o.ApplyT(func(v LookupProjectResult) []ProjectArtifacts { return v.SecondaryArtifacts }).(ProjectArtifactsArrayOutput)
}

func (o LookupProjectResultOutput) SecondarySourceVersions() ProjectSourceVersionArrayOutput {
	return o.ApplyT(func(v LookupProjectResult) []ProjectSourceVersion { return v.SecondarySourceVersions }).(ProjectSourceVersionArrayOutput)
}

func (o LookupProjectResultOutput) SecondarySources() ProjectSourceArrayOutput {
	return o.ApplyT(func(v LookupProjectResult) []ProjectSource { return v.SecondarySources }).(ProjectSourceArrayOutput)
}

func (o LookupProjectResultOutput) ServiceRole() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupProjectResult) *string { return v.ServiceRole }).(pulumi.StringPtrOutput)
}

func (o LookupProjectResultOutput) Source() ProjectSourcePtrOutput {
	return o.ApplyT(func(v LookupProjectResult) *ProjectSource { return v.Source }).(ProjectSourcePtrOutput)
}

func (o LookupProjectResultOutput) SourceVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupProjectResult) *string { return v.SourceVersion }).(pulumi.StringPtrOutput)
}

func (o LookupProjectResultOutput) Tags() ProjectTagArrayOutput {
	return o.ApplyT(func(v LookupProjectResult) []ProjectTag { return v.Tags }).(ProjectTagArrayOutput)
}

func (o LookupProjectResultOutput) TimeoutInMinutes() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupProjectResult) *int { return v.TimeoutInMinutes }).(pulumi.IntPtrOutput)
}

func (o LookupProjectResultOutput) Triggers() ProjectTriggersPtrOutput {
	return o.ApplyT(func(v LookupProjectResult) *ProjectTriggers { return v.Triggers }).(ProjectTriggersPtrOutput)
}

func (o LookupProjectResultOutput) Visibility() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupProjectResult) *string { return v.Visibility }).(pulumi.StringPtrOutput)
}

func (o LookupProjectResultOutput) VpcConfig() ProjectVpcConfigPtrOutput {
	return o.ApplyT(func(v LookupProjectResult) *ProjectVpcConfig { return v.VpcConfig }).(ProjectVpcConfigPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupProjectResultOutput{})
}