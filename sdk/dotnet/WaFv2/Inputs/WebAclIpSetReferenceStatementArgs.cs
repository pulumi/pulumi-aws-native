// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.WaFv2.Inputs
{

    public sealed class WebAclIpSetReferenceStatementArgs : global::Pulumi.ResourceArgs
    {
        [Input("arn", required: true)]
        public Input<string> Arn { get; set; } = null!;

        [Input("ipSetForwardedIpConfig")]
        public Input<Inputs.WebAclIpSetForwardedIpConfigurationArgs>? IpSetForwardedIpConfig { get; set; }

        public WebAclIpSetReferenceStatementArgs()
        {
        }
        public static new WebAclIpSetReferenceStatementArgs Empty => new WebAclIpSetReferenceStatementArgs();
    }
}