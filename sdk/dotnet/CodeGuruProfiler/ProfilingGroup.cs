// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.CodeGuruProfiler
{
    /// <summary>
    /// This resource schema represents the Profiling Group resource in the Amazon CodeGuru Profiler service.
    /// </summary>
    [AwsNativeResourceType("aws-native:codeguruprofiler:ProfilingGroup")]
    public partial class ProfilingGroup : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The agent permissions attached to this profiling group.
        /// </summary>
        [Output("agentPermissions")]
        public Output<Outputs.AgentPermissionsProperties?> AgentPermissions { get; private set; } = null!;

        /// <summary>
        /// Configuration for Notification Channels for Anomaly Detection feature in CodeGuru Profiler which enables customers to detect anomalies in the application profile for those methods that represent the highest proportion of CPU time or latency
        /// </summary>
        [Output("anomalyDetectionNotificationConfiguration")]
        public Output<ImmutableArray<Outputs.ProfilingGroupChannel>> AnomalyDetectionNotificationConfiguration { get; private set; } = null!;

        /// <summary>
        /// The Amazon Resource Name (ARN) of the specified profiling group.
        /// </summary>
        [Output("arn")]
        public Output<string> Arn { get; private set; } = null!;

        /// <summary>
        /// The compute platform of the profiling group.
        /// </summary>
        [Output("computePlatform")]
        public Output<Pulumi.AwsNative.CodeGuruProfiler.ProfilingGroupComputePlatform?> ComputePlatform { get; private set; } = null!;

        /// <summary>
        /// The name of the profiling group.
        /// </summary>
        [Output("profilingGroupName")]
        public Output<string> ProfilingGroupName { get; private set; } = null!;

        /// <summary>
        /// The tags associated with a profiling group.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableArray<Pulumi.AwsNative.Outputs.Tag>> Tags { get; private set; } = null!;


        /// <summary>
        /// Create a ProfilingGroup resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ProfilingGroup(string name, ProfilingGroupArgs? args = null, CustomResourceOptions? options = null)
            : base("aws-native:codeguruprofiler:ProfilingGroup", name, args ?? new ProfilingGroupArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ProfilingGroup(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("aws-native:codeguruprofiler:ProfilingGroup", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                ReplaceOnChanges =
                {
                    "computePlatform",
                    "profilingGroupName",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing ProfilingGroup resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ProfilingGroup Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new ProfilingGroup(name, id, options);
        }
    }

    public sealed class ProfilingGroupArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The agent permissions attached to this profiling group.
        /// </summary>
        [Input("agentPermissions")]
        public Input<Inputs.AgentPermissionsPropertiesArgs>? AgentPermissions { get; set; }

        [Input("anomalyDetectionNotificationConfiguration")]
        private InputList<Inputs.ProfilingGroupChannelArgs>? _anomalyDetectionNotificationConfiguration;

        /// <summary>
        /// Configuration for Notification Channels for Anomaly Detection feature in CodeGuru Profiler which enables customers to detect anomalies in the application profile for those methods that represent the highest proportion of CPU time or latency
        /// </summary>
        public InputList<Inputs.ProfilingGroupChannelArgs> AnomalyDetectionNotificationConfiguration
        {
            get => _anomalyDetectionNotificationConfiguration ?? (_anomalyDetectionNotificationConfiguration = new InputList<Inputs.ProfilingGroupChannelArgs>());
            set => _anomalyDetectionNotificationConfiguration = value;
        }

        /// <summary>
        /// The compute platform of the profiling group.
        /// </summary>
        [Input("computePlatform")]
        public Input<Pulumi.AwsNative.CodeGuruProfiler.ProfilingGroupComputePlatform>? ComputePlatform { get; set; }

        /// <summary>
        /// The name of the profiling group.
        /// </summary>
        [Input("profilingGroupName")]
        public Input<string>? ProfilingGroupName { get; set; }

        [Input("tags")]
        private InputList<Pulumi.AwsNative.Inputs.TagArgs>? _tags;

        /// <summary>
        /// The tags associated with a profiling group.
        /// </summary>
        public InputList<Pulumi.AwsNative.Inputs.TagArgs> Tags
        {
            get => _tags ?? (_tags = new InputList<Pulumi.AwsNative.Inputs.TagArgs>());
            set => _tags = value;
        }

        public ProfilingGroupArgs()
        {
        }
        public static new ProfilingGroupArgs Empty => new ProfilingGroupArgs();
    }
}
