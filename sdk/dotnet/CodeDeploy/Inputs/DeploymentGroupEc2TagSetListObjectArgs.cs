// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.CodeDeploy.Inputs
{

    public sealed class DeploymentGroupEc2TagSetListObjectArgs : global::Pulumi.ResourceArgs
    {
        [Input("ec2TagGroup")]
        private InputList<Inputs.DeploymentGroupEc2TagFilterArgs>? _ec2TagGroup;
        public InputList<Inputs.DeploymentGroupEc2TagFilterArgs> Ec2TagGroup
        {
            get => _ec2TagGroup ?? (_ec2TagGroup = new InputList<Inputs.DeploymentGroupEc2TagFilterArgs>());
            set => _ec2TagGroup = value;
        }

        public DeploymentGroupEc2TagSetListObjectArgs()
        {
        }
        public static new DeploymentGroupEc2TagSetListObjectArgs Empty => new DeploymentGroupEc2TagSetListObjectArgs();
    }
}