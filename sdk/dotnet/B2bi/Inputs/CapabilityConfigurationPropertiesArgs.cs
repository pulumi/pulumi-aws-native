// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.B2bi.Inputs
{

    public sealed class CapabilityConfigurationPropertiesArgs : global::Pulumi.ResourceArgs
    {
        [Input("edi", required: true)]
        public Input<Inputs.CapabilityEdiConfigurationArgs> Edi { get; set; } = null!;

        public CapabilityConfigurationPropertiesArgs()
        {
        }
        public static new CapabilityConfigurationPropertiesArgs Empty => new CapabilityConfigurationPropertiesArgs();
    }
}
