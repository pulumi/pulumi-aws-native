// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.SecurityLake
{
    public static class GetAwsLogSource
    {
        /// <summary>
        /// Resource Type definition for AWS::SecurityLake::AwsLogSource
        /// </summary>
        public static Task<GetAwsLogSourceResult> InvokeAsync(GetAwsLogSourceArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetAwsLogSourceResult>("aws-native:securitylake:getAwsLogSource", args ?? new GetAwsLogSourceArgs(), options.WithDefaults());

        /// <summary>
        /// Resource Type definition for AWS::SecurityLake::AwsLogSource
        /// </summary>
        public static Output<GetAwsLogSourceResult> Invoke(GetAwsLogSourceInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetAwsLogSourceResult>("aws-native:securitylake:getAwsLogSource", args ?? new GetAwsLogSourceInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Resource Type definition for AWS::SecurityLake::AwsLogSource
        /// </summary>
        public static Output<GetAwsLogSourceResult> Invoke(GetAwsLogSourceInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetAwsLogSourceResult>("aws-native:securitylake:getAwsLogSource", args ?? new GetAwsLogSourceInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetAwsLogSourceArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name for a AWS source. This must be a Regionally unique value.
        /// </summary>
        [Input("sourceName", required: true)]
        public string SourceName { get; set; } = null!;

        /// <summary>
        /// The version for a AWS source. This must be a Regionally unique value.
        /// </summary>
        [Input("sourceVersion", required: true)]
        public string SourceVersion { get; set; } = null!;

        public GetAwsLogSourceArgs()
        {
        }
        public static new GetAwsLogSourceArgs Empty => new GetAwsLogSourceArgs();
    }

    public sealed class GetAwsLogSourceInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name for a AWS source. This must be a Regionally unique value.
        /// </summary>
        [Input("sourceName", required: true)]
        public Input<string> SourceName { get; set; } = null!;

        /// <summary>
        /// The version for a AWS source. This must be a Regionally unique value.
        /// </summary>
        [Input("sourceVersion", required: true)]
        public Input<string> SourceVersion { get; set; } = null!;

        public GetAwsLogSourceInvokeArgs()
        {
        }
        public static new GetAwsLogSourceInvokeArgs Empty => new GetAwsLogSourceInvokeArgs();
    }


    [OutputType]
    public sealed class GetAwsLogSourceResult
    {
        /// <summary>
        /// AWS account where you want to collect logs from.
        /// </summary>
        public readonly ImmutableArray<string> Accounts;

        [OutputConstructor]
        private GetAwsLogSourceResult(ImmutableArray<string> accounts)
        {
            Accounts = accounts;
        }
    }
}
