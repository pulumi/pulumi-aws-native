// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.GreengrassV2.Outputs
{

    [OutputType]
    public sealed class DeploymentComponentRunWith
    {
        public readonly string? PosixUser;
        public readonly Outputs.DeploymentSystemResourceLimits? SystemResourceLimits;
        public readonly string? WindowsUser;

        [OutputConstructor]
        private DeploymentComponentRunWith(
            string? posixUser,

            Outputs.DeploymentSystemResourceLimits? systemResourceLimits,

            string? windowsUser)
        {
            PosixUser = posixUser;
            SystemResourceLimits = systemResourceLimits;
            WindowsUser = windowsUser;
        }
    }
}