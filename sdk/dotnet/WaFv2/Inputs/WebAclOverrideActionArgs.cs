// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.WaFv2.Inputs
{

    /// <summary>
    /// Override a RuleGroup or ManagedRuleGroup behavior. This can only be applied to Rule that has RuleGroupReferenceStatement or ManagedRuleGroupReferenceStatement.
    /// </summary>
    public sealed class WebAclOverrideActionArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Count traffic towards application.
        /// </summary>
        [Input("count")]
        public Input<object>? Count { get; set; }

        /// <summary>
        /// Keep the RuleGroup or ManagedRuleGroup behavior as is.
        /// </summary>
        [Input("none")]
        public Input<object>? None { get; set; }

        public WebAclOverrideActionArgs()
        {
        }
        public static new WebAclOverrideActionArgs Empty => new WebAclOverrideActionArgs();
    }
}
