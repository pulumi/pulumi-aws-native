// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Glue
{
    public static class GetDataQualityRuleset
    {
        /// <summary>
        /// Resource Type definition for AWS::Glue::DataQualityRuleset
        /// </summary>
        public static Task<GetDataQualityRulesetResult> InvokeAsync(GetDataQualityRulesetArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetDataQualityRulesetResult>("aws-native:glue:getDataQualityRuleset", args ?? new GetDataQualityRulesetArgs(), options.WithDefaults());

        /// <summary>
        /// Resource Type definition for AWS::Glue::DataQualityRuleset
        /// </summary>
        public static Output<GetDataQualityRulesetResult> Invoke(GetDataQualityRulesetInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetDataQualityRulesetResult>("aws-native:glue:getDataQualityRuleset", args ?? new GetDataQualityRulesetInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetDataQualityRulesetArgs : global::Pulumi.InvokeArgs
    {
        [Input("id", required: true)]
        public string Id { get; set; } = null!;

        public GetDataQualityRulesetArgs()
        {
        }
        public static new GetDataQualityRulesetArgs Empty => new GetDataQualityRulesetArgs();
    }

    public sealed class GetDataQualityRulesetInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("id", required: true)]
        public Input<string> Id { get; set; } = null!;

        public GetDataQualityRulesetInvokeArgs()
        {
        }
        public static new GetDataQualityRulesetInvokeArgs Empty => new GetDataQualityRulesetInvokeArgs();
    }


    [OutputType]
    public sealed class GetDataQualityRulesetResult
    {
        public readonly string? ClientToken;
        public readonly string? Description;
        public readonly string? Id;
        public readonly string? Name;
        public readonly string? Ruleset;
        public readonly object? Tags;
        public readonly Outputs.DataQualityRulesetDataQualityTargetTable? TargetTable;

        [OutputConstructor]
        private GetDataQualityRulesetResult(
            string? clientToken,

            string? description,

            string? id,

            string? name,

            string? ruleset,

            object? tags,

            Outputs.DataQualityRulesetDataQualityTargetTable? targetTable)
        {
            ClientToken = clientToken;
            Description = description;
            Id = id;
            Name = name;
            Ruleset = ruleset;
            Tags = tags;
            TargetTable = targetTable;
        }
    }
}