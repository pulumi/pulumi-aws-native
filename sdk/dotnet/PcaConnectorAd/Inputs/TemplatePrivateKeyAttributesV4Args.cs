// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.PcaConnectorAd.Inputs
{

    public sealed class TemplatePrivateKeyAttributesV4Args : global::Pulumi.ResourceArgs
    {
        [Input("algorithm")]
        public Input<Pulumi.AwsNative.PcaConnectorAd.TemplatePrivateKeyAlgorithm>? Algorithm { get; set; }

        [Input("cryptoProviders")]
        private InputList<string>? _cryptoProviders;
        public InputList<string> CryptoProviders
        {
            get => _cryptoProviders ?? (_cryptoProviders = new InputList<string>());
            set => _cryptoProviders = value;
        }

        [Input("keySpec", required: true)]
        public Input<Pulumi.AwsNative.PcaConnectorAd.TemplateKeySpec> KeySpec { get; set; } = null!;

        [Input("keyUsageProperty")]
        public InputUnion<Inputs.TemplateKeyUsageProperty0PropertiesArgs, Inputs.TemplateKeyUsageProperty1PropertiesArgs>? KeyUsageProperty { get; set; }

        [Input("minimalKeyLength", required: true)]
        public Input<double> MinimalKeyLength { get; set; } = null!;

        public TemplatePrivateKeyAttributesV4Args()
        {
        }
        public static new TemplatePrivateKeyAttributesV4Args Empty => new TemplatePrivateKeyAttributesV4Args();
    }
}