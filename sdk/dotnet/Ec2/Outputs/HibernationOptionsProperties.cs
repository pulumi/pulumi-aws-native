// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Ec2.Outputs
{

    /// <summary>
    /// Indicates whether an instance is enabled for hibernation.
    /// </summary>
    [OutputType]
    public sealed class HibernationOptionsProperties
    {
        /// <summary>
        /// If you set this parameter to true, your instance is enabled for hibernation.
        /// </summary>
        public readonly bool? Configured;

        [OutputConstructor]
        private HibernationOptionsProperties(bool? configured)
        {
            Configured = configured;
        }
    }
}