// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Ec2.Inputs
{

    /// <summary>
    /// Specifies whether your instance is configured for hibernation.
    /// </summary>
    public sealed class LaunchTemplateHibernationOptionsArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// TIf you set this parameter to true, the instance is enabled for hibernation.
        /// </summary>
        [Input("configured")]
        public Input<bool>? Configured { get; set; }

        public LaunchTemplateHibernationOptionsArgs()
        {
        }
        public static new LaunchTemplateHibernationOptionsArgs Empty => new LaunchTemplateHibernationOptionsArgs();
    }
}