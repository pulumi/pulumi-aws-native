// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.OpenSearchServerless
{
    public static class GetSecurityConfig
    {
        /// <summary>
        /// Amazon OpenSearchServerless security config resource
        /// </summary>
        public static Task<GetSecurityConfigResult> InvokeAsync(GetSecurityConfigArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetSecurityConfigResult>("aws-native:opensearchserverless:getSecurityConfig", args ?? new GetSecurityConfigArgs(), options.WithDefaults());

        /// <summary>
        /// Amazon OpenSearchServerless security config resource
        /// </summary>
        public static Output<GetSecurityConfigResult> Invoke(GetSecurityConfigInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetSecurityConfigResult>("aws-native:opensearchserverless:getSecurityConfig", args ?? new GetSecurityConfigInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetSecurityConfigArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The identifier of the security config
        /// </summary>
        [Input("id", required: true)]
        public string Id { get; set; } = null!;

        public GetSecurityConfigArgs()
        {
        }
        public static new GetSecurityConfigArgs Empty => new GetSecurityConfigArgs();
    }

    public sealed class GetSecurityConfigInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The identifier of the security config
        /// </summary>
        [Input("id", required: true)]
        public Input<string> Id { get; set; } = null!;

        public GetSecurityConfigInvokeArgs()
        {
        }
        public static new GetSecurityConfigInvokeArgs Empty => new GetSecurityConfigInvokeArgs();
    }


    [OutputType]
    public sealed class GetSecurityConfigResult
    {
        /// <summary>
        /// Security config description
        /// </summary>
        public readonly string? Description;
        /// <summary>
        /// The identifier of the security config
        /// </summary>
        public readonly string? Id;
        public readonly Outputs.SecurityConfigSamlConfigOptions? SamlOptions;

        [OutputConstructor]
        private GetSecurityConfigResult(
            string? description,

            string? id,

            Outputs.SecurityConfigSamlConfigOptions? samlOptions)
        {
            Description = description;
            Id = id;
            SamlOptions = samlOptions;
        }
    }
}