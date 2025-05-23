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
    /// &lt;p&gt;A data transformation on a logical table. This is a variant type structure. For this
    ///             structure to be valid, only one of the attributes can be non-null.&lt;/p&gt;
    /// </summary>
    [OutputType]
    public sealed class DataSetTransformOperation
    {
        /// <summary>
        /// A transform operation that casts a column to a different type.
        /// </summary>
        public readonly Outputs.DataSetCastColumnTypeOperation? CastColumnTypeOperation;
        /// <summary>
        /// An operation that creates calculated columns. Columns created in one such operation form a lexical closure.
        /// </summary>
        public readonly Outputs.DataSetCreateColumnsOperation? CreateColumnsOperation;
        /// <summary>
        /// An operation that filters rows based on some condition.
        /// </summary>
        public readonly Outputs.DataSetFilterOperation? FilterOperation;
        /// <summary>
        /// A transform operation that overrides the dataset parameter values that are defined in another dataset.
        /// </summary>
        public readonly Outputs.DataSetOverrideDatasetParameterOperation? OverrideDatasetParameterOperation;
        /// <summary>
        /// An operation that projects columns. Operations that come after a projection can only refer to projected columns.
        /// </summary>
        public readonly Outputs.DataSetProjectOperation? ProjectOperation;
        /// <summary>
        /// An operation that renames a column.
        /// </summary>
        public readonly Outputs.DataSetRenameColumnOperation? RenameColumnOperation;
        /// <summary>
        /// An operation that tags a column with additional information.
        /// </summary>
        public readonly Outputs.DataSetTagColumnOperation? TagColumnOperation;
        public readonly Outputs.DataSetUntagColumnOperation? UntagColumnOperation;

        [OutputConstructor]
        private DataSetTransformOperation(
            Outputs.DataSetCastColumnTypeOperation? castColumnTypeOperation,

            Outputs.DataSetCreateColumnsOperation? createColumnsOperation,

            Outputs.DataSetFilterOperation? filterOperation,

            Outputs.DataSetOverrideDatasetParameterOperation? overrideDatasetParameterOperation,

            Outputs.DataSetProjectOperation? projectOperation,

            Outputs.DataSetRenameColumnOperation? renameColumnOperation,

            Outputs.DataSetTagColumnOperation? tagColumnOperation,

            Outputs.DataSetUntagColumnOperation? untagColumnOperation)
        {
            CastColumnTypeOperation = castColumnTypeOperation;
            CreateColumnsOperation = createColumnsOperation;
            FilterOperation = filterOperation;
            OverrideDatasetParameterOperation = overrideDatasetParameterOperation;
            ProjectOperation = projectOperation;
            RenameColumnOperation = renameColumnOperation;
            TagColumnOperation = tagColumnOperation;
            UntagColumnOperation = untagColumnOperation;
        }
    }
}
