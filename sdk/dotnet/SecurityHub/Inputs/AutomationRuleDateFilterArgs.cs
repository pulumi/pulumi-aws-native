// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.SecurityHub.Inputs
{

    /// <summary>
    /// A date filter for querying findings.
    /// </summary>
    public sealed class AutomationRuleDateFilterArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// A date range for the date filter.
        /// </summary>
        [Input("dateRange")]
        public Input<Inputs.AutomationRuleDateRangeArgs>? DateRange { get; set; }

        /// <summary>
        /// A timestamp that provides the end date for the date filter.
        ///  For more information about the validation and formatting of timestamp fields in ASHlong, see [Timestamps](https://docs.aws.amazon.com/securityhub/1.0/APIReference/Welcome.html#timestamps).
        /// </summary>
        [Input("end")]
        public Input<string>? End { get; set; }

        /// <summary>
        /// A timestamp that provides the start date for the date filter.
        ///  For more information about the validation and formatting of timestamp fields in ASHlong, see [Timestamps](https://docs.aws.amazon.com/securityhub/1.0/APIReference/Welcome.html#timestamps).
        /// </summary>
        [Input("start")]
        public Input<string>? Start { get; set; }

        public AutomationRuleDateFilterArgs()
        {
        }
        public static new AutomationRuleDateFilterArgs Empty => new AutomationRuleDateFilterArgs();
    }
}
