// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.WaFv2.Outputs
{

    [OutputType]
    public sealed class RuleGroupForwardedIpConfiguration
    {
        /// <summary>
        /// The match status to assign to the web request if the request doesn't have a valid IP address in the specified position.
        /// 
        /// &gt; If the specified header isn't present in the request, AWS WAF doesn't apply the rule to the web request at all. 
        /// 
        /// You can specify the following fallback behaviors:
        /// 
        /// - `MATCH` - Treat the web request as matching the rule statement. AWS WAF applies the rule action to the request.
        /// - `NO_MATCH` - Treat the web request as not matching the rule statement.
        /// </summary>
        public readonly Pulumi.AwsNative.WaFv2.RuleGroupForwardedIpConfigurationFallbackBehavior FallbackBehavior;
        /// <summary>
        /// The name of the HTTP header to use for the IP address. For example, to use the X-Forwarded-For (XFF) header, set this to `X-Forwarded-For` .
        /// 
        /// &gt; If the specified header isn't present in the request, AWS WAF doesn't apply the rule to the web request at all.
        /// </summary>
        public readonly string HeaderName;

        [OutputConstructor]
        private RuleGroupForwardedIpConfiguration(
            Pulumi.AwsNative.WaFv2.RuleGroupForwardedIpConfigurationFallbackBehavior fallbackBehavior,

            string headerName)
        {
            FallbackBehavior = fallbackBehavior;
            HeaderName = headerName;
        }
    }
}
