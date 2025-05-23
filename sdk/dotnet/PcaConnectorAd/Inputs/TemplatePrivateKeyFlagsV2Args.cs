// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.PcaConnectorAd.Inputs
{

    public sealed class TemplatePrivateKeyFlagsV2Args : global::Pulumi.ResourceArgs
    {
        [Input("clientVersion", required: true)]
        public Input<Pulumi.AwsNative.PcaConnectorAd.TemplateClientCompatibilityV2> ClientVersion { get; set; } = null!;

        [Input("exportableKey")]
        public Input<bool>? ExportableKey { get; set; }

        [Input("strongKeyProtectionRequired")]
        public Input<bool>? StrongKeyProtectionRequired { get; set; }

        public TemplatePrivateKeyFlagsV2Args()
        {
        }
        public static new TemplatePrivateKeyFlagsV2Args Empty => new TemplatePrivateKeyFlagsV2Args();
    }
}
