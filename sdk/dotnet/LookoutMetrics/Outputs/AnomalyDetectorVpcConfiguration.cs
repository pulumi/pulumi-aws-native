// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.LookoutMetrics.Outputs
{

    [OutputType]
    public sealed class AnomalyDetectorVpcConfiguration
    {
        public readonly ImmutableArray<string> SecurityGroupIdList;
        public readonly ImmutableArray<string> SubnetIdList;

        [OutputConstructor]
        private AnomalyDetectorVpcConfiguration(
            ImmutableArray<string> securityGroupIdList,

            ImmutableArray<string> subnetIdList)
        {
            SecurityGroupIdList = securityGroupIdList;
            SubnetIdList = subnetIdList;
        }
    }
}