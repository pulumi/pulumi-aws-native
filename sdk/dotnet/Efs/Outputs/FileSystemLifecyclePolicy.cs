// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Efs.Outputs
{

    [OutputType]
    public sealed class FileSystemLifecyclePolicy
    {
        public readonly string? TransitionToArchive;
        public readonly string? TransitionToIa;
        public readonly string? TransitionToPrimaryStorageClass;

        [OutputConstructor]
        private FileSystemLifecyclePolicy(
            string? transitionToArchive,

            string? transitionToIa,

            string? transitionToPrimaryStorageClass)
        {
            TransitionToArchive = transitionToArchive;
            TransitionToIa = transitionToIa;
            TransitionToPrimaryStorageClass = transitionToPrimaryStorageClass;
        }
    }
}