// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.QuickSight.Outputs
{

    [OutputType]
    public sealed class TemplateDefaultTextAreaControlOptions
    {
        /// <summary>
        /// The delimiter that is used to separate the lines in text.
        /// </summary>
        public readonly string? Delimiter;
        /// <summary>
        /// The display options of a control.
        /// </summary>
        public readonly Outputs.TemplateTextAreaControlDisplayOptions? DisplayOptions;

        [OutputConstructor]
        private TemplateDefaultTextAreaControlOptions(
            string? delimiter,

            Outputs.TemplateTextAreaControlDisplayOptions? displayOptions)
        {
            Delimiter = delimiter;
            DisplayOptions = displayOptions;
        }
    }
}