// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.AmplifyUIBuilder
{
    public static class GetTheme
    {
        /// <summary>
        /// Definition of AWS::AmplifyUIBuilder::Theme Resource Type
        /// </summary>
        public static Task<GetThemeResult> InvokeAsync(GetThemeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetThemeResult>("aws-native:amplifyuibuilder:getTheme", args ?? new GetThemeArgs(), options.WithDefaults());

        /// <summary>
        /// Definition of AWS::AmplifyUIBuilder::Theme Resource Type
        /// </summary>
        public static Output<GetThemeResult> Invoke(GetThemeInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetThemeResult>("aws-native:amplifyuibuilder:getTheme", args ?? new GetThemeInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetThemeArgs : global::Pulumi.InvokeArgs
    {
        [Input("appId", required: true)]
        public string AppId { get; set; } = null!;

        [Input("environmentName", required: true)]
        public string EnvironmentName { get; set; } = null!;

        [Input("id", required: true)]
        public string Id { get; set; } = null!;

        public GetThemeArgs()
        {
        }
        public static new GetThemeArgs Empty => new GetThemeArgs();
    }

    public sealed class GetThemeInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("appId", required: true)]
        public Input<string> AppId { get; set; } = null!;

        [Input("environmentName", required: true)]
        public Input<string> EnvironmentName { get; set; } = null!;

        [Input("id", required: true)]
        public Input<string> Id { get; set; } = null!;

        public GetThemeInvokeArgs()
        {
        }
        public static new GetThemeInvokeArgs Empty => new GetThemeInvokeArgs();
    }


    [OutputType]
    public sealed class GetThemeResult
    {
        public readonly string? AppId;
        public readonly string? EnvironmentName;
        public readonly string? Id;
        public readonly string? Name;
        public readonly ImmutableArray<Outputs.ThemeValues> Overrides;
        public readonly ImmutableArray<Outputs.ThemeValues> Values;

        [OutputConstructor]
        private GetThemeResult(
            string? appId,

            string? environmentName,

            string? id,

            string? name,

            ImmutableArray<Outputs.ThemeValues> overrides,

            ImmutableArray<Outputs.ThemeValues> values)
        {
            AppId = appId;
            EnvironmentName = environmentName;
            Id = id;
            Name = name;
            Overrides = overrides;
            Values = values;
        }
    }
}