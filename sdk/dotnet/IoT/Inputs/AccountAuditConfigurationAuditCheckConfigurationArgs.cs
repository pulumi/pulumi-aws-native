// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.IoT.Inputs
{

    /// <summary>
    /// The configuration for a specific audit check.
    /// </summary>
    public sealed class AccountAuditConfigurationAuditCheckConfigurationArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// True if the check is enabled.
        /// </summary>
        [Input("enabled")]
        public Input<bool>? Enabled { get; set; }

        public AccountAuditConfigurationAuditCheckConfigurationArgs()
        {
        }
        public static new AccountAuditConfigurationAuditCheckConfigurationArgs Empty => new AccountAuditConfigurationAuditCheckConfigurationArgs();
    }
}
