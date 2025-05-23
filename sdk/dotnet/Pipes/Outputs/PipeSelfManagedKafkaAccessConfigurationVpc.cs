// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Pipes.Outputs
{

    [OutputType]
    public sealed class PipeSelfManagedKafkaAccessConfigurationVpc
    {
        /// <summary>
        /// List of SecurityGroupId.
        /// </summary>
        public readonly ImmutableArray<string> SecurityGroup;
        /// <summary>
        /// List of SubnetId.
        /// </summary>
        public readonly ImmutableArray<string> Subnets;

        [OutputConstructor]
        private PipeSelfManagedKafkaAccessConfigurationVpc(
            ImmutableArray<string> securityGroup,

            ImmutableArray<string> subnets)
        {
            SecurityGroup = securityGroup;
            Subnets = subnets;
        }
    }
}
