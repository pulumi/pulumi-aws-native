// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.WaFv2.Inputs
{

    /// <summary>
    /// Configures the inspection of login requests
    /// </summary>
    public sealed class WebAclRequestInspectionArgs : global::Pulumi.ResourceArgs
    {
        [Input("passwordField", required: true)]
        public Input<Inputs.WebAclFieldIdentifierArgs> PasswordField { get; set; } = null!;

        [Input("payloadType", required: true)]
        public Input<Pulumi.AwsNative.WaFv2.WebAclRequestInspectionPayloadType> PayloadType { get; set; } = null!;

        [Input("usernameField", required: true)]
        public Input<Inputs.WebAclFieldIdentifierArgs> UsernameField { get; set; } = null!;

        public WebAclRequestInspectionArgs()
        {
        }
        public static new WebAclRequestInspectionArgs Empty => new WebAclRequestInspectionArgs();
    }
}