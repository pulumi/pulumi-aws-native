// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Pinpoint.Outputs
{

    [OutputType]
    public sealed class InAppTemplateHeaderConfig
    {
        public readonly Pulumi.AwsNative.Pinpoint.InAppTemplateAlignment? Alignment;
        public readonly string? Header;
        public readonly string? TextColor;

        [OutputConstructor]
        private InAppTemplateHeaderConfig(
            Pulumi.AwsNative.Pinpoint.InAppTemplateAlignment? alignment,

            string? header,

            string? textColor)
        {
            Alignment = alignment;
            Header = header;
            TextColor = textColor;
        }
    }
}