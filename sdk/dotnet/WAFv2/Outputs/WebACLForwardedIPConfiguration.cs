// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.WAFv2.Outputs
{

    [OutputType]
    public sealed class WebACLForwardedIPConfiguration
    {
        public readonly Pulumi.AwsNative.WAFv2.WebACLForwardedIPConfigurationFallbackBehavior FallbackBehavior;
        public readonly string HeaderName;

        [OutputConstructor]
        private WebACLForwardedIPConfiguration(
            Pulumi.AwsNative.WAFv2.WebACLForwardedIPConfigurationFallbackBehavior fallbackBehavior,

            string headerName)
        {
            FallbackBehavior = fallbackBehavior;
            HeaderName = headerName;
        }
    }
}