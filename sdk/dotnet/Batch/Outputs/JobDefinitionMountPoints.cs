// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Batch.Outputs
{

    [OutputType]
    public sealed class JobDefinitionMountPoints
    {
        public readonly string? ContainerPath;
        public readonly bool? ReadOnly;
        public readonly string? SourceVolume;

        [OutputConstructor]
        private JobDefinitionMountPoints(
            string? containerPath,

            bool? readOnly,

            string? sourceVolume)
        {
            ContainerPath = containerPath;
            ReadOnly = readOnly;
            SourceVolume = sourceVolume;
        }
    }
}