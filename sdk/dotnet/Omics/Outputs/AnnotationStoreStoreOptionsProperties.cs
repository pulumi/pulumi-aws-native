// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Omics.Outputs
{

    [OutputType]
    public sealed class AnnotationStoreStoreOptionsProperties
    {
        public readonly Outputs.AnnotationStoreTsvStoreOptions TsvStoreOptions;

        [OutputConstructor]
        private AnnotationStoreStoreOptionsProperties(Outputs.AnnotationStoreTsvStoreOptions tsvStoreOptions)
        {
            TsvStoreOptions = tsvStoreOptions;
        }
    }
}
