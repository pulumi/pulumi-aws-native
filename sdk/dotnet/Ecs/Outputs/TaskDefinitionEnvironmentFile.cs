// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Ecs.Outputs
{

    [OutputType]
    public sealed class TaskDefinitionEnvironmentFile
    {
        public readonly string? Type;
        public readonly string? Value;

        [OutputConstructor]
        private TaskDefinitionEnvironmentFile(
            string? type,

            string? value)
        {
            Type = type;
            Value = value;
        }
    }
}