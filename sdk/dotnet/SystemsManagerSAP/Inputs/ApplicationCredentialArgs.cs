// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.SystemsManagerSAP.Inputs
{

    public sealed class ApplicationCredentialArgs : global::Pulumi.ResourceArgs
    {
        [Input("credentialType")]
        public Input<Pulumi.AwsNative.SystemsManagerSAP.ApplicationCredentialCredentialType>? CredentialType { get; set; }

        [Input("databaseName")]
        public Input<string>? DatabaseName { get; set; }

        [Input("secretId")]
        public Input<string>? SecretId { get; set; }

        public ApplicationCredentialArgs()
        {
        }
        public static new ApplicationCredentialArgs Empty => new ApplicationCredentialArgs();
    }
}