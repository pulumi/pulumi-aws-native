// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Proton
{
    public static class GetEnvironmentTemplate
    {
        /// <summary>
        /// Definition of AWS::Proton::EnvironmentTemplate Resource Type
        /// </summary>
        public static Task<GetEnvironmentTemplateResult> InvokeAsync(GetEnvironmentTemplateArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetEnvironmentTemplateResult>("aws-native:proton:getEnvironmentTemplate", args ?? new GetEnvironmentTemplateArgs(), options.WithDefaults());

        /// <summary>
        /// Definition of AWS::Proton::EnvironmentTemplate Resource Type
        /// </summary>
        public static Output<GetEnvironmentTemplateResult> Invoke(GetEnvironmentTemplateInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetEnvironmentTemplateResult>("aws-native:proton:getEnvironmentTemplate", args ?? new GetEnvironmentTemplateInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetEnvironmentTemplateArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// &lt;p&gt;The Amazon Resource Name (ARN) of the environment template.&lt;/p&gt;
        /// </summary>
        [Input("arn", required: true)]
        public string Arn { get; set; } = null!;

        public GetEnvironmentTemplateArgs()
        {
        }
        public static new GetEnvironmentTemplateArgs Empty => new GetEnvironmentTemplateArgs();
    }

    public sealed class GetEnvironmentTemplateInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// &lt;p&gt;The Amazon Resource Name (ARN) of the environment template.&lt;/p&gt;
        /// </summary>
        [Input("arn", required: true)]
        public Input<string> Arn { get; set; } = null!;

        public GetEnvironmentTemplateInvokeArgs()
        {
        }
        public static new GetEnvironmentTemplateInvokeArgs Empty => new GetEnvironmentTemplateInvokeArgs();
    }


    [OutputType]
    public sealed class GetEnvironmentTemplateResult
    {
        /// <summary>
        /// &lt;p&gt;The Amazon Resource Name (ARN) of the environment template.&lt;/p&gt;
        /// </summary>
        public readonly string? Arn;
        /// <summary>
        /// &lt;p&gt;A description of the environment template.&lt;/p&gt;
        /// </summary>
        public readonly string? Description;
        /// <summary>
        /// &lt;p&gt;The environment template name as displayed in the developer interface.&lt;/p&gt;
        /// </summary>
        public readonly string? DisplayName;
        /// <summary>
        /// &lt;p&gt;An optional list of metadata items that you can associate with the Proton environment template. A tag is a key-value pair.&lt;/p&gt;
        ///          &lt;p&gt;For more information, see &lt;a href="https://docs.aws.amazon.com/proton/latest/userguide/resources.html"&gt;Proton resources and tagging&lt;/a&gt; in the
        ///         &lt;i&gt;Proton User Guide&lt;/i&gt;.&lt;/p&gt;
        /// </summary>
        public readonly ImmutableArray<Outputs.EnvironmentTemplateTag> Tags;

        [OutputConstructor]
        private GetEnvironmentTemplateResult(
            string? arn,

            string? description,

            string? displayName,

            ImmutableArray<Outputs.EnvironmentTemplateTag> tags)
        {
            Arn = arn;
            Description = description;
            DisplayName = displayName;
            Tags = tags;
        }
    }
}