// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.SES.Outputs
{

    /// <summary>
    /// Used to associate a configuration set with an email identity.
    /// </summary>
    [OutputType]
    public sealed class EmailIdentityConfigurationSetAttributes
    {
        /// <summary>
        /// The configuration set to use by default when sending from this identity. Note that any configuration set defined in the email sending request takes precedence.
        /// </summary>
        public readonly string? ConfigurationSetName;

        [OutputConstructor]
        private EmailIdentityConfigurationSetAttributes(string? configurationSetName)
        {
            ConfigurationSetName = configurationSetName;
        }
    }
}