// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.CodePipeline.Outputs
{

    /// <summary>
    /// Represents information about a rule type.
    /// </summary>
    [OutputType]
    public sealed class PipelineRuleTypeId
    {
        /// <summary>
        /// A category for the provider type for the rule.
        /// </summary>
        public readonly string? Category;
        /// <summary>
        /// The creator of the rule being called. Only AWS is supported.
        /// </summary>
        public readonly string? Owner;
        /// <summary>
        /// The provider of the service being called by the rule.
        /// </summary>
        public readonly string? Provider;
        /// <summary>
        /// A string that describes the rule version.
        /// </summary>
        public readonly string? Version;

        [OutputConstructor]
        private PipelineRuleTypeId(
            string? category,

            string? owner,

            string? provider,

            string? version)
        {
            Category = category;
            Owner = owner;
            Provider = provider;
            Version = version;
        }
    }
}