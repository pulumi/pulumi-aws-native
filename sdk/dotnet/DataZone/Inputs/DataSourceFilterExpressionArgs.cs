// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.DataZone.Inputs
{

    /// <summary>
    /// The search filter expression.
    /// </summary>
    public sealed class DataSourceFilterExpressionArgs : global::Pulumi.ResourceArgs
    {
        [Input("expression", required: true)]
        public Input<string> Expression { get; set; } = null!;

        [Input("type", required: true)]
        public Input<Pulumi.AwsNative.DataZone.DataSourceFilterExpressionType> Type { get; set; } = null!;

        public DataSourceFilterExpressionArgs()
        {
        }
        public static new DataSourceFilterExpressionArgs Empty => new DataSourceFilterExpressionArgs();
    }
}