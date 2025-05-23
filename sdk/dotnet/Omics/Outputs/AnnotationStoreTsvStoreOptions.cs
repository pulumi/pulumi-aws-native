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
    public sealed class AnnotationStoreTsvStoreOptions
    {
        public readonly Pulumi.AwsNative.Omics.AnnotationStoreAnnotationType? AnnotationType;
        public readonly ImmutableDictionary<string, string>? FormatToHeader;
        public readonly ImmutableArray<ImmutableDictionary<string, Pulumi.AwsNative.Omics.AnnotationStoreSchemaValueType>> Schema;

        [OutputConstructor]
        private AnnotationStoreTsvStoreOptions(
            Pulumi.AwsNative.Omics.AnnotationStoreAnnotationType? annotationType,

            ImmutableDictionary<string, string>? formatToHeader,

            ImmutableArray<ImmutableDictionary<string, Pulumi.AwsNative.Omics.AnnotationStoreSchemaValueType>> schema)
        {
            AnnotationType = annotationType;
            FormatToHeader = formatToHeader;
            Schema = schema;
        }
    }
}
