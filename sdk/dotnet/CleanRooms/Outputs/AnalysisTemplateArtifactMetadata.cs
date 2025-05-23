// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.CleanRooms.Outputs
{

    [OutputType]
    public sealed class AnalysisTemplateArtifactMetadata
    {
        public readonly ImmutableArray<Outputs.AnalysisTemplateHash> AdditionalArtifactHashes;
        public readonly Outputs.AnalysisTemplateHash EntryPointHash;

        [OutputConstructor]
        private AnalysisTemplateArtifactMetadata(
            ImmutableArray<Outputs.AnalysisTemplateHash> additionalArtifactHashes,

            Outputs.AnalysisTemplateHash entryPointHash)
        {
            AdditionalArtifactHashes = additionalArtifactHashes;
            EntryPointHash = entryPointHash;
        }
    }
}
