// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.QuickSight.Inputs
{

    public sealed class DashboardDataPathValueArgs : global::Pulumi.ResourceArgs
    {
        [Input("fieldId", required: true)]
        public Input<string> FieldId { get; set; } = null!;

        [Input("fieldValue", required: true)]
        public Input<string> FieldValue { get; set; } = null!;

        public DashboardDataPathValueArgs()
        {
        }
        public static new DashboardDataPathValueArgs Empty => new DashboardDataPathValueArgs();
    }
}