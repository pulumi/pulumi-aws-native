// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Glue
{
    /// <summary>
    /// Resource Type definition for AWS::Glue::Workflow
    /// </summary>
    [Obsolete(@"Workflow is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.")]
    [AwsNativeResourceType("aws-native:glue:Workflow")]
    public partial class Workflow : global::Pulumi.CustomResource
    {
        [Output("defaultRunProperties")]
        public Output<object?> DefaultRunProperties { get; private set; } = null!;

        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        [Output("maxConcurrentRuns")]
        public Output<int?> MaxConcurrentRuns { get; private set; } = null!;

        [Output("name")]
        public Output<string?> Name { get; private set; } = null!;

        [Output("tags")]
        public Output<object?> Tags { get; private set; } = null!;


        /// <summary>
        /// Create a Workflow resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Workflow(string name, WorkflowArgs? args = null, CustomResourceOptions? options = null)
            : base("aws-native:glue:Workflow", name, args ?? new WorkflowArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Workflow(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("aws-native:glue:Workflow", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing Workflow resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Workflow Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new Workflow(name, id, options);
        }
    }

    public sealed class WorkflowArgs : global::Pulumi.ResourceArgs
    {
        [Input("defaultRunProperties")]
        public Input<object>? DefaultRunProperties { get; set; }

        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("maxConcurrentRuns")]
        public Input<int>? MaxConcurrentRuns { get; set; }

        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("tags")]
        public Input<object>? Tags { get; set; }

        public WorkflowArgs()
        {
        }
        public static new WorkflowArgs Empty => new WorkflowArgs();
    }
}