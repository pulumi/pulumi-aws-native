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
    public sealed class JobDefinitionEksContainerVolumeMount
    {
        public readonly string? MountPath;
        public readonly string? Name;
        public readonly bool? ReadOnly;

        [OutputConstructor]
        private JobDefinitionEksContainerVolumeMount(
            string? mountPath,

            string? name,

            bool? readOnly)
        {
            MountPath = mountPath;
            Name = name;
            ReadOnly = readOnly;
        }
    }
}