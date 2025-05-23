// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.GameLift.Outputs
{

    /// <summary>
    /// A set of one or more port numbers that can be opened on the container.
    /// </summary>
    [OutputType]
    public sealed class ContainerGroupDefinitionContainerPortRange
    {
        /// <summary>
        /// A starting value for the range of allowed port numbers.
        /// </summary>
        public readonly int FromPort;
        /// <summary>
        /// Defines the protocol of these ports.
        /// </summary>
        public readonly Pulumi.AwsNative.GameLift.ContainerGroupDefinitionContainerPortRangeProtocol Protocol;
        /// <summary>
        /// An ending value for the range of allowed port numbers. Port numbers are end-inclusive. This value must be equal to or greater than FromPort.
        /// </summary>
        public readonly int ToPort;

        [OutputConstructor]
        private ContainerGroupDefinitionContainerPortRange(
            int fromPort,

            Pulumi.AwsNative.GameLift.ContainerGroupDefinitionContainerPortRangeProtocol protocol,

            int toPort)
        {
            FromPort = fromPort;
            Protocol = protocol;
            ToPort = toPort;
        }
    }
}
