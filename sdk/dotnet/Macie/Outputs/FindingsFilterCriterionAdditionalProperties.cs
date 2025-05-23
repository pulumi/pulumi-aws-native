// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Macie.Outputs
{

    [OutputType]
    public sealed class FindingsFilterCriterionAdditionalProperties
    {
        /// <summary>
        /// The value for the specified property matches (equals) the specified value. If you specify multiple values, Amazon Macie uses OR logic to join the values.
        /// </summary>
        public readonly ImmutableArray<string> Eq;
        /// <summary>
        /// The value for the specified property is greater than the specified value.
        /// </summary>
        public readonly int? Gt;
        /// <summary>
        /// The value for the specified property is greater than or equal to the specified value.
        /// </summary>
        public readonly int? Gte;
        /// <summary>
        /// The value for the specified property is less than the specified value.
        /// </summary>
        public readonly int? Lt;
        /// <summary>
        /// The value for the specified property is less than or equal to the specified value.
        /// </summary>
        public readonly int? Lte;
        /// <summary>
        /// The value for the specified property doesn't match (doesn't equal) the specified value. If you specify multiple values, Amazon Macie uses OR logic to join the values.
        /// </summary>
        public readonly ImmutableArray<string> Neq;

        [OutputConstructor]
        private FindingsFilterCriterionAdditionalProperties(
            ImmutableArray<string> eq,

            int? gt,

            int? gte,

            int? lt,

            int? lte,

            ImmutableArray<string> neq)
        {
            Eq = eq;
            Gt = gt;
            Gte = gte;
            Lt = lt;
            Lte = lte;
            Neq = neq;
        }
    }
}
