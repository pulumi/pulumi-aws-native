// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fsx

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::FSx::FileSystem
func LookupFileSystem(ctx *pulumi.Context, args *LookupFileSystemArgs, opts ...pulumi.InvokeOption) (*LookupFileSystemResult, error) {
	var rv LookupFileSystemResult
	err := ctx.Invoke("aws-native:fsx:getFileSystem", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupFileSystemArgs struct {
	Id string `pulumi:"id"`
}

type LookupFileSystemResult struct {
	DNSName              *string                         `pulumi:"dNSName"`
	Id                   *string                         `pulumi:"id"`
	LustreConfiguration  *FileSystemLustreConfiguration  `pulumi:"lustreConfiguration"`
	LustreMountName      *string                         `pulumi:"lustreMountName"`
	OntapConfiguration   *FileSystemOntapConfiguration   `pulumi:"ontapConfiguration"`
	OpenZFSConfiguration *FileSystemOpenZFSConfiguration `pulumi:"openZFSConfiguration"`
	RootVolumeId         *string                         `pulumi:"rootVolumeId"`
	StorageCapacity      *int                            `pulumi:"storageCapacity"`
	Tags                 []FileSystemTag                 `pulumi:"tags"`
	WindowsConfiguration *FileSystemWindowsConfiguration `pulumi:"windowsConfiguration"`
}

func LookupFileSystemOutput(ctx *pulumi.Context, args LookupFileSystemOutputArgs, opts ...pulumi.InvokeOption) LookupFileSystemResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupFileSystemResult, error) {
			args := v.(LookupFileSystemArgs)
			r, err := LookupFileSystem(ctx, &args, opts...)
			var s LookupFileSystemResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupFileSystemResultOutput)
}

type LookupFileSystemOutputArgs struct {
	Id pulumi.StringInput `pulumi:"id"`
}

func (LookupFileSystemOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupFileSystemArgs)(nil)).Elem()
}

type LookupFileSystemResultOutput struct{ *pulumi.OutputState }

func (LookupFileSystemResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupFileSystemResult)(nil)).Elem()
}

func (o LookupFileSystemResultOutput) ToLookupFileSystemResultOutput() LookupFileSystemResultOutput {
	return o
}

func (o LookupFileSystemResultOutput) ToLookupFileSystemResultOutputWithContext(ctx context.Context) LookupFileSystemResultOutput {
	return o
}

func (o LookupFileSystemResultOutput) DNSName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupFileSystemResult) *string { return v.DNSName }).(pulumi.StringPtrOutput)
}

func (o LookupFileSystemResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupFileSystemResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o LookupFileSystemResultOutput) LustreConfiguration() FileSystemLustreConfigurationPtrOutput {
	return o.ApplyT(func(v LookupFileSystemResult) *FileSystemLustreConfiguration { return v.LustreConfiguration }).(FileSystemLustreConfigurationPtrOutput)
}

func (o LookupFileSystemResultOutput) LustreMountName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupFileSystemResult) *string { return v.LustreMountName }).(pulumi.StringPtrOutput)
}

func (o LookupFileSystemResultOutput) OntapConfiguration() FileSystemOntapConfigurationPtrOutput {
	return o.ApplyT(func(v LookupFileSystemResult) *FileSystemOntapConfiguration { return v.OntapConfiguration }).(FileSystemOntapConfigurationPtrOutput)
}

func (o LookupFileSystemResultOutput) OpenZFSConfiguration() FileSystemOpenZFSConfigurationPtrOutput {
	return o.ApplyT(func(v LookupFileSystemResult) *FileSystemOpenZFSConfiguration { return v.OpenZFSConfiguration }).(FileSystemOpenZFSConfigurationPtrOutput)
}

func (o LookupFileSystemResultOutput) RootVolumeId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupFileSystemResult) *string { return v.RootVolumeId }).(pulumi.StringPtrOutput)
}

func (o LookupFileSystemResultOutput) StorageCapacity() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupFileSystemResult) *int { return v.StorageCapacity }).(pulumi.IntPtrOutput)
}

func (o LookupFileSystemResultOutput) Tags() FileSystemTagArrayOutput {
	return o.ApplyT(func(v LookupFileSystemResult) []FileSystemTag { return v.Tags }).(FileSystemTagArrayOutput)
}

func (o LookupFileSystemResultOutput) WindowsConfiguration() FileSystemWindowsConfigurationPtrOutput {
	return o.ApplyT(func(v LookupFileSystemResult) *FileSystemWindowsConfiguration { return v.WindowsConfiguration }).(FileSystemWindowsConfigurationPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupFileSystemResultOutput{})
}