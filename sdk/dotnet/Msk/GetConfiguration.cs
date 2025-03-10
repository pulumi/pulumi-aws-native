// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Msk
{
    public static class GetConfiguration
    {
        /// <summary>
        /// Resource Type definition for AWS::MSK::Configuration
        /// </summary>
        public static Task<GetConfigurationResult> InvokeAsync(GetConfigurationArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetConfigurationResult>("aws-native:msk:getConfiguration", args ?? new GetConfigurationArgs(), options.WithDefaults());

        /// <summary>
        /// Resource Type definition for AWS::MSK::Configuration
        /// </summary>
        public static Output<GetConfigurationResult> Invoke(GetConfigurationInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetConfigurationResult>("aws-native:msk:getConfiguration", args ?? new GetConfigurationInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Resource Type definition for AWS::MSK::Configuration
        /// </summary>
        public static Output<GetConfigurationResult> Invoke(GetConfigurationInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetConfigurationResult>("aws-native:msk:getConfiguration", args ?? new GetConfigurationInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetConfigurationArgs : global::Pulumi.InvokeArgs
    {
        [Input("arn", required: true)]
        public string Arn { get; set; } = null!;

        public GetConfigurationArgs()
        {
        }
        public static new GetConfigurationArgs Empty => new GetConfigurationArgs();
    }

    public sealed class GetConfigurationInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("arn", required: true)]
        public Input<string> Arn { get; set; } = null!;

        public GetConfigurationInvokeArgs()
        {
        }
        public static new GetConfigurationInvokeArgs Empty => new GetConfigurationInvokeArgs();
    }


    [OutputType]
    public sealed class GetConfigurationResult
    {
        public readonly string? Arn;
        public readonly string? Description;
        public readonly Outputs.ConfigurationLatestRevision? LatestRevision;

        [OutputConstructor]
        private GetConfigurationResult(
            string? arn,

            string? description,

            Outputs.ConfigurationLatestRevision? latestRevision)
        {
            Arn = arn;
            Description = description;
            LatestRevision = latestRevision;
        }
    }
}
