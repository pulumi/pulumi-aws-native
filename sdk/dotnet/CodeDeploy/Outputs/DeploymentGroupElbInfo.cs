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
    public sealed class DeploymentGroupElbInfo
    {
        public readonly string? Name;

        [OutputConstructor]
        private DeploymentGroupElbInfo(string? name)
        {
            Name = name;
        }
    }
}