// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Deadline
{
    /// <summary>
    /// Definition of AWS::Deadline::Farm Resource Type
    /// </summary>
    [AwsNativeResourceType("aws-native:deadline:Farm")]
    public partial class Farm : global::Pulumi.CustomResource
    {
        [Output("arn")]
        public Output<string> Arn { get; private set; } = null!;

        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        [Output("displayName")]
        public Output<string> DisplayName { get; private set; } = null!;

        [Output("farmId")]
        public Output<string> FarmId { get; private set; } = null!;

        [Output("kmsKeyArn")]
        public Output<string?> KmsKeyArn { get; private set; } = null!;


        /// <summary>
        /// Create a Farm resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Farm(string name, FarmArgs args, CustomResourceOptions? options = null)
            : base("aws-native:deadline:Farm", name, args ?? new FarmArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Farm(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("aws-native:deadline:Farm", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                ReplaceOnChanges =
                {
                    "kmsKeyArn",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Farm resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Farm Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new Farm(name, id, options);
        }
    }

    public sealed class FarmArgs : global::Pulumi.ResourceArgs
    {
        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("displayName", required: true)]
        public Input<string> DisplayName { get; set; } = null!;

        [Input("kmsKeyArn")]
        public Input<string>? KmsKeyArn { get; set; }

        public FarmArgs()
        {
        }
        public static new FarmArgs Empty => new FarmArgs();
    }
}