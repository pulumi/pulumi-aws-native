// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.EC2
{
    /// <summary>
    /// An example resource schema demonstrating some basic constructs and validation rules.
    /// </summary>
    [AwsNativeResourceType("aws-native:ec2:NetworkPerformanceMetricSubscription")]
    public partial class NetworkPerformanceMetricSubscription : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The destination is a mandatory element for the metric subscription.
        /// </summary>
        [Output("destination")]
        public Output<string> Destination { get; private set; } = null!;

        /// <summary>
        /// The metric type for the metric subscription.
        /// </summary>
        [Output("metric")]
        public Output<string> Metric { get; private set; } = null!;

        /// <summary>
        /// The source is a mandatory element for the metric subscription.
        /// </summary>
        [Output("source")]
        public Output<string> Source { get; private set; } = null!;

        /// <summary>
        /// The statistic type for the metric subscription.
        /// </summary>
        [Output("statistic")]
        public Output<string> Statistic { get; private set; } = null!;


        /// <summary>
        /// Create a NetworkPerformanceMetricSubscription resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public NetworkPerformanceMetricSubscription(string name, NetworkPerformanceMetricSubscriptionArgs args, CustomResourceOptions? options = null)
            : base("aws-native:ec2:NetworkPerformanceMetricSubscription", name, args ?? new NetworkPerformanceMetricSubscriptionArgs(), MakeResourceOptions(options, ""))
        {
        }

        private NetworkPerformanceMetricSubscription(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("aws-native:ec2:NetworkPerformanceMetricSubscription", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing NetworkPerformanceMetricSubscription resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static NetworkPerformanceMetricSubscription Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new NetworkPerformanceMetricSubscription(name, id, options);
        }
    }

    public sealed class NetworkPerformanceMetricSubscriptionArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The destination is a mandatory element for the metric subscription.
        /// </summary>
        [Input("destination", required: true)]
        public Input<string> Destination { get; set; } = null!;

        /// <summary>
        /// The metric type for the metric subscription.
        /// </summary>
        [Input("metric", required: true)]
        public Input<string> Metric { get; set; } = null!;

        /// <summary>
        /// The source is a mandatory element for the metric subscription.
        /// </summary>
        [Input("source", required: true)]
        public Input<string> Source { get; set; } = null!;

        /// <summary>
        /// The statistic type for the metric subscription.
        /// </summary>
        [Input("statistic", required: true)]
        public Input<string> Statistic { get; set; } = null!;

        public NetworkPerformanceMetricSubscriptionArgs()
        {
        }
        public static new NetworkPerformanceMetricSubscriptionArgs Empty => new NetworkPerformanceMetricSubscriptionArgs();
    }
}