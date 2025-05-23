// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Ec2.Outputs
{

    /// <summary>
    /// Options for sending VPN tunnel logs to CloudWatch.
    /// </summary>
    [OutputType]
    public sealed class VpnConnectionCloudwatchLogOptionsSpecification
    {
        /// <summary>
        /// Enable or disable VPN tunnel logging feature. Default value is ``False``.
        ///  Valid values: ``True`` | ``False``
        /// </summary>
        public readonly bool? LogEnabled;
        /// <summary>
        /// The Amazon Resource Name (ARN) of the CloudWatch log group to send logs to.
        /// </summary>
        public readonly string? LogGroupArn;
        /// <summary>
        /// Set log format. Default format is ``json``.
        ///  Valid values: ``json`` | ``text``
        /// </summary>
        public readonly Pulumi.AwsNative.Ec2.VpnConnectionCloudwatchLogOptionsSpecificationLogOutputFormat? LogOutputFormat;

        [OutputConstructor]
        private VpnConnectionCloudwatchLogOptionsSpecification(
            bool? logEnabled,

            string? logGroupArn,

            Pulumi.AwsNative.Ec2.VpnConnectionCloudwatchLogOptionsSpecificationLogOutputFormat? logOutputFormat)
        {
            LogEnabled = logEnabled;
            LogGroupArn = logGroupArn;
            LogOutputFormat = logOutputFormat;
        }
    }
}
