// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative
{
    /// <summary>
    /// A special resource that enables deploying CloudFormation Extensions (third-party resources). An extension has to be pre-registered in your AWS account in order to use this resource.
    /// </summary>
    [AwsNativeResourceType("aws-native:index:ExtensionResource")]
    public partial class ExtensionResource : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Dictionary of the extension resource attributes.
        /// </summary>
        [Output("outputs")]
        public Output<ImmutableDictionary<string, object>> Outputs { get; private set; } = null!;


        /// <summary>
        /// Create a ExtensionResource resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ExtensionResource(string name, ExtensionResourceArgs args, CustomResourceOptions? options = null)
            : base("aws-native:index:ExtensionResource", name, args ?? new ExtensionResourceArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ExtensionResource(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("aws-native:index:ExtensionResource", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing ExtensionResource resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ExtensionResource Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new ExtensionResource(name, id, options);
        }
    }

    public sealed class ExtensionResourceArgs : global::Pulumi.ResourceArgs
    {
        [Input("properties", required: true)]
        private InputMap<object>? _properties;

        /// <summary>
        /// Dictionary of the extension resource properties.
        /// </summary>
        public InputMap<object> Properties
        {
            get => _properties ?? (_properties = new InputMap<object>());
            set => _properties = value;
        }

        /// <summary>
        /// CloudFormation type name.
        /// </summary>
        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        public ExtensionResourceArgs()
        {
        }
        public static new ExtensionResourceArgs Empty => new ExtensionResourceArgs();
    }
}