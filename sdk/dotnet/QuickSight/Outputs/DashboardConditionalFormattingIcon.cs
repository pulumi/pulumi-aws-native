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
    public sealed class DashboardConditionalFormattingIcon
    {
        public readonly Outputs.DashboardConditionalFormattingCustomIconCondition? CustomCondition;
        public readonly Outputs.DashboardConditionalFormattingIconSet? IconSet;

        [OutputConstructor]
        private DashboardConditionalFormattingIcon(
            Outputs.DashboardConditionalFormattingCustomIconCondition? customCondition,

            Outputs.DashboardConditionalFormattingIconSet? iconSet)
        {
            CustomCondition = customCondition;
            IconSet = iconSet;
        }
    }
}