// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Bedrock
{
    public static class GetFlowVersion
    {
        /// <summary>
        /// Definition of AWS::Bedrock::FlowVersion Resource Type
        /// </summary>
        public static Task<GetFlowVersionResult> InvokeAsync(GetFlowVersionArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetFlowVersionResult>("aws-native:bedrock:getFlowVersion", args ?? new GetFlowVersionArgs(), options.WithDefaults());

        /// <summary>
        /// Definition of AWS::Bedrock::FlowVersion Resource Type
        /// </summary>
        public static Output<GetFlowVersionResult> Invoke(GetFlowVersionInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetFlowVersionResult>("aws-native:bedrock:getFlowVersion", args ?? new GetFlowVersionInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Definition of AWS::Bedrock::FlowVersion Resource Type
        /// </summary>
        public static Output<GetFlowVersionResult> Invoke(GetFlowVersionInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetFlowVersionResult>("aws-native:bedrock:getFlowVersion", args ?? new GetFlowVersionInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetFlowVersionArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Arn representation of the Flow
        /// </summary>
        [Input("flowArn", required: true)]
        public string FlowArn { get; set; } = null!;

        /// <summary>
        /// Numerical Version.
        /// </summary>
        [Input("version", required: true)]
        public string Version { get; set; } = null!;

        public GetFlowVersionArgs()
        {
        }
        public static new GetFlowVersionArgs Empty => new GetFlowVersionArgs();
    }

    public sealed class GetFlowVersionInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Arn representation of the Flow
        /// </summary>
        [Input("flowArn", required: true)]
        public Input<string> FlowArn { get; set; } = null!;

        /// <summary>
        /// Numerical Version.
        /// </summary>
        [Input("version", required: true)]
        public Input<string> Version { get; set; } = null!;

        public GetFlowVersionInvokeArgs()
        {
        }
        public static new GetFlowVersionInvokeArgs Empty => new GetFlowVersionInvokeArgs();
    }


    [OutputType]
    public sealed class GetFlowVersionResult
    {
        /// <summary>
        /// Time Stamp.
        /// </summary>
        public readonly string? CreatedAt;
        /// <summary>
        /// A KMS key ARN
        /// </summary>
        public readonly string? CustomerEncryptionKeyArn;
        public readonly Outputs.FlowVersionFlowDefinition? Definition;
        /// <summary>
        /// ARN of a IAM role
        /// </summary>
        public readonly string? ExecutionRoleArn;
        /// <summary>
        /// Identifier for a Flow
        /// </summary>
        public readonly string? FlowId;
        /// <summary>
        /// Name for the flow
        /// </summary>
        public readonly string? Name;
        /// <summary>
        /// The status of the flow.
        /// </summary>
        public readonly Pulumi.AwsNative.Bedrock.FlowVersionFlowStatus? Status;
        /// <summary>
        /// Numerical Version.
        /// </summary>
        public readonly string? Version;

        [OutputConstructor]
        private GetFlowVersionResult(
            string? createdAt,

            string? customerEncryptionKeyArn,

            Outputs.FlowVersionFlowDefinition? definition,

            string? executionRoleArn,

            string? flowId,

            string? name,

            Pulumi.AwsNative.Bedrock.FlowVersionFlowStatus? status,

            string? version)
        {
            CreatedAt = createdAt;
            CustomerEncryptionKeyArn = customerEncryptionKeyArn;
            Definition = definition;
            ExecutionRoleArn = executionRoleArn;
            FlowId = flowId;
            Name = name;
            Status = status;
            Version = version;
        }
    }
}
