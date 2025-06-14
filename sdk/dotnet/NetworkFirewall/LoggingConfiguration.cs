// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.NetworkFirewall
{
    /// <summary>
    /// Resource type definition for AWS::NetworkFirewall::LoggingConfiguration
    /// </summary>
    [AwsNativeResourceType("aws-native:networkfirewall:LoggingConfiguration")]
    public partial class LoggingConfiguration : global::Pulumi.CustomResource
    {
        [Output("enableMonitoringDashboard")]
        public Output<bool?> EnableMonitoringDashboard { get; private set; } = null!;

        /// <summary>
        /// The Amazon Resource Name (ARN) of the `Firewall` that the logging configuration is associated with. You can't change the firewall specification after you create the logging configuration.
        /// </summary>
        [Output("firewallArn")]
        public Output<string> FirewallArn { get; private set; } = null!;

        /// <summary>
        /// The name of the firewall that the logging configuration is associated with. You can't change the firewall specification after you create the logging configuration.
        /// </summary>
        [Output("firewallName")]
        public Output<string?> FirewallName { get; private set; } = null!;

        /// <summary>
        /// Defines how AWS Network Firewall performs logging for a `Firewall` .
        /// </summary>
        [Output("loggingConfiguration")]
        public Output<Outputs.LoggingConfiguration> LoggingConfigurationValue { get; private set; } = null!;


        /// <summary>
        /// Create a LoggingConfiguration resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public LoggingConfiguration(string name, LoggingConfigurationArgs args, CustomResourceOptions? options = null)
            : base("aws-native:networkfirewall:LoggingConfiguration", name, args ?? new LoggingConfigurationArgs(), MakeResourceOptions(options, ""))
        {
        }

        private LoggingConfiguration(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("aws-native:networkfirewall:LoggingConfiguration", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                ReplaceOnChanges =
                {
                    "firewallArn",
                    "firewallName",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing LoggingConfiguration resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static LoggingConfiguration Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new LoggingConfiguration(name, id, options);
        }
    }

    public sealed class LoggingConfigurationArgs : global::Pulumi.ResourceArgs
    {
        [Input("enableMonitoringDashboard")]
        public Input<bool>? EnableMonitoringDashboard { get; set; }

        /// <summary>
        /// The Amazon Resource Name (ARN) of the `Firewall` that the logging configuration is associated with. You can't change the firewall specification after you create the logging configuration.
        /// </summary>
        [Input("firewallArn", required: true)]
        public Input<string> FirewallArn { get; set; } = null!;

        /// <summary>
        /// The name of the firewall that the logging configuration is associated with. You can't change the firewall specification after you create the logging configuration.
        /// </summary>
        [Input("firewallName")]
        public Input<string>? FirewallName { get; set; }

        /// <summary>
        /// Defines how AWS Network Firewall performs logging for a `Firewall` .
        /// </summary>
        [Input("loggingConfiguration", required: true)]
        public Input<Inputs.LoggingConfigurationArgs> LoggingConfigurationValue { get; set; } = null!;

        public LoggingConfigurationArgs()
        {
        }
        public static new LoggingConfigurationArgs Empty => new LoggingConfigurationArgs();
    }
}
