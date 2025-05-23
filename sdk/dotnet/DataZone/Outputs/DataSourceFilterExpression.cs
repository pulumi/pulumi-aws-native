// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.DataZone.Outputs
{

    /// <summary>
    /// The search filter expression.
    /// </summary>
    [OutputType]
    public sealed class DataSourceFilterExpression
    {
        public readonly string Expression;
        public readonly Pulumi.AwsNative.DataZone.DataSourceFilterExpressionType Type;

        [OutputConstructor]
        private DataSourceFilterExpression(
            string expression,

            Pulumi.AwsNative.DataZone.DataSourceFilterExpressionType type)
        {
            Expression = expression;
            Type = type;
        }
    }
}
