// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.CodeArtifact.Outputs
{

    [OutputType]
    public sealed class PackageGroupRestrictionType
    {
        /// <summary>
        /// The repositories to add to the allowed repositories list. The allowed repositories list is used when the `RestrictionMode` is set to `ALLOW_SPECIFIC_REPOSITORIES` .
        /// </summary>
        public readonly ImmutableArray<string> Repositories;
        /// <summary>
        /// The package group origin restriction setting. When the value is `INHERIT` , the value is set to the value of the first parent package group which does not have a value of `INHERIT` .
        /// </summary>
        public readonly Pulumi.AwsNative.CodeArtifact.PackageGroupRestrictionTypeRestrictionMode RestrictionMode;

        [OutputConstructor]
        private PackageGroupRestrictionType(
            ImmutableArray<string> repositories,

            Pulumi.AwsNative.CodeArtifact.PackageGroupRestrictionTypeRestrictionMode restrictionMode)
        {
            Repositories = repositories;
            RestrictionMode = restrictionMode;
        }
    }
}
