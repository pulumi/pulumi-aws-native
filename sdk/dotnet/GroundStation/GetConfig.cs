// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.GroundStation
{
    public static class GetConfig
    {
        /// <summary>
        /// AWS Ground Station config resource type for CloudFormation.
        /// </summary>
        public static Task<GetConfigResult> InvokeAsync(GetConfigArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetConfigResult>("aws-native:groundstation:getConfig", args ?? new GetConfigArgs(), options.WithDefaults());

        /// <summary>
        /// AWS Ground Station config resource type for CloudFormation.
        /// </summary>
        public static Output<GetConfigResult> Invoke(GetConfigInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetConfigResult>("aws-native:groundstation:getConfig", args ?? new GetConfigInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetConfigArgs : global::Pulumi.InvokeArgs
    {
        [Input("arn", required: true)]
        public string Arn { get; set; } = null!;

        public GetConfigArgs()
        {
        }
        public static new GetConfigArgs Empty => new GetConfigArgs();
    }

    public sealed class GetConfigInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("arn", required: true)]
        public Input<string> Arn { get; set; } = null!;

        public GetConfigInvokeArgs()
        {
        }
        public static new GetConfigInvokeArgs Empty => new GetConfigInvokeArgs();
    }


    [OutputType]
    public sealed class GetConfigResult
    {
        public readonly string? Arn;
        public readonly Outputs.ConfigData? ConfigData;
        public readonly string? Id;
        public readonly string? Name;
        public readonly ImmutableArray<Outputs.ConfigTag> Tags;
        public readonly string? Type;

        [OutputConstructor]
        private GetConfigResult(
            string? arn,

            Outputs.ConfigData? configData,

            string? id,

            string? name,

            ImmutableArray<Outputs.ConfigTag> tags,

            string? type)
        {
            Arn = arn;
            ConfigData = configData;
            Id = id;
            Name = name;
            Tags = tags;
            Type = type;
        }
    }
}