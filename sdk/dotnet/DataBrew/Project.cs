// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.DataBrew
{
    /// <summary>
    /// Resource schema for AWS::DataBrew::Project.
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
    ///     var testDataBrewProject = new AwsNative.DataBrew.Project("testDataBrewProject", new()
    ///     {
    ///         Name = "project-name",
    ///         RecipeName = "recipe-name",
    ///         DatasetName = "dataset-name",
    ///         RoleArn = "arn:aws:iam::12345678910:role/PassRoleAdmin",
    ///         Sample = new AwsNative.DataBrew.Inputs.ProjectSampleArgs
    ///         {
    ///             Size = 500,
    ///             Type = AwsNative.DataBrew.ProjectSampleType.LastN,
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
    ///     var myDataBrewProject = new AwsNative.DataBrew.Project("myDataBrewProject", new()
    ///     {
    ///         Name = "test-project",
    ///         RecipeName = "test-project-recipe",
    ///         DatasetName = "test-dataset",
    ///         RoleArn = "arn:aws:iam::1234567891011:role/PassRoleAdmin",
    ///         Sample = new AwsNative.DataBrew.Inputs.ProjectSampleArgs
    ///         {
    ///             Size = 500,
    ///             Type = AwsNative.DataBrew.ProjectSampleType.LastN,
    ///         },
    ///         Tags = new[]
    ///         {
    ///             new AwsNative.Inputs.TagArgs
    ///             {
    ///                 Key = "key00AtCreate",
    ///                 Value = "value001AtCreate",
    ///             },
    ///         },
    ///     });
    /// 
    /// });
    /// 
    /// 
    /// ```
    /// </summary>
    [AwsNativeResourceType("aws-native:databrew:Project")]
    public partial class Project : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Dataset name
        /// </summary>
        [Output("datasetName")]
        public Output<string> DatasetName { get; private set; } = null!;

        /// <summary>
        /// Project name
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Recipe name
        /// </summary>
        [Output("recipeName")]
        public Output<string> RecipeName { get; private set; } = null!;

        /// <summary>
        /// Role arn
        /// </summary>
        [Output("roleArn")]
        public Output<string> RoleArn { get; private set; } = null!;

        /// <summary>
        /// Sample
        /// </summary>
        [Output("sample")]
        public Output<Outputs.ProjectSample?> Sample { get; private set; } = null!;

        /// <summary>
        /// Metadata tags that have been applied to the project.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableArray<Pulumi.AwsNative.Outputs.Tag>> Tags { get; private set; } = null!;


        /// <summary>
        /// Create a Project resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Project(string name, ProjectArgs args, CustomResourceOptions? options = null)
            : base("aws-native:databrew:Project", name, args ?? new ProjectArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Project(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("aws-native:databrew:Project", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                ReplaceOnChanges =
                {
                    "name",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Project resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Project Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new Project(name, id, options);
        }
    }

    public sealed class ProjectArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Dataset name
        /// </summary>
        [Input("datasetName", required: true)]
        public Input<string> DatasetName { get; set; } = null!;

        /// <summary>
        /// Project name
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Recipe name
        /// </summary>
        [Input("recipeName", required: true)]
        public Input<string> RecipeName { get; set; } = null!;

        /// <summary>
        /// Role arn
        /// </summary>
        [Input("roleArn", required: true)]
        public Input<string> RoleArn { get; set; } = null!;

        /// <summary>
        /// Sample
        /// </summary>
        [Input("sample")]
        public Input<Inputs.ProjectSampleArgs>? Sample { get; set; }

        [Input("tags")]
        private InputList<Pulumi.AwsNative.Inputs.TagArgs>? _tags;

        /// <summary>
        /// Metadata tags that have been applied to the project.
        /// </summary>
        public InputList<Pulumi.AwsNative.Inputs.TagArgs> Tags
        {
            get => _tags ?? (_tags = new InputList<Pulumi.AwsNative.Inputs.TagArgs>());
            set => _tags = value;
        }

        public ProjectArgs()
        {
        }
        public static new ProjectArgs Empty => new ProjectArgs();
    }
}
