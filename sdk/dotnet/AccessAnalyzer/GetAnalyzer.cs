// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.AccessAnalyzer
{
    public static class GetAnalyzer
    {
        /// <summary>
        /// The AWS::AccessAnalyzer::Analyzer type specifies an analyzer of the user's account
        /// </summary>
        public static Task<GetAnalyzerResult> InvokeAsync(GetAnalyzerArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetAnalyzerResult>("aws-native:accessanalyzer:getAnalyzer", args ?? new GetAnalyzerArgs(), options.WithDefaults());

        /// <summary>
        /// The AWS::AccessAnalyzer::Analyzer type specifies an analyzer of the user's account
        /// </summary>
        public static Output<GetAnalyzerResult> Invoke(GetAnalyzerInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetAnalyzerResult>("aws-native:accessanalyzer:getAnalyzer", args ?? new GetAnalyzerInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// The AWS::AccessAnalyzer::Analyzer type specifies an analyzer of the user's account
        /// </summary>
        public static Output<GetAnalyzerResult> Invoke(GetAnalyzerInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetAnalyzerResult>("aws-native:accessanalyzer:getAnalyzer", args ?? new GetAnalyzerInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetAnalyzerArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Amazon Resource Name (ARN) of the analyzer
        /// </summary>
        [Input("arn", required: true)]
        public string Arn { get; set; } = null!;

        public GetAnalyzerArgs()
        {
        }
        public static new GetAnalyzerArgs Empty => new GetAnalyzerArgs();
    }

    public sealed class GetAnalyzerInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Amazon Resource Name (ARN) of the analyzer
        /// </summary>
        [Input("arn", required: true)]
        public Input<string> Arn { get; set; } = null!;

        public GetAnalyzerInvokeArgs()
        {
        }
        public static new GetAnalyzerInvokeArgs Empty => new GetAnalyzerInvokeArgs();
    }


    [OutputType]
    public sealed class GetAnalyzerResult
    {
        /// <summary>
        /// The configuration for the analyzer
        /// </summary>
        public readonly Outputs.AnalyzerConfigurationProperties? AnalyzerConfiguration;
        /// <summary>
        /// Specifies the archive rules to add for the analyzer. Archive rules automatically archive findings that meet the criteria you define for the rule.
        /// </summary>
        public readonly ImmutableArray<Outputs.AnalyzerArchiveRule> ArchiveRules;
        /// <summary>
        /// Amazon Resource Name (ARN) of the analyzer
        /// </summary>
        public readonly string? Arn;
        /// <summary>
        /// An array of key-value pairs to apply to this resource.
        /// </summary>
        public readonly ImmutableArray<Pulumi.AwsNative.Outputs.Tag> Tags;

        [OutputConstructor]
        private GetAnalyzerResult(
            Outputs.AnalyzerConfigurationProperties? analyzerConfiguration,

            ImmutableArray<Outputs.AnalyzerArchiveRule> archiveRules,

            string? arn,

            ImmutableArray<Pulumi.AwsNative.Outputs.Tag> tags)
        {
            AnalyzerConfiguration = analyzerConfiguration;
            ArchiveRules = archiveRules;
            Arn = arn;
            Tags = tags;
        }
    }
}
