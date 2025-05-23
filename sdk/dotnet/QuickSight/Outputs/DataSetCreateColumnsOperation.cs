// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.QuickSight.Outputs
{

    /// <summary>
    /// &lt;p&gt;A transform operation that creates calculated columns. Columns created in one such
    ///             operation form a lexical closure.&lt;/p&gt;
    /// </summary>
    [OutputType]
    public sealed class DataSetCreateColumnsOperation
    {
        /// <summary>
        /// &lt;p&gt;Calculated columns to create.&lt;/p&gt;
        /// </summary>
        public readonly ImmutableArray<Outputs.DataSetCalculatedColumn> Columns;

        [OutputConstructor]
        private DataSetCreateColumnsOperation(ImmutableArray<Outputs.DataSetCalculatedColumn> columns)
        {
            Columns = columns;
        }
    }
}
