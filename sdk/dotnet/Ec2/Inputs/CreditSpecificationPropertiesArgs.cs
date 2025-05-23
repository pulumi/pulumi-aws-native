// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Ec2.Inputs
{

    /// <summary>
    /// The credit option for CPU usage of the burstable performance instance. Valid values are standard and unlimited.
    /// </summary>
    public sealed class CreditSpecificationPropertiesArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The credit option for CPU usage of the instance.
        /// 
        /// Valid values: `standard` | `unlimited`
        /// 
        /// T3 instances with `host` tenancy do not support the `unlimited` CPU credit option.
        /// </summary>
        [Input("cpuCredits")]
        public Input<string>? CpuCredits { get; set; }

        public CreditSpecificationPropertiesArgs()
        {
        }
        public static new CreditSpecificationPropertiesArgs Empty => new CreditSpecificationPropertiesArgs();
    }
}
