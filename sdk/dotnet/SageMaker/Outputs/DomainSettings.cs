// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.SageMaker.Outputs
{

    /// <summary>
    /// A collection of Domain settings.
    /// </summary>
    [OutputType]
    public sealed class DomainSettings
    {
        /// <summary>
        /// A collection of settings that configure the domain's Docker interaction.
        /// </summary>
        public readonly Outputs.DomainDockerSettings? DockerSettings;
        /// <summary>
        /// A collection of settings that configure the `RStudioServerPro` Domain-level app.
        /// </summary>
        public readonly Outputs.DomainRStudioServerProDomainSettings? RStudioServerProDomainSettings;
        /// <summary>
        /// The security groups for the Amazon Virtual Private Cloud that the Domain uses for communication between Domain-level apps and user apps.
        /// </summary>
        public readonly ImmutableArray<string> SecurityGroupIds;

        [OutputConstructor]
        private DomainSettings(
            Outputs.DomainDockerSettings? dockerSettings,

            Outputs.DomainRStudioServerProDomainSettings? rStudioServerProDomainSettings,

            ImmutableArray<string> securityGroupIds)
        {
            DockerSettings = dockerSettings;
            RStudioServerProDomainSettings = rStudioServerProDomainSettings;
            SecurityGroupIds = securityGroupIds;
        }
    }
}
