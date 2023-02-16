// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Omics
{
    /// <summary>
    /// Definition of AWS::Omics::Workflow Resource Type
    /// </summary>
    [AwsNativeResourceType("aws-native:omics:Workflow")]
    public partial class Workflow : global::Pulumi.CustomResource
    {
        [Output("arn")]
        public Output<string> Arn { get; private set; } = null!;

        [Output("creationTime")]
        public Output<string> CreationTime { get; private set; } = null!;

        [Output("definitionUri")]
        public Output<string?> DefinitionUri { get; private set; } = null!;

        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        [Output("engine")]
        public Output<Pulumi.AwsNative.Omics.WorkflowEngine?> Engine { get; private set; } = null!;

        [Output("main")]
        public Output<string?> Main { get; private set; } = null!;

        [Output("name")]
        public Output<string?> Name { get; private set; } = null!;

        [Output("parameterTemplate")]
        public Output<Outputs.WorkflowParameterTemplate?> ParameterTemplate { get; private set; } = null!;

        [Output("status")]
        public Output<Pulumi.AwsNative.Omics.WorkflowStatus> Status { get; private set; } = null!;

        [Output("storageCapacity")]
        public Output<double?> StorageCapacity { get; private set; } = null!;

        [Output("tags")]
        public Output<Outputs.WorkflowTagMap?> Tags { get; private set; } = null!;

        [Output("type")]
        public Output<Pulumi.AwsNative.Omics.WorkflowType> Type { get; private set; } = null!;


        /// <summary>
        /// Create a Workflow resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Workflow(string name, WorkflowArgs? args = null, CustomResourceOptions? options = null)
            : base("aws-native:omics:Workflow", name, args ?? new WorkflowArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Workflow(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("aws-native:omics:Workflow", name, null, MakeResourceOptions(options, id))
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
        [Input("definitionUri")]
        public Input<string>? DefinitionUri { get; set; }

        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("engine")]
        public Input<Pulumi.AwsNative.Omics.WorkflowEngine>? Engine { get; set; }

        [Input("main")]
        public Input<string>? Main { get; set; }

        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("parameterTemplate")]
        public Input<Inputs.WorkflowParameterTemplateArgs>? ParameterTemplate { get; set; }

        [Input("storageCapacity")]
        public Input<double>? StorageCapacity { get; set; }

        [Input("tags")]
        public Input<Inputs.WorkflowTagMapArgs>? Tags { get; set; }

        public WorkflowArgs()
        {
        }
        public static new WorkflowArgs Empty => new WorkflowArgs();
    }
}