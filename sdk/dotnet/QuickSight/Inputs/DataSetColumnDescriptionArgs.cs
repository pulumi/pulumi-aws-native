// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.QuickSight.Inputs
{

    /// <summary>
    /// &lt;p&gt;Metadata that contains a description for a column.&lt;/p&gt;
    /// </summary>
    public sealed class DataSetColumnDescriptionArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// &lt;p&gt;The text of a description for a column.&lt;/p&gt;
        /// </summary>
        [Input("text")]
        public Input<string>? Text { get; set; }

        public DataSetColumnDescriptionArgs()
        {
        }
        public static new DataSetColumnDescriptionArgs Empty => new DataSetColumnDescriptionArgs();
    }
}