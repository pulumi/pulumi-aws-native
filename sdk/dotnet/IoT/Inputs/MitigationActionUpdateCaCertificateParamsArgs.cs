// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.IoT.Inputs
{

    /// <summary>
    /// Parameters to define a mitigation action that changes the state of the CA certificate to inactive.
    /// </summary>
    public sealed class MitigationActionUpdateCaCertificateParamsArgs : global::Pulumi.ResourceArgs
    {
        [Input("action", required: true)]
        public Input<Pulumi.AwsNative.IoT.MitigationActionUpdateCaCertificateParamsAction> Action { get; set; } = null!;

        public MitigationActionUpdateCaCertificateParamsArgs()
        {
        }
        public static new MitigationActionUpdateCaCertificateParamsArgs Empty => new MitigationActionUpdateCaCertificateParamsArgs();
    }
}