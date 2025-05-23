// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.DataBrew.Outputs
{

    /// <summary>
    /// Configuration to attach Rulesets to the job
    /// </summary>
    [OutputType]
    public sealed class JobValidationConfiguration
    {
        /// <summary>
        /// Arn of the Ruleset
        /// </summary>
        public readonly string RulesetArn;
        /// <summary>
        /// Mode of data quality validation. Default mode is "CHECK_ALL" which verifies all rules defined in the selected ruleset.
        /// </summary>
        public readonly Pulumi.AwsNative.DataBrew.JobValidationMode? ValidationMode;

        [OutputConstructor]
        private JobValidationConfiguration(
            string rulesetArn,

            Pulumi.AwsNative.DataBrew.JobValidationMode? validationMode)
        {
            RulesetArn = rulesetArn;
            ValidationMode = validationMode;
        }
    }
}
