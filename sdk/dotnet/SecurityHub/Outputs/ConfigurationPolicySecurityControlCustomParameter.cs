// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.SecurityHub.Outputs
{

    /// <summary>
    /// An object of security control and control parameter value that are included in a configuration policy.
    /// </summary>
    [OutputType]
    public sealed class ConfigurationPolicySecurityControlCustomParameter
    {
        /// <summary>
        /// An object that specifies parameter values for a control in a configuration policy.
        /// </summary>
        public readonly ImmutableDictionary<string, Outputs.ConfigurationPolicyParameterConfiguration>? Parameters;
        /// <summary>
        /// The ID of the security control.
        /// </summary>
        public readonly string? SecurityControlId;

        [OutputConstructor]
        private ConfigurationPolicySecurityControlCustomParameter(
            ImmutableDictionary<string, Outputs.ConfigurationPolicyParameterConfiguration>? parameters,

            string? securityControlId)
        {
            Parameters = parameters;
            SecurityControlId = securityControlId;
        }
    }
}