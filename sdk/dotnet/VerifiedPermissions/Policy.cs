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
    /// Definition of AWS::VerifiedPermissions::Policy Resource Type
    /// </summary>
    [AwsNativeResourceType("aws-native:verifiedpermissions:Policy")]
    public partial class Policy : global::Pulumi.CustomResource
    {
        [Output("definition")]
        public Output<Outputs.PolicyDefinition> Definition { get; private set; } = null!;

        [Output("policyId")]
        public Output<string> PolicyId { get; private set; } = null!;

        [Output("policyStoreId")]
        public Output<string?> PolicyStoreId { get; private set; } = null!;

        [Output("policyType")]
        public Output<Pulumi.AwsNative.VerifiedPermissions.PolicyType> PolicyType { get; private set; } = null!;


        /// <summary>
        /// Create a Policy resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Policy(string name, PolicyArgs args, CustomResourceOptions? options = null)
            : base("aws-native:verifiedpermissions:Policy", name, args ?? new PolicyArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Policy(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("aws-native:verifiedpermissions:Policy", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing Policy resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Policy Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new Policy(name, id, options);
        }
    }

    public sealed class PolicyArgs : global::Pulumi.ResourceArgs
    {
        [Input("definition", required: true)]
        public Input<Inputs.PolicyDefinitionArgs> Definition { get; set; } = null!;

        [Input("policyStoreId")]
        public Input<string>? PolicyStoreId { get; set; }

        public PolicyArgs()
        {
        }
        public static new PolicyArgs Empty => new PolicyArgs();
    }
}