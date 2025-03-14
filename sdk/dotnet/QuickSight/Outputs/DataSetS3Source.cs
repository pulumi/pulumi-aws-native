// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.QuickSight.Outputs
{

    /// <summary>
    /// &lt;p&gt;A physical table type for an S3 data source.&lt;/p&gt;
    /// </summary>
    [OutputType]
    public sealed class DataSetS3Source
    {
        /// <summary>
        /// &lt;p&gt;The Amazon Resource Name (ARN) for the data source.&lt;/p&gt;
        /// </summary>
        public readonly string DataSourceArn;
        /// <summary>
        /// &lt;p&gt;A physical table type for an S3 data source.&lt;/p&gt;
        ///          &lt;note&gt;
        ///             &lt;p&gt;For files that aren't JSON, only &lt;code&gt;STRING&lt;/code&gt; data types are supported in input columns.&lt;/p&gt;
        ///          &lt;/note&gt;
        /// </summary>
        public readonly ImmutableArray<Outputs.DataSetInputColumn> InputColumns;
        /// <summary>
        /// Information about the format for the S3 source file or files.
        /// </summary>
        public readonly object? UploadSettings;

        [OutputConstructor]
        private DataSetS3Source(
            string dataSourceArn,

            ImmutableArray<Outputs.DataSetInputColumn> inputColumns,

            object? uploadSettings)
        {
            DataSourceArn = dataSourceArn;
            InputColumns = inputColumns;
            UploadSettings = uploadSettings;
        }
    }
}
