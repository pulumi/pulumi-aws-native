// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Rds
{
    public static class GetDbProxy
    {
        /// <summary>
        /// Resource schema for AWS::RDS::DBProxy
        /// </summary>
        public static Task<GetDbProxyResult> InvokeAsync(GetDbProxyArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetDbProxyResult>("aws-native:rds:getDbProxy", args ?? new GetDbProxyArgs(), options.WithDefaults());

        /// <summary>
        /// Resource schema for AWS::RDS::DBProxy
        /// </summary>
        public static Output<GetDbProxyResult> Invoke(GetDbProxyInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetDbProxyResult>("aws-native:rds:getDbProxy", args ?? new GetDbProxyInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetDbProxyArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The identifier for the proxy. This name must be unique for all proxies owned by your AWS account in the specified AWS Region.
        /// </summary>
        [Input("dbProxyName", required: true)]
        public string DbProxyName { get; set; } = null!;

        public GetDbProxyArgs()
        {
        }
        public static new GetDbProxyArgs Empty => new GetDbProxyArgs();
    }

    public sealed class GetDbProxyInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The identifier for the proxy. This name must be unique for all proxies owned by your AWS account in the specified AWS Region.
        /// </summary>
        [Input("dbProxyName", required: true)]
        public Input<string> DbProxyName { get; set; } = null!;

        public GetDbProxyInvokeArgs()
        {
        }
        public static new GetDbProxyInvokeArgs Empty => new GetDbProxyInvokeArgs();
    }


    [OutputType]
    public sealed class GetDbProxyResult
    {
        /// <summary>
        /// The authorization mechanism that the proxy uses.
        /// </summary>
        public readonly ImmutableArray<Outputs.DbProxyAuthFormat> Auth;
        /// <summary>
        /// The Amazon Resource Name (ARN) for the proxy.
        /// </summary>
        public readonly string? DbProxyArn;
        /// <summary>
        /// Whether the proxy includes detailed information about SQL statements in its logs.
        /// </summary>
        public readonly bool? DebugLogging;
        /// <summary>
        /// The endpoint that you can use to connect to the proxy. You include the endpoint value in the connection string for a database client application.
        /// </summary>
        public readonly string? Endpoint;
        /// <summary>
        /// The number of seconds that a connection to the proxy can be inactive before the proxy disconnects it.
        /// </summary>
        public readonly int? IdleClientTimeout;
        /// <summary>
        /// A Boolean parameter that specifies whether Transport Layer Security (TLS) encryption is required for connections to the proxy.
        /// </summary>
        public readonly bool? RequireTls;
        /// <summary>
        /// The Amazon Resource Name (ARN) of the IAM role that the proxy uses to access secrets in AWS Secrets Manager.
        /// </summary>
        public readonly string? RoleArn;
        /// <summary>
        /// An optional set of key-value pairs to associate arbitrary data of your choosing with the proxy.
        /// </summary>
        public readonly ImmutableArray<Outputs.DbProxyTagFormat> Tags;
        /// <summary>
        /// VPC ID to associate with the new DB proxy.
        /// </summary>
        public readonly string? VpcId;
        /// <summary>
        /// VPC security group IDs to associate with the new proxy.
        /// </summary>
        public readonly ImmutableArray<string> VpcSecurityGroupIds;

        [OutputConstructor]
        private GetDbProxyResult(
            ImmutableArray<Outputs.DbProxyAuthFormat> auth,

            string? dbProxyArn,

            bool? debugLogging,

            string? endpoint,

            int? idleClientTimeout,

            bool? requireTls,

            string? roleArn,

            ImmutableArray<Outputs.DbProxyTagFormat> tags,

            string? vpcId,

            ImmutableArray<string> vpcSecurityGroupIds)
        {
            Auth = auth;
            DbProxyArn = dbProxyArn;
            DebugLogging = debugLogging;
            Endpoint = endpoint;
            IdleClientTimeout = idleClientTimeout;
            RequireTls = requireTls;
            RoleArn = roleArn;
            Tags = tags;
            VpcId = vpcId;
            VpcSecurityGroupIds = vpcSecurityGroupIds;
        }
    }
}