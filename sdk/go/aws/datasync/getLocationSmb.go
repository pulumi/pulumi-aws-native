// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package datasync

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws"
	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource schema for AWS::DataSync::LocationSMB.
func LookupLocationSmb(ctx *pulumi.Context, args *LookupLocationSmbArgs, opts ...pulumi.InvokeOption) (*LookupLocationSmbResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupLocationSmbResult
	err := ctx.Invoke("aws-native:datasync:getLocationSmb", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupLocationSmbArgs struct {
	// The Amazon Resource Name (ARN) of the SMB location that is created.
	LocationArn string `pulumi:"locationArn"`
}

type LookupLocationSmbResult struct {
	// The Amazon Resource Names (ARNs) of agents to use for a Simple Message Block (SMB) location.
	AgentArns []string `pulumi:"agentArns"`
	// The authentication mode used to determine identity of user.
	AuthenticationType *LocationSmbAuthenticationType `pulumi:"authenticationType"`
	// Specifies the IPv4 addresses for the DNS servers that your SMB file server belongs to. This parameter applies only if AuthenticationType is set to KERBEROS. If you have multiple domains in your environment, configuring this parameter makes sure that DataSync connects to the right SMB file server.
	DnsIpAddresses []string `pulumi:"dnsIpAddresses"`
	// The name of the Windows domain that the SMB server belongs to.
	Domain *string `pulumi:"domain"`
	// Specifies a service principal name (SPN), which is an identity in your Kerberos realm that has permission to access the files, folders, and file metadata in your SMB file server. SPNs are case sensitive and must include a prepended cifs/. For example, an SPN might look like cifs/kerberosuser@EXAMPLE.COM. Your task execution will fail if the SPN that you provide for this parameter doesn't match exactly what's in your keytab or krb5.conf files.
	KerberosPrincipal *string `pulumi:"kerberosPrincipal"`
	// The Amazon Resource Name (ARN) of the SMB location that is created.
	LocationArn *string `pulumi:"locationArn"`
	// The URL of the SMB location that was described.
	LocationUri *string `pulumi:"locationUri"`
	// Specifies the version of the SMB protocol that DataSync uses to access your SMB file server.
	MountOptions *LocationSmbMountOptions `pulumi:"mountOptions"`
	// An array of key-value pairs to apply to this resource.
	Tags []aws.Tag `pulumi:"tags"`
	// The user who can mount the share, has the permissions to access files and folders in the SMB share.
	User *string `pulumi:"user"`
}

func LookupLocationSmbOutput(ctx *pulumi.Context, args LookupLocationSmbOutputArgs, opts ...pulumi.InvokeOption) LookupLocationSmbResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupLocationSmbResultOutput, error) {
			args := v.(LookupLocationSmbArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("aws-native:datasync:getLocationSmb", args, LookupLocationSmbResultOutput{}, options).(LookupLocationSmbResultOutput), nil
		}).(LookupLocationSmbResultOutput)
}

type LookupLocationSmbOutputArgs struct {
	// The Amazon Resource Name (ARN) of the SMB location that is created.
	LocationArn pulumi.StringInput `pulumi:"locationArn"`
}

func (LookupLocationSmbOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupLocationSmbArgs)(nil)).Elem()
}

type LookupLocationSmbResultOutput struct{ *pulumi.OutputState }

func (LookupLocationSmbResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupLocationSmbResult)(nil)).Elem()
}

func (o LookupLocationSmbResultOutput) ToLookupLocationSmbResultOutput() LookupLocationSmbResultOutput {
	return o
}

func (o LookupLocationSmbResultOutput) ToLookupLocationSmbResultOutputWithContext(ctx context.Context) LookupLocationSmbResultOutput {
	return o
}

// The Amazon Resource Names (ARNs) of agents to use for a Simple Message Block (SMB) location.
func (o LookupLocationSmbResultOutput) AgentArns() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupLocationSmbResult) []string { return v.AgentArns }).(pulumi.StringArrayOutput)
}

// The authentication mode used to determine identity of user.
func (o LookupLocationSmbResultOutput) AuthenticationType() LocationSmbAuthenticationTypePtrOutput {
	return o.ApplyT(func(v LookupLocationSmbResult) *LocationSmbAuthenticationType { return v.AuthenticationType }).(LocationSmbAuthenticationTypePtrOutput)
}

// Specifies the IPv4 addresses for the DNS servers that your SMB file server belongs to. This parameter applies only if AuthenticationType is set to KERBEROS. If you have multiple domains in your environment, configuring this parameter makes sure that DataSync connects to the right SMB file server.
func (o LookupLocationSmbResultOutput) DnsIpAddresses() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupLocationSmbResult) []string { return v.DnsIpAddresses }).(pulumi.StringArrayOutput)
}

// The name of the Windows domain that the SMB server belongs to.
func (o LookupLocationSmbResultOutput) Domain() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupLocationSmbResult) *string { return v.Domain }).(pulumi.StringPtrOutput)
}

// Specifies a service principal name (SPN), which is an identity in your Kerberos realm that has permission to access the files, folders, and file metadata in your SMB file server. SPNs are case sensitive and must include a prepended cifs/. For example, an SPN might look like cifs/kerberosuser@EXAMPLE.COM. Your task execution will fail if the SPN that you provide for this parameter doesn't match exactly what's in your keytab or krb5.conf files.
func (o LookupLocationSmbResultOutput) KerberosPrincipal() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupLocationSmbResult) *string { return v.KerberosPrincipal }).(pulumi.StringPtrOutput)
}

// The Amazon Resource Name (ARN) of the SMB location that is created.
func (o LookupLocationSmbResultOutput) LocationArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupLocationSmbResult) *string { return v.LocationArn }).(pulumi.StringPtrOutput)
}

// The URL of the SMB location that was described.
func (o LookupLocationSmbResultOutput) LocationUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupLocationSmbResult) *string { return v.LocationUri }).(pulumi.StringPtrOutput)
}

// Specifies the version of the SMB protocol that DataSync uses to access your SMB file server.
func (o LookupLocationSmbResultOutput) MountOptions() LocationSmbMountOptionsPtrOutput {
	return o.ApplyT(func(v LookupLocationSmbResult) *LocationSmbMountOptions { return v.MountOptions }).(LocationSmbMountOptionsPtrOutput)
}

// An array of key-value pairs to apply to this resource.
func (o LookupLocationSmbResultOutput) Tags() aws.TagArrayOutput {
	return o.ApplyT(func(v LookupLocationSmbResult) []aws.Tag { return v.Tags }).(aws.TagArrayOutput)
}

// The user who can mount the share, has the permissions to access files and folders in the SMB share.
func (o LookupLocationSmbResultOutput) User() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupLocationSmbResult) *string { return v.User }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupLocationSmbResultOutput{})
}
