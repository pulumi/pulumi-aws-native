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
    /// &lt;p&gt;Permission for the resource.&lt;/p&gt;
    /// </summary>
    public sealed class DataSetRowLevelPermissionTagRuleArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// &lt;p&gt;The column name that a tag key is assigned to.&lt;/p&gt;
        /// </summary>
        [Input("columnName", required: true)]
        public Input<string> ColumnName { get; set; } = null!;

        /// <summary>
        /// &lt;p&gt;A string that you want to use to filter by all the values in a column in the dataset and don’t want to list the values one by one. For example, you can use an asterisk as your match all value.&lt;/p&gt;
        /// </summary>
        [Input("matchAllValue")]
        public Input<string>? MatchAllValue { get; set; }

        /// <summary>
        /// &lt;p&gt;The unique key for a tag.&lt;/p&gt;
        /// </summary>
        [Input("tagKey", required: true)]
        public Input<string> TagKey { get; set; } = null!;

        /// <summary>
        /// &lt;p&gt;A string that you want to use to delimit the values when you pass the values at run time. For example, you can delimit the values with a comma.&lt;/p&gt;
        /// </summary>
        [Input("tagMultiValueDelimiter")]
        public Input<string>? TagMultiValueDelimiter { get; set; }

        public DataSetRowLevelPermissionTagRuleArgs()
        {
        }
        public static new DataSetRowLevelPermissionTagRuleArgs Empty => new DataSetRowLevelPermissionTagRuleArgs();
    }
}