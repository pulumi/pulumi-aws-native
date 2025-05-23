// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Msk.Outputs
{

    [OutputType]
    public sealed class ConfigurationLatestRevision
    {
        /// <summary>
        /// The time when the configuration revision was created.
        /// </summary>
        public readonly string? CreationTime;
        /// <summary>
        /// The description of the configuration revision.
        /// </summary>
        public readonly string? Description;
        /// <summary>
        /// The revision number.
        /// </summary>
        public readonly int? Revision;

        [OutputConstructor]
        private ConfigurationLatestRevision(
            string? creationTime,

            string? description,

            int? revision)
        {
            CreationTime = creationTime;
            Description = description;
            Revision = revision;
        }
    }
}
