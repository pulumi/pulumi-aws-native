// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.GuardDuty.Inputs
{

    public sealed class DetectorCfnScanEc2InstanceWithFindingsConfigurationArgs : global::Pulumi.ResourceArgs
    {
        [Input("ebsVolumes")]
        public Input<bool>? EbsVolumes { get; set; }

        public DetectorCfnScanEc2InstanceWithFindingsConfigurationArgs()
        {
        }
        public static new DetectorCfnScanEc2InstanceWithFindingsConfigurationArgs Empty => new DetectorCfnScanEc2InstanceWithFindingsConfigurationArgs();
    }
}