// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.GreengrassV2.Outputs
{

    [OutputType]
    public sealed class DeploymentComponentConfigurationUpdate
    {
        public readonly string? Merge;
        public readonly ImmutableArray<string> Reset;

        [OutputConstructor]
        private DeploymentComponentConfigurationUpdate(
            string? merge,

            ImmutableArray<string> reset)
        {
            Merge = merge;
            Reset = reset;
        }
    }
}
