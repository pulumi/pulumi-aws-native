// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Bedrock.Inputs
{

    /// <summary>
    /// Pii entity configuration.
    /// </summary>
    public sealed class GuardrailPiiEntityConfigArgs : global::Pulumi.ResourceArgs
    {
        [Input("action", required: true)]
        public Input<Pulumi.AwsNative.Bedrock.GuardrailSensitiveInformationAction> Action { get; set; } = null!;

        [Input("type", required: true)]
        public Input<Pulumi.AwsNative.Bedrock.GuardrailPiiEntityType> Type { get; set; } = null!;

        public GuardrailPiiEntityConfigArgs()
        {
        }
        public static new GuardrailPiiEntityConfigArgs Empty => new GuardrailPiiEntityConfigArgs();
    }
}