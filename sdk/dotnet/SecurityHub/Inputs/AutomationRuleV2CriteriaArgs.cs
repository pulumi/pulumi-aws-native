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
    /// Defines the parameters and conditions used to evaluate and filter security findings
    /// </summary>
    public sealed class AutomationRuleV2CriteriaArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The filtering conditions that align with OCSF standards.
        /// </summary>
        [Input("ocsfFindingCriteria")]
        public Input<Inputs.AutomationRuleV2OcsfFindingFiltersArgs>? OcsfFindingCriteria { get; set; }

        public AutomationRuleV2CriteriaArgs()
        {
        }
        public static new AutomationRuleV2CriteriaArgs Empty => new AutomationRuleV2CriteriaArgs();
    }
}
