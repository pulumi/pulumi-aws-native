// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Shield.Outputs
{

    [OutputType]
    public sealed class ProtectionApplicationLayerAutomaticResponseConfigurationAction1Properties
    {
        /// <summary>
        /// Specifies that Shield Advanced should configure its AWS WAF rules with the AWS WAF `Block` action.
        /// You must specify exactly one action, either `Block` or `Count`.
        /// </summary>
        public readonly object? Block;

        [OutputConstructor]
        private ProtectionApplicationLayerAutomaticResponseConfigurationAction1Properties(object? block)
        {
            Block = block;
        }
    }
}