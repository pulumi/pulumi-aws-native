// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.AppRunner.Outputs
{

    /// <summary>
    /// Instance Configuration
    /// </summary>
    [OutputType]
    public sealed class ServiceInstanceConfiguration
    {
        /// <summary>
        /// CPU
        /// </summary>
        public readonly string? Cpu;
        /// <summary>
        /// Instance Role Arn
        /// </summary>
        public readonly string? InstanceRoleArn;
        /// <summary>
        /// Memory
        /// </summary>
        public readonly string? Memory;

        [OutputConstructor]
        private ServiceInstanceConfiguration(
            string? cpu,

            string? instanceRoleArn,

            string? memory)
        {
            Cpu = cpu;
            InstanceRoleArn = instanceRoleArn;
            Memory = memory;
        }
    }
}
