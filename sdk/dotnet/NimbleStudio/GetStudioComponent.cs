// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.NimbleStudio
{
    public static class GetStudioComponent
    {
        /// <summary>
        /// Resource Type definition for AWS::NimbleStudio::StudioComponent
        /// </summary>
        public static Task<GetStudioComponentResult> InvokeAsync(GetStudioComponentArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetStudioComponentResult>("aws-native:nimblestudio:getStudioComponent", args ?? new GetStudioComponentArgs(), options.WithDefaults());

        /// <summary>
        /// Resource Type definition for AWS::NimbleStudio::StudioComponent
        /// </summary>
        public static Output<GetStudioComponentResult> Invoke(GetStudioComponentInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetStudioComponentResult>("aws-native:nimblestudio:getStudioComponent", args ?? new GetStudioComponentInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Resource Type definition for AWS::NimbleStudio::StudioComponent
        /// </summary>
        public static Output<GetStudioComponentResult> Invoke(GetStudioComponentInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetStudioComponentResult>("aws-native:nimblestudio:getStudioComponent", args ?? new GetStudioComponentInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetStudioComponentArgs : global::Pulumi.InvokeArgs
    {
        [Input("studioComponentId", required: true)]
        public string StudioComponentId { get; set; } = null!;

        public GetStudioComponentArgs()
        {
        }
        public static new GetStudioComponentArgs Empty => new GetStudioComponentArgs();
    }

    public sealed class GetStudioComponentInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("studioComponentId", required: true)]
        public Input<string> StudioComponentId { get; set; } = null!;

        public GetStudioComponentInvokeArgs()
        {
        }
        public static new GetStudioComponentInvokeArgs Empty => new GetStudioComponentInvokeArgs();
    }


    [OutputType]
    public sealed class GetStudioComponentResult
    {
        public readonly Outputs.StudioComponentConfiguration? Configuration;
        public readonly string? Description;
        public readonly ImmutableArray<string> Ec2SecurityGroupIds;
        public readonly ImmutableArray<Outputs.StudioComponentInitializationScript> InitializationScripts;
        public readonly string? Name;
        public readonly ImmutableArray<Outputs.StudioComponentScriptParameterKeyValue> ScriptParameters;
        public readonly string? StudioComponentId;
        public readonly string? Type;

        [OutputConstructor]
        private GetStudioComponentResult(
            Outputs.StudioComponentConfiguration? configuration,

            string? description,

            ImmutableArray<string> ec2SecurityGroupIds,

            ImmutableArray<Outputs.StudioComponentInitializationScript> initializationScripts,

            string? name,

            ImmutableArray<Outputs.StudioComponentScriptParameterKeyValue> scriptParameters,

            string? studioComponentId,

            string? type)
        {
            Configuration = configuration;
            Description = description;
            Ec2SecurityGroupIds = ec2SecurityGroupIds;
            InitializationScripts = initializationScripts;
            Name = name;
            ScriptParameters = scriptParameters;
            StudioComponentId = studioComponentId;
            Type = type;
        }
    }
}
