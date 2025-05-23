// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.QuickSight.Outputs
{

    [OutputType]
    public sealed class TemplateSectionLayoutConfiguration
    {
        /// <summary>
        /// The free-form layout configuration of a section.
        /// </summary>
        public readonly Outputs.TemplateFreeFormSectionLayoutConfiguration FreeFormLayout;

        [OutputConstructor]
        private TemplateSectionLayoutConfiguration(Outputs.TemplateFreeFormSectionLayoutConfiguration freeFormLayout)
        {
            FreeFormLayout = freeFormLayout;
        }
    }
}
