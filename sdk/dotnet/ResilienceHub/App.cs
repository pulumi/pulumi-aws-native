// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.ResilienceHub
{
    /// <summary>
    /// Resource Type Definition for AWS::ResilienceHub::App.
    /// </summary>
    [AwsNativeResourceType("aws-native:resiliencehub:App")]
    public partial class App : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Amazon Resource Name (ARN) of the App.
        /// </summary>
        [Output("appArn")]
        public Output<string> AppArn { get; private set; } = null!;

        /// <summary>
        /// Assessment execution schedule.
        /// </summary>
        [Output("appAssessmentSchedule")]
        public Output<Pulumi.AwsNative.ResilienceHub.AppAssessmentSchedule?> AppAssessmentSchedule { get; private set; } = null!;

        /// <summary>
        /// A string containing full ResilienceHub app template body.
        /// </summary>
        [Output("appTemplateBody")]
        public Output<string> AppTemplateBody { get; private set; } = null!;

        /// <summary>
        /// App description.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// Name of the app.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Amazon Resource Name (ARN) of the Resiliency Policy.
        /// </summary>
        [Output("resiliencyPolicyArn")]
        public Output<string?> ResiliencyPolicyArn { get; private set; } = null!;

        /// <summary>
        /// An array of ResourceMapping objects.
        /// </summary>
        [Output("resourceMappings")]
        public Output<ImmutableArray<Outputs.AppResourceMapping>> ResourceMappings { get; private set; } = null!;

        [Output("tags")]
        public Output<Outputs.AppTagMap?> Tags { get; private set; } = null!;


        /// <summary>
        /// Create a App resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public App(string name, AppArgs args, CustomResourceOptions? options = null)
            : base("aws-native:resiliencehub:App", name, args ?? new AppArgs(), MakeResourceOptions(options, ""))
        {
        }

        private App(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("aws-native:resiliencehub:App", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing App resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static App Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new App(name, id, options);
        }
    }

    public sealed class AppArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Assessment execution schedule.
        /// </summary>
        [Input("appAssessmentSchedule")]
        public Input<Pulumi.AwsNative.ResilienceHub.AppAssessmentSchedule>? AppAssessmentSchedule { get; set; }

        /// <summary>
        /// A string containing full ResilienceHub app template body.
        /// </summary>
        [Input("appTemplateBody", required: true)]
        public Input<string> AppTemplateBody { get; set; } = null!;

        /// <summary>
        /// App description.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Name of the app.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Amazon Resource Name (ARN) of the Resiliency Policy.
        /// </summary>
        [Input("resiliencyPolicyArn")]
        public Input<string>? ResiliencyPolicyArn { get; set; }

        [Input("resourceMappings", required: true)]
        private InputList<Inputs.AppResourceMappingArgs>? _resourceMappings;

        /// <summary>
        /// An array of ResourceMapping objects.
        /// </summary>
        public InputList<Inputs.AppResourceMappingArgs> ResourceMappings
        {
            get => _resourceMappings ?? (_resourceMappings = new InputList<Inputs.AppResourceMappingArgs>());
            set => _resourceMappings = value;
        }

        [Input("tags")]
        public Input<Inputs.AppTagMapArgs>? Tags { get; set; }

        public AppArgs()
        {
        }
        public static new AppArgs Empty => new AppArgs();
    }
}