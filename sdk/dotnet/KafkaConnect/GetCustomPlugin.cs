// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.KafkaConnect
{
    public static class GetCustomPlugin
    {
        /// <summary>
        /// An example resource schema demonstrating some basic constructs and validation rules.
        /// </summary>
        public static Task<GetCustomPluginResult> InvokeAsync(GetCustomPluginArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetCustomPluginResult>("aws-native:kafkaconnect:getCustomPlugin", args ?? new GetCustomPluginArgs(), options.WithDefaults());

        /// <summary>
        /// An example resource schema demonstrating some basic constructs and validation rules.
        /// </summary>
        public static Output<GetCustomPluginResult> Invoke(GetCustomPluginInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetCustomPluginResult>("aws-native:kafkaconnect:getCustomPlugin", args ?? new GetCustomPluginInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// An example resource schema demonstrating some basic constructs and validation rules.
        /// </summary>
        public static Output<GetCustomPluginResult> Invoke(GetCustomPluginInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetCustomPluginResult>("aws-native:kafkaconnect:getCustomPlugin", args ?? new GetCustomPluginInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetCustomPluginArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The Amazon Resource Name (ARN) of the custom plugin to use.
        /// </summary>
        [Input("customPluginArn", required: true)]
        public string CustomPluginArn { get; set; } = null!;

        public GetCustomPluginArgs()
        {
        }
        public static new GetCustomPluginArgs Empty => new GetCustomPluginArgs();
    }

    public sealed class GetCustomPluginInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The Amazon Resource Name (ARN) of the custom plugin to use.
        /// </summary>
        [Input("customPluginArn", required: true)]
        public Input<string> CustomPluginArn { get; set; } = null!;

        public GetCustomPluginInvokeArgs()
        {
        }
        public static new GetCustomPluginInvokeArgs Empty => new GetCustomPluginInvokeArgs();
    }


    [OutputType]
    public sealed class GetCustomPluginResult
    {
        /// <summary>
        /// The Amazon Resource Name (ARN) of the custom plugin to use.
        /// </summary>
        public readonly string? CustomPluginArn;
        public readonly Outputs.CustomPluginFileDescription? FileDescription;
        /// <summary>
        /// The revision of the custom plugin.
        /// </summary>
        public readonly int? Revision;
        /// <summary>
        /// An array of key-value pairs to apply to this resource.
        /// </summary>
        public readonly ImmutableArray<Pulumi.AwsNative.Outputs.Tag> Tags;

        [OutputConstructor]
        private GetCustomPluginResult(
            string? customPluginArn,

            Outputs.CustomPluginFileDescription? fileDescription,

            int? revision,

            ImmutableArray<Pulumi.AwsNative.Outputs.Tag> tags)
        {
            CustomPluginArn = customPluginArn;
            FileDescription = fileDescription;
            Revision = revision;
            Tags = tags;
        }
    }
}
