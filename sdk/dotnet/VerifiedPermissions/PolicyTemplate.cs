// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.VerifiedPermissions
{
    /// <summary>
    /// Definition of AWS::VerifiedPermissions::PolicyTemplate Resource Type
    /// </summary>
    [AwsNativeResourceType("aws-native:verifiedpermissions:PolicyTemplate")]
    public partial class PolicyTemplate : global::Pulumi.CustomResource
    {
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        [Output("policyStoreId")]
        public Output<string?> PolicyStoreId { get; private set; } = null!;

        [Output("policyTemplateId")]
        public Output<string> PolicyTemplateId { get; private set; } = null!;

        [Output("statement")]
        public Output<string> Statement { get; private set; } = null!;


        /// <summary>
        /// Create a PolicyTemplate resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public PolicyTemplate(string name, PolicyTemplateArgs args, CustomResourceOptions? options = null)
            : base("aws-native:verifiedpermissions:PolicyTemplate", name, args ?? new PolicyTemplateArgs(), MakeResourceOptions(options, ""))
        {
        }

        private PolicyTemplate(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("aws-native:verifiedpermissions:PolicyTemplate", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing PolicyTemplate resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static PolicyTemplate Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new PolicyTemplate(name, id, options);
        }
    }

    public sealed class PolicyTemplateArgs : global::Pulumi.ResourceArgs
    {
        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("policyStoreId")]
        public Input<string>? PolicyStoreId { get; set; }

        [Input("statement", required: true)]
        public Input<string> Statement { get; set; } = null!;

        public PolicyTemplateArgs()
        {
        }
        public static new PolicyTemplateArgs Empty => new PolicyTemplateArgs();
    }
}