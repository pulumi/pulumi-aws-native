// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.DataBrew.Outputs
{

    [OutputType]
    public sealed class DatasetFilterExpression
    {
        /// <summary>
        /// Filtering expression for a parameter
        /// </summary>
        public readonly string Expression;
        /// <summary>
        /// The map of substitution variable names to their values used in this filter expression.
        /// </summary>
        public readonly ImmutableArray<Outputs.DatasetFilterValue> ValuesMap;

        [OutputConstructor]
        private DatasetFilterExpression(
            string expression,

            ImmutableArray<Outputs.DatasetFilterValue> valuesMap)
        {
            Expression = expression;
            ValuesMap = valuesMap;
        }
    }
}
