// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.ApplicationSignals.Inputs
{

    /// <summary>
    /// This object defines how often to repeat a time exclusion window.
    /// </summary>
    public sealed class ServiceLevelObjectiveRecurrenceRuleArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// A cron or rate expression denoting how often to repeat this exclusion window.
        /// </summary>
        [Input("expression", required: true)]
        public Input<string> Expression { get; set; } = null!;

        public ServiceLevelObjectiveRecurrenceRuleArgs()
        {
        }
        public static new ServiceLevelObjectiveRecurrenceRuleArgs Empty => new ServiceLevelObjectiveRecurrenceRuleArgs();
    }
}
