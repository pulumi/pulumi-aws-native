// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.EC2.Inputs
{

    public sealed class LaunchTemplateMaintenanceOptionsArgs : global::Pulumi.ResourceArgs
    {
        [Input("autoRecovery")]
        public Input<string>? AutoRecovery { get; set; }

        public LaunchTemplateMaintenanceOptionsArgs()
        {
        }
        public static new LaunchTemplateMaintenanceOptionsArgs Empty => new LaunchTemplateMaintenanceOptionsArgs();
    }
}