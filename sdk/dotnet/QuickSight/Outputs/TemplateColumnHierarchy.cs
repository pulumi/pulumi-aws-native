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
    public sealed class TemplateColumnHierarchy
    {
        public readonly Outputs.TemplateDateTimeHierarchy? DateTimeHierarchy;
        public readonly Outputs.TemplateExplicitHierarchy? ExplicitHierarchy;
        public readonly Outputs.TemplatePredefinedHierarchy? PredefinedHierarchy;

        [OutputConstructor]
        private TemplateColumnHierarchy(
            Outputs.TemplateDateTimeHierarchy? dateTimeHierarchy,

            Outputs.TemplateExplicitHierarchy? explicitHierarchy,

            Outputs.TemplatePredefinedHierarchy? predefinedHierarchy)
        {
            DateTimeHierarchy = dateTimeHierarchy;
            ExplicitHierarchy = explicitHierarchy;
            PredefinedHierarchy = predefinedHierarchy;
        }
    }
}