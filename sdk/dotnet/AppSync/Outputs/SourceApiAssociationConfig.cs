// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.AppSync.Outputs
{

    [OutputType]
    public sealed class SourceApiAssociationConfig
    {
        /// <summary>
        /// Configuration of the merged behavior for the association. For example when it could be auto or has to be manual.
        /// </summary>
        public readonly Pulumi.AwsNative.AppSync.SourceApiAssociationConfigMergeType? MergeType;

        [OutputConstructor]
        private SourceApiAssociationConfig(Pulumi.AwsNative.AppSync.SourceApiAssociationConfigMergeType? mergeType)
        {
            MergeType = mergeType;
        }
    }
}
