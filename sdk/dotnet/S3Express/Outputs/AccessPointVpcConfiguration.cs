// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.S3Express.Outputs
{

    /// <summary>
    /// The Virtual Private Cloud (VPC) configuration for a bucket access point.
    /// </summary>
    [OutputType]
    public sealed class AccessPointVpcConfiguration
    {
        /// <summary>
        /// If this field is specified, this access point will only allow connections from the specified VPC ID.
        /// </summary>
        public readonly string? VpcId;

        [OutputConstructor]
        private AccessPointVpcConfiguration(string? vpcId)
        {
            VpcId = vpcId;
        }
    }
}
