// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Lightsail
{
    /// <summary>
    /// Resource Type definition for AWS::Lightsail::LoadBalancer
    /// </summary>
    [AwsNativeResourceType("aws-native:lightsail:LoadBalancer")]
    public partial class LoadBalancer : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The names of the instances attached to the load balancer.
        /// </summary>
        [Output("attachedInstances")]
        public Output<ImmutableArray<string>> AttachedInstances { get; private set; } = null!;

        /// <summary>
        /// The path you provided to perform the load balancer health check. If you didn't specify a health check path, Lightsail uses the root path of your website (e.g., "/").
        /// </summary>
        [Output("healthCheckPath")]
        public Output<string?> HealthCheckPath { get; private set; } = null!;

        /// <summary>
        /// The instance port where you're creating your load balancer.
        /// </summary>
        [Output("instancePort")]
        public Output<int> InstancePort { get; private set; } = null!;

        /// <summary>
        /// The IP address type for the load balancer. The possible values are ipv4 for IPv4 only, and dualstack for IPv4 and IPv6. The default value is dualstack.
        /// </summary>
        [Output("ipAddressType")]
        public Output<string?> IpAddressType { get; private set; } = null!;

        [Output("loadBalancerArn")]
        public Output<string> LoadBalancerArn { get; private set; } = null!;

        /// <summary>
        /// The name of your load balancer.
        /// </summary>
        [Output("loadBalancerName")]
        public Output<string> LoadBalancerName { get; private set; } = null!;

        /// <summary>
        /// Configuration option to enable session stickiness.
        /// </summary>
        [Output("sessionStickinessEnabled")]
        public Output<bool?> SessionStickinessEnabled { get; private set; } = null!;

        /// <summary>
        /// Configuration option to adjust session stickiness cookie duration parameter.
        /// </summary>
        [Output("sessionStickinessLBCookieDurationSeconds")]
        public Output<string?> SessionStickinessLBCookieDurationSeconds { get; private set; } = null!;

        /// <summary>
        /// An array of key-value pairs to apply to this resource.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableArray<Outputs.LoadBalancerTag>> Tags { get; private set; } = null!;

        /// <summary>
        /// The name of the TLS policy to apply to the load balancer.
        /// </summary>
        [Output("tlsPolicyName")]
        public Output<string?> TlsPolicyName { get; private set; } = null!;


        /// <summary>
        /// Create a LoadBalancer resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public LoadBalancer(string name, LoadBalancerArgs args, CustomResourceOptions? options = null)
            : base("aws-native:lightsail:LoadBalancer", name, args ?? new LoadBalancerArgs(), MakeResourceOptions(options, ""))
        {
        }

        private LoadBalancer(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("aws-native:lightsail:LoadBalancer", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing LoadBalancer resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static LoadBalancer Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new LoadBalancer(name, id, options);
        }
    }

    public sealed class LoadBalancerArgs : global::Pulumi.ResourceArgs
    {
        [Input("attachedInstances")]
        private InputList<string>? _attachedInstances;

        /// <summary>
        /// The names of the instances attached to the load balancer.
        /// </summary>
        public InputList<string> AttachedInstances
        {
            get => _attachedInstances ?? (_attachedInstances = new InputList<string>());
            set => _attachedInstances = value;
        }

        /// <summary>
        /// The path you provided to perform the load balancer health check. If you didn't specify a health check path, Lightsail uses the root path of your website (e.g., "/").
        /// </summary>
        [Input("healthCheckPath")]
        public Input<string>? HealthCheckPath { get; set; }

        /// <summary>
        /// The instance port where you're creating your load balancer.
        /// </summary>
        [Input("instancePort", required: true)]
        public Input<int> InstancePort { get; set; } = null!;

        /// <summary>
        /// The IP address type for the load balancer. The possible values are ipv4 for IPv4 only, and dualstack for IPv4 and IPv6. The default value is dualstack.
        /// </summary>
        [Input("ipAddressType")]
        public Input<string>? IpAddressType { get; set; }

        /// <summary>
        /// The name of your load balancer.
        /// </summary>
        [Input("loadBalancerName")]
        public Input<string>? LoadBalancerName { get; set; }

        /// <summary>
        /// Configuration option to enable session stickiness.
        /// </summary>
        [Input("sessionStickinessEnabled")]
        public Input<bool>? SessionStickinessEnabled { get; set; }

        /// <summary>
        /// Configuration option to adjust session stickiness cookie duration parameter.
        /// </summary>
        [Input("sessionStickinessLBCookieDurationSeconds")]
        public Input<string>? SessionStickinessLBCookieDurationSeconds { get; set; }

        [Input("tags")]
        private InputList<Inputs.LoadBalancerTagArgs>? _tags;

        /// <summary>
        /// An array of key-value pairs to apply to this resource.
        /// </summary>
        public InputList<Inputs.LoadBalancerTagArgs> Tags
        {
            get => _tags ?? (_tags = new InputList<Inputs.LoadBalancerTagArgs>());
            set => _tags = value;
        }

        /// <summary>
        /// The name of the TLS policy to apply to the load balancer.
        /// </summary>
        [Input("tlsPolicyName")]
        public Input<string>? TlsPolicyName { get; set; }

        public LoadBalancerArgs()
        {
        }
        public static new LoadBalancerArgs Empty => new LoadBalancerArgs();
    }
}