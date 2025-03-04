// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Msk.Outputs
{

    [OutputType]
    public sealed class ClusterConnectivityInfo
    {
        public readonly Outputs.ClusterPublicAccess? PublicAccess;
        public readonly Outputs.ClusterVpcConnectivity? VpcConnectivity;

        [OutputConstructor]
        private ClusterConnectivityInfo(
            Outputs.ClusterPublicAccess? publicAccess,

            Outputs.ClusterVpcConnectivity? vpcConnectivity)
        {
            PublicAccess = publicAccess;
            VpcConnectivity = vpcConnectivity;
        }
    }
}
