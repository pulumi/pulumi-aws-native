// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Connect.Outputs
{

    /// <summary>
    /// The definition for create case action.
    /// </summary>
    [OutputType]
    public sealed class RuleCreateCaseAction
    {
        public readonly ImmutableArray<Outputs.RuleField> Fields;
        /// <summary>
        /// The Id of template.
        /// </summary>
        public readonly string TemplateId;

        [OutputConstructor]
        private RuleCreateCaseAction(
            ImmutableArray<Outputs.RuleField> fields,

            string templateId)
        {
            Fields = fields;
            TemplateId = templateId;
        }
    }
}