// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Neptune
{
    /// <summary>
    /// Resource Type definition for AWS::Neptune::DBParameterGroup
    /// </summary>
    [Obsolete(@"DBParameterGroup is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.")]
    [AwsNativeResourceType("aws-native:neptune:DBParameterGroup")]
    public partial class DBParameterGroup : global::Pulumi.CustomResource
    {
        [Output("description")]
        public Output<string> Description { get; private set; } = null!;

        [Output("family")]
        public Output<string> Family { get; private set; } = null!;

        [Output("name")]
        public Output<string?> Name { get; private set; } = null!;

        [Output("parameters")]
        public Output<object> Parameters { get; private set; } = null!;

        [Output("tags")]
        public Output<ImmutableArray<Outputs.DBParameterGroupTag>> Tags { get; private set; } = null!;


        /// <summary>
        /// Create a DBParameterGroup resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public DBParameterGroup(string name, DBParameterGroupArgs args, CustomResourceOptions? options = null)
            : base("aws-native:neptune:DBParameterGroup", name, args ?? new DBParameterGroupArgs(), MakeResourceOptions(options, ""))
        {
        }

        private DBParameterGroup(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("aws-native:neptune:DBParameterGroup", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing DBParameterGroup resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static DBParameterGroup Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new DBParameterGroup(name, id, options);
        }
    }

    public sealed class DBParameterGroupArgs : global::Pulumi.ResourceArgs
    {
        [Input("description", required: true)]
        public Input<string> Description { get; set; } = null!;

        [Input("family", required: true)]
        public Input<string> Family { get; set; } = null!;

        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("parameters", required: true)]
        public Input<object> Parameters { get; set; } = null!;

        [Input("tags")]
        private InputList<Inputs.DBParameterGroupTagArgs>? _tags;
        public InputList<Inputs.DBParameterGroupTagArgs> Tags
        {
            get => _tags ?? (_tags = new InputList<Inputs.DBParameterGroupTagArgs>());
            set => _tags = value;
        }

        public DBParameterGroupArgs()
        {
        }
        public static new DBParameterGroupArgs Empty => new DBParameterGroupArgs();
    }
}