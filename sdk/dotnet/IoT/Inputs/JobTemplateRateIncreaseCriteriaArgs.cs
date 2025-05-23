// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.IoT.Inputs
{

    public sealed class JobTemplateRateIncreaseCriteriaArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The threshold for number of notified things that will initiate the increase in rate of rollout.
        /// </summary>
        [Input("numberOfNotifiedThings")]
        public Input<int>? NumberOfNotifiedThings { get; set; }

        /// <summary>
        /// The threshold for number of succeeded things that will initiate the increase in rate of rollout.
        /// </summary>
        [Input("numberOfSucceededThings")]
        public Input<int>? NumberOfSucceededThings { get; set; }

        public JobTemplateRateIncreaseCriteriaArgs()
        {
        }
        public static new JobTemplateRateIncreaseCriteriaArgs Empty => new JobTemplateRateIncreaseCriteriaArgs();
    }
}
