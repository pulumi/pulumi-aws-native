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
    /// A collection of settings that apply to spaces of Amazon SageMaker Studio. These settings are specified when the Create/Update Domain API is called.
    /// </summary>
    [OutputType]
    public sealed class DomainDefaultSpaceSettings
    {
        /// <summary>
        /// The execution role for the space.
        /// </summary>
        public readonly string ExecutionRole;
        /// <summary>
        /// The Jupyter server's app settings.
        /// </summary>
        public readonly Outputs.DomainJupyterServerAppSettings? JupyterServerAppSettings;
        /// <summary>
        /// The kernel gateway app settings.
        /// </summary>
        public readonly Outputs.DomainKernelGatewayAppSettings? KernelGatewayAppSettings;
        /// <summary>
        /// The security groups for the Amazon Virtual Private Cloud (VPC) that Studio uses for communication.
        /// </summary>
        public readonly ImmutableArray<string> SecurityGroups;

        [OutputConstructor]
        private DomainDefaultSpaceSettings(
            string executionRole,

            Outputs.DomainJupyterServerAppSettings? jupyterServerAppSettings,

            Outputs.DomainKernelGatewayAppSettings? kernelGatewayAppSettings,

            ImmutableArray<string> securityGroups)
        {
            ExecutionRole = executionRole;
            JupyterServerAppSettings = jupyterServerAppSettings;
            KernelGatewayAppSettings = kernelGatewayAppSettings;
            SecurityGroups = securityGroups;
        }
    }
}