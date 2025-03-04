// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ec2

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Adds the specified outbound (egress) rule to a security group.
//
//	An outbound rule permits instances to send traffic to the specified IPv4 or IPv6 address range, the IP addresses that are specified by a prefix list, or the instances that are associated with a destination security group. For more information, see [Security group rules](https://docs.aws.amazon.com/vpc/latest/userguide/security-group-rules.html).
//	You must specify exactly one of the following destinations: an IPv4 address range, an IPv6 address range, a prefix list, or a security group.
//	You must specify a protocol for each rule (for example, TCP). If the protocol is TCP or UDP, you must also specify a port or port range. If the protocol is ICMP or ICMPv6, you must also specify the ICMP/ICMPv6 type and code. To specify all types or all codes, use -1.
//	Rule changes are propagated to instances associated with the security group as quickly as possible. However, a small delay might occur.
type SecurityGroupEgress struct {
	pulumi.CustomResourceState

	AwsId pulumi.StringOutput `pulumi:"awsId"`
	// The IPv4 address range, in CIDR format.
	//  You must specify exactly one of the following: ``CidrIp``, ``CidrIpv6``, ``DestinationPrefixListId``, or ``DestinationSecurityGroupId``.
	//  For examples of rules that you can add to security groups for specific access scenarios, see [Security group rules for different use cases](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/security-group-rules-reference.html) in the *User Guide*.
	CidrIp pulumi.StringPtrOutput `pulumi:"cidrIp"`
	// The IPv6 address range, in CIDR format.
	//  You must specify exactly one of the following: ``CidrIp``, ``CidrIpv6``, ``DestinationPrefixListId``, or ``DestinationSecurityGroupId``.
	//  For examples of rules that you can add to security groups for specific access scenarios, see [Security group rules for different use cases](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/security-group-rules-reference.html) in the *User Guide*.
	CidrIpv6 pulumi.StringPtrOutput `pulumi:"cidrIpv6"`
	// The description of an egress (outbound) security group rule.
	//  Constraints: Up to 255 characters in length. Allowed characters are a-z, A-Z, 0-9, spaces, and ._-:/()#,@[]+=;{}!$*
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The prefix list IDs for an AWS service. This is the AWS service to access through a VPC endpoint from instances associated with the security group.
	//  You must specify exactly one of the following: ``CidrIp``, ``CidrIpv6``, ``DestinationPrefixListId``, or ``DestinationSecurityGroupId``.
	DestinationPrefixListId pulumi.StringPtrOutput `pulumi:"destinationPrefixListId"`
	// The ID of the security group.
	//  You must specify exactly one of the following: ``CidrIp``, ``CidrIpv6``, ``DestinationPrefixListId``, or ``DestinationSecurityGroupId``.
	DestinationSecurityGroupId pulumi.StringPtrOutput `pulumi:"destinationSecurityGroupId"`
	// If the protocol is TCP or UDP, this is the start of the port range. If the protocol is ICMP or ICMPv6, this is the ICMP type or -1 (all ICMP types).
	FromPort pulumi.IntPtrOutput `pulumi:"fromPort"`
	// The ID of the security group. You must specify either the security group ID or the security group name in the request. For security groups in a nondefault VPC, you must specify the security group ID.
	GroupId pulumi.StringOutput `pulumi:"groupId"`
	// The IP protocol name (``tcp``, ``udp``, ``icmp``, ``icmpv6``) or number (see [Protocol Numbers](https://docs.aws.amazon.com/http://www.iana.org/assignments/protocol-numbers/protocol-numbers.xhtml)).
	//  Use ``-1`` to specify all protocols. When authorizing security group rules, specifying ``-1`` or a protocol number other than ``tcp``, ``udp``, ``icmp``, or ``icmpv6`` allows traffic on all ports, regardless of any port range you specify. For ``tcp``, ``udp``, and ``icmp``, you must specify a port range. For ``icmpv6``, the port range is optional; if you omit the port range, traffic for all types and codes is allowed.
	IpProtocol pulumi.StringOutput `pulumi:"ipProtocol"`
	// If the protocol is TCP or UDP, this is the end of the port range. If the protocol is ICMP or ICMPv6, this is the ICMP code or -1 (all ICMP codes). If the start port is -1 (all ICMP types), then the end port must be -1 (all ICMP codes).
	ToPort pulumi.IntPtrOutput `pulumi:"toPort"`
}

// NewSecurityGroupEgress registers a new resource with the given unique name, arguments, and options.
func NewSecurityGroupEgress(ctx *pulumi.Context,
	name string, args *SecurityGroupEgressArgs, opts ...pulumi.ResourceOption) (*SecurityGroupEgress, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.GroupId == nil {
		return nil, errors.New("invalid value for required argument 'GroupId'")
	}
	if args.IpProtocol == nil {
		return nil, errors.New("invalid value for required argument 'IpProtocol'")
	}
	replaceOnChanges := pulumi.ReplaceOnChanges([]string{
		"cidrIp",
		"cidrIpv6",
		"destinationPrefixListId",
		"destinationSecurityGroupId",
		"fromPort",
		"groupId",
		"ipProtocol",
		"toPort",
	})
	opts = append(opts, replaceOnChanges)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource SecurityGroupEgress
	err := ctx.RegisterResource("aws-native:ec2:SecurityGroupEgress", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSecurityGroupEgress gets an existing SecurityGroupEgress resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSecurityGroupEgress(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SecurityGroupEgressState, opts ...pulumi.ResourceOption) (*SecurityGroupEgress, error) {
	var resource SecurityGroupEgress
	err := ctx.ReadResource("aws-native:ec2:SecurityGroupEgress", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SecurityGroupEgress resources.
type securityGroupEgressState struct {
}

type SecurityGroupEgressState struct {
}

func (SecurityGroupEgressState) ElementType() reflect.Type {
	return reflect.TypeOf((*securityGroupEgressState)(nil)).Elem()
}

type securityGroupEgressArgs struct {
	// The IPv4 address range, in CIDR format.
	//  You must specify exactly one of the following: ``CidrIp``, ``CidrIpv6``, ``DestinationPrefixListId``, or ``DestinationSecurityGroupId``.
	//  For examples of rules that you can add to security groups for specific access scenarios, see [Security group rules for different use cases](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/security-group-rules-reference.html) in the *User Guide*.
	CidrIp *string `pulumi:"cidrIp"`
	// The IPv6 address range, in CIDR format.
	//  You must specify exactly one of the following: ``CidrIp``, ``CidrIpv6``, ``DestinationPrefixListId``, or ``DestinationSecurityGroupId``.
	//  For examples of rules that you can add to security groups for specific access scenarios, see [Security group rules for different use cases](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/security-group-rules-reference.html) in the *User Guide*.
	CidrIpv6 *string `pulumi:"cidrIpv6"`
	// The description of an egress (outbound) security group rule.
	//  Constraints: Up to 255 characters in length. Allowed characters are a-z, A-Z, 0-9, spaces, and ._-:/()#,@[]+=;{}!$*
	Description *string `pulumi:"description"`
	// The prefix list IDs for an AWS service. This is the AWS service to access through a VPC endpoint from instances associated with the security group.
	//  You must specify exactly one of the following: ``CidrIp``, ``CidrIpv6``, ``DestinationPrefixListId``, or ``DestinationSecurityGroupId``.
	DestinationPrefixListId *string `pulumi:"destinationPrefixListId"`
	// The ID of the security group.
	//  You must specify exactly one of the following: ``CidrIp``, ``CidrIpv6``, ``DestinationPrefixListId``, or ``DestinationSecurityGroupId``.
	DestinationSecurityGroupId *string `pulumi:"destinationSecurityGroupId"`
	// If the protocol is TCP or UDP, this is the start of the port range. If the protocol is ICMP or ICMPv6, this is the ICMP type or -1 (all ICMP types).
	FromPort *int `pulumi:"fromPort"`
	// The ID of the security group. You must specify either the security group ID or the security group name in the request. For security groups in a nondefault VPC, you must specify the security group ID.
	GroupId string `pulumi:"groupId"`
	// The IP protocol name (``tcp``, ``udp``, ``icmp``, ``icmpv6``) or number (see [Protocol Numbers](https://docs.aws.amazon.com/http://www.iana.org/assignments/protocol-numbers/protocol-numbers.xhtml)).
	//  Use ``-1`` to specify all protocols. When authorizing security group rules, specifying ``-1`` or a protocol number other than ``tcp``, ``udp``, ``icmp``, or ``icmpv6`` allows traffic on all ports, regardless of any port range you specify. For ``tcp``, ``udp``, and ``icmp``, you must specify a port range. For ``icmpv6``, the port range is optional; if you omit the port range, traffic for all types and codes is allowed.
	IpProtocol string `pulumi:"ipProtocol"`
	// If the protocol is TCP or UDP, this is the end of the port range. If the protocol is ICMP or ICMPv6, this is the ICMP code or -1 (all ICMP codes). If the start port is -1 (all ICMP types), then the end port must be -1 (all ICMP codes).
	ToPort *int `pulumi:"toPort"`
}

// The set of arguments for constructing a SecurityGroupEgress resource.
type SecurityGroupEgressArgs struct {
	// The IPv4 address range, in CIDR format.
	//  You must specify exactly one of the following: ``CidrIp``, ``CidrIpv6``, ``DestinationPrefixListId``, or ``DestinationSecurityGroupId``.
	//  For examples of rules that you can add to security groups for specific access scenarios, see [Security group rules for different use cases](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/security-group-rules-reference.html) in the *User Guide*.
	CidrIp pulumi.StringPtrInput
	// The IPv6 address range, in CIDR format.
	//  You must specify exactly one of the following: ``CidrIp``, ``CidrIpv6``, ``DestinationPrefixListId``, or ``DestinationSecurityGroupId``.
	//  For examples of rules that you can add to security groups for specific access scenarios, see [Security group rules for different use cases](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/security-group-rules-reference.html) in the *User Guide*.
	CidrIpv6 pulumi.StringPtrInput
	// The description of an egress (outbound) security group rule.
	//  Constraints: Up to 255 characters in length. Allowed characters are a-z, A-Z, 0-9, spaces, and ._-:/()#,@[]+=;{}!$*
	Description pulumi.StringPtrInput
	// The prefix list IDs for an AWS service. This is the AWS service to access through a VPC endpoint from instances associated with the security group.
	//  You must specify exactly one of the following: ``CidrIp``, ``CidrIpv6``, ``DestinationPrefixListId``, or ``DestinationSecurityGroupId``.
	DestinationPrefixListId pulumi.StringPtrInput
	// The ID of the security group.
	//  You must specify exactly one of the following: ``CidrIp``, ``CidrIpv6``, ``DestinationPrefixListId``, or ``DestinationSecurityGroupId``.
	DestinationSecurityGroupId pulumi.StringPtrInput
	// If the protocol is TCP or UDP, this is the start of the port range. If the protocol is ICMP or ICMPv6, this is the ICMP type or -1 (all ICMP types).
	FromPort pulumi.IntPtrInput
	// The ID of the security group. You must specify either the security group ID or the security group name in the request. For security groups in a nondefault VPC, you must specify the security group ID.
	GroupId pulumi.StringInput
	// The IP protocol name (``tcp``, ``udp``, ``icmp``, ``icmpv6``) or number (see [Protocol Numbers](https://docs.aws.amazon.com/http://www.iana.org/assignments/protocol-numbers/protocol-numbers.xhtml)).
	//  Use ``-1`` to specify all protocols. When authorizing security group rules, specifying ``-1`` or a protocol number other than ``tcp``, ``udp``, ``icmp``, or ``icmpv6`` allows traffic on all ports, regardless of any port range you specify. For ``tcp``, ``udp``, and ``icmp``, you must specify a port range. For ``icmpv6``, the port range is optional; if you omit the port range, traffic for all types and codes is allowed.
	IpProtocol pulumi.StringInput
	// If the protocol is TCP or UDP, this is the end of the port range. If the protocol is ICMP or ICMPv6, this is the ICMP code or -1 (all ICMP codes). If the start port is -1 (all ICMP types), then the end port must be -1 (all ICMP codes).
	ToPort pulumi.IntPtrInput
}

func (SecurityGroupEgressArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*securityGroupEgressArgs)(nil)).Elem()
}

type SecurityGroupEgressInput interface {
	pulumi.Input

	ToSecurityGroupEgressOutput() SecurityGroupEgressOutput
	ToSecurityGroupEgressOutputWithContext(ctx context.Context) SecurityGroupEgressOutput
}

func (*SecurityGroupEgress) ElementType() reflect.Type {
	return reflect.TypeOf((**SecurityGroupEgress)(nil)).Elem()
}

func (i *SecurityGroupEgress) ToSecurityGroupEgressOutput() SecurityGroupEgressOutput {
	return i.ToSecurityGroupEgressOutputWithContext(context.Background())
}

func (i *SecurityGroupEgress) ToSecurityGroupEgressOutputWithContext(ctx context.Context) SecurityGroupEgressOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecurityGroupEgressOutput)
}

type SecurityGroupEgressOutput struct{ *pulumi.OutputState }

func (SecurityGroupEgressOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SecurityGroupEgress)(nil)).Elem()
}

func (o SecurityGroupEgressOutput) ToSecurityGroupEgressOutput() SecurityGroupEgressOutput {
	return o
}

func (o SecurityGroupEgressOutput) ToSecurityGroupEgressOutputWithContext(ctx context.Context) SecurityGroupEgressOutput {
	return o
}

func (o SecurityGroupEgressOutput) AwsId() pulumi.StringOutput {
	return o.ApplyT(func(v *SecurityGroupEgress) pulumi.StringOutput { return v.AwsId }).(pulumi.StringOutput)
}

// The IPv4 address range, in CIDR format.
//
//	You must specify exactly one of the following: ``CidrIp``, ``CidrIpv6``, ``DestinationPrefixListId``, or ``DestinationSecurityGroupId``.
//	For examples of rules that you can add to security groups for specific access scenarios, see [Security group rules for different use cases](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/security-group-rules-reference.html) in the *User Guide*.
func (o SecurityGroupEgressOutput) CidrIp() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SecurityGroupEgress) pulumi.StringPtrOutput { return v.CidrIp }).(pulumi.StringPtrOutput)
}

// The IPv6 address range, in CIDR format.
//
//	You must specify exactly one of the following: ``CidrIp``, ``CidrIpv6``, ``DestinationPrefixListId``, or ``DestinationSecurityGroupId``.
//	For examples of rules that you can add to security groups for specific access scenarios, see [Security group rules for different use cases](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/security-group-rules-reference.html) in the *User Guide*.
func (o SecurityGroupEgressOutput) CidrIpv6() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SecurityGroupEgress) pulumi.StringPtrOutput { return v.CidrIpv6 }).(pulumi.StringPtrOutput)
}

// The description of an egress (outbound) security group rule.
//
//	Constraints: Up to 255 characters in length. Allowed characters are a-z, A-Z, 0-9, spaces, and ._-:/()#,@[]+=;{}!$*
func (o SecurityGroupEgressOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SecurityGroupEgress) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// The prefix list IDs for an AWS service. This is the AWS service to access through a VPC endpoint from instances associated with the security group.
//
//	You must specify exactly one of the following: ``CidrIp``, ``CidrIpv6``, ``DestinationPrefixListId``, or ``DestinationSecurityGroupId``.
func (o SecurityGroupEgressOutput) DestinationPrefixListId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SecurityGroupEgress) pulumi.StringPtrOutput { return v.DestinationPrefixListId }).(pulumi.StringPtrOutput)
}

// The ID of the security group.
//
//	You must specify exactly one of the following: ``CidrIp``, ``CidrIpv6``, ``DestinationPrefixListId``, or ``DestinationSecurityGroupId``.
func (o SecurityGroupEgressOutput) DestinationSecurityGroupId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SecurityGroupEgress) pulumi.StringPtrOutput { return v.DestinationSecurityGroupId }).(pulumi.StringPtrOutput)
}

// If the protocol is TCP or UDP, this is the start of the port range. If the protocol is ICMP or ICMPv6, this is the ICMP type or -1 (all ICMP types).
func (o SecurityGroupEgressOutput) FromPort() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *SecurityGroupEgress) pulumi.IntPtrOutput { return v.FromPort }).(pulumi.IntPtrOutput)
}

// The ID of the security group. You must specify either the security group ID or the security group name in the request. For security groups in a nondefault VPC, you must specify the security group ID.
func (o SecurityGroupEgressOutput) GroupId() pulumi.StringOutput {
	return o.ApplyT(func(v *SecurityGroupEgress) pulumi.StringOutput { return v.GroupId }).(pulumi.StringOutput)
}

// The IP protocol name (“tcp“, “udp“, “icmp“, “icmpv6“) or number (see [Protocol Numbers](https://docs.aws.amazon.com/http://www.iana.org/assignments/protocol-numbers/protocol-numbers.xhtml)).
//
//	Use ``-1`` to specify all protocols. When authorizing security group rules, specifying ``-1`` or a protocol number other than ``tcp``, ``udp``, ``icmp``, or ``icmpv6`` allows traffic on all ports, regardless of any port range you specify. For ``tcp``, ``udp``, and ``icmp``, you must specify a port range. For ``icmpv6``, the port range is optional; if you omit the port range, traffic for all types and codes is allowed.
func (o SecurityGroupEgressOutput) IpProtocol() pulumi.StringOutput {
	return o.ApplyT(func(v *SecurityGroupEgress) pulumi.StringOutput { return v.IpProtocol }).(pulumi.StringOutput)
}

// If the protocol is TCP or UDP, this is the end of the port range. If the protocol is ICMP or ICMPv6, this is the ICMP code or -1 (all ICMP codes). If the start port is -1 (all ICMP types), then the end port must be -1 (all ICMP codes).
func (o SecurityGroupEgressOutput) ToPort() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *SecurityGroupEgress) pulumi.IntPtrOutput { return v.ToPort }).(pulumi.IntPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SecurityGroupEgressInput)(nil)).Elem(), &SecurityGroupEgress{})
	pulumi.RegisterOutputType(SecurityGroupEgressOutput{})
}
