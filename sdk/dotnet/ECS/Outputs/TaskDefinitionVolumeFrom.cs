// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.ECS.Outputs
{

    [OutputType]
    public sealed class TaskDefinitionVolumeFrom
    {
        public readonly bool? ReadOnly;
        public readonly string? SourceContainer;

        [OutputConstructor]
        private TaskDefinitionVolumeFrom(
            bool? readOnly,

            string? sourceContainer)
        {
            ReadOnly = readOnly;
            SourceContainer = sourceContainer;
        }
    }
}