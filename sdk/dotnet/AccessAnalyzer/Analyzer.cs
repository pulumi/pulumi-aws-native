// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.AccessAnalyzer
{
    /// <summary>
    /// The AWS::AccessAnalyzer::Analyzer type specifies an analyzer of the user's account
    /// </summary>
    [AwsNativeResourceType("aws-native:accessanalyzer:Analyzer")]
    public partial class Analyzer : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Analyzer name
        /// </summary>
        [Output("analyzerName")]
        public Output<string?> AnalyzerName { get; private set; } = null!;

        [Output("archiveRules")]
        public Output<ImmutableArray<Outputs.AnalyzerArchiveRule>> ArchiveRules { get; private set; } = null!;

        /// <summary>
        /// Amazon Resource Name (ARN) of the analyzer
        /// </summary>
        [Output("arn")]
        public Output<string> Arn { get; private set; } = null!;

        /// <summary>
        /// An array of key-value pairs to apply to this resource.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableArray<Outputs.AnalyzerTag>> Tags { get; private set; } = null!;

        /// <summary>
        /// The type of the analyzer, must be ACCOUNT or ORGANIZATION
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;


        /// <summary>
        /// Create a Analyzer resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Analyzer(string name, AnalyzerArgs args, CustomResourceOptions? options = null)
            : base("aws-native:accessanalyzer:Analyzer", name, args ?? new AnalyzerArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Analyzer(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("aws-native:accessanalyzer:Analyzer", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing Analyzer resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Analyzer Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new Analyzer(name, id, options);
        }
    }

    public sealed class AnalyzerArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Analyzer name
        /// </summary>
        [Input("analyzerName")]
        public Input<string>? AnalyzerName { get; set; }

        [Input("archiveRules")]
        private InputList<Inputs.AnalyzerArchiveRuleArgs>? _archiveRules;
        public InputList<Inputs.AnalyzerArchiveRuleArgs> ArchiveRules
        {
            get => _archiveRules ?? (_archiveRules = new InputList<Inputs.AnalyzerArchiveRuleArgs>());
            set => _archiveRules = value;
        }

        [Input("tags")]
        private InputList<Inputs.AnalyzerTagArgs>? _tags;

        /// <summary>
        /// An array of key-value pairs to apply to this resource.
        /// </summary>
        public InputList<Inputs.AnalyzerTagArgs> Tags
        {
            get => _tags ?? (_tags = new InputList<Inputs.AnalyzerTagArgs>());
            set => _tags = value;
        }

        /// <summary>
        /// The type of the analyzer, must be ACCOUNT or ORGANIZATION
        /// </summary>
        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        public AnalyzerArgs()
        {
        }
        public static new AnalyzerArgs Empty => new AnalyzerArgs();
    }
}