// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.EC2
{
    public static class GetSecurityGroup
    {
        /// <summary>
        /// Resource Type definition for AWS::EC2::SecurityGroup
        /// </summary>
        public static Task<GetSecurityGroupResult> InvokeAsync(GetSecurityGroupArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetSecurityGroupResult>("aws-native:ec2:getSecurityGroup", args ?? new GetSecurityGroupArgs(), options.WithDefaults());

        /// <summary>
        /// Resource Type definition for AWS::EC2::SecurityGroup
        /// </summary>
        public static Output<GetSecurityGroupResult> Invoke(GetSecurityGroupInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetSecurityGroupResult>("aws-native:ec2:getSecurityGroup", args ?? new GetSecurityGroupInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetSecurityGroupArgs : global::Pulumi.InvokeArgs
    {
        [Input("id", required: true)]
        public string Id { get; set; } = null!;

        public GetSecurityGroupArgs()
        {
        }
        public static new GetSecurityGroupArgs Empty => new GetSecurityGroupArgs();
    }

    public sealed class GetSecurityGroupInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("id", required: true)]
        public Input<string> Id { get; set; } = null!;

        public GetSecurityGroupInvokeArgs()
        {
        }
        public static new GetSecurityGroupInvokeArgs Empty => new GetSecurityGroupInvokeArgs();
    }


    [OutputType]
    public sealed class GetSecurityGroupResult
    {
        public readonly string? GroupId;
        public readonly string? Id;
        public readonly ImmutableArray<Outputs.SecurityGroupEgress> SecurityGroupEgress;
        public readonly ImmutableArray<Outputs.SecurityGroupIngress> SecurityGroupIngress;
        public readonly ImmutableArray<Outputs.SecurityGroupTag> Tags;

        [OutputConstructor]
        private GetSecurityGroupResult(
            string? groupId,

            string? id,

            ImmutableArray<Outputs.SecurityGroupEgress> securityGroupEgress,

            ImmutableArray<Outputs.SecurityGroupIngress> securityGroupIngress,

            ImmutableArray<Outputs.SecurityGroupTag> tags)
        {
            GroupId = groupId;
            Id = id;
            SecurityGroupEgress = securityGroupEgress;
            SecurityGroupIngress = securityGroupIngress;
            Tags = tags;
        }
    }
}