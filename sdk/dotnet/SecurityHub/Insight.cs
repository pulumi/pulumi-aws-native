// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.SecurityHub
{
    /// <summary>
    /// The AWS::SecurityHub::Insight resource represents the AWS Security Hub Insight in your account. An AWS Security Hub insight is a collection of related findings.
    /// </summary>
    [AwsNativeResourceType("aws-native:securityhub:Insight")]
    public partial class Insight : global::Pulumi.CustomResource
    {
        /// <summary>
        /// One or more attributes used to filter the findings included in the insight
        /// </summary>
        [Output("filters")]
        public Output<Outputs.InsightAwsSecurityFindingFilters> Filters { get; private set; } = null!;

        /// <summary>
        /// The grouping attribute for the insight's findings
        /// </summary>
        [Output("groupByAttribute")]
        public Output<string> GroupByAttribute { get; private set; } = null!;

        /// <summary>
        /// The ARN of a Security Hub insight
        /// </summary>
        [Output("insightArn")]
        public Output<string> InsightArn { get; private set; } = null!;

        /// <summary>
        /// The name of a Security Hub insight
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;


        /// <summary>
        /// Create a Insight resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Insight(string name, InsightArgs args, CustomResourceOptions? options = null)
            : base("aws-native:securityhub:Insight", name, args ?? new InsightArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Insight(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("aws-native:securityhub:Insight", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Insight resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Insight Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new Insight(name, id, options);
        }
    }

    public sealed class InsightArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// One or more attributes used to filter the findings included in the insight
        /// </summary>
        [Input("filters", required: true)]
        public Input<Inputs.InsightAwsSecurityFindingFiltersArgs> Filters { get; set; } = null!;

        /// <summary>
        /// The grouping attribute for the insight's findings
        /// </summary>
        [Input("groupByAttribute", required: true)]
        public Input<string> GroupByAttribute { get; set; } = null!;

        /// <summary>
        /// The name of a Security Hub insight
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        public InsightArgs()
        {
        }
        public static new InsightArgs Empty => new InsightArgs();
    }
}