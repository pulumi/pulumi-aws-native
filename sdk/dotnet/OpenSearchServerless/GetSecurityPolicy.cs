// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.OpenSearchServerless
{
    public static class GetSecurityPolicy
    {
        /// <summary>
        /// Amazon OpenSearchServerless security policy resource
        /// </summary>
        public static Task<GetSecurityPolicyResult> InvokeAsync(GetSecurityPolicyArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetSecurityPolicyResult>("aws-native:opensearchserverless:getSecurityPolicy", args ?? new GetSecurityPolicyArgs(), options.WithDefaults());

        /// <summary>
        /// Amazon OpenSearchServerless security policy resource
        /// </summary>
        public static Output<GetSecurityPolicyResult> Invoke(GetSecurityPolicyInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetSecurityPolicyResult>("aws-native:opensearchserverless:getSecurityPolicy", args ?? new GetSecurityPolicyInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Amazon OpenSearchServerless security policy resource
        /// </summary>
        public static Output<GetSecurityPolicyResult> Invoke(GetSecurityPolicyInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetSecurityPolicyResult>("aws-native:opensearchserverless:getSecurityPolicy", args ?? new GetSecurityPolicyInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetSecurityPolicyArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the policy
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        /// <summary>
        /// The type of security policy. Can be either `encryption` or `network` .
        /// </summary>
        [Input("type", required: true)]
        public Pulumi.AwsNative.OpenSearchServerless.SecurityPolicyType Type { get; set; }

        public GetSecurityPolicyArgs()
        {
        }
        public static new GetSecurityPolicyArgs Empty => new GetSecurityPolicyArgs();
    }

    public sealed class GetSecurityPolicyInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the policy
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// The type of security policy. Can be either `encryption` or `network` .
        /// </summary>
        [Input("type", required: true)]
        public Input<Pulumi.AwsNative.OpenSearchServerless.SecurityPolicyType> Type { get; set; } = null!;

        public GetSecurityPolicyInvokeArgs()
        {
        }
        public static new GetSecurityPolicyInvokeArgs Empty => new GetSecurityPolicyInvokeArgs();
    }


    [OutputType]
    public sealed class GetSecurityPolicyResult
    {
        /// <summary>
        /// The description of the policy
        /// </summary>
        public readonly string? Description;
        /// <summary>
        /// The JSON policy document that is the content for the policy
        /// </summary>
        public readonly string? Policy;

        [OutputConstructor]
        private GetSecurityPolicyResult(
            string? description,

            string? policy)
        {
            Description = description;
            Policy = policy;
        }
    }
}
