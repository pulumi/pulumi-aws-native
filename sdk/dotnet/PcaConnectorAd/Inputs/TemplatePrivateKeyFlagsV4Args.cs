// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.PcaConnectorAd.Inputs
{

    public sealed class TemplatePrivateKeyFlagsV4Args : global::Pulumi.ResourceArgs
    {
        [Input("clientVersion", required: true)]
        public Input<Pulumi.AwsNative.PcaConnectorAd.TemplateClientCompatibilityV4> ClientVersion { get; set; } = null!;

        [Input("exportableKey")]
        public Input<bool>? ExportableKey { get; set; }

        [Input("requireAlternateSignatureAlgorithm")]
        public Input<bool>? RequireAlternateSignatureAlgorithm { get; set; }

        [Input("requireSameKeyRenewal")]
        public Input<bool>? RequireSameKeyRenewal { get; set; }

        [Input("strongKeyProtectionRequired")]
        public Input<bool>? StrongKeyProtectionRequired { get; set; }

        [Input("useLegacyProvider")]
        public Input<bool>? UseLegacyProvider { get; set; }

        public TemplatePrivateKeyFlagsV4Args()
        {
        }
        public static new TemplatePrivateKeyFlagsV4Args Empty => new TemplatePrivateKeyFlagsV4Args();
    }
}
