// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Logs
{
    public static class GetResourcePolicy
    {
        /// <summary>
        /// The resource schema for AWSLogs ResourcePolicy
        /// </summary>
        public static Task<GetResourcePolicyResult> InvokeAsync(GetResourcePolicyArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetResourcePolicyResult>("aws-native:logs:getResourcePolicy", args ?? new GetResourcePolicyArgs(), options.WithDefaults());

        /// <summary>
        /// The resource schema for AWSLogs ResourcePolicy
        /// </summary>
        public static Output<GetResourcePolicyResult> Invoke(GetResourcePolicyInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetResourcePolicyResult>("aws-native:logs:getResourcePolicy", args ?? new GetResourcePolicyInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetResourcePolicyArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// A name for resource policy
        /// </summary>
        [Input("policyName", required: true)]
        public string PolicyName { get; set; } = null!;

        public GetResourcePolicyArgs()
        {
        }
        public static new GetResourcePolicyArgs Empty => new GetResourcePolicyArgs();
    }

    public sealed class GetResourcePolicyInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// A name for resource policy
        /// </summary>
        [Input("policyName", required: true)]
        public Input<string> PolicyName { get; set; } = null!;

        public GetResourcePolicyInvokeArgs()
        {
        }
        public static new GetResourcePolicyInvokeArgs Empty => new GetResourcePolicyInvokeArgs();
    }


    [OutputType]
    public sealed class GetResourcePolicyResult
    {
        /// <summary>
        /// The policy document
        /// </summary>
        public readonly string? PolicyDocument;

        [OutputConstructor]
        private GetResourcePolicyResult(string? policyDocument)
        {
            PolicyDocument = policyDocument;
        }
    }
}