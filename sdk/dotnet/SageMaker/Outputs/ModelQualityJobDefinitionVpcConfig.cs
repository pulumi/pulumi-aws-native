// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.SageMaker.Outputs
{

    /// <summary>
    /// Specifies a VPC that your training jobs and hosted models have access to. Control access to and from your training and model containers by configuring the VPC.
    /// </summary>
    [OutputType]
    public sealed class ModelQualityJobDefinitionVpcConfig
    {
        /// <summary>
        /// The VPC security group IDs, in the form sg-xxxxxxxx. Specify the security groups for the VPC that is specified in the Subnets field.
        /// </summary>
        public readonly ImmutableArray<string> SecurityGroupIds;
        /// <summary>
        /// The ID of the subnets in the VPC to which you want to connect to your monitoring jobs.
        /// </summary>
        public readonly ImmutableArray<string> Subnets;

        [OutputConstructor]
        private ModelQualityJobDefinitionVpcConfig(
            ImmutableArray<string> securityGroupIds,

            ImmutableArray<string> subnets)
        {
            SecurityGroupIds = securityGroupIds;
            Subnets = subnets;
        }
    }
}