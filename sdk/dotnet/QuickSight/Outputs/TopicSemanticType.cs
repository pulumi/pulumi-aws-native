// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.QuickSight.Outputs
{

    [OutputType]
    public sealed class TopicSemanticType
    {
        /// <summary>
        /// The semantic type falsey cell value.
        /// </summary>
        public readonly string? FalseyCellValue;
        /// <summary>
        /// The other names or aliases for the false cell value.
        /// </summary>
        public readonly ImmutableArray<string> FalseyCellValueSynonyms;
        /// <summary>
        /// The semantic type sub type name.
        /// </summary>
        public readonly string? SubTypeName;
        /// <summary>
        /// The semantic type truthy cell value.
        /// </summary>
        public readonly string? TruthyCellValue;
        /// <summary>
        /// The other names or aliases for the true cell value.
        /// </summary>
        public readonly ImmutableArray<string> TruthyCellValueSynonyms;
        /// <summary>
        /// The semantic type name.
        /// </summary>
        public readonly string? TypeName;
        /// <summary>
        /// The semantic type parameters.
        /// </summary>
        public readonly ImmutableDictionary<string, string>? TypeParameters;

        [OutputConstructor]
        private TopicSemanticType(
            string? falseyCellValue,

            ImmutableArray<string> falseyCellValueSynonyms,

            string? subTypeName,

            string? truthyCellValue,

            ImmutableArray<string> truthyCellValueSynonyms,

            string? typeName,

            ImmutableDictionary<string, string>? typeParameters)
        {
            FalseyCellValue = falseyCellValue;
            FalseyCellValueSynonyms = falseyCellValueSynonyms;
            SubTypeName = subTypeName;
            TruthyCellValue = truthyCellValue;
            TruthyCellValueSynonyms = truthyCellValueSynonyms;
            TypeName = typeName;
            TypeParameters = typeParameters;
        }
    }
}
