// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Greengrass
{
    public static class GetLoggerDefinition
    {
        /// <summary>
        /// Resource Type definition for AWS::Greengrass::LoggerDefinition
        /// </summary>
        public static Task<GetLoggerDefinitionResult> InvokeAsync(GetLoggerDefinitionArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetLoggerDefinitionResult>("aws-native:greengrass:getLoggerDefinition", args ?? new GetLoggerDefinitionArgs(), options.WithDefaults());

        /// <summary>
        /// Resource Type definition for AWS::Greengrass::LoggerDefinition
        /// </summary>
        public static Output<GetLoggerDefinitionResult> Invoke(GetLoggerDefinitionInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetLoggerDefinitionResult>("aws-native:greengrass:getLoggerDefinition", args ?? new GetLoggerDefinitionInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetLoggerDefinitionArgs : global::Pulumi.InvokeArgs
    {
        [Input("id", required: true)]
        public string Id { get; set; } = null!;

        public GetLoggerDefinitionArgs()
        {
        }
        public static new GetLoggerDefinitionArgs Empty => new GetLoggerDefinitionArgs();
    }

    public sealed class GetLoggerDefinitionInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("id", required: true)]
        public Input<string> Id { get; set; } = null!;

        public GetLoggerDefinitionInvokeArgs()
        {
        }
        public static new GetLoggerDefinitionInvokeArgs Empty => new GetLoggerDefinitionInvokeArgs();
    }


    [OutputType]
    public sealed class GetLoggerDefinitionResult
    {
        public readonly string? Arn;
        public readonly string? Id;
        public readonly string? LatestVersionArn;
        public readonly string? Name;
        public readonly object? Tags;

        [OutputConstructor]
        private GetLoggerDefinitionResult(
            string? arn,

            string? id,

            string? latestVersionArn,

            string? name,

            object? tags)
        {
            Arn = arn;
            Id = id;
            LatestVersionArn = latestVersionArn;
            Name = name;
            Tags = tags;
        }
    }
}