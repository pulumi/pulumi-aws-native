// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.WorkSpaces
{
    /// <summary>
    /// Resource Type definition for AWS::WorkSpaces::WorkspacesPool
    /// </summary>
    [AwsNativeResourceType("aws-native:workspaces:WorkspacesPool")]
    public partial class WorkspacesPool : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The persistent application settings for users of the pool.
        /// </summary>
        [Output("applicationSettings")]
        public Output<Outputs.WorkspacesPoolApplicationSettings?> ApplicationSettings { get; private set; } = null!;

        /// <summary>
        /// The identifier of the bundle used by the pool.
        /// </summary>
        [Output("bundleId")]
        public Output<string> BundleId { get; private set; } = null!;

        /// <summary>
        /// Describes the user capacity for the pool.
        /// </summary>
        [Output("capacity")]
        public Output<Outputs.WorkspacesPoolCapacity> Capacity { get; private set; } = null!;

        /// <summary>
        /// The time the pool was created.
        /// </summary>
        [Output("createdAt")]
        public Output<string> CreatedAt { get; private set; } = null!;

        /// <summary>
        /// The description of the pool.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// The identifier of the directory used by the pool.
        /// </summary>
        [Output("directoryId")]
        public Output<string> DirectoryId { get; private set; } = null!;

        /// <summary>
        /// The Amazon Resource Name (ARN) for the pool.
        /// </summary>
        [Output("poolArn")]
        public Output<string> PoolArn { get; private set; } = null!;

        /// <summary>
        /// The identifier of the pool.
        /// </summary>
        [Output("poolId")]
        public Output<string> PoolId { get; private set; } = null!;

        /// <summary>
        /// The name of the pool.
        /// </summary>
        [Output("poolName")]
        public Output<string> PoolName { get; private set; } = null!;

        /// <summary>
        /// The running mode of the pool.
        /// </summary>
        [Output("runningMode")]
        public Output<Pulumi.AwsNative.WorkSpaces.WorkspacesPoolRunningMode?> RunningMode { get; private set; } = null!;

        [Output("tags")]
        public Output<ImmutableArray<Pulumi.AwsNative.Outputs.Tag>> Tags { get; private set; } = null!;

        /// <summary>
        /// The amount of time that a pool session remains active after users disconnect. If they try to reconnect to the pool session after a disconnection or network interruption within this time interval, they are connected to their previous session. Otherwise, they are connected to a new session with a new pool instance.
        /// </summary>
        [Output("timeoutSettings")]
        public Output<Outputs.WorkspacesPoolTimeoutSettings?> TimeoutSettings { get; private set; } = null!;


        /// <summary>
        /// Create a WorkspacesPool resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public WorkspacesPool(string name, WorkspacesPoolArgs args, CustomResourceOptions? options = null)
            : base("aws-native:workspaces:WorkspacesPool", name, args ?? new WorkspacesPoolArgs(), MakeResourceOptions(options, ""))
        {
        }

        private WorkspacesPool(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("aws-native:workspaces:WorkspacesPool", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                ReplaceOnChanges =
                {
                    "poolName",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing WorkspacesPool resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static WorkspacesPool Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new WorkspacesPool(name, id, options);
        }
    }

    public sealed class WorkspacesPoolArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The persistent application settings for users of the pool.
        /// </summary>
        [Input("applicationSettings")]
        public Input<Inputs.WorkspacesPoolApplicationSettingsArgs>? ApplicationSettings { get; set; }

        /// <summary>
        /// The identifier of the bundle used by the pool.
        /// </summary>
        [Input("bundleId", required: true)]
        public Input<string> BundleId { get; set; } = null!;

        /// <summary>
        /// Describes the user capacity for the pool.
        /// </summary>
        [Input("capacity", required: true)]
        public Input<Inputs.WorkspacesPoolCapacityArgs> Capacity { get; set; } = null!;

        /// <summary>
        /// The description of the pool.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The identifier of the directory used by the pool.
        /// </summary>
        [Input("directoryId", required: true)]
        public Input<string> DirectoryId { get; set; } = null!;

        /// <summary>
        /// The name of the pool.
        /// </summary>
        [Input("poolName")]
        public Input<string>? PoolName { get; set; }

        /// <summary>
        /// The running mode of the pool.
        /// </summary>
        [Input("runningMode")]
        public Input<Pulumi.AwsNative.WorkSpaces.WorkspacesPoolRunningMode>? RunningMode { get; set; }

        [Input("tags")]
        private InputList<Pulumi.AwsNative.Inputs.TagArgs>? _tags;
        public InputList<Pulumi.AwsNative.Inputs.TagArgs> Tags
        {
            get => _tags ?? (_tags = new InputList<Pulumi.AwsNative.Inputs.TagArgs>());
            set => _tags = value;
        }

        /// <summary>
        /// The amount of time that a pool session remains active after users disconnect. If they try to reconnect to the pool session after a disconnection or network interruption within this time interval, they are connected to their previous session. Otherwise, they are connected to a new session with a new pool instance.
        /// </summary>
        [Input("timeoutSettings")]
        public Input<Inputs.WorkspacesPoolTimeoutSettingsArgs>? TimeoutSettings { get; set; }

        public WorkspacesPoolArgs()
        {
        }
        public static new WorkspacesPoolArgs Empty => new WorkspacesPoolArgs();
    }
}
