// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.QuickSight.Inputs
{

    public sealed class DashboardColumnIdentifierArgs : global::Pulumi.ResourceArgs
    {
        [Input("columnName", required: true)]
        public Input<string> ColumnName { get; set; } = null!;

        [Input("dataSetIdentifier", required: true)]
        public Input<string> DataSetIdentifier { get; set; } = null!;

        public DashboardColumnIdentifierArgs()
        {
        }
        public static new DashboardColumnIdentifierArgs Empty => new DashboardColumnIdentifierArgs();
    }
}