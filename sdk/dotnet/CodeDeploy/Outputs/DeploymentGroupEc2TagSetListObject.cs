// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.CodeDeploy.Outputs
{

    [OutputType]
    public sealed class DeploymentGroupEc2TagSetListObject
    {
        public readonly ImmutableArray<Outputs.DeploymentGroupEc2TagFilter> Ec2TagGroup;

        [OutputConstructor]
        private DeploymentGroupEc2TagSetListObject(ImmutableArray<Outputs.DeploymentGroupEc2TagFilter> ec2TagGroup)
        {
            Ec2TagGroup = ec2TagGroup;
        }
    }
}