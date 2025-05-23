// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Configuration
{
    /// <summary>
    /// A conformance pack is a collection of AWS Config rules and remediation actions that can be easily deployed as a single entity in an account and a region or across an entire AWS Organization.
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
    ///     var conformancePack = new AwsNative.Configuration.ConformancePack("conformancePack", new()
    ///     {
    ///         ConformancePackName = "ConformancePackName",
    ///         DeliveryS3Bucket = "DeliveryS3Bucket",
    ///         TemplateS3Uri = "s3://bucketname/prefix",
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
    ///     var cloudFormationCanaryPack = new AwsNative.Configuration.ConformancePack("cloudFormationCanaryPack", new()
    ///     {
    ///         ConformancePackName = "ConformancePackName",
    ///         DeliveryS3Bucket = "DeliveryS3Bucket",
    ///         TemplateS3Uri = "s3://bucketname/prefix",
    ///     });
    /// 
    /// });
    /// 
    /// 
    /// ```
    /// </summary>
    [AwsNativeResourceType("aws-native:configuration:ConformancePack")]
    public partial class ConformancePack : global::Pulumi.CustomResource
    {
        /// <summary>
        /// A list of ConformancePackInputParameter objects.
        /// </summary>
        [Output("conformancePackInputParameters")]
        public Output<ImmutableArray<Outputs.ConformancePackInputParameter>> ConformancePackInputParameters { get; private set; } = null!;

        /// <summary>
        /// Name of the conformance pack which will be assigned as the unique identifier.
        /// </summary>
        [Output("conformancePackName")]
        public Output<string> ConformancePackName { get; private set; } = null!;

        /// <summary>
        /// AWS Config stores intermediate files while processing conformance pack template.
        /// </summary>
        [Output("deliveryS3Bucket")]
        public Output<string?> DeliveryS3Bucket { get; private set; } = null!;

        /// <summary>
        /// The prefix for delivery S3 bucket.
        /// </summary>
        [Output("deliveryS3KeyPrefix")]
        public Output<string?> DeliveryS3KeyPrefix { get; private set; } = null!;

        /// <summary>
        /// A string containing full conformance pack template body. You can only specify one of the template body or template S3Uri fields.
        /// </summary>
        [Output("templateBody")]
        public Output<string?> TemplateBody { get; private set; } = null!;

        /// <summary>
        /// Location of file containing the template body which points to the conformance pack template that is located in an Amazon S3 bucket. You can only specify one of the template body or template S3Uri fields.
        /// </summary>
        [Output("templateS3Uri")]
        public Output<string?> TemplateS3Uri { get; private set; } = null!;

        /// <summary>
        /// The TemplateSSMDocumentDetails object contains the name of the SSM document and the version of the SSM document.
        /// </summary>
        [Output("templateSsmDocumentDetails")]
        public Output<Outputs.TemplateSsmDocumentDetailsProperties?> TemplateSsmDocumentDetails { get; private set; } = null!;


        /// <summary>
        /// Create a ConformancePack resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ConformancePack(string name, ConformancePackArgs? args = null, CustomResourceOptions? options = null)
            : base("aws-native:configuration:ConformancePack", name, args ?? new ConformancePackArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ConformancePack(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("aws-native:configuration:ConformancePack", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                ReplaceOnChanges =
                {
                    "conformancePackName",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing ConformancePack resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ConformancePack Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new ConformancePack(name, id, options);
        }
    }

    public sealed class ConformancePackArgs : global::Pulumi.ResourceArgs
    {
        [Input("conformancePackInputParameters")]
        private InputList<Inputs.ConformancePackInputParameterArgs>? _conformancePackInputParameters;

        /// <summary>
        /// A list of ConformancePackInputParameter objects.
        /// </summary>
        public InputList<Inputs.ConformancePackInputParameterArgs> ConformancePackInputParameters
        {
            get => _conformancePackInputParameters ?? (_conformancePackInputParameters = new InputList<Inputs.ConformancePackInputParameterArgs>());
            set => _conformancePackInputParameters = value;
        }

        /// <summary>
        /// Name of the conformance pack which will be assigned as the unique identifier.
        /// </summary>
        [Input("conformancePackName")]
        public Input<string>? ConformancePackName { get; set; }

        /// <summary>
        /// AWS Config stores intermediate files while processing conformance pack template.
        /// </summary>
        [Input("deliveryS3Bucket")]
        public Input<string>? DeliveryS3Bucket { get; set; }

        /// <summary>
        /// The prefix for delivery S3 bucket.
        /// </summary>
        [Input("deliveryS3KeyPrefix")]
        public Input<string>? DeliveryS3KeyPrefix { get; set; }

        /// <summary>
        /// A string containing full conformance pack template body. You can only specify one of the template body or template S3Uri fields.
        /// </summary>
        [Input("templateBody")]
        public Input<string>? TemplateBody { get; set; }

        /// <summary>
        /// Location of file containing the template body which points to the conformance pack template that is located in an Amazon S3 bucket. You can only specify one of the template body or template S3Uri fields.
        /// </summary>
        [Input("templateS3Uri")]
        public Input<string>? TemplateS3Uri { get; set; }

        /// <summary>
        /// The TemplateSSMDocumentDetails object contains the name of the SSM document and the version of the SSM document.
        /// </summary>
        [Input("templateSsmDocumentDetails")]
        public Input<Inputs.TemplateSsmDocumentDetailsPropertiesArgs>? TemplateSsmDocumentDetails { get; set; }

        public ConformancePackArgs()
        {
        }
        public static new ConformancePackArgs Empty => new ConformancePackArgs();
    }
}
