// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Ses
{
    /// <summary>
    /// Resource Type definition for AWS::SES::Template
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
    ///     var config = new Config();
    ///     var templateName = config.Require("templateName");
    ///     var subjectPart = config.Require("subjectPart");
    ///     var textPart = config.Require("textPart");
    ///     var htmlPart = config.Require("htmlPart");
    ///     var template = new AwsNative.Ses.Template("template", new()
    ///     {
    ///         TemplateValue = new AwsNative.Ses.Inputs.TemplateArgs
    ///         {
    ///             TemplateName = templateName,
    ///             SubjectPart = subjectPart,
    ///             TextPart = textPart,
    ///             HtmlPart = htmlPart,
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
    ///     var config = new Config();
    ///     var templateName = config.Require("templateName");
    ///     var subjectPart = config.Require("subjectPart");
    ///     var textPart = config.Require("textPart");
    ///     var htmlPart = config.Require("htmlPart");
    ///     var template = new AwsNative.Ses.Template("template", new()
    ///     {
    ///         TemplateValue = new AwsNative.Ses.Inputs.TemplateArgs
    ///         {
    ///             TemplateName = templateName,
    ///             SubjectPart = subjectPart,
    ///             TextPart = textPart,
    ///             HtmlPart = htmlPart,
    ///         },
    ///     });
    /// 
    /// });
    /// 
    /// 
    /// ```
    /// </summary>
    [AwsNativeResourceType("aws-native:ses:Template")]
    public partial class Template : global::Pulumi.CustomResource
    {
        [Output("awsId")]
        public Output<string> AwsId { get; private set; } = null!;

        /// <summary>
        /// The content of the email, composed of a subject line and either an HTML part or a text-only part.
        /// </summary>
        [Output("template")]
        public Output<Outputs.Template?> TemplateValue { get; private set; } = null!;


        /// <summary>
        /// Create a Template resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Template(string name, TemplateArgs? args = null, CustomResourceOptions? options = null)
            : base("aws-native:ses:Template", name, args ?? new TemplateArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Template(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("aws-native:ses:Template", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                ReplaceOnChanges =
                {
                    "template.templateName",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Template resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Template Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new Template(name, id, options);
        }
    }

    public sealed class TemplateArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The content of the email, composed of a subject line and either an HTML part or a text-only part.
        /// </summary>
        [Input("template")]
        public Input<Inputs.TemplateArgs>? TemplateValue { get; set; }

        public TemplateArgs()
        {
        }
        public static new TemplateArgs Empty => new TemplateArgs();
    }
}
