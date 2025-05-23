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
    public sealed class TemplateFontWeight
    {
        /// <summary>
        /// The lexical name for the level of boldness of the text display.
        /// </summary>
        public readonly Pulumi.AwsNative.QuickSight.TemplateFontWeightName? Name;

        [OutputConstructor]
        private TemplateFontWeight(Pulumi.AwsNative.QuickSight.TemplateFontWeightName? name)
        {
            Name = name;
        }
    }
}
