// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Connect.Outputs
{

    /// <summary>
    /// An item at the root level. All items must be sections.
    /// </summary>
    [OutputType]
    public sealed class EvaluationFormBaseItem
    {
        /// <summary>
        /// A subsection or inner section of an item.
        /// </summary>
        public readonly Outputs.EvaluationFormSection Section;

        [OutputConstructor]
        private EvaluationFormBaseItem(Outputs.EvaluationFormSection section)
        {
            Section = section;
        }
    }
}
