// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.QuickSight.Outputs
{

    [OutputType]
    public sealed class DashboardFieldLabelType
    {
        public readonly string? FieldId;
        public readonly Pulumi.AwsNative.QuickSight.DashboardVisibility? Visibility;

        [OutputConstructor]
        private DashboardFieldLabelType(
            string? fieldId,

            Pulumi.AwsNative.QuickSight.DashboardVisibility? visibility)
        {
            FieldId = fieldId;
            Visibility = visibility;
        }
    }
}