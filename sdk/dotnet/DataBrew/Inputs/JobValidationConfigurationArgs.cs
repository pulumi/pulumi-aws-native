// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.DataBrew.Inputs
{

    /// <summary>
    /// Configuration to attach Rulesets to the job
    /// </summary>
    public sealed class JobValidationConfigurationArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Arn of the Ruleset
        /// </summary>
        [Input("rulesetArn", required: true)]
        public Input<string> RulesetArn { get; set; } = null!;

        [Input("validationMode")]
        public Input<Pulumi.AwsNative.DataBrew.JobValidationMode>? ValidationMode { get; set; }

        public JobValidationConfigurationArgs()
        {
        }
        public static new JobValidationConfigurationArgs Empty => new JobValidationConfigurationArgs();
    }
}