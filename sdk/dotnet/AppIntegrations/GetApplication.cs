// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.AppIntegrations
{
    public static class GetApplication
    {
        /// <summary>
        /// Resource Type definition for AWS:AppIntegrations::Application
        /// </summary>
        public static Task<GetApplicationResult> InvokeAsync(GetApplicationArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetApplicationResult>("aws-native:appintegrations:getApplication", args ?? new GetApplicationArgs(), options.WithDefaults());

        /// <summary>
        /// Resource Type definition for AWS:AppIntegrations::Application
        /// </summary>
        public static Output<GetApplicationResult> Invoke(GetApplicationInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetApplicationResult>("aws-native:appintegrations:getApplication", args ?? new GetApplicationInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Resource Type definition for AWS:AppIntegrations::Application
        /// </summary>
        public static Output<GetApplicationResult> Invoke(GetApplicationInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetApplicationResult>("aws-native:appintegrations:getApplication", args ?? new GetApplicationInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetApplicationArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The Amazon Resource Name (ARN) of the application.
        /// </summary>
        [Input("applicationArn", required: true)]
        public string ApplicationArn { get; set; } = null!;

        public GetApplicationArgs()
        {
        }
        public static new GetApplicationArgs Empty => new GetApplicationArgs();
    }

    public sealed class GetApplicationInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The Amazon Resource Name (ARN) of the application.
        /// </summary>
        [Input("applicationArn", required: true)]
        public Input<string> ApplicationArn { get; set; } = null!;

        public GetApplicationInvokeArgs()
        {
        }
        public static new GetApplicationInvokeArgs Empty => new GetApplicationInvokeArgs();
    }


    [OutputType]
    public sealed class GetApplicationResult
    {
        /// <summary>
        /// The Amazon Resource Name (ARN) of the application.
        /// </summary>
        public readonly string? ApplicationArn;
        /// <summary>
        /// Application source config
        /// </summary>
        public readonly Outputs.ApplicationSourceConfigProperties? ApplicationSourceConfig;
        /// <summary>
        /// The application description.
        /// </summary>
        public readonly string? Description;
        /// <summary>
        /// The id of the application.
        /// </summary>
        public readonly string? Id;
        /// <summary>
        /// The name of the application.
        /// </summary>
        public readonly string? Name;
        /// <summary>
        /// The namespace of the application.
        /// </summary>
        public readonly string? Namespace;
        /// <summary>
        /// The configuration of events or requests that the application has access to.
        /// </summary>
        public readonly ImmutableArray<string> Permissions;
        /// <summary>
        /// The tags (keys and values) associated with the application.
        /// </summary>
        public readonly ImmutableArray<Pulumi.AwsNative.Outputs.Tag> Tags;

        [OutputConstructor]
        private GetApplicationResult(
            string? applicationArn,

            Outputs.ApplicationSourceConfigProperties? applicationSourceConfig,

            string? description,

            string? id,

            string? name,

            string? @namespace,

            ImmutableArray<string> permissions,

            ImmutableArray<Pulumi.AwsNative.Outputs.Tag> tags)
        {
            ApplicationArn = applicationArn;
            ApplicationSourceConfig = applicationSourceConfig;
            Description = description;
            Id = id;
            Name = name;
            Namespace = @namespace;
            Permissions = permissions;
            Tags = tags;
        }
    }
}
