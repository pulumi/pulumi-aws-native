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
    /// &lt;p&gt;A transform operation that filters rows based on a condition.&lt;/p&gt;
    /// </summary>
    public sealed class DataSetFilterOperationArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// &lt;p&gt;An expression that must evaluate to a Boolean value. Rows for which the expression
        ///             evaluates to true are kept in the dataset.&lt;/p&gt;
        /// </summary>
        [Input("conditionExpression", required: true)]
        public Input<string> ConditionExpression { get; set; } = null!;

        public DataSetFilterOperationArgs()
        {
        }
        public static new DataSetFilterOperationArgs Empty => new DataSetFilterOperationArgs();
    }
}