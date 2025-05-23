// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Personalize
{
    /// <summary>
    /// Resource schema for AWS::Personalize::Solution.
    /// 
    /// ## Example Usage
    /// ### Example
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using AwsNative = Pulumi.AwsNative;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var mySolution = new AwsNative.Personalize.Solution("mySolution", new()
    ///     {
    ///         Name = "my-solution-name",
    ///         DatasetGroupArn = "arn:aws:personalize:us-west-2:123456789012:dataset-group/my-dataset-group-name",
    ///         RecipeArn = "arn:aws:personalize:::recipe/aws-user-personalization",
    ///         SolutionConfig = new AwsNative.Personalize.Inputs.SolutionConfigArgs
    ///         {
    ///             EventValueThreshold = ".05",
    ///         },
    ///     });
    /// 
    /// });
    /// 
    /// 
    /// ```
    /// ### Example
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using AwsNative = Pulumi.AwsNative;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var mySolution = new AwsNative.Personalize.Solution("mySolution", new()
    ///     {
    ///         Name = "my-solution-name",
    ///         DatasetGroupArn = "arn:aws:personalize:us-west-2:123456789012:dataset-group/my-dataset-group-name",
    ///         RecipeArn = "arn:aws:personalize:::recipe/aws-user-personalization",
    ///         SolutionConfig = new AwsNative.Personalize.Inputs.SolutionConfigArgs
    ///         {
    ///             EventValueThreshold = ".05",
    ///         },
    ///     });
    /// 
    /// });
    /// 
    /// 
    /// ```
    /// </summary>
    [AwsNativeResourceType("aws-native:personalize:Solution")]
    public partial class Solution : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The ARN of the dataset group that provides the training data.
        /// </summary>
        [Output("datasetGroupArn")]
        public Output<string> DatasetGroupArn { get; private set; } = null!;

        /// <summary>
        /// When your have multiple event types (using an EVENT_TYPE schema field), this parameter specifies which event type (for example, 'click' or 'like') is used for training the model. If you do not provide an eventType, Amazon Personalize will use all interactions for training with equal weight regardless of type.
        /// </summary>
        [Output("eventType")]
        public Output<string?> EventType { get; private set; } = null!;

        /// <summary>
        /// The name for the solution
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Whether to perform automated machine learning (AutoML). The default is false. For this case, you must specify recipeArn.
        /// </summary>
        [Output("performAutoMl")]
        public Output<bool?> PerformAutoMl { get; private set; } = null!;

        /// <summary>
        /// Whether to perform hyperparameter optimization (HPO) on the specified or selected recipe. The default is false. When performing AutoML, this parameter is always true and you should not set it to false.
        /// </summary>
        [Output("performHpo")]
        public Output<bool?> PerformHpo { get; private set; } = null!;

        /// <summary>
        /// The ARN of the recipe to use for model training. Only specified when performAutoML is false.
        /// </summary>
        [Output("recipeArn")]
        public Output<string?> RecipeArn { get; private set; } = null!;

        /// <summary>
        /// The Amazon Resource Name (ARN) of the solution.
        /// </summary>
        [Output("solutionArn")]
        public Output<string> SolutionArn { get; private set; } = null!;

        /// <summary>
        /// Describes the configuration properties for the solution.
        /// </summary>
        [Output("solutionConfig")]
        public Output<Outputs.SolutionConfig?> SolutionConfig { get; private set; } = null!;


        /// <summary>
        /// Create a Solution resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Solution(string name, SolutionArgs args, CustomResourceOptions? options = null)
            : base("aws-native:personalize:Solution", name, args ?? new SolutionArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Solution(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("aws-native:personalize:Solution", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                ReplaceOnChanges =
                {
                    "datasetGroupArn",
                    "eventType",
                    "name",
                    "performAutoMl",
                    "performHpo",
                    "recipeArn",
                    "solutionConfig",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Solution resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Solution Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new Solution(name, id, options);
        }
    }

    public sealed class SolutionArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ARN of the dataset group that provides the training data.
        /// </summary>
        [Input("datasetGroupArn", required: true)]
        public Input<string> DatasetGroupArn { get; set; } = null!;

        /// <summary>
        /// When your have multiple event types (using an EVENT_TYPE schema field), this parameter specifies which event type (for example, 'click' or 'like') is used for training the model. If you do not provide an eventType, Amazon Personalize will use all interactions for training with equal weight regardless of type.
        /// </summary>
        [Input("eventType")]
        public Input<string>? EventType { get; set; }

        /// <summary>
        /// The name for the solution
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Whether to perform automated machine learning (AutoML). The default is false. For this case, you must specify recipeArn.
        /// </summary>
        [Input("performAutoMl")]
        public Input<bool>? PerformAutoMl { get; set; }

        /// <summary>
        /// Whether to perform hyperparameter optimization (HPO) on the specified or selected recipe. The default is false. When performing AutoML, this parameter is always true and you should not set it to false.
        /// </summary>
        [Input("performHpo")]
        public Input<bool>? PerformHpo { get; set; }

        /// <summary>
        /// The ARN of the recipe to use for model training. Only specified when performAutoML is false.
        /// </summary>
        [Input("recipeArn")]
        public Input<string>? RecipeArn { get; set; }

        /// <summary>
        /// Describes the configuration properties for the solution.
        /// </summary>
        [Input("solutionConfig")]
        public Input<Inputs.SolutionConfigArgs>? SolutionConfig { get; set; }

        public SolutionArgs()
        {
        }
        public static new SolutionArgs Empty => new SolutionArgs();
    }
}
