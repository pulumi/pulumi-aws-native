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
    /// &lt;p&gt;Metadata that contains a description for a column.&lt;/p&gt;
    /// </summary>
    [OutputType]
    public sealed class DataSetColumnDescription
    {
        /// <summary>
        /// &lt;p&gt;The text of a description for a column.&lt;/p&gt;
        /// </summary>
        public readonly string? Text;

        [OutputConstructor]
        private DataSetColumnDescription(string? text)
        {
            Text = text;
        }
    }
}